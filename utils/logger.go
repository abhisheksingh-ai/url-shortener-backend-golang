package utils

import (
	"log/slog"
	"os"
	"sync"
)

var (
	onceLogger sync.Once
	Logger     *slog.Logger
)

func InitLogger() *slog.Logger {
	onceLogger.Do(func() {
		// create the log folder
		err := os.MkdirAll("logs", 0755)
		if err != nil {
			panic("logs folder not created: " + err.Error())
		}

		// create the json file
		file, err := os.OpenFile("logs/app.josn", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic("app.json file creation/open issue, error: " + err.Error())
		}

		handler := slog.NewJSONHandler(file, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})

		Logger = slog.New(handler)
	})
	return Logger
}
