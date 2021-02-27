package xRPC

type Connection interface {
    Send([]byte) error
}
