package plist

import (
	"github.com/pkg6/appgo/seekbuf"
	"github.com/pkg6/go-plist"
	"io"
)

func Decode(r io.Reader, d interface{}) error {
	buf, err := seekbuf.Open(r, seekbuf.MemoryMode)
	if err != nil {
		return err
	}
	if err := plist.NewDecoder(buf).Decode(d); err != nil {
		return err
	}
	return nil
}
