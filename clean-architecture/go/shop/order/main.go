package main

import(
	"order/order"
)

func main() {
	
	order.OrderDBInit()
	order.InitProducer()
	e := order.RouteInit()

	e.Logger.Fatal(e.Start(":8081"))
}