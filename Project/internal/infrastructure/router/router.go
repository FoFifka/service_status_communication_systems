package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"skillsbox-diploma/internal/interface/controller"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/sms", func(ctx echo.Context) error {
		return c.GetSMSData(ctx)
	})

	e.GET("/mms", func(ctx echo.Context) error { // TODO routes
		return c.GetMMSData(ctx)
	})

	e.GET("/email", func(ctx echo.Context) error {
		return c.GetEmailData(ctx)
	})

	e.GET("/voicecall", func(ctx echo.Context) error {
		return c.GetVoiceCallData(ctx)
	})

	e.GET("/billing", func(ctx echo.Context) error {
		return c.GetBillingData(ctx)
	})

	e.GET("/support", func(ctx echo.Context) error {
		return c.GetSupportData(ctx)
	})

	e.GET("/incident", func(ctx echo.Context) error {
		return c.GetIncidentData(ctx)
	})

	e.GET("/result", func(ctx echo.Context) error {
		return c.GetResultData(ctx)
	})

	return e
}
