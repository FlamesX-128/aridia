package middlewares

import "github.com/labstack/echo"

func JSON(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().Header().Set(echo.HeaderAccept, echo.MIMEApplicationJSON)

		return next(c)
	}
}
