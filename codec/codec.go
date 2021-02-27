package codec

type (
    Codec interface {
        Encode(r interface{}) ([]byte, error)
        Decode(data[]byte, r interface{}) error
    }
)
var (
    JSON Codec = &jsonCodec{}
    MSGPack Codec = &msgpackCodec{}
)
