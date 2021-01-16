package xRPC

type Request struct {
    ServiceName string
    Method string
    request interface{}
    reponse interface{}
}
