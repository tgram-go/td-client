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

// AccountPasswordSettings represents TL type `account.passwordSettings#9a5c33e5`.
// Private info associated to the password info (recovery email, telegram passport¹ info & so on)
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/constructor/account.passwordSettings for reference.
type AccountPasswordSettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// 2FA Recovery email¹
	//
	// Links:
	//  1) https://core.telegram.org/api/srp#email-verification
	//
	// Use SetEmail and GetEmail helpers.
	Email string
	// Telegram passport¹ settings
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetSecureSettings and GetSecureSettings helpers.
	SecureSettings SecureSecretSettings
}

// AccountPasswordSettingsTypeID is TL type id of AccountPasswordSettings.
const AccountPasswordSettingsTypeID = 0x9a5c33e5

// Encode implements bin.Encoder.
func (p *AccountPasswordSettings) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode account.passwordSettings#9a5c33e5 as nil")
	}
	b.PutID(AccountPasswordSettingsTypeID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.passwordSettings#9a5c33e5: field flags: %w", err)
	}
	if p.Flags.Has(0) {
		b.PutString(p.Email)
	}
	if p.Flags.Has(1) {
		if err := p.SecureSettings.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.passwordSettings#9a5c33e5: field secure_settings: %w", err)
		}
	}
	return nil
}

// SetEmail sets value of Email conditional field.
func (p *AccountPasswordSettings) SetEmail(value string) {
	p.Flags.Set(0)
	p.Email = value
}

// GetEmail returns value of Email conditional field and
// boolean which is true if field was set.
func (p *AccountPasswordSettings) GetEmail() (value string, ok bool) {
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.Email, true
}

// SetSecureSettings sets value of SecureSettings conditional field.
func (p *AccountPasswordSettings) SetSecureSettings(value SecureSecretSettings) {
	p.Flags.Set(1)
	p.SecureSettings = value
}

// GetSecureSettings returns value of SecureSettings conditional field and
// boolean which is true if field was set.
func (p *AccountPasswordSettings) GetSecureSettings() (value SecureSecretSettings, ok bool) {
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.SecureSettings, true
}

// Decode implements bin.Decoder.
func (p *AccountPasswordSettings) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode account.passwordSettings#9a5c33e5 to nil")
	}
	if err := b.ConsumeID(AccountPasswordSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode account.passwordSettings#9a5c33e5: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.passwordSettings#9a5c33e5: field flags: %w", err)
		}
	}
	if p.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.passwordSettings#9a5c33e5: field email: %w", err)
		}
		p.Email = value
	}
	if p.Flags.Has(1) {
		if err := p.SecureSettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.passwordSettings#9a5c33e5: field secure_settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountPasswordSettings.
var (
	_ bin.Encoder = &AccountPasswordSettings{}
	_ bin.Decoder = &AccountPasswordSettings{}
)
