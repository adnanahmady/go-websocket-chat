package request

import (
	"context"
	"github.com/adnanahmady/go-websocket-chat/pkg/applog"
	"time"
)

func ToWsCtx(ctx context.Context) context.Context {
	newCtx := context.Background()
	newCtx = SetLogger(newCtx, GetLogger(ctx))
	newCtx = SetRequestID(newCtx, GetRequestID(ctx))
	return newCtx
}

func SetLogger(ctx context.Context, val applog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, val)
}

func GetLogger(ctx context.Context) applog.Logger {
	return ctx.Value(loggerKey).(applog.Logger)
}

func SetUserName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, userNameKey, name)
}

func SetRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestKey, id)
}

func GetRequestID(ctx context.Context) string {
	id, exists := ctx.Value(requestKey).(string)
	if !exists {
		return "anonymouse-id"
	}
	return id
}

func NewWithTimeout(duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), duration)
}
