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

// MessagesGetStatsURLRequest represents TL type `messages.getStatsURL#812c2ae6`.
// Returns URL with the chat statistics. Currently this method can be used only for channels
//
// See https://core.telegram.org/method/messages.getStatsURL for reference.
type MessagesGetStatsURLRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Pass true if a URL with the dark theme must be returned
	Dark bool
	// Chat identifier
	Peer InputPeerClass
	// Parameters from tg://statsrefresh?params=****** link
	Params string
}

// MessagesGetStatsURLRequestTypeID is TL type id of MessagesGetStatsURLRequest.
const MessagesGetStatsURLRequestTypeID = 0x812c2ae6

// Encode implements bin.Encoder.
func (g *MessagesGetStatsURLRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getStatsURL#812c2ae6 as nil")
	}
	b.PutID(MessagesGetStatsURLRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getStatsURL#812c2ae6: field flags: %w", err)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getStatsURL#812c2ae6: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getStatsURL#812c2ae6: field peer: %w", err)
	}
	b.PutString(g.Params)
	return nil
}

// SetDark sets value of Dark conditional field.
func (g *MessagesGetStatsURLRequest) SetDark(value bool) {
	if value {
		g.Flags.Set(0)
	} else {
		g.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (g *MessagesGetStatsURLRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getStatsURL#812c2ae6 to nil")
	}
	if err := b.ConsumeID(MessagesGetStatsURLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getStatsURL#812c2ae6: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getStatsURL#812c2ae6: field flags: %w", err)
		}
	}
	g.Dark = g.Flags.Has(0)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getStatsURL#812c2ae6: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getStatsURL#812c2ae6: field params: %w", err)
		}
		g.Params = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetStatsURLRequest.
var (
	_ bin.Encoder = &MessagesGetStatsURLRequest{}
	_ bin.Decoder = &MessagesGetStatsURLRequest{}
)

// MessagesGetStatsURL invokes method messages.getStatsURL#812c2ae6 returning error if any.
// Returns URL with the chat statistics. Currently this method can be used only for channels
//
// See https://core.telegram.org/method/messages.getStatsURL for reference.
func (c *Client) MessagesGetStatsURL(ctx context.Context, request *MessagesGetStatsURLRequest) (*StatsURL, error) {
	var result StatsURL

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
