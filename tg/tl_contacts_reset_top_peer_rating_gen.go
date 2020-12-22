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

// ContactsResetTopPeerRatingRequest represents TL type `contacts.resetTopPeerRating#1ae373ac`.
// Reset rating¹ of top peer
//
// Links:
//  1) https://core.telegram.org/api/top-rating
//
// See https://core.telegram.org/method/contacts.resetTopPeerRating for reference.
type ContactsResetTopPeerRatingRequest struct {
	// Top peer category
	Category TopPeerCategoryClass
	// Peer whose rating should be reset
	Peer InputPeerClass
}

// ContactsResetTopPeerRatingRequestTypeID is TL type id of ContactsResetTopPeerRatingRequest.
const ContactsResetTopPeerRatingRequestTypeID = 0x1ae373ac

// Encode implements bin.Encoder.
func (r *ContactsResetTopPeerRatingRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode contacts.resetTopPeerRating#1ae373ac as nil")
	}
	b.PutID(ContactsResetTopPeerRatingRequestTypeID)
	if r.Category == nil {
		return fmt.Errorf("unable to encode contacts.resetTopPeerRating#1ae373ac: field category is nil")
	}
	if err := r.Category.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.resetTopPeerRating#1ae373ac: field category: %w", err)
	}
	if r.Peer == nil {
		return fmt.Errorf("unable to encode contacts.resetTopPeerRating#1ae373ac: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.resetTopPeerRating#1ae373ac: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ContactsResetTopPeerRatingRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode contacts.resetTopPeerRating#1ae373ac to nil")
	}
	if err := b.ConsumeID(ContactsResetTopPeerRatingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.resetTopPeerRating#1ae373ac: %w", err)
	}
	{
		value, err := DecodeTopPeerCategory(b)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.resetTopPeerRating#1ae373ac: field category: %w", err)
		}
		r.Category = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.resetTopPeerRating#1ae373ac: field peer: %w", err)
		}
		r.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsResetTopPeerRatingRequest.
var (
	_ bin.Encoder = &ContactsResetTopPeerRatingRequest{}
	_ bin.Decoder = &ContactsResetTopPeerRatingRequest{}
)

// ContactsResetTopPeerRating invokes method contacts.resetTopPeerRating#1ae373ac returning error if any.
// Reset rating¹ of top peer
//
// Links:
//  1) https://core.telegram.org/api/top-rating
//
// See https://core.telegram.org/method/contacts.resetTopPeerRating for reference.
func (c *Client) ContactsResetTopPeerRating(ctx context.Context, request *ContactsResetTopPeerRatingRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
