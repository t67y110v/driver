package model

import (
	"encoding/binary"
	"os"
	"reflect"

	"github.com/t67y110v/driver/internal/driver/converter"
)

type Util interface {
	Get(n int) []byte
	Read(buf []byte) error
	FillMeta()
}

func (r *RawResponse) Get(n int) []byte {
	result := r.Bytes[r.Offset : r.Offset+n]
	r.Offset += n
	return result
}

func (r *Response) Read(buf []byte) error {
	file := os.NewFile(0, "")
	defer file.Close()
	n, err := file.Read(buf)
	if err != nil {
		return err
	}

	r.raw.Bytes = make([]byte, n)
	copy(r.raw.Bytes, buf)

	r.raw.FillMeta()

	return nil
}
func (r *RawResponse) FillMeta() {
	header := []byte{0xF8, 0x55, 0xCE}
	responses := map[byte]string{
		0x12: "CMD_ACK_SET_TARE",
		0x27: "CMD_ACK_SET",
		0x15: "CMD_ERROR",
		0x24: "CMD_ACK_MASSA",
	}

	r.Meta.Valid = true

	if !reflect.DeepEqual(r.Bytes[:3], header) {
		r.Meta.Valid = false
		return
	}

	r.Meta.Len = binary.LittleEndian.Uint16(r.Bytes[3:5])

	r.Meta.CommandCode = r.Bytes[5]
	r.Meta.CommandName = responses[r.Meta.CommandCode]

	r.Meta.CRC = binary.LittleEndian.Uint16(r.Bytes[6+r.Meta.Len-1 : 6+r.Meta.Len+1])

	if r.Meta.CRC != converter.CRC16(0, r.Bytes[5:6+r.Meta.Len-1], r.Meta.Len) {
		r.Meta.Valid = false
	}

}

func (r *Response) Raw() RawResponse {
	return r.raw
}
