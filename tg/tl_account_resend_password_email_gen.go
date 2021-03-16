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

// AccountResendPasswordEmailRequest represents TL type `account.resendPasswordEmail#7a7f2a15`.
// Resend the code to verify an email to use as 2FA recovery method¹.
//
// Links:
//  1) https://core.telegram.org/api/srp
//
// See https://core.telegram.org/method/account.resendPasswordEmail for reference.
type AccountResendPasswordEmailRequest struct {
}

// AccountResendPasswordEmailRequestTypeID is TL type id of AccountResendPasswordEmailRequest.
const AccountResendPasswordEmailRequestTypeID = 0x7a7f2a15

func (r *AccountResendPasswordEmailRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *AccountResendPasswordEmailRequest) String() string {
	if r == nil {
		return "AccountResendPasswordEmailRequest(nil)"
	}
	type Alias AccountResendPasswordEmailRequest
	return fmt.Sprintf("AccountResendPasswordEmailRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountResendPasswordEmailRequest) TypeID() uint32 {
	return AccountResendPasswordEmailRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountResendPasswordEmailRequest) TypeName() string {
	return "account.resendPasswordEmail"
}

// TypeInfo returns info about TL type.
func (r *AccountResendPasswordEmailRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.resendPasswordEmail",
		ID:   AccountResendPasswordEmailRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *AccountResendPasswordEmailRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resendPasswordEmail#7a7f2a15 as nil")
	}
	b.PutID(AccountResendPasswordEmailRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AccountResendPasswordEmailRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resendPasswordEmail#7a7f2a15 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountResendPasswordEmailRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resendPasswordEmail#7a7f2a15 to nil")
	}
	if err := b.ConsumeID(AccountResendPasswordEmailRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.resendPasswordEmail#7a7f2a15: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AccountResendPasswordEmailRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resendPasswordEmail#7a7f2a15 to nil")
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountResendPasswordEmailRequest.
var (
	_ bin.Encoder     = &AccountResendPasswordEmailRequest{}
	_ bin.Decoder     = &AccountResendPasswordEmailRequest{}
	_ bin.BareEncoder = &AccountResendPasswordEmailRequest{}
	_ bin.BareDecoder = &AccountResendPasswordEmailRequest{}
)

// AccountResendPasswordEmail invokes method account.resendPasswordEmail#7a7f2a15 returning error if any.
// Resend the code to verify an email to use as 2FA recovery method¹.
//
// Links:
//  1) https://core.telegram.org/api/srp
//
// See https://core.telegram.org/method/account.resendPasswordEmail for reference.
func (c *Client) AccountResendPasswordEmail(ctx context.Context) (bool, error) {
	var result BoolBox

	request := &AccountResendPasswordEmailRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
