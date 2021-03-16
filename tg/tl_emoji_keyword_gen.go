// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

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
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// EmojiKeyword represents TL type `emojiKeyword#d5b3b9f9`.
// Emoji keyword
//
// See https://core.telegram.org/constructor/emojiKeyword for reference.
type EmojiKeyword struct {
	// Keyword
	Keyword string
	// Emojis associated to keyword
	Emoticons []string
}

// EmojiKeywordTypeID is TL type id of EmojiKeyword.
const EmojiKeywordTypeID = 0xd5b3b9f9

func (e *EmojiKeyword) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Keyword == "") {
		return false
	}
	if !(e.Emoticons == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EmojiKeyword) String() string {
	if e == nil {
		return "EmojiKeyword(nil)"
	}
	type Alias EmojiKeyword
	return fmt.Sprintf("EmojiKeyword%+v", Alias(*e))
}

// FillFrom fills EmojiKeyword from given interface.
func (e *EmojiKeyword) FillFrom(from interface {
	GetKeyword() (value string)
	GetEmoticons() (value []string)
}) {
	e.Keyword = from.GetKeyword()
	e.Emoticons = from.GetEmoticons()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EmojiKeyword) TypeID() uint32 {
	return EmojiKeywordTypeID
}

// TypeName returns name of type in TL schema.
func (*EmojiKeyword) TypeName() string {
	return "emojiKeyword"
}

// TypeInfo returns info about TL type.
func (e *EmojiKeyword) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "emojiKeyword",
		ID:   EmojiKeywordTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Keyword",
			SchemaName: "keyword",
		},
		{
			Name:       "Emoticons",
			SchemaName: "emoticons",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EmojiKeyword) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiKeyword#d5b3b9f9 as nil")
	}
	b.PutID(EmojiKeywordTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EmojiKeyword) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiKeyword#d5b3b9f9 as nil")
	}
	b.PutString(e.Keyword)
	b.PutVectorHeader(len(e.Emoticons))
	for _, v := range e.Emoticons {
		b.PutString(v)
	}
	return nil
}

// GetKeyword returns value of Keyword field.
func (e *EmojiKeyword) GetKeyword() (value string) {
	return e.Keyword
}

// GetEmoticons returns value of Emoticons field.
func (e *EmojiKeyword) GetEmoticons() (value []string) {
	return e.Emoticons
}

// Decode implements bin.Decoder.
func (e *EmojiKeyword) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiKeyword#d5b3b9f9 to nil")
	}
	if err := b.ConsumeID(EmojiKeywordTypeID); err != nil {
		return fmt.Errorf("unable to decode emojiKeyword#d5b3b9f9: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EmojiKeyword) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiKeyword#d5b3b9f9 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode emojiKeyword#d5b3b9f9: field keyword: %w", err)
		}
		e.Keyword = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode emojiKeyword#d5b3b9f9: field emoticons: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode emojiKeyword#d5b3b9f9: field emoticons: %w", err)
			}
			e.Emoticons = append(e.Emoticons, value)
		}
	}
	return nil
}

// construct implements constructor of EmojiKeywordClass.
func (e EmojiKeyword) construct() EmojiKeywordClass { return &e }

// Ensuring interfaces in compile-time for EmojiKeyword.
var (
	_ bin.Encoder     = &EmojiKeyword{}
	_ bin.Decoder     = &EmojiKeyword{}
	_ bin.BareEncoder = &EmojiKeyword{}
	_ bin.BareDecoder = &EmojiKeyword{}

	_ EmojiKeywordClass = &EmojiKeyword{}
)

// EmojiKeywordDeleted represents TL type `emojiKeywordDeleted#236df622`.
// Deleted emoji keyword
//
// See https://core.telegram.org/constructor/emojiKeywordDeleted for reference.
type EmojiKeywordDeleted struct {
	// Keyword
	Keyword string
	// Emojis that were associated to keyword
	Emoticons []string
}

// EmojiKeywordDeletedTypeID is TL type id of EmojiKeywordDeleted.
const EmojiKeywordDeletedTypeID = 0x236df622

func (e *EmojiKeywordDeleted) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Keyword == "") {
		return false
	}
	if !(e.Emoticons == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EmojiKeywordDeleted) String() string {
	if e == nil {
		return "EmojiKeywordDeleted(nil)"
	}
	type Alias EmojiKeywordDeleted
	return fmt.Sprintf("EmojiKeywordDeleted%+v", Alias(*e))
}

// FillFrom fills EmojiKeywordDeleted from given interface.
func (e *EmojiKeywordDeleted) FillFrom(from interface {
	GetKeyword() (value string)
	GetEmoticons() (value []string)
}) {
	e.Keyword = from.GetKeyword()
	e.Emoticons = from.GetEmoticons()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EmojiKeywordDeleted) TypeID() uint32 {
	return EmojiKeywordDeletedTypeID
}

// TypeName returns name of type in TL schema.
func (*EmojiKeywordDeleted) TypeName() string {
	return "emojiKeywordDeleted"
}

// TypeInfo returns info about TL type.
func (e *EmojiKeywordDeleted) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "emojiKeywordDeleted",
		ID:   EmojiKeywordDeletedTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Keyword",
			SchemaName: "keyword",
		},
		{
			Name:       "Emoticons",
			SchemaName: "emoticons",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EmojiKeywordDeleted) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiKeywordDeleted#236df622 as nil")
	}
	b.PutID(EmojiKeywordDeletedTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EmojiKeywordDeleted) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiKeywordDeleted#236df622 as nil")
	}
	b.PutString(e.Keyword)
	b.PutVectorHeader(len(e.Emoticons))
	for _, v := range e.Emoticons {
		b.PutString(v)
	}
	return nil
}

// GetKeyword returns value of Keyword field.
func (e *EmojiKeywordDeleted) GetKeyword() (value string) {
	return e.Keyword
}

// GetEmoticons returns value of Emoticons field.
func (e *EmojiKeywordDeleted) GetEmoticons() (value []string) {
	return e.Emoticons
}

// Decode implements bin.Decoder.
func (e *EmojiKeywordDeleted) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiKeywordDeleted#236df622 to nil")
	}
	if err := b.ConsumeID(EmojiKeywordDeletedTypeID); err != nil {
		return fmt.Errorf("unable to decode emojiKeywordDeleted#236df622: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EmojiKeywordDeleted) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiKeywordDeleted#236df622 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode emojiKeywordDeleted#236df622: field keyword: %w", err)
		}
		e.Keyword = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode emojiKeywordDeleted#236df622: field emoticons: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode emojiKeywordDeleted#236df622: field emoticons: %w", err)
			}
			e.Emoticons = append(e.Emoticons, value)
		}
	}
	return nil
}

// construct implements constructor of EmojiKeywordClass.
func (e EmojiKeywordDeleted) construct() EmojiKeywordClass { return &e }

// Ensuring interfaces in compile-time for EmojiKeywordDeleted.
var (
	_ bin.Encoder     = &EmojiKeywordDeleted{}
	_ bin.Decoder     = &EmojiKeywordDeleted{}
	_ bin.BareEncoder = &EmojiKeywordDeleted{}
	_ bin.BareDecoder = &EmojiKeywordDeleted{}

	_ EmojiKeywordClass = &EmojiKeywordDeleted{}
)

// EmojiKeywordClass represents EmojiKeyword generic type.
//
// See https://core.telegram.org/type/EmojiKeyword for reference.
//
// Example:
//  g, err := tg.DecodeEmojiKeyword(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.EmojiKeyword: // emojiKeyword#d5b3b9f9
//  case *tg.EmojiKeywordDeleted: // emojiKeywordDeleted#236df622
//  default: panic(v)
//  }
type EmojiKeywordClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() EmojiKeywordClass

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

	// Keyword
	GetKeyword() (value string)
	// Emojis associated to keyword
	GetEmoticons() (value []string)
}

// DecodeEmojiKeyword implements binary de-serialization for EmojiKeywordClass.
func DecodeEmojiKeyword(buf *bin.Buffer) (EmojiKeywordClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case EmojiKeywordTypeID:
		// Decoding emojiKeyword#d5b3b9f9.
		v := EmojiKeyword{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EmojiKeywordClass: %w", err)
		}
		return &v, nil
	case EmojiKeywordDeletedTypeID:
		// Decoding emojiKeywordDeleted#236df622.
		v := EmojiKeywordDeleted{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EmojiKeywordClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode EmojiKeywordClass: %w", bin.NewUnexpectedID(id))
	}
}

// EmojiKeyword boxes the EmojiKeywordClass providing a helper.
type EmojiKeywordBox struct {
	EmojiKeyword EmojiKeywordClass
}

// Decode implements bin.Decoder for EmojiKeywordBox.
func (b *EmojiKeywordBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode EmojiKeywordBox to nil")
	}
	v, err := DecodeEmojiKeyword(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.EmojiKeyword = v
	return nil
}

// Encode implements bin.Encode for EmojiKeywordBox.
func (b *EmojiKeywordBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.EmojiKeyword == nil {
		return fmt.Errorf("unable to encode EmojiKeywordClass as nil")
	}
	return b.EmojiKeyword.Encode(buf)
}

// EmojiKeywordClassArray is adapter for slice of EmojiKeywordClass.
type EmojiKeywordClassArray []EmojiKeywordClass

// Sort sorts slice of EmojiKeywordClass.
func (s EmojiKeywordClassArray) Sort(less func(a, b EmojiKeywordClass) bool) EmojiKeywordClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of EmojiKeywordClass.
func (s EmojiKeywordClassArray) SortStable(less func(a, b EmojiKeywordClass) bool) EmojiKeywordClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of EmojiKeywordClass.
func (s EmojiKeywordClassArray) Retain(keep func(x EmojiKeywordClass) bool) EmojiKeywordClassArray {
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
func (s EmojiKeywordClassArray) First() (v EmojiKeywordClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s EmojiKeywordClassArray) Last() (v EmojiKeywordClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *EmojiKeywordClassArray) PopFirst() (v EmojiKeywordClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero EmojiKeywordClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *EmojiKeywordClassArray) Pop() (v EmojiKeywordClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsEmojiKeyword returns copy with only EmojiKeyword constructors.
func (s EmojiKeywordClassArray) AsEmojiKeyword() (to EmojiKeywordArray) {
	for _, elem := range s {
		value, ok := elem.(*EmojiKeyword)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsEmojiKeywordDeleted returns copy with only EmojiKeywordDeleted constructors.
func (s EmojiKeywordClassArray) AsEmojiKeywordDeleted() (to EmojiKeywordDeletedArray) {
	for _, elem := range s {
		value, ok := elem.(*EmojiKeywordDeleted)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// EmojiKeywordArray is adapter for slice of EmojiKeyword.
type EmojiKeywordArray []EmojiKeyword

// Sort sorts slice of EmojiKeyword.
func (s EmojiKeywordArray) Sort(less func(a, b EmojiKeyword) bool) EmojiKeywordArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of EmojiKeyword.
func (s EmojiKeywordArray) SortStable(less func(a, b EmojiKeyword) bool) EmojiKeywordArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of EmojiKeyword.
func (s EmojiKeywordArray) Retain(keep func(x EmojiKeyword) bool) EmojiKeywordArray {
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
func (s EmojiKeywordArray) First() (v EmojiKeyword, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s EmojiKeywordArray) Last() (v EmojiKeyword, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *EmojiKeywordArray) PopFirst() (v EmojiKeyword, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero EmojiKeyword
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *EmojiKeywordArray) Pop() (v EmojiKeyword, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// EmojiKeywordDeletedArray is adapter for slice of EmojiKeywordDeleted.
type EmojiKeywordDeletedArray []EmojiKeywordDeleted

// Sort sorts slice of EmojiKeywordDeleted.
func (s EmojiKeywordDeletedArray) Sort(less func(a, b EmojiKeywordDeleted) bool) EmojiKeywordDeletedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of EmojiKeywordDeleted.
func (s EmojiKeywordDeletedArray) SortStable(less func(a, b EmojiKeywordDeleted) bool) EmojiKeywordDeletedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of EmojiKeywordDeleted.
func (s EmojiKeywordDeletedArray) Retain(keep func(x EmojiKeywordDeleted) bool) EmojiKeywordDeletedArray {
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
func (s EmojiKeywordDeletedArray) First() (v EmojiKeywordDeleted, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s EmojiKeywordDeletedArray) Last() (v EmojiKeywordDeleted, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *EmojiKeywordDeletedArray) PopFirst() (v EmojiKeywordDeleted, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero EmojiKeywordDeleted
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *EmojiKeywordDeletedArray) Pop() (v EmojiKeywordDeleted, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
