package sqlite

import (
	"time"

	"github.com/IceWhaleTech/CasaOS-Common/utils/file"
	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	model2 "github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/service/model"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var gdb *gorm.DB

func GetDB(dbPath string) *gorm.DB {
	if gdb != nil {
		return gdb
	}

	_ = file.IsNotExistMkDir(dbPath)
	db, err := gorm.Open(sqlite.Open(dbPath+"/user.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	c, _ := db.DB()
	c.SetMaxIdleConns(10)
	c.SetMaxOpenConns(1)
	c.SetConnMaxIdleTime(time.Second * 1000)

	gdb = db

	err = db.AutoMigrate(model2.MetricDBModel{})
	if err != nil {
		logger.Error("check or create db error", zap.Any("error", err))
	}

	return db
}
