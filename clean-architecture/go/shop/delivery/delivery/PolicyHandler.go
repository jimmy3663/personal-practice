package delivery

import (
	"github.com/mitchellh/mapstructure"
)
var deliveryrepository = DeliveryRepository()

func PolicyHandler(data map[string]interface{}){
	wheneverOrderPlaced_StartDelivery(data)

}

func wheneverOrderPlaced_StartDelivery(data map[string]interface{}){
	event := NewOrderPlaced()
	mapstructure.Decode(data, &event)

	if(event.IsMe()){
		delivery := &Delivery{}
		deliveryrepository.save(delivery)
	}
}
