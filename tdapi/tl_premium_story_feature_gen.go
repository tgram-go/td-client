// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// PremiumStoryFeaturePriorityOrder represents TL type `premiumStoryFeaturePriorityOrder#8ff172c7`.
type PremiumStoryFeaturePriorityOrder struct {
}

// PremiumStoryFeaturePriorityOrderTypeID is TL type id of PremiumStoryFeaturePriorityOrder.
const PremiumStoryFeaturePriorityOrderTypeID = 0x8ff172c7

// construct implements constructor of PremiumStoryFeatureClass.
func (p PremiumStoryFeaturePriorityOrder) construct() PremiumStoryFeatureClass { return &p }

// Ensuring interfaces in compile-time for PremiumStoryFeaturePriorityOrder.
var (
	_ bin.Encoder     = &PremiumStoryFeaturePriorityOrder{}
	_ bin.Decoder     = &PremiumStoryFeaturePriorityOrder{}
	_ bin.BareEncoder = &PremiumStoryFeaturePriorityOrder{}
	_ bin.BareDecoder = &PremiumStoryFeaturePriorityOrder{}

	_ PremiumStoryFeatureClass = &PremiumStoryFeaturePriorityOrder{}
)

func (p *PremiumStoryFeaturePriorityOrder) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PremiumStoryFeaturePriorityOrder) String() string {
	if p == nil {
		return "PremiumStoryFeaturePriorityOrder(nil)"
	}
	type Alias PremiumStoryFeaturePriorityOrder
	return fmt.Sprintf("PremiumStoryFeaturePriorityOrder%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumStoryFeaturePriorityOrder) TypeID() uint32 {
	return PremiumStoryFeaturePriorityOrderTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumStoryFeaturePriorityOrder) TypeName() string {
	return "premiumStoryFeaturePriorityOrder"
}

// TypeInfo returns info about TL type.
func (p *PremiumStoryFeaturePriorityOrder) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premiumStoryFeaturePriorityOrder",
		ID:   PremiumStoryFeaturePriorityOrderTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PremiumStoryFeaturePriorityOrder) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeaturePriorityOrder#8ff172c7 as nil")
	}
	b.PutID(PremiumStoryFeaturePriorityOrderTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PremiumStoryFeaturePriorityOrder) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeaturePriorityOrder#8ff172c7 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PremiumStoryFeaturePriorityOrder) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeaturePriorityOrder#8ff172c7 to nil")
	}
	if err := b.ConsumeID(PremiumStoryFeaturePriorityOrderTypeID); err != nil {
		return fmt.Errorf("unable to decode premiumStoryFeaturePriorityOrder#8ff172c7: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PremiumStoryFeaturePriorityOrder) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeaturePriorityOrder#8ff172c7 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PremiumStoryFeaturePriorityOrder) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeaturePriorityOrder#8ff172c7 as nil")
	}
	b.ObjStart()
	b.PutID("premiumStoryFeaturePriorityOrder")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PremiumStoryFeaturePriorityOrder) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeaturePriorityOrder#8ff172c7 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("premiumStoryFeaturePriorityOrder"); err != nil {
				return fmt.Errorf("unable to decode premiumStoryFeaturePriorityOrder#8ff172c7: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// PremiumStoryFeatureStealthMode represents TL type `premiumStoryFeatureStealthMode#47343da4`.
type PremiumStoryFeatureStealthMode struct {
}

// PremiumStoryFeatureStealthModeTypeID is TL type id of PremiumStoryFeatureStealthMode.
const PremiumStoryFeatureStealthModeTypeID = 0x47343da4

// construct implements constructor of PremiumStoryFeatureClass.
func (p PremiumStoryFeatureStealthMode) construct() PremiumStoryFeatureClass { return &p }

// Ensuring interfaces in compile-time for PremiumStoryFeatureStealthMode.
var (
	_ bin.Encoder     = &PremiumStoryFeatureStealthMode{}
	_ bin.Decoder     = &PremiumStoryFeatureStealthMode{}
	_ bin.BareEncoder = &PremiumStoryFeatureStealthMode{}
	_ bin.BareDecoder = &PremiumStoryFeatureStealthMode{}

	_ PremiumStoryFeatureClass = &PremiumStoryFeatureStealthMode{}
)

func (p *PremiumStoryFeatureStealthMode) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PremiumStoryFeatureStealthMode) String() string {
	if p == nil {
		return "PremiumStoryFeatureStealthMode(nil)"
	}
	type Alias PremiumStoryFeatureStealthMode
	return fmt.Sprintf("PremiumStoryFeatureStealthMode%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumStoryFeatureStealthMode) TypeID() uint32 {
	return PremiumStoryFeatureStealthModeTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumStoryFeatureStealthMode) TypeName() string {
	return "premiumStoryFeatureStealthMode"
}

// TypeInfo returns info about TL type.
func (p *PremiumStoryFeatureStealthMode) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premiumStoryFeatureStealthMode",
		ID:   PremiumStoryFeatureStealthModeTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PremiumStoryFeatureStealthMode) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeatureStealthMode#47343da4 as nil")
	}
	b.PutID(PremiumStoryFeatureStealthModeTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PremiumStoryFeatureStealthMode) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeatureStealthMode#47343da4 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PremiumStoryFeatureStealthMode) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeatureStealthMode#47343da4 to nil")
	}
	if err := b.ConsumeID(PremiumStoryFeatureStealthModeTypeID); err != nil {
		return fmt.Errorf("unable to decode premiumStoryFeatureStealthMode#47343da4: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PremiumStoryFeatureStealthMode) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeatureStealthMode#47343da4 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PremiumStoryFeatureStealthMode) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeatureStealthMode#47343da4 as nil")
	}
	b.ObjStart()
	b.PutID("premiumStoryFeatureStealthMode")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PremiumStoryFeatureStealthMode) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeatureStealthMode#47343da4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("premiumStoryFeatureStealthMode"); err != nil {
				return fmt.Errorf("unable to decode premiumStoryFeatureStealthMode#47343da4: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// PremiumStoryFeaturePermanentViewsHistory represents TL type `premiumStoryFeaturePermanentViewsHistory#c2a047a0`.
type PremiumStoryFeaturePermanentViewsHistory struct {
}

// PremiumStoryFeaturePermanentViewsHistoryTypeID is TL type id of PremiumStoryFeaturePermanentViewsHistory.
const PremiumStoryFeaturePermanentViewsHistoryTypeID = 0xc2a047a0

// construct implements constructor of PremiumStoryFeatureClass.
func (p PremiumStoryFeaturePermanentViewsHistory) construct() PremiumStoryFeatureClass { return &p }

// Ensuring interfaces in compile-time for PremiumStoryFeaturePermanentViewsHistory.
var (
	_ bin.Encoder     = &PremiumStoryFeaturePermanentViewsHistory{}
	_ bin.Decoder     = &PremiumStoryFeaturePermanentViewsHistory{}
	_ bin.BareEncoder = &PremiumStoryFeaturePermanentViewsHistory{}
	_ bin.BareDecoder = &PremiumStoryFeaturePermanentViewsHistory{}

	_ PremiumStoryFeatureClass = &PremiumStoryFeaturePermanentViewsHistory{}
)

func (p *PremiumStoryFeaturePermanentViewsHistory) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PremiumStoryFeaturePermanentViewsHistory) String() string {
	if p == nil {
		return "PremiumStoryFeaturePermanentViewsHistory(nil)"
	}
	type Alias PremiumStoryFeaturePermanentViewsHistory
	return fmt.Sprintf("PremiumStoryFeaturePermanentViewsHistory%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumStoryFeaturePermanentViewsHistory) TypeID() uint32 {
	return PremiumStoryFeaturePermanentViewsHistoryTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumStoryFeaturePermanentViewsHistory) TypeName() string {
	return "premiumStoryFeaturePermanentViewsHistory"
}

// TypeInfo returns info about TL type.
func (p *PremiumStoryFeaturePermanentViewsHistory) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premiumStoryFeaturePermanentViewsHistory",
		ID:   PremiumStoryFeaturePermanentViewsHistoryTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PremiumStoryFeaturePermanentViewsHistory) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeaturePermanentViewsHistory#c2a047a0 as nil")
	}
	b.PutID(PremiumStoryFeaturePermanentViewsHistoryTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PremiumStoryFeaturePermanentViewsHistory) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeaturePermanentViewsHistory#c2a047a0 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PremiumStoryFeaturePermanentViewsHistory) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeaturePermanentViewsHistory#c2a047a0 to nil")
	}
	if err := b.ConsumeID(PremiumStoryFeaturePermanentViewsHistoryTypeID); err != nil {
		return fmt.Errorf("unable to decode premiumStoryFeaturePermanentViewsHistory#c2a047a0: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PremiumStoryFeaturePermanentViewsHistory) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeaturePermanentViewsHistory#c2a047a0 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PremiumStoryFeaturePermanentViewsHistory) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeaturePermanentViewsHistory#c2a047a0 as nil")
	}
	b.ObjStart()
	b.PutID("premiumStoryFeaturePermanentViewsHistory")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PremiumStoryFeaturePermanentViewsHistory) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeaturePermanentViewsHistory#c2a047a0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("premiumStoryFeaturePermanentViewsHistory"); err != nil {
				return fmt.Errorf("unable to decode premiumStoryFeaturePermanentViewsHistory#c2a047a0: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// PremiumStoryFeatureCustomExpirationDuration represents TL type `premiumStoryFeatureCustomExpirationDuration#dca40a96`.
type PremiumStoryFeatureCustomExpirationDuration struct {
}

// PremiumStoryFeatureCustomExpirationDurationTypeID is TL type id of PremiumStoryFeatureCustomExpirationDuration.
const PremiumStoryFeatureCustomExpirationDurationTypeID = 0xdca40a96

// construct implements constructor of PremiumStoryFeatureClass.
func (p PremiumStoryFeatureCustomExpirationDuration) construct() PremiumStoryFeatureClass { return &p }

// Ensuring interfaces in compile-time for PremiumStoryFeatureCustomExpirationDuration.
var (
	_ bin.Encoder     = &PremiumStoryFeatureCustomExpirationDuration{}
	_ bin.Decoder     = &PremiumStoryFeatureCustomExpirationDuration{}
	_ bin.BareEncoder = &PremiumStoryFeatureCustomExpirationDuration{}
	_ bin.BareDecoder = &PremiumStoryFeatureCustomExpirationDuration{}

	_ PremiumStoryFeatureClass = &PremiumStoryFeatureCustomExpirationDuration{}
)

func (p *PremiumStoryFeatureCustomExpirationDuration) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PremiumStoryFeatureCustomExpirationDuration) String() string {
	if p == nil {
		return "PremiumStoryFeatureCustomExpirationDuration(nil)"
	}
	type Alias PremiumStoryFeatureCustomExpirationDuration
	return fmt.Sprintf("PremiumStoryFeatureCustomExpirationDuration%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumStoryFeatureCustomExpirationDuration) TypeID() uint32 {
	return PremiumStoryFeatureCustomExpirationDurationTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumStoryFeatureCustomExpirationDuration) TypeName() string {
	return "premiumStoryFeatureCustomExpirationDuration"
}

// TypeInfo returns info about TL type.
func (p *PremiumStoryFeatureCustomExpirationDuration) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premiumStoryFeatureCustomExpirationDuration",
		ID:   PremiumStoryFeatureCustomExpirationDurationTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PremiumStoryFeatureCustomExpirationDuration) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeatureCustomExpirationDuration#dca40a96 as nil")
	}
	b.PutID(PremiumStoryFeatureCustomExpirationDurationTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PremiumStoryFeatureCustomExpirationDuration) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeatureCustomExpirationDuration#dca40a96 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PremiumStoryFeatureCustomExpirationDuration) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeatureCustomExpirationDuration#dca40a96 to nil")
	}
	if err := b.ConsumeID(PremiumStoryFeatureCustomExpirationDurationTypeID); err != nil {
		return fmt.Errorf("unable to decode premiumStoryFeatureCustomExpirationDuration#dca40a96: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PremiumStoryFeatureCustomExpirationDuration) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeatureCustomExpirationDuration#dca40a96 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PremiumStoryFeatureCustomExpirationDuration) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeatureCustomExpirationDuration#dca40a96 as nil")
	}
	b.ObjStart()
	b.PutID("premiumStoryFeatureCustomExpirationDuration")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PremiumStoryFeatureCustomExpirationDuration) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeatureCustomExpirationDuration#dca40a96 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("premiumStoryFeatureCustomExpirationDuration"); err != nil {
				return fmt.Errorf("unable to decode premiumStoryFeatureCustomExpirationDuration#dca40a96: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// PremiumStoryFeatureSaveStories represents TL type `premiumStoryFeatureSaveStories#a6842fbd`.
type PremiumStoryFeatureSaveStories struct {
}

// PremiumStoryFeatureSaveStoriesTypeID is TL type id of PremiumStoryFeatureSaveStories.
const PremiumStoryFeatureSaveStoriesTypeID = 0xa6842fbd

// construct implements constructor of PremiumStoryFeatureClass.
func (p PremiumStoryFeatureSaveStories) construct() PremiumStoryFeatureClass { return &p }

// Ensuring interfaces in compile-time for PremiumStoryFeatureSaveStories.
var (
	_ bin.Encoder     = &PremiumStoryFeatureSaveStories{}
	_ bin.Decoder     = &PremiumStoryFeatureSaveStories{}
	_ bin.BareEncoder = &PremiumStoryFeatureSaveStories{}
	_ bin.BareDecoder = &PremiumStoryFeatureSaveStories{}

	_ PremiumStoryFeatureClass = &PremiumStoryFeatureSaveStories{}
)

func (p *PremiumStoryFeatureSaveStories) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PremiumStoryFeatureSaveStories) String() string {
	if p == nil {
		return "PremiumStoryFeatureSaveStories(nil)"
	}
	type Alias PremiumStoryFeatureSaveStories
	return fmt.Sprintf("PremiumStoryFeatureSaveStories%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumStoryFeatureSaveStories) TypeID() uint32 {
	return PremiumStoryFeatureSaveStoriesTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumStoryFeatureSaveStories) TypeName() string {
	return "premiumStoryFeatureSaveStories"
}

// TypeInfo returns info about TL type.
func (p *PremiumStoryFeatureSaveStories) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premiumStoryFeatureSaveStories",
		ID:   PremiumStoryFeatureSaveStoriesTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PremiumStoryFeatureSaveStories) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeatureSaveStories#a6842fbd as nil")
	}
	b.PutID(PremiumStoryFeatureSaveStoriesTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PremiumStoryFeatureSaveStories) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeatureSaveStories#a6842fbd as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PremiumStoryFeatureSaveStories) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeatureSaveStories#a6842fbd to nil")
	}
	if err := b.ConsumeID(PremiumStoryFeatureSaveStoriesTypeID); err != nil {
		return fmt.Errorf("unable to decode premiumStoryFeatureSaveStories#a6842fbd: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PremiumStoryFeatureSaveStories) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeatureSaveStories#a6842fbd to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PremiumStoryFeatureSaveStories) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeatureSaveStories#a6842fbd as nil")
	}
	b.ObjStart()
	b.PutID("premiumStoryFeatureSaveStories")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PremiumStoryFeatureSaveStories) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeatureSaveStories#a6842fbd to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("premiumStoryFeatureSaveStories"); err != nil {
				return fmt.Errorf("unable to decode premiumStoryFeatureSaveStories#a6842fbd: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// PremiumStoryFeatureLinksAndFormatting represents TL type `premiumStoryFeatureLinksAndFormatting#dae383f7`.
type PremiumStoryFeatureLinksAndFormatting struct {
}

// PremiumStoryFeatureLinksAndFormattingTypeID is TL type id of PremiumStoryFeatureLinksAndFormatting.
const PremiumStoryFeatureLinksAndFormattingTypeID = 0xdae383f7

// construct implements constructor of PremiumStoryFeatureClass.
func (p PremiumStoryFeatureLinksAndFormatting) construct() PremiumStoryFeatureClass { return &p }

// Ensuring interfaces in compile-time for PremiumStoryFeatureLinksAndFormatting.
var (
	_ bin.Encoder     = &PremiumStoryFeatureLinksAndFormatting{}
	_ bin.Decoder     = &PremiumStoryFeatureLinksAndFormatting{}
	_ bin.BareEncoder = &PremiumStoryFeatureLinksAndFormatting{}
	_ bin.BareDecoder = &PremiumStoryFeatureLinksAndFormatting{}

	_ PremiumStoryFeatureClass = &PremiumStoryFeatureLinksAndFormatting{}
)

func (p *PremiumStoryFeatureLinksAndFormatting) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PremiumStoryFeatureLinksAndFormatting) String() string {
	if p == nil {
		return "PremiumStoryFeatureLinksAndFormatting(nil)"
	}
	type Alias PremiumStoryFeatureLinksAndFormatting
	return fmt.Sprintf("PremiumStoryFeatureLinksAndFormatting%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumStoryFeatureLinksAndFormatting) TypeID() uint32 {
	return PremiumStoryFeatureLinksAndFormattingTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumStoryFeatureLinksAndFormatting) TypeName() string {
	return "premiumStoryFeatureLinksAndFormatting"
}

// TypeInfo returns info about TL type.
func (p *PremiumStoryFeatureLinksAndFormatting) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premiumStoryFeatureLinksAndFormatting",
		ID:   PremiumStoryFeatureLinksAndFormattingTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PremiumStoryFeatureLinksAndFormatting) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeatureLinksAndFormatting#dae383f7 as nil")
	}
	b.PutID(PremiumStoryFeatureLinksAndFormattingTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PremiumStoryFeatureLinksAndFormatting) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeatureLinksAndFormatting#dae383f7 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PremiumStoryFeatureLinksAndFormatting) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeatureLinksAndFormatting#dae383f7 to nil")
	}
	if err := b.ConsumeID(PremiumStoryFeatureLinksAndFormattingTypeID); err != nil {
		return fmt.Errorf("unable to decode premiumStoryFeatureLinksAndFormatting#dae383f7: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PremiumStoryFeatureLinksAndFormatting) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeatureLinksAndFormatting#dae383f7 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PremiumStoryFeatureLinksAndFormatting) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumStoryFeatureLinksAndFormatting#dae383f7 as nil")
	}
	b.ObjStart()
	b.PutID("premiumStoryFeatureLinksAndFormatting")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PremiumStoryFeatureLinksAndFormatting) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumStoryFeatureLinksAndFormatting#dae383f7 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("premiumStoryFeatureLinksAndFormatting"); err != nil {
				return fmt.Errorf("unable to decode premiumStoryFeatureLinksAndFormatting#dae383f7: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// PremiumStoryFeatureClassName is schema name of PremiumStoryFeatureClass.
const PremiumStoryFeatureClassName = "PremiumStoryFeature"

// PremiumStoryFeatureClass represents PremiumStoryFeature generic type.
//
// Example:
//
//	g, err := tdapi.DecodePremiumStoryFeature(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.PremiumStoryFeaturePriorityOrder: // premiumStoryFeaturePriorityOrder#8ff172c7
//	case *tdapi.PremiumStoryFeatureStealthMode: // premiumStoryFeatureStealthMode#47343da4
//	case *tdapi.PremiumStoryFeaturePermanentViewsHistory: // premiumStoryFeaturePermanentViewsHistory#c2a047a0
//	case *tdapi.PremiumStoryFeatureCustomExpirationDuration: // premiumStoryFeatureCustomExpirationDuration#dca40a96
//	case *tdapi.PremiumStoryFeatureSaveStories: // premiumStoryFeatureSaveStories#a6842fbd
//	case *tdapi.PremiumStoryFeatureLinksAndFormatting: // premiumStoryFeatureLinksAndFormatting#dae383f7
//	default: panic(v)
//	}
type PremiumStoryFeatureClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PremiumStoryFeatureClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodePremiumStoryFeature implements binary de-serialization for PremiumStoryFeatureClass.
func DecodePremiumStoryFeature(buf *bin.Buffer) (PremiumStoryFeatureClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PremiumStoryFeaturePriorityOrderTypeID:
		// Decoding premiumStoryFeaturePriorityOrder#8ff172c7.
		v := PremiumStoryFeaturePriorityOrder{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumStoryFeatureClass: %w", err)
		}
		return &v, nil
	case PremiumStoryFeatureStealthModeTypeID:
		// Decoding premiumStoryFeatureStealthMode#47343da4.
		v := PremiumStoryFeatureStealthMode{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumStoryFeatureClass: %w", err)
		}
		return &v, nil
	case PremiumStoryFeaturePermanentViewsHistoryTypeID:
		// Decoding premiumStoryFeaturePermanentViewsHistory#c2a047a0.
		v := PremiumStoryFeaturePermanentViewsHistory{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumStoryFeatureClass: %w", err)
		}
		return &v, nil
	case PremiumStoryFeatureCustomExpirationDurationTypeID:
		// Decoding premiumStoryFeatureCustomExpirationDuration#dca40a96.
		v := PremiumStoryFeatureCustomExpirationDuration{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumStoryFeatureClass: %w", err)
		}
		return &v, nil
	case PremiumStoryFeatureSaveStoriesTypeID:
		// Decoding premiumStoryFeatureSaveStories#a6842fbd.
		v := PremiumStoryFeatureSaveStories{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumStoryFeatureClass: %w", err)
		}
		return &v, nil
	case PremiumStoryFeatureLinksAndFormattingTypeID:
		// Decoding premiumStoryFeatureLinksAndFormatting#dae383f7.
		v := PremiumStoryFeatureLinksAndFormatting{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumStoryFeatureClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PremiumStoryFeatureClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONPremiumStoryFeature implements binary de-serialization for PremiumStoryFeatureClass.
func DecodeTDLibJSONPremiumStoryFeature(buf tdjson.Decoder) (PremiumStoryFeatureClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "premiumStoryFeaturePriorityOrder":
		// Decoding premiumStoryFeaturePriorityOrder#8ff172c7.
		v := PremiumStoryFeaturePriorityOrder{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumStoryFeatureClass: %w", err)
		}
		return &v, nil
	case "premiumStoryFeatureStealthMode":
		// Decoding premiumStoryFeatureStealthMode#47343da4.
		v := PremiumStoryFeatureStealthMode{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumStoryFeatureClass: %w", err)
		}
		return &v, nil
	case "premiumStoryFeaturePermanentViewsHistory":
		// Decoding premiumStoryFeaturePermanentViewsHistory#c2a047a0.
		v := PremiumStoryFeaturePermanentViewsHistory{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumStoryFeatureClass: %w", err)
		}
		return &v, nil
	case "premiumStoryFeatureCustomExpirationDuration":
		// Decoding premiumStoryFeatureCustomExpirationDuration#dca40a96.
		v := PremiumStoryFeatureCustomExpirationDuration{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumStoryFeatureClass: %w", err)
		}
		return &v, nil
	case "premiumStoryFeatureSaveStories":
		// Decoding premiumStoryFeatureSaveStories#a6842fbd.
		v := PremiumStoryFeatureSaveStories{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumStoryFeatureClass: %w", err)
		}
		return &v, nil
	case "premiumStoryFeatureLinksAndFormatting":
		// Decoding premiumStoryFeatureLinksAndFormatting#dae383f7.
		v := PremiumStoryFeatureLinksAndFormatting{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumStoryFeatureClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PremiumStoryFeatureClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// PremiumStoryFeature boxes the PremiumStoryFeatureClass providing a helper.
type PremiumStoryFeatureBox struct {
	PremiumStoryFeature PremiumStoryFeatureClass
}

// Decode implements bin.Decoder for PremiumStoryFeatureBox.
func (b *PremiumStoryFeatureBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PremiumStoryFeatureBox to nil")
	}
	v, err := DecodePremiumStoryFeature(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PremiumStoryFeature = v
	return nil
}

// Encode implements bin.Encode for PremiumStoryFeatureBox.
func (b *PremiumStoryFeatureBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PremiumStoryFeature == nil {
		return fmt.Errorf("unable to encode PremiumStoryFeatureClass as nil")
	}
	return b.PremiumStoryFeature.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for PremiumStoryFeatureBox.
func (b *PremiumStoryFeatureBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode PremiumStoryFeatureBox to nil")
	}
	v, err := DecodeTDLibJSONPremiumStoryFeature(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PremiumStoryFeature = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for PremiumStoryFeatureBox.
func (b *PremiumStoryFeatureBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.PremiumStoryFeature == nil {
		return fmt.Errorf("unable to encode PremiumStoryFeatureClass as nil")
	}
	return b.PremiumStoryFeature.EncodeTDLibJSON(buf)
}