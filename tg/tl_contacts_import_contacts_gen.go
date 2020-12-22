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

// ContactsImportContactsRequest represents TL type `contacts.importContacts#2c800be5`.
// Imports contacts: saves a full list on the server, adds already registered contacts to the contact list, returns added contacts and their info.
// Use contacts.addContact¹ to add Telegram contacts without actually using their phone number.
//
// Links:
//  1) https://core.telegram.org/method/contacts.addContact
//
// See https://core.telegram.org/method/contacts.importContacts for reference.
type ContactsImportContactsRequest struct {
	// List of contacts to import
	Contacts []InputPhoneContact
}

// ContactsImportContactsRequestTypeID is TL type id of ContactsImportContactsRequest.
const ContactsImportContactsRequestTypeID = 0x2c800be5

// Encode implements bin.Encoder.
func (i *ContactsImportContactsRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode contacts.importContacts#2c800be5 as nil")
	}
	b.PutID(ContactsImportContactsRequestTypeID)
	b.PutVectorHeader(len(i.Contacts))
	for idx, v := range i.Contacts {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.importContacts#2c800be5: field contacts element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *ContactsImportContactsRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode contacts.importContacts#2c800be5 to nil")
	}
	if err := b.ConsumeID(ContactsImportContactsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.importContacts#2c800be5: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.importContacts#2c800be5: field contacts: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value InputPhoneContact
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode contacts.importContacts#2c800be5: field contacts: %w", err)
			}
			i.Contacts = append(i.Contacts, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsImportContactsRequest.
var (
	_ bin.Encoder = &ContactsImportContactsRequest{}
	_ bin.Decoder = &ContactsImportContactsRequest{}
)

// ContactsImportContacts invokes method contacts.importContacts#2c800be5 returning error if any.
// Imports contacts: saves a full list on the server, adds already registered contacts to the contact list, returns added contacts and their info.
// Use contacts.addContact¹ to add Telegram contacts without actually using their phone number.
//
// Links:
//  1) https://core.telegram.org/method/contacts.addContact
//
// See https://core.telegram.org/method/contacts.importContacts for reference.
func (c *Client) ContactsImportContacts(ctx context.Context, contacts []InputPhoneContact) (*ContactsImportedContacts, error) {
	var result ContactsImportedContacts

	request := &ContactsImportContactsRequest{
		Contacts: contacts,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
