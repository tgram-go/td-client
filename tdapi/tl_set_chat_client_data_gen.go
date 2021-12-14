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

// SetChatClientDataRequest represents TL type `setChatClientData#ceb3273d`.
type SetChatClientDataRequest struct {
	// Chat identifier
	ChatID int64
	// New value of client_data
	ClientData string
}

// SetChatClientDataRequestTypeID is TL type id of SetChatClientDataRequest.
const SetChatClientDataRequestTypeID = 0xceb3273d

// Ensuring interfaces in compile-time for SetChatClientDataRequest.
var (
	_ bin.Encoder     = &SetChatClientDataRequest{}
	_ bin.Decoder     = &SetChatClientDataRequest{}
	_ bin.BareEncoder = &SetChatClientDataRequest{}
	_ bin.BareDecoder = &SetChatClientDataRequest{}
)

func (s *SetChatClientDataRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.ClientData == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetChatClientDataRequest) String() string {
	if s == nil {
		return "SetChatClientDataRequest(nil)"
	}
	type Alias SetChatClientDataRequest
	return fmt.Sprintf("SetChatClientDataRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetChatClientDataRequest) TypeID() uint32 {
	return SetChatClientDataRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetChatClientDataRequest) TypeName() string {
	return "setChatClientData"
}

// TypeInfo returns info about TL type.
func (s *SetChatClientDataRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setChatClientData",
		ID:   SetChatClientDataRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "ClientData",
			SchemaName: "client_data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetChatClientDataRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatClientData#ceb3273d as nil")
	}
	b.PutID(SetChatClientDataRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetChatClientDataRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatClientData#ceb3273d as nil")
	}
	b.PutInt53(s.ChatID)
	b.PutString(s.ClientData)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetChatClientDataRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatClientData#ceb3273d to nil")
	}
	if err := b.ConsumeID(SetChatClientDataRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setChatClientData#ceb3273d: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetChatClientDataRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatClientData#ceb3273d to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setChatClientData#ceb3273d: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setChatClientData#ceb3273d: field client_data: %w", err)
		}
		s.ClientData = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetChatClientDataRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatClientData#ceb3273d as nil")
	}
	b.ObjStart()
	b.PutID("setChatClientData")
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.FieldStart("client_data")
	b.PutString(s.ClientData)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetChatClientDataRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatClientData#ceb3273d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setChatClientData"); err != nil {
				return fmt.Errorf("unable to decode setChatClientData#ceb3273d: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setChatClientData#ceb3273d: field chat_id: %w", err)
			}
			s.ChatID = value
		case "client_data":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setChatClientData#ceb3273d: field client_data: %w", err)
			}
			s.ClientData = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SetChatClientDataRequest) GetChatID() (value int64) {
	return s.ChatID
}

// GetClientData returns value of ClientData field.
func (s *SetChatClientDataRequest) GetClientData() (value string) {
	return s.ClientData
}

// SetChatClientData invokes method setChatClientData#ceb3273d returning error if any.
func (c *Client) SetChatClientData(ctx context.Context, request *SetChatClientDataRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}