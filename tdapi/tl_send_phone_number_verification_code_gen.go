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

// SendPhoneNumberVerificationCodeRequest represents TL type `sendPhoneNumberVerificationCode#7c140dcb`.
type SendPhoneNumberVerificationCodeRequest struct {
	// The phone number of the user, in international format
	PhoneNumber string
	// Settings for the authentication of the user's phone number; pass null to use default
	// settings
	Settings PhoneNumberAuthenticationSettings
}

// SendPhoneNumberVerificationCodeRequestTypeID is TL type id of SendPhoneNumberVerificationCodeRequest.
const SendPhoneNumberVerificationCodeRequestTypeID = 0x7c140dcb

// Ensuring interfaces in compile-time for SendPhoneNumberVerificationCodeRequest.
var (
	_ bin.Encoder     = &SendPhoneNumberVerificationCodeRequest{}
	_ bin.Decoder     = &SendPhoneNumberVerificationCodeRequest{}
	_ bin.BareEncoder = &SendPhoneNumberVerificationCodeRequest{}
	_ bin.BareDecoder = &SendPhoneNumberVerificationCodeRequest{}
)

func (s *SendPhoneNumberVerificationCodeRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.PhoneNumber == "") {
		return false
	}
	if !(s.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendPhoneNumberVerificationCodeRequest) String() string {
	if s == nil {
		return "SendPhoneNumberVerificationCodeRequest(nil)"
	}
	type Alias SendPhoneNumberVerificationCodeRequest
	return fmt.Sprintf("SendPhoneNumberVerificationCodeRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SendPhoneNumberVerificationCodeRequest) TypeID() uint32 {
	return SendPhoneNumberVerificationCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SendPhoneNumberVerificationCodeRequest) TypeName() string {
	return "sendPhoneNumberVerificationCode"
}

// TypeInfo returns info about TL type.
func (s *SendPhoneNumberVerificationCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sendPhoneNumberVerificationCode",
		ID:   SendPhoneNumberVerificationCodeRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhoneNumber",
			SchemaName: "phone_number",
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SendPhoneNumberVerificationCodeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendPhoneNumberVerificationCode#7c140dcb as nil")
	}
	b.PutID(SendPhoneNumberVerificationCodeRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SendPhoneNumberVerificationCodeRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendPhoneNumberVerificationCode#7c140dcb as nil")
	}
	b.PutString(s.PhoneNumber)
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendPhoneNumberVerificationCode#7c140dcb: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SendPhoneNumberVerificationCodeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendPhoneNumberVerificationCode#7c140dcb to nil")
	}
	if err := b.ConsumeID(SendPhoneNumberVerificationCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode sendPhoneNumberVerificationCode#7c140dcb: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SendPhoneNumberVerificationCodeRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendPhoneNumberVerificationCode#7c140dcb to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sendPhoneNumberVerificationCode#7c140dcb: field phone_number: %w", err)
		}
		s.PhoneNumber = value
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sendPhoneNumberVerificationCode#7c140dcb: field settings: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SendPhoneNumberVerificationCodeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sendPhoneNumberVerificationCode#7c140dcb as nil")
	}
	b.ObjStart()
	b.PutID("sendPhoneNumberVerificationCode")
	b.FieldStart("phone_number")
	b.PutString(s.PhoneNumber)
	b.FieldStart("settings")
	if err := s.Settings.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendPhoneNumberVerificationCode#7c140dcb: field settings: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SendPhoneNumberVerificationCodeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sendPhoneNumberVerificationCode#7c140dcb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sendPhoneNumberVerificationCode"); err != nil {
				return fmt.Errorf("unable to decode sendPhoneNumberVerificationCode#7c140dcb: %w", err)
			}
		case "phone_number":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode sendPhoneNumberVerificationCode#7c140dcb: field phone_number: %w", err)
			}
			s.PhoneNumber = value
		case "settings":
			if err := s.Settings.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode sendPhoneNumberVerificationCode#7c140dcb: field settings: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPhoneNumber returns value of PhoneNumber field.
func (s *SendPhoneNumberVerificationCodeRequest) GetPhoneNumber() (value string) {
	return s.PhoneNumber
}

// GetSettings returns value of Settings field.
func (s *SendPhoneNumberVerificationCodeRequest) GetSettings() (value PhoneNumberAuthenticationSettings) {
	return s.Settings
}

// SendPhoneNumberVerificationCode invokes method sendPhoneNumberVerificationCode#7c140dcb returning error if any.
func (c *Client) SendPhoneNumberVerificationCode(ctx context.Context, request *SendPhoneNumberVerificationCodeRequest) (*AuthenticationCodeInfo, error) {
	var result AuthenticationCodeInfo

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}