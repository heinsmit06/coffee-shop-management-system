package internal

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func init() {
	file, err := os.OpenFile("internal/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		panic("failed to open log file: " + err.Error())
	}

	// Configure slog to write to the file
	Logger = slog.New(slog.NewJSONHandler(file, nil))
}
