package xRPC

import (
    "context"
    "net"
)

type Client interface {
    //Dial(network string, address string) error
    Plugins() PluginContainer
    Call(ctx context.Context, serviceName string, method string, r interface{}, w interface{})
}

func NewClient(discovery Discovery) Client {
    c := &client{
        discovery: discovery,
    }
    return c
}

type client struct {
    conn      net.Conn
    plugins   pluginContainer
    discovery Discovery
}

func (c *client) Call(ctx context.Context, serviceName string, method string, r interface{}, w interface{}) {

}

func (c *client) Plugins() PluginContainer {
    return &c.plugins
}
