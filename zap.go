package log

import (
	"os"
	"strings"

	enc "github.com/lunuan/go-log/encoder"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	DebugLevel = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel

	RotateMaxSize    = 300
	RotateMaxAge     = 7
	RotateMaxBackups = 15
)

func NewDefaultEncoderConfig() *zapcore.EncoderConfig {
	DefaultEncoderConfig := zap.NewProductionEncoderConfig()
	DefaultEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05,000")
	DefaultEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	DefaultEncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	DefaultEncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	DefaultEncoderConfig.EncodeName = zapcore.FullNameEncoder
	DefaultEncoderConfig.ConsoleSeparator = " "
	DefaultEncoderConfig.LineEnding = zapcore.DefaultLineEnding
	DefaultEncoderConfig.MessageKey = "message"
	DefaultEncoderConfig.StacktraceKey = "stacktrace"
	DefaultEncoderConfig.CallerKey = "caller"
	DefaultEncoderConfig.FunctionKey = "function"
	return &DefaultEncoderConfig
}

func NewLogger(conf *Config) *zap.Logger {
	return initZapLogger(conf)
}

func NewSugaredLogger(conf *Config) *zap.SugaredLogger {
	log := initZapLogger(conf)
	return log.Sugar()
}

func initZapLogger(conf *Config) *zap.Logger {
	//log rotate
	var rotateHook *lumberjack.Logger
	if conf.FilePath != "" {
		maxSize := RotateMaxSize
		if conf.Rotate.MaxSize > 100 {
			maxSize = conf.Rotate.MaxSize
		}

		maxAge := RotateMaxAge
		if conf.Rotate.MaxAge >= 3 {
			maxAge = conf.Rotate.MaxAge
		}

		maxBackups := RotateMaxBackups
		if conf.Rotate.MaxBackups >= 3 {
			maxBackups = conf.Rotate.MaxBackups
		}

		rotateHook = &lumberjack.Logger{
			Filename:   conf.FilePath,
			MaxSize:    maxSize,
			MaxBackups: maxBackups,
			MaxAge:     maxAge,
			Compress:   conf.Rotate.Compress,
			LocalTime:  true,
		}
	}

	//log encoder
	var encoder zapcore.Encoder
	if conf.Encoder == nil {
		conf.Encoder = NewDefaultEncoderConfig()
	}

	switch conf.Format {
	case "json":
		encoder = zapcore.NewJSONEncoder(*conf.Encoder)
	case "console":
		encoder = zapcore.NewConsoleEncoder(*conf.Encoder)
	case "common":
		encoder = enc.NewCommonEncoder(*conf.Encoder)
	default:
		encoder = enc.NewCommonEncoder(*conf.Encoder)
	}

	//log Level
	logLevel := zap.NewAtomicLevel()
	logLevel.SetLevel(toZapLevel(conf.Level))

	//log fileWrites consoleWrites
	var fileWrites zapcore.WriteSyncer
	var consoleWrites zapcore.WriteSyncer
	var core zapcore.Core

	consoleWrites = zapcore.AddSync(os.Stdout)
	if rotateHook != nil {
		fileWrites = zapcore.AddSync(rotateHook)
		fileCore := zapcore.NewCore(encoder, fileWrites, logLevel)
		consoleCore := zapcore.NewCore(encoder, consoleWrites, logLevel)
		core = zapcore.NewTee(fileCore, consoleCore)
	} else {
		core = zapcore.NewCore(encoder, consoleWrites, logLevel)
	}

	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return zapLogger
}

func toZapLevel(level string) zapcore.Level {
	level = strings.ToLower(level)
	logLevel := zap.NewAtomicLevel()
	err := logLevel.UnmarshalText([]byte(level))
	if err != nil {
		logLevel.SetLevel(zapcore.InfoLevel)
	}
	return logLevel.Level()
}
