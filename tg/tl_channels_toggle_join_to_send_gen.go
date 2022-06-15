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

// ChannelsToggleJoinToSendRequest represents TL type `channels.toggleJoinToSend#e4cb9580`.
//
// See https://core.telegram.org/method/channels.toggleJoinToSend for reference.
type ChannelsToggleJoinToSendRequest struct {
	// Channel field of ChannelsToggleJoinToSendRequest.
	Channel InputChannelClass
	// Enabled field of ChannelsToggleJoinToSendRequest.
	Enabled bool
}

// ChannelsToggleJoinToSendRequestTypeID is TL type id of ChannelsToggleJoinToSendRequest.
const ChannelsToggleJoinToSendRequestTypeID = 0xe4cb9580

// Ensuring interfaces in compile-time for ChannelsToggleJoinToSendRequest.
var (
	_ bin.Encoder     = &ChannelsToggleJoinToSendRequest{}
	_ bin.Decoder     = &ChannelsToggleJoinToSendRequest{}
	_ bin.BareEncoder = &ChannelsToggleJoinToSendRequest{}
	_ bin.BareDecoder = &ChannelsToggleJoinToSendRequest{}
)

func (t *ChannelsToggleJoinToSendRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Channel == nil) {
		return false
	}
	if !(t.Enabled == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ChannelsToggleJoinToSendRequest) String() string {
	if t == nil {
		return "ChannelsToggleJoinToSendRequest(nil)"
	}
	type Alias ChannelsToggleJoinToSendRequest
	return fmt.Sprintf("ChannelsToggleJoinToSendRequest%+v", Alias(*t))
}

// FillFrom fills ChannelsToggleJoinToSendRequest from given interface.
func (t *ChannelsToggleJoinToSendRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetEnabled() (value bool)
}) {
	t.Channel = from.GetChannel()
	t.Enabled = from.GetEnabled()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsToggleJoinToSendRequest) TypeID() uint32 {
	return ChannelsToggleJoinToSendRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsToggleJoinToSendRequest) TypeName() string {
	return "channels.toggleJoinToSend"
}

// TypeInfo returns info about TL type.
func (t *ChannelsToggleJoinToSendRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.toggleJoinToSend",
		ID:   ChannelsToggleJoinToSendRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Enabled",
			SchemaName: "enabled",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ChannelsToggleJoinToSendRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.toggleJoinToSend#e4cb9580 as nil")
	}
	b.PutID(ChannelsToggleJoinToSendRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ChannelsToggleJoinToSendRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.toggleJoinToSend#e4cb9580 as nil")
	}
	if t.Channel == nil {
		return fmt.Errorf("unable to encode channels.toggleJoinToSend#e4cb9580: field channel is nil")
	}
	if err := t.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.toggleJoinToSend#e4cb9580: field channel: %w", err)
	}
	b.PutBool(t.Enabled)
	return nil
}

// Decode implements bin.Decoder.
func (t *ChannelsToggleJoinToSendRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.toggleJoinToSend#e4cb9580 to nil")
	}
	if err := b.ConsumeID(ChannelsToggleJoinToSendRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.toggleJoinToSend#e4cb9580: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ChannelsToggleJoinToSendRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.toggleJoinToSend#e4cb9580 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleJoinToSend#e4cb9580: field channel: %w", err)
		}
		t.Channel = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleJoinToSend#e4cb9580: field enabled: %w", err)
		}
		t.Enabled = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (t *ChannelsToggleJoinToSendRequest) GetChannel() (value InputChannelClass) {
	if t == nil {
		return
	}
	return t.Channel
}

// GetEnabled returns value of Enabled field.
func (t *ChannelsToggleJoinToSendRequest) GetEnabled() (value bool) {
	if t == nil {
		return
	}
	return t.Enabled
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (t *ChannelsToggleJoinToSendRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return t.Channel.AsNotEmpty()
}

// ChannelsToggleJoinToSend invokes method channels.toggleJoinToSend#e4cb9580 returning error if any.
//
// See https://core.telegram.org/method/channels.toggleJoinToSend for reference.
func (c *Client) ChannelsToggleJoinToSend(ctx context.Context, request *ChannelsToggleJoinToSendRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
