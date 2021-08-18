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

// StickerSet represents TL type `stickerSet#d7df217a`.
// Represents a stickerset (stickerpack)
//
// See https://core.telegram.org/constructor/stickerSet for reference.
type StickerSet struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this stickerset was archived (due to too many saved stickers in the current
	// account)
	Archived bool
	// Is this stickerset official
	Official bool
	// Is this a mask stickerset
	Masks bool
	// Is this an animated stickerpack
	Animated bool
	// When was this stickerset installed
	//
	// Use SetInstalledDate and GetInstalledDate helpers.
	InstalledDate int
	// ID of the stickerset
	ID int64
	// Access hash of stickerset
	AccessHash int64
	// Title of stickerset
	Title string
	// Short name of stickerset to use in tg://addstickers?set=short_name
	ShortName string
	// Thumbs field of StickerSet.
	//
	// Use SetThumbs and GetThumbs helpers.
	Thumbs []PhotoSizeClass
	// DC ID of thumbnail
	//
	// Use SetThumbDCID and GetThumbDCID helpers.
	ThumbDCID int
	// ThumbVersion field of StickerSet.
	//
	// Use SetThumbVersion and GetThumbVersion helpers.
	ThumbVersion int
	// Number of stickers in pack
	Count int
	// Hash
	Hash int
}

// StickerSetTypeID is TL type id of StickerSet.
const StickerSetTypeID = 0xd7df217a

func (s *StickerSet) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Archived == false) {
		return false
	}
	if !(s.Official == false) {
		return false
	}
	if !(s.Masks == false) {
		return false
	}
	if !(s.Animated == false) {
		return false
	}
	if !(s.InstalledDate == 0) {
		return false
	}
	if !(s.ID == 0) {
		return false
	}
	if !(s.AccessHash == 0) {
		return false
	}
	if !(s.Title == "") {
		return false
	}
	if !(s.ShortName == "") {
		return false
	}
	if !(s.Thumbs == nil) {
		return false
	}
	if !(s.ThumbDCID == 0) {
		return false
	}
	if !(s.ThumbVersion == 0) {
		return false
	}
	if !(s.Count == 0) {
		return false
	}
	if !(s.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StickerSet) String() string {
	if s == nil {
		return "StickerSet(nil)"
	}
	type Alias StickerSet
	return fmt.Sprintf("StickerSet%+v", Alias(*s))
}

// FillFrom fills StickerSet from given interface.
func (s *StickerSet) FillFrom(from interface {
	GetArchived() (value bool)
	GetOfficial() (value bool)
	GetMasks() (value bool)
	GetAnimated() (value bool)
	GetInstalledDate() (value int, ok bool)
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetTitle() (value string)
	GetShortName() (value string)
	GetThumbs() (value []PhotoSizeClass, ok bool)
	GetThumbDCID() (value int, ok bool)
	GetThumbVersion() (value int, ok bool)
	GetCount() (value int)
	GetHash() (value int)
}) {
	s.Archived = from.GetArchived()
	s.Official = from.GetOfficial()
	s.Masks = from.GetMasks()
	s.Animated = from.GetAnimated()
	if val, ok := from.GetInstalledDate(); ok {
		s.InstalledDate = val
	}

	s.ID = from.GetID()
	s.AccessHash = from.GetAccessHash()
	s.Title = from.GetTitle()
	s.ShortName = from.GetShortName()
	if val, ok := from.GetThumbs(); ok {
		s.Thumbs = val
	}

	if val, ok := from.GetThumbDCID(); ok {
		s.ThumbDCID = val
	}

	if val, ok := from.GetThumbVersion(); ok {
		s.ThumbVersion = val
	}

	s.Count = from.GetCount()
	s.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickerSet) TypeID() uint32 {
	return StickerSetTypeID
}

// TypeName returns name of type in TL schema.
func (*StickerSet) TypeName() string {
	return "stickerSet"
}

// TypeInfo returns info about TL type.
func (s *StickerSet) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickerSet",
		ID:   StickerSetTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Archived",
			SchemaName: "archived",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "Official",
			SchemaName: "official",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "Masks",
			SchemaName: "masks",
			Null:       !s.Flags.Has(3),
		},
		{
			Name:       "Animated",
			SchemaName: "animated",
			Null:       !s.Flags.Has(5),
		},
		{
			Name:       "InstalledDate",
			SchemaName: "installed_date",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "ShortName",
			SchemaName: "short_name",
		},
		{
			Name:       "Thumbs",
			SchemaName: "thumbs",
			Null:       !s.Flags.Has(4),
		},
		{
			Name:       "ThumbDCID",
			SchemaName: "thumb_dc_id",
			Null:       !s.Flags.Has(4),
		},
		{
			Name:       "ThumbVersion",
			SchemaName: "thumb_version",
			Null:       !s.Flags.Has(4),
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StickerSet) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerSet#d7df217a as nil")
	}
	b.PutID(StickerSetTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StickerSet) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerSet#d7df217a as nil")
	}
	if !(s.Archived == false) {
		s.Flags.Set(1)
	}
	if !(s.Official == false) {
		s.Flags.Set(2)
	}
	if !(s.Masks == false) {
		s.Flags.Set(3)
	}
	if !(s.Animated == false) {
		s.Flags.Set(5)
	}
	if !(s.InstalledDate == 0) {
		s.Flags.Set(0)
	}
	if !(s.Thumbs == nil) {
		s.Flags.Set(4)
	}
	if !(s.ThumbDCID == 0) {
		s.Flags.Set(4)
	}
	if !(s.ThumbVersion == 0) {
		s.Flags.Set(4)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickerSet#d7df217a: field flags: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.InstalledDate)
	}
	b.PutLong(s.ID)
	b.PutLong(s.AccessHash)
	b.PutString(s.Title)
	b.PutString(s.ShortName)
	if s.Flags.Has(4) {
		b.PutVectorHeader(len(s.Thumbs))
		for idx, v := range s.Thumbs {
			if v == nil {
				return fmt.Errorf("unable to encode stickerSet#d7df217a: field thumbs element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode stickerSet#d7df217a: field thumbs element with index %d: %w", idx, err)
			}
		}
	}
	if s.Flags.Has(4) {
		b.PutInt(s.ThumbDCID)
	}
	if s.Flags.Has(4) {
		b.PutInt(s.ThumbVersion)
	}
	b.PutInt(s.Count)
	b.PutInt(s.Hash)
	return nil
}

// SetArchived sets value of Archived conditional field.
func (s *StickerSet) SetArchived(value bool) {
	if value {
		s.Flags.Set(1)
		s.Archived = true
	} else {
		s.Flags.Unset(1)
		s.Archived = false
	}
}

// GetArchived returns value of Archived conditional field.
func (s *StickerSet) GetArchived() (value bool) {
	return s.Flags.Has(1)
}

// SetOfficial sets value of Official conditional field.
func (s *StickerSet) SetOfficial(value bool) {
	if value {
		s.Flags.Set(2)
		s.Official = true
	} else {
		s.Flags.Unset(2)
		s.Official = false
	}
}

// GetOfficial returns value of Official conditional field.
func (s *StickerSet) GetOfficial() (value bool) {
	return s.Flags.Has(2)
}

// SetMasks sets value of Masks conditional field.
func (s *StickerSet) SetMasks(value bool) {
	if value {
		s.Flags.Set(3)
		s.Masks = true
	} else {
		s.Flags.Unset(3)
		s.Masks = false
	}
}

// GetMasks returns value of Masks conditional field.
func (s *StickerSet) GetMasks() (value bool) {
	return s.Flags.Has(3)
}

// SetAnimated sets value of Animated conditional field.
func (s *StickerSet) SetAnimated(value bool) {
	if value {
		s.Flags.Set(5)
		s.Animated = true
	} else {
		s.Flags.Unset(5)
		s.Animated = false
	}
}

// GetAnimated returns value of Animated conditional field.
func (s *StickerSet) GetAnimated() (value bool) {
	return s.Flags.Has(5)
}

// SetInstalledDate sets value of InstalledDate conditional field.
func (s *StickerSet) SetInstalledDate(value int) {
	s.Flags.Set(0)
	s.InstalledDate = value
}

// GetInstalledDate returns value of InstalledDate conditional field and
// boolean which is true if field was set.
func (s *StickerSet) GetInstalledDate() (value int, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.InstalledDate, true
}

// GetID returns value of ID field.
func (s *StickerSet) GetID() (value int64) {
	return s.ID
}

// GetAccessHash returns value of AccessHash field.
func (s *StickerSet) GetAccessHash() (value int64) {
	return s.AccessHash
}

// GetTitle returns value of Title field.
func (s *StickerSet) GetTitle() (value string) {
	return s.Title
}

// GetShortName returns value of ShortName field.
func (s *StickerSet) GetShortName() (value string) {
	return s.ShortName
}

// SetThumbs sets value of Thumbs conditional field.
func (s *StickerSet) SetThumbs(value []PhotoSizeClass) {
	s.Flags.Set(4)
	s.Thumbs = value
}

// GetThumbs returns value of Thumbs conditional field and
// boolean which is true if field was set.
func (s *StickerSet) GetThumbs() (value []PhotoSizeClass, ok bool) {
	if !s.Flags.Has(4) {
		return value, false
	}
	return s.Thumbs, true
}

// MapThumbs returns field Thumbs wrapped in PhotoSizeClassArray helper.
func (s *StickerSet) MapThumbs() (value PhotoSizeClassArray, ok bool) {
	if !s.Flags.Has(4) {
		return value, false
	}
	return PhotoSizeClassArray(s.Thumbs), true
}

// SetThumbDCID sets value of ThumbDCID conditional field.
func (s *StickerSet) SetThumbDCID(value int) {
	s.Flags.Set(4)
	s.ThumbDCID = value
}

// GetThumbDCID returns value of ThumbDCID conditional field and
// boolean which is true if field was set.
func (s *StickerSet) GetThumbDCID() (value int, ok bool) {
	if !s.Flags.Has(4) {
		return value, false
	}
	return s.ThumbDCID, true
}

// SetThumbVersion sets value of ThumbVersion conditional field.
func (s *StickerSet) SetThumbVersion(value int) {
	s.Flags.Set(4)
	s.ThumbVersion = value
}

// GetThumbVersion returns value of ThumbVersion conditional field and
// boolean which is true if field was set.
func (s *StickerSet) GetThumbVersion() (value int, ok bool) {
	if !s.Flags.Has(4) {
		return value, false
	}
	return s.ThumbVersion, true
}

// GetCount returns value of Count field.
func (s *StickerSet) GetCount() (value int) {
	return s.Count
}

// GetHash returns value of Hash field.
func (s *StickerSet) GetHash() (value int) {
	return s.Hash
}

// Decode implements bin.Decoder.
func (s *StickerSet) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerSet#d7df217a to nil")
	}
	if err := b.ConsumeID(StickerSetTypeID); err != nil {
		return fmt.Errorf("unable to decode stickerSet#d7df217a: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StickerSet) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerSet#d7df217a to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stickerSet#d7df217a: field flags: %w", err)
		}
	}
	s.Archived = s.Flags.Has(1)
	s.Official = s.Flags.Has(2)
	s.Masks = s.Flags.Has(3)
	s.Animated = s.Flags.Has(5)
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#d7df217a: field installed_date: %w", err)
		}
		s.InstalledDate = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#d7df217a: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#d7df217a: field access_hash: %w", err)
		}
		s.AccessHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#d7df217a: field title: %w", err)
		}
		s.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#d7df217a: field short_name: %w", err)
		}
		s.ShortName = value
	}
	if s.Flags.Has(4) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#d7df217a: field thumbs: %w", err)
		}

		if headerLen != 0 {
			s.Thumbs = make([]PhotoSizeClass, 0, headerLen)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePhotoSize(b)
			if err != nil {
				return fmt.Errorf("unable to decode stickerSet#d7df217a: field thumbs: %w", err)
			}
			s.Thumbs = append(s.Thumbs, value)
		}
	}
	if s.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#d7df217a: field thumb_dc_id: %w", err)
		}
		s.ThumbDCID = value
	}
	if s.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#d7df217a: field thumb_version: %w", err)
		}
		s.ThumbVersion = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#d7df217a: field count: %w", err)
		}
		s.Count = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#d7df217a: field hash: %w", err)
		}
		s.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StickerSet.
var (
	_ bin.Encoder     = &StickerSet{}
	_ bin.Decoder     = &StickerSet{}
	_ bin.BareEncoder = &StickerSet{}
	_ bin.BareDecoder = &StickerSet{}
)
