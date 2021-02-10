package xRPC

import (
    "bufio"
    "context"
    "crypto/tls"
    "github.com/DGHeroin/xRPC/protocol"
    "net"
    "sync"
    "time"
)

type Server interface {
    ListenAndSerer(ln net.Listener)
    Plugins() PluginContainer
}
type Option struct{}

func NewServer(opts ...Option) Server {
    s := &server{
        activeConn: make(map[net.Conn]struct{}),
    }

    return s
}

type server struct {
    plugins      pluginContainer
    activeConn   map[net.Conn]struct{}
    mu           sync.RWMutex
    readTimeout  time.Duration
    writeTimeout time.Duration
}

func (s *server) Plugins() PluginContainer {
    return &s.plugins
}

func (s *server) ListenAndSerer(ln net.Listener) {
    for {
        conn, err := ln.Accept()
        if err != nil {
            break
        }
        go s.handleConn(conn)
    }
}

func (s *server) handleConn(conn net.Conn) {
    if err := s.plugins.DoConnCreated(conn); err != nil {
        return
    }
    s.mu.Lock()
    s.activeConn[conn] = struct{}{}
    s.mu.Unlock()

    defer func() {
        closeConn(s, conn)
        if err := s.plugins.DoConnClosed(conn); err != nil {
            return
        }
    }()

    if tlsConn, ok := conn.(*tls.Conn); ok {
        if d := s.readTimeout; d != 0 {
            conn.SetReadDeadline(time.Now().Add(d))
        }
        if d := s.writeTimeout; d != 0 {
            conn.SetWriteDeadline(time.Now().Add(d))
        }
        if err := tlsConn.Handshake(); err != nil {
            return
        }
    }
    r := bufio.NewReaderSize(conn, 1024)
    ctx := context.TODO()
    //
    for {
        t0 := time.Now()
        if s.readTimeout != 0 {
            conn.SetReadDeadline(t0.Add(s.readTimeout))
        }

        req, err := s.readRequest(ctx, r)
        if err != nil {
            return
        }
        // handle
        if req != nil {

        }
    }
}

func (s *server) readRequest(ctx context.Context, r *bufio.Reader) (interface{}, error) {
    msg, err := protocol.Read(r)
    if err != nil {
        return nil, err
    }
    m := &MessagePayload{Payload: msg.GetPayload()}
    if err := s.plugins.DoBeforeEncode(m); err != nil {
        return nil, err
    }

    return nil, nil
}

func closeConn(s *server, conn net.Conn) {
    s.mu.Lock()
    delete(s.activeConn, conn)
    s.mu.Unlock()
    _ = conn.Close()
}
