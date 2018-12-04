package server

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func connectLogger() echo.MiddlewareFunc {
	logger := newLogger()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			fields := []zapcore.Field{
				zap.Int("status", res.Status),
				zap.String("latency", time.Since(start).String()),
				zap.String("id", id),
				zap.String("method", req.Method),
				zap.String("uri", req.RequestURI),
				zap.String("host", req.Host),
				zap.String("remote_ip", c.RealIP()),
			}

			n := res.Status
			switch {
			case n >= 500:
				logger.Error("Server error", fields...)
			case n >= 400:
				logger.Warn("Client error", fields...)
			case n >= 300:
				logger.Info("Redirection", fields...)
			default:
				logger.Info("Success", fields...)
			}

			return nil
		}
	}
}

func cognitoAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		isAuth, sub, err := checkToken(ctx.Request().Header.Get("Authorization"))
		fromLambda := ctx.Request().Header.Get("X-From-Lambda")

		if isAuth {
			ctx.Set("sub", sub)
			return next(ctx)
		} else if fromLambda == os.Getenv("COGNITO_USER_POOL_ID") {
			return next(ctx)
		}

		return ctx.JSON(response(false, http.StatusUnauthorized, map[string]interface{}{"server_error": "Unauthorized", "error_message": err}))

	}
}
