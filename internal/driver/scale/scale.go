package scale

import "encoding/binary"

type Scale struct {
}

func NewScale() *Scale {
	return &Scale{}
}

func (s *Scale) CRC16(crc uint16, buf []byte, len uint16) uint16 {
	var bits, k, a, temp uint16

	for k = 0; k < len; k++ {
		a = 0
		temp = (crc >> 8) << 8
		for bits = 0; bits < 8; bits++ {
			if (temp^a)&0x8000 != 0 {
				a = (a << 1) ^ 0x1021
			} else {
				a <<= 1
			}
			temp <<= 1
		}
		crc = a ^ (crc << 8) ^ (binary.LittleEndian.Uint16([]byte{buf[k], 0x0}) & 0xFF)
	}
	return crc
}
func (s *Scale) MakeMessage(data []byte) []byte {
	var result []byte

	result = append(result, []byte{0xF8, 0x55, 0xCE}...)

	lenBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(lenBytes, uint16(len(data)))
	result = append(result, lenBytes...)

	result = append(result, data...)

	crcBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(crcBytes, s.CRC16(0, data, uint16(len(data))))
	result = append(result, crcBytes...)

	return result
}

func (s *Scale) MakeMessagefForSetValue(data []byte, value []byte) []byte {
	var result []byte

	result = append(result, []byte{0xF8, 0x55, 0xCE}...)

	lenBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(lenBytes, uint16(len(data)))
	result = append(result, lenBytes...)

	result = append(result, data...)

	result = append(result, value...)

	crcBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(crcBytes, s.CRC16(0, data, uint16(len(data))))
	result = append(result, crcBytes...)

	return result
}
