package packet

import (
	"fmt"

	legacyprotocol "github.com/didntpot/multiversion/multiversion/protocols/v486/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// CraftingData is sent by the server to let the client know all crafting data that the server maintains. This
// includes shapeless crafting, crafting table recipes, furnace recipes etc. Each crafting station's recipes
// are included in it.
type CraftingData struct {
	Recipes                      []protocol.Recipe
	PotionRecipes                []protocol.PotionRecipe
	PotionContainerChangeRecipes []protocol.PotionContainerChangeRecipe
	MaterialReducers             []protocol.MaterialReducer
	ClearRecipes                 bool
}

// ID ...
func (*CraftingData) ID() uint32 {
	return IDCraftingData
}

// Marshal ...
func (pk *CraftingData) Marshal(io protocol.IO) {
	l, potRecipesLen, containerRecipesLen, materialReducersLen := uint32(len(pk.Recipes)), uint32(len(pk.PotionRecipes)), uint32(len(pk.PotionContainerChangeRecipes)), uint32(len(pk.MaterialReducers))
	io.Varuint32(&l)
	for _, recipe := range pk.Recipes {
		var c int32
		switch recipe.(type) {
		case *protocol.ShapelessRecipe:
			c = protocol.RecipeShapeless
		case *protocol.ShapedRecipe:
			c = protocol.RecipeShaped
		case *protocol.FurnaceRecipe:
			c = protocol.RecipeFurnace
		case *protocol.FurnaceDataRecipe:
			c = protocol.RecipeFurnaceData
		case *protocol.MultiRecipe:
			c = protocol.RecipeMulti
		case *protocol.ShulkerBoxRecipe:
			c = protocol.RecipeShulkerBox
		case *protocol.ShapelessChemistryRecipe:
			c = protocol.RecipeShapelessChemistry
		case *protocol.ShapedChemistryRecipe:
			c = protocol.RecipeShapedChemistry
		default:
			io.UnknownEnumOption(fmt.Sprintf("%T", recipe), "crafting recipe type")
		}
		io.Varint32(&c)
		recipe.Marshal(io)
	}
	io.Varuint32(&potRecipesLen)
	for _, mix := range pk.PotionRecipes {
		legacyprotocol.PotRecipe(io, (*legacyprotocol.PotionRecipe)(&mix))
	}
	io.Varuint32(&containerRecipesLen)
	for _, mix := range pk.PotionContainerChangeRecipes {
		legacyprotocol.PotContainerChangeRecipe(io, (*legacyprotocol.PotionContainerChangeRecipe)(&mix))
	}
	io.Varuint32(&materialReducersLen)
	for _, mat := range pk.MaterialReducers {
		io.MaterialReducer(&mat)
	}

	io.Bool(&pk.ClearRecipes)
}
