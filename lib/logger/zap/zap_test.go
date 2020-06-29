package zap_test

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

func Test(t *testing.T) {
	logger, _ := zap.NewProduction()
	zap.Fields()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "/xxx"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
