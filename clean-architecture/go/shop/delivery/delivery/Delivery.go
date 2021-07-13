package delivery

import (
	"gopkg.in/jeevatkm/go-model.v1"
	
	"gorm.io/gorm"
	
)
type Delivery struct {
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	OrderId int `json:"orderId" type:"int"`
	ProductName string `json:"productName" type:"string"`
	ProductId int `json:"productId" type:"int"`
	Qty int `json:"qty" type:"int"`
	
	
}

func (self *Delivery) AfterCreate(tx *gorm.DB) (err error){
	deliveryStarted := NewDeliveryStarted()
	model.Copy(deliveryStarted, self)

	deliveryStarted.Publish()
	
	return nil
}
