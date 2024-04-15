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

// SetBusinessIntroRequest represents TL type `setBusinessIntro#d2f2e117`.
type SetBusinessIntroRequest struct {
	// The new intro of the business; pass null to remove the intro
	Intro InputBusinessIntro
}

// SetBusinessIntroRequestTypeID is TL type id of SetBusinessIntroRequest.
const SetBusinessIntroRequestTypeID = 0xd2f2e117

// Ensuring interfaces in compile-time for SetBusinessIntroRequest.
var (
	_ bin.Encoder     = &SetBusinessIntroRequest{}
	_ bin.Decoder     = &SetBusinessIntroRequest{}
	_ bin.BareEncoder = &SetBusinessIntroRequest{}
	_ bin.BareDecoder = &SetBusinessIntroRequest{}
)

func (s *SetBusinessIntroRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Intro.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetBusinessIntroRequest) String() string {
	if s == nil {
		return "SetBusinessIntroRequest(nil)"
	}
	type Alias SetBusinessIntroRequest
	return fmt.Sprintf("SetBusinessIntroRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetBusinessIntroRequest) TypeID() uint32 {
	return SetBusinessIntroRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetBusinessIntroRequest) TypeName() string {
	return "setBusinessIntro"
}

// TypeInfo returns info about TL type.
func (s *SetBusinessIntroRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setBusinessIntro",
		ID:   SetBusinessIntroRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Intro",
			SchemaName: "intro",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetBusinessIntroRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBusinessIntro#d2f2e117 as nil")
	}
	b.PutID(SetBusinessIntroRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetBusinessIntroRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBusinessIntro#d2f2e117 as nil")
	}
	if err := s.Intro.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setBusinessIntro#d2f2e117: field intro: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetBusinessIntroRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBusinessIntro#d2f2e117 to nil")
	}
	if err := b.ConsumeID(SetBusinessIntroRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setBusinessIntro#d2f2e117: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetBusinessIntroRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBusinessIntro#d2f2e117 to nil")
	}
	{
		if err := s.Intro.Decode(b); err != nil {
			return fmt.Errorf("unable to decode setBusinessIntro#d2f2e117: field intro: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetBusinessIntroRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setBusinessIntro#d2f2e117 as nil")
	}
	b.ObjStart()
	b.PutID("setBusinessIntro")
	b.Comma()
	b.FieldStart("intro")
	if err := s.Intro.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setBusinessIntro#d2f2e117: field intro: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetBusinessIntroRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setBusinessIntro#d2f2e117 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setBusinessIntro"); err != nil {
				return fmt.Errorf("unable to decode setBusinessIntro#d2f2e117: %w", err)
			}
		case "intro":
			if err := s.Intro.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode setBusinessIntro#d2f2e117: field intro: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetIntro returns value of Intro field.
func (s *SetBusinessIntroRequest) GetIntro() (value InputBusinessIntro) {
	if s == nil {
		return
	}
	return s.Intro
}

// SetBusinessIntro invokes method setBusinessIntro#d2f2e117 returning error if any.
func (c *Client) SetBusinessIntro(ctx context.Context, intro InputBusinessIntro) error {
	var ok Ok

	request := &SetBusinessIntroRequest{
		Intro: intro,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}