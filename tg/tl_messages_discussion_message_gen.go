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

// MessagesDiscussionMessage represents TL type `messages.discussionMessage#f5dd8f9d`.
// Information about a message thread¹
//
// Links:
//  1) https://core.telegram.org/api/threads
//
// See https://core.telegram.org/constructor/messages.discussionMessage for reference.
type MessagesDiscussionMessage struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Discussion messages
	Messages []MessageClass
	// Message ID of latest reply in this thread¹
	//
	// Links:
	//  1) https://core.telegram.org/api/threads
	//
	// Use SetMaxID and GetMaxID helpers.
	MaxID int
	// Message ID of latest read incoming message in this thread¹
	//
	// Links:
	//  1) https://core.telegram.org/api/threads
	//
	// Use SetReadInboxMaxID and GetReadInboxMaxID helpers.
	ReadInboxMaxID int
	// Message ID of latest read outgoing message in this thread¹
	//
	// Links:
	//  1) https://core.telegram.org/api/threads
	//
	// Use SetReadOutboxMaxID and GetReadOutboxMaxID helpers.
	ReadOutboxMaxID int
	// Chats mentioned in constructor
	Chats []ChatClass
	// Users mentioned in constructor
	Users []UserClass
}

// MessagesDiscussionMessageTypeID is TL type id of MessagesDiscussionMessage.
const MessagesDiscussionMessageTypeID = 0xf5dd8f9d

// Encode implements bin.Encoder.
func (d *MessagesDiscussionMessage) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.discussionMessage#f5dd8f9d as nil")
	}
	b.PutID(MessagesDiscussionMessageTypeID)
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.discussionMessage#f5dd8f9d: field flags: %w", err)
	}
	b.PutVectorHeader(len(d.Messages))
	for idx, v := range d.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode messages.discussionMessage#f5dd8f9d: field messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.discussionMessage#f5dd8f9d: field messages element with index %d: %w", idx, err)
		}
	}
	if d.Flags.Has(0) {
		b.PutInt(d.MaxID)
	}
	if d.Flags.Has(1) {
		b.PutInt(d.ReadInboxMaxID)
	}
	if d.Flags.Has(2) {
		b.PutInt(d.ReadOutboxMaxID)
	}
	b.PutVectorHeader(len(d.Chats))
	for idx, v := range d.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.discussionMessage#f5dd8f9d: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.discussionMessage#f5dd8f9d: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Users))
	for idx, v := range d.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.discussionMessage#f5dd8f9d: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.discussionMessage#f5dd8f9d: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// SetMaxID sets value of MaxID conditional field.
func (d *MessagesDiscussionMessage) SetMaxID(value int) {
	d.Flags.Set(0)
	d.MaxID = value
}

// GetMaxID returns value of MaxID conditional field and
// boolean which is true if field was set.
func (d *MessagesDiscussionMessage) GetMaxID() (value int, ok bool) {
	if !d.Flags.Has(0) {
		return value, false
	}
	return d.MaxID, true
}

// SetReadInboxMaxID sets value of ReadInboxMaxID conditional field.
func (d *MessagesDiscussionMessage) SetReadInboxMaxID(value int) {
	d.Flags.Set(1)
	d.ReadInboxMaxID = value
}

// GetReadInboxMaxID returns value of ReadInboxMaxID conditional field and
// boolean which is true if field was set.
func (d *MessagesDiscussionMessage) GetReadInboxMaxID() (value int, ok bool) {
	if !d.Flags.Has(1) {
		return value, false
	}
	return d.ReadInboxMaxID, true
}

// SetReadOutboxMaxID sets value of ReadOutboxMaxID conditional field.
func (d *MessagesDiscussionMessage) SetReadOutboxMaxID(value int) {
	d.Flags.Set(2)
	d.ReadOutboxMaxID = value
}

// GetReadOutboxMaxID returns value of ReadOutboxMaxID conditional field and
// boolean which is true if field was set.
func (d *MessagesDiscussionMessage) GetReadOutboxMaxID() (value int, ok bool) {
	if !d.Flags.Has(2) {
		return value, false
	}
	return d.ReadOutboxMaxID, true
}

// Decode implements bin.Decoder.
func (d *MessagesDiscussionMessage) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.discussionMessage#f5dd8f9d to nil")
	}
	if err := b.ConsumeID(MessagesDiscussionMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.discussionMessage#f5dd8f9d: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.discussionMessage#f5dd8f9d: field flags: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.discussionMessage#f5dd8f9d: field messages: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.discussionMessage#f5dd8f9d: field messages: %w", err)
			}
			d.Messages = append(d.Messages, value)
		}
	}
	if d.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.discussionMessage#f5dd8f9d: field max_id: %w", err)
		}
		d.MaxID = value
	}
	if d.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.discussionMessage#f5dd8f9d: field read_inbox_max_id: %w", err)
		}
		d.ReadInboxMaxID = value
	}
	if d.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.discussionMessage#f5dd8f9d: field read_outbox_max_id: %w", err)
		}
		d.ReadOutboxMaxID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.discussionMessage#f5dd8f9d: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.discussionMessage#f5dd8f9d: field chats: %w", err)
			}
			d.Chats = append(d.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.discussionMessage#f5dd8f9d: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.discussionMessage#f5dd8f9d: field users: %w", err)
			}
			d.Users = append(d.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesDiscussionMessage.
var (
	_ bin.Encoder = &MessagesDiscussionMessage{}
	_ bin.Decoder = &MessagesDiscussionMessage{}
)
