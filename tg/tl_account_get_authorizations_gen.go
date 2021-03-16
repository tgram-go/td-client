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

// AccountGetAuthorizationsRequest represents TL type `account.getAuthorizations#e320c158`.
// Get logged-in sessions
//
// See https://core.telegram.org/method/account.getAuthorizations for reference.
type AccountGetAuthorizationsRequest struct {
}

// AccountGetAuthorizationsRequestTypeID is TL type id of AccountGetAuthorizationsRequest.
const AccountGetAuthorizationsRequestTypeID = 0xe320c158

func (g *AccountGetAuthorizationsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetAuthorizationsRequest) String() string {
	if g == nil {
		return "AccountGetAuthorizationsRequest(nil)"
	}
	type Alias AccountGetAuthorizationsRequest
	return fmt.Sprintf("AccountGetAuthorizationsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetAuthorizationsRequest) TypeID() uint32 {
	return AccountGetAuthorizationsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetAuthorizationsRequest) TypeName() string {
	return "account.getAuthorizations"
}

// TypeInfo returns info about TL type.
func (g *AccountGetAuthorizationsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getAuthorizations",
		ID:   AccountGetAuthorizationsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetAuthorizationsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getAuthorizations#e320c158 as nil")
	}
	b.PutID(AccountGetAuthorizationsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetAuthorizationsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getAuthorizations#e320c158 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetAuthorizationsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getAuthorizations#e320c158 to nil")
	}
	if err := b.ConsumeID(AccountGetAuthorizationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getAuthorizations#e320c158: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetAuthorizationsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getAuthorizations#e320c158 to nil")
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetAuthorizationsRequest.
var (
	_ bin.Encoder     = &AccountGetAuthorizationsRequest{}
	_ bin.Decoder     = &AccountGetAuthorizationsRequest{}
	_ bin.BareEncoder = &AccountGetAuthorizationsRequest{}
	_ bin.BareDecoder = &AccountGetAuthorizationsRequest{}
)

// AccountGetAuthorizations invokes method account.getAuthorizations#e320c158 returning error if any.
// Get logged-in sessions
//
// See https://core.telegram.org/method/account.getAuthorizations for reference.
func (c *Client) AccountGetAuthorizations(ctx context.Context) (*AccountAuthorizations, error) {
	var result AccountAuthorizations

	request := &AccountGetAuthorizationsRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
