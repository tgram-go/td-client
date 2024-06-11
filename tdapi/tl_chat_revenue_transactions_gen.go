// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// ChatRevenueTransactions represents TL type `chatRevenueTransactions#95e61144`.
type ChatRevenueTransactions struct {
	// Total number of transactions
	TotalCount int32
	// List of transactions
	Transactions []ChatRevenueTransaction
}

// ChatRevenueTransactionsTypeID is TL type id of ChatRevenueTransactions.
const ChatRevenueTransactionsTypeID = 0x95e61144

// Ensuring interfaces in compile-time for ChatRevenueTransactions.
var (
	_ bin.Encoder     = &ChatRevenueTransactions{}
	_ bin.Decoder     = &ChatRevenueTransactions{}
	_ bin.BareEncoder = &ChatRevenueTransactions{}
	_ bin.BareDecoder = &ChatRevenueTransactions{}
)

func (c *ChatRevenueTransactions) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.TotalCount == 0) {
		return false
	}
	if !(c.Transactions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatRevenueTransactions) String() string {
	if c == nil {
		return "ChatRevenueTransactions(nil)"
	}
	type Alias ChatRevenueTransactions
	return fmt.Sprintf("ChatRevenueTransactions%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatRevenueTransactions) TypeID() uint32 {
	return ChatRevenueTransactionsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatRevenueTransactions) TypeName() string {
	return "chatRevenueTransactions"
}

// TypeInfo returns info about TL type.
func (c *ChatRevenueTransactions) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatRevenueTransactions",
		ID:   ChatRevenueTransactionsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TotalCount",
			SchemaName: "total_count",
		},
		{
			Name:       "Transactions",
			SchemaName: "transactions",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatRevenueTransactions) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueTransactions#95e61144 as nil")
	}
	b.PutID(ChatRevenueTransactionsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatRevenueTransactions) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueTransactions#95e61144 as nil")
	}
	b.PutInt32(c.TotalCount)
	b.PutInt(len(c.Transactions))
	for idx, v := range c.Transactions {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare chatRevenueTransactions#95e61144: field transactions element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatRevenueTransactions) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueTransactions#95e61144 to nil")
	}
	if err := b.ConsumeID(ChatRevenueTransactionsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatRevenueTransactions#95e61144: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatRevenueTransactions) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueTransactions#95e61144 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueTransactions#95e61144: field total_count: %w", err)
		}
		c.TotalCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueTransactions#95e61144: field transactions: %w", err)
		}

		if headerLen > 0 {
			c.Transactions = make([]ChatRevenueTransaction, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ChatRevenueTransaction
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare chatRevenueTransactions#95e61144: field transactions: %w", err)
			}
			c.Transactions = append(c.Transactions, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatRevenueTransactions) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueTransactions#95e61144 as nil")
	}
	b.ObjStart()
	b.PutID("chatRevenueTransactions")
	b.Comma()
	b.FieldStart("total_count")
	b.PutInt32(c.TotalCount)
	b.Comma()
	b.FieldStart("transactions")
	b.ArrStart()
	for idx, v := range c.Transactions {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode chatRevenueTransactions#95e61144: field transactions element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatRevenueTransactions) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueTransactions#95e61144 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatRevenueTransactions"); err != nil {
				return fmt.Errorf("unable to decode chatRevenueTransactions#95e61144: %w", err)
			}
		case "total_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueTransactions#95e61144: field total_count: %w", err)
			}
			c.TotalCount = value
		case "transactions":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value ChatRevenueTransaction
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode chatRevenueTransactions#95e61144: field transactions: %w", err)
				}
				c.Transactions = append(c.Transactions, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode chatRevenueTransactions#95e61144: field transactions: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTotalCount returns value of TotalCount field.
func (c *ChatRevenueTransactions) GetTotalCount() (value int32) {
	if c == nil {
		return
	}
	return c.TotalCount
}

// GetTransactions returns value of Transactions field.
func (c *ChatRevenueTransactions) GetTransactions() (value []ChatRevenueTransaction) {
	if c == nil {
		return
	}
	return c.Transactions
}