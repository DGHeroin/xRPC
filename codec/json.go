package codec

import "encoding/json"

type jsonCodec struct {}

func (j jsonCodec) Encode(r interface{}) ([]byte, error) {
    return json.Marshal(r)
}

func (j jsonCodec) Decode(data []byte, r interface{}) error {
    return json.Unmarshal(data, r)
}

