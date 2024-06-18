//go:generate bash -c "mkdir -p codegen && go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0 -generate types,server,spec -package codegen api/zimacube-metrics/openapi.yaml > codegen/zimacube_metrics_api.go"

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/IceWhaleTech/CasaOS-Common/model"
	"github.com/IceWhaleTech/CasaOS-Common/utils/file"
	util_http "github.com/IceWhaleTech/CasaOS-Common/utils/http"
	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/common"
	"github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/config"
	"github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/pkg/sqlite"
	"github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/route"
	"github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/service"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	commit = "private build"
	date   = "private build"

	//go:embed api/index.html
	_docHTML string

	//go:embed api/zimacube-metrics/openapi.yaml
	_docYAML string

	sqliteDB *gorm.DB
)

func main() {
	dbFlag := flag.String("db", "", "db path")
	reportFlag := flag.Bool("r", false, "report")
	versionFlag := flag.Bool("v", false, "version")
	webhookFlag := flag.String("webhook", "", "webhook url")

	flag.Parse()

	if *versionFlag {
		fmt.Printf("v%s\n", common.Version)
		os.Exit(0)
	}

	println("git commit:", commit)
	println("build date:", date)

	if len(*dbFlag) == 0 {
		*dbFlag = config.AppInfo.DBPath
	}

	logger.LogInit(config.AppInfo.LogPath, config.AppInfo.LogSaveName, config.AppInfo.LogFileExt)

	sqliteDB = sqlite.GetDB(*dbFlag)

	service.Initialize(config.CommonInfo.RuntimePath)

	// setup listener
	listener, _err := net.Listen("tcp", net.JoinHostPort("127.0.0.1", "0"))
	if _err != nil { // using `_err` to avoid shadowing the `err` variable
		panic(_err)
	}

	urlFilePath := filepath.Join(config.CommonInfo.RuntimePath, "zimacube-metrics.url")
	if err := file.CreateFileAndWriteContent(urlFilePath, "http://"+listener.Addr().String()); err != nil {
		logger.Error("error when creating address file", zap.Error(err),
			zap.Any("address", listener.Addr().String()),
			zap.Any("filepath", urlFilePath),
		)
	}

	// initialize routers and register at gateway
	apiPaths := []string{
		route.APIPath,
		route.DocPath,
	}

	for _, apiPath := range apiPaths {
		if err := service.MyService.Gateway().CreateRoute(&model.Route{
			Path:   apiPath,
			Target: "http://" + listener.Addr().String(),
		}); err != nil {
			panic(err)
		}
	}

	service.MyService.Metrics().DB = sqliteDB
	go func() {
		time.After(5 * time.Second)
		service.MyService.Metrics().Monitor()
	}()
	if *reportFlag {
		go service.MyService.Metrics().ReportFeiShu(*webhookFlag)
	}

	router := route.GetRouter()
	docRouter := route.GetDocRouter(_docHTML, _docYAML)

	mux := &util_http.HandlerMultiplexer{
		HandlerMap: map[string]http.Handler{
			"v2":  router,
			"doc": docRouter,
		},
	}

	s := &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second, // fix G112: Potential slowloris attack (see https://github.com/securego/gosec)
	}

	_err = s.Serve(listener) // not using http.serve() to fix G114: Use of net/http serve function that has no support for setting timeouts (see
	if _err != nil {
		panic(_err)
	}
}
