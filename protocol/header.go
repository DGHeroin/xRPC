package protocol

import "encoding/binary"

const (
    flagBody      = 0x1
    flagHeartbeat = 0x2
    flagOneway    = 0x4
)

// format
// [0] magic number
// [1] 0000 || oneway flag | heartbeat | has body |
// [2] reserved byte
// [3] reserved byte
// [4-8] seq
// [8-12] size

type Header [12]byte

var zeroHeaderArray Header
var zeroHeader = zeroHeaderArray[1:]

func resetHeader(h *Header) {
    copy(h[1:], zeroHeader)
}

func (h Header) CheckMagicNumber() bool {
    return h[0] == magicNumber
}
func (h Header) IsHeartbeat() bool {
    return h[1]&flagHeartbeat == flagHeartbeat
}
func (h *Header) SetHeartbeat(hb bool) {
    if hb {
        h[1] = h[1] | flagHeartbeat
    } else {
        h[1] = h[1] &^ flagHeartbeat
    }
}

func (h Header) IsOneway() bool {
    return h[1]&flagOneway == flagOneway
}
func (h *Header) SetOneway(oneway bool) {
    if oneway {
        h[1] = h[1] | flagOneway
    } else {
        h[1] = h[1] &^ flagOneway
    }
}
func (h Header) Seq() uint32 {
    return binary.BigEndian.Uint32(h[4:])
}
func (h *Header) SetSeq(seq uint32) {
    binary.BigEndian.PutUint32(h[4:], seq)
}

func (h *Header) HasBody() bool {
    return h[2]&flagBody == flagBody
}

func (h *Header) BodyLength() uint32 {
    return binary.BigEndian.Uint32(h[8:])
}
