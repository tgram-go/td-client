// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// VideoSize represents TL type `videoSize#e831c556`.
// Animated profile picture¹ in MPEG4 format
//
// Links:
//  1) https://core.telegram.org/api/files#animated-profile-pictures
//
// See https://core.telegram.org/constructor/videoSize for reference.
type VideoSize struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// u for animated profile pictures, and v for trimmed and downscaled video previews
	Type string
	// File location
	Location FileLocationToBeDeprecated
	// Video width
	W int
	// Video height
	H int
	// File size
	Size int
	// Timestamp that should be shown as static preview to the user (seconds)
	//
	// Use SetVideoStartTs and GetVideoStartTs helpers.
	VideoStartTs float64
}

// VideoSizeTypeID is TL type id of VideoSize.
const VideoSizeTypeID = 0xe831c556

// Encode implements bin.Encoder.
func (v *VideoSize) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode videoSize#e831c556 as nil")
	}
	b.PutID(VideoSizeTypeID)
	if err := v.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode videoSize#e831c556: field flags: %w", err)
	}
	b.PutString(v.Type)
	if err := v.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode videoSize#e831c556: field location: %w", err)
	}
	b.PutInt(v.W)
	b.PutInt(v.H)
	b.PutInt(v.Size)
	if v.Flags.Has(0) {
		b.PutDouble(v.VideoStartTs)
	}
	return nil
}

// SetVideoStartTs sets value of VideoStartTs conditional field.
func (v *VideoSize) SetVideoStartTs(value float64) {
	v.Flags.Set(0)
	v.VideoStartTs = value
}

// GetVideoStartTs returns value of VideoStartTs conditional field and
// boolean which is true if field was set.
func (v *VideoSize) GetVideoStartTs() (value float64, ok bool) {
	if !v.Flags.Has(0) {
		return value, false
	}
	return v.VideoStartTs, true
}

// Decode implements bin.Decoder.
func (v *VideoSize) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode videoSize#e831c556 to nil")
	}
	if err := b.ConsumeID(VideoSizeTypeID); err != nil {
		return fmt.Errorf("unable to decode videoSize#e831c556: %w", err)
	}
	{
		if err := v.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode videoSize#e831c556: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#e831c556: field type: %w", err)
		}
		v.Type = value
	}
	{
		if err := v.Location.Decode(b); err != nil {
			return fmt.Errorf("unable to decode videoSize#e831c556: field location: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#e831c556: field w: %w", err)
		}
		v.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#e831c556: field h: %w", err)
		}
		v.H = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#e831c556: field size: %w", err)
		}
		v.Size = value
	}
	if v.Flags.Has(0) {
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#e831c556: field video_start_ts: %w", err)
		}
		v.VideoStartTs = value
	}
	return nil
}

// Ensuring interfaces in compile-time for VideoSize.
var (
	_ bin.Encoder = &VideoSize{}
	_ bin.Decoder = &VideoSize{}
)
