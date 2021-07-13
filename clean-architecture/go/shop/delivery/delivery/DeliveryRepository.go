package delivery

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
)

func (self *Delivery) Get(c echo.Context) error {
	repository := DeliveryRepository()
	entities := repository.GetList()
	return c.JSON(http.StatusOK, entities)
}

func (self *Delivery) GetbyId(c echo.Context) error{
	repository := DeliveryRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	self = repository.GetID(id)

	return c.JSON(http.StatusOK, self)
}

func (self *Delivery) Persist(c echo.Context) error{
	repository := DeliveryRepository()
	params := make(map[string] string)
	
	c.Bind(&params)
	ObjectMapping(self, params)
	
	
	repository.save(self)

	return c.JSON(http.StatusOK, self)
}

func (self *Delivery) Put(c echo.Context) error{
	repository := DeliveryRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	params := make(map[string] string)
	
	c.Bind(&params)

	err := repository.Update(id, params)

	return c.JSON(http.StatusOK, err)
}

func (self *Delivery) Remove(c echo.Context) error{
	repository := DeliveryRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	self = repository.GetID(id)

	err := repository.Delete(self)

	return c.JSON(http.StatusOK, err)
}