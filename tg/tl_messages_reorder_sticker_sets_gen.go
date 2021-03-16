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

// MessagesReorderStickerSetsRequest represents TL type `messages.reorderStickerSets#78337739`.
// Reorder installed stickersets
//
// See https://core.telegram.org/method/messages.reorderStickerSets for reference.
type MessagesReorderStickerSetsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Reorder mask stickersets
	Masks bool
	// New stickerset order by stickerset IDs
	Order []int64
}

// MessagesReorderStickerSetsRequestTypeID is TL type id of MessagesReorderStickerSetsRequest.
const MessagesReorderStickerSetsRequestTypeID = 0x78337739

func (r *MessagesReorderStickerSetsRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Masks == false) {
		return false
	}
	if !(r.Order == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReorderStickerSetsRequest) String() string {
	if r == nil {
		return "MessagesReorderStickerSetsRequest(nil)"
	}
	type Alias MessagesReorderStickerSetsRequest
	return fmt.Sprintf("MessagesReorderStickerSetsRequest%+v", Alias(*r))
}

// FillFrom fills MessagesReorderStickerSetsRequest from given interface.
func (r *MessagesReorderStickerSetsRequest) FillFrom(from interface {
	GetMasks() (value bool)
	GetOrder() (value []int64)
}) {
	r.Masks = from.GetMasks()
	r.Order = from.GetOrder()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesReorderStickerSetsRequest) TypeID() uint32 {
	return MessagesReorderStickerSetsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesReorderStickerSetsRequest) TypeName() string {
	return "messages.reorderStickerSets"
}

// TypeInfo returns info about TL type.
func (r *MessagesReorderStickerSetsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.reorderStickerSets",
		ID:   MessagesReorderStickerSetsRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Masks",
			SchemaName: "masks",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "Order",
			SchemaName: "order",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *MessagesReorderStickerSetsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reorderStickerSets#78337739 as nil")
	}
	b.PutID(MessagesReorderStickerSetsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesReorderStickerSetsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reorderStickerSets#78337739 as nil")
	}
	if !(r.Masks == false) {
		r.Flags.Set(0)
	}
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.reorderStickerSets#78337739: field flags: %w", err)
	}
	b.PutVectorHeader(len(r.Order))
	for _, v := range r.Order {
		b.PutLong(v)
	}
	return nil
}

// SetMasks sets value of Masks conditional field.
func (r *MessagesReorderStickerSetsRequest) SetMasks(value bool) {
	if value {
		r.Flags.Set(0)
		r.Masks = true
	} else {
		r.Flags.Unset(0)
		r.Masks = false
	}
}

// GetMasks returns value of Masks conditional field.
func (r *MessagesReorderStickerSetsRequest) GetMasks() (value bool) {
	return r.Flags.Has(0)
}

// GetOrder returns value of Order field.
func (r *MessagesReorderStickerSetsRequest) GetOrder() (value []int64) {
	return r.Order
}

// Decode implements bin.Decoder.
func (r *MessagesReorderStickerSetsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reorderStickerSets#78337739 to nil")
	}
	if err := b.ConsumeID(MessagesReorderStickerSetsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.reorderStickerSets#78337739: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesReorderStickerSetsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reorderStickerSets#78337739 to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.reorderStickerSets#78337739: field flags: %w", err)
		}
	}
	r.Masks = r.Flags.Has(0)
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.reorderStickerSets#78337739: field order: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode messages.reorderStickerSets#78337739: field order: %w", err)
			}
			r.Order = append(r.Order, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesReorderStickerSetsRequest.
var (
	_ bin.Encoder     = &MessagesReorderStickerSetsRequest{}
	_ bin.Decoder     = &MessagesReorderStickerSetsRequest{}
	_ bin.BareEncoder = &MessagesReorderStickerSetsRequest{}
	_ bin.BareDecoder = &MessagesReorderStickerSetsRequest{}
)

// MessagesReorderStickerSets invokes method messages.reorderStickerSets#78337739 returning error if any.
// Reorder installed stickersets
//
// See https://core.telegram.org/method/messages.reorderStickerSets for reference.
func (c *Client) MessagesReorderStickerSets(ctx context.Context, request *MessagesReorderStickerSetsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
