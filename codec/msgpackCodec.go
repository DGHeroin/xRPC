package codec

import (
    "github.com/vmihailenco/msgpack"
)

type msgpackCodec struct {

}

func (m msgpackCodec) Encode(r interface{}) ([]byte, error) {
    return msgpack.Marshal(r)
}

func (m msgpackCodec) Decode(data []byte, r interface{}) error {
    return msgpack.Unmarshal(data, r)
}

