package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// GUIDataPickItem is sent by the server to make the client 'select' a hot bar slot. It currently appears to
// be broken however, and does not actually set the selected slot to the hot bar slot set in the packet.
type GUIDataPickItem struct {
	ItemName    string
	ItemEffects string
	HotBarSlot  int32
}

// ID ...
func (*GUIDataPickItem) ID() uint32 {
	return IDGUIDataPickItem
}

// Marshal ...
func (pk *GUIDataPickItem) Marshal(io protocol.IO) {
	io.String(&pk.ItemName)
	io.String(&pk.ItemEffects)
	io.Int32(&pk.HotBarSlot)
}
