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

// MessagesToggleDialogPinRequest represents TL type `messages.toggleDialogPin#a731e257`.
// Pin/unpin a dialog
//
// See https://core.telegram.org/method/messages.toggleDialogPin for reference.
type MessagesToggleDialogPinRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to pin or unpin the dialog
	Pinned bool
	// The dialog to pin
	Peer InputDialogPeerClass
}

// MessagesToggleDialogPinRequestTypeID is TL type id of MessagesToggleDialogPinRequest.
const MessagesToggleDialogPinRequestTypeID = 0xa731e257

// Encode implements bin.Encoder.
func (t *MessagesToggleDialogPinRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.toggleDialogPin#a731e257 as nil")
	}
	b.PutID(MessagesToggleDialogPinRequestTypeID)
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.toggleDialogPin#a731e257: field flags: %w", err)
	}
	if t.Peer == nil {
		return fmt.Errorf("unable to encode messages.toggleDialogPin#a731e257: field peer is nil")
	}
	if err := t.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.toggleDialogPin#a731e257: field peer: %w", err)
	}
	return nil
}

// SetPinned sets value of Pinned conditional field.
func (t *MessagesToggleDialogPinRequest) SetPinned(value bool) {
	if value {
		t.Flags.Set(0)
	} else {
		t.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (t *MessagesToggleDialogPinRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.toggleDialogPin#a731e257 to nil")
	}
	if err := b.ConsumeID(MessagesToggleDialogPinRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.toggleDialogPin#a731e257: %w", err)
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.toggleDialogPin#a731e257: field flags: %w", err)
		}
	}
	t.Pinned = t.Flags.Has(0)
	{
		value, err := DecodeInputDialogPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.toggleDialogPin#a731e257: field peer: %w", err)
		}
		t.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesToggleDialogPinRequest.
var (
	_ bin.Encoder = &MessagesToggleDialogPinRequest{}
	_ bin.Decoder = &MessagesToggleDialogPinRequest{}
)

// MessagesToggleDialogPin invokes method messages.toggleDialogPin#a731e257 returning error if any.
// Pin/unpin a dialog
//
// See https://core.telegram.org/method/messages.toggleDialogPin for reference.
func (c *Client) MessagesToggleDialogPin(ctx context.Context, request *MessagesToggleDialogPinRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
