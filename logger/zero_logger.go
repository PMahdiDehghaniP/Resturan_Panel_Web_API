package logger

import (
	"log"
	"os"

	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var ZeroLogLevelMap = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
}

type ZeroLogger struct {
	cfg    *config.Config
	logger *zerolog.Logger
}

func (zl *ZeroLogger) GetLogLevel() zerolog.Level {
	level, exists := ZeroLogLevelMap[zl.cfg.Logger.Level]
	if !exists {
		level = zerolog.DebugLevel
	}
	return level
}

func NewZeroLogger(cfg *config.Config) *ZeroLogger {
	logger := &ZeroLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (zl *ZeroLogger) Init() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	level := zl.GetLogLevel()
	zerolog.SetGlobalLevel(level)

	file, err := os.OpenFile(
		zl.cfg.Logger.FilePath,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)
	if err != nil {
		log.Fatalf("cannot open log file: %v", err)
	}
	var logger = zerolog.New(file).With().Timestamp().Str("App Name", "My App").
		Str("Logger Name", "ZeroLog").Logger()

	zl.logger = &logger
}

func (zl *ZeroLogger) Debug(category Category, subCategory SubCategory,
	message string, extra map[ExtraKey]interface{}) {
	zl.logger.Debug().Str("Category", string(category)).
		Str("SubCategory", string(subCategory)).
		Fields(mapToZeroParams(extra)).
		Msg(message)
}

func (zl *ZeroLogger) DebugF(template string, arg ...interface{}) {
	zl.logger.
		Debug().
		Msgf(template, arg...)
}

func (zl *ZeroLogger) Info(category Category, subCategory SubCategory,
	message string, extra map[ExtraKey]interface{}) {
	zl.logger.Info().Str("Category", string(category)).
		Str("SubCategory", string(subCategory)).
		Fields(mapToZeroParams(extra)).
		Msg(message)
}

func (zl *ZeroLogger) InfoF(template string, arg ...interface{}) {
	zl.logger.
		Info().
		Msgf(template, arg...)
}

func (zl *ZeroLogger) Warn(category Category, subCategory SubCategory,
	message string, extra map[ExtraKey]interface{}) {
	zl.logger.Warn().Str("Category", string(category)).
		Str("SubCategory", string(subCategory)).
		Fields(mapToZeroParams(extra)).
		Msg(message)
}

func (zl *ZeroLogger) WarnF(template string, arg ...interface{}) {
	zl.logger.
		Warn().
		Msgf(template, arg...)
}

func (zl *ZeroLogger) Error(category Category, subCategory SubCategory,
	message string, extra map[ExtraKey]interface{}) {
	zl.logger.Error().Str("Category", string(category)).
		Str("SubCategory", string(subCategory)).
		Fields(mapToZeroParams(extra)).
		Msg(message)
}

func (zl *ZeroLogger) ErrorF(template string, arg ...interface{}) {
	zl.logger.
		Error().
		Msgf(template, arg...)
}

func (zl *ZeroLogger) Fatal(category Category, subCategory SubCategory,
	message string, extra map[ExtraKey]interface{}) {
	zl.logger.Fatal().Str("Category", string(category)).
		Str("SubCategory", string(subCategory)).
		Fields(mapToZeroParams(extra)).
		Msg(message)
}

func (zl *ZeroLogger) FatalF(template string, arg ...interface{}) {
	zl.logger.
		Fatal().
		Msgf(template, arg...)
}
