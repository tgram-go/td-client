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

// AccountContentSettings represents TL type `account.contentSettings#57e28221`.
// Sensitive content settings
//
// See https://core.telegram.org/constructor/account.contentSettings for reference.
type AccountContentSettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether viewing of sensitive (NSFW) content is enabled
	SensitiveEnabled bool
	// Whether the current client can change the sensitive content settings to view NSFW content
	SensitiveCanChange bool
}

// AccountContentSettingsTypeID is TL type id of AccountContentSettings.
const AccountContentSettingsTypeID = 0x57e28221

// Encode implements bin.Encoder.
func (c *AccountContentSettings) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.contentSettings#57e28221 as nil")
	}
	b.PutID(AccountContentSettingsTypeID)
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.contentSettings#57e28221: field flags: %w", err)
	}
	return nil
}

// SetSensitiveEnabled sets value of SensitiveEnabled conditional field.
func (c *AccountContentSettings) SetSensitiveEnabled(value bool) {
	if value {
		c.Flags.Set(0)
	} else {
		c.Flags.Unset(0)
	}
}

// SetSensitiveCanChange sets value of SensitiveCanChange conditional field.
func (c *AccountContentSettings) SetSensitiveCanChange(value bool) {
	if value {
		c.Flags.Set(1)
	} else {
		c.Flags.Unset(1)
	}
}

// Decode implements bin.Decoder.
func (c *AccountContentSettings) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.contentSettings#57e28221 to nil")
	}
	if err := b.ConsumeID(AccountContentSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode account.contentSettings#57e28221: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.contentSettings#57e28221: field flags: %w", err)
		}
	}
	c.SensitiveEnabled = c.Flags.Has(0)
	c.SensitiveCanChange = c.Flags.Has(1)
	return nil
}

// Ensuring interfaces in compile-time for AccountContentSettings.
var (
	_ bin.Encoder = &AccountContentSettings{}
	_ bin.Decoder = &AccountContentSettings{}
)
