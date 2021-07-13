package order

type EventInterface interface {
	ToJson() string
	IsMe() bool
	Publish() 
}
