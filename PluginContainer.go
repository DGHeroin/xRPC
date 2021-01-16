package xRPC

import (
    "context"
    "net"
)

type PluginContainer interface {
    Add(plugin Plugin)
    Remove(plugin Plugin)
    All() []Plugin

    DoConnCreated(conn net.Conn) error
    DoConnClosed(conn net.Conn) error

    DoPreCall(ctx context.Context, domain, name, method string, args interface{}) error
    DoPostCall(ctx context.Context, domain, name, method string, args interface{}) error

    DoBeforeEncode(m *MessagePayload) error
    DoAfterEncode(m *MessagePayload) error
}

type (
    MessagePayload struct {
        Payload []byte
    }
    ConnCreatedPlugin interface {
        ConnCreated(conn net.Conn) error
    }

    DoConnClosedPlugin interface {
        DoConnClosed(conn net.Conn) error
    }

    DoPreCallPlugin interface {
        DoPreCall(ctx context.Context, domain, name, method string, args interface{}) error
    }

    DoPostCallPlugin interface {
        DoPostCall(ctx context.Context, domain, name, method string, args interface{}) error
    }

    DoBeforeEncodePlugin interface {
        DoBeforeEncode(m *MessagePayload) error
    }

    DoAfterEncodePlugin interface {
        DoAfterEncode(m *MessagePayload) error
    }
)

type pluginContainer struct {
    plugins []Plugin
}

func (s *pluginContainer) Add(plugin Plugin) {
    s.plugins = append(s.plugins, plugin)
}

func (s *pluginContainer) Remove(plugin Plugin) {
    var keep []Plugin
    for _, p := range s.plugins {
        if p != plugin {
            keep = append(keep, p)
        }
    }
    s.plugins = keep
}

func (s *pluginContainer) All() []Plugin {
    return s.plugins
}

func (s *pluginContainer) DoConnCreated(conn net.Conn) error {
    for _, p := range s.plugins {
        if obj, ok := p.(ConnCreatedPlugin); ok {
            if err := obj.ConnCreated(conn); err != nil {
                return err
            }

        }
    }
    return nil
}

func (s *pluginContainer) DoConnClosed(conn net.Conn) error {
    for _, p := range s.plugins {
        if obj, ok := p.(DoConnClosedPlugin); ok {
            if err := obj.DoConnClosed(conn); err != nil {
                return err
            }

        }
    }
    return nil
}

func (s *pluginContainer) DoPreCall(ctx context.Context, domain, name, method string, args interface{}) error {
    for _, p := range s.plugins {
        if obj, ok := p.(DoPreCallPlugin); ok {
            if err := obj.DoPreCall(ctx, domain, name, method, args); err != nil {
                return err
            }

        }
    }
    return nil
}

func (s *pluginContainer) DoPostCall(ctx context.Context, domain, name, method string, args interface{}) error {
    for _, p := range s.plugins {
        if obj, ok := p.(DoPostCallPlugin); ok {
            if err := obj.DoPostCall(ctx, domain, name, method, args); err != nil {
                return err
            }

        }
    }
    return nil
}
func (s *pluginContainer) DoBeforeEncode(m *MessagePayload) error {
    for _, p := range s.plugins {
        if obj, ok := p.(DoBeforeEncodePlugin); ok {
            if err := obj.DoBeforeEncode(m); err != nil {
                return err
            }

        }
    }
    return nil
}

func (s *pluginContainer) DoAfterEncode(m *MessagePayload) error {
    for _, p := range s.plugins {
        if obj, ok := p.(DoAfterEncodePlugin); ok {
            if err := obj.DoAfterEncode(m); err != nil {
                return err
            }

        }
    }
    return nil
}
