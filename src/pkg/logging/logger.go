package logging

import "golang-web-api/config"

type Logger interface {
	Init()
	Info(cat Category, sub subCategory, msg string, extera map[ExtraKey]interface{})
	Infof(template string, args ...interface{})

	Debug(cat Category, sub subCategory, msg string, extera map[ExtraKey]interface{})
	Debugf(template string, args ...interface{})

	Warn(cat Category, sub subCategory, msg string, extera map[ExtraKey]interface{})
	Warnf(template string, args ...interface{})

	Error(cat Category, sub subCategory, msg string, extera map[ExtraKey]interface{})
	Errorf(template string, args ...interface{})

	Fatal(cat Category, sub subCategory, msg string, extera map[ExtraKey]interface{})
	Fatalf(template string, args ...interface{})
}

func NewLogger(cfg *config.Logger) Logger {
	if cfg.Logger == "zap" {
		return newZapLogger(cfg)
	}else if cfg.Logger == "zerolog" {
		return NewZeroLogger(cfg)
	}
	return nil
}