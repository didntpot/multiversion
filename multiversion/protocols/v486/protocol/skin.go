package protocol

import (
	"fmt"

	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Skin represents the skin of a player as sent over network. The skin holds a texture and a model, and
// optional animations which may be present when the skin is created using persona or bought from the
// marketplace.
type Skin struct {
	protocol.Skin
}

// Marshal ...
func (x *Skin) Marshal(io protocol.IO) {
	io.String(&x.SkinID)
	io.String(&x.PlayFabID)
	io.ByteSlice(&x.SkinResourcePatch)
	io.Uint32(&x.SkinImageWidth)
	io.Uint32(&x.SkinImageHeight)
	io.ByteSlice(&x.SkinData)
	protocol.SliceUint32Length(io, &x.Animations)
	io.Uint32(&x.CapeImageWidth)
	io.Uint32(&x.CapeImageHeight)
	io.ByteSlice(&x.CapeData)
	io.ByteSlice(&x.SkinGeometry)
	io.ByteSlice(&x.GeometryDataEngineVersion)
	io.ByteSlice(&x.AnimationData)
	io.String(&x.CapeID)
	io.String(&x.FullID)
	io.String(&x.ArmSize)
	io.String(&x.SkinColour)
	protocol.SliceUint32Length(io, &x.PersonaPieces)
	protocol.SliceUint32Length(io, &x.PieceTintColours)
	if err := x.validate(); err != nil {
		io.InvalidValue(fmt.Sprintf("Skin %v", x.SkinID), "serialised skin", err.Error())
	}
	io.Bool(&x.PremiumSkin)
	io.Bool(&x.PersonaSkin)
	io.Bool(&x.PersonaCapeOnClassicSkin)
	io.Bool(&x.PrimaryUser)
}

// validate ...
func (x *Skin) validate() error {
	if x.SkinImageHeight*x.SkinImageWidth*4 != uint32(len(x.SkinData)) {
		return fmt.Errorf("expected size of skin is %vx%v (%v bytes total), but got %v bytes", x.SkinImageWidth, x.SkinImageHeight, x.SkinImageHeight*x.SkinImageWidth*4, len(x.SkinData))
	}
	if x.CapeImageHeight*x.CapeImageWidth*4 != uint32(len(x.CapeData)) {
		return fmt.Errorf("expected size of cape is %vx%v (%v bytes total), but got %v bytes", x.CapeImageWidth, x.CapeImageHeight, x.CapeImageHeight*x.CapeImageWidth*4, len(x.CapeData))
	}
	for i, animation := range x.Animations {
		if animation.ImageHeight*animation.ImageWidth*4 != uint32(len(animation.ImageData)) {
			return fmt.Errorf("expected size of animation %v is %vx%v (%v bytes total), but got %v bytes", i, animation.ImageWidth, animation.ImageHeight, animation.ImageHeight*animation.ImageWidth*4, len(animation.ImageData))
		}
	}
	return nil
}
