package message

import (
	"context"
	"time"

	"github.com/gotd/td/tg"
)

// DocumentBuilder is a Document media option.
type DocumentBuilder struct {
	doc     tg.InputMediaDocument
	caption []StyledTextOption
}

// TTL sets time to live of self-destructing document.
func (u *DocumentBuilder) TTL(ttl time.Duration) *DocumentBuilder {
	return u.TTLSeconds(int(ttl.Seconds()))
}

// TTLSeconds sets time to live in seconds of self-destructing document.
func (u *DocumentBuilder) TTLSeconds(ttl int) *DocumentBuilder {
	u.doc.TTLSeconds = ttl
	return u
}

// Query sets query field of InputMediaDocument.
func (u *DocumentBuilder) Query(query string) *DocumentBuilder {
	u.doc.Query = query
	return u
}

// apply implements MediaOption.
func (u *DocumentBuilder) apply(ctx context.Context, b multiMediaBuilder) error {
	return Media(&u.doc, u.caption...).apply(ctx, b)
}

// Document adds doc attachment.
func Document(doc FileLocation, caption ...StyledTextOption) *DocumentBuilder {
	v := new(tg.InputDocument)
	v.FillFrom(doc)

	return &DocumentBuilder{
		doc: tg.InputMediaDocument{
			ID: v,
		},
		caption: caption,
	}
}

// DocumentExternalBuilder is a DocumentExternal media option.
type DocumentExternalBuilder struct {
	doc     tg.InputMediaDocumentExternal
	caption []StyledTextOption
}

// TTL sets time to live of self-destructing document.
func (u *DocumentExternalBuilder) TTL(ttl time.Duration) *DocumentExternalBuilder {
	return u.TTLSeconds(int(ttl.Seconds()))
}

// TTLSeconds sets time to live in seconds of self-destructing document.
func (u *DocumentExternalBuilder) TTLSeconds(ttl int) *DocumentExternalBuilder {
	u.doc.TTLSeconds = ttl
	return u
}

// apply implements MediaOption.
func (u *DocumentExternalBuilder) apply(ctx context.Context, b multiMediaBuilder) error {
	return Media(&u.doc, u.caption...).apply(ctx, b)
}

// DocumentExternal adds document attachment that will be downloaded by the Telegram servers.
func DocumentExternal(url string, caption ...StyledTextOption) *DocumentExternalBuilder {
	return &DocumentExternalBuilder{
		doc: tg.InputMediaDocumentExternal{
			URL: url,
		},
		caption: caption,
	}
}

// UploadedDocumentBuilder is a UploadedDocument media option.
type UploadedDocumentBuilder struct {
	doc     tg.InputMediaUploadedDocument
	caption []StyledTextOption
}

// NosoundVideo sets flag that the specified document is a video file with no audio tracks
// (a GIF animation (even as MPEG4), for example).
func (u *UploadedDocumentBuilder) NosoundVideo(v bool) *UploadedDocumentBuilder {
	u.doc.NosoundVideo = v
	return u
}

// ForceFile sets flag to force the media file to be uploaded as document.
func (u *UploadedDocumentBuilder) ForceFile(v bool) *UploadedDocumentBuilder {
	u.doc.ForceFile = v
	return u
}

// Thumb sets thumbnail of the document, uploaded as for the file.
func (u *UploadedDocumentBuilder) Thumb(file tg.InputFileClass) *UploadedDocumentBuilder {
	u.doc.Thumb = file
	return u
}

// MIME sets MIME type of document.
func (u *UploadedDocumentBuilder) MIME(mime string) *UploadedDocumentBuilder {
	u.doc.MimeType = mime
	return u
}

// Attributes adds given attributes to the document.
// Attribute specify the type of the document (video, audio, voice, sticker, etc.).
func (u *UploadedDocumentBuilder) Attributes(attrs ...tg.DocumentAttributeClass) *UploadedDocumentBuilder {
	u.doc.Attributes = append(u.doc.Attributes, attrs...)
	return u
}

// Stickers adds attached mask stickers.
func (u *UploadedDocumentBuilder) Stickers(stickers ...FileLocation) *UploadedDocumentBuilder {
	u.doc.Stickers = append(u.doc.Stickers, inputDocuments(stickers...)...)
	return u
}

// TTL sets time to live of self-destructing document.
func (u *UploadedDocumentBuilder) TTL(ttl time.Duration) *UploadedDocumentBuilder {
	return u.TTLSeconds(int(ttl.Seconds()))
}

// TTLSeconds sets time to live in seconds of self-destructing document.
func (u *UploadedDocumentBuilder) TTLSeconds(ttl int) *UploadedDocumentBuilder {
	u.doc.TTLSeconds = ttl
	return u
}

// apply implements MediaOption.
func (u *UploadedDocumentBuilder) apply(ctx context.Context, b multiMediaBuilder) error {
	return Media(&u.doc, u.caption...).apply(ctx, b)
}

// UploadedDocument adds document attachment.
func UploadedDocument(file tg.InputFileClass, caption ...StyledTextOption) *UploadedDocumentBuilder {
	return &UploadedDocumentBuilder{
		doc: tg.InputMediaUploadedDocument{
			File: file,
		},
		caption: caption,
	}
}

// File adds document attachment and forces it to be used as plain file, not media.
func File(file tg.InputFileClass, caption ...StyledTextOption) *UploadedDocumentBuilder {
	return UploadedDocument(file, caption...).ForceFile(true)
}
