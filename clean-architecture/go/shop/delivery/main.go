package main

import(
	"delivery/delivery"
)

func main() {
	
	delivery.DeliveryDBInit()
	delivery.InitProducer()
	delivery.InitConsumer()
	e := delivery.RouteInit()

	e.Logger.Fatal(e.Start(":8082"))
}