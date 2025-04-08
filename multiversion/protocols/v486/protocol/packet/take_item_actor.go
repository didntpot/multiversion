package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// TakeItemActor is sent by the server when a player picks up an item entity. It makes the item entity
// disappear to viewers and shows the pick-up animation.
type TakeItemActor struct {
	ItemEntityRuntimeID  uint64
	TakerEntityRuntimeID uint64
}

// ID ...
func (*TakeItemActor) ID() uint32 {
	return IDTakeItemActor
}

// Marshal ...
func (pk *TakeItemActor) Marshal(io protocol.IO) {
	io.Varuint64(&pk.ItemEntityRuntimeID)
	io.Varuint64(&pk.TakerEntityRuntimeID)
}
