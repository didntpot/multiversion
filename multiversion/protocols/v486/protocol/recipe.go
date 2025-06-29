package protocol

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

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

// RecipeIngredientItem represents an item that may be used as a recipe ingredient.
type RecipeIngredientItem struct {
	NetworkID     int32
	MetadataValue int32
	Count         int32
}

// RecipeIngredient ...
func RecipeIngredient(r protocol.IO, x *RecipeIngredientItem) {
	r.Varint32(&x.NetworkID)
	if x.NetworkID == 0 {
		return
	}
	r.Varint32(&x.MetadataValue)
	r.Varint32(&x.Count)
}

const (
	RecipeShapeless int32 = iota
	RecipeShaped
	RecipeFurnace
	RecipeFurnaceData
	RecipeMulti
	RecipeShulkerBox
	RecipeShapelessChemistry
	RecipeShapedChemistry
)

// Recipe represents a recipe that may be sent in a CraftingData packet to let the client know what recipes
// are available server-side.
type Recipe interface {
	Marshal(io protocol.IO)
}

// ShapelessRecipe is a recipe that has no particular shape. Its functionality is shared with the
// RecipeShulkerBox and RecipeShapelessChemistry types.
type ShapelessRecipe struct {
	RecipeID        string
	Input           []RecipeIngredientItem
	Output          []protocol.ItemStack
	UUID            uuid.UUID
	Block           string
	Priority        int32
	RecipeNetworkID uint32
}

// ShulkerBoxRecipe is a shapeless recipe made specifically for shulker box crafting, so that they don't lose
// their user data when dyeing a shulker box.
type ShulkerBoxRecipe ShapelessRecipe

// ShapelessChemistryRecipe is a recipe specifically made for chemistry related features, which exist only in
// the Education Edition. They function the same as shapeless recipes do.
type ShapelessChemistryRecipe ShapelessRecipe

// ShapedRecipe is a recipe that has a specific shape that must be used to craft the output of the recipe.
// Trying to craft the item in any other shape will not work. The ShapedRecipe is of the same structure as the
// ShapedChemistryRecipe.
type ShapedRecipe struct {
	RecipeID        string
	Width           int32
	Height          int32
	Input           []RecipeIngredientItem
	Output          []protocol.ItemStack
	UUID            uuid.UUID
	Block           string
	Priority        int32
	RecipeNetworkID uint32
}

// ShapedChemistryRecipe is a recipe specifically made for chemistry related features, which exist only in the
// Education Edition. It functions the same as a normal ShapedRecipe.
type ShapedChemistryRecipe ShapedRecipe

// FurnaceRecipe is a recipe that is specifically used for all kinds of furnaces. These recipes don't just
// apply to furnaces, but also blast furnaces and smokers.
type FurnaceRecipe struct {
	InputType protocol.ItemType
	Output    protocol.ItemStack
	Block     string
}

// FurnaceDataRecipe is a recipe specifically used for furnace-type crafting stations. It is equal to
// FurnaceRecipe, except it has an input item with a specific metadata value, instead of any metadata value.
type FurnaceDataRecipe FurnaceRecipe

// MultiRecipe serves as an 'enable' switch for multi-shape recipes.
type MultiRecipe struct {
	UUID            uuid.UUID
	RecipeNetworkID uint32
}

// Marshal ...
func (recipe *ShapelessRecipe) Marshal(io protocol.IO) {
	marshalShapeless(io, recipe)
}

// Marshal ...
func (recipe *ShulkerBoxRecipe) Marshal(io protocol.IO) {
	r := ShapelessRecipe(*recipe)
	marshalShapeless(io, &r)
}

// Marshal ...
func (recipe *ShapelessChemistryRecipe) Marshal(io protocol.IO) {
	r := ShapelessRecipe(*recipe)
	marshalShapeless(io, &r)
}

// Marshal ...
func (recipe *ShapedRecipe) Marshal(io protocol.IO) {
	marshalShaped(io, recipe)
}

// Marshal ...
func (recipe *ShapedChemistryRecipe) Marshal(io protocol.IO) {
	r := ShapedRecipe(*recipe)
	marshalShaped(io, &r)
}

// Marshal ...
func (recipe *FurnaceRecipe) Marshal(io protocol.IO) {
	io.Varint32(&recipe.InputType.NetworkID)
	io.Item(&recipe.Output)
	io.String(&recipe.Block)
}

// Marshal ...
func (recipe *FurnaceDataRecipe) Marshal(io protocol.IO) {
	io.Varint32(&recipe.InputType.NetworkID)
	aux := int32(recipe.InputType.MetadataValue)
	io.Varint32(&aux)
	io.Item(&recipe.Output)
	io.String(&recipe.Block)
}

// Marshal ...
func (recipe *MultiRecipe) Marshal(io protocol.IO) {
	io.UUID(&recipe.UUID)
	io.Varuint32(&recipe.RecipeNetworkID)
}

// marshalShaped ...
func marshalShaped(io protocol.IO, recipe *ShapedRecipe) {
	io.String(&recipe.RecipeID)
	io.Varint32(&recipe.Width)
	io.Varint32(&recipe.Height)

	itemCount := int(recipe.Width * recipe.Height)
	if len(recipe.Input) != itemCount {
		// We got an input count that was not as as big as the full size of the recipe, so we panic as this is
		// a user error.
		panic(fmt.Sprintf("shaped recipe must have exactly %vx%v input items, but got %v", recipe.Width, recipe.Height, len(recipe.Input)))
	}
	for _, input := range recipe.Input {
		RecipeIngredient(io, &input)
	}

	protocol.FuncSlice(io, &recipe.Output, io.Item)
	io.UUID(&recipe.UUID)
	io.String(&recipe.Block)
	io.Varint32(&recipe.Priority)
	io.Varuint32(&recipe.RecipeNetworkID)
}

// marshalShapeless ...
func marshalShapeless(io protocol.IO, recipe *ShapelessRecipe) {
	io.String(&recipe.RecipeID)

	inputLen := uint32(len(recipe.Input))
	io.Varuint32(&inputLen)
	for _, input := range recipe.Input {
		RecipeIngredient(io, &input)
	}

	protocol.FuncSlice(io, &recipe.Output, io.Item)
	io.UUID(&recipe.UUID)
	io.String(&recipe.Block)
	io.Varint32(&recipe.Priority)
	io.Varuint32(&recipe.RecipeNetworkID)
}
