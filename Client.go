package xRPC

import (
    "context"
    "fmt"
    "github.com/DGHeroin/xRPC/codec"
    "github.com/DGHeroin/xRPC/protocol"
    "sync/atomic"
)

type Client interface {
    Plugins() PluginContainer
    Call(ctx context.Context, serviceName string, method string, r interface{}, w interface{}) error
}

func NewClient(discovery Discovery) Client {
    c := &client{
        discovery: discovery,
        Codec:     codec.JSON,
    }
    return c
}

type client struct {
    plugins   pluginContainer
    discovery Discovery
    msgSeq    uint32
    Codec     codec.Codec
}

func (c *client) Call(ctx context.Context, serviceName string, method string, r interface{}, w interface{}) error {
    call := c.Go(ctx, serviceName, method, r, w)

    select {
    case <- ctx.Done():
        return ctx.Err()
        case <-call.Done:
    }
    return nil
}

func (c *client) Go(ctx context.Context, serviceName string, method string, r interface{}, w interface{}) *Call {
    call := new(Call)
    call.Done = make(chan *Call)
    call.ServiceName = serviceName
    call.Method = method
    call.Seq = atomic.AddUint32(&c.msgSeq, 1)
    call.Args = r
    call.Reply = w

    c.send(ctx, call)
    return call
}

func (c *client) Plugins() PluginContainer {
    return &c.plugins
}

func (c *client) send(ctx context.Context, call *Call)  {
    msg := protocol.GetPooledMessage()
    if call.ServiceName == "" && call.Method == "" {
        msg.SetHeartbeat(true)
    }
    msg.Method = call.Method
    data, err := c.Codec.Encode(call.Args)
    if err != nil {
        // xxx
    }
    msg.Payload = data
    bin, _ := msg.Encode(c.Codec)
    fmt.Println(string(bin))
    conn := c.discovery.GetConnection()
    if conn == nil {
        //
    }
    conn.Send(bin)
    //c.conn.Write(bin)
}