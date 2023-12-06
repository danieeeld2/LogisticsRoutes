package internal

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"time"
)

var logger *zerolog.Logger

func initLogger() zerolog.Logger {
	consoleWriter := zerolog.ConsoleWriter{Out: nil}
	fileWriter := zerolog.ConsoleWriter{Out: nil}

	zerolog.TimeFieldFormat = time.RFC3339Nano

	conf, err := loadConfig("../koanf_config.yaml")
	if err != nil {
		panic(err)
	}

	if (conf.Log.Config.Enable.Console) {
		consoleWriter = zerolog.ConsoleWriter{Out: os.Stdout}
	}
	if (conf.Log.Config.Enable.File) {
		logFile, err := os.OpenFile(conf.Log.Config.Logfile.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			defer logFile.Close()
			panic(err)
		}
		fileWriter = zerolog.ConsoleWriter{Out: logFile, NoColor: true}
	}

	multi := zerolog.MultiLevelWriter(consoleWriter, fileWriter)
	logger := zerolog.New(multi).
		With().
		Timestamp().
		Caller().
		Logger()

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	logger.Info().Msg("Logger initialized")

	return logger
}
	
func GetLogger() *zerolog.Logger {
	if logger == nil {
		var log = initLogger()
		logger = &log
	}
	return logger
}