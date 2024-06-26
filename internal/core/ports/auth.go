package ports

import "github.com/labstack/echo/v4"

type AuthHandler interface {
	Register(c echo.Context) error
	Verify(c echo.Context) error
	ResendCode(c echo.Context) error
	Login(c echo.Context) error
}
