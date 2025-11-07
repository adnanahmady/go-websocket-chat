package request

import (
	"github.com/adnanahmady/go-websocket-chat/pkg/applog"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	loggerKey  = &struct{ uint8 }{}
	requestKey = &struct{ uint8 }{}
)

func logMiddleware(logger applog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.NewString()
		ctx := c.Request.Context()
		ctx = SetLogger(ctx, logger.New())
		ctx = SetRequestID(ctx, id)
		c.Request.WithContext(ctx)
		lgr := GetLogger(ctx)
		lgr.Info("Incoming request")
		lgr.Info("Request URL: %s", c.Request.URL.Path)

		c.Next()

		lgr.Info("Response Code: %d", c.Writer.Status())
		lgr.Info("Response returned")
	}
}
