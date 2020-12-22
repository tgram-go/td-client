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

// KeyboardButton represents TL type `keyboardButton#a2fa4880`.
// Bot keyboard button
//
// See https://core.telegram.org/constructor/keyboardButton for reference.
type KeyboardButton struct {
	// Button text
	Text string
}

// KeyboardButtonTypeID is TL type id of KeyboardButton.
const KeyboardButtonTypeID = 0xa2fa4880

// Encode implements bin.Encoder.
func (k *KeyboardButton) Encode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButton#a2fa4880 as nil")
	}
	b.PutID(KeyboardButtonTypeID)
	b.PutString(k.Text)
	return nil
}

// Decode implements bin.Decoder.
func (k *KeyboardButton) Decode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButton#a2fa4880 to nil")
	}
	if err := b.ConsumeID(KeyboardButtonTypeID); err != nil {
		return fmt.Errorf("unable to decode keyboardButton#a2fa4880: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButton#a2fa4880: field text: %w", err)
		}
		k.Text = value
	}
	return nil
}

// construct implements constructor of KeyboardButtonClass.
func (k KeyboardButton) construct() KeyboardButtonClass { return &k }

// Ensuring interfaces in compile-time for KeyboardButton.
var (
	_ bin.Encoder = &KeyboardButton{}
	_ bin.Decoder = &KeyboardButton{}

	_ KeyboardButtonClass = &KeyboardButton{}
)

// KeyboardButtonUrl represents TL type `keyboardButtonUrl#258aff05`.
// URL button
//
// See https://core.telegram.org/constructor/keyboardButtonUrl for reference.
type KeyboardButtonUrl struct {
	// Button label
	Text string
	// URL
	URL string
}

// KeyboardButtonUrlTypeID is TL type id of KeyboardButtonUrl.
const KeyboardButtonUrlTypeID = 0x258aff05

// Encode implements bin.Encoder.
func (k *KeyboardButtonUrl) Encode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonUrl#258aff05 as nil")
	}
	b.PutID(KeyboardButtonUrlTypeID)
	b.PutString(k.Text)
	b.PutString(k.URL)
	return nil
}

// Decode implements bin.Decoder.
func (k *KeyboardButtonUrl) Decode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonUrl#258aff05 to nil")
	}
	if err := b.ConsumeID(KeyboardButtonUrlTypeID); err != nil {
		return fmt.Errorf("unable to decode keyboardButtonUrl#258aff05: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonUrl#258aff05: field text: %w", err)
		}
		k.Text = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonUrl#258aff05: field url: %w", err)
		}
		k.URL = value
	}
	return nil
}

// construct implements constructor of KeyboardButtonClass.
func (k KeyboardButtonUrl) construct() KeyboardButtonClass { return &k }

// Ensuring interfaces in compile-time for KeyboardButtonUrl.
var (
	_ bin.Encoder = &KeyboardButtonUrl{}
	_ bin.Decoder = &KeyboardButtonUrl{}

	_ KeyboardButtonClass = &KeyboardButtonUrl{}
)

// KeyboardButtonCallback represents TL type `keyboardButtonCallback#35bbdb6b`.
// Callback button
//
// See https://core.telegram.org/constructor/keyboardButtonCallback for reference.
type KeyboardButtonCallback struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the user should verify his identity by entering his 2FA SRP parameters¹ to the messages.getBotCallbackAnswer² method. NOTE: telegram and the bot WILL NOT have access to the plaintext password, thanks to SRP³. This button is mainly used by the official @botfather⁴ bot, for verifying the user's identity before transferring ownership of a bot to another user.
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	//  2) https://core.telegram.org/method/messages.getBotCallbackAnswer
	//  3) https://core.telegram.org/api/srp
	//  4) https://t.me/botfather
	RequiresPassword bool
	// Button text
	Text string
	// Callback data
	Data []byte
}

// KeyboardButtonCallbackTypeID is TL type id of KeyboardButtonCallback.
const KeyboardButtonCallbackTypeID = 0x35bbdb6b

// Encode implements bin.Encoder.
func (k *KeyboardButtonCallback) Encode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonCallback#35bbdb6b as nil")
	}
	b.PutID(KeyboardButtonCallbackTypeID)
	if err := k.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode keyboardButtonCallback#35bbdb6b: field flags: %w", err)
	}
	b.PutString(k.Text)
	b.PutBytes(k.Data)
	return nil
}

// SetRequiresPassword sets value of RequiresPassword conditional field.
func (k *KeyboardButtonCallback) SetRequiresPassword(value bool) {
	if value {
		k.Flags.Set(0)
	} else {
		k.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (k *KeyboardButtonCallback) Decode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonCallback#35bbdb6b to nil")
	}
	if err := b.ConsumeID(KeyboardButtonCallbackTypeID); err != nil {
		return fmt.Errorf("unable to decode keyboardButtonCallback#35bbdb6b: %w", err)
	}
	{
		if err := k.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode keyboardButtonCallback#35bbdb6b: field flags: %w", err)
		}
	}
	k.RequiresPassword = k.Flags.Has(0)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonCallback#35bbdb6b: field text: %w", err)
		}
		k.Text = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonCallback#35bbdb6b: field data: %w", err)
		}
		k.Data = value
	}
	return nil
}

// construct implements constructor of KeyboardButtonClass.
func (k KeyboardButtonCallback) construct() KeyboardButtonClass { return &k }

// Ensuring interfaces in compile-time for KeyboardButtonCallback.
var (
	_ bin.Encoder = &KeyboardButtonCallback{}
	_ bin.Decoder = &KeyboardButtonCallback{}

	_ KeyboardButtonClass = &KeyboardButtonCallback{}
)

// KeyboardButtonRequestPhone represents TL type `keyboardButtonRequestPhone#b16a6c29`.
// Button to request a user's phone number
//
// See https://core.telegram.org/constructor/keyboardButtonRequestPhone for reference.
type KeyboardButtonRequestPhone struct {
	// Button text
	Text string
}

// KeyboardButtonRequestPhoneTypeID is TL type id of KeyboardButtonRequestPhone.
const KeyboardButtonRequestPhoneTypeID = 0xb16a6c29

// Encode implements bin.Encoder.
func (k *KeyboardButtonRequestPhone) Encode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonRequestPhone#b16a6c29 as nil")
	}
	b.PutID(KeyboardButtonRequestPhoneTypeID)
	b.PutString(k.Text)
	return nil
}

// Decode implements bin.Decoder.
func (k *KeyboardButtonRequestPhone) Decode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonRequestPhone#b16a6c29 to nil")
	}
	if err := b.ConsumeID(KeyboardButtonRequestPhoneTypeID); err != nil {
		return fmt.Errorf("unable to decode keyboardButtonRequestPhone#b16a6c29: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonRequestPhone#b16a6c29: field text: %w", err)
		}
		k.Text = value
	}
	return nil
}

// construct implements constructor of KeyboardButtonClass.
func (k KeyboardButtonRequestPhone) construct() KeyboardButtonClass { return &k }

// Ensuring interfaces in compile-time for KeyboardButtonRequestPhone.
var (
	_ bin.Encoder = &KeyboardButtonRequestPhone{}
	_ bin.Decoder = &KeyboardButtonRequestPhone{}

	_ KeyboardButtonClass = &KeyboardButtonRequestPhone{}
)

// KeyboardButtonRequestGeoLocation represents TL type `keyboardButtonRequestGeoLocation#fc796b3f`.
// Button to request a user's geolocation
//
// See https://core.telegram.org/constructor/keyboardButtonRequestGeoLocation for reference.
type KeyboardButtonRequestGeoLocation struct {
	// Button text
	Text string
}

// KeyboardButtonRequestGeoLocationTypeID is TL type id of KeyboardButtonRequestGeoLocation.
const KeyboardButtonRequestGeoLocationTypeID = 0xfc796b3f

// Encode implements bin.Encoder.
func (k *KeyboardButtonRequestGeoLocation) Encode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonRequestGeoLocation#fc796b3f as nil")
	}
	b.PutID(KeyboardButtonRequestGeoLocationTypeID)
	b.PutString(k.Text)
	return nil
}

// Decode implements bin.Decoder.
func (k *KeyboardButtonRequestGeoLocation) Decode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonRequestGeoLocation#fc796b3f to nil")
	}
	if err := b.ConsumeID(KeyboardButtonRequestGeoLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode keyboardButtonRequestGeoLocation#fc796b3f: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonRequestGeoLocation#fc796b3f: field text: %w", err)
		}
		k.Text = value
	}
	return nil
}

// construct implements constructor of KeyboardButtonClass.
func (k KeyboardButtonRequestGeoLocation) construct() KeyboardButtonClass { return &k }

// Ensuring interfaces in compile-time for KeyboardButtonRequestGeoLocation.
var (
	_ bin.Encoder = &KeyboardButtonRequestGeoLocation{}
	_ bin.Decoder = &KeyboardButtonRequestGeoLocation{}

	_ KeyboardButtonClass = &KeyboardButtonRequestGeoLocation{}
)

// KeyboardButtonSwitchInline represents TL type `keyboardButtonSwitchInline#568a748`.
// Button to force a user to switch to inline mode Pressing the button will prompt the user to select one of their chats, open that chat and insert the bot‘s username and the specified inline query in the input field.
//
// See https://core.telegram.org/constructor/keyboardButtonSwitchInline for reference.
type KeyboardButtonSwitchInline struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, pressing the button will insert the bot‘s username and the specified inline query in the current chat's input field.
	SamePeer bool
	// Button label
	Text string
	// The inline query to use
	Query string
}

// KeyboardButtonSwitchInlineTypeID is TL type id of KeyboardButtonSwitchInline.
const KeyboardButtonSwitchInlineTypeID = 0x568a748

// Encode implements bin.Encoder.
func (k *KeyboardButtonSwitchInline) Encode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonSwitchInline#568a748 as nil")
	}
	b.PutID(KeyboardButtonSwitchInlineTypeID)
	if err := k.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode keyboardButtonSwitchInline#568a748: field flags: %w", err)
	}
	b.PutString(k.Text)
	b.PutString(k.Query)
	return nil
}

// SetSamePeer sets value of SamePeer conditional field.
func (k *KeyboardButtonSwitchInline) SetSamePeer(value bool) {
	if value {
		k.Flags.Set(0)
	} else {
		k.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (k *KeyboardButtonSwitchInline) Decode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonSwitchInline#568a748 to nil")
	}
	if err := b.ConsumeID(KeyboardButtonSwitchInlineTypeID); err != nil {
		return fmt.Errorf("unable to decode keyboardButtonSwitchInline#568a748: %w", err)
	}
	{
		if err := k.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode keyboardButtonSwitchInline#568a748: field flags: %w", err)
		}
	}
	k.SamePeer = k.Flags.Has(0)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonSwitchInline#568a748: field text: %w", err)
		}
		k.Text = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonSwitchInline#568a748: field query: %w", err)
		}
		k.Query = value
	}
	return nil
}

// construct implements constructor of KeyboardButtonClass.
func (k KeyboardButtonSwitchInline) construct() KeyboardButtonClass { return &k }

// Ensuring interfaces in compile-time for KeyboardButtonSwitchInline.
var (
	_ bin.Encoder = &KeyboardButtonSwitchInline{}
	_ bin.Decoder = &KeyboardButtonSwitchInline{}

	_ KeyboardButtonClass = &KeyboardButtonSwitchInline{}
)

// KeyboardButtonGame represents TL type `keyboardButtonGame#50f41ccf`.
// Button to start a game
//
// See https://core.telegram.org/constructor/keyboardButtonGame for reference.
type KeyboardButtonGame struct {
	// Button text
	Text string
}

// KeyboardButtonGameTypeID is TL type id of KeyboardButtonGame.
const KeyboardButtonGameTypeID = 0x50f41ccf

// Encode implements bin.Encoder.
func (k *KeyboardButtonGame) Encode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonGame#50f41ccf as nil")
	}
	b.PutID(KeyboardButtonGameTypeID)
	b.PutString(k.Text)
	return nil
}

// Decode implements bin.Decoder.
func (k *KeyboardButtonGame) Decode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonGame#50f41ccf to nil")
	}
	if err := b.ConsumeID(KeyboardButtonGameTypeID); err != nil {
		return fmt.Errorf("unable to decode keyboardButtonGame#50f41ccf: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonGame#50f41ccf: field text: %w", err)
		}
		k.Text = value
	}
	return nil
}

// construct implements constructor of KeyboardButtonClass.
func (k KeyboardButtonGame) construct() KeyboardButtonClass { return &k }

// Ensuring interfaces in compile-time for KeyboardButtonGame.
var (
	_ bin.Encoder = &KeyboardButtonGame{}
	_ bin.Decoder = &KeyboardButtonGame{}

	_ KeyboardButtonClass = &KeyboardButtonGame{}
)

// KeyboardButtonBuy represents TL type `keyboardButtonBuy#afd93fbb`.
// Button to buy a product
//
// See https://core.telegram.org/constructor/keyboardButtonBuy for reference.
type KeyboardButtonBuy struct {
	// Button text
	Text string
}

// KeyboardButtonBuyTypeID is TL type id of KeyboardButtonBuy.
const KeyboardButtonBuyTypeID = 0xafd93fbb

// Encode implements bin.Encoder.
func (k *KeyboardButtonBuy) Encode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonBuy#afd93fbb as nil")
	}
	b.PutID(KeyboardButtonBuyTypeID)
	b.PutString(k.Text)
	return nil
}

// Decode implements bin.Decoder.
func (k *KeyboardButtonBuy) Decode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonBuy#afd93fbb to nil")
	}
	if err := b.ConsumeID(KeyboardButtonBuyTypeID); err != nil {
		return fmt.Errorf("unable to decode keyboardButtonBuy#afd93fbb: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonBuy#afd93fbb: field text: %w", err)
		}
		k.Text = value
	}
	return nil
}

// construct implements constructor of KeyboardButtonClass.
func (k KeyboardButtonBuy) construct() KeyboardButtonClass { return &k }

// Ensuring interfaces in compile-time for KeyboardButtonBuy.
var (
	_ bin.Encoder = &KeyboardButtonBuy{}
	_ bin.Decoder = &KeyboardButtonBuy{}

	_ KeyboardButtonClass = &KeyboardButtonBuy{}
)

// KeyboardButtonUrlAuth represents TL type `keyboardButtonUrlAuth#10b78d29`.
// Button to request a user to authorize via URL using Seamless Telegram Login¹. When the user clicks on such a button, messages.requestUrlAuth² should be called, providing the button_id and the ID of the container message. The returned urlAuthResultRequest³ object will contain more details about the authorization request (request_write_access if the bot would like to send messages to the user along with the username of the bot which will be used for user authorization). Finally, the user can choose to call messages.acceptUrlAuth⁴ to get a urlAuthResultAccepted⁵ with the URL to open instead of the url of this constructor, or a urlAuthResultDefault⁶, in which case the url of this constructor must be opened, instead. If the user refuses the authorization request but still wants to open the link, the url of this constructor must be used.
//
// Links:
//  1) https://telegram.org/blog/privacy-discussions-web-bots#meet-seamless-web-bots
//  2) https://core.telegram.org/method/messages.requestUrlAuth
//  3) https://core.telegram.org/constructor/urlAuthResultRequest
//  4) https://core.telegram.org/method/messages.acceptUrlAuth
//  5) https://core.telegram.org/constructor/urlAuthResultAccepted
//  6) https://core.telegram.org/constructor/urlAuthResultDefault
//
// See https://core.telegram.org/constructor/keyboardButtonUrlAuth for reference.
type KeyboardButtonUrlAuth struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Button label
	Text string
	// New text of the button in forwarded messages.
	//
	// Use SetFwdText and GetFwdText helpers.
	FwdText string
	// An HTTP URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in Receiving authorization data¹.NOTE: Services must always check the hash of the received data to verify the authentication and the integrity of the data as described in Checking authorization².
	//
	// Links:
	//  1) https://core.telegram.org/widgets/login#receiving-authorization-data
	//  2) https://core.telegram.org/widgets/login#checking-authorization
	URL string
	// ID of the button to pass to messages.requestUrlAuth¹
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.requestUrlAuth
	ButtonID int
}

// KeyboardButtonUrlAuthTypeID is TL type id of KeyboardButtonUrlAuth.
const KeyboardButtonUrlAuthTypeID = 0x10b78d29

// Encode implements bin.Encoder.
func (k *KeyboardButtonUrlAuth) Encode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonUrlAuth#10b78d29 as nil")
	}
	b.PutID(KeyboardButtonUrlAuthTypeID)
	if err := k.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode keyboardButtonUrlAuth#10b78d29: field flags: %w", err)
	}
	b.PutString(k.Text)
	if k.Flags.Has(0) {
		b.PutString(k.FwdText)
	}
	b.PutString(k.URL)
	b.PutInt(k.ButtonID)
	return nil
}

// SetFwdText sets value of FwdText conditional field.
func (k *KeyboardButtonUrlAuth) SetFwdText(value string) {
	k.Flags.Set(0)
	k.FwdText = value
}

// GetFwdText returns value of FwdText conditional field and
// boolean which is true if field was set.
func (k *KeyboardButtonUrlAuth) GetFwdText() (value string, ok bool) {
	if !k.Flags.Has(0) {
		return value, false
	}
	return k.FwdText, true
}

// Decode implements bin.Decoder.
func (k *KeyboardButtonUrlAuth) Decode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonUrlAuth#10b78d29 to nil")
	}
	if err := b.ConsumeID(KeyboardButtonUrlAuthTypeID); err != nil {
		return fmt.Errorf("unable to decode keyboardButtonUrlAuth#10b78d29: %w", err)
	}
	{
		if err := k.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode keyboardButtonUrlAuth#10b78d29: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonUrlAuth#10b78d29: field text: %w", err)
		}
		k.Text = value
	}
	if k.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonUrlAuth#10b78d29: field fwd_text: %w", err)
		}
		k.FwdText = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonUrlAuth#10b78d29: field url: %w", err)
		}
		k.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonUrlAuth#10b78d29: field button_id: %w", err)
		}
		k.ButtonID = value
	}
	return nil
}

// construct implements constructor of KeyboardButtonClass.
func (k KeyboardButtonUrlAuth) construct() KeyboardButtonClass { return &k }

// Ensuring interfaces in compile-time for KeyboardButtonUrlAuth.
var (
	_ bin.Encoder = &KeyboardButtonUrlAuth{}
	_ bin.Decoder = &KeyboardButtonUrlAuth{}

	_ KeyboardButtonClass = &KeyboardButtonUrlAuth{}
)

// InputKeyboardButtonUrlAuth represents TL type `inputKeyboardButtonUrlAuth#d02e7fd4`.
// Button to request a user to authorize¹ via URL using Seamless Telegram Login².
//
// Links:
//  1) https://core.telegram.org/method/messages.acceptUrlAuth
//  2) https://telegram.org/blog/privacy-discussions-web-bots#meet-seamless-web-bots
//
// See https://core.telegram.org/constructor/inputKeyboardButtonUrlAuth for reference.
type InputKeyboardButtonUrlAuth struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Set this flag to request the permission for your bot to send messages to the user.
	RequestWriteAccess bool
	// Button text
	Text string
	// New text of the button in forwarded messages.
	//
	// Use SetFwdText and GetFwdText helpers.
	FwdText string
	// An HTTP URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in Receiving authorization data¹.NOTE: You must always check the hash of the received data to verify the authentication and the integrity of the data as described in Checking authorization².
	//
	// Links:
	//  1) https://core.telegram.org/widgets/login#receiving-authorization-data
	//  2) https://core.telegram.org/widgets/login#checking-authorization
	URL string
	// Username of a bot, which will be used for user authorization. See Setting up a bot¹ for more details. If not specified, the current bot's username will be assumed. The url's domain must be the same as the domain linked with the bot. See Linking your domain to the bot² for more details.
	//
	// Links:
	//  1) https://core.telegram.org/widgets/login#setting-up-a-bot
	//  2) https://core.telegram.org/widgets/login#linking-your-domain-to-the-bot
	Bot InputUserClass
}

// InputKeyboardButtonUrlAuthTypeID is TL type id of InputKeyboardButtonUrlAuth.
const InputKeyboardButtonUrlAuthTypeID = 0xd02e7fd4

// Encode implements bin.Encoder.
func (i *InputKeyboardButtonUrlAuth) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputKeyboardButtonUrlAuth#d02e7fd4 as nil")
	}
	b.PutID(InputKeyboardButtonUrlAuthTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputKeyboardButtonUrlAuth#d02e7fd4: field flags: %w", err)
	}
	b.PutString(i.Text)
	if i.Flags.Has(1) {
		b.PutString(i.FwdText)
	}
	b.PutString(i.URL)
	if i.Bot == nil {
		return fmt.Errorf("unable to encode inputKeyboardButtonUrlAuth#d02e7fd4: field bot is nil")
	}
	if err := i.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputKeyboardButtonUrlAuth#d02e7fd4: field bot: %w", err)
	}
	return nil
}

// SetRequestWriteAccess sets value of RequestWriteAccess conditional field.
func (i *InputKeyboardButtonUrlAuth) SetRequestWriteAccess(value bool) {
	if value {
		i.Flags.Set(0)
	} else {
		i.Flags.Unset(0)
	}
}

// SetFwdText sets value of FwdText conditional field.
func (i *InputKeyboardButtonUrlAuth) SetFwdText(value string) {
	i.Flags.Set(1)
	i.FwdText = value
}

// GetFwdText returns value of FwdText conditional field and
// boolean which is true if field was set.
func (i *InputKeyboardButtonUrlAuth) GetFwdText() (value string, ok bool) {
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.FwdText, true
}

// Decode implements bin.Decoder.
func (i *InputKeyboardButtonUrlAuth) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputKeyboardButtonUrlAuth#d02e7fd4 to nil")
	}
	if err := b.ConsumeID(InputKeyboardButtonUrlAuthTypeID); err != nil {
		return fmt.Errorf("unable to decode inputKeyboardButtonUrlAuth#d02e7fd4: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputKeyboardButtonUrlAuth#d02e7fd4: field flags: %w", err)
		}
	}
	i.RequestWriteAccess = i.Flags.Has(0)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputKeyboardButtonUrlAuth#d02e7fd4: field text: %w", err)
		}
		i.Text = value
	}
	if i.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputKeyboardButtonUrlAuth#d02e7fd4: field fwd_text: %w", err)
		}
		i.FwdText = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputKeyboardButtonUrlAuth#d02e7fd4: field url: %w", err)
		}
		i.URL = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputKeyboardButtonUrlAuth#d02e7fd4: field bot: %w", err)
		}
		i.Bot = value
	}
	return nil
}

// construct implements constructor of KeyboardButtonClass.
func (i InputKeyboardButtonUrlAuth) construct() KeyboardButtonClass { return &i }

// Ensuring interfaces in compile-time for InputKeyboardButtonUrlAuth.
var (
	_ bin.Encoder = &InputKeyboardButtonUrlAuth{}
	_ bin.Decoder = &InputKeyboardButtonUrlAuth{}

	_ KeyboardButtonClass = &InputKeyboardButtonUrlAuth{}
)

// KeyboardButtonRequestPoll represents TL type `keyboardButtonRequestPoll#bbc7515d`.
// A button that allows the user to create and send a poll when pressed; available only in private
//
// See https://core.telegram.org/constructor/keyboardButtonRequestPoll for reference.
type KeyboardButtonRequestPoll struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, only quiz polls can be sent
	Quiz bool
	// Button text
	Text string
}

// KeyboardButtonRequestPollTypeID is TL type id of KeyboardButtonRequestPoll.
const KeyboardButtonRequestPollTypeID = 0xbbc7515d

// Encode implements bin.Encoder.
func (k *KeyboardButtonRequestPoll) Encode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonRequestPoll#bbc7515d as nil")
	}
	b.PutID(KeyboardButtonRequestPollTypeID)
	if err := k.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode keyboardButtonRequestPoll#bbc7515d: field flags: %w", err)
	}
	b.PutString(k.Text)
	return nil
}

// SetQuiz sets value of Quiz conditional field.
func (k *KeyboardButtonRequestPoll) SetQuiz(value bool) {
	if value {
		k.Flags.Set(0)
	} else {
		k.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (k *KeyboardButtonRequestPoll) Decode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonRequestPoll#bbc7515d to nil")
	}
	if err := b.ConsumeID(KeyboardButtonRequestPollTypeID); err != nil {
		return fmt.Errorf("unable to decode keyboardButtonRequestPoll#bbc7515d: %w", err)
	}
	{
		if err := k.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode keyboardButtonRequestPoll#bbc7515d: field flags: %w", err)
		}
	}
	k.Quiz = k.Flags.Has(0)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonRequestPoll#bbc7515d: field text: %w", err)
		}
		k.Text = value
	}
	return nil
}

// construct implements constructor of KeyboardButtonClass.
func (k KeyboardButtonRequestPoll) construct() KeyboardButtonClass { return &k }

// Ensuring interfaces in compile-time for KeyboardButtonRequestPoll.
var (
	_ bin.Encoder = &KeyboardButtonRequestPoll{}
	_ bin.Decoder = &KeyboardButtonRequestPoll{}

	_ KeyboardButtonClass = &KeyboardButtonRequestPoll{}
)

// KeyboardButtonClass represents KeyboardButton generic type.
//
// See https://core.telegram.org/type/KeyboardButton for reference.
//
// Example:
//  g, err := DecodeKeyboardButton(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *KeyboardButton: // keyboardButton#a2fa4880
//  case *KeyboardButtonUrl: // keyboardButtonUrl#258aff05
//  case *KeyboardButtonCallback: // keyboardButtonCallback#35bbdb6b
//  case *KeyboardButtonRequestPhone: // keyboardButtonRequestPhone#b16a6c29
//  case *KeyboardButtonRequestGeoLocation: // keyboardButtonRequestGeoLocation#fc796b3f
//  case *KeyboardButtonSwitchInline: // keyboardButtonSwitchInline#568a748
//  case *KeyboardButtonGame: // keyboardButtonGame#50f41ccf
//  case *KeyboardButtonBuy: // keyboardButtonBuy#afd93fbb
//  case *KeyboardButtonUrlAuth: // keyboardButtonUrlAuth#10b78d29
//  case *InputKeyboardButtonUrlAuth: // inputKeyboardButtonUrlAuth#d02e7fd4
//  case *KeyboardButtonRequestPoll: // keyboardButtonRequestPoll#bbc7515d
//  default: panic(v)
//  }
type KeyboardButtonClass interface {
	bin.Encoder
	bin.Decoder
	construct() KeyboardButtonClass
}

// DecodeKeyboardButton implements binary de-serialization for KeyboardButtonClass.
func DecodeKeyboardButton(buf *bin.Buffer) (KeyboardButtonClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case KeyboardButtonTypeID:
		// Decoding keyboardButton#a2fa4880.
		v := KeyboardButton{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonClass: %w", err)
		}
		return &v, nil
	case KeyboardButtonUrlTypeID:
		// Decoding keyboardButtonUrl#258aff05.
		v := KeyboardButtonUrl{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonClass: %w", err)
		}
		return &v, nil
	case KeyboardButtonCallbackTypeID:
		// Decoding keyboardButtonCallback#35bbdb6b.
		v := KeyboardButtonCallback{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonClass: %w", err)
		}
		return &v, nil
	case KeyboardButtonRequestPhoneTypeID:
		// Decoding keyboardButtonRequestPhone#b16a6c29.
		v := KeyboardButtonRequestPhone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonClass: %w", err)
		}
		return &v, nil
	case KeyboardButtonRequestGeoLocationTypeID:
		// Decoding keyboardButtonRequestGeoLocation#fc796b3f.
		v := KeyboardButtonRequestGeoLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonClass: %w", err)
		}
		return &v, nil
	case KeyboardButtonSwitchInlineTypeID:
		// Decoding keyboardButtonSwitchInline#568a748.
		v := KeyboardButtonSwitchInline{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonClass: %w", err)
		}
		return &v, nil
	case KeyboardButtonGameTypeID:
		// Decoding keyboardButtonGame#50f41ccf.
		v := KeyboardButtonGame{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonClass: %w", err)
		}
		return &v, nil
	case KeyboardButtonBuyTypeID:
		// Decoding keyboardButtonBuy#afd93fbb.
		v := KeyboardButtonBuy{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonClass: %w", err)
		}
		return &v, nil
	case KeyboardButtonUrlAuthTypeID:
		// Decoding keyboardButtonUrlAuth#10b78d29.
		v := KeyboardButtonUrlAuth{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonClass: %w", err)
		}
		return &v, nil
	case InputKeyboardButtonUrlAuthTypeID:
		// Decoding inputKeyboardButtonUrlAuth#d02e7fd4.
		v := InputKeyboardButtonUrlAuth{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonClass: %w", err)
		}
		return &v, nil
	case KeyboardButtonRequestPollTypeID:
		// Decoding keyboardButtonRequestPoll#bbc7515d.
		v := KeyboardButtonRequestPoll{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode KeyboardButtonClass: %w", bin.NewUnexpectedID(id))
	}
}

// KeyboardButton boxes the KeyboardButtonClass providing a helper.
type KeyboardButtonBox struct {
	KeyboardButton KeyboardButtonClass
}

// Decode implements bin.Decoder for KeyboardButtonBox.
func (b *KeyboardButtonBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode KeyboardButtonBox to nil")
	}
	v, err := DecodeKeyboardButton(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.KeyboardButton = v
	return nil
}

// Encode implements bin.Encode for KeyboardButtonBox.
func (b *KeyboardButtonBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.KeyboardButton == nil {
		return fmt.Errorf("unable to encode KeyboardButtonClass as nil")
	}
	return b.KeyboardButton.Encode(buf)
}
