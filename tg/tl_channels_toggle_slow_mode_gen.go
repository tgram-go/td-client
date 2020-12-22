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

// ChannelsToggleSlowModeRequest represents TL type `channels.toggleSlowMode#edd49ef0`.
// Toggle supergroup slow mode: if enabled, users will only be able to send one message every seconds seconds
//
// See https://core.telegram.org/method/channels.toggleSlowMode for reference.
type ChannelsToggleSlowModeRequest struct {
	// The supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass
	// Users will only be able to send one message every seconds seconds, 0 to disable the limitation
	Seconds int
}

// ChannelsToggleSlowModeRequestTypeID is TL type id of ChannelsToggleSlowModeRequest.
const ChannelsToggleSlowModeRequestTypeID = 0xedd49ef0

// Encode implements bin.Encoder.
func (t *ChannelsToggleSlowModeRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.toggleSlowMode#edd49ef0 as nil")
	}
	b.PutID(ChannelsToggleSlowModeRequestTypeID)
	if t.Channel == nil {
		return fmt.Errorf("unable to encode channels.toggleSlowMode#edd49ef0: field channel is nil")
	}
	if err := t.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.toggleSlowMode#edd49ef0: field channel: %w", err)
	}
	b.PutInt(t.Seconds)
	return nil
}

// Decode implements bin.Decoder.
func (t *ChannelsToggleSlowModeRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.toggleSlowMode#edd49ef0 to nil")
	}
	if err := b.ConsumeID(ChannelsToggleSlowModeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.toggleSlowMode#edd49ef0: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleSlowMode#edd49ef0: field channel: %w", err)
		}
		t.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleSlowMode#edd49ef0: field seconds: %w", err)
		}
		t.Seconds = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsToggleSlowModeRequest.
var (
	_ bin.Encoder = &ChannelsToggleSlowModeRequest{}
	_ bin.Decoder = &ChannelsToggleSlowModeRequest{}
)

// ChannelsToggleSlowMode invokes method channels.toggleSlowMode#edd49ef0 returning error if any.
// Toggle supergroup slow mode: if enabled, users will only be able to send one message every seconds seconds
//
// See https://core.telegram.org/method/channels.toggleSlowMode for reference.
func (c *Client) ChannelsToggleSlowMode(ctx context.Context, request *ChannelsToggleSlowModeRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
