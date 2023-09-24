package settings

import (
	"time"

	"github.com/spf13/viper"
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

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

func Setup() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("ini")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Set the values of the AppSetting variable
	AppSetting.JwtSecret = viper.GetString("app.jwt_secret")
	AppSetting.PageSize = viper.GetInt("app.page_size")
	AppSetting.RuntimeRootPath = viper.GetString("app.runtime_root_path")
	AppSetting.LogSavePath = viper.GetString("app.log_save_path")
	AppSetting.LogSaveName = viper.GetString("app.log_save_name")
	AppSetting.LogFileExt = viper.GetString("app.log_file_ext")
	AppSetting.TimeFormat = viper.GetString("app.time_format")

	// Set the values of the ServerSetting variable
	ServerSetting.RunMode = viper.GetString("server.run_mode")
	ServerSetting.HttpPort = viper.GetInt("server.http_port")
	ServerSetting.ReadTimeout = viper.GetDuration("server.read_timeout") * time.Second
	ServerSetting.WriteTimeout = viper.GetDuration("server.write_timeout") * time.Second
}
