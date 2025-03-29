package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	InventoryTransactionTypeNormal = iota
	InventoryTransactionTypeMismatch
	InventoryTransactionTypeUseItem
	InventoryTransactionTypeUseItemOnEntity
	InventoryTransactionTypeReleaseItem
)

// InventoryTransaction is a packet sent by the client. It essentially exists out of multiple sub-packets,
// each of which have something to do with the inventory in one way or another. Some of these sub-packets
// directly relate to the inventory, others relate to interaction with the world, that could potentially
// result in a change in the inventory.
type InventoryTransaction struct {
	LegacyRequestID    int32
	LegacySetItemSlots []protocol.LegacySetItemSlot
	Actions            []protocol.InventoryAction
	TransactionData    protocol.InventoryTransactionData
}

// ID ...
func (*InventoryTransaction) ID() uint32 {
	return IDInventoryTransaction
}

// Marshal ...
func (pk *InventoryTransaction) Marshal(io protocol.IO) {
	io.Varint32(&pk.LegacyRequestID)
	if pk.LegacyRequestID != 0 {
		protocol.Slice(io, &pk.LegacySetItemSlots)
	}
	io.TransactionDataType(&pk.TransactionData)
	protocol.Slice(io, &pk.Actions)
	pk.TransactionData.Marshal(io)
}
