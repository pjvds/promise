package messaging

type Bus interface {
	Publish(msg *Message) error
}
