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

// MessagesMessageEditData represents TL type `messages.messageEditData#26b5dde6`.
// Message edit data for media
//
// See https://core.telegram.org/constructor/messages.messageEditData for reference.
type MessagesMessageEditData struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Media caption, if the specified media's caption can be edited
	Caption bool
}

// MessagesMessageEditDataTypeID is TL type id of MessagesMessageEditData.
const MessagesMessageEditDataTypeID = 0x26b5dde6

// Encode implements bin.Encoder.
func (m *MessagesMessageEditData) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.messageEditData#26b5dde6 as nil")
	}
	b.PutID(MessagesMessageEditDataTypeID)
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.messageEditData#26b5dde6: field flags: %w", err)
	}
	return nil
}

// SetCaption sets value of Caption conditional field.
func (m *MessagesMessageEditData) SetCaption(value bool) {
	if value {
		m.Flags.Set(0)
	} else {
		m.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (m *MessagesMessageEditData) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.messageEditData#26b5dde6 to nil")
	}
	if err := b.ConsumeID(MessagesMessageEditDataTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.messageEditData#26b5dde6: %w", err)
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.messageEditData#26b5dde6: field flags: %w", err)
		}
	}
	m.Caption = m.Flags.Has(0)
	return nil
}

// Ensuring interfaces in compile-time for MessagesMessageEditData.
var (
	_ bin.Encoder = &MessagesMessageEditData{}
	_ bin.Decoder = &MessagesMessageEditData{}
)
