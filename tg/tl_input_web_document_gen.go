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

// InputWebDocument represents TL type `inputWebDocument#9bed434d`.
// The document
//
// See https://core.telegram.org/constructor/inputWebDocument for reference.
type InputWebDocument struct {
	// Remote document URL to be downloaded using the appropriate method¹
	//
	// Links:
	//  1) https://core.telegram.org/api/files
	URL string
	// Remote file size
	Size int
	// Mime type
	MimeType string
	// Attributes for media types
	Attributes []DocumentAttributeClass
}

// InputWebDocumentTypeID is TL type id of InputWebDocument.
const InputWebDocumentTypeID = 0x9bed434d

// Encode implements bin.Encoder.
func (i *InputWebDocument) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputWebDocument#9bed434d as nil")
	}
	b.PutID(InputWebDocumentTypeID)
	b.PutString(i.URL)
	b.PutInt(i.Size)
	b.PutString(i.MimeType)
	b.PutVectorHeader(len(i.Attributes))
	for idx, v := range i.Attributes {
		if v == nil {
			return fmt.Errorf("unable to encode inputWebDocument#9bed434d: field attributes element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputWebDocument#9bed434d: field attributes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputWebDocument) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputWebDocument#9bed434d to nil")
	}
	if err := b.ConsumeID(InputWebDocumentTypeID); err != nil {
		return fmt.Errorf("unable to decode inputWebDocument#9bed434d: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebDocument#9bed434d: field url: %w", err)
		}
		i.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebDocument#9bed434d: field size: %w", err)
		}
		i.Size = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebDocument#9bed434d: field mime_type: %w", err)
		}
		i.MimeType = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebDocument#9bed434d: field attributes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocumentAttribute(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputWebDocument#9bed434d: field attributes: %w", err)
			}
			i.Attributes = append(i.Attributes, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for InputWebDocument.
var (
	_ bin.Encoder = &InputWebDocument{}
	_ bin.Decoder = &InputWebDocument{}
)
