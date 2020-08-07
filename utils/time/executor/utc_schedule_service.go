package executor

import (
	"context"
	"time"
)

type ScheduleExecutor interface {
	Schedule(ctx context.Context)
}

const defaultDuration = 1 * time.Second

type UtcScheduleExecutor struct {
	start     *time.Time
	duration  time.Duration
	executeFn func()
}

type Option func(*UtcScheduleExecutor)

func WithUtcStart(start time.Time) Option {
	return func(executor *UtcScheduleExecutor) {
		utc := start.UTC()
		executor.start = &utc
	}
}

func NewScheduleExecutor(duration time.Duration, execute func(), opts ...Option) ScheduleExecutor {
	executor := &UtcScheduleExecutor{duration: duration, executeFn: execute}
	for _, opt := range opts {
		opt(executor)
	}
	return executor
}

func (s *UtcScheduleExecutor) Schedule(ctx context.Context) {
	endTime := time.Now().UTC()
	duration := betweenOrDefault(s.start, &endTime, defaultDuration)
	duration = defaultIfNegative(duration, defaultDuration)
	ticker := time.NewTicker(duration)
	for {
		select {
		case <-ticker.C:
			ticker = time.NewTicker(s.duration)
			s.executeFn()
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func betweenOrDefault(start *time.Time, end *time.Time, def time.Duration) time.Duration {
	if start == nil || end == nil {
		return def
	}
	return start.Sub(*end)
}

func defaultIfNegative(duration time.Duration, def time.Duration) time.Duration {
	if duration <= 0 {
		duration = def
	}
	return duration
}
