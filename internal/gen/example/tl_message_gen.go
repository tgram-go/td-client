// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// Message represents TL type `message#ec200d96`.
//
// See https://localhost:80/doc/constructor/message for reference.
type Message struct {
	// Err field of Message.
	Err Error
}

// MessageTypeID is TL type id of Message.
const MessageTypeID = 0xec200d96

// Encode implements bin.Encoder.
func (m *Message) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode message#ec200d96 as nil")
	}
	b.PutID(MessageTypeID)
	if err := m.Err.Encode(b); err != nil {
		return fmt.Errorf("unable to encode message#ec200d96: field err: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *Message) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode message#ec200d96 to nil")
	}
	if err := b.ConsumeID(MessageTypeID); err != nil {
		return fmt.Errorf("unable to decode message#ec200d96: %w", err)
	}

	{
		if err := m.Err.Decode(b); err != nil {
			return fmt.Errorf("unable to decode message#ec200d96: field err: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for Message.
var (
	_ bin.Encoder = &Message{}
	_ bin.Decoder = &Message{}
)
