package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// PhotoInfoRequest is sent by the client to request photo information from the server.
type PhotoInfoRequest struct {
	PhotoID int64
}

// ID ...
func (*PhotoInfoRequest) ID() uint32 {
	return IDPhotoInfoRequest
}

// Marshal ...
func (pk *PhotoInfoRequest) Marshal(io protocol.IO) {
	io.Varint64(&pk.PhotoID)
}
