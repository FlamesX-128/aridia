package middlewares

import (
	"github.com/labstack/echo"
)

func JSON(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		ctx.Response().Header().Set(echo.HeaderAccept, echo.MIMEApplicationJSON)

		return next(ctx)
	}
}
