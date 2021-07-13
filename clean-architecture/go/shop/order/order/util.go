package order

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func ObjectMapping(anything interface{}, request map[string]string) {

	target := reflect.ValueOf(anything)
	elements := target.Elem()

	for i := 0; i < elements.NumField(); i++ {
		//mValue := elements.Field(i)
		mType := elements.Type().Field(i)
		tag := mType.Tag
		structFieldValue := elements.FieldByName(mType.Name)
		tagName := tag.Get("json")
		if tag.Get("type") == "int" {
			temp, _ := strconv.Atoi(request[tagName])
			val := reflect.ValueOf(temp)
			structFieldValue.Set(val)
		} else if tag.Get("type") == "string" {
			val := reflect.ValueOf(request[tagName])
			structFieldValue.Set(val)
		} else {
			fmt.Println("ELSE IN OBJECTMAPPING")
		}
	}
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

func ToJson(event interface{}) string {
	e, err := json.Marshal(event)
	if err != nil {

		return "ToJson error"
	}

	return string(e)
}

func IsMe(event interface{}) bool {

	if getType(event) == event.EventType {
		return true
	} else {
		return false
	}
}

func Publish(event interface{}) {
	message := ToJson(event)

	streamhandler(message)
}
