package log

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ilya1st/rotatewriter"
	"github.com/rs/zerolog"
)

var zlog zerolog.Logger

func init() {
	//TODO get file path form YAML Config

	writer, err := rotatewriter.NewRotateWriter("./test.log", 2)
	if err != nil {
		panic(err)
	}
	sighupChan := make(chan os.Signal, 1)
	signal.Notify(sighupChan, syscall.SIGHUP)
	go func() {
		for {
			_, ok := <-sighupChan
			if !ok {
				return
			}
			fmt.Println("Log rotation")
			writer.Rotate(nil)
		}
	}()
	zlog = zerolog.New(writer).With().Timestamp().Logger()
	zlog.Warn().Msg("fluent logger initialize!")
}

type FluentLogger struct {
	zlog zerolog.Logger
}

type FluentLoggerIF interface {
	Message(message string)
}

func GetFluentLogger() FluentLoggerIF {
	return &FluentLogger{}
}

func (self *FluentLogger) Message(message string) {
	zlog.Info().Msg(message)
}
