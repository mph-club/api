package server

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func cognitoAuth(ctx iris.Context) {
	isAuth, sub, err := checkToken(ctx.Request().Header.Get("Authorization"))

	if isAuth {
		ctx.Values().Set("sub", sub)
		ctx.Next()
	} else {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(generateJSONResponse(false, iris.Map{"error": iris.Map{"server_error": "Unauthorized", "error_message": err}}))
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
