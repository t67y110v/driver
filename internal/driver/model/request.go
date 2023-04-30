package model

type Request struct {
	raw RawResponse
}

type RawRequest struct {
	Bytes  []byte
	Offset int
	Meta   RequestMeta
}

type SetTareRequest struct {
	Request
	Tare [4]byte
}

type RequestMeta struct {
	Valid       bool
	Len         uint16
	CommandCode byte
	CommandName string
	CRC         uint16
}
