package io

import (
	"fmt"

	legacyprotocol "github.com/didntpot/multiversion/protocols/v486/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Writer ...
type Writer struct {
	*protocol.Writer
}

// NewWriter ...
func NewWriter(w *protocol.Writer) *Writer {
	return &Writer{w}
}

// ItemDescriptorCount ...
func (w *Writer) ItemDescriptorCount(x *protocol.ItemDescriptorCount) {
	var id byte
	switch descriptor := x.Descriptor.(type) {
	case *protocol.InvalidItemDescriptor:
		id = protocol.ItemDescriptorInvalid
	case *protocol.DefaultItemDescriptor:
		id = protocol.ItemDescriptorDefault
	case *protocol.MoLangItemDescriptor:
		id = protocol.ItemDescriptorMoLang
	case *protocol.ItemTagItemDescriptor:
		id = protocol.ItemDescriptorItemTag
	case *protocol.DeferredItemDescriptor:
		id = protocol.ItemDescriptorDeferred
	case *protocol.ComplexAliasItemDescriptor:
		id = protocol.ItemDescriptorComplexAlias
	case *legacyprotocol.DefaultItemDescriptor:
		descriptor.Marshal(w)
		if descriptor.NetworkID != 0 {
			w.Varint32(&x.Count)
		}
		return
	default:
		w.UnknownEnumOption(fmt.Sprintf("%T", x.Descriptor), "item descriptor type")
		return
	}
	w.Uint8(&id)

	x.Descriptor.Marshal(w)
	w.Varint32(&x.Count)
}

// Recipe ...
func (w *Writer) Recipe(x *protocol.Recipe) {
	var recipeType int32
	if !w.lookupRecipeType(*x, &recipeType) {
		w.UnknownEnumOption(fmt.Sprintf("%T", *x), "crafting recipe type")
	}
	w.Varint32(&recipeType)
	(*x).Marshal(w)
}

// lookupRecipeType ...
func (w *Writer) lookupRecipeType(x protocol.Recipe, recipeType *int32) bool {
	switch x.(type) {
	case *protocol.ShapelessRecipe:
		*recipeType = protocol.RecipeShapeless
	case *protocol.ShapedRecipe:
		*recipeType = protocol.RecipeShaped
	case *protocol.FurnaceRecipe:
		*recipeType = protocol.RecipeFurnace
	case *protocol.FurnaceDataRecipe:
		*recipeType = protocol.RecipeFurnaceData
	case *protocol.MultiRecipe:
		*recipeType = protocol.RecipeMulti
	case *protocol.ShulkerBoxRecipe:
		*recipeType = protocol.RecipeShulkerBox
	case *protocol.ShapelessChemistryRecipe:
		*recipeType = protocol.RecipeShapelessChemistry
	case *protocol.ShapedChemistryRecipe:
		*recipeType = protocol.RecipeShapedChemistry
	case *protocol.SmithingTransformRecipe:
		*recipeType = protocol.RecipeSmithingTransform
	case *protocol.SmithingTrimRecipe:
		*recipeType = protocol.RecipeSmithingTrim
	default:
		return false
	}
	return true
}

// TransactionDataType ...
func (w *Writer) TransactionDataType(x *protocol.InventoryTransactionData) {
	var id uint32
	if !w.lookupTransactionDataType(*x, &id) {
		w.UnknownEnumOption(fmt.Sprintf("%T", x), "inventory transaction data type")
	}
	w.Varuint32(&id)
}

// lookupTransactionDataType ...
func (w *Writer) lookupTransactionDataType(x protocol.InventoryTransactionData, id *uint32) bool {
	switch x.(type) {
	case *protocol.NormalTransactionData:
		*id = protocol.InventoryTransactionTypeNormal
	case *protocol.MismatchTransactionData:
		*id = protocol.InventoryTransactionTypeMismatch
	case *legacyprotocol.UseItemTransactionData:
		*id = protocol.InventoryTransactionTypeUseItem
	case *protocol.UseItemOnEntityTransactionData:
		*id = protocol.InventoryTransactionTypeUseItemOnEntity
	case *protocol.ReleaseItemTransactionData:
		*id = protocol.InventoryTransactionTypeReleaseItem
	default:
		return false
	}
	return true
}
