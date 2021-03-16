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

// MessagesUpdateDialogFilterRequest represents TL type `messages.updateDialogFilter#1ad4a04a`.
// Update folder¹
//
// Links:
//  1) https://core.telegram.org/api/folders
//
// See https://core.telegram.org/method/messages.updateDialogFilter for reference.
type MessagesUpdateDialogFilterRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Folder¹ ID
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	ID int
	// Folder¹ info
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	//
	// Use SetFilter and GetFilter helpers.
	Filter DialogFilter
}

// MessagesUpdateDialogFilterRequestTypeID is TL type id of MessagesUpdateDialogFilterRequest.
const MessagesUpdateDialogFilterRequestTypeID = 0x1ad4a04a

func (u *MessagesUpdateDialogFilterRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.ID == 0) {
		return false
	}
	if !(u.Filter.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *MessagesUpdateDialogFilterRequest) String() string {
	if u == nil {
		return "MessagesUpdateDialogFilterRequest(nil)"
	}
	type Alias MessagesUpdateDialogFilterRequest
	return fmt.Sprintf("MessagesUpdateDialogFilterRequest%+v", Alias(*u))
}

// FillFrom fills MessagesUpdateDialogFilterRequest from given interface.
func (u *MessagesUpdateDialogFilterRequest) FillFrom(from interface {
	GetID() (value int)
	GetFilter() (value DialogFilter, ok bool)
}) {
	u.ID = from.GetID()
	if val, ok := from.GetFilter(); ok {
		u.Filter = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesUpdateDialogFilterRequest) TypeID() uint32 {
	return MessagesUpdateDialogFilterRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesUpdateDialogFilterRequest) TypeName() string {
	return "messages.updateDialogFilter"
}

// TypeInfo returns info about TL type.
func (u *MessagesUpdateDialogFilterRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.updateDialogFilter",
		ID:   MessagesUpdateDialogFilterRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
			Null:       !u.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *MessagesUpdateDialogFilterRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.updateDialogFilter#1ad4a04a as nil")
	}
	b.PutID(MessagesUpdateDialogFilterRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *MessagesUpdateDialogFilterRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.updateDialogFilter#1ad4a04a as nil")
	}
	if !(u.Filter.Zero()) {
		u.Flags.Set(0)
	}
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.updateDialogFilter#1ad4a04a: field flags: %w", err)
	}
	b.PutInt(u.ID)
	if u.Flags.Has(0) {
		if err := u.Filter.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.updateDialogFilter#1ad4a04a: field filter: %w", err)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (u *MessagesUpdateDialogFilterRequest) GetID() (value int) {
	return u.ID
}

// SetFilter sets value of Filter conditional field.
func (u *MessagesUpdateDialogFilterRequest) SetFilter(value DialogFilter) {
	u.Flags.Set(0)
	u.Filter = value
}

// GetFilter returns value of Filter conditional field and
// boolean which is true if field was set.
func (u *MessagesUpdateDialogFilterRequest) GetFilter() (value DialogFilter, ok bool) {
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.Filter, true
}

// Decode implements bin.Decoder.
func (u *MessagesUpdateDialogFilterRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.updateDialogFilter#1ad4a04a to nil")
	}
	if err := b.ConsumeID(MessagesUpdateDialogFilterRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.updateDialogFilter#1ad4a04a: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *MessagesUpdateDialogFilterRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.updateDialogFilter#1ad4a04a to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.updateDialogFilter#1ad4a04a: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.updateDialogFilter#1ad4a04a: field id: %w", err)
		}
		u.ID = value
	}
	if u.Flags.Has(0) {
		if err := u.Filter.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.updateDialogFilter#1ad4a04a: field filter: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesUpdateDialogFilterRequest.
var (
	_ bin.Encoder     = &MessagesUpdateDialogFilterRequest{}
	_ bin.Decoder     = &MessagesUpdateDialogFilterRequest{}
	_ bin.BareEncoder = &MessagesUpdateDialogFilterRequest{}
	_ bin.BareDecoder = &MessagesUpdateDialogFilterRequest{}
)

// MessagesUpdateDialogFilter invokes method messages.updateDialogFilter#1ad4a04a returning error if any.
// Update folder¹
//
// Links:
//  1) https://core.telegram.org/api/folders
//
// Possible errors:
//  400 FILTER_ID_INVALID: The specified filter ID is invalid
//
// See https://core.telegram.org/method/messages.updateDialogFilter for reference.
func (c *Client) MessagesUpdateDialogFilter(ctx context.Context, request *MessagesUpdateDialogFilterRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
