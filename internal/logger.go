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

	config, err := loadConfig("../koanf_config.yaml")
	if err != nil {
		panic(err)
	}

	if (config.LogConfig.EnableConsole) {
		consoleWriter = zerolog.ConsoleWriter{Out: os.Stdout}
	}
	if (config.LogConfig.EnableFile) {
		logFile, err := os.OpenFile(config.LogConfig.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
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
		return &log
	}
	return logger
}