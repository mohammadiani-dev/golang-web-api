package logging

import (
	"fmt"
	"golang-web-api/config"
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var once sync.Once
var zeroLoggerInstance *zerolog.Logger

type ZeroLogger struct {
	cfg    *config.Logger
	logger *zerolog.Logger
}

var ZeroLogLevelMap = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
}

func (l *ZeroLogger) getLogLevel(level string) zerolog.Level {
	if logLevel, ok := ZeroLogLevelMap[level]; ok {
		return logLevel
	}
	return zerolog.DebugLevel
}

func NewZeroLogger(cfg *config.Logger) *ZeroLogger {
	logger := &ZeroLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *ZeroLogger) Init() {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack


		file, err := os.OpenFile(l.cfg.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(fmt.Sprintf("error opening file: %v", err))
		}

		var logger = zerolog.New(file).
			With().
			Timestamp().
			Str("Appname", "car app").
			Str("LoggerName", "Zerolog").
			Logger()

		zerolog.SetGlobalLevel(l.getLogLevel(l.cfg.Level))

		zeroLoggerInstance = &logger
	})
	l.logger = zeroLoggerInstance
}

func (l *ZeroLogger) Debug(cat Category, sub subCategory, msg string, extera map[ExtraKey]interface{}) {
	l.logger.Debug().
		Str("category", string(cat)).
		Str("subcategory", string(sub)).
		Fields(mapToZeroParams(extera)).
		Msg(msg)
}

func (l *ZeroLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debug().Msgf(template, args...)
}

func (l *ZeroLogger) Info(cat Category, sub subCategory, msg string, extera map[ExtraKey]interface{}) {
	l.logger.Info().
		Str("category", string(cat)).
		Str("subcategory", string(sub)).
		Fields(mapToZeroParams(extera)).
		Msg(msg)
}

func (l *ZeroLogger) Infof(template string, args ...interface{}) {
	l.logger.Info().Msgf(template, args...)
}

func (l *ZeroLogger) Warn(cat Category, sub subCategory, msg string, extera map[ExtraKey]interface{}) {
	l.logger.Warn().
		Str("category", string(cat)).
		Str("subcategory", string(sub)).
		Fields(mapToZeroParams(extera)).
		Msg(msg)
}

func (l *ZeroLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warn().Msgf(template, args...)
}

func (l *ZeroLogger) Error(cat Category, sub subCategory, msg string, extera map[ExtraKey]interface{}) {
	l.logger.Error().
		Str("category", string(cat)).
		Str("subcategory", string(sub)).
		Fields(mapToZeroParams(extera)).
		Msg(msg)
}

func (l *ZeroLogger) Errorf(template string, args ...interface{}) {
	l.logger.Error().Msgf(template, args...)
}

func (l *ZeroLogger) Fatal(cat Category, sub subCategory, msg string, extera map[ExtraKey]interface{}) {
	l.logger.Fatal().
		Str("category", string(cat)).
		Str("subcategory", string(sub)).
		Fields(mapToZeroParams(extera)).
		Msg(msg)
}

func (l *ZeroLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatal().Msgf(template, args...)
}
