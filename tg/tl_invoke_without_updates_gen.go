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

// InvokeWithoutUpdatesRequest represents TL type `invokeWithoutUpdates#bf9459b7`.
// Invoke a request without subscribing the used connection for updates¹ (this is enabled by default for file queries²).
//
// Links:
//  1) https://core.telegram.org/api/updates
//  2) https://core.telegram.org/api/files
//
// See https://core.telegram.org/constructor/invokeWithoutUpdates for reference.
type InvokeWithoutUpdatesRequest struct {
	// The query
	Query bin.Object
}

// InvokeWithoutUpdatesRequestTypeID is TL type id of InvokeWithoutUpdatesRequest.
const InvokeWithoutUpdatesRequestTypeID = 0xbf9459b7

func (i *InvokeWithoutUpdatesRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Query == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InvokeWithoutUpdatesRequest) String() string {
	if i == nil {
		return "InvokeWithoutUpdatesRequest(nil)"
	}
	type Alias InvokeWithoutUpdatesRequest
	return fmt.Sprintf("InvokeWithoutUpdatesRequest%+v", Alias(*i))
}

// FillFrom fills InvokeWithoutUpdatesRequest from given interface.
func (i *InvokeWithoutUpdatesRequest) FillFrom(from interface {
	GetQuery() (value bin.Object)
}) {
	i.Query = from.GetQuery()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InvokeWithoutUpdatesRequest) TypeID() uint32 {
	return InvokeWithoutUpdatesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*InvokeWithoutUpdatesRequest) TypeName() string {
	return "invokeWithoutUpdates"
}

// TypeInfo returns info about TL type.
func (i *InvokeWithoutUpdatesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "invokeWithoutUpdates",
		ID:   InvokeWithoutUpdatesRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Query",
			SchemaName: "query",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InvokeWithoutUpdatesRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invokeWithoutUpdates#bf9459b7 as nil")
	}
	b.PutID(InvokeWithoutUpdatesRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InvokeWithoutUpdatesRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invokeWithoutUpdates#bf9459b7 as nil")
	}
	if err := i.Query.Encode(b); err != nil {
		return fmt.Errorf("unable to encode invokeWithoutUpdates#bf9459b7: field query: %w", err)
	}
	return nil
}

// GetQuery returns value of Query field.
func (i *InvokeWithoutUpdatesRequest) GetQuery() (value bin.Object) {
	return i.Query
}

// Decode implements bin.Decoder.
func (i *InvokeWithoutUpdatesRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invokeWithoutUpdates#bf9459b7 to nil")
	}
	if err := b.ConsumeID(InvokeWithoutUpdatesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode invokeWithoutUpdates#bf9459b7: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InvokeWithoutUpdatesRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invokeWithoutUpdates#bf9459b7 to nil")
	}
	{
		if err := i.Query.Decode(b); err != nil {
			return fmt.Errorf("unable to decode invokeWithoutUpdates#bf9459b7: field query: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for InvokeWithoutUpdatesRequest.
var (
	_ bin.Encoder     = &InvokeWithoutUpdatesRequest{}
	_ bin.Decoder     = &InvokeWithoutUpdatesRequest{}
	_ bin.BareEncoder = &InvokeWithoutUpdatesRequest{}
	_ bin.BareDecoder = &InvokeWithoutUpdatesRequest{}
)
