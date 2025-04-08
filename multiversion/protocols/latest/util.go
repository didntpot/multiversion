package latest

import "github.com/sandertv/gophertunnel/minecraft/nbt"

// itemData ...
func itemData() []byte {
	var items map[string]struct {
		RuntimeID      int32          `nbt:"runtime_id"`
		ComponentBased bool           `nbt:"component_based"`
		Version        int32          `nbt:"version"`
		Data           map[string]any `nbt:"data,omitempty"`
	}
	err := nbt.Unmarshal(itemRuntimeIDData, &items)
	if err != nil {
		panic(err)
	}
	var legacyItems = make(map[string]int32)
	for name, e := range items {
		legacyItems[name] = e.RuntimeID
	}
	data, err := nbt.Marshal(legacyItems)
	if err != nil {
		panic(err)
	}
	return data
}
