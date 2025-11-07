package request

import (
	"context"
	"github.com/adnanahmady/go-websocket-chat/pkg/applog"
	"time"
)

func SetLogger(ctx context.Context, val applog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, val)
}

func GetLogger(ctx context.Context) applog.Logger {
	return ctx.Value(loggerKey).(applog.Logger)
}

func SetRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestKey, id)
}

func NewWithTimeout(duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), duration)
}
