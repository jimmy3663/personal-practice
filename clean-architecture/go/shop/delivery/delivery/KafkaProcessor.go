package delivery

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"fmt"
    "encoding/json"
)

var producer *kafka.Producer
var consumer *kafka.Consumer
var topic string

func InitProducer(){

    producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
    if err != nil {
        panic(err)
    }
	topic = "shop"
}

func InitConsumer(){
    consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost",
        "group.id":          "delivery",
        "auto.offset.reset": "earliest",
    })

    if err != nil {
        panic(err)
    }

    go KafkaConsumer()
	
}

func KafkaProducer() (*kafka.Producer, string){
	
	return producer, topic
}

func KafkaConsumer(){
    
	
    consumer.SubscribeTopics([]string{topic}, nil)
    //defer c.Close()
	var dat map[string]interface{}
    for {
        msg, err := consumer.ReadMessage(-1)
        if err == nil {
			if err := json.Unmarshal(msg.Value, &dat); err != nil {
				panic(err)
			}
			PolicyHandler(dat)
        } else {
            // The client will automatically try to recover from all errors.
            fmt.Printf("Consumer error: %v (%v)\n", err, msg)
        }
    }
}