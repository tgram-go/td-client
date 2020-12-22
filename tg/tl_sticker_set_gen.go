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

// StickerSet represents TL type `stickerSet#eeb46f27`.
// Represents a stickerset (stickerpack)
//
// See https://core.telegram.org/constructor/stickerSet for reference.
type StickerSet struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this stickerset was archived (due to too many saved stickers in the current account)
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
	// Thumbnail for stickerset
	//
	// Use SetThumb and GetThumb helpers.
	Thumb PhotoSizeClass
	// DC ID of thumbnail
	//
	// Use SetThumbDCID and GetThumbDCID helpers.
	ThumbDCID int
	// Number of stickers in pack
	Count int
	// Hash
	Hash int
}

// StickerSetTypeID is TL type id of StickerSet.
const StickerSetTypeID = 0xeeb46f27

// Encode implements bin.Encoder.
func (s *StickerSet) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerSet#eeb46f27 as nil")
	}
	b.PutID(StickerSetTypeID)
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickerSet#eeb46f27: field flags: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.InstalledDate)
	}
	b.PutLong(s.ID)
	b.PutLong(s.AccessHash)
	b.PutString(s.Title)
	b.PutString(s.ShortName)
	if s.Flags.Has(4) {
		if s.Thumb == nil {
			return fmt.Errorf("unable to encode stickerSet#eeb46f27: field thumb is nil")
		}
		if err := s.Thumb.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stickerSet#eeb46f27: field thumb: %w", err)
		}
	}
	if s.Flags.Has(4) {
		b.PutInt(s.ThumbDCID)
	}
	b.PutInt(s.Count)
	b.PutInt(s.Hash)
	return nil
}

// SetArchived sets value of Archived conditional field.
func (s *StickerSet) SetArchived(value bool) {
	if value {
		s.Flags.Set(1)
	} else {
		s.Flags.Unset(1)
	}
}

// SetOfficial sets value of Official conditional field.
func (s *StickerSet) SetOfficial(value bool) {
	if value {
		s.Flags.Set(2)
	} else {
		s.Flags.Unset(2)
	}
}

// SetMasks sets value of Masks conditional field.
func (s *StickerSet) SetMasks(value bool) {
	if value {
		s.Flags.Set(3)
	} else {
		s.Flags.Unset(3)
	}
}

// SetAnimated sets value of Animated conditional field.
func (s *StickerSet) SetAnimated(value bool) {
	if value {
		s.Flags.Set(5)
	} else {
		s.Flags.Unset(5)
	}
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

// SetThumb sets value of Thumb conditional field.
func (s *StickerSet) SetThumb(value PhotoSizeClass) {
	s.Flags.Set(4)
	s.Thumb = value
}

// GetThumb returns value of Thumb conditional field and
// boolean which is true if field was set.
func (s *StickerSet) GetThumb() (value PhotoSizeClass, ok bool) {
	if !s.Flags.Has(4) {
		return value, false
	}
	return s.Thumb, true
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

// Decode implements bin.Decoder.
func (s *StickerSet) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerSet#eeb46f27 to nil")
	}
	if err := b.ConsumeID(StickerSetTypeID); err != nil {
		return fmt.Errorf("unable to decode stickerSet#eeb46f27: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stickerSet#eeb46f27: field flags: %w", err)
		}
	}
	s.Archived = s.Flags.Has(1)
	s.Official = s.Flags.Has(2)
	s.Masks = s.Flags.Has(3)
	s.Animated = s.Flags.Has(5)
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#eeb46f27: field installed_date: %w", err)
		}
		s.InstalledDate = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#eeb46f27: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#eeb46f27: field access_hash: %w", err)
		}
		s.AccessHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#eeb46f27: field title: %w", err)
		}
		s.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#eeb46f27: field short_name: %w", err)
		}
		s.ShortName = value
	}
	if s.Flags.Has(4) {
		value, err := DecodePhotoSize(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#eeb46f27: field thumb: %w", err)
		}
		s.Thumb = value
	}
	if s.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#eeb46f27: field thumb_dc_id: %w", err)
		}
		s.ThumbDCID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#eeb46f27: field count: %w", err)
		}
		s.Count = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSet#eeb46f27: field hash: %w", err)
		}
		s.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StickerSet.
var (
	_ bin.Encoder = &StickerSet{}
	_ bin.Decoder = &StickerSet{}
)
