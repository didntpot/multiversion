package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ServerSettingsResponse is optionally sent by the server in response to a ServerSettingsRequest from the
// client. It is structured the same as a ModalFormRequest packet, and if filled out correctly, will show
// a specific tab for the server in the settings of the client. A ModalFormResponse packet is sent by the
// client in response to a ServerSettingsResponse, when the client fills out the settings and closes the
// settings again.
type ServerSettingsResponse struct {
	FormID   uint32
	FormData []byte
}

// ID ...
func (*ServerSettingsResponse) ID() uint32 {
	return IDServerSettingsResponse
}

// Marshal ...
func (pk *ServerSettingsResponse) Marshal(io protocol.IO) {
	io.Varuint32(&pk.FormID)
	io.ByteSlice(&pk.FormData)
}
