package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func GetLogger(env string, name string) *slog.Logger {
	if logger == nil {
		loggerObj := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logger = loggerObj.With(slog.Group("app", slog.String("env", env), slog.String("name", name)))
	}
	return logger
}