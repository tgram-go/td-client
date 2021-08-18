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

// MessagesFoundStickerSetsNotModified represents TL type `messages.foundStickerSetsNotModified#d54b65d`.
// No further results were found
//
// See https://core.telegram.org/constructor/messages.foundStickerSetsNotModified for reference.
type MessagesFoundStickerSetsNotModified struct {
}

// MessagesFoundStickerSetsNotModifiedTypeID is TL type id of MessagesFoundStickerSetsNotModified.
const MessagesFoundStickerSetsNotModifiedTypeID = 0xd54b65d

func (f *MessagesFoundStickerSetsNotModified) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *MessagesFoundStickerSetsNotModified) String() string {
	if f == nil {
		return "MessagesFoundStickerSetsNotModified(nil)"
	}
	type Alias MessagesFoundStickerSetsNotModified
	return fmt.Sprintf("MessagesFoundStickerSetsNotModified%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesFoundStickerSetsNotModified) TypeID() uint32 {
	return MessagesFoundStickerSetsNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesFoundStickerSetsNotModified) TypeName() string {
	return "messages.foundStickerSetsNotModified"
}

// TypeInfo returns info about TL type.
func (f *MessagesFoundStickerSetsNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.foundStickerSetsNotModified",
		ID:   MessagesFoundStickerSetsNotModifiedTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (f *MessagesFoundStickerSetsNotModified) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.foundStickerSetsNotModified#d54b65d as nil")
	}
	b.PutID(MessagesFoundStickerSetsNotModifiedTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *MessagesFoundStickerSetsNotModified) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.foundStickerSetsNotModified#d54b65d as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *MessagesFoundStickerSetsNotModified) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.foundStickerSetsNotModified#d54b65d to nil")
	}
	if err := b.ConsumeID(MessagesFoundStickerSetsNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.foundStickerSetsNotModified#d54b65d: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *MessagesFoundStickerSetsNotModified) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.foundStickerSetsNotModified#d54b65d to nil")
	}
	return nil
}

// construct implements constructor of MessagesFoundStickerSetsClass.
func (f MessagesFoundStickerSetsNotModified) construct() MessagesFoundStickerSetsClass { return &f }

// Ensuring interfaces in compile-time for MessagesFoundStickerSetsNotModified.
var (
	_ bin.Encoder     = &MessagesFoundStickerSetsNotModified{}
	_ bin.Decoder     = &MessagesFoundStickerSetsNotModified{}
	_ bin.BareEncoder = &MessagesFoundStickerSetsNotModified{}
	_ bin.BareDecoder = &MessagesFoundStickerSetsNotModified{}

	_ MessagesFoundStickerSetsClass = &MessagesFoundStickerSetsNotModified{}
)

// MessagesFoundStickerSets represents TL type `messages.foundStickerSets#5108d648`.
// Found stickersets
//
// See https://core.telegram.org/constructor/messages.foundStickerSets for reference.
type MessagesFoundStickerSets struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
	// Found stickersets
	Sets []StickerSetCoveredClass
}

// MessagesFoundStickerSetsTypeID is TL type id of MessagesFoundStickerSets.
const MessagesFoundStickerSetsTypeID = 0x5108d648

func (f *MessagesFoundStickerSets) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Hash == 0) {
		return false
	}
	if !(f.Sets == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *MessagesFoundStickerSets) String() string {
	if f == nil {
		return "MessagesFoundStickerSets(nil)"
	}
	type Alias MessagesFoundStickerSets
	return fmt.Sprintf("MessagesFoundStickerSets%+v", Alias(*f))
}

// FillFrom fills MessagesFoundStickerSets from given interface.
func (f *MessagesFoundStickerSets) FillFrom(from interface {
	GetHash() (value int)
	GetSets() (value []StickerSetCoveredClass)
}) {
	f.Hash = from.GetHash()
	f.Sets = from.GetSets()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesFoundStickerSets) TypeID() uint32 {
	return MessagesFoundStickerSetsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesFoundStickerSets) TypeName() string {
	return "messages.foundStickerSets"
}

// TypeInfo returns info about TL type.
func (f *MessagesFoundStickerSets) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.foundStickerSets",
		ID:   MessagesFoundStickerSetsTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Sets",
			SchemaName: "sets",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *MessagesFoundStickerSets) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.foundStickerSets#5108d648 as nil")
	}
	b.PutID(MessagesFoundStickerSetsTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *MessagesFoundStickerSets) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.foundStickerSets#5108d648 as nil")
	}
	b.PutInt(f.Hash)
	b.PutVectorHeader(len(f.Sets))
	for idx, v := range f.Sets {
		if v == nil {
			return fmt.Errorf("unable to encode messages.foundStickerSets#5108d648: field sets element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.foundStickerSets#5108d648: field sets element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (f *MessagesFoundStickerSets) GetHash() (value int) {
	return f.Hash
}

// GetSets returns value of Sets field.
func (f *MessagesFoundStickerSets) GetSets() (value []StickerSetCoveredClass) {
	return f.Sets
}

// MapSets returns field Sets wrapped in StickerSetCoveredClassArray helper.
func (f *MessagesFoundStickerSets) MapSets() (value StickerSetCoveredClassArray) {
	return StickerSetCoveredClassArray(f.Sets)
}

// Decode implements bin.Decoder.
func (f *MessagesFoundStickerSets) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.foundStickerSets#5108d648 to nil")
	}
	if err := b.ConsumeID(MessagesFoundStickerSetsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.foundStickerSets#5108d648: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *MessagesFoundStickerSets) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.foundStickerSets#5108d648 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.foundStickerSets#5108d648: field hash: %w", err)
		}
		f.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.foundStickerSets#5108d648: field sets: %w", err)
		}

		if headerLen != 0 {
			f.Sets = make([]StickerSetCoveredClass, 0, headerLen)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeStickerSetCovered(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.foundStickerSets#5108d648: field sets: %w", err)
			}
			f.Sets = append(f.Sets, value)
		}
	}
	return nil
}

// construct implements constructor of MessagesFoundStickerSetsClass.
func (f MessagesFoundStickerSets) construct() MessagesFoundStickerSetsClass { return &f }

// Ensuring interfaces in compile-time for MessagesFoundStickerSets.
var (
	_ bin.Encoder     = &MessagesFoundStickerSets{}
	_ bin.Decoder     = &MessagesFoundStickerSets{}
	_ bin.BareEncoder = &MessagesFoundStickerSets{}
	_ bin.BareDecoder = &MessagesFoundStickerSets{}

	_ MessagesFoundStickerSetsClass = &MessagesFoundStickerSets{}
)

// MessagesFoundStickerSetsClass represents messages.FoundStickerSets generic type.
//
// See https://core.telegram.org/type/messages.FoundStickerSets for reference.
//
// Example:
//  g, err := tg.DecodeMessagesFoundStickerSets(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.MessagesFoundStickerSetsNotModified: // messages.foundStickerSetsNotModified#d54b65d
//  case *tg.MessagesFoundStickerSets: // messages.foundStickerSets#5108d648
//  default: panic(v)
//  }
type MessagesFoundStickerSetsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesFoundStickerSetsClass

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

	// AsModified tries to map MessagesFoundStickerSetsClass to MessagesFoundStickerSets.
	AsModified() (*MessagesFoundStickerSets, bool)
}

// AsModified tries to map MessagesFoundStickerSetsNotModified to MessagesFoundStickerSets.
func (f *MessagesFoundStickerSetsNotModified) AsModified() (*MessagesFoundStickerSets, bool) {
	return nil, false
}

// AsModified tries to map MessagesFoundStickerSets to MessagesFoundStickerSets.
func (f *MessagesFoundStickerSets) AsModified() (*MessagesFoundStickerSets, bool) {
	return f, true
}

// DecodeMessagesFoundStickerSets implements binary de-serialization for MessagesFoundStickerSetsClass.
func DecodeMessagesFoundStickerSets(buf *bin.Buffer) (MessagesFoundStickerSetsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesFoundStickerSetsNotModifiedTypeID:
		// Decoding messages.foundStickerSetsNotModified#d54b65d.
		v := MessagesFoundStickerSetsNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFoundStickerSetsClass: %w", err)
		}
		return &v, nil
	case MessagesFoundStickerSetsTypeID:
		// Decoding messages.foundStickerSets#5108d648.
		v := MessagesFoundStickerSets{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFoundStickerSetsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesFoundStickerSetsClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesFoundStickerSets boxes the MessagesFoundStickerSetsClass providing a helper.
type MessagesFoundStickerSetsBox struct {
	FoundStickerSets MessagesFoundStickerSetsClass
}

// Decode implements bin.Decoder for MessagesFoundStickerSetsBox.
func (b *MessagesFoundStickerSetsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesFoundStickerSetsBox to nil")
	}
	v, err := DecodeMessagesFoundStickerSets(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.FoundStickerSets = v
	return nil
}

// Encode implements bin.Encode for MessagesFoundStickerSetsBox.
func (b *MessagesFoundStickerSetsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.FoundStickerSets == nil {
		return fmt.Errorf("unable to encode MessagesFoundStickerSetsClass as nil")
	}
	return b.FoundStickerSets.Encode(buf)
}

// MessagesFoundStickerSetsClassArray is adapter for slice of MessagesFoundStickerSetsClass.
type MessagesFoundStickerSetsClassArray []MessagesFoundStickerSetsClass

// Sort sorts slice of MessagesFoundStickerSetsClass.
func (s MessagesFoundStickerSetsClassArray) Sort(less func(a, b MessagesFoundStickerSetsClass) bool) MessagesFoundStickerSetsClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesFoundStickerSetsClass.
func (s MessagesFoundStickerSetsClassArray) SortStable(less func(a, b MessagesFoundStickerSetsClass) bool) MessagesFoundStickerSetsClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesFoundStickerSetsClass.
func (s MessagesFoundStickerSetsClassArray) Retain(keep func(x MessagesFoundStickerSetsClass) bool) MessagesFoundStickerSetsClassArray {
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
func (s MessagesFoundStickerSetsClassArray) First() (v MessagesFoundStickerSetsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesFoundStickerSetsClassArray) Last() (v MessagesFoundStickerSetsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesFoundStickerSetsClassArray) PopFirst() (v MessagesFoundStickerSetsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesFoundStickerSetsClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesFoundStickerSetsClassArray) Pop() (v MessagesFoundStickerSetsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsMessagesFoundStickerSets returns copy with only MessagesFoundStickerSets constructors.
func (s MessagesFoundStickerSetsClassArray) AsMessagesFoundStickerSets() (to MessagesFoundStickerSetsArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesFoundStickerSets)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s MessagesFoundStickerSetsClassArray) AppendOnlyModified(to []*MessagesFoundStickerSets) []*MessagesFoundStickerSets {
	for _, elem := range s {
		value, ok := elem.AsModified()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsModified returns copy with only Modified constructors.
func (s MessagesFoundStickerSetsClassArray) AsModified() (to []*MessagesFoundStickerSets) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s MessagesFoundStickerSetsClassArray) FirstAsModified() (v *MessagesFoundStickerSets, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s MessagesFoundStickerSetsClassArray) LastAsModified() (v *MessagesFoundStickerSets, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *MessagesFoundStickerSetsClassArray) PopFirstAsModified() (v *MessagesFoundStickerSets, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *MessagesFoundStickerSetsClassArray) PopAsModified() (v *MessagesFoundStickerSets, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// MessagesFoundStickerSetsArray is adapter for slice of MessagesFoundStickerSets.
type MessagesFoundStickerSetsArray []MessagesFoundStickerSets

// Sort sorts slice of MessagesFoundStickerSets.
func (s MessagesFoundStickerSetsArray) Sort(less func(a, b MessagesFoundStickerSets) bool) MessagesFoundStickerSetsArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesFoundStickerSets.
func (s MessagesFoundStickerSetsArray) SortStable(less func(a, b MessagesFoundStickerSets) bool) MessagesFoundStickerSetsArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesFoundStickerSets.
func (s MessagesFoundStickerSetsArray) Retain(keep func(x MessagesFoundStickerSets) bool) MessagesFoundStickerSetsArray {
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
func (s MessagesFoundStickerSetsArray) First() (v MessagesFoundStickerSets, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesFoundStickerSetsArray) Last() (v MessagesFoundStickerSets, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesFoundStickerSetsArray) PopFirst() (v MessagesFoundStickerSets, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesFoundStickerSets
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesFoundStickerSetsArray) Pop() (v MessagesFoundStickerSets, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
