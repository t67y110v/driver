package model

type SetZeroResponse struct {
	Response
}
type ResponseMeta struct {
	Valid       bool
	Len         uint16
	CommandCode byte
	CommandName string
	CRC         uint16
}

type RawResponse struct {
	Bytes  []byte
	Offset int
	Meta   ResponseMeta
}

type Response struct {
	raw RawResponse
}

type GetMassaResponse struct {
	Response
	Weight   uint32
	Division byte
	Stable   bool
	Net      bool
	Zero     bool
}

type SetTareResponse struct {
	Response
}

type SetTareZeroResponse struct {
	Response
}
