package logger

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

const (
	correlationID = "correlationID"
)

// CreateLogger initialises default logger with system ID.
func CreateLogger(systemID string) *logrus.Entry {
	l := CreateDefaultLogger()
	logger = l.WithField("systemID", systemID)
	return logger
}

// CreateDefaultLogger initialises a default logger.
func CreateDefaultLogger() *logrus.Entry {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
		FieldMap: logrus.FieldMap{
			"file": "caller",
		},
	})
	l.SetReportCaller(true)
	l.SetLevel(logrus.InfoLevel)

	logger = logrus.NewEntry(l)
	return logger
}

// Get returns logger global variable.
func Get() *logrus.Entry {
	return logger
}

// FromContext returns logger from context.
func FromContext(ctx context.Context) *logrus.Entry {
	if logger == nil {
		logger = CreateDefaultLogger()
	}

	return logger.WithField(correlationID, fromContext(ctx, correlationID))
}

func fromContext(ctx context.Context, key string) string {
	if ctx == nil {
		return ""
	}

	if val, ok := ctx.Value(key).(string); ok {
		return val
	}

	return ""
}
