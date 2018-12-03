package server

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

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
