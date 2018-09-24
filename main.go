package main

import (
	"echo-production-example/common"
	"echo-production-example/log"
	"echo-production-example/route"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

const ymlPrefix = "config" //config.yml

var yml common.Yaml

var appLogger *log.AppLogger

func init() {
	appLogger := log.GetApplogger()

	viper.SetConfigName(ymlPrefix) // actually 'fileName + yml'
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		appLogger.Fatal("Error reading config file : ")
	}

	err := viper.Unmarshal(&yml)
	if err != nil {
		appLogger.Fatal("unable to decode into struct : ")
	}

}

func main() {

	appLogger.Info("HTTP Server Initialize...")

	/* sqlx start */
	_, err := common.SqlxInit(yml)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(common.TraceLogInterceptor())

	route.MsAuthorityActivate(e)
	route.MsUserActivate(e)

	e.Static("/", yml.Echo.Static)
	e.Logger.Fatal(e.Start(yml.Echo.Port))
}
