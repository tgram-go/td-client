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

// GroupCallParticipant represents TL type `groupCallParticipant#19adba89`.
//
// See https://core.telegram.org/constructor/groupCallParticipant for reference.
type GroupCallParticipant struct {
	// Flags field of GroupCallParticipant.
	Flags bin.Fields
	// Muted field of GroupCallParticipant.
	Muted bool
	// Left field of GroupCallParticipant.
	Left bool
	// CanSelfUnmute field of GroupCallParticipant.
	CanSelfUnmute bool
	// JustJoined field of GroupCallParticipant.
	JustJoined bool
	// Versioned field of GroupCallParticipant.
	Versioned bool
	// Min field of GroupCallParticipant.
	Min bool
	// MutedByYou field of GroupCallParticipant.
	MutedByYou bool
	// VolumeByAdmin field of GroupCallParticipant.
	VolumeByAdmin bool
	// Self field of GroupCallParticipant.
	Self bool
	// Peer field of GroupCallParticipant.
	Peer PeerClass
	// Date field of GroupCallParticipant.
	Date int
	// ActiveDate field of GroupCallParticipant.
	//
	// Use SetActiveDate and GetActiveDate helpers.
	ActiveDate int
	// Source field of GroupCallParticipant.
	Source int
	// Volume field of GroupCallParticipant.
	//
	// Use SetVolume and GetVolume helpers.
	Volume int
	// About field of GroupCallParticipant.
	//
	// Use SetAbout and GetAbout helpers.
	About string
	// RaiseHandRating field of GroupCallParticipant.
	//
	// Use SetRaiseHandRating and GetRaiseHandRating helpers.
	RaiseHandRating int64
}

// GroupCallParticipantTypeID is TL type id of GroupCallParticipant.
const GroupCallParticipantTypeID = 0x19adba89

func (g *GroupCallParticipant) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Muted == false) {
		return false
	}
	if !(g.Left == false) {
		return false
	}
	if !(g.CanSelfUnmute == false) {
		return false
	}
	if !(g.JustJoined == false) {
		return false
	}
	if !(g.Versioned == false) {
		return false
	}
	if !(g.Min == false) {
		return false
	}
	if !(g.MutedByYou == false) {
		return false
	}
	if !(g.VolumeByAdmin == false) {
		return false
	}
	if !(g.Self == false) {
		return false
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.Date == 0) {
		return false
	}
	if !(g.ActiveDate == 0) {
		return false
	}
	if !(g.Source == 0) {
		return false
	}
	if !(g.Volume == 0) {
		return false
	}
	if !(g.About == "") {
		return false
	}
	if !(g.RaiseHandRating == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GroupCallParticipant) String() string {
	if g == nil {
		return "GroupCallParticipant(nil)"
	}
	type Alias GroupCallParticipant
	return fmt.Sprintf("GroupCallParticipant%+v", Alias(*g))
}

// FillFrom fills GroupCallParticipant from given interface.
func (g *GroupCallParticipant) FillFrom(from interface {
	GetMuted() (value bool)
	GetLeft() (value bool)
	GetCanSelfUnmute() (value bool)
	GetJustJoined() (value bool)
	GetVersioned() (value bool)
	GetMin() (value bool)
	GetMutedByYou() (value bool)
	GetVolumeByAdmin() (value bool)
	GetSelf() (value bool)
	GetPeer() (value PeerClass)
	GetDate() (value int)
	GetActiveDate() (value int, ok bool)
	GetSource() (value int)
	GetVolume() (value int, ok bool)
	GetAbout() (value string, ok bool)
	GetRaiseHandRating() (value int64, ok bool)
}) {
	g.Muted = from.GetMuted()
	g.Left = from.GetLeft()
	g.CanSelfUnmute = from.GetCanSelfUnmute()
	g.JustJoined = from.GetJustJoined()
	g.Versioned = from.GetVersioned()
	g.Min = from.GetMin()
	g.MutedByYou = from.GetMutedByYou()
	g.VolumeByAdmin = from.GetVolumeByAdmin()
	g.Self = from.GetSelf()
	g.Peer = from.GetPeer()
	g.Date = from.GetDate()
	if val, ok := from.GetActiveDate(); ok {
		g.ActiveDate = val
	}

	g.Source = from.GetSource()
	if val, ok := from.GetVolume(); ok {
		g.Volume = val
	}

	if val, ok := from.GetAbout(); ok {
		g.About = val
	}

	if val, ok := from.GetRaiseHandRating(); ok {
		g.RaiseHandRating = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GroupCallParticipant) TypeID() uint32 {
	return GroupCallParticipantTypeID
}

// TypeName returns name of type in TL schema.
func (*GroupCallParticipant) TypeName() string {
	return "groupCallParticipant"
}

// TypeInfo returns info about TL type.
func (g *GroupCallParticipant) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "groupCallParticipant",
		ID:   GroupCallParticipantTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Muted",
			SchemaName: "muted",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "Left",
			SchemaName: "left",
			Null:       !g.Flags.Has(1),
		},
		{
			Name:       "CanSelfUnmute",
			SchemaName: "can_self_unmute",
			Null:       !g.Flags.Has(2),
		},
		{
			Name:       "JustJoined",
			SchemaName: "just_joined",
			Null:       !g.Flags.Has(4),
		},
		{
			Name:       "Versioned",
			SchemaName: "versioned",
			Null:       !g.Flags.Has(5),
		},
		{
			Name:       "Min",
			SchemaName: "min",
			Null:       !g.Flags.Has(8),
		},
		{
			Name:       "MutedByYou",
			SchemaName: "muted_by_you",
			Null:       !g.Flags.Has(9),
		},
		{
			Name:       "VolumeByAdmin",
			SchemaName: "volume_by_admin",
			Null:       !g.Flags.Has(10),
		},
		{
			Name:       "Self",
			SchemaName: "self",
			Null:       !g.Flags.Has(12),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "ActiveDate",
			SchemaName: "active_date",
			Null:       !g.Flags.Has(3),
		},
		{
			Name:       "Source",
			SchemaName: "source",
		},
		{
			Name:       "Volume",
			SchemaName: "volume",
			Null:       !g.Flags.Has(7),
		},
		{
			Name:       "About",
			SchemaName: "about",
			Null:       !g.Flags.Has(11),
		},
		{
			Name:       "RaiseHandRating",
			SchemaName: "raise_hand_rating",
			Null:       !g.Flags.Has(13),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GroupCallParticipant) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCallParticipant#19adba89 as nil")
	}
	b.PutID(GroupCallParticipantTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GroupCallParticipant) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCallParticipant#19adba89 as nil")
	}
	if !(g.Muted == false) {
		g.Flags.Set(0)
	}
	if !(g.Left == false) {
		g.Flags.Set(1)
	}
	if !(g.CanSelfUnmute == false) {
		g.Flags.Set(2)
	}
	if !(g.JustJoined == false) {
		g.Flags.Set(4)
	}
	if !(g.Versioned == false) {
		g.Flags.Set(5)
	}
	if !(g.Min == false) {
		g.Flags.Set(8)
	}
	if !(g.MutedByYou == false) {
		g.Flags.Set(9)
	}
	if !(g.VolumeByAdmin == false) {
		g.Flags.Set(10)
	}
	if !(g.Self == false) {
		g.Flags.Set(12)
	}
	if !(g.ActiveDate == 0) {
		g.Flags.Set(3)
	}
	if !(g.Volume == 0) {
		g.Flags.Set(7)
	}
	if !(g.About == "") {
		g.Flags.Set(11)
	}
	if !(g.RaiseHandRating == 0) {
		g.Flags.Set(13)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode groupCallParticipant#19adba89: field flags: %w", err)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode groupCallParticipant#19adba89: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode groupCallParticipant#19adba89: field peer: %w", err)
	}
	b.PutInt(g.Date)
	if g.Flags.Has(3) {
		b.PutInt(g.ActiveDate)
	}
	b.PutInt(g.Source)
	if g.Flags.Has(7) {
		b.PutInt(g.Volume)
	}
	if g.Flags.Has(11) {
		b.PutString(g.About)
	}
	if g.Flags.Has(13) {
		b.PutLong(g.RaiseHandRating)
	}
	return nil
}

// SetMuted sets value of Muted conditional field.
func (g *GroupCallParticipant) SetMuted(value bool) {
	if value {
		g.Flags.Set(0)
		g.Muted = true
	} else {
		g.Flags.Unset(0)
		g.Muted = false
	}
}

// GetMuted returns value of Muted conditional field.
func (g *GroupCallParticipant) GetMuted() (value bool) {
	return g.Flags.Has(0)
}

// SetLeft sets value of Left conditional field.
func (g *GroupCallParticipant) SetLeft(value bool) {
	if value {
		g.Flags.Set(1)
		g.Left = true
	} else {
		g.Flags.Unset(1)
		g.Left = false
	}
}

// GetLeft returns value of Left conditional field.
func (g *GroupCallParticipant) GetLeft() (value bool) {
	return g.Flags.Has(1)
}

// SetCanSelfUnmute sets value of CanSelfUnmute conditional field.
func (g *GroupCallParticipant) SetCanSelfUnmute(value bool) {
	if value {
		g.Flags.Set(2)
		g.CanSelfUnmute = true
	} else {
		g.Flags.Unset(2)
		g.CanSelfUnmute = false
	}
}

// GetCanSelfUnmute returns value of CanSelfUnmute conditional field.
func (g *GroupCallParticipant) GetCanSelfUnmute() (value bool) {
	return g.Flags.Has(2)
}

// SetJustJoined sets value of JustJoined conditional field.
func (g *GroupCallParticipant) SetJustJoined(value bool) {
	if value {
		g.Flags.Set(4)
		g.JustJoined = true
	} else {
		g.Flags.Unset(4)
		g.JustJoined = false
	}
}

// GetJustJoined returns value of JustJoined conditional field.
func (g *GroupCallParticipant) GetJustJoined() (value bool) {
	return g.Flags.Has(4)
}

// SetVersioned sets value of Versioned conditional field.
func (g *GroupCallParticipant) SetVersioned(value bool) {
	if value {
		g.Flags.Set(5)
		g.Versioned = true
	} else {
		g.Flags.Unset(5)
		g.Versioned = false
	}
}

// GetVersioned returns value of Versioned conditional field.
func (g *GroupCallParticipant) GetVersioned() (value bool) {
	return g.Flags.Has(5)
}

// SetMin sets value of Min conditional field.
func (g *GroupCallParticipant) SetMin(value bool) {
	if value {
		g.Flags.Set(8)
		g.Min = true
	} else {
		g.Flags.Unset(8)
		g.Min = false
	}
}

// GetMin returns value of Min conditional field.
func (g *GroupCallParticipant) GetMin() (value bool) {
	return g.Flags.Has(8)
}

// SetMutedByYou sets value of MutedByYou conditional field.
func (g *GroupCallParticipant) SetMutedByYou(value bool) {
	if value {
		g.Flags.Set(9)
		g.MutedByYou = true
	} else {
		g.Flags.Unset(9)
		g.MutedByYou = false
	}
}

// GetMutedByYou returns value of MutedByYou conditional field.
func (g *GroupCallParticipant) GetMutedByYou() (value bool) {
	return g.Flags.Has(9)
}

// SetVolumeByAdmin sets value of VolumeByAdmin conditional field.
func (g *GroupCallParticipant) SetVolumeByAdmin(value bool) {
	if value {
		g.Flags.Set(10)
		g.VolumeByAdmin = true
	} else {
		g.Flags.Unset(10)
		g.VolumeByAdmin = false
	}
}

// GetVolumeByAdmin returns value of VolumeByAdmin conditional field.
func (g *GroupCallParticipant) GetVolumeByAdmin() (value bool) {
	return g.Flags.Has(10)
}

// SetSelf sets value of Self conditional field.
func (g *GroupCallParticipant) SetSelf(value bool) {
	if value {
		g.Flags.Set(12)
		g.Self = true
	} else {
		g.Flags.Unset(12)
		g.Self = false
	}
}

// GetSelf returns value of Self conditional field.
func (g *GroupCallParticipant) GetSelf() (value bool) {
	return g.Flags.Has(12)
}

// GetPeer returns value of Peer field.
func (g *GroupCallParticipant) GetPeer() (value PeerClass) {
	return g.Peer
}

// GetDate returns value of Date field.
func (g *GroupCallParticipant) GetDate() (value int) {
	return g.Date
}

// SetActiveDate sets value of ActiveDate conditional field.
func (g *GroupCallParticipant) SetActiveDate(value int) {
	g.Flags.Set(3)
	g.ActiveDate = value
}

// GetActiveDate returns value of ActiveDate conditional field and
// boolean which is true if field was set.
func (g *GroupCallParticipant) GetActiveDate() (value int, ok bool) {
	if !g.Flags.Has(3) {
		return value, false
	}
	return g.ActiveDate, true
}

// GetSource returns value of Source field.
func (g *GroupCallParticipant) GetSource() (value int) {
	return g.Source
}

// SetVolume sets value of Volume conditional field.
func (g *GroupCallParticipant) SetVolume(value int) {
	g.Flags.Set(7)
	g.Volume = value
}

// GetVolume returns value of Volume conditional field and
// boolean which is true if field was set.
func (g *GroupCallParticipant) GetVolume() (value int, ok bool) {
	if !g.Flags.Has(7) {
		return value, false
	}
	return g.Volume, true
}

// SetAbout sets value of About conditional field.
func (g *GroupCallParticipant) SetAbout(value string) {
	g.Flags.Set(11)
	g.About = value
}

// GetAbout returns value of About conditional field and
// boolean which is true if field was set.
func (g *GroupCallParticipant) GetAbout() (value string, ok bool) {
	if !g.Flags.Has(11) {
		return value, false
	}
	return g.About, true
}

// SetRaiseHandRating sets value of RaiseHandRating conditional field.
func (g *GroupCallParticipant) SetRaiseHandRating(value int64) {
	g.Flags.Set(13)
	g.RaiseHandRating = value
}

// GetRaiseHandRating returns value of RaiseHandRating conditional field and
// boolean which is true if field was set.
func (g *GroupCallParticipant) GetRaiseHandRating() (value int64, ok bool) {
	if !g.Flags.Has(13) {
		return value, false
	}
	return g.RaiseHandRating, true
}

// Decode implements bin.Decoder.
func (g *GroupCallParticipant) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCallParticipant#19adba89 to nil")
	}
	if err := b.ConsumeID(GroupCallParticipantTypeID); err != nil {
		return fmt.Errorf("unable to decode groupCallParticipant#19adba89: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GroupCallParticipant) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCallParticipant#19adba89 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode groupCallParticipant#19adba89: field flags: %w", err)
		}
	}
	g.Muted = g.Flags.Has(0)
	g.Left = g.Flags.Has(1)
	g.CanSelfUnmute = g.Flags.Has(2)
	g.JustJoined = g.Flags.Has(4)
	g.Versioned = g.Flags.Has(5)
	g.Min = g.Flags.Has(8)
	g.MutedByYou = g.Flags.Has(9)
	g.VolumeByAdmin = g.Flags.Has(10)
	g.Self = g.Flags.Has(12)
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode groupCallParticipant#19adba89: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallParticipant#19adba89: field date: %w", err)
		}
		g.Date = value
	}
	if g.Flags.Has(3) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallParticipant#19adba89: field active_date: %w", err)
		}
		g.ActiveDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallParticipant#19adba89: field source: %w", err)
		}
		g.Source = value
	}
	if g.Flags.Has(7) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallParticipant#19adba89: field volume: %w", err)
		}
		g.Volume = value
	}
	if g.Flags.Has(11) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallParticipant#19adba89: field about: %w", err)
		}
		g.About = value
	}
	if g.Flags.Has(13) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallParticipant#19adba89: field raise_hand_rating: %w", err)
		}
		g.RaiseHandRating = value
	}
	return nil
}

// Ensuring interfaces in compile-time for GroupCallParticipant.
var (
	_ bin.Encoder     = &GroupCallParticipant{}
	_ bin.Decoder     = &GroupCallParticipant{}
	_ bin.BareEncoder = &GroupCallParticipant{}
	_ bin.BareDecoder = &GroupCallParticipant{}
)
