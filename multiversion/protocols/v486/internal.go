package v486

import (
	"math"

	"github.com/didntpot/multiversion/multiversion/mapping"
	legacyprotocol "github.com/didntpot/multiversion/multiversion/protocols/v486/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Internal ...
type Internal struct{}

// downgradeBlockActorData ...
func (Internal) downgradeBlockActorData(data map[string]any) map[string]any {
	switch data["id"] {
	case "Sign":
		delete(data, "BackText")
		frontRaw, ok := data["FrontText"]
		if !ok {
			frontRaw = map[string]any{"Text": ""}
		}
		front, ok := frontRaw.(map[string]any)
		if !ok {
			front = map[string]any{"Text": ""}
		}
		textRaw, ok := front["Text"]
		if !ok {
			textRaw = ""
		}
		text, ok := textRaw.(string)
		if !ok {
			text = ""
		}
		data["Text"] = text
	}
	return data
}

// upgradeBlockActorData ...
func (Internal) upgradeBlockActorData(data map[string]any) map[string]any {
	switch data["id"] {
	case "Sign":
		textRaw, ok := data["Text"]
		if !ok {
			textRaw = ""
		}
		text, ok := textRaw.(string)
		if !ok {
			text = ""
		}
		data["FrontText"] = map[string]any{"Text": text}
		data["BackText"] = map[string]any{"Text": ""}
	}
	return data
}

// downgradeEntityMetadata ...
func (Internal) downgradeEntityMetadata(data map[uint32]any) map[uint32]any {
	var flag1, flag2 int64
	if v, ok := data[protocol.EntityDataKeyFlags]; ok {
		flag1 = v.(int64)
	}
	if v, ok := data[protocol.EntityDataKeyFlagsTwo]; ok {
		flag2 = v.(int64)
	}
	if flag1 == 0 && flag2 == 0 {
		return data
	}

	newFlag1 := flag1 & ^(^0 << (protocol.EntityDataFlagDash - 1))
	lastHalf := flag1 & (^0 << protocol.EntityDataFlagDash)
	lastHalf >>= 1
	lastHalf &= math.MaxInt64

	newFlag1 |= lastHalf

	if flag2 != 0 {
		newFlag1 ^= (flag2 & 1) << 63
		flag2 >>= 1
		flag2 &= math.MaxInt64

		data[protocol.EntityDataKeyFlagsTwo] = flag2
	}

	data[protocol.EntityDataKeyFlags] = newFlag1
	return data
}

// upgradeEntityMetadata ...
func (Internal) upgradeEntityMetadata(data map[uint32]any) map[uint32]any {
	var flag1, flag2 int64
	if v, ok := data[protocol.EntityDataKeyFlags]; ok {
		flag1 = v.(int64)
	}
	if v, ok := data[protocol.EntityDataKeyFlagsTwo]; ok {
		flag2 = v.(int64)
	}

	flag2 <<= 1
	flag2 |= (flag1 >> 63) & 1

	newFlag1 := flag1 & ^(^0 << (protocol.EntityDataFlagDash - 1))
	lastHalf := flag1 & (^0 << (protocol.EntityDataFlagDash - 1))
	lastHalf <<= 1
	newFlag1 |= lastHalf

	data[protocol.EntityDataKeyFlagsTwo] = flag2
	data[protocol.EntityDataKeyFlags] = newFlag1
	return data
}

// downgradeCraftingDescription ...
func (Internal) downgradeCraftingDescription(descriptor protocol.ItemDescriptor, m mapping.Item) (networkId, metadata int32) { // TODO: ..?
	switch descriptor := descriptor.(type) {
	case *protocol.DefaultItemDescriptor:
		networkId = int32(descriptor.NetworkID)
		metadata = int32(descriptor.MetadataValue)
	case *protocol.DeferredItemDescriptor:
		if rid, ok := m.ItemNameToRuntimeID(descriptor.Name); ok {
			networkId = rid
			metadata = int32(descriptor.MetadataValue)
		}
	case *protocol.ItemTagItemDescriptor:
		/// ?????
	case *protocol.ComplexAliasItemDescriptor:
		/// ?????
	}
	return
}

// upgradeCraftingDescription ...
func (Internal) upgradeCraftingDescription(descriptor *legacyprotocol.DefaultItemDescriptor) protocol.ItemDescriptor { // TODO: ..?
	return &protocol.DefaultItemDescriptor{
		NetworkID:     int16(descriptor.NetworkID),
		MetadataValue: int16(descriptor.MetadataValue),
	}
}
