package packet

import (
	"image/color"

	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	MapUpdateFlagTexture = 1 << (iota + 1)
	MapUpdateFlagDecoration
	MapUpdateFlagInitialisation
)

// ClientBoundMapItemData is sent by the server to the client to update the data of a map shown to the client.
// It is sent with a combination of flags that specify what data is updated.
// The ClientBoundMapItemData packet may be used to update specific parts of the map only. It is not required
// to send the entire map each time when updating one part.
type ClientBoundMapItemData struct {
	MapID          int64
	UpdateFlags    uint32
	Dimension      byte
	LockedMap      bool
	Scale          byte
	MapsIncludedIn []int64
	TrackedObjects []protocol.MapTrackedObject
	Decorations    []protocol.MapDecoration
	Height         int32
	Width          int32
	XOffset        int32
	YOffset        int32
	Pixels         [][]color.RGBA
}

// ID ...
func (*ClientBoundMapItemData) ID() uint32 {
	return IDClientBoundMapItemData
}

// Marshal ...
func (pk *ClientBoundMapItemData) Marshal(io protocol.IO) {
	io.Varint64(&pk.MapID)
	io.Varuint32(&pk.UpdateFlags)
	io.Uint8(&pk.Dimension)
	io.Bool(&pk.LockedMap)

	if pk.UpdateFlags&MapUpdateFlagInitialisation != 0 {
		l := uint32(len(pk.MapsIncludedIn))
		io.Varuint32(&l)
		for _, mapID := range pk.MapsIncludedIn {
			io.Varint64(&mapID)
		}
	}
	if pk.UpdateFlags&(MapUpdateFlagInitialisation|MapUpdateFlagDecoration|MapUpdateFlagTexture) != 0 {
		io.Uint8(&pk.Scale)
	}
	if pk.UpdateFlags&MapUpdateFlagDecoration != 0 {
		protocol.Slice(io, &pk.TrackedObjects)
		protocol.Slice(io, &pk.Decorations)
	}
	if pk.UpdateFlags&MapUpdateFlagTexture != 0 {
		// Some basic validation for the values passed into the packet.
		if pk.Width <= 0 || pk.Height <= 0 {
			panic("invalid map texture update: width and height must be at least 1")
		}

		io.Varint32(&pk.Width)
		io.Varint32(&pk.Height)
		io.Varint32(&pk.XOffset)
		io.Varint32(&pk.YOffset)

		l := uint32(pk.Width * pk.Height)
		io.Varuint32(&l)

		if len(pk.Pixels) != int(pk.Height) {
			panic("invalid map texture update: length of outer pixels array must be equal to height")
		}
		for y := int32(0); y < pk.Height; y++ {
			if len(pk.Pixels[y]) != int(pk.Width) {
				panic("invalid map texture update: length of inner pixels array must be equal to width")
			}
			for x := int32(0); x < pk.Width; x++ {
				io.VarRGBA(&pk.Pixels[y][x])
			}
		}
	}
}
