package log

import (
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestInitLogger(t *testing.T) {
	c := &Config{
		FilePath: "/tmp/test-init-logger.log",
		Format:   "",
		Level:    "debug",
	}

	encoderConfig := NewDefaultEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	c.Encoder = encoderConfig
	Init(c)

	Debug("test debug")
	Debugf("test debugf %s", "test")
	Debugw("test debugw", "key", "value")

	Info("test info")
	Infof("test infof %s", "test")
	Infow("test infow", "key", "value")

	Warn("test warn")
	Warnf("test warnf %s", "test")
	Warnw("test warnw", "key", "value")

	Error("test error")
	Errorf("test errorf %s", "test")
	Errorw("test errorw", "key", "value")

	// Fatal("test fatal")
	// Fatalf("test fatalf %s", "test")
	// Fatalw("test fatalw", "key", "value")

	// Panic("test panic")
	// Panicf("test panicf %s", "test")
	// Panicw("test panic", "key", "value")

}
