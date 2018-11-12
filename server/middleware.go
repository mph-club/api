package server

import (
	"net/http"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/labstack/echo"
)

func cognitoAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		isAuth, sub, err := checkToken(ctx.Request().Header.Get("Authorization"))

		if isAuth {
			ctx.Set("sub", sub)
			return next(ctx)
		} else {
			return ctx.JSON(generateJSONResponse(false, http.StatusUnauthorized, map[string]interface{}{"server_error": "Unauthorized", "error_message": err}))
		}
	}
}

func requestLogger() iris.Handler {
	return logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,

		//Columns: true,

		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})
}
