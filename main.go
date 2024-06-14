//go:generate bash -c "mkdir -p codegen && go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0 -generate types,server,spec -package codegen api/zimacube-metrics/openapi.yaml > codegen/zimacube_monitoring_api.go"

package main

import (
	_ "embed"
	"net"
	"net/http"
	"time"

	"github.com/IceWhaleTech/CasaOS-Common/model"
	"github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/config"
	"github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/route"
	"github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/service"
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
	service.Initialize(config.CommonInfo.RuntimePath)

	// setup listener
	listener, _err := net.Listen("tcp", net.JoinHostPort("127.0.0.1", "0"))
	if _err != nil { // using `_err` to avoid shadowing the `err` variable
		panic(_err)
	}

	// initialize routers and register at gateway
	if err := service.Gateway.CreateRoute(&model.Route{
		Path:   route.APIPath,
		Target: "http://" + listener.Addr().String(),
	}); err != nil {
		panic(err)
	}

	s := &http.Server{
		Handler:           route.GetRouter(),
		ReadHeaderTimeout: 5 * time.Second, // fix G112: Potential slowloris attack (see https://github.com/securego/gosec)
	}

	_err = s.Serve(listener) // not using http.serve() to fix G114: Use of net/http serve function that has no support for setting timeouts (see https://github.com/securego/gosec)
	if _err != nil {
		panic(_err)
	}
}
