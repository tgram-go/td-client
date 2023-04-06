// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// AccountGetAutoSaveSettingsRequest represents TL type `account.getAutoSaveSettings#adcbbcda`.
// Get autosave settings
//
// See https://core.telegram.org/method/account.getAutoSaveSettings for reference.
type AccountGetAutoSaveSettingsRequest struct {
}

// AccountGetAutoSaveSettingsRequestTypeID is TL type id of AccountGetAutoSaveSettingsRequest.
const AccountGetAutoSaveSettingsRequestTypeID = 0xadcbbcda

// Ensuring interfaces in compile-time for AccountGetAutoSaveSettingsRequest.
var (
	_ bin.Encoder     = &AccountGetAutoSaveSettingsRequest{}
	_ bin.Decoder     = &AccountGetAutoSaveSettingsRequest{}
	_ bin.BareEncoder = &AccountGetAutoSaveSettingsRequest{}
	_ bin.BareDecoder = &AccountGetAutoSaveSettingsRequest{}
)

func (g *AccountGetAutoSaveSettingsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetAutoSaveSettingsRequest) String() string {
	if g == nil {
		return "AccountGetAutoSaveSettingsRequest(nil)"
	}
	type Alias AccountGetAutoSaveSettingsRequest
	return fmt.Sprintf("AccountGetAutoSaveSettingsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetAutoSaveSettingsRequest) TypeID() uint32 {
	return AccountGetAutoSaveSettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetAutoSaveSettingsRequest) TypeName() string {
	return "account.getAutoSaveSettings"
}

// TypeInfo returns info about TL type.
func (g *AccountGetAutoSaveSettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getAutoSaveSettings",
		ID:   AccountGetAutoSaveSettingsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetAutoSaveSettingsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getAutoSaveSettings#adcbbcda as nil")
	}
	b.PutID(AccountGetAutoSaveSettingsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetAutoSaveSettingsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getAutoSaveSettings#adcbbcda as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetAutoSaveSettingsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getAutoSaveSettings#adcbbcda to nil")
	}
	if err := b.ConsumeID(AccountGetAutoSaveSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getAutoSaveSettings#adcbbcda: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetAutoSaveSettingsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getAutoSaveSettings#adcbbcda to nil")
	}
	return nil
}

// AccountGetAutoSaveSettings invokes method account.getAutoSaveSettings#adcbbcda returning error if any.
// Get autosave settings
//
// See https://core.telegram.org/method/account.getAutoSaveSettings for reference.
func (c *Client) AccountGetAutoSaveSettings(ctx context.Context) (*AccountAutoSaveSettings, error) {
	var result AccountAutoSaveSettings

	request := &AccountGetAutoSaveSettingsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}