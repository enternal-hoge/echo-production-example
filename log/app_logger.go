package log

import (
	"fmt"
	"os"

	"github.com/francoispqt/onelog"
)

var applog *onelog.Logger

type AppLogger struct {
	applog *onelog.Logger
}

type AppLoggerIF interface {
	Trace(message string)
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
}

func GetApplogger() AppLoggerIF {
	if applog == nil {
		//ToDo : Change Environment Log Level
		//Production : onelog.DEBUG|onelog.WARN|onelog.ERROR|onelog.FATAL,
		applog = onelog.New(os.Stdout, onelog.ALL)
		return &AppLogger{
			applog: applog,
		}
	} else {
		return &AppLogger{}
	}
}

// Trace do'nt use normaly execution.
func (self *AppLogger) Trace(message string) {
	onelog.LevelText(onelog.DEBUG, "trace")
	applog.Debug(message)
}

func (self *AppLogger) Debug(message string) {
	//onelog.LevelText(onelog.DEBUG, "debug")
	//applog.Debug(message)
	fmt.Println(message)
}

func (self *AppLogger) Info(message string) {
	//applog.Info(message)
	fmt.Println(message)
}

func (self *AppLogger) Warn(message string) {
	applog.Warn(message)
}

func (self *AppLogger) Error(message string) {
	applog.Error(message)
}

func (self *AppLogger) Fatal(message string) {
	applog.Fatal(message)
}
