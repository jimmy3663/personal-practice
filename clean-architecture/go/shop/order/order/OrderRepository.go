package order

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (self *Order) Get(c echo.Context) error {
	repository := OrderRepository()
	entities := repository.GetList()
	return c.JSON(http.StatusOK, entities)
}

func (self *Order) GetbyId(c echo.Context) error {
	repository := OrderRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	self = repository.GetID(id)

	return c.JSON(http.StatusOK, self)
}

func (self *Order) Persist(c echo.Context) error {
	repository := OrderRepository()
	params := make(map[string]string)

	c.Bind(&params)
	ObjectMapping(self, params)

	repository.save(self)

	return c.JSON(http.StatusOK, self)
}

func (self *Order) Put(c echo.Context) error {
	repository := OrderRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	params := make(map[string]string)

	c.Bind(&params)

	err := repository.Update(id, params)

	return c.JSON(http.StatusOK, err)
}

func (self *Order) Remove(c echo.Context) error {
	repository := OrderRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	self = repository.GetID(id)

	err := repository.Delete(self)

	return c.JSON(http.StatusOK, err)
}
