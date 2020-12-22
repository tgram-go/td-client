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

// DialogFilterSuggested represents TL type `dialogFilterSuggested#77744d4a`.
// Suggested folders¹
//
// Links:
//  1) https://core.telegram.org/api/folders
//
// See https://core.telegram.org/constructor/dialogFilterSuggested for reference.
type DialogFilterSuggested struct {
	// Folder info¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	Filter DialogFilter
	// Folder¹ description
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	Description string
}

// DialogFilterSuggestedTypeID is TL type id of DialogFilterSuggested.
const DialogFilterSuggestedTypeID = 0x77744d4a

// Encode implements bin.Encoder.
func (d *DialogFilterSuggested) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dialogFilterSuggested#77744d4a as nil")
	}
	b.PutID(DialogFilterSuggestedTypeID)
	if err := d.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode dialogFilterSuggested#77744d4a: field filter: %w", err)
	}
	b.PutString(d.Description)
	return nil
}

// Decode implements bin.Decoder.
func (d *DialogFilterSuggested) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dialogFilterSuggested#77744d4a to nil")
	}
	if err := b.ConsumeID(DialogFilterSuggestedTypeID); err != nil {
		return fmt.Errorf("unable to decode dialogFilterSuggested#77744d4a: %w", err)
	}
	{
		if err := d.Filter.Decode(b); err != nil {
			return fmt.Errorf("unable to decode dialogFilterSuggested#77744d4a: field filter: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilterSuggested#77744d4a: field description: %w", err)
		}
		d.Description = value
	}
	return nil
}

// Ensuring interfaces in compile-time for DialogFilterSuggested.
var (
	_ bin.Encoder = &DialogFilterSuggested{}
	_ bin.Decoder = &DialogFilterSuggested{}
)
