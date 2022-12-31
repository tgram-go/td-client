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

// MessageLinkInfo represents TL type `messageLinkInfo#2b96fb50`.
type MessageLinkInfo struct {
	// True, if the link is a public link for a message or a forum topic in a chat
	IsPublic bool
	// If found, identifier of the chat to which the link points, 0 otherwise
	ChatID int64
	// If found, identifier of the message thread in which to open the message, or a forum
	// topic to open if the message is missing
	MessageThreadID int64
	// If found, the linked message; may be null
	Message Message
	// Timestamp from which the video/audio/video note/voice note playing must start, in
	// seconds; 0 if not specified. The media can be in the message content or in its web
	// page preview
	MediaTimestamp int32
	// True, if the whole media album to which the message belongs is linked
	ForAlbum bool
}

// MessageLinkInfoTypeID is TL type id of MessageLinkInfo.
const MessageLinkInfoTypeID = 0x2b96fb50

// Ensuring interfaces in compile-time for MessageLinkInfo.
var (
	_ bin.Encoder     = &MessageLinkInfo{}
	_ bin.Decoder     = &MessageLinkInfo{}
	_ bin.BareEncoder = &MessageLinkInfo{}
	_ bin.BareDecoder = &MessageLinkInfo{}
)

func (m *MessageLinkInfo) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.IsPublic == false) {
		return false
	}
	if !(m.ChatID == 0) {
		return false
	}
	if !(m.MessageThreadID == 0) {
		return false
	}
	if !(m.Message.Zero()) {
		return false
	}
	if !(m.MediaTimestamp == 0) {
		return false
	}
	if !(m.ForAlbum == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageLinkInfo) String() string {
	if m == nil {
		return "MessageLinkInfo(nil)"
	}
	type Alias MessageLinkInfo
	return fmt.Sprintf("MessageLinkInfo%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageLinkInfo) TypeID() uint32 {
	return MessageLinkInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageLinkInfo) TypeName() string {
	return "messageLinkInfo"
}

// TypeInfo returns info about TL type.
func (m *MessageLinkInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageLinkInfo",
		ID:   MessageLinkInfoTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "IsPublic",
			SchemaName: "is_public",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageThreadID",
			SchemaName: "message_thread_id",
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "MediaTimestamp",
			SchemaName: "media_timestamp",
		},
		{
			Name:       "ForAlbum",
			SchemaName: "for_album",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageLinkInfo) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageLinkInfo#2b96fb50 as nil")
	}
	b.PutID(MessageLinkInfoTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageLinkInfo) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageLinkInfo#2b96fb50 as nil")
	}
	b.PutBool(m.IsPublic)
	b.PutInt53(m.ChatID)
	b.PutInt53(m.MessageThreadID)
	if err := m.Message.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageLinkInfo#2b96fb50: field message: %w", err)
	}
	b.PutInt32(m.MediaTimestamp)
	b.PutBool(m.ForAlbum)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageLinkInfo) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageLinkInfo#2b96fb50 to nil")
	}
	if err := b.ConsumeID(MessageLinkInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode messageLinkInfo#2b96fb50: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageLinkInfo) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageLinkInfo#2b96fb50 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messageLinkInfo#2b96fb50: field is_public: %w", err)
		}
		m.IsPublic = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode messageLinkInfo#2b96fb50: field chat_id: %w", err)
		}
		m.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode messageLinkInfo#2b96fb50: field message_thread_id: %w", err)
		}
		m.MessageThreadID = value
	}
	{
		if err := m.Message.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageLinkInfo#2b96fb50: field message: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode messageLinkInfo#2b96fb50: field media_timestamp: %w", err)
		}
		m.MediaTimestamp = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messageLinkInfo#2b96fb50: field for_album: %w", err)
		}
		m.ForAlbum = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageLinkInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageLinkInfo#2b96fb50 as nil")
	}
	b.ObjStart()
	b.PutID("messageLinkInfo")
	b.Comma()
	b.FieldStart("is_public")
	b.PutBool(m.IsPublic)
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(m.ChatID)
	b.Comma()
	b.FieldStart("message_thread_id")
	b.PutInt53(m.MessageThreadID)
	b.Comma()
	b.FieldStart("message")
	if err := m.Message.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode messageLinkInfo#2b96fb50: field message: %w", err)
	}
	b.Comma()
	b.FieldStart("media_timestamp")
	b.PutInt32(m.MediaTimestamp)
	b.Comma()
	b.FieldStart("for_album")
	b.PutBool(m.ForAlbum)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageLinkInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageLinkInfo#2b96fb50 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageLinkInfo"); err != nil {
				return fmt.Errorf("unable to decode messageLinkInfo#2b96fb50: %w", err)
			}
		case "is_public":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode messageLinkInfo#2b96fb50: field is_public: %w", err)
			}
			m.IsPublic = value
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode messageLinkInfo#2b96fb50: field chat_id: %w", err)
			}
			m.ChatID = value
		case "message_thread_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode messageLinkInfo#2b96fb50: field message_thread_id: %w", err)
			}
			m.MessageThreadID = value
		case "message":
			if err := m.Message.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode messageLinkInfo#2b96fb50: field message: %w", err)
			}
		case "media_timestamp":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode messageLinkInfo#2b96fb50: field media_timestamp: %w", err)
			}
			m.MediaTimestamp = value
		case "for_album":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode messageLinkInfo#2b96fb50: field for_album: %w", err)
			}
			m.ForAlbum = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetIsPublic returns value of IsPublic field.
func (m *MessageLinkInfo) GetIsPublic() (value bool) {
	if m == nil {
		return
	}
	return m.IsPublic
}

// GetChatID returns value of ChatID field.
func (m *MessageLinkInfo) GetChatID() (value int64) {
	if m == nil {
		return
	}
	return m.ChatID
}

// GetMessageThreadID returns value of MessageThreadID field.
func (m *MessageLinkInfo) GetMessageThreadID() (value int64) {
	if m == nil {
		return
	}
	return m.MessageThreadID
}

// GetMessage returns value of Message field.
func (m *MessageLinkInfo) GetMessage() (value Message) {
	if m == nil {
		return
	}
	return m.Message
}

// GetMediaTimestamp returns value of MediaTimestamp field.
func (m *MessageLinkInfo) GetMediaTimestamp() (value int32) {
	if m == nil {
		return
	}
	return m.MediaTimestamp
}

// GetForAlbum returns value of ForAlbum field.
func (m *MessageLinkInfo) GetForAlbum() (value bool) {
	if m == nil {
		return
	}
	return m.ForAlbum
}
