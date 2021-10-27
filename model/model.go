package model

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"go-api/lib"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type ModelTime struct {
	AddTime  sql.NullInt64 `gorm:"column:add_time;type:int;autoUpdateTime" json:"add_time"`
	LastTime sql.NullInt64 `gorm:"column:last_time;type:int;autoCreateTime" json:"last_time"`
}

var (
	DB *gorm.DB
)

func init() {
	master_dsn := lib.Db.MasterDsn[0]
	logrus.Debug(master_dsn)
	master_dsns := lib.Db.MasterDsn
	slave_dsns := lib.Db.SlaveDsn
	var dialectors []gorm.Dialector
	var replicas []gorm.Dialector

	//主从配置支持
	DB, _ = gorm.Open(mysql.Open(master_dsn), &gorm.Config{})
	//关联主库
	for _, val := range master_dsns {
		dialectors = append(dialectors, mysql.Open(val))
	}
	//关联从库
	for _, val := range slave_dsns {
		if val != "" {
			replicas = append(replicas, mysql.Open(val))
		}
	}
	dbResolverCfg := dbresolver.Config{
		Sources:  dialectors,
		Replicas: replicas,
		Policy:   dbresolver.RandomPolicy{}}
	readWritePlugin := dbresolver.Register(dbResolverCfg)
	_ = DB.Use(readWritePlugin)
}

//Table
func Table(value interface{}) *gorm.DB {
	DB.Model(value)
	return DB
}
