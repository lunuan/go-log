package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

type Config struct {
	Format   string
	Level    string
	FilePath string
	Rotate   RotateConfig
	Encoder  *zapcore.EncoderConfig
}

type RotateConfig struct {
	Compress   bool
	MaxSize    int
	MaxAge     int
	MaxBackups int
}

func Init(config *Config) {
	logger = NewSugaredLogger(config)
}

func Debug(msg string) {
	logger.Debug(msg)
}

func Debugf(msg string, args ...any) {
	logger.Debugf(msg, args...)
}

func Debugw(msg string, keysAndValues ...any) {
	logger.Debugw(msg, keysAndValues...)
}

func Info(msg string) {
	logger.Info(msg)
}

func Infof(msg string, args ...any) {
	logger.Infof(msg, args...)
}

func Infow(msg string, keysAndValues ...any) {
	logger.Infow(msg, keysAndValues...)
}

func Warn(msg string) {
	logger.Warn(msg)
}

func Warnf(msg string, args ...any) {
	logger.Warnf(msg, args...)
}

func Warnw(msg string, keysAndValues ...any) {
	logger.Warnw(msg, keysAndValues...)
}

func Error(msg string) {
	logger.Error(msg)
}

func Errorf(msg string, args ...any) {
	logger.Errorf(msg, args...)
}

func Errorw(msg string, keysAndValues ...any) {
	logger.Errorw(msg, keysAndValues...)
}

func Fatal(args ...any) {
	logger.Fatal(args)
}

func Fatalf(template string, args ...any) {
	logger.Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...any) {
	logger.Fatalw(msg, keysAndValues...)
}

func Panic(args ...any) {
	logger.Panic(args)
}

func Panicf(template string, args ...any) {
	logger.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...any) {
	logger.Panicw(msg, keysAndValues...)
}
