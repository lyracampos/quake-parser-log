package logger

import "go.uber.org/zap"

var Log *zap.SugaredLogger

// nolint: errcheck
func InitLog() {
	logger := zap.Must(zap.NewProduction())

	defer logger.Sync()

	Log = logger.Sugar()
}
