package mtproto

import (
	"golang.org/x/xerrors"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/internal/mt"
)

func (c *Conn) handleFutureSalts(b *bin.Buffer) error {
	var res mt.FutureSalts

	if err := res.Decode(b); err != nil {
		return xerrors.Errorf("error decode: %w", err)
	}

	c.salts.Store(res.Salts)
	return nil
}
