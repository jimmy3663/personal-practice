package delivery

type EventInterface interface {
	ToJson() string
	IsMe() bool
	Publish() 
}
