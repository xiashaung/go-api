package lib

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"log"
	"strings"
)

type db struct {
	MasterDsn []string
	SlaveDsn  []string
}

type server struct {
	Host       string
	Port       int
	Model      string
	TimeFormat string
}

type redis struct {
	Host     string
	Port     int
	password string
}

type mpConf struct {
	AppId     string
	AppSecret string
}

var (
	cfg       *ini.File
	Server    server
	Db        db
	RedisConf redis
	MpConf    mpConf
)

func init() {
	var err error
	var config_file = GetCurrentPath() + "/etc/config.ini"
	cfg, err = ini.Load(config_file)
	if err != nil {
		logrus.Info(err)
	}
	loadServer()
	loadDb()
	loadRedis()
	loadMp()
}

func loadServer() {
	sec := GetSection("server")
	Server.Host = sec.Key("host").MustString("0.0.0.0")
	Server.Port = sec.Key("port").MustInt(8901)
	Server.Model = sec.Key("model").MustString("debug")
	Server.TimeFormat = sec.Key("time_format").MustString("2006-01-02 15:04:05")
}

func loadDb() {
	sec := GetSection("database")
	masterDsn := sec.Key("master_dsn").MustString("")
	slaveDsn := sec.Key("slave_dsn").MustString("")
	Db.MasterDsn = strings.Split(masterDsn, ";")
	Db.SlaveDsn = strings.Split(slaveDsn, ";")
}

func loadRedis() {
	sec := GetSection("redis")
	RedisConf.Host = sec.Key("host").MustString("127.0.0.1")
	RedisConf.Port = sec.Key("port").MustInt(6379)
	RedisConf.password = sec.Key("password").MustString("")
}

func loadMp() {
	sec := GetSection("miniprogram")
	MpConf.AppId = sec.Key("app_id").MustString("")
	MpConf.AppSecret = sec.Key("app_secret").MustString("")
}

func loadApp() {
	_, err := cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
}

func GetConfig(section string, key string, def string) string {
	sec := GetSection(section)
	return sec.Key(key).MustString(def)
}

func GetSection(name string) *ini.Section {
	r, _ := cfg.GetSection(name)
	return r
}
