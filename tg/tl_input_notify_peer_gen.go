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

// InputNotifyPeer represents TL type `inputNotifyPeer#b8bc5b0c`.
// Notifications generated by a certain user or group.
//
// See https://core.telegram.org/constructor/inputNotifyPeer for reference.
type InputNotifyPeer struct {
	// User or group
	Peer InputPeerClass
}

// InputNotifyPeerTypeID is TL type id of InputNotifyPeer.
const InputNotifyPeerTypeID = 0xb8bc5b0c

func (i *InputNotifyPeer) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputNotifyPeer) String() string {
	if i == nil {
		return "InputNotifyPeer(nil)"
	}
	type Alias InputNotifyPeer
	return fmt.Sprintf("InputNotifyPeer%+v", Alias(*i))
}

// FillFrom fills InputNotifyPeer from given interface.
func (i *InputNotifyPeer) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
}) {
	i.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputNotifyPeer) TypeID() uint32 {
	return InputNotifyPeerTypeID
}

// TypeName returns name of type in TL schema.
func (*InputNotifyPeer) TypeName() string {
	return "inputNotifyPeer"
}

// TypeInfo returns info about TL type.
func (i *InputNotifyPeer) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputNotifyPeer",
		ID:   InputNotifyPeerTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputNotifyPeer) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputNotifyPeer#b8bc5b0c as nil")
	}
	b.PutID(InputNotifyPeerTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputNotifyPeer) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputNotifyPeer#b8bc5b0c as nil")
	}
	if i.Peer == nil {
		return fmt.Errorf("unable to encode inputNotifyPeer#b8bc5b0c: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputNotifyPeer#b8bc5b0c: field peer: %w", err)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (i *InputNotifyPeer) GetPeer() (value InputPeerClass) {
	return i.Peer
}

// Decode implements bin.Decoder.
func (i *InputNotifyPeer) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputNotifyPeer#b8bc5b0c to nil")
	}
	if err := b.ConsumeID(InputNotifyPeerTypeID); err != nil {
		return fmt.Errorf("unable to decode inputNotifyPeer#b8bc5b0c: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputNotifyPeer) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputNotifyPeer#b8bc5b0c to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputNotifyPeer#b8bc5b0c: field peer: %w", err)
		}
		i.Peer = value
	}
	return nil
}

// construct implements constructor of InputNotifyPeerClass.
func (i InputNotifyPeer) construct() InputNotifyPeerClass { return &i }

// Ensuring interfaces in compile-time for InputNotifyPeer.
var (
	_ bin.Encoder     = &InputNotifyPeer{}
	_ bin.Decoder     = &InputNotifyPeer{}
	_ bin.BareEncoder = &InputNotifyPeer{}
	_ bin.BareDecoder = &InputNotifyPeer{}

	_ InputNotifyPeerClass = &InputNotifyPeer{}
)

// InputNotifyUsers represents TL type `inputNotifyUsers#193b4417`.
// Notifications generated by all users.
//
// See https://core.telegram.org/constructor/inputNotifyUsers for reference.
type InputNotifyUsers struct {
}

// InputNotifyUsersTypeID is TL type id of InputNotifyUsers.
const InputNotifyUsersTypeID = 0x193b4417

func (i *InputNotifyUsers) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputNotifyUsers) String() string {
	if i == nil {
		return "InputNotifyUsers(nil)"
	}
	type Alias InputNotifyUsers
	return fmt.Sprintf("InputNotifyUsers%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputNotifyUsers) TypeID() uint32 {
	return InputNotifyUsersTypeID
}

// TypeName returns name of type in TL schema.
func (*InputNotifyUsers) TypeName() string {
	return "inputNotifyUsers"
}

// TypeInfo returns info about TL type.
func (i *InputNotifyUsers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputNotifyUsers",
		ID:   InputNotifyUsersTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputNotifyUsers) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputNotifyUsers#193b4417 as nil")
	}
	b.PutID(InputNotifyUsersTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputNotifyUsers) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputNotifyUsers#193b4417 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputNotifyUsers) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputNotifyUsers#193b4417 to nil")
	}
	if err := b.ConsumeID(InputNotifyUsersTypeID); err != nil {
		return fmt.Errorf("unable to decode inputNotifyUsers#193b4417: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputNotifyUsers) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputNotifyUsers#193b4417 to nil")
	}
	return nil
}

// construct implements constructor of InputNotifyPeerClass.
func (i InputNotifyUsers) construct() InputNotifyPeerClass { return &i }

// Ensuring interfaces in compile-time for InputNotifyUsers.
var (
	_ bin.Encoder     = &InputNotifyUsers{}
	_ bin.Decoder     = &InputNotifyUsers{}
	_ bin.BareEncoder = &InputNotifyUsers{}
	_ bin.BareDecoder = &InputNotifyUsers{}

	_ InputNotifyPeerClass = &InputNotifyUsers{}
)

// InputNotifyChats represents TL type `inputNotifyChats#4a95e84e`.
// Notifications generated by all groups.
//
// See https://core.telegram.org/constructor/inputNotifyChats for reference.
type InputNotifyChats struct {
}

// InputNotifyChatsTypeID is TL type id of InputNotifyChats.
const InputNotifyChatsTypeID = 0x4a95e84e

func (i *InputNotifyChats) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputNotifyChats) String() string {
	if i == nil {
		return "InputNotifyChats(nil)"
	}
	type Alias InputNotifyChats
	return fmt.Sprintf("InputNotifyChats%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputNotifyChats) TypeID() uint32 {
	return InputNotifyChatsTypeID
}

// TypeName returns name of type in TL schema.
func (*InputNotifyChats) TypeName() string {
	return "inputNotifyChats"
}

// TypeInfo returns info about TL type.
func (i *InputNotifyChats) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputNotifyChats",
		ID:   InputNotifyChatsTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputNotifyChats) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputNotifyChats#4a95e84e as nil")
	}
	b.PutID(InputNotifyChatsTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputNotifyChats) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputNotifyChats#4a95e84e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputNotifyChats) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputNotifyChats#4a95e84e to nil")
	}
	if err := b.ConsumeID(InputNotifyChatsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputNotifyChats#4a95e84e: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputNotifyChats) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputNotifyChats#4a95e84e to nil")
	}
	return nil
}

// construct implements constructor of InputNotifyPeerClass.
func (i InputNotifyChats) construct() InputNotifyPeerClass { return &i }

// Ensuring interfaces in compile-time for InputNotifyChats.
var (
	_ bin.Encoder     = &InputNotifyChats{}
	_ bin.Decoder     = &InputNotifyChats{}
	_ bin.BareEncoder = &InputNotifyChats{}
	_ bin.BareDecoder = &InputNotifyChats{}

	_ InputNotifyPeerClass = &InputNotifyChats{}
)

// InputNotifyBroadcasts represents TL type `inputNotifyBroadcasts#b1db7c7e`.
// All channels¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/constructor/inputNotifyBroadcasts for reference.
type InputNotifyBroadcasts struct {
}

// InputNotifyBroadcastsTypeID is TL type id of InputNotifyBroadcasts.
const InputNotifyBroadcastsTypeID = 0xb1db7c7e

func (i *InputNotifyBroadcasts) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputNotifyBroadcasts) String() string {
	if i == nil {
		return "InputNotifyBroadcasts(nil)"
	}
	type Alias InputNotifyBroadcasts
	return fmt.Sprintf("InputNotifyBroadcasts%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputNotifyBroadcasts) TypeID() uint32 {
	return InputNotifyBroadcastsTypeID
}

// TypeName returns name of type in TL schema.
func (*InputNotifyBroadcasts) TypeName() string {
	return "inputNotifyBroadcasts"
}

// TypeInfo returns info about TL type.
func (i *InputNotifyBroadcasts) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputNotifyBroadcasts",
		ID:   InputNotifyBroadcastsTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputNotifyBroadcasts) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputNotifyBroadcasts#b1db7c7e as nil")
	}
	b.PutID(InputNotifyBroadcastsTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputNotifyBroadcasts) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputNotifyBroadcasts#b1db7c7e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputNotifyBroadcasts) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputNotifyBroadcasts#b1db7c7e to nil")
	}
	if err := b.ConsumeID(InputNotifyBroadcastsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputNotifyBroadcasts#b1db7c7e: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputNotifyBroadcasts) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputNotifyBroadcasts#b1db7c7e to nil")
	}
	return nil
}

// construct implements constructor of InputNotifyPeerClass.
func (i InputNotifyBroadcasts) construct() InputNotifyPeerClass { return &i }

// Ensuring interfaces in compile-time for InputNotifyBroadcasts.
var (
	_ bin.Encoder     = &InputNotifyBroadcasts{}
	_ bin.Decoder     = &InputNotifyBroadcasts{}
	_ bin.BareEncoder = &InputNotifyBroadcasts{}
	_ bin.BareDecoder = &InputNotifyBroadcasts{}

	_ InputNotifyPeerClass = &InputNotifyBroadcasts{}
)

// InputNotifyPeerClass represents InputNotifyPeer generic type.
//
// See https://core.telegram.org/type/InputNotifyPeer for reference.
//
// Example:
//  g, err := tg.DecodeInputNotifyPeer(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputNotifyPeer: // inputNotifyPeer#b8bc5b0c
//  case *tg.InputNotifyUsers: // inputNotifyUsers#193b4417
//  case *tg.InputNotifyChats: // inputNotifyChats#4a95e84e
//  case *tg.InputNotifyBroadcasts: // inputNotifyBroadcasts#b1db7c7e
//  default: panic(v)
//  }
type InputNotifyPeerClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputNotifyPeerClass

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

// DecodeInputNotifyPeer implements binary de-serialization for InputNotifyPeerClass.
func DecodeInputNotifyPeer(buf *bin.Buffer) (InputNotifyPeerClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputNotifyPeerTypeID:
		// Decoding inputNotifyPeer#b8bc5b0c.
		v := InputNotifyPeer{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputNotifyPeerClass: %w", err)
		}
		return &v, nil
	case InputNotifyUsersTypeID:
		// Decoding inputNotifyUsers#193b4417.
		v := InputNotifyUsers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputNotifyPeerClass: %w", err)
		}
		return &v, nil
	case InputNotifyChatsTypeID:
		// Decoding inputNotifyChats#4a95e84e.
		v := InputNotifyChats{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputNotifyPeerClass: %w", err)
		}
		return &v, nil
	case InputNotifyBroadcastsTypeID:
		// Decoding inputNotifyBroadcasts#b1db7c7e.
		v := InputNotifyBroadcasts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputNotifyPeerClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputNotifyPeerClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputNotifyPeer boxes the InputNotifyPeerClass providing a helper.
type InputNotifyPeerBox struct {
	InputNotifyPeer InputNotifyPeerClass
}

// Decode implements bin.Decoder for InputNotifyPeerBox.
func (b *InputNotifyPeerBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputNotifyPeerBox to nil")
	}
	v, err := DecodeInputNotifyPeer(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputNotifyPeer = v
	return nil
}

// Encode implements bin.Encode for InputNotifyPeerBox.
func (b *InputNotifyPeerBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputNotifyPeer == nil {
		return fmt.Errorf("unable to encode InputNotifyPeerClass as nil")
	}
	return b.InputNotifyPeer.Encode(buf)
}

// InputNotifyPeerClassArray is adapter for slice of InputNotifyPeerClass.
type InputNotifyPeerClassArray []InputNotifyPeerClass

// Sort sorts slice of InputNotifyPeerClass.
func (s InputNotifyPeerClassArray) Sort(less func(a, b InputNotifyPeerClass) bool) InputNotifyPeerClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputNotifyPeerClass.
func (s InputNotifyPeerClassArray) SortStable(less func(a, b InputNotifyPeerClass) bool) InputNotifyPeerClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputNotifyPeerClass.
func (s InputNotifyPeerClassArray) Retain(keep func(x InputNotifyPeerClass) bool) InputNotifyPeerClassArray {
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
func (s InputNotifyPeerClassArray) First() (v InputNotifyPeerClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputNotifyPeerClassArray) Last() (v InputNotifyPeerClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputNotifyPeerClassArray) PopFirst() (v InputNotifyPeerClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputNotifyPeerClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputNotifyPeerClassArray) Pop() (v InputNotifyPeerClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputNotifyPeer returns copy with only InputNotifyPeer constructors.
func (s InputNotifyPeerClassArray) AsInputNotifyPeer() (to InputNotifyPeerArray) {
	for _, elem := range s {
		value, ok := elem.(*InputNotifyPeer)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// InputNotifyPeerArray is adapter for slice of InputNotifyPeer.
type InputNotifyPeerArray []InputNotifyPeer

// Sort sorts slice of InputNotifyPeer.
func (s InputNotifyPeerArray) Sort(less func(a, b InputNotifyPeer) bool) InputNotifyPeerArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputNotifyPeer.
func (s InputNotifyPeerArray) SortStable(less func(a, b InputNotifyPeer) bool) InputNotifyPeerArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputNotifyPeer.
func (s InputNotifyPeerArray) Retain(keep func(x InputNotifyPeer) bool) InputNotifyPeerArray {
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
func (s InputNotifyPeerArray) First() (v InputNotifyPeer, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputNotifyPeerArray) Last() (v InputNotifyPeer, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputNotifyPeerArray) PopFirst() (v InputNotifyPeer, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputNotifyPeer
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputNotifyPeerArray) Pop() (v InputNotifyPeer, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
