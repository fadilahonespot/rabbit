package message

type RabbitService interface {
	Publish(topick string, message []byte) (err error)
	Consume(topict string, f func([]byte) error) (err error)
}
