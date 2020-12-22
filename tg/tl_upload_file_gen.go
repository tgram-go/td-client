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

// UploadFile represents TL type `upload.file#96a18d5`.
// File content.
//
// See https://core.telegram.org/constructor/upload.file for reference.
type UploadFile struct {
	// File type
	Type StorageFileTypeClass
	// Modification type
	Mtime int
	// Binary data, file content
	Bytes []byte
}

// UploadFileTypeID is TL type id of UploadFile.
const UploadFileTypeID = 0x96a18d5

// Encode implements bin.Encoder.
func (f *UploadFile) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode upload.file#96a18d5 as nil")
	}
	b.PutID(UploadFileTypeID)
	if f.Type == nil {
		return fmt.Errorf("unable to encode upload.file#96a18d5: field type is nil")
	}
	if err := f.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode upload.file#96a18d5: field type: %w", err)
	}
	b.PutInt(f.Mtime)
	b.PutBytes(f.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (f *UploadFile) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode upload.file#96a18d5 to nil")
	}
	if err := b.ConsumeID(UploadFileTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.file#96a18d5: %w", err)
	}
	{
		value, err := DecodeStorageFileType(b)
		if err != nil {
			return fmt.Errorf("unable to decode upload.file#96a18d5: field type: %w", err)
		}
		f.Type = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.file#96a18d5: field mtime: %w", err)
		}
		f.Mtime = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.file#96a18d5: field bytes: %w", err)
		}
		f.Bytes = value
	}
	return nil
}

// construct implements constructor of UploadFileClass.
func (f UploadFile) construct() UploadFileClass { return &f }

// Ensuring interfaces in compile-time for UploadFile.
var (
	_ bin.Encoder = &UploadFile{}
	_ bin.Decoder = &UploadFile{}

	_ UploadFileClass = &UploadFile{}
)

// UploadFileCdnRedirect represents TL type `upload.fileCdnRedirect#f18cda44`.
// The file must be downloaded from a CDN DC¹.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/constructor/upload.fileCdnRedirect for reference.
type UploadFileCdnRedirect struct {
	// CDN DC¹ ID
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	DCID int
	// File token (see CDN files¹)
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	FileToken []byte
	// Encryption key (see CDN files¹)
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	EncryptionKey []byte
	// Encryption IV (see CDN files¹)
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	EncryptionIv []byte
	// File hashes (see CDN files¹)
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	FileHashes []FileHash
}

// UploadFileCdnRedirectTypeID is TL type id of UploadFileCdnRedirect.
const UploadFileCdnRedirectTypeID = 0xf18cda44

// Encode implements bin.Encoder.
func (f *UploadFileCdnRedirect) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode upload.fileCdnRedirect#f18cda44 as nil")
	}
	b.PutID(UploadFileCdnRedirectTypeID)
	b.PutInt(f.DCID)
	b.PutBytes(f.FileToken)
	b.PutBytes(f.EncryptionKey)
	b.PutBytes(f.EncryptionIv)
	b.PutVectorHeader(len(f.FileHashes))
	for idx, v := range f.FileHashes {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode upload.fileCdnRedirect#f18cda44: field file_hashes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *UploadFileCdnRedirect) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode upload.fileCdnRedirect#f18cda44 to nil")
	}
	if err := b.ConsumeID(UploadFileCdnRedirectTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.fileCdnRedirect#f18cda44: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.fileCdnRedirect#f18cda44: field dc_id: %w", err)
		}
		f.DCID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.fileCdnRedirect#f18cda44: field file_token: %w", err)
		}
		f.FileToken = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.fileCdnRedirect#f18cda44: field encryption_key: %w", err)
		}
		f.EncryptionKey = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.fileCdnRedirect#f18cda44: field encryption_iv: %w", err)
		}
		f.EncryptionIv = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode upload.fileCdnRedirect#f18cda44: field file_hashes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value FileHash
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode upload.fileCdnRedirect#f18cda44: field file_hashes: %w", err)
			}
			f.FileHashes = append(f.FileHashes, value)
		}
	}
	return nil
}

// construct implements constructor of UploadFileClass.
func (f UploadFileCdnRedirect) construct() UploadFileClass { return &f }

// Ensuring interfaces in compile-time for UploadFileCdnRedirect.
var (
	_ bin.Encoder = &UploadFileCdnRedirect{}
	_ bin.Decoder = &UploadFileCdnRedirect{}

	_ UploadFileClass = &UploadFileCdnRedirect{}
)

// UploadFileClass represents upload.File generic type.
//
// See https://core.telegram.org/type/upload.File for reference.
//
// Example:
//  g, err := DecodeUploadFile(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *UploadFile: // upload.file#96a18d5
//  case *UploadFileCdnRedirect: // upload.fileCdnRedirect#f18cda44
//  default: panic(v)
//  }
type UploadFileClass interface {
	bin.Encoder
	bin.Decoder
	construct() UploadFileClass
}

// DecodeUploadFile implements binary de-serialization for UploadFileClass.
func DecodeUploadFile(buf *bin.Buffer) (UploadFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UploadFileTypeID:
		// Decoding upload.file#96a18d5.
		v := UploadFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UploadFileClass: %w", err)
		}
		return &v, nil
	case UploadFileCdnRedirectTypeID:
		// Decoding upload.fileCdnRedirect#f18cda44.
		v := UploadFileCdnRedirect{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UploadFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UploadFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// UploadFile boxes the UploadFileClass providing a helper.
type UploadFileBox struct {
	File UploadFileClass
}

// Decode implements bin.Decoder for UploadFileBox.
func (b *UploadFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UploadFileBox to nil")
	}
	v, err := DecodeUploadFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.File = v
	return nil
}

// Encode implements bin.Encode for UploadFileBox.
func (b *UploadFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.File == nil {
		return fmt.Errorf("unable to encode UploadFileClass as nil")
	}
	return b.File.Encode(buf)
}
