package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func cognitoAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		isAuth, sub, err := checkToken(ctx.Request().Header.Get("Authorization"))

		if isAuth {
			ctx.Set("sub", sub)
			return next(ctx)
		}
		return ctx.JSON(generateJSONResponse(false, http.StatusUnauthorized, map[string]interface{}{"server_error": "Unauthorized", "error_message": err}))

	}
}
