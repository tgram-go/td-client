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

// AccountSaveSecureValueRequest represents TL type `account.saveSecureValue#899fe31d`.
// Securely save Telegram Passport¹ document, for more info see the passport docs »²
//
// Links:
//  1) https://core.telegram.org/passport
//  2) https://core.telegram.org/passport/encryption#encryption
//
// See https://core.telegram.org/method/account.saveSecureValue for reference.
type AccountSaveSecureValueRequest struct {
	// Secure value, for more info see the passport docs »¹
	//
	// Links:
	//  1) https://core.telegram.org/passport/encryption#encryption
	Value InputSecureValue
	// Passport secret hash, for more info see the passport docs »¹
	//
	// Links:
	//  1) https://core.telegram.org/passport/encryption#encryption
	SecureSecretID int64
}

// AccountSaveSecureValueRequestTypeID is TL type id of AccountSaveSecureValueRequest.
const AccountSaveSecureValueRequestTypeID = 0x899fe31d

// Encode implements bin.Encoder.
func (s *AccountSaveSecureValueRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.saveSecureValue#899fe31d as nil")
	}
	b.PutID(AccountSaveSecureValueRequestTypeID)
	if err := s.Value.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.saveSecureValue#899fe31d: field value: %w", err)
	}
	b.PutLong(s.SecureSecretID)
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSaveSecureValueRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.saveSecureValue#899fe31d to nil")
	}
	if err := b.ConsumeID(AccountSaveSecureValueRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.saveSecureValue#899fe31d: %w", err)
	}
	{
		if err := s.Value.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.saveSecureValue#899fe31d: field value: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.saveSecureValue#899fe31d: field secure_secret_id: %w", err)
		}
		s.SecureSecretID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSaveSecureValueRequest.
var (
	_ bin.Encoder = &AccountSaveSecureValueRequest{}
	_ bin.Decoder = &AccountSaveSecureValueRequest{}
)

// AccountSaveSecureValue invokes method account.saveSecureValue#899fe31d returning error if any.
// Securely save Telegram Passport¹ document, for more info see the passport docs »²
//
// Links:
//  1) https://core.telegram.org/passport
//  2) https://core.telegram.org/passport/encryption#encryption
//
// See https://core.telegram.org/method/account.saveSecureValue for reference.
func (c *Client) AccountSaveSecureValue(ctx context.Context, request *AccountSaveSecureValueRequest) (*SecureValue, error) {
	var result SecureValue

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
