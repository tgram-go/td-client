// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// AddLocalMessageRequest represents TL type `addLocalMessage#88db87fb`.
type AddLocalMessageRequest struct {
	// Target chat
	ChatID int64
	// Identifier of the sender of the message
	SenderID MessageSenderClass
	// Identifier of the message to reply to or 0
	ReplyToMessageID int64
	// Pass true to disable notification for the message
	DisableNotification bool
	// The content of the message to be added
	InputMessageContent InputMessageContentClass
}

// AddLocalMessageRequestTypeID is TL type id of AddLocalMessageRequest.
const AddLocalMessageRequestTypeID = 0x88db87fb

// Ensuring interfaces in compile-time for AddLocalMessageRequest.
var (
	_ bin.Encoder     = &AddLocalMessageRequest{}
	_ bin.Decoder     = &AddLocalMessageRequest{}
	_ bin.BareEncoder = &AddLocalMessageRequest{}
	_ bin.BareDecoder = &AddLocalMessageRequest{}
)

func (a *AddLocalMessageRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.ChatID == 0) {
		return false
	}
	if !(a.SenderID == nil) {
		return false
	}
	if !(a.ReplyToMessageID == 0) {
		return false
	}
	if !(a.DisableNotification == false) {
		return false
	}
	if !(a.InputMessageContent == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AddLocalMessageRequest) String() string {
	if a == nil {
		return "AddLocalMessageRequest(nil)"
	}
	type Alias AddLocalMessageRequest
	return fmt.Sprintf("AddLocalMessageRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AddLocalMessageRequest) TypeID() uint32 {
	return AddLocalMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AddLocalMessageRequest) TypeName() string {
	return "addLocalMessage"
}

// TypeInfo returns info about TL type.
func (a *AddLocalMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "addLocalMessage",
		ID:   AddLocalMessageRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "SenderID",
			SchemaName: "sender_id",
		},
		{
			Name:       "ReplyToMessageID",
			SchemaName: "reply_to_message_id",
		},
		{
			Name:       "DisableNotification",
			SchemaName: "disable_notification",
		},
		{
			Name:       "InputMessageContent",
			SchemaName: "input_message_content",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AddLocalMessageRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addLocalMessage#88db87fb as nil")
	}
	b.PutID(AddLocalMessageRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AddLocalMessageRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addLocalMessage#88db87fb as nil")
	}
	b.PutInt53(a.ChatID)
	if a.SenderID == nil {
		return fmt.Errorf("unable to encode addLocalMessage#88db87fb: field sender_id is nil")
	}
	if err := a.SenderID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode addLocalMessage#88db87fb: field sender_id: %w", err)
	}
	b.PutInt53(a.ReplyToMessageID)
	b.PutBool(a.DisableNotification)
	if a.InputMessageContent == nil {
		return fmt.Errorf("unable to encode addLocalMessage#88db87fb: field input_message_content is nil")
	}
	if err := a.InputMessageContent.Encode(b); err != nil {
		return fmt.Errorf("unable to encode addLocalMessage#88db87fb: field input_message_content: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AddLocalMessageRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addLocalMessage#88db87fb to nil")
	}
	if err := b.ConsumeID(AddLocalMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode addLocalMessage#88db87fb: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AddLocalMessageRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addLocalMessage#88db87fb to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode addLocalMessage#88db87fb: field chat_id: %w", err)
		}
		a.ChatID = value
	}
	{
		value, err := DecodeMessageSender(b)
		if err != nil {
			return fmt.Errorf("unable to decode addLocalMessage#88db87fb: field sender_id: %w", err)
		}
		a.SenderID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode addLocalMessage#88db87fb: field reply_to_message_id: %w", err)
		}
		a.ReplyToMessageID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode addLocalMessage#88db87fb: field disable_notification: %w", err)
		}
		a.DisableNotification = value
	}
	{
		value, err := DecodeInputMessageContent(b)
		if err != nil {
			return fmt.Errorf("unable to decode addLocalMessage#88db87fb: field input_message_content: %w", err)
		}
		a.InputMessageContent = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AddLocalMessageRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode addLocalMessage#88db87fb as nil")
	}
	b.ObjStart()
	b.PutID("addLocalMessage")
	b.FieldStart("chat_id")
	b.PutInt53(a.ChatID)
	b.FieldStart("sender_id")
	if a.SenderID == nil {
		return fmt.Errorf("unable to encode addLocalMessage#88db87fb: field sender_id is nil")
	}
	if err := a.SenderID.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode addLocalMessage#88db87fb: field sender_id: %w", err)
	}
	b.FieldStart("reply_to_message_id")
	b.PutInt53(a.ReplyToMessageID)
	b.FieldStart("disable_notification")
	b.PutBool(a.DisableNotification)
	b.FieldStart("input_message_content")
	if a.InputMessageContent == nil {
		return fmt.Errorf("unable to encode addLocalMessage#88db87fb: field input_message_content is nil")
	}
	if err := a.InputMessageContent.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode addLocalMessage#88db87fb: field input_message_content: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AddLocalMessageRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode addLocalMessage#88db87fb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("addLocalMessage"); err != nil {
				return fmt.Errorf("unable to decode addLocalMessage#88db87fb: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode addLocalMessage#88db87fb: field chat_id: %w", err)
			}
			a.ChatID = value
		case "sender_id":
			value, err := DecodeTDLibJSONMessageSender(b)
			if err != nil {
				return fmt.Errorf("unable to decode addLocalMessage#88db87fb: field sender_id: %w", err)
			}
			a.SenderID = value
		case "reply_to_message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode addLocalMessage#88db87fb: field reply_to_message_id: %w", err)
			}
			a.ReplyToMessageID = value
		case "disable_notification":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode addLocalMessage#88db87fb: field disable_notification: %w", err)
			}
			a.DisableNotification = value
		case "input_message_content":
			value, err := DecodeTDLibJSONInputMessageContent(b)
			if err != nil {
				return fmt.Errorf("unable to decode addLocalMessage#88db87fb: field input_message_content: %w", err)
			}
			a.InputMessageContent = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (a *AddLocalMessageRequest) GetChatID() (value int64) {
	return a.ChatID
}

// GetSenderID returns value of SenderID field.
func (a *AddLocalMessageRequest) GetSenderID() (value MessageSenderClass) {
	return a.SenderID
}

// GetReplyToMessageID returns value of ReplyToMessageID field.
func (a *AddLocalMessageRequest) GetReplyToMessageID() (value int64) {
	return a.ReplyToMessageID
}

// GetDisableNotification returns value of DisableNotification field.
func (a *AddLocalMessageRequest) GetDisableNotification() (value bool) {
	return a.DisableNotification
}

// GetInputMessageContent returns value of InputMessageContent field.
func (a *AddLocalMessageRequest) GetInputMessageContent() (value InputMessageContentClass) {
	return a.InputMessageContent
}

// AddLocalMessage invokes method addLocalMessage#88db87fb returning error if any.
func (c *Client) AddLocalMessage(ctx context.Context, request *AddLocalMessageRequest) (*Message, error) {
	var result Message

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}