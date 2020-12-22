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

// HelpUserInfoEmpty represents TL type `help.userInfoEmpty#f3ae2eed`.
// Internal use
//
// See https://core.telegram.org/constructor/help.userInfoEmpty for reference.
type HelpUserInfoEmpty struct {
}

// HelpUserInfoEmptyTypeID is TL type id of HelpUserInfoEmpty.
const HelpUserInfoEmptyTypeID = 0xf3ae2eed

// Encode implements bin.Encoder.
func (u *HelpUserInfoEmpty) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode help.userInfoEmpty#f3ae2eed as nil")
	}
	b.PutID(HelpUserInfoEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (u *HelpUserInfoEmpty) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode help.userInfoEmpty#f3ae2eed to nil")
	}
	if err := b.ConsumeID(HelpUserInfoEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode help.userInfoEmpty#f3ae2eed: %w", err)
	}
	return nil
}

// construct implements constructor of HelpUserInfoClass.
func (u HelpUserInfoEmpty) construct() HelpUserInfoClass { return &u }

// Ensuring interfaces in compile-time for HelpUserInfoEmpty.
var (
	_ bin.Encoder = &HelpUserInfoEmpty{}
	_ bin.Decoder = &HelpUserInfoEmpty{}

	_ HelpUserInfoClass = &HelpUserInfoEmpty{}
)

// HelpUserInfo represents TL type `help.userInfo#1eb3758`.
// Internal use
//
// See https://core.telegram.org/constructor/help.userInfo for reference.
type HelpUserInfo struct {
	// Info
	Message string
	// Message entities for styled text¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	Entities []MessageEntityClass
	// Author
	Author string
	// Date
	Date int
}

// HelpUserInfoTypeID is TL type id of HelpUserInfo.
const HelpUserInfoTypeID = 0x1eb3758

// Encode implements bin.Encoder.
func (u *HelpUserInfo) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode help.userInfo#1eb3758 as nil")
	}
	b.PutID(HelpUserInfoTypeID)
	b.PutString(u.Message)
	b.PutVectorHeader(len(u.Entities))
	for idx, v := range u.Entities {
		if v == nil {
			return fmt.Errorf("unable to encode help.userInfo#1eb3758: field entities element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.userInfo#1eb3758: field entities element with index %d: %w", idx, err)
		}
	}
	b.PutString(u.Author)
	b.PutInt(u.Date)
	return nil
}

// Decode implements bin.Decoder.
func (u *HelpUserInfo) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode help.userInfo#1eb3758 to nil")
	}
	if err := b.ConsumeID(HelpUserInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode help.userInfo#1eb3758: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.userInfo#1eb3758: field message: %w", err)
		}
		u.Message = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.userInfo#1eb3758: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.userInfo#1eb3758: field entities: %w", err)
			}
			u.Entities = append(u.Entities, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.userInfo#1eb3758: field author: %w", err)
		}
		u.Author = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.userInfo#1eb3758: field date: %w", err)
		}
		u.Date = value
	}
	return nil
}

// construct implements constructor of HelpUserInfoClass.
func (u HelpUserInfo) construct() HelpUserInfoClass { return &u }

// Ensuring interfaces in compile-time for HelpUserInfo.
var (
	_ bin.Encoder = &HelpUserInfo{}
	_ bin.Decoder = &HelpUserInfo{}

	_ HelpUserInfoClass = &HelpUserInfo{}
)

// HelpUserInfoClass represents help.UserInfo generic type.
//
// See https://core.telegram.org/type/help.UserInfo for reference.
//
// Example:
//  g, err := DecodeHelpUserInfo(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *HelpUserInfoEmpty: // help.userInfoEmpty#f3ae2eed
//  case *HelpUserInfo: // help.userInfo#1eb3758
//  default: panic(v)
//  }
type HelpUserInfoClass interface {
	bin.Encoder
	bin.Decoder
	construct() HelpUserInfoClass
}

// DecodeHelpUserInfo implements binary de-serialization for HelpUserInfoClass.
func DecodeHelpUserInfo(buf *bin.Buffer) (HelpUserInfoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case HelpUserInfoEmptyTypeID:
		// Decoding help.userInfoEmpty#f3ae2eed.
		v := HelpUserInfoEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpUserInfoClass: %w", err)
		}
		return &v, nil
	case HelpUserInfoTypeID:
		// Decoding help.userInfo#1eb3758.
		v := HelpUserInfo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpUserInfoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode HelpUserInfoClass: %w", bin.NewUnexpectedID(id))
	}
}

// HelpUserInfo boxes the HelpUserInfoClass providing a helper.
type HelpUserInfoBox struct {
	UserInfo HelpUserInfoClass
}

// Decode implements bin.Decoder for HelpUserInfoBox.
func (b *HelpUserInfoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode HelpUserInfoBox to nil")
	}
	v, err := DecodeHelpUserInfo(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.UserInfo = v
	return nil
}

// Encode implements bin.Encode for HelpUserInfoBox.
func (b *HelpUserInfoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.UserInfo == nil {
		return fmt.Errorf("unable to encode HelpUserInfoClass as nil")
	}
	return b.UserInfo.Encode(buf)
}
