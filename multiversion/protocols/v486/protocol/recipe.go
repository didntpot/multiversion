package protocol

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// PotionContainerChangeRecipe represents a recipe to turn a potion from one type to another. This means from
// a drinkable potion + gunpowder -> splash potion, and from a splash potion + dragon breath -> lingering
// potion.
type PotionContainerChangeRecipe struct {
	InputItemID   int32
	ReagentItemID int32
	OutputItemID  int32
}

// PotContainerChangeRecipe ...
func PotContainerChangeRecipe(io protocol.IO, x *PotionContainerChangeRecipe) {
	io.Varint32(&x.InputItemID)
	io.Varint32(&x.ReagentItemID)
	io.Varint32(&x.OutputItemID)
}

// PotionRecipe represents a potion mixing recipe which may be used in a brewing stand.
type PotionRecipe struct {
	InputPotionID        int32
	InputPotionMetadata  int32
	ReagentItemID        int32
	ReagentItemMetadata  int32
	OutputPotionID       int32
	OutputPotionMetadata int32
}

// PotRecipe ...
func PotRecipe(io protocol.IO, x *PotionRecipe) {
	io.Varint32(&x.InputPotionID)
	io.Varint32(&x.InputPotionMetadata)
	io.Varint32(&x.ReagentItemID)
	io.Varint32(&x.ReagentItemMetadata)
	io.Varint32(&x.OutputPotionID)
	io.Varint32(&x.OutputPotionMetadata)
}
