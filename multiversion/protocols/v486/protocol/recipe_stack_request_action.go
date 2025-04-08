package protocol

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// TakeStackRequestAction is sent by the client to the server to take x amount of items from one slot in a
// container to the cursor.
type TakeStackRequestAction struct {
	protocol.TakeStackRequestAction
}

// Marshal ...
func (x *TakeStackRequestAction) Marshal(io protocol.IO) {
	io.Uint8(&x.Count)
	StackReqSlotInfo(io, &x.Source)
	StackReqSlotInfo(io, &x.Destination)
}

// PlaceStackRequestAction is sent by the client to the server to place x amount of items from one slot into
// another slot, such as when shift clicking an item in the inventory to move it around or when moving an item
// in the cursor into a slot.
type PlaceStackRequestAction struct {
	protocol.PlaceStackRequestAction
}

// Marshal ...
func (x *PlaceStackRequestAction) Marshal(io protocol.IO) {
	io.Uint8(&x.Count)
	StackReqSlotInfo(io, &x.Source)
	StackReqSlotInfo(io, &x.Destination)
}

// SwapStackRequestAction is sent by the client to swap the item in its cursor with an item present in another
// container. The two item stacks swap places.
type SwapStackRequestAction struct {
	protocol.SwapStackRequestAction
}

// Marshal ...
func (x *SwapStackRequestAction) Marshal(io protocol.IO) {
	StackReqSlotInfo(io, &x.Source)
	StackReqSlotInfo(io, &x.Destination)
}

// DropStackRequestAction is sent by the client when it drops an item out of the inventory when it has its
// inventory opened. This action is not sent when a player drops an item out of the hotbar using the Q button
// (or the equivalent on mobile). The InventoryTransaction packet is still used for that action, regardless of
// whether the item stack network IDs are used or not.
type DropStackRequestAction struct {
	protocol.DropStackRequestAction
}

// Marshal ...
func (x *DropStackRequestAction) Marshal(io protocol.IO) {
	io.Uint8(&x.Count)
	StackReqSlotInfo(io, &x.Source)
	io.Bool(&x.Randomly)
}

// DestroyStackRequestAction is sent by the client when it destroys an item in creative mode by moving it
// back into the creative inventory.
type DestroyStackRequestAction struct {
	protocol.DestroyStackRequestAction
}

// Marshal ...
func (x *DestroyStackRequestAction) Marshal(io protocol.IO) {
	io.Uint8(&x.Count)
	StackReqSlotInfo(io, &x.Source)
}

// ConsumeStackRequestAction is sent by the client when it uses an item to craft another item. The original
// item is 'consumed'.
type ConsumeStackRequestAction struct {
	protocol.DestroyStackRequestAction
}

// PlaceInContainerStackRequestAction currently has no known purpose.
type PlaceInContainerStackRequestAction struct {
	protocol.PlaceInContainerStackRequestAction
}

// Marshal ...
func (x *PlaceInContainerStackRequestAction) Marshal(io protocol.IO) {
	io.Uint8(&x.Count)
	StackReqSlotInfo(io, &x.Source)
	StackReqSlotInfo(io, &x.Destination)
}

// TakeOutContainerStackRequestAction currently has no known purpose.
type TakeOutContainerStackRequestAction struct {
	protocol.TakeOutContainerStackRequestAction
}

// Marshal ...
func (x *TakeOutContainerStackRequestAction) Marshal(io protocol.IO) {
	io.Uint8(&x.Count)
	StackReqSlotInfo(io, &x.Source)
	StackReqSlotInfo(io, &x.Destination)
}

// AutoCraftRecipeStackRequestAction is sent by the client similarly to the CraftRecipeStackRequestAction. The
// only difference is that the recipe is automatically created and crafted by shift clicking the recipe book.
type AutoCraftRecipeStackRequestAction struct {
	protocol.AutoCraftRecipeStackRequestAction
}

// Marshal ...
func (x *AutoCraftRecipeStackRequestAction) Marshal(io protocol.IO) {
	io.Varuint32(&x.RecipeNetworkID)
	io.Uint8(&x.TimesCrafted)
}

// StackReqSlotInfo ...
func StackReqSlotInfo(io protocol.IO, x *protocol.StackRequestSlotInfo) {
	if _, ok := io.(interface{ Reads() bool }); !ok && x.Container.ContainerID > 21 {
		x.Container.ContainerID -= 1
	}
	io.Uint8(&x.Container.ContainerID)
	if _, ok := io.(interface{ Reads() bool }); ok && x.Container.ContainerID >= 21 { // RECIPE_BOOK
		x.Container.ContainerID += 1
	}
	io.Uint8(&x.Slot)
	io.Varint32(&x.StackNetworkID)
}
