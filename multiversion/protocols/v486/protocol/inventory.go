package protocol

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// UseItemTransactionData represents an inventory transaction data object sent when the client uses an item on
// a block.
type UseItemTransactionData struct {
	LegacyRequestID    int32
	LegacySetItemSlots []protocol.LegacySetItemSlot
	Actions            []protocol.InventoryAction
	ActionType         uint32
	BlockPosition      protocol.BlockPos
	BlockFace          int32
	HotBarSlot         int32
	HeldItem           protocol.ItemInstance
	Position           mgl32.Vec3
	ClickedPosition    mgl32.Vec3
	BlockRuntimeID     uint32
}

// Marshal ...
func (x *UseItemTransactionData) Marshal(io protocol.IO) {
	io.Varuint32(&x.ActionType)
	io.UBlockPos(&x.BlockPosition)
	io.Varint32(&x.BlockFace)
	io.Varint32(&x.HotBarSlot)
	io.ItemInstance(&x.HeldItem)
	io.Vec3(&x.Position)
	io.Vec3(&x.ClickedPosition)
	io.Varuint32(&x.BlockRuntimeID)
}
