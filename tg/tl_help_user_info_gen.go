// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
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
)

// HelpUserInfoEmpty represents TL type `help.userInfoEmpty#f3ae2eed`.
// Internal use
//
// See https://core.telegram.org/constructor/help.userInfoEmpty for reference.
type HelpUserInfoEmpty struct {
}

// HelpUserInfoEmptyTypeID is TL type id of HelpUserInfoEmpty.
const HelpUserInfoEmptyTypeID = 0xf3ae2eed

func (u *HelpUserInfoEmpty) Zero() bool {
	if u == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (u *HelpUserInfoEmpty) String() string {
	if u == nil {
		return "HelpUserInfoEmpty(nil)"
	}
	type Alias HelpUserInfoEmpty
	return fmt.Sprintf("HelpUserInfoEmpty%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpUserInfoEmpty) TypeID() uint32 {
	return HelpUserInfoEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpUserInfoEmpty) TypeName() string {
	return "help.userInfoEmpty"
}

// TypeInfo returns info about TL type.
func (u *HelpUserInfoEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.userInfoEmpty",
		ID:   HelpUserInfoEmptyTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (u *HelpUserInfoEmpty) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode help.userInfoEmpty#f3ae2eed as nil")
	}
	b.PutID(HelpUserInfoEmptyTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *HelpUserInfoEmpty) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode help.userInfoEmpty#f3ae2eed as nil")
	}
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
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *HelpUserInfoEmpty) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode help.userInfoEmpty#f3ae2eed to nil")
	}
	return nil
}

// construct implements constructor of HelpUserInfoClass.
func (u HelpUserInfoEmpty) construct() HelpUserInfoClass { return &u }

// Ensuring interfaces in compile-time for HelpUserInfoEmpty.
var (
	_ bin.Encoder     = &HelpUserInfoEmpty{}
	_ bin.Decoder     = &HelpUserInfoEmpty{}
	_ bin.BareEncoder = &HelpUserInfoEmpty{}
	_ bin.BareDecoder = &HelpUserInfoEmpty{}

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

func (u *HelpUserInfo) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Message == "") {
		return false
	}
	if !(u.Entities == nil) {
		return false
	}
	if !(u.Author == "") {
		return false
	}
	if !(u.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *HelpUserInfo) String() string {
	if u == nil {
		return "HelpUserInfo(nil)"
	}
	type Alias HelpUserInfo
	return fmt.Sprintf("HelpUserInfo%+v", Alias(*u))
}

// FillFrom fills HelpUserInfo from given interface.
func (u *HelpUserInfo) FillFrom(from interface {
	GetMessage() (value string)
	GetEntities() (value []MessageEntityClass)
	GetAuthor() (value string)
	GetDate() (value int)
}) {
	u.Message = from.GetMessage()
	u.Entities = from.GetEntities()
	u.Author = from.GetAuthor()
	u.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpUserInfo) TypeID() uint32 {
	return HelpUserInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpUserInfo) TypeName() string {
	return "help.userInfo"
}

// TypeInfo returns info about TL type.
func (u *HelpUserInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.userInfo",
		ID:   HelpUserInfoTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
		},
		{
			Name:       "Author",
			SchemaName: "author",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *HelpUserInfo) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode help.userInfo#1eb3758 as nil")
	}
	b.PutID(HelpUserInfoTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *HelpUserInfo) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode help.userInfo#1eb3758 as nil")
	}
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

// GetMessage returns value of Message field.
func (u *HelpUserInfo) GetMessage() (value string) {
	return u.Message
}

// GetEntities returns value of Entities field.
func (u *HelpUserInfo) GetEntities() (value []MessageEntityClass) {
	return u.Entities
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (u *HelpUserInfo) MapEntities() (value MessageEntityClassArray) {
	return MessageEntityClassArray(u.Entities)
}

// GetAuthor returns value of Author field.
func (u *HelpUserInfo) GetAuthor() (value string) {
	return u.Author
}

// GetDate returns value of Date field.
func (u *HelpUserInfo) GetDate() (value int) {
	return u.Date
}

// Decode implements bin.Decoder.
func (u *HelpUserInfo) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode help.userInfo#1eb3758 to nil")
	}
	if err := b.ConsumeID(HelpUserInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode help.userInfo#1eb3758: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *HelpUserInfo) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode help.userInfo#1eb3758 to nil")
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

		if headerLen != 0 {
			u.Entities = make([]MessageEntityClass, 0, headerLen)
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
	_ bin.Encoder     = &HelpUserInfo{}
	_ bin.Decoder     = &HelpUserInfo{}
	_ bin.BareEncoder = &HelpUserInfo{}
	_ bin.BareDecoder = &HelpUserInfo{}

	_ HelpUserInfoClass = &HelpUserInfo{}
)

// HelpUserInfoClass represents help.UserInfo generic type.
//
// See https://core.telegram.org/type/help.UserInfo for reference.
//
// Example:
//  g, err := tg.DecodeHelpUserInfo(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.HelpUserInfoEmpty: // help.userInfoEmpty#f3ae2eed
//  case *tg.HelpUserInfo: // help.userInfo#1eb3758
//  default: panic(v)
//  }
type HelpUserInfoClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() HelpUserInfoClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// AsNotEmpty tries to map HelpUserInfoClass to HelpUserInfo.
	AsNotEmpty() (*HelpUserInfo, bool)
}

// AsNotEmpty tries to map HelpUserInfoEmpty to HelpUserInfo.
func (u *HelpUserInfoEmpty) AsNotEmpty() (*HelpUserInfo, bool) {
	return nil, false
}

// AsNotEmpty tries to map HelpUserInfo to HelpUserInfo.
func (u *HelpUserInfo) AsNotEmpty() (*HelpUserInfo, bool) {
	return u, true
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

// HelpUserInfoClassArray is adapter for slice of HelpUserInfoClass.
type HelpUserInfoClassArray []HelpUserInfoClass

// Sort sorts slice of HelpUserInfoClass.
func (s HelpUserInfoClassArray) Sort(less func(a, b HelpUserInfoClass) bool) HelpUserInfoClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpUserInfoClass.
func (s HelpUserInfoClassArray) SortStable(less func(a, b HelpUserInfoClass) bool) HelpUserInfoClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpUserInfoClass.
func (s HelpUserInfoClassArray) Retain(keep func(x HelpUserInfoClass) bool) HelpUserInfoClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s HelpUserInfoClassArray) First() (v HelpUserInfoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpUserInfoClassArray) Last() (v HelpUserInfoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpUserInfoClassArray) PopFirst() (v HelpUserInfoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpUserInfoClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpUserInfoClassArray) Pop() (v HelpUserInfoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsHelpUserInfo returns copy with only HelpUserInfo constructors.
func (s HelpUserInfoClassArray) AsHelpUserInfo() (to HelpUserInfoArray) {
	for _, elem := range s {
		value, ok := elem.(*HelpUserInfo)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s HelpUserInfoClassArray) AppendOnlyNotEmpty(to []*HelpUserInfo) []*HelpUserInfo {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s HelpUserInfoClassArray) AsNotEmpty() (to []*HelpUserInfo) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s HelpUserInfoClassArray) FirstAsNotEmpty() (v *HelpUserInfo, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s HelpUserInfoClassArray) LastAsNotEmpty() (v *HelpUserInfo, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *HelpUserInfoClassArray) PopFirstAsNotEmpty() (v *HelpUserInfo, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *HelpUserInfoClassArray) PopAsNotEmpty() (v *HelpUserInfo, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// HelpUserInfoArray is adapter for slice of HelpUserInfo.
type HelpUserInfoArray []HelpUserInfo

// Sort sorts slice of HelpUserInfo.
func (s HelpUserInfoArray) Sort(less func(a, b HelpUserInfo) bool) HelpUserInfoArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpUserInfo.
func (s HelpUserInfoArray) SortStable(less func(a, b HelpUserInfo) bool) HelpUserInfoArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpUserInfo.
func (s HelpUserInfoArray) Retain(keep func(x HelpUserInfo) bool) HelpUserInfoArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s HelpUserInfoArray) First() (v HelpUserInfo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpUserInfoArray) Last() (v HelpUserInfo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpUserInfoArray) PopFirst() (v HelpUserInfo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpUserInfo
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpUserInfoArray) Pop() (v HelpUserInfo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByDate sorts slice of HelpUserInfo by Date.
func (s HelpUserInfoArray) SortByDate() HelpUserInfoArray {
	return s.Sort(func(a, b HelpUserInfo) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of HelpUserInfo by Date.
func (s HelpUserInfoArray) SortStableByDate() HelpUserInfoArray {
	return s.SortStable(func(a, b HelpUserInfo) bool {
		return a.GetDate() < b.GetDate()
	})
}
