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

// CodeSettings represents TL type `codeSettings#debebe83`.
// Settings used by telegram servers for sending the confirm code.
// Example implementations: telegram for android¹, tdlib².
//
// Links:
//  1) https://github.com/DrKLO/Telegram/blob/master/TMessagesProj/src/main/java/org/telegram/ui/LoginActivity.java
//  2) https://github.com/tdlib/td/tree/master/td/telegram/SendCodeHelper.cpp
//
// See https://core.telegram.org/constructor/codeSettings for reference.
type CodeSettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to allow phone verification via phone calls¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/auth
	AllowFlashcall bool
	// Pass true if the phone number is used on the current device. Ignored if allow_flashcall is not set.
	CurrentNumber bool
	// If a token that will be included in eventually sent SMSs is required: required in newer versions of android, to use the android SMS receiver APIs¹
	//
	// Links:
	//  1) https://developers.google.com/identity/sms-retriever/overview
	AllowAppHash bool
}

// CodeSettingsTypeID is TL type id of CodeSettings.
const CodeSettingsTypeID = 0xdebebe83

// Encode implements bin.Encoder.
func (c *CodeSettings) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode codeSettings#debebe83 as nil")
	}
	b.PutID(CodeSettingsTypeID)
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode codeSettings#debebe83: field flags: %w", err)
	}
	return nil
}

// SetAllowFlashcall sets value of AllowFlashcall conditional field.
func (c *CodeSettings) SetAllowFlashcall(value bool) {
	if value {
		c.Flags.Set(0)
	} else {
		c.Flags.Unset(0)
	}
}

// SetCurrentNumber sets value of CurrentNumber conditional field.
func (c *CodeSettings) SetCurrentNumber(value bool) {
	if value {
		c.Flags.Set(1)
	} else {
		c.Flags.Unset(1)
	}
}

// SetAllowAppHash sets value of AllowAppHash conditional field.
func (c *CodeSettings) SetAllowAppHash(value bool) {
	if value {
		c.Flags.Set(4)
	} else {
		c.Flags.Unset(4)
	}
}

// Decode implements bin.Decoder.
func (c *CodeSettings) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode codeSettings#debebe83 to nil")
	}
	if err := b.ConsumeID(CodeSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode codeSettings#debebe83: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode codeSettings#debebe83: field flags: %w", err)
		}
	}
	c.AllowFlashcall = c.Flags.Has(0)
	c.CurrentNumber = c.Flags.Has(1)
	c.AllowAppHash = c.Flags.Has(4)
	return nil
}

// Ensuring interfaces in compile-time for CodeSettings.
var (
	_ bin.Encoder = &CodeSettings{}
	_ bin.Decoder = &CodeSettings{}
)
