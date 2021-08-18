// Code generated by gotdgen, DO NOT EDIT.

package mt

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

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
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// MsgsAllInfo represents TL type `msgs_all_info#8cc0d131`.
type MsgsAllInfo struct {
	// MsgIDs field of MsgsAllInfo.
	MsgIDs []int64
	// Info field of MsgsAllInfo.
	Info []byte
}

// MsgsAllInfoTypeID is TL type id of MsgsAllInfo.
const MsgsAllInfoTypeID = 0x8cc0d131

func (m *MsgsAllInfo) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.MsgIDs == nil) {
		return false
	}
	if !(m.Info == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MsgsAllInfo) String() string {
	if m == nil {
		return "MsgsAllInfo(nil)"
	}
	type Alias MsgsAllInfo
	return fmt.Sprintf("MsgsAllInfo%+v", Alias(*m))
}

// FillFrom fills MsgsAllInfo from given interface.
func (m *MsgsAllInfo) FillFrom(from interface {
	GetMsgIDs() (value []int64)
	GetInfo() (value []byte)
}) {
	m.MsgIDs = from.GetMsgIDs()
	m.Info = from.GetInfo()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MsgsAllInfo) TypeID() uint32 {
	return MsgsAllInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*MsgsAllInfo) TypeName() string {
	return "msgs_all_info"
}

// TypeInfo returns info about TL type.
func (m *MsgsAllInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "msgs_all_info",
		ID:   MsgsAllInfoTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MsgIDs",
			SchemaName: "msg_ids",
		},
		{
			Name:       "Info",
			SchemaName: "info",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MsgsAllInfo) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msgs_all_info#8cc0d131 as nil")
	}
	b.PutID(MsgsAllInfoTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MsgsAllInfo) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msgs_all_info#8cc0d131 as nil")
	}
	b.PutVectorHeader(len(m.MsgIDs))
	for _, v := range m.MsgIDs {
		b.PutLong(v)
	}
	b.PutBytes(m.Info)
	return nil
}

// GetMsgIDs returns value of MsgIDs field.
func (m *MsgsAllInfo) GetMsgIDs() (value []int64) {
	return m.MsgIDs
}

// GetInfo returns value of Info field.
func (m *MsgsAllInfo) GetInfo() (value []byte) {
	return m.Info
}

// Decode implements bin.Decoder.
func (m *MsgsAllInfo) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msgs_all_info#8cc0d131 to nil")
	}
	if err := b.ConsumeID(MsgsAllInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode msgs_all_info#8cc0d131: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MsgsAllInfo) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msgs_all_info#8cc0d131 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode msgs_all_info#8cc0d131: field msg_ids: %w", err)
		}

		if headerLen != 0 {
			m.MsgIDs = make([]int64, 0, headerLen)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode msgs_all_info#8cc0d131: field msg_ids: %w", err)
			}
			m.MsgIDs = append(m.MsgIDs, value)
		}
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode msgs_all_info#8cc0d131: field info: %w", err)
		}
		m.Info = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MsgsAllInfo.
var (
	_ bin.Encoder     = &MsgsAllInfo{}
	_ bin.Decoder     = &MsgsAllInfo{}
	_ bin.BareEncoder = &MsgsAllInfo{}
	_ bin.BareDecoder = &MsgsAllInfo{}
)
