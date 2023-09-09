package internal

import (
	"io"

	"github.com/charmbracelet/lipgloss"
	log "github.com/charmbracelet/log"
)

type MyLogger struct {
	*log.Logger
}

var Logger *MyLogger

func NewLogger(level string, w io.Writer) {
	if Logger != nil {
		return
	}
	log.DebugLevelStyle = lipgloss.NewStyle().SetString("DEBUG")
	log.ErrorLevelStyle = lipgloss.NewStyle().SetString("ERROR")
	l := log.NewWithOptions(w, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
	})
	// Code to set the log level
	switch level {
	case "debug":
		l.SetLevel(log.DebugLevel)
	case "info":
		l.SetLevel(log.InfoLevel)
	case "warn":
		l.SetLevel(log.WarnLevel)
	case "error":
		l.SetLevel(log.ErrorLevel)
	default:
		l.SetLevel(log.InfoLevel)
	}
	Logger = &MyLogger{l}
}
