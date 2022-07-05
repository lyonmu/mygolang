package main

import (
	"github.com/gookit/slog"
)

func main() {
	//lunar := calendar.NewLunarFromYmd(1998, 8, 11)
	//fmt.Println(lunar.ToFullString())
	//fmt.Println(lunar.GetSolar().ToFullString())
	// use JSON formatter
	slog.Configure(func(logger *slog.SugaredLogger) {
		f := logger.Formatter.(*slog.TextFormatter)
		f.EnableColor = true
	})

	slog.Trace("this is a simple log message")
	slog.Debug("this is a simple log message")
	slog.Info("this is a simple log message")
	slog.Notice("this is a simple log message")
	slog.Warn("this is a simple log message")
	slog.Error("this is a simple log message")
	slog.Fatal("this is a simple log message")
}
