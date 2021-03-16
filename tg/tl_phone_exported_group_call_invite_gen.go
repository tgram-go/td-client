// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
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
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// PhoneExportedGroupCallInvite represents TL type `phone.exportedGroupCallInvite#204bd158`.
//
// See https://core.telegram.org/constructor/phone.exportedGroupCallInvite for reference.
type PhoneExportedGroupCallInvite struct {
	// Link field of PhoneExportedGroupCallInvite.
	Link string
}

// PhoneExportedGroupCallInviteTypeID is TL type id of PhoneExportedGroupCallInvite.
const PhoneExportedGroupCallInviteTypeID = 0x204bd158

func (e *PhoneExportedGroupCallInvite) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Link == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *PhoneExportedGroupCallInvite) String() string {
	if e == nil {
		return "PhoneExportedGroupCallInvite(nil)"
	}
	type Alias PhoneExportedGroupCallInvite
	return fmt.Sprintf("PhoneExportedGroupCallInvite%+v", Alias(*e))
}

// FillFrom fills PhoneExportedGroupCallInvite from given interface.
func (e *PhoneExportedGroupCallInvite) FillFrom(from interface {
	GetLink() (value string)
}) {
	e.Link = from.GetLink()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneExportedGroupCallInvite) TypeID() uint32 {
	return PhoneExportedGroupCallInviteTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneExportedGroupCallInvite) TypeName() string {
	return "phone.exportedGroupCallInvite"
}

// TypeInfo returns info about TL type.
func (e *PhoneExportedGroupCallInvite) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.exportedGroupCallInvite",
		ID:   PhoneExportedGroupCallInviteTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Link",
			SchemaName: "link",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *PhoneExportedGroupCallInvite) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode phone.exportedGroupCallInvite#204bd158 as nil")
	}
	b.PutID(PhoneExportedGroupCallInviteTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *PhoneExportedGroupCallInvite) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode phone.exportedGroupCallInvite#204bd158 as nil")
	}
	b.PutString(e.Link)
	return nil
}

// GetLink returns value of Link field.
func (e *PhoneExportedGroupCallInvite) GetLink() (value string) {
	return e.Link
}

// Decode implements bin.Decoder.
func (e *PhoneExportedGroupCallInvite) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode phone.exportedGroupCallInvite#204bd158 to nil")
	}
	if err := b.ConsumeID(PhoneExportedGroupCallInviteTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.exportedGroupCallInvite#204bd158: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *PhoneExportedGroupCallInvite) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode phone.exportedGroupCallInvite#204bd158 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phone.exportedGroupCallInvite#204bd158: field link: %w", err)
		}
		e.Link = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneExportedGroupCallInvite.
var (
	_ bin.Encoder     = &PhoneExportedGroupCallInvite{}
	_ bin.Decoder     = &PhoneExportedGroupCallInvite{}
	_ bin.BareEncoder = &PhoneExportedGroupCallInvite{}
	_ bin.BareDecoder = &PhoneExportedGroupCallInvite{}
)
