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

// UserFull represents TL type `userFull#139a9a77`.
// Extended user info
//
// See https://core.telegram.org/constructor/userFull for reference.
type UserFull struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether you have blocked this user
	Blocked bool
	// Whether this user can make VoIP calls
	PhoneCallsAvailable bool
	// Whether this user's privacy settings allow you to call him
	PhoneCallsPrivate bool
	// Whether you can pin messages in the chat with this user, you can do this only for a chat with yourself
	CanPinMessage bool
	// Whether scheduled messages¹ are available
	//
	// Links:
	//  1) https://core.telegram.org/api/scheduled-messages
	HasScheduled bool
	// Whether the user can receive video calls
	VideoCallsAvailable bool
	// Remaining user info
	User UserClass
	// Bio of the user
	//
	// Use SetAbout and GetAbout helpers.
	About string
	// Peer settings
	Settings PeerSettings
	// Profile photo
	//
	// Use SetProfilePhoto and GetProfilePhoto helpers.
	ProfilePhoto PhotoClass
	// Notification settings
	NotifySettings PeerNotifySettings
	// For bots, info about the bot (bot commands, etc)
	//
	// Use SetBotInfo and GetBotInfo helpers.
	BotInfo BotInfo
	// Message ID of the last pinned message¹
	//
	// Links:
	//  1) https://core.telegram.org/api/pin
	//
	// Use SetPinnedMsgID and GetPinnedMsgID helpers.
	PinnedMsgID int
	// Chats in common with this user
	CommonChatsCount int
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	//
	// Use SetFolderID and GetFolderID helpers.
	FolderID int
	// TTLPeriod field of UserFull.
	//
	// Use SetTTLPeriod and GetTTLPeriod helpers.
	TTLPeriod int
}

// UserFullTypeID is TL type id of UserFull.
const UserFullTypeID = 0x139a9a77

func (u *UserFull) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.Blocked == false) {
		return false
	}
	if !(u.PhoneCallsAvailable == false) {
		return false
	}
	if !(u.PhoneCallsPrivate == false) {
		return false
	}
	if !(u.CanPinMessage == false) {
		return false
	}
	if !(u.HasScheduled == false) {
		return false
	}
	if !(u.VideoCallsAvailable == false) {
		return false
	}
	if !(u.User == nil) {
		return false
	}
	if !(u.About == "") {
		return false
	}
	if !(u.Settings.Zero()) {
		return false
	}
	if !(u.ProfilePhoto == nil) {
		return false
	}
	if !(u.NotifySettings.Zero()) {
		return false
	}
	if !(u.BotInfo.Zero()) {
		return false
	}
	if !(u.PinnedMsgID == 0) {
		return false
	}
	if !(u.CommonChatsCount == 0) {
		return false
	}
	if !(u.FolderID == 0) {
		return false
	}
	if !(u.TTLPeriod == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserFull) String() string {
	if u == nil {
		return "UserFull(nil)"
	}
	type Alias UserFull
	return fmt.Sprintf("UserFull%+v", Alias(*u))
}

// FillFrom fills UserFull from given interface.
func (u *UserFull) FillFrom(from interface {
	GetBlocked() (value bool)
	GetPhoneCallsAvailable() (value bool)
	GetPhoneCallsPrivate() (value bool)
	GetCanPinMessage() (value bool)
	GetHasScheduled() (value bool)
	GetVideoCallsAvailable() (value bool)
	GetUser() (value UserClass)
	GetAbout() (value string, ok bool)
	GetSettings() (value PeerSettings)
	GetProfilePhoto() (value PhotoClass, ok bool)
	GetNotifySettings() (value PeerNotifySettings)
	GetBotInfo() (value BotInfo, ok bool)
	GetPinnedMsgID() (value int, ok bool)
	GetCommonChatsCount() (value int)
	GetFolderID() (value int, ok bool)
	GetTTLPeriod() (value int, ok bool)
}) {
	u.Blocked = from.GetBlocked()
	u.PhoneCallsAvailable = from.GetPhoneCallsAvailable()
	u.PhoneCallsPrivate = from.GetPhoneCallsPrivate()
	u.CanPinMessage = from.GetCanPinMessage()
	u.HasScheduled = from.GetHasScheduled()
	u.VideoCallsAvailable = from.GetVideoCallsAvailable()
	u.User = from.GetUser()
	if val, ok := from.GetAbout(); ok {
		u.About = val
	}

	u.Settings = from.GetSettings()
	if val, ok := from.GetProfilePhoto(); ok {
		u.ProfilePhoto = val
	}

	u.NotifySettings = from.GetNotifySettings()
	if val, ok := from.GetBotInfo(); ok {
		u.BotInfo = val
	}

	if val, ok := from.GetPinnedMsgID(); ok {
		u.PinnedMsgID = val
	}

	u.CommonChatsCount = from.GetCommonChatsCount()
	if val, ok := from.GetFolderID(); ok {
		u.FolderID = val
	}

	if val, ok := from.GetTTLPeriod(); ok {
		u.TTLPeriod = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserFull) TypeID() uint32 {
	return UserFullTypeID
}

// TypeName returns name of type in TL schema.
func (*UserFull) TypeName() string {
	return "userFull"
}

// TypeInfo returns info about TL type.
func (u *UserFull) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "userFull",
		ID:   UserFullTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Blocked",
			SchemaName: "blocked",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "PhoneCallsAvailable",
			SchemaName: "phone_calls_available",
			Null:       !u.Flags.Has(4),
		},
		{
			Name:       "PhoneCallsPrivate",
			SchemaName: "phone_calls_private",
			Null:       !u.Flags.Has(5),
		},
		{
			Name:       "CanPinMessage",
			SchemaName: "can_pin_message",
			Null:       !u.Flags.Has(7),
		},
		{
			Name:       "HasScheduled",
			SchemaName: "has_scheduled",
			Null:       !u.Flags.Has(12),
		},
		{
			Name:       "VideoCallsAvailable",
			SchemaName: "video_calls_available",
			Null:       !u.Flags.Has(13),
		},
		{
			Name:       "User",
			SchemaName: "user",
		},
		{
			Name:       "About",
			SchemaName: "about",
			Null:       !u.Flags.Has(1),
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
		{
			Name:       "ProfilePhoto",
			SchemaName: "profile_photo",
			Null:       !u.Flags.Has(2),
		},
		{
			Name:       "NotifySettings",
			SchemaName: "notify_settings",
		},
		{
			Name:       "BotInfo",
			SchemaName: "bot_info",
			Null:       !u.Flags.Has(3),
		},
		{
			Name:       "PinnedMsgID",
			SchemaName: "pinned_msg_id",
			Null:       !u.Flags.Has(6),
		},
		{
			Name:       "CommonChatsCount",
			SchemaName: "common_chats_count",
		},
		{
			Name:       "FolderID",
			SchemaName: "folder_id",
			Null:       !u.Flags.Has(11),
		},
		{
			Name:       "TTLPeriod",
			SchemaName: "ttl_period",
			Null:       !u.Flags.Has(14),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UserFull) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userFull#139a9a77 as nil")
	}
	b.PutID(UserFullTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UserFull) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userFull#139a9a77 as nil")
	}
	if !(u.Blocked == false) {
		u.Flags.Set(0)
	}
	if !(u.PhoneCallsAvailable == false) {
		u.Flags.Set(4)
	}
	if !(u.PhoneCallsPrivate == false) {
		u.Flags.Set(5)
	}
	if !(u.CanPinMessage == false) {
		u.Flags.Set(7)
	}
	if !(u.HasScheduled == false) {
		u.Flags.Set(12)
	}
	if !(u.VideoCallsAvailable == false) {
		u.Flags.Set(13)
	}
	if !(u.About == "") {
		u.Flags.Set(1)
	}
	if !(u.ProfilePhoto == nil) {
		u.Flags.Set(2)
	}
	if !(u.BotInfo.Zero()) {
		u.Flags.Set(3)
	}
	if !(u.PinnedMsgID == 0) {
		u.Flags.Set(6)
	}
	if !(u.FolderID == 0) {
		u.Flags.Set(11)
	}
	if !(u.TTLPeriod == 0) {
		u.Flags.Set(14)
	}
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFull#139a9a77: field flags: %w", err)
	}
	if u.User == nil {
		return fmt.Errorf("unable to encode userFull#139a9a77: field user is nil")
	}
	if err := u.User.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFull#139a9a77: field user: %w", err)
	}
	if u.Flags.Has(1) {
		b.PutString(u.About)
	}
	if err := u.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFull#139a9a77: field settings: %w", err)
	}
	if u.Flags.Has(2) {
		if u.ProfilePhoto == nil {
			return fmt.Errorf("unable to encode userFull#139a9a77: field profile_photo is nil")
		}
		if err := u.ProfilePhoto.Encode(b); err != nil {
			return fmt.Errorf("unable to encode userFull#139a9a77: field profile_photo: %w", err)
		}
	}
	if err := u.NotifySettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFull#139a9a77: field notify_settings: %w", err)
	}
	if u.Flags.Has(3) {
		if err := u.BotInfo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode userFull#139a9a77: field bot_info: %w", err)
		}
	}
	if u.Flags.Has(6) {
		b.PutInt(u.PinnedMsgID)
	}
	b.PutInt(u.CommonChatsCount)
	if u.Flags.Has(11) {
		b.PutInt(u.FolderID)
	}
	if u.Flags.Has(14) {
		b.PutInt(u.TTLPeriod)
	}
	return nil
}

// SetBlocked sets value of Blocked conditional field.
func (u *UserFull) SetBlocked(value bool) {
	if value {
		u.Flags.Set(0)
		u.Blocked = true
	} else {
		u.Flags.Unset(0)
		u.Blocked = false
	}
}

// GetBlocked returns value of Blocked conditional field.
func (u *UserFull) GetBlocked() (value bool) {
	return u.Flags.Has(0)
}

// SetPhoneCallsAvailable sets value of PhoneCallsAvailable conditional field.
func (u *UserFull) SetPhoneCallsAvailable(value bool) {
	if value {
		u.Flags.Set(4)
		u.PhoneCallsAvailable = true
	} else {
		u.Flags.Unset(4)
		u.PhoneCallsAvailable = false
	}
}

// GetPhoneCallsAvailable returns value of PhoneCallsAvailable conditional field.
func (u *UserFull) GetPhoneCallsAvailable() (value bool) {
	return u.Flags.Has(4)
}

// SetPhoneCallsPrivate sets value of PhoneCallsPrivate conditional field.
func (u *UserFull) SetPhoneCallsPrivate(value bool) {
	if value {
		u.Flags.Set(5)
		u.PhoneCallsPrivate = true
	} else {
		u.Flags.Unset(5)
		u.PhoneCallsPrivate = false
	}
}

// GetPhoneCallsPrivate returns value of PhoneCallsPrivate conditional field.
func (u *UserFull) GetPhoneCallsPrivate() (value bool) {
	return u.Flags.Has(5)
}

// SetCanPinMessage sets value of CanPinMessage conditional field.
func (u *UserFull) SetCanPinMessage(value bool) {
	if value {
		u.Flags.Set(7)
		u.CanPinMessage = true
	} else {
		u.Flags.Unset(7)
		u.CanPinMessage = false
	}
}

// GetCanPinMessage returns value of CanPinMessage conditional field.
func (u *UserFull) GetCanPinMessage() (value bool) {
	return u.Flags.Has(7)
}

// SetHasScheduled sets value of HasScheduled conditional field.
func (u *UserFull) SetHasScheduled(value bool) {
	if value {
		u.Flags.Set(12)
		u.HasScheduled = true
	} else {
		u.Flags.Unset(12)
		u.HasScheduled = false
	}
}

// GetHasScheduled returns value of HasScheduled conditional field.
func (u *UserFull) GetHasScheduled() (value bool) {
	return u.Flags.Has(12)
}

// SetVideoCallsAvailable sets value of VideoCallsAvailable conditional field.
func (u *UserFull) SetVideoCallsAvailable(value bool) {
	if value {
		u.Flags.Set(13)
		u.VideoCallsAvailable = true
	} else {
		u.Flags.Unset(13)
		u.VideoCallsAvailable = false
	}
}

// GetVideoCallsAvailable returns value of VideoCallsAvailable conditional field.
func (u *UserFull) GetVideoCallsAvailable() (value bool) {
	return u.Flags.Has(13)
}

// GetUser returns value of User field.
func (u *UserFull) GetUser() (value UserClass) {
	return u.User
}

// GetUserAsNotEmpty returns mapped value of User field.
func (u *UserFull) GetUserAsNotEmpty() (*User, bool) {
	return u.User.AsNotEmpty()
}

// SetAbout sets value of About conditional field.
func (u *UserFull) SetAbout(value string) {
	u.Flags.Set(1)
	u.About = value
}

// GetAbout returns value of About conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetAbout() (value string, ok bool) {
	if !u.Flags.Has(1) {
		return value, false
	}
	return u.About, true
}

// GetSettings returns value of Settings field.
func (u *UserFull) GetSettings() (value PeerSettings) {
	return u.Settings
}

// SetProfilePhoto sets value of ProfilePhoto conditional field.
func (u *UserFull) SetProfilePhoto(value PhotoClass) {
	u.Flags.Set(2)
	u.ProfilePhoto = value
}

// GetProfilePhoto returns value of ProfilePhoto conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetProfilePhoto() (value PhotoClass, ok bool) {
	if !u.Flags.Has(2) {
		return value, false
	}
	return u.ProfilePhoto, true
}

// GetProfilePhotoAsNotEmpty returns mapped value of ProfilePhoto conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetProfilePhotoAsNotEmpty() (*Photo, bool) {
	if value, ok := u.GetProfilePhoto(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}

// GetNotifySettings returns value of NotifySettings field.
func (u *UserFull) GetNotifySettings() (value PeerNotifySettings) {
	return u.NotifySettings
}

// SetBotInfo sets value of BotInfo conditional field.
func (u *UserFull) SetBotInfo(value BotInfo) {
	u.Flags.Set(3)
	u.BotInfo = value
}

// GetBotInfo returns value of BotInfo conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetBotInfo() (value BotInfo, ok bool) {
	if !u.Flags.Has(3) {
		return value, false
	}
	return u.BotInfo, true
}

// SetPinnedMsgID sets value of PinnedMsgID conditional field.
func (u *UserFull) SetPinnedMsgID(value int) {
	u.Flags.Set(6)
	u.PinnedMsgID = value
}

// GetPinnedMsgID returns value of PinnedMsgID conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetPinnedMsgID() (value int, ok bool) {
	if !u.Flags.Has(6) {
		return value, false
	}
	return u.PinnedMsgID, true
}

// GetCommonChatsCount returns value of CommonChatsCount field.
func (u *UserFull) GetCommonChatsCount() (value int) {
	return u.CommonChatsCount
}

// SetFolderID sets value of FolderID conditional field.
func (u *UserFull) SetFolderID(value int) {
	u.Flags.Set(11)
	u.FolderID = value
}

// GetFolderID returns value of FolderID conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetFolderID() (value int, ok bool) {
	if !u.Flags.Has(11) {
		return value, false
	}
	return u.FolderID, true
}

// SetTTLPeriod sets value of TTLPeriod conditional field.
func (u *UserFull) SetTTLPeriod(value int) {
	u.Flags.Set(14)
	u.TTLPeriod = value
}

// GetTTLPeriod returns value of TTLPeriod conditional field and
// boolean which is true if field was set.
func (u *UserFull) GetTTLPeriod() (value int, ok bool) {
	if !u.Flags.Has(14) {
		return value, false
	}
	return u.TTLPeriod, true
}

// Decode implements bin.Decoder.
func (u *UserFull) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userFull#139a9a77 to nil")
	}
	if err := b.ConsumeID(UserFullTypeID); err != nil {
		return fmt.Errorf("unable to decode userFull#139a9a77: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UserFull) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userFull#139a9a77 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFull#139a9a77: field flags: %w", err)
		}
	}
	u.Blocked = u.Flags.Has(0)
	u.PhoneCallsAvailable = u.Flags.Has(4)
	u.PhoneCallsPrivate = u.Flags.Has(5)
	u.CanPinMessage = u.Flags.Has(7)
	u.HasScheduled = u.Flags.Has(12)
	u.VideoCallsAvailable = u.Flags.Has(13)
	{
		value, err := DecodeUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode userFull#139a9a77: field user: %w", err)
		}
		u.User = value
	}
	if u.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#139a9a77: field about: %w", err)
		}
		u.About = value
	}
	{
		if err := u.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFull#139a9a77: field settings: %w", err)
		}
	}
	if u.Flags.Has(2) {
		value, err := DecodePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode userFull#139a9a77: field profile_photo: %w", err)
		}
		u.ProfilePhoto = value
	}
	{
		if err := u.NotifySettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFull#139a9a77: field notify_settings: %w", err)
		}
	}
	if u.Flags.Has(3) {
		if err := u.BotInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFull#139a9a77: field bot_info: %w", err)
		}
	}
	if u.Flags.Has(6) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#139a9a77: field pinned_msg_id: %w", err)
		}
		u.PinnedMsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#139a9a77: field common_chats_count: %w", err)
		}
		u.CommonChatsCount = value
	}
	if u.Flags.Has(11) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#139a9a77: field folder_id: %w", err)
		}
		u.FolderID = value
	}
	if u.Flags.Has(14) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userFull#139a9a77: field ttl_period: %w", err)
		}
		u.TTLPeriod = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UserFull.
var (
	_ bin.Encoder     = &UserFull{}
	_ bin.Decoder     = &UserFull{}
	_ bin.BareEncoder = &UserFull{}
	_ bin.BareDecoder = &UserFull{}
)
