package protocol

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// CraftRecipeStackRequestAction is sent by the client the moment it begins crafting an item. This is the
// first action sent, before the Consume and Create item stack request actions.
// This action is also sent when an item is enchanted. Enchanting should be treated mostly the same way as
// crafting, where the old item is consumed.
type CraftRecipeStackRequestAction struct {
	RecipeNetworkID uint32
}

// Marshal ...
func (x *CraftRecipeStackRequestAction) Marshal(io protocol.IO) {
	io.Varuint32(&x.RecipeNetworkID)
}

// CraftCreativeStackRequestAction is sent by the client when it takes an item out fo the creative inventory.
// The item is thus not really crafted, but instantly created.
type CraftCreativeStackRequestAction struct {
	protocol.CraftCreativeStackRequestAction
}

// Marshal ...
func (x *CraftCreativeStackRequestAction) Marshal(io protocol.IO) {
	io.Varuint32(&x.CreativeItemNetworkID)
}

// CraftRecipeOptionalStackRequestAction is sent when using an anvil. When this action is sent, the
// FilterStrings field in the respective stack request is non-empty and contains the name of the item created
// using the anvil or cartography table.
type CraftRecipeOptionalStackRequestAction struct {
	RecipeNetworkID   uint32
	FilterStringIndex int32
}

// Marshal ...
func (x *CraftRecipeOptionalStackRequestAction) Marshal(io protocol.IO) {
	io.Varuint32(&x.RecipeNetworkID)
	io.Int32(&x.FilterStringIndex)
}

// CraftGrindstoneRecipeStackRequestAction is sent when a grindstone recipe is crafted. It contains the RecipeNetworkID
// to identify the recipe crafted, and the cost for crafting the recipe.
type CraftGrindstoneRecipeStackRequestAction struct {
	RecipeNetworkID uint32
	Cost            int32
}

// Marshal ...
func (x *CraftGrindstoneRecipeStackRequestAction) Marshal(io protocol.IO) {
	io.Varuint32(&x.RecipeNetworkID)
	io.Varint32(&x.Cost)
}

// StackRequestSlotInfo holds information on a specific slot client-side.
type StackRequestSlotInfo struct { // TODO: do we use this?
	ContainerID    byte
	Slot           byte
	StackNetworkID int32
}

// StackResponseSlotInfo holds information on what item stack should be present in a specific slot.
type StackResponseSlotInfo struct {
	protocol.StackResponseSlotInfo
}

// Marshal ...
func (x *StackResponseSlotInfo) Marshal(io protocol.IO) {
	io.Uint8(&x.Slot)
	io.Uint8(&x.HotbarSlot)
	io.Uint8(&x.Count)
	io.Varint32(&x.StackNetworkID)
	if x.Slot != x.HotbarSlot {
		io.InvalidValue(x.HotbarSlot, "hotbar slot", "hot bar slot must be equal to normal slot")
	}
	io.String(&x.CustomName)
	io.Varint32(&x.DurabilityCorrection)
}

// StackResponseContainerInfo holds information on what slots in a container have what item stack in them.
type StackResponseContainerInfo struct {
	ContainerID byte
	SlotInfo    []StackResponseSlotInfo
}

// Marshal ...
func (x *StackResponseContainerInfo) Marshal(io protocol.IO) {
	io.Uint8(&x.ContainerID)
	protocol.Slice(io, &x.SlotInfo)
}

// ItemStackResponse is a response to an individual ItemStackRequest.
type ItemStackResponse struct {
	Status        uint8
	RequestID     int32
	ContainerInfo []StackResponseContainerInfo
}

// Marshal ...
func (x *ItemStackResponse) Marshal(io protocol.IO) {
	io.Uint8(&x.Status)
	io.Varint32(&x.RequestID)
	if x.Status == protocol.ItemStackResponseStatusOK {
		protocol.Slice(io, &x.ContainerInfo)
	}
}
