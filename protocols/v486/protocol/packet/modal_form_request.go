package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ModalFormRequest is sent by the server to make the client open a form. This form may be either a modal form
// which has two options, a menu form for a selection of options and a custom form for properties.
type ModalFormRequest struct {
	FormID   uint32
	FormData []byte
}

// ID ...
func (*ModalFormRequest) ID() uint32 {
	return IDModalFormRequest
}

// Marshal ...
func (pk *ModalFormRequest) Marshal(io protocol.IO) {
	io.Varuint32(&pk.FormID)
	io.ByteSlice(&pk.FormData)
}
