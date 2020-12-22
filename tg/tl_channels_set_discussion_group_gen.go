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

// ChannelsSetDiscussionGroupRequest represents TL type `channels.setDiscussionGroup#40582bb2`.
// Associate a group to a channel as discussion group¹ for that channel
//
// Links:
//  1) https://core.telegram.org/api/discussion
//
// See https://core.telegram.org/method/channels.setDiscussionGroup for reference.
type ChannelsSetDiscussionGroupRequest struct {
	// Channel
	Broadcast InputChannelClass
	// Discussion group¹ to associate to the channel
	//
	// Links:
	//  1) https://core.telegram.org/api/discussion
	Group InputChannelClass
}

// ChannelsSetDiscussionGroupRequestTypeID is TL type id of ChannelsSetDiscussionGroupRequest.
const ChannelsSetDiscussionGroupRequestTypeID = 0x40582bb2

// Encode implements bin.Encoder.
func (s *ChannelsSetDiscussionGroupRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode channels.setDiscussionGroup#40582bb2 as nil")
	}
	b.PutID(ChannelsSetDiscussionGroupRequestTypeID)
	if s.Broadcast == nil {
		return fmt.Errorf("unable to encode channels.setDiscussionGroup#40582bb2: field broadcast is nil")
	}
	if err := s.Broadcast.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.setDiscussionGroup#40582bb2: field broadcast: %w", err)
	}
	if s.Group == nil {
		return fmt.Errorf("unable to encode channels.setDiscussionGroup#40582bb2: field group is nil")
	}
	if err := s.Group.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.setDiscussionGroup#40582bb2: field group: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *ChannelsSetDiscussionGroupRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode channels.setDiscussionGroup#40582bb2 to nil")
	}
	if err := b.ConsumeID(ChannelsSetDiscussionGroupRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.setDiscussionGroup#40582bb2: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.setDiscussionGroup#40582bb2: field broadcast: %w", err)
		}
		s.Broadcast = value
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.setDiscussionGroup#40582bb2: field group: %w", err)
		}
		s.Group = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsSetDiscussionGroupRequest.
var (
	_ bin.Encoder = &ChannelsSetDiscussionGroupRequest{}
	_ bin.Decoder = &ChannelsSetDiscussionGroupRequest{}
)

// ChannelsSetDiscussionGroup invokes method channels.setDiscussionGroup#40582bb2 returning error if any.
// Associate a group to a channel as discussion group¹ for that channel
//
// Links:
//  1) https://core.telegram.org/api/discussion
//
// See https://core.telegram.org/method/channels.setDiscussionGroup for reference.
func (c *Client) ChannelsSetDiscussionGroup(ctx context.Context, request *ChannelsSetDiscussionGroupRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
