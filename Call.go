package xRPC

type Call struct {
    Done        chan *Call
    ServiceName string
    Method      string
    Seq         uint32
    Args        interface{}
    Reply       interface{}
}
