package xRPC

import (
    "bufio"
    "bytes"
    "encoding/binary"
)

type (
    Protocol interface {
    }
    ProtocolMessage interface {
        SetType(t uint8)
        SetBody(data []byte)
        SetHeader(key, value string)
        DelHeader(key string)
        Encode() []byte
        Decode(data []byte)
    }
)
type (
    _protocol struct {
    }
    _protocolMessage struct {
        t       uint8
        body    []byte
        headers map[string]string
    }
)

const (
    maxHeaderSize = 16 << 10
)

func (pm *_protocolMessage) SetType(t uint8) {
    pm.t = t
}

func (pm *_protocolMessage) SetBody(data []byte) {
    pm.body = data
}

func (pm *_protocolMessage) SetHeader(key, value string) {
    if len(key) > maxHeaderSize { // 16k
        return
    }
    if len(value) > maxHeaderSize { // 16k
        return
    }
    pm.headers[key] = value
}

func (pm *_protocolMessage) DelHeader(key string) {
    delete(pm.headers, key)
}

func (pm *_protocolMessage) Encode() []byte {
    buf := bytes.NewBuffer(nil)
    { // type sec
        h := make([]byte, 2)
        binary.BigEndian.PutUint16(h, uint16(pm.t))
        buf.Write(h)
    }
    { // header sec
        // number of header
        h := make([]byte, 4)
        binary.BigEndian.PutUint32(h, uint32(len(pm.headers)))
        buf.Write(h)
        for k, v := range pm.headers {
            binary.BigEndian.PutUint32(h, uint32(len(k)))
            buf.Write(h)
            binary.BigEndian.PutUint32(h, uint32(len(v)))
            buf.Write(h)
        }
    }
    { // body sec
        h := make([]byte, 4)
        binary.BigEndian.PutUint32(h, uint32(len(pm.body)))
        buf.Write(h)
        if len(pm.body) > 0 {
            buf.Write(pm.body)
        }
    }

    return buf.Bytes()
}

func (pm *_protocolMessage) Decode(data []byte) {
    r := bufio.NewReader(bytes.NewBuffer(data))
    { // type sec
        h := make([]byte, 2)
        r.Read(h)
        n := binary.BigEndian.Uint16(h)
        pm.t = uint8(n)
    }
    { // header sec
        h := make([]byte, 4)
        r.Read(h)
        n := binary.BigEndian.Uint32(h)
        for i := 0; i < int(n); i++ {
            // key
            h = make([]byte, 4)
            r.Read(h)
            n := binary.BigEndian.Uint32(h)
            if n > maxHeaderSize {
                // error
                return
            }
            h = make([]byte, n)
            r.Read(h)
            key := string(h)
            // value
            h = make([]byte, 4)
            r.Read(h)
            n = binary.BigEndian.Uint32(h)
            if n > maxHeaderSize {
                // error
                return
            }
            h = make([]byte, n)
            r.Read(h)
            value := string(h)
            // done
            pm.headers[key] = value
        }
    }
    { // body sec
        h := make([]byte, 4)
        r.Read(h)
        n := binary.BigEndian.Uint32(h)
        if n > maxHeaderSize {
            // error
            return
        }
        h = make([]byte, n)
        r.Read(h)
        pm.body = h
    }
}
func NewProtocol() Protocol {
    p := &_protocol{}
    return p
}
func NewProtocolMessage() ProtocolMessage {
    pm := &_protocolMessage{}
    return pm
}
