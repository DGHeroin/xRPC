package protocol

import "sync"

var (
    syncPool = sync.Pool{
        New: func() interface{} {
            header := Header([12]byte{})
            header[0] = magicNumber

            return &Message{
                Header: &header,
            }
        },
    }
)

func GetPooledMessage() *Message {
    return syncPool.Get().(*Message)
}

func FreeMessage(msg*Message)  {
    if msg != nil {
        msg.Reset()
        syncPool.Put(msg)
    }
}
