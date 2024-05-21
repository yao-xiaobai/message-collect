package kafka

// Handler
type Handler interface {
	handle(message []byte) error
}
