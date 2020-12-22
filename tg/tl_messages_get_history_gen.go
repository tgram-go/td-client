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

// MessagesGetHistoryRequest represents TL type `messages.getHistory#dcbb8260`.
// Gets back the conversation history with one interlocutor / within a chat
//
// See https://core.telegram.org/method/messages.getHistory for reference.
type MessagesGetHistoryRequest struct {
	// Target peer
	Peer InputPeerClass
	// Only return messages starting from the specified message ID
	OffsetID int
	// Only return messages sent before the specified date
	OffsetDate int
	// Number of list elements to be skipped, negative values are also accepted.
	AddOffset int
	// Number of results to return
	Limit int
	// If a positive value was transferred, the method will return only messages with IDs less than max_id
	MaxID int
	// If a positive value was transferred, the method will return only messages with IDs more than min_id
	MinID int
	// Result hash¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Hash int
}

// MessagesGetHistoryRequestTypeID is TL type id of MessagesGetHistoryRequest.
const MessagesGetHistoryRequestTypeID = 0xdcbb8260

// Encode implements bin.Encoder.
func (g *MessagesGetHistoryRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getHistory#dcbb8260 as nil")
	}
	b.PutID(MessagesGetHistoryRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getHistory#dcbb8260: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getHistory#dcbb8260: field peer: %w", err)
	}
	b.PutInt(g.OffsetID)
	b.PutInt(g.OffsetDate)
	b.PutInt(g.AddOffset)
	b.PutInt(g.Limit)
	b.PutInt(g.MaxID)
	b.PutInt(g.MinID)
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetHistoryRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getHistory#dcbb8260 to nil")
	}
	if err := b.ConsumeID(MessagesGetHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: field offset_date: %w", err)
		}
		g.OffsetDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: field add_offset: %w", err)
		}
		g.AddOffset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: field max_id: %w", err)
		}
		g.MaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: field min_id: %w", err)
		}
		g.MinID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetHistoryRequest.
var (
	_ bin.Encoder = &MessagesGetHistoryRequest{}
	_ bin.Decoder = &MessagesGetHistoryRequest{}
)

// MessagesGetHistory invokes method messages.getHistory#dcbb8260 returning error if any.
// Gets back the conversation history with one interlocutor / within a chat
//
// See https://core.telegram.org/method/messages.getHistory for reference.
func (c *Client) MessagesGetHistory(ctx context.Context, request *MessagesGetHistoryRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
