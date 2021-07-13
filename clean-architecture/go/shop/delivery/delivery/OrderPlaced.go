package delivery

import (
	"encoding/json"
	"time"
	"reflect"
)

type OrderPlaced struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	ProductId int `json:"productId" type:"int"` 
	Qty int `json:"qty" type:"int"` 
	ProductName string `json:"productName" type:"string"` 
	
}

func NewOrderPlaced() *OrderPlaced{
	event := &OrderPlaced{EventType:"OrderPlaced", TimeStamp:time.Now().String()}

	return event
}

func (self *OrderPlaced) ToJson() string {
	e, err := json.Marshal(self)
	if err != nil{
		fmt.Println(err)
		return "ToJson error"
	}
	
	return string(e)
}

func (self *OrderPlaced) IsMe() bool{
	
	if getType(self) == self.EventType {
		return true
	} else{
		return false
	}
}

func (self *OrderPlaced) Publish(){
	message := self.ToJson()
	
	streamhandler(message)
}
