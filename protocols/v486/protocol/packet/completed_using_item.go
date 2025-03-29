package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	UseItemEquipArmor = iota
	UseItemEat
	UseItemAttack
	UseItemConsume
	UseItemThrow
	UseItemShoot
	UseItemPlace
	UseItemFillBottle
	UseItemFillBucket
	UseItemPourBucket
	UseItemUseTool
	UseItemInteract
	UseItemRetrieved
	UseItemDyed
	UseItemTraded
)

// CompletedUsingItem is sent by the server to tell the client that it should be done using the item it is
// currently using.
type CompletedUsingItem struct {
	UsedItemID int16
	UseMethod  int32
}

// ID ...
func (*CompletedUsingItem) ID() uint32 {
	return IDCompletedUsingItem
}

// Marshal ...
func (pk *CompletedUsingItem) Marshal(io protocol.IO) {
	io.Int16(&pk.UsedItemID)
	io.Int32(&pk.UseMethod)
}
