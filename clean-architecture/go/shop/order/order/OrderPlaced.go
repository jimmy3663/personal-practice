package order

import (
	"time"
)

type OrderPlaced struct {
	EventType   string `json:"eventType" type:"string"`
	TimeStamp   string `json:"timeStamp" type:"string"`
	Id          int    `json:"id" type:"int"`
	ProductId   int    `json:"productId" type:"int"`
	Qty         int    `json:"qty" type:"int"`
	ProductName string `json:"productName" type:"string"`
}

func NewOrderPlaced() *OrderPlaced {
	event := &OrderPlaced{EventType: "OrderPlaced", TimeStamp: time.Now().String()}

	return event
}
