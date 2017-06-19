package zapgoogle

import (
	"cloud.google.com/go/logging"
	"go.uber.org/zap/zapcore"
)

var (
	severity = map[zapcore.Level]logging.Severity{
		zapcore.DebugLevel: logging.Debug,
		zapcore.InfoLevel:  logging.Info,
		zapcore.WarnLevel:  logging.Warning,
		zapcore.ErrorLevel: logging.Error,
		zapcore.FatalLevel: logging.Critical,
	}
)
