package internal

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
)

func initLogger() zerolog.Logger {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}

	logFile, err := os.OpenFile("logisticsroutes.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	fileWriter := zerolog.ConsoleWriter{Out: logFile}

	multi := zerolog.MultiLevelWriter(consoleWriter, fileWriter)
	logger := zerolog.New(multi).
		With().
		Timestamp().
		.Str("format", "15:04:05 Jan 02").
		Caller().
		Logger()

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	logger.Info().Msg("Logger initialized")

	return logger
}
	
func (lw *LoggerWrapper) GetLogger() *zerolog.Logger {
	return &lw.logger
}