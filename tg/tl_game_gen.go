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

// Game represents TL type `game#bdf9653b`.
// Indicates an already sent game
//
// See https://core.telegram.org/constructor/game for reference.
type Game struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// ID of the game
	ID int64
	// Access hash of the game
	AccessHash int64
	// Short name for the game
	ShortName string
	// Title of the game
	Title string
	// Game description
	Description string
	// Game preview
	Photo PhotoClass
	// Optional attached document
	//
	// Use SetDocument and GetDocument helpers.
	Document DocumentClass
}

// GameTypeID is TL type id of Game.
const GameTypeID = 0xbdf9653b

// Encode implements bin.Encoder.
func (g *Game) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode game#bdf9653b as nil")
	}
	b.PutID(GameTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode game#bdf9653b: field flags: %w", err)
	}
	b.PutLong(g.ID)
	b.PutLong(g.AccessHash)
	b.PutString(g.ShortName)
	b.PutString(g.Title)
	b.PutString(g.Description)
	if g.Photo == nil {
		return fmt.Errorf("unable to encode game#bdf9653b: field photo is nil")
	}
	if err := g.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode game#bdf9653b: field photo: %w", err)
	}
	if g.Flags.Has(0) {
		if g.Document == nil {
			return fmt.Errorf("unable to encode game#bdf9653b: field document is nil")
		}
		if err := g.Document.Encode(b); err != nil {
			return fmt.Errorf("unable to encode game#bdf9653b: field document: %w", err)
		}
	}
	return nil
}

// SetDocument sets value of Document conditional field.
func (g *Game) SetDocument(value DocumentClass) {
	g.Flags.Set(0)
	g.Document = value
}

// GetDocument returns value of Document conditional field and
// boolean which is true if field was set.
func (g *Game) GetDocument() (value DocumentClass, ok bool) {
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.Document, true
}

// Decode implements bin.Decoder.
func (g *Game) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode game#bdf9653b to nil")
	}
	if err := b.ConsumeID(GameTypeID); err != nil {
		return fmt.Errorf("unable to decode game#bdf9653b: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode game#bdf9653b: field flags: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode game#bdf9653b: field id: %w", err)
		}
		g.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode game#bdf9653b: field access_hash: %w", err)
		}
		g.AccessHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode game#bdf9653b: field short_name: %w", err)
		}
		g.ShortName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode game#bdf9653b: field title: %w", err)
		}
		g.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode game#bdf9653b: field description: %w", err)
		}
		g.Description = value
	}
	{
		value, err := DecodePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode game#bdf9653b: field photo: %w", err)
		}
		g.Photo = value
	}
	if g.Flags.Has(0) {
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode game#bdf9653b: field document: %w", err)
		}
		g.Document = value
	}
	return nil
}

// Ensuring interfaces in compile-time for Game.
var (
	_ bin.Encoder = &Game{}
	_ bin.Decoder = &Game{}
)
