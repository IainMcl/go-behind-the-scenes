package settings

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
	PageSize  int

	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSettings = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSettings = &Server{}

type Database struct {
	Type         string
	User         string
	Password     string
	Host         string
	Port         string
	DbName       string
	SSLMode      string
	MaxIdleConns int
	MaxOpenConns int
}

var DatabaseSettings = &Database{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("config.ini")
	if err != nil {
		log.Fatal("Fail to parse 'config.ini': %v", err)
	}

	mapTo("app", AppSettings)
	mapTo("server", ServerSettings)
	mapTo("database", DatabaseSettings)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatal("Cfg.MapTo %s err: %v", section, err)
	}
}
