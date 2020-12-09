package _select

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestMultiSelectWithCancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			log.Printf("fist Done")
			return
		}
	}(ctx)

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			log.Printf("second Done")
			return
		}
	}(ctx)

	done := make(chan struct{})
	go func(c chan struct{}) {
		cancel()
		time.Sleep(1 * time.Second)
		c <- struct{}{}
	}(done)
	<-done
}
