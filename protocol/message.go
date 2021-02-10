package protocol

import (
    "fmt"
    "github.com/DGHeroin/xRPC/codec"
    "io"
)

type Message struct {
    *Header
    Method  string
    Payload []byte
}

const (
    magicNumber byte = 0x88
)

type MessageType byte

const (
    Request MessageType = iota
    Response
)

func NewMessage() *Message {
    header := Header([12]byte{})
    header[0] = magicNumber

    return &Message{
        Header: &header,
    }
}
func Read(r io.Reader) (*Message, error) {
    msg := NewMessage()
    err := msg.Decode(r)
    if err != nil {
        return nil, err
    }
    return msg, nil
}

func (m *Message) Decode(r io.Reader) error {
    // header

    if _, err := io.ReadFull(r, m.Header[:1]); err != nil {
        return err
    }
    if !m.Header.CheckMagicNumber() {
        return fmt.Errorf("worong magic number:%v", m.Header[0])
    }
    //
    if _, err := io.ReadFull(r, m.Header[1:]); err != nil {
        return err
    }
    // data len
    if !m.Header.HasBody() {
        return nil
    }
    // read body
    sz := m.Header.BodyLength()
    if sz > 0 {
        m.Payload = make([]byte, sz)
        if _, err := io.ReadFull(r, m.Payload); err != nil {
            return err
        }
    }
    return nil
}
func (m *Message) GetPayload() []byte {
    return m.Payload
}

func (m *Message) Reset() {
    resetHeader(m.Header)
    m.Method = ""
    m.Payload = nil
}
func (m *Message) Encode(codec codec.Codec) ([]byte,error) {
    return codec.Encode(m)
}
