package v486

import (
	"bytes"

	"github.com/didntpot/multiversion/multiversion/internal/chunk"
	"github.com/didntpot/multiversion/multiversion/mapping"
	"github.com/didntpot/multiversion/multiversion/mapping/translator"
	"github.com/didntpot/multiversion/multiversion/protocols/latest"
	legacyprotocol "github.com/didntpot/multiversion/multiversion/protocols/v486/protocol"
	legacyio "github.com/didntpot/multiversion/multiversion/protocols/v486/protocol/io"
	legacypacket "github.com/didntpot/multiversion/multiversion/protocols/v486/protocol/packet"
	"github.com/samber/lo"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// Protocol ...
type Protocol struct {
	itemMapping  mapping.Item
	blockMapping mapping.Block

	itemTranslator  translator.ItemTranslator
	blockTranslator translator.BlockTranslator

	internal *Internal
}

// New ...
func New(direct bool) *Protocol {
	internal := &Internal{}
	itemMapping := mapping.NewItemMapping(itemRuntimeIDData, requiredItemList, ItemVersion, direct)
	blockMapping := mapping.NewBlockMapping(blockStateData).WithBlockActorRemapper(internal.downgradeBlockActorData, internal.upgradeBlockActorData)
	latestBlockMapping := latest.NewBlockMapping()
	return &Protocol{
		itemMapping:  itemMapping,
		blockMapping: blockMapping,

		itemTranslator: translator.NewItemTranslator(
			itemMapping,
			latest.NewItemMapping(direct),
			blockMapping, latestBlockMapping,
		),
		blockTranslator: translator.NewBlockTranslator(
			blockMapping, latestBlockMapping,
			chunk.NewNetworkPersistentEncoding(blockMapping, BlockVersion),
			chunk.NewBlockPaletteEncoding(blockMapping, BlockVersion),
			false,
		),

		internal: internal,
	}
}

// ID ...
func (Protocol) ID() int32 {
	return ProtocolID
}

// Ver ...
func (Protocol) Ver() string {
	return ProtocolVersion
}

// Encryption ...
func (Protocol) Encryption(key [32]byte) packet.Encryption {
	return packet.NewCTREncryption(key[:])
}

// NewReader ...
func (Protocol) NewReader(r minecraft.ByteReader, shieldID int32, enableLimits bool) protocol.IO {
	return legacyio.NewReader(protocol.NewReader(r, shieldID, enableLimits))
}

// NewWriter ...
func (Protocol) NewWriter(w minecraft.ByteWriter, shieldID int32) protocol.IO {
	return legacyio.NewWriter(protocol.NewWriter(w, shieldID))
}

// Packets ...
func (Protocol) Packets(bool) packet.Pool {
	pool := packet.NewClientPool()
	for k, v := range packet.NewServerPool() {
		pool[k] = v
	}
	packetMap := map[uint32]packet.Packet{
		packet.IDChangeDimension:      &legacypacket.ChangeDimension{},
		packet.IDCommandRequest:       &legacypacket.CommandRequest{},
		packet.IDContainerClose:       &legacypacket.ContainerClose{},
		packet.IDDisconnect:           &legacypacket.Disconnect{},
		packet.IDEmote:                &legacypacket.Emote{},
		packet.IDInventoryContent:     &legacypacket.InventoryContent{},
		packet.IDInventorySlot:        &legacypacket.InventorySlot{},
		packet.IDInventoryTransaction: &legacypacket.InventoryTransaction{},
		packet.IDItemStackRequest:     &legacypacket.ItemStackRequest{},
		packet.IDItemStackResponse:    &legacypacket.ItemStackResponse{},
		packet.IDLevelSoundEvent:      &legacypacket.LevelSoundEvent{},
		packet.IDMobArmourEquipment:   &legacypacket.MobArmourEquipment{},
		packet.IDModalFormResponse:    &legacypacket.ModalFormResponse{},
		packet.IDPlayerAction:         &legacypacket.PlayerAction{},
		packet.IDPlayerArmourDamage:   &legacypacket.PlayerArmourDamage{},
		packet.IDPlayerAuthInput:      &legacypacket.PlayerAuthInput{},
		packet.IDPlayerSkin:           &legacypacket.PlayerSkin{},
		packet.IDRequestChunkRadius:   &legacypacket.RequestChunkRadius{},
		packet.IDSetTitle:             &legacypacket.SetTitle{},
		packet.IDStopSound:            &legacypacket.StopSound{},
		packet.IDText:                 &legacypacket.Text{},
		legacypacket.IDTickSync:       &legacypacket.TickSync{},
	}
	for id, pk := range packetMap {
		pool[id] = func() packet.Packet { return pk }
	}
	return pool
}

// ConvertToLatest ...
func (p Protocol) ConvertToLatest(pk packet.Packet, conn *minecraft.Conn) []packet.Packet {
	var newPks []packet.Packet
	switch pk := pk.(type) {
	case *packet.ClientCacheStatus:
		pk.Enabled = false
		newPks = append(newPks, pk)
	case *legacypacket.CommandRequest:
		newPks = append(newPks, &packet.CommandRequest{
			CommandLine:   pk.CommandLine,
			CommandOrigin: pk.CommandOrigin,
			Internal:      pk.Internal,
		})
	case *legacypacket.ContainerClose:
		newPks = append(newPks, &packet.ContainerClose{
			WindowID:   pk.WindowID,
			ServerSide: pk.ServerSide,
		})
	case *legacypacket.Disconnect:
		newPks = append(newPks, &packet.Disconnect{
			HideDisconnectionScreen: pk.HideDisconnectionScreen,
			Message:                 pk.Message,
		})
	case *legacypacket.Emote:
		newPks = append(newPks, &packet.Emote{
			EntityRuntimeID: pk.EntityRuntimeID,
			EmoteID:         pk.EmoteID,
			XUID:            conn.IdentityData().XUID,
			PlatformID:      conn.ClientData().PlatformOnlineID,
			Flags:           pk.Flags,
		})
	case *legacypacket.InventoryTransaction:
		transactionData := pk.TransactionData
		if useItemData, ok := transactionData.(*legacyprotocol.UseItemTransactionData); ok {
			transactionData = &protocol.UseItemTransactionData{
				LegacyRequestID:    useItemData.LegacyRequestID,
				LegacySetItemSlots: useItemData.LegacySetItemSlots,
				Actions:            useItemData.Actions,
				ActionType:         useItemData.ActionType,
				BlockPosition:      useItemData.BlockPosition,
				BlockFace:          useItemData.BlockFace,
				HotBarSlot:         useItemData.HotBarSlot,
				HeldItem:           useItemData.HeldItem,
				Position:           useItemData.Position,
				ClickedPosition:    useItemData.ClickedPosition,
				BlockRuntimeID:     useItemData.BlockRuntimeID,
			}
		}
		newPks = append(newPks, &packet.InventoryTransaction{
			LegacyRequestID: pk.LegacyRequestID,
			LegacySetItemSlots: lo.Map(pk.LegacySetItemSlots, func(item protocol.LegacySetItemSlot, _ int) protocol.LegacySetItemSlot {
				if item.ContainerID >= 21 { // RECIPE_BOOK
					item.ContainerID += 1
				}
				return item
			}),
			Actions:         pk.Actions,
			TransactionData: transactionData,
		})
	case *legacypacket.ItemStackRequest:
		newPks = append(newPks, &packet.ItemStackRequest{
			Requests: lo.Map(pk.Requests, func(item legacyprotocol.ItemStackRequest, _ int) protocol.ItemStackRequest {
				return protocol.ItemStackRequest{
					RequestID: item.RequestID,
					Actions: lo.Map(item.Actions, func(item protocol.StackRequestAction, _ int) protocol.StackRequestAction {
						switch action := item.(type) {
						case *legacyprotocol.AutoCraftRecipeStackRequestAction:
							return &action.AutoCraftRecipeStackRequestAction
						case *legacyprotocol.ConsumeStackRequestAction:
							return &action.DestroyStackRequestAction
						case *legacyprotocol.CraftCreativeStackRequestAction:
							return &action.CraftCreativeStackRequestAction
						case *legacyprotocol.CraftRecipeStackRequestAction:
							return &action.CraftRecipeStackRequestAction
						case *legacyprotocol.DestroyStackRequestAction:
							return &action.DestroyStackRequestAction
						case *legacyprotocol.DropStackRequestAction:
							return &action.DropStackRequestAction
						case *legacyprotocol.PlaceInContainerStackRequestAction:
							return &action.PlaceInContainerStackRequestAction
						case *legacyprotocol.PlaceStackRequestAction:
							return &action.PlaceStackRequestAction
						case *legacyprotocol.SwapStackRequestAction:
							return &action.SwapStackRequestAction
						case *legacyprotocol.TakeOutContainerStackRequestAction:
							return &action.TakeOutContainerStackRequestAction
						case *legacyprotocol.TakeStackRequestAction:
							return &action.TakeStackRequestAction
						}
						return item
					}),
					FilterStrings: item.FilterStrings,
				}
			}),
		})
	case *legacypacket.LevelSoundEvent:
		return nil
	case *legacypacket.ModalFormResponse:
		responseData := protocol.Optional[[]byte]{}
		cancelReason := protocol.Optional[uint8]{}
		if bytes.Equal(pk.ResponseData, []byte{110, 117, 108, 108, 10}) {
			cancelReason = protocol.Option(uint8(packet.ModalFormCancelReasonUserClosed))
		} else {
			responseData = protocol.Option(pk.ResponseData)
		}
		newPks = append(newPks, &packet.ModalFormResponse{
			FormID:       pk.FormID,
			ResponseData: responseData,
			CancelReason: cancelReason,
		})
	case *legacypacket.PlayerAction:
		newPks = append(newPks, &packet.PlayerAction{
			EntityRuntimeID: pk.EntityRuntimeID,
			ActionType:      pk.ActionType,
			BlockPosition:   pk.BlockPosition,
			ResultPosition:  pk.BlockPosition,
			BlockFace:       pk.BlockFace,
		})
	case *legacypacket.PlayerAuthInput:
		newPks = append(newPks, &packet.PlayerAuthInput{
			Pitch:      pk.Pitch,
			Yaw:        pk.Yaw,
			Position:   pk.Position,
			MoveVector: pk.MoveVector,
			HeadYaw:    pk.HeadYaw,
			InputData:  legacyprotocol.BitSet(pk.InputData, packet.PlayerAuthInputBitsetSize),
			InputMode:  pk.InputMode,
			PlayMode:   pk.PlayMode,
			Tick:       pk.Tick,
			Delta:      pk.Delta,
			ItemInteractionData: func(data protocol.UseItemTransactionData) protocol.UseItemTransactionData {
				data.LegacySetItemSlots = lo.Map(data.LegacySetItemSlots, func(item protocol.LegacySetItemSlot, _ int) protocol.LegacySetItemSlot {
					if item.ContainerID >= 21 { // RECIPE_BOOK
						item.ContainerID += 1
					}
					return item
				})
				return data
			}(pk.ItemInteractionData),
			ItemStackRequest: protocol.ItemStackRequest{
				RequestID: pk.ItemStackRequest.RequestID,
				Actions: lo.Map(pk.ItemStackRequest.Actions, func(item protocol.StackRequestAction, _ int) protocol.StackRequestAction {
					switch action := item.(type) {
					case *legacyprotocol.AutoCraftRecipeStackRequestAction:
						return &action.AutoCraftRecipeStackRequestAction
					case *legacyprotocol.ConsumeStackRequestAction:
						return &action.DestroyStackRequestAction
					case *legacyprotocol.CraftCreativeStackRequestAction:
						return &action.CraftCreativeStackRequestAction
					case *legacyprotocol.CraftRecipeStackRequestAction:
						return &action.CraftRecipeStackRequestAction
					case *legacyprotocol.DestroyStackRequestAction:
						return &action.DestroyStackRequestAction
					case *legacyprotocol.DropStackRequestAction:
						return &action.DropStackRequestAction
					case *legacyprotocol.PlaceInContainerStackRequestAction:
						return &action.PlaceInContainerStackRequestAction
					case *legacyprotocol.PlaceStackRequestAction:
						return &action.PlaceStackRequestAction
					case *legacyprotocol.SwapStackRequestAction:
						return &action.SwapStackRequestAction
					case *legacyprotocol.TakeOutContainerStackRequestAction:
						return &action.TakeOutContainerStackRequestAction
					case *legacyprotocol.TakeStackRequestAction:
						return &action.TakeStackRequestAction
					}
					return item
				}),
				FilterStrings: pk.ItemStackRequest.FilterStrings,
			},
			BlockActions:       pk.BlockActions,
			AnalogueMoveVector: pk.MoveVector,
		})
	case *legacypacket.PlayerSkin:
		newPks = append(newPks, &packet.PlayerSkin{
			UUID:        pk.UUID,
			Skin:        pk.Skin.Skin,
			NewSkinName: pk.NewSkinName,
			OldSkinName: pk.OldSkinName,
		})
	case *legacypacket.RequestChunkRadius:
		newPks = append(newPks, &packet.RequestChunkRadius{
			ChunkRadius:    pk.ChunkRadius,
			MaxChunkRadius: pk.ChunkRadius,
		})
	case *legacypacket.Text:
		newPks = append(newPks, &packet.Text{
			TextType:         pk.TextType,
			NeedsTranslation: pk.NeedsTranslation,
			SourceName:       pk.SourceName,
			Message:          pk.Message,
			Parameters:       pk.Parameters,
			XUID:             pk.XUID,
			PlatformChatID:   pk.PlatformChatID,
		})
	case *legacypacket.TickSync:
		return nil
	default:
		newPks = append(newPks, pk)
	}
	return p.blockTranslator.UpgradeBlockPackets(p.itemTranslator.UpgradeItemPackets(newPks, conn), conn)
}

// ConvertFromLatest ...
func (p Protocol) ConvertFromLatest(pk packet.Packet, conn *minecraft.Conn) (result []packet.Packet) {
	result = p.blockTranslator.DowngradeBlockPackets(p.itemTranslator.DowngradeItemPackets([]packet.Packet{pk}, conn), conn)
	for i, pk := range result {
		switch pk := pk.(type) {
		case *packet.AddActor:
			result[i] = &legacypacket.AddActor{
				EntityMetadata:  p.internal.downgradeEntityMetadata(pk.EntityMetadata),
				EntityRuntimeID: pk.EntityRuntimeID,
				EntityType:      pk.EntityType,
				EntityUniqueID:  pk.EntityUniqueID,
				HeadYaw:         pk.HeadYaw,
				Pitch:           pk.Pitch,
				Position:        pk.Position,
				Velocity:        pk.Velocity,
				Yaw:             pk.Yaw,
				Attributes:      pk.Attributes,
				EntityLinks: lo.Map(pk.EntityLinks, func(i protocol.EntityLink, _ int) legacyprotocol.EntityLink {
					return legacyprotocol.EntityLink{EntityLink: i}
				}),
			}
		case *packet.AddItemActor:
			result[i] = &packet.AddItemActor{
				EntityMetadata:  p.internal.downgradeEntityMetadata(pk.EntityMetadata),
				EntityRuntimeID: pk.EntityRuntimeID,
				EntityUniqueID:  pk.EntityUniqueID,
				FromFishing:     pk.FromFishing,
				Item:            pk.Item,
				Position:        pk.Position,
				Velocity:        pk.Velocity,
			}
		case *packet.AddPlayer:
			result[i] = &legacypacket.AddPlayer{
				UUID:            pk.UUID,
				Username:        pk.Username,
				EntityUniqueID:  pk.AbilityData.EntityUniqueID,
				EntityRuntimeID: pk.EntityRuntimeID,
				PlatformChatID:  pk.PlatformChatID,
				Position:        pk.Position,
				Velocity:        pk.Velocity,
				Pitch:           pk.Pitch,
				Yaw:             pk.Yaw,
				HeadYaw:         pk.HeadYaw,
				HeldItem:        pk.HeldItem,
				EntityMetadata:  p.internal.downgradeEntityMetadata(pk.EntityMetadata),
				AdventureSettings: packet.AdventureSettings{
					CommandPermissionLevel: uint32(pk.AbilityData.CommandPermissions),
					PermissionLevel:        uint32(pk.AbilityData.PlayerPermissions),
					PlayerUniqueID:         pk.AbilityData.EntityUniqueID,
				},
				DeviceID:    pk.DeviceID,
				EntityLinks: pk.EntityLinks,
			}
		case *packet.AddVolumeEntity:
			result[i] = &legacypacket.AddVolumeEntity{
				EntityRuntimeID:    pk.EntityRuntimeID,
				EntityMetadata:     pk.EntityMetadata,
				EncodingIdentifier: pk.EncodingIdentifier,
				InstanceIdentifier: pk.InstanceIdentifier,
				EngineVersion:      pk.EngineVersion,
			}
		case *packet.AvailableActorIdentifiers:
			result[i] = &packet.AvailableActorIdentifiers{
				SerialisedEntityIdentifiers: entityIdentifierData,
			}
		case *packet.AvailableCommands:
			for ind1, command := range pk.Commands {
				for ind2, overload := range command.Overloads {
					for ind3, parameter := range overload.Parameters {
						parameterType := uint32(parameter.Type) | protocol.CommandArgValid
						switch parameter.Type | protocol.CommandArgValid {
						case protocol.CommandArgTypeBlockPosition, protocol.CommandArgTypePosition:
							parameterType = legacyprotocol.CommandArgTypePosition
						case protocol.CommandArgTypeCommand:
							parameterType = legacyprotocol.CommandArgTypeCommand
						case protocol.CommandArgTypeCompareOperator:
							parameterType = legacyprotocol.CommandArgTypeOperator
						case protocol.CommandArgTypeFilepath:
							parameterType = legacyprotocol.CommandArgTypeFilepath
						case protocol.CommandArgTypeJSON:
							parameterType = legacyprotocol.CommandArgTypeJSON
						case protocol.CommandArgTypeMessage:
							parameterType = legacyprotocol.CommandArgTypeMessage
						case protocol.CommandArgTypeRawText:
							parameterType = legacyprotocol.CommandArgTypeRawText
						case protocol.CommandArgTypeString:
							parameterType = legacyprotocol.CommandArgTypeString
						case protocol.CommandArgTypeTarget:
							parameterType = legacyprotocol.CommandArgTypeTarget
						case protocol.CommandArgTypeWildcardTarget:
							parameterType = legacyprotocol.CommandArgTypeWildcardTarget
						}
						parameter.Type = parameterType | protocol.CommandArgValid
						pk.Commands[ind1].Overloads[ind2].Parameters[ind3] = parameter
					}
				}
			}
			result[i] = &legacypacket.AvailableCommands{
				EnumValues: pk.EnumValues,
				Suffixes:   pk.Suffixes,
				Enums:      pk.Enums,
				Commands: lo.Map(pk.Commands, func(item protocol.Command, _ int) legacyprotocol.Command {
					return legacyprotocol.Command{
						Name:            item.Name,
						Description:     item.Description,
						Flags:           item.Flags,
						PermissionLevel: item.PermissionLevel,
						AliasesOffset:   item.AliasesOffset,
						Overloads: lo.Map(item.Overloads, func(item protocol.CommandOverload, _ int) legacyprotocol.CommandOverload {
							return legacyprotocol.CommandOverload{
								Parameters: item.Parameters,
							}
						}),
					}
				}),
				DynamicEnums: pk.DynamicEnums,
				Constraints:  pk.Constraints,
			}
		case *packet.BiomeDefinitionList:
			return []packet.Packet{
				&legacypacket.BiomeDefinitionList{SerialisedBiomeDefinitions: legacySerialisedBiomeDefinitions},
			}
		case *packet.BlockActorData:
			pk.NBTData = p.internal.downgradeBlockActorData(pk.NBTData)
			result[i] = pk
		case *packet.BossEvent:
			result[i] = &legacypacket.BossEvent{
				BossEntityUniqueID: pk.BossEntityUniqueID,
				EventType:          pk.EventType,
				PlayerUniqueID:     pk.PlayerUniqueID,
				BossBarTitle:       pk.BossBarTitle,
				HealthPercentage:   pk.HealthPercentage,
				ScreenDarkening:    int16(pk.ScreenDarkening),
				Colour:             pk.Colour,
				Overlay:            pk.Overlay,
			}
		case *packet.ContainerClose:
			result[i] = &legacypacket.ContainerClose{
				WindowID:   pk.WindowID,
				ServerSide: pk.ServerSide,
			}
		case *packet.CraftingData:
			recipes := make([]legacyprotocol.Recipe, 0, len(pk.Recipes))
			for _, recipe := range pk.Recipes {
				switch recipe := recipe.(type) {
				case *protocol.ShapedChemistryRecipe:
					recipes = append(recipes, &legacyprotocol.ShapedChemistryRecipe{
						RecipeID: recipe.RecipeID,
						Width:    recipe.Width,
						Height:   recipe.Height,
						Input: lo.Map(recipe.Input, func(item protocol.ItemDescriptorCount, _ int) legacyprotocol.RecipeIngredientItem {
							networkId, metadata := p.internal.downgradeCraftingDescription(item.Descriptor, p.itemMapping)
							return legacyprotocol.RecipeIngredientItem{
								NetworkID:     networkId,
								MetadataValue: metadata,
								Count:         item.Count,
							}
						}),
						Output:          recipe.Output,
						UUID:            recipe.UUID,
						Block:           recipe.Block,
						Priority:        recipe.Priority,
						RecipeNetworkID: recipe.RecipeNetworkID,
					})
				case *protocol.ShapedRecipe:
					recipes = append(recipes, &legacyprotocol.ShapedRecipe{
						RecipeID: recipe.RecipeID,
						Width:    recipe.Width,
						Height:   recipe.Height,
						Input: lo.Map(recipe.Input, func(item protocol.ItemDescriptorCount, _ int) legacyprotocol.RecipeIngredientItem {
							networkId, metadata := p.internal.downgradeCraftingDescription(item.Descriptor, p.itemMapping)
							return legacyprotocol.RecipeIngredientItem{
								NetworkID:     networkId,
								MetadataValue: metadata,
								Count:         item.Count,
							}
						}),
						Output:          recipe.Output,
						UUID:            recipe.UUID,
						Block:           recipe.Block,
						Priority:        recipe.Priority,
						RecipeNetworkID: recipe.RecipeNetworkID,
					})
				case *protocol.ShapelessChemistryRecipe:
					recipes = append(recipes, &legacyprotocol.ShapelessChemistryRecipe{
						RecipeID: recipe.RecipeID,
						Input: lo.Map(recipe.Input, func(item protocol.ItemDescriptorCount, _ int) legacyprotocol.RecipeIngredientItem {
							networkId, metadata := p.internal.downgradeCraftingDescription(item.Descriptor, p.itemMapping)
							return legacyprotocol.RecipeIngredientItem{
								NetworkID:     networkId,
								MetadataValue: metadata,
								Count:         item.Count,
							}
						}),
						Output:          recipe.Output,
						UUID:            recipe.UUID,
						Block:           recipe.Block,
						Priority:        recipe.Priority,
						RecipeNetworkID: recipe.RecipeNetworkID,
					})
				case *protocol.ShapelessRecipe:
					recipes = append(recipes, &legacyprotocol.ShapelessRecipe{
						RecipeID: recipe.RecipeID,
						Input: lo.Map(recipe.Input, func(item protocol.ItemDescriptorCount, _ int) legacyprotocol.RecipeIngredientItem {
							networkId, metadata := p.internal.downgradeCraftingDescription(item.Descriptor, p.itemMapping)
							return legacyprotocol.RecipeIngredientItem{
								NetworkID:     networkId,
								MetadataValue: metadata,
								Count:         item.Count,
							}
						}),
						Output:          recipe.Output,
						UUID:            recipe.UUID,
						Block:           recipe.Block,
						Priority:        recipe.Priority,
						RecipeNetworkID: recipe.RecipeNetworkID,
					})
				case *protocol.ShulkerBoxRecipe:
					recipes = append(recipes, &legacyprotocol.ShulkerBoxRecipe{
						RecipeID: recipe.RecipeID,
						Input: lo.Map(recipe.Input, func(item protocol.ItemDescriptorCount, _ int) legacyprotocol.RecipeIngredientItem {
							networkId, metadata := p.internal.downgradeCraftingDescription(item.Descriptor, p.itemMapping)
							return legacyprotocol.RecipeIngredientItem{
								NetworkID:     networkId,
								MetadataValue: metadata,
								Count:         item.Count,
							}
						}),
						Output:          recipe.Output,
						UUID:            recipe.UUID,
						Block:           recipe.Block,
						Priority:        recipe.Priority,
						RecipeNetworkID: recipe.RecipeNetworkID,
					})
				case *protocol.SmithingTransformRecipe, *protocol.SmithingTrimRecipe:
					// Just ignore these entries completely.
				}
			}
			result[i] = &legacypacket.CraftingData{
				Recipes:                      recipes,
				PotionRecipes:                pk.PotionRecipes,
				PotionContainerChangeRecipes: pk.PotionContainerChangeRecipes,
				MaterialReducers:             pk.MaterialReducers,
				ClearRecipes:                 pk.ClearRecipes,
			}
		case *packet.CreativeContent:
			result[i] = &legacypacket.CreativeContent{
				Items: lo.Map(pk.Items, func(it protocol.CreativeItem, _ int) legacyprotocol.CreativeItem {
					return legacyprotocol.CreativeItem{
						CreativeItemNetworkID: it.CreativeItemNetworkID,
						Item:                  it.Item,
					}
				}),
			}
			continue
		case *packet.Disconnect:
			result[i] = &legacypacket.Disconnect{
				HideDisconnectionScreen: pk.HideDisconnectionScreen,
				Message:                 pk.Message,
			}
		case *packet.Emote:
			result[i] = &legacypacket.Emote{
				EntityRuntimeID: pk.EntityRuntimeID,
				EmoteID:         pk.EmoteID,
				Flags:           pk.Flags,
			}
		case *packet.InventoryContent:
			result[i] = &legacypacket.InventoryContent{
				WindowID: pk.WindowID,
				Content:  pk.Content,
			}
		case *packet.InventorySlot:
			result[i] = &legacypacket.InventorySlot{
				WindowID: pk.WindowID,
				Slot:     pk.Slot,
				NewItem:  pk.NewItem,
			}
		case *packet.InventoryTransaction:
			result[i] = &legacypacket.InventoryTransaction{
				LegacyRequestID: pk.LegacyRequestID,
				LegacySetItemSlots: lo.Map(pk.LegacySetItemSlots, func(item protocol.LegacySetItemSlot, _ int) protocol.LegacySetItemSlot {
					if item.ContainerID > 21 { // RECIPE_BOOK
						item.ContainerID -= 1
					}
					return item
				}),
				Actions:         pk.Actions,
				TransactionData: pk.TransactionData,
			}
		case *packet.ItemStackResponse:
			result[i] = &legacypacket.ItemStackResponse{
				Responses: lo.Map(pk.Responses, func(response protocol.ItemStackResponse, _ int) legacyprotocol.ItemStackResponse {
					return legacyprotocol.ItemStackResponse{
						Status:    response.Status,
						RequestID: response.RequestID,
						ContainerInfo: lo.Map(response.ContainerInfo, func(info protocol.StackResponseContainerInfo, _ int) legacyprotocol.StackResponseContainerInfo {
							if info.Container.ContainerID > 21 { // RECIPE_BOOK
								info.Container.ContainerID -= 1
							}
							return legacyprotocol.StackResponseContainerInfo{
								ContainerID: info.Container.ContainerID,
								SlotInfo: lo.Map(info.SlotInfo, func(slot protocol.StackResponseSlotInfo, _ int) legacyprotocol.StackResponseSlotInfo {
									return legacyprotocol.StackResponseSlotInfo{StackResponseSlotInfo: slot}
								}),
							}
						}),
					}
				}),
			}
		case *packet.LevelChunk:
			result[i] = &legacypacket.LevelChunk{
				Position:        pk.Position,
				HighestSubChunk: pk.HighestSubChunk,
				SubChunkCount:   pk.SubChunkCount,
				CacheEnabled:    pk.CacheEnabled,
				BlobHashes:      pk.BlobHashes,
				RawPayload:      pk.RawPayload,
			}
		case *packet.LevelSoundEvent:
			result[i] = &legacypacket.LevelSoundEvent{
				SoundType:             pk.SoundType,
				Position:              pk.Position,
				ExtraData:             pk.ExtraData,
				EntityType:            pk.EntityType,
				BabyMob:               pk.BabyMob,
				DisableRelativeVolume: pk.DisableRelativeVolume,
			}
		case *packet.MobArmourEquipment:
			result[i] = &legacypacket.MobArmourEquipment{
				EntityRuntimeID: pk.EntityRuntimeID,
				Helmet:          pk.Helmet,
				Chestplate:      pk.Chestplate,
				Leggings:        pk.Leggings,
				Boots:           pk.Boots,
			}
		case *packet.MobEffect:
			result[i] = &legacypacket.MobEffect{
				EntityRuntimeID: pk.EntityRuntimeID,
				Operation:       pk.Operation,
				EffectType:      pk.EffectType,
				Amplifier:       pk.Amplifier,
				Particles:       pk.Particles,
				Duration:        pk.Duration,
			}
		case *packet.NetworkChunkPublisherUpdate:
			result[i] = &legacypacket.NetworkChunkPublisherUpdate{
				Position: pk.Position,
				Radius:   pk.Radius,
			}
		case *packet.PlayerArmourDamage:
			result[i] = &legacypacket.PlayerArmourDamage{
				Bitset:           pk.Bitset,
				HelmetDamage:     pk.HelmetDamage,
				ChestplateDamage: pk.ChestplateDamage,
				LeggingsDamage:   pk.LeggingsDamage,
				BootsDamage:      pk.BootsDamage,
			}
		case *packet.PlayerList:
			result[i] = &legacypacket.PlayerList{
				ActionType: pk.ActionType,
				Entries: lo.Map(pk.Entries, func(item protocol.PlayerListEntry, _ int) legacyprotocol.PlayerListEntry {
					return legacyprotocol.PlayerListEntry{PlayerListEntry: item}
				}),
			}
		case *packet.PlayerSkin:
			result[i] = &legacypacket.PlayerSkin{
				UUID:        pk.UUID,
				Skin:        legacyprotocol.Skin{Skin: pk.Skin},
				NewSkinName: pk.NewSkinName,
				OldSkinName: pk.OldSkinName,
			}
		case *packet.RemoveVolumeEntity:
			result[i] = &legacypacket.RemoveVolumeEntity{
				EntityRuntimeID: pk.EntityRuntimeID,
			}
		case *packet.ResourcePackStack:
			result[i] = &legacypacket.ResourcePackStack{
				TexturePackRequired:          pk.TexturePackRequired,
				BehaviourPacks:               pk.BehaviourPacks,
				TexturePacks:                 pk.TexturePacks,
				BaseGameVersion:              pk.BaseGameVersion,
				Experiments:                  pk.Experiments,
				ExperimentsPreviouslyToggled: pk.ExperimentsPreviouslyToggled,
			}
		case *packet.ResourcePacksInfo:
			result[i] = &legacypacket.ResourcePacksInfo{
				TexturePackRequired: pk.TexturePackRequired,
				HasScripts:          pk.HasScripts,
				TexturePacks: lo.Map(pk.TexturePacks, func(pack protocol.TexturePackInfo, _ int) legacyprotocol.TexturePackInfo {
					return legacyprotocol.TexturePackInfo{
						UUID:            pack.UUID.String(),
						Version:         pack.Version,
						Size:            pack.Size,
						ContentKey:      pack.ContentKey,
						SubPackName:     pack.SubPackName,
						ContentIdentity: pack.ContentIdentity,
						HasScripts:      pack.HasScripts,
						RTXEnabled:      pack.RTXEnabled,
					}
				}),
			}
		case *packet.SetActorData:
			result[i] = &legacypacket.SetActorData{
				EntityRuntimeID: pk.EntityRuntimeID,
				EntityMetadata:  p.internal.downgradeEntityMetadata(pk.EntityMetadata),
				Tick:            pk.Tick,
			}
		case *packet.SetActorLink:
			result[i] = &legacypacket.SetActorLink{
				EntityLink: legacyprotocol.EntityLink{EntityLink: pk.EntityLink},
			}
		case *packet.SetActorMotion:
			result[i] = &legacypacket.SetActorMotion{
				EntityRuntimeID: pk.EntityRuntimeID,
				Velocity:        pk.Velocity,
			}
		case *packet.SetTitle:
			result[i] = &legacypacket.SetTitle{
				ActionType:      pk.ActionType,
				Text:            pk.Text,
				FadeInDuration:  pk.FadeInDuration,
				RemainDuration:  pk.RemainDuration,
				FadeOutDuration: pk.FadeOutDuration,
			}
		case *packet.SpawnParticleEffect:
			result[i] = &legacypacket.SpawnParticleEffect{
				Dimension:      pk.Dimension,
				EntityUniqueID: pk.EntityUniqueID,
				Position:       pk.Position,
				ParticleName:   pk.ParticleName,
			}
		case *packet.StartGame:
			_, enabled := pk.ForceExperimentalGameplay.Value()
			result[i] = &legacypacket.StartGame{
				EntityUniqueID:                 pk.EntityUniqueID,
				EntityRuntimeID:                pk.EntityRuntimeID,
				PlayerGameMode:                 pk.PlayerGameMode,
				PlayerPosition:                 pk.PlayerPosition,
				Pitch:                          pk.Pitch,
				Yaw:                            pk.Yaw,
				WorldSeed:                      int32(pk.WorldSeed),
				SpawnBiomeType:                 pk.SpawnBiomeType,
				UserDefinedBiomeName:           pk.UserDefinedBiomeName,
				Dimension:                      pk.Dimension,
				Generator:                      pk.Generator,
				WorldGameMode:                  pk.WorldGameMode,
				Difficulty:                     pk.Difficulty,
				WorldSpawn:                     pk.WorldSpawn,
				AchievementsDisabled:           pk.AchievementsDisabled,
				DayCycleLockTime:               pk.DayCycleLockTime,
				EducationEditionOffer:          pk.EducationEditionOffer,
				EducationFeaturesEnabled:       pk.EducationFeaturesEnabled,
				EducationProductID:             pk.EducationProductID,
				RainLevel:                      pk.RainLevel,
				LightningLevel:                 pk.LightningLevel,
				ConfirmedPlatformLockedContent: pk.ConfirmedPlatformLockedContent,
				MultiPlayerGame:                pk.MultiPlayerGame,
				LANBroadcastEnabled:            pk.LANBroadcastEnabled,
				XBLBroadcastMode:               pk.XBLBroadcastMode,
				PlatformBroadcastMode:          pk.PlatformBroadcastMode,
				CommandsEnabled:                pk.CommandsEnabled,
				TexturePackRequired:            pk.TexturePackRequired,
				GameRules:                      pk.GameRules,
				Experiments:                    pk.Experiments,
				ExperimentsPreviouslyToggled:   pk.ExperimentsPreviouslyToggled,
				BonusChestEnabled:              pk.BonusChestEnabled,
				StartWithMapEnabled:            pk.StartWithMapEnabled,
				PlayerPermissions:              pk.PlayerPermissions,
				ServerChunkTickRadius:          pk.ServerChunkTickRadius,
				HasLockedBehaviourPack:         pk.HasLockedBehaviourPack,
				HasLockedTexturePack:           pk.HasLockedTexturePack,
				FromLockedWorldTemplate:        pk.FromLockedWorldTemplate,
				MSAGamerTagsOnly:               pk.MSAGamerTagsOnly,
				FromWorldTemplate:              pk.FromWorldTemplate,
				WorldTemplateSettingsLocked:    pk.WorldTemplateSettingsLocked,
				OnlySpawnV1Villagers:           pk.OnlySpawnV1Villagers,
				BaseGameVersion:                pk.BaseGameVersion,
				LimitedWorldWidth:              pk.LimitedWorldWidth,
				LimitedWorldDepth:              pk.LimitedWorldDepth,
				NewNether:                      pk.NewNether,
				EducationSharedResourceURI:     pk.EducationSharedResourceURI,
				ForceExperimentalGameplay:      enabled,
				LevelID:                        pk.LevelID,
				WorldName:                      pk.WorldName,
				TemplateContentIdentity:        pk.TemplateContentIdentity,
				Trial:                          pk.Trial,
				PlayerMovementSettings: legacyprotocol.PlayerMovementSettings{
					MovementType:                     legacyprotocol.PlayerMovementModeServer,
					RewindHistorySize:                pk.PlayerMovementSettings.RewindHistorySize,
					ServerAuthoritativeBlockBreaking: pk.PlayerMovementSettings.ServerAuthoritativeBlockBreaking,
				},
				Time:            pk.Time,
				EnchantmentSeed: pk.EnchantmentSeed,
				Blocks:          pk.Blocks,
				Items: lo.Map(pk.Items, func(entry protocol.ItemEntry, _ int) legacyprotocol.ItemEntry {
					return legacyprotocol.ItemEntry{
						Name:           entry.Name,
						RuntimeID:      entry.RuntimeID,
						ComponentBased: entry.ComponentBased,
					}
				}),
				MultiPlayerCorrelationID:     pk.MultiPlayerCorrelationID,
				ServerAuthoritativeInventory: pk.ServerAuthoritativeInventory,
				GameVersion:                  pk.GameVersion,
				ServerBlockStateChecksum:     pk.ServerBlockStateChecksum,
			}
		case *packet.StopSound:
			result[i] = &legacypacket.StopSound{
				SoundName: pk.SoundName,
				StopAll:   pk.StopAll,
			}
		case *packet.SubChunk:
			result[i] = &legacypacket.SubChunk{
				CacheEnabled: pk.CacheEnabled,
				Dimension:    pk.Dimension,
				Position:     pk.Position,
				SubChunkEntries: lo.Map(pk.SubChunkEntries, func(entry protocol.SubChunkEntry, _ int) legacyprotocol.SubChunkEntry {
					return legacyprotocol.SubChunkEntry{
						Offset:        entry.Offset,
						Result:        entry.Result,
						RawPayload:    entry.RawPayload,
						HeightMapType: entry.HeightMapType,
						HeightMapData: entry.HeightMapData,
						BlobHash:      entry.BlobHash,
					}
				}),
			}
		case *packet.Text:
			result[i] = &legacypacket.Text{
				TextType:         pk.TextType,
				NeedsTranslation: pk.NeedsTranslation,
				SourceName:       pk.SourceName,
				Message:          pk.Message,
				Parameters:       pk.Parameters,
				XUID:             pk.XUID,
				PlatformChatID:   pk.PlatformChatID,
			}
		case *packet.Transfer:
			result[i] = &legacypacket.Transfer{
				Address: pk.Address,
				Port:    pk.Port,
			}
		case *packet.UpdateAbilities:
			handleFlag := func(layers []protocol.AbilityLayer, secondFlag bool) uint32 {
				layerMapping := map[uint32]uint32{
					protocol.AbilityAttackPlayers: packet.AdventureSettingsFlagsNoPvM,
					protocol.AbilityFlying:        packet.AdventureFlagFlying,
					protocol.AbilityMayFly:        packet.AdventureFlagAllowFlight,
					protocol.AbilityMuted:         packet.AdventureFlagMuted,
					protocol.AbilityNoClip:        packet.AdventureFlagNoClip,
					protocol.AbilityWorldBuilder:  packet.AdventureFlagWorldBuilder,
				}
				if secondFlag {
					layerMapping = map[uint32]uint32{
						protocol.AbilityAttackMobs:       packet.ActionPermissionAttackMobs,
						protocol.AbilityAttackPlayers:    packet.ActionPermissionAttackPlayers,
						protocol.AbilityBuild:            packet.ActionPermissionBuild,
						protocol.AbilityDoorsAndSwitches: packet.ActionPermissionDoorsAndSwitches,
						protocol.AbilityMine:             packet.ActionPermissionMine,
						protocol.AbilityOpenContainers:   packet.ActionPermissionOpenContainers,
						protocol.AbilityOperatorCommands: packet.ActionPermissionOperator,
					}
				}

				out := uint32(0)
				for _, layer := range layers {
					for flag, mapped := range layerMapping {
						if (layer.Values & flag) != 0 {
							out |= mapped
						}
					}
				}
				return out
			}

			result[i] = &packet.AdventureSettings{
				Flags:                  handleFlag(pk.AbilityData.Layers, false),
				CommandPermissionLevel: uint32(pk.AbilityData.CommandPermissions),
				ActionPermissions:      handleFlag(pk.AbilityData.Layers, true),
				PermissionLevel:        uint32(pk.AbilityData.PlayerPermissions),
				PlayerUniqueID:         pk.AbilityData.EntityUniqueID,
			}
		case *packet.UpdateAttributes:
			result[i] = &legacypacket.UpdateAttributes{
				EntityRuntimeID: pk.EntityRuntimeID,
				Attributes: lo.Map(pk.Attributes, func(item protocol.Attribute, _ int) legacyprotocol.Attribute {
					return legacyprotocol.Attribute{Attribute: item}
				}),
				Tick: pk.Tick,
			}
		case *packet.UpdatePlayerGameType:
			result[i] = &legacypacket.UpdatePlayerGameType{
				GameType:       pk.GameType,
				PlayerUniqueID: pk.PlayerUniqueID,
			}
		}
	}
	return result
}
