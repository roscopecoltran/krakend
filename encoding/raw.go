package encoding

import (
	"bytes"
	"io"
)

// raw will not encode just generate a map[string]inteface{} with one element body
const RAW = "raw"

// NewRawDecoder returns the raw decoder
func NewRawDecoder() Decoder {
	return func(r io.Reader, v *map[string]interface{}, targets []map[string]string) error {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r)
		*(v) = map[string]interface{}{
			"body": buf.String(),
		}
		return nil
	}
}
