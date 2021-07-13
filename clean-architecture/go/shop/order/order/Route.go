package order

import (
	"github.com/labstack/echo"
)

func RouteInit() *echo.Echo {
	e := echo.New()
	order := &Order{}
	e.GET("/orders", order.Get)
	e.GET("/orders/:id", order.GetbyId)
	e.POST("/orders", order.Persist)
	e.PUT("/orders/:id", order.Put)
	e.DELETE("/orders/:id", order.Remove)
	return e
}
