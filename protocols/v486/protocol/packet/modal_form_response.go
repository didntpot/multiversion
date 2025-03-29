package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ModalFormResponse is sent by the client in response to a ModalFormRequest, after the player has submitted
// the form sent. It contains the options/properties selected by the player, or a JSON encoded 'null' if
// the form was closed by clicking the X in the top right corner of the form.
type ModalFormResponse struct {
	FormID       uint32
	ResponseData []byte
}

// ID ...
func (*ModalFormResponse) ID() uint32 {
	return IDModalFormResponse
}

// Marshal ...
func (pk *ModalFormResponse) Marshal(io protocol.IO) {
	io.Varuint32(&pk.FormID)
	io.ByteSlice(&pk.ResponseData)
}
