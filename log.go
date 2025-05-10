package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

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

func init() {
	encoder := NewDefaultEncoderConfig()
	encoder.EncodeLevel = zapcore.CapitalColorLevelEncoder
	conf := &Config{
		Format:   "console",
		Level:    "debug",
		FilePath: "",
		Encoder:  encoder,
	}
	logger = NewLogger(conf)
	sugar = logger.Sugar()
}

func Init(config *Config) {
	sugar = NewSugaredLogger(config)
}

func GetLogger() *zap.Logger {
	return logger
}

func GetSugaredLogger() *zap.SugaredLogger {
	return sugar
}

func Debug(msg string) {
	sugar.Debug(msg)
}

func Debugf(msg string, args ...any) {
	sugar.Debugf(msg, args...)
}

func Debugw(msg string, keysAndValues ...any) {
	sugar.Debugw(msg, keysAndValues...)
}

func Info(msg string) {
	sugar.Info(msg)
}

func Infof(msg string, args ...any) {
	sugar.Infof(msg, args...)
}

func Infow(msg string, keysAndValues ...any) {
	sugar.Infow(msg, keysAndValues...)
}

func Warn(msg string) {
	sugar.Warn(msg)
}

func Warnf(msg string, args ...any) {
	sugar.Warnf(msg, args...)
}

func Warnw(msg string, keysAndValues ...any) {
	sugar.Warnw(msg, keysAndValues...)
}

func Error(msg string) {
	sugar.Error(msg)
}

func Errorf(msg string, args ...any) {
	sugar.Errorf(msg, args...)
}

func Errorw(msg string, keysAndValues ...any) {
	sugar.Errorw(msg, keysAndValues...)
}

func Fatal(args ...any) {
	sugar.Fatal(args)
}

func Fatalf(template string, args ...any) {
	sugar.Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...any) {
	sugar.Fatalw(msg, keysAndValues...)
}

func Panic(args ...any) {
	sugar.Panic(args)
}

func Panicf(template string, args ...any) {
	sugar.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...any) {
	sugar.Panicw(msg, keysAndValues...)
}
