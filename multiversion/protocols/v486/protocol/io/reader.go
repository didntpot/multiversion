package io

import (
	legacyprotocol "github.com/didntpot/multiversion/multiversion/protocols/v486/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Reader ...
type Reader struct {
	*protocol.Reader
}

// NewReader ...
func NewReader(r *protocol.Reader) *Reader {
	return &Reader{r}
}

// Reads ...
func (r *Reader) Reads() bool {
	return true
}

// LimitsEnabled ...
func (r *Reader) LimitsEnabled() bool {
	return r.Reader.LimitsEnabled()
}

// PlayerInventoryAction ...
func (r *Reader) PlayerInventoryAction(x *protocol.UseItemTransactionData) {
	r.Varint32(&x.LegacyRequestID)
	if x.LegacyRequestID < -1 && (x.LegacyRequestID&1) == 0 {
		protocol.Slice(r, &x.LegacySetItemSlots)
	}
	protocol.Slice(r, &x.Actions)
	r.Varuint32(&x.ActionType)
	r.BlockPos(&x.BlockPosition)
	r.Varint32(&x.BlockFace)
	r.Varint32(&x.HotBarSlot)
	r.ItemInstance(&x.HeldItem)
	r.Vec3(&x.Position)
	r.Vec3(&x.ClickedPosition)
	r.Varuint32(&x.BlockRuntimeID)
}

// StackRequestAction ...
func (r *Reader) StackRequestAction(x *protocol.StackRequestAction) {
	var id uint8
	r.Uint8(&id)
	if !r.lookupStackRequestAction(id, x) {
		r.UnknownEnumOption(id, "stack request action type")
		return
	}
	(*x).Marshal(r)
}

// lookupStackRequestAction ...
func (r *Reader) lookupStackRequestAction(id uint8, x *protocol.StackRequestAction) bool {
	switch id {
	case protocol.StackRequestActionTake:
		*x = &legacyprotocol.TakeStackRequestAction{TakeStackRequestAction: protocol.TakeStackRequestAction{}}
	case protocol.StackRequestActionPlace:
		*x = &legacyprotocol.PlaceStackRequestAction{PlaceStackRequestAction: protocol.PlaceStackRequestAction{}}
	case protocol.StackRequestActionSwap:
		*x = &legacyprotocol.SwapStackRequestAction{SwapStackRequestAction: protocol.SwapStackRequestAction{}}
	case protocol.StackRequestActionDrop:
		*x = &legacyprotocol.DropStackRequestAction{DropStackRequestAction: protocol.DropStackRequestAction{}}
	case protocol.StackRequestActionDestroy:
		*x = &legacyprotocol.DestroyStackRequestAction{DestroyStackRequestAction: protocol.DestroyStackRequestAction{}}
	case protocol.StackRequestActionConsume:
		*x = &legacyprotocol.ConsumeStackRequestAction{DestroyStackRequestAction: protocol.DestroyStackRequestAction{}}
	case protocol.StackRequestActionCreate:
		*x = &protocol.CreateStackRequestAction{}
	case protocol.StackRequestActionPlaceInContainer:
		*x = &legacyprotocol.PlaceInContainerStackRequestAction{PlaceInContainerStackRequestAction: protocol.PlaceInContainerStackRequestAction{}}
	case protocol.StackRequestActionTakeOutContainer:
		*x = &legacyprotocol.TakeOutContainerStackRequestAction{TakeOutContainerStackRequestAction: protocol.TakeOutContainerStackRequestAction{}}
	case protocol.StackRequestActionLabTableCombine:
		*x = &protocol.LabTableCombineStackRequestAction{}
	case protocol.StackRequestActionBeaconPayment:
		*x = &protocol.BeaconPaymentStackRequestAction{}
	case protocol.StackRequestActionMineBlock:
		*x = &protocol.MineBlockStackRequestAction{}
	case protocol.StackRequestActionCraftRecipe:
		*x = &legacyprotocol.CraftRecipeStackRequestAction{CraftRecipeStackRequestAction: protocol.CraftRecipeStackRequestAction{NumberOfCrafts: 1}}
	case protocol.StackRequestActionCraftRecipeAuto:
		*x = &legacyprotocol.AutoCraftRecipeStackRequestAction{AutoCraftRecipeStackRequestAction: protocol.AutoCraftRecipeStackRequestAction{}}
	case protocol.StackRequestActionCraftCreative:
		*x = &legacyprotocol.CraftCreativeStackRequestAction{CraftCreativeStackRequestAction: protocol.CraftCreativeStackRequestAction{}}
	case protocol.StackRequestActionCraftRecipeOptional:
		*x = &protocol.CraftRecipeOptionalStackRequestAction{}
	case protocol.StackRequestActionCraftGrindstone:
		*x = &protocol.CraftGrindstoneRecipeStackRequestAction{}
	case protocol.StackRequestActionCraftLoom:
		*x = &protocol.CraftLoomRecipeStackRequestAction{}
	case protocol.StackRequestActionCraftNonImplementedDeprecated:
		*x = &protocol.CraftNonImplementedStackRequestAction{}
	case protocol.StackRequestActionCraftResultsDeprecated:
		*x = &protocol.CraftResultsDeprecatedStackRequestAction{}
	default:
		return false
	}
	return true
}

// TransactionDataType ...
func (r *Reader) TransactionDataType(x *protocol.InventoryTransactionData) {
	var transactionType uint32
	r.Varuint32(&transactionType)
	if !r.lookupTransactionData(transactionType, x) {
		r.UnknownEnumOption(transactionType, "inventory transaction data type")
	}
}

// lookupTransactionData ...
func (r *Reader) lookupTransactionData(id uint32, x *protocol.InventoryTransactionData) bool {
	switch id {
	case protocol.InventoryTransactionTypeNormal:
		*x = &protocol.NormalTransactionData{}
	case protocol.InventoryTransactionTypeMismatch:
		*x = &protocol.MismatchTransactionData{}
	case protocol.InventoryTransactionTypeUseItem:
		*x = &legacyprotocol.UseItemTransactionData{}
	case protocol.InventoryTransactionTypeUseItemOnEntity:
		*x = &protocol.UseItemOnEntityTransactionData{}
	case protocol.InventoryTransactionTypeReleaseItem:
		*x = &protocol.ReleaseItemTransactionData{}
	default:
		return false
	}
	return true
}
