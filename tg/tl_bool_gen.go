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

// BoolFalse represents TL type `boolFalse#bc799737`.
// Constructor may be interpreted as a booleanfalse value.
//
// See https://core.telegram.org/constructor/boolFalse for reference.
type BoolFalse struct {
}

// BoolFalseTypeID is TL type id of BoolFalse.
const BoolFalseTypeID = 0xbc799737

func (b *BoolFalse) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BoolFalse) String() string {
	if b == nil {
		return "BoolFalse(nil)"
	}
	type Alias BoolFalse
	return fmt.Sprintf("BoolFalse%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BoolFalse) TypeID() uint32 {
	return BoolFalseTypeID
}

// TypeName returns name of type in TL schema.
func (*BoolFalse) TypeName() string {
	return "boolFalse"
}

// TypeInfo returns info about TL type.
func (b *BoolFalse) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "boolFalse",
		ID:   BoolFalseTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BoolFalse) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode boolFalse#bc799737 as nil")
	}
	buf.PutID(BoolFalseTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BoolFalse) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode boolFalse#bc799737 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BoolFalse) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode boolFalse#bc799737 to nil")
	}
	if err := buf.ConsumeID(BoolFalseTypeID); err != nil {
		return fmt.Errorf("unable to decode boolFalse#bc799737: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BoolFalse) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode boolFalse#bc799737 to nil")
	}
	return nil
}

// construct implements constructor of BoolClass.
func (b BoolFalse) construct() BoolClass { return &b }

// Ensuring interfaces in compile-time for BoolFalse.
var (
	_ bin.Encoder     = &BoolFalse{}
	_ bin.Decoder     = &BoolFalse{}
	_ bin.BareEncoder = &BoolFalse{}
	_ bin.BareDecoder = &BoolFalse{}

	_ BoolClass = &BoolFalse{}
)

// BoolTrue represents TL type `boolTrue#997275b5`.
// The constructor can be interpreted as a booleantrue value.
//
// See https://core.telegram.org/constructor/boolTrue for reference.
type BoolTrue struct {
}

// BoolTrueTypeID is TL type id of BoolTrue.
const BoolTrueTypeID = 0x997275b5

func (b *BoolTrue) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BoolTrue) String() string {
	if b == nil {
		return "BoolTrue(nil)"
	}
	type Alias BoolTrue
	return fmt.Sprintf("BoolTrue%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BoolTrue) TypeID() uint32 {
	return BoolTrueTypeID
}

// TypeName returns name of type in TL schema.
func (*BoolTrue) TypeName() string {
	return "boolTrue"
}

// TypeInfo returns info about TL type.
func (b *BoolTrue) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "boolTrue",
		ID:   BoolTrueTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BoolTrue) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode boolTrue#997275b5 as nil")
	}
	buf.PutID(BoolTrueTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BoolTrue) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode boolTrue#997275b5 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BoolTrue) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode boolTrue#997275b5 to nil")
	}
	if err := buf.ConsumeID(BoolTrueTypeID); err != nil {
		return fmt.Errorf("unable to decode boolTrue#997275b5: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BoolTrue) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode boolTrue#997275b5 to nil")
	}
	return nil
}

// construct implements constructor of BoolClass.
func (b BoolTrue) construct() BoolClass { return &b }

// Ensuring interfaces in compile-time for BoolTrue.
var (
	_ bin.Encoder     = &BoolTrue{}
	_ bin.Decoder     = &BoolTrue{}
	_ bin.BareEncoder = &BoolTrue{}
	_ bin.BareDecoder = &BoolTrue{}

	_ BoolClass = &BoolTrue{}
)

// BoolClass represents Bool generic type.
//
// See https://core.telegram.org/type/Bool for reference.
//
// Example:
//  g, err := tg.DecodeBool(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.BoolFalse: // boolFalse#bc799737
//  case *tg.BoolTrue: // boolTrue#997275b5
//  default: panic(v)
//  }
type BoolClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() BoolClass

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
}

// DecodeBool implements binary de-serialization for BoolClass.
func DecodeBool(buf *bin.Buffer) (BoolClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case BoolFalseTypeID:
		// Decoding boolFalse#bc799737.
		v := BoolFalse{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BoolClass: %w", err)
		}
		return &v, nil
	case BoolTrueTypeID:
		// Decoding boolTrue#997275b5.
		v := BoolTrue{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BoolClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BoolClass: %w", bin.NewUnexpectedID(id))
	}
}

// Bool boxes the BoolClass providing a helper.
type BoolBox struct {
	Bool BoolClass
}

// Decode implements bin.Decoder for BoolBox.
func (b *BoolBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode BoolBox to nil")
	}
	v, err := DecodeBool(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Bool = v
	return nil
}

// Encode implements bin.Encode for BoolBox.
func (b *BoolBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Bool == nil {
		return fmt.Errorf("unable to encode BoolClass as nil")
	}
	return b.Bool.Encode(buf)
}

// BoolClassArray is adapter for slice of BoolClass.
type BoolClassArray []BoolClass

// Sort sorts slice of BoolClass.
func (s BoolClassArray) Sort(less func(a, b BoolClass) bool) BoolClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of BoolClass.
func (s BoolClassArray) SortStable(less func(a, b BoolClass) bool) BoolClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of BoolClass.
func (s BoolClassArray) Retain(keep func(x BoolClass) bool) BoolClassArray {
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
func (s BoolClassArray) First() (v BoolClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s BoolClassArray) Last() (v BoolClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *BoolClassArray) PopFirst() (v BoolClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero BoolClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *BoolClassArray) Pop() (v BoolClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
