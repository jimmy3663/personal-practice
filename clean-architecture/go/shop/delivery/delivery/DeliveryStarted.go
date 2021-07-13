package delivery

import (
	"encoding/json"
	"time"
	"fmt"
)

type DeliveryStarted struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	OrderId int `json:"orderId" type:"int"` 
	ProductName string `json:"productName" type:"string"` 
	ProductId int `json:"productId" type:"int"` 
	Qty int `json:"qty" type:"int"` 
	
}

func NewDeliveryStarted() *DeliveryStarted{
	event := &DeliveryStarted{EventType:"DeliveryStarted", TimeStamp:time.Now().String()}

	return event
}

func (self *DeliveryStarted) ToJson() string {
	e, err := json.Marshal(self)
	if err != nil{
		
		return "ToJson error"
	}
	
	return string(e)
}

func (self *DeliveryStarted) IsMe() bool{
	
	if getType(self) == self.EventType {
		return true
	} else{
		return false
	}
}

func (self *DeliveryStarted) Publish(){
	message := self.ToJson()
	
	streamhandler(message)
}
