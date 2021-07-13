package delivery

import (
	"github.com/labstack/echo"
)

func RouteInit() *echo.Echo {
	e := echo.New()
	delivery := &Delivery{}
	e.GET("/deliveries", delivery.Get)
	e.GET("/deliveries/:id", delivery.GetbyId)
	e.POST("/deliveries", delivery.Persist)
	e.PUT("/deliveries/:id", delivery.Put)
	e.DELETE("/deliveries/:id", delivery.Remove)
	return e
}
