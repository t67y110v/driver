package converter

var Commands = map[string]byte{
	"CMD_GET_NAME":     0x20,
	"CMD_SET_NAME":     0x22,
	"CMD_GET_ETHERNET": 0x2D,
	"CMD_SET_ETHERNET": 0x39,
	"CMD_GET_WIFI":     0x3A,
	"CMD_SET_WIFI":     0x3C,
	"CMD_GET_MASSA":    0x23,
	"CMD_SET_TARE":     0xA3,
	"CMD_SET_ZERO":     0x72,
	"CMD_ACK_SET":      0x27,
}

var Errors = map[byte]string{
	0x07: "CMD_ERROR",
	0x08: "CMD_ERROR",
	0x09: "CMD_ERROR",
	0x0A: "CMD_ERROR",
	0x0B: "CMD_ERROR",
	0x10: "CMD_ERROR",
	0x11: "CMD_ERROR",
	0x15: "CMD_ERROR",
	0x17: "CMD_ERROR",
	0x18: "CMD_ERROR",
	0x19: "CMD_ERROR",
	0xF0: "CMD_ERROR",
}

var Responses = map[string]byte{
	"CMD_ACK_SET_TARE": 0x12,
	"CMD_NACK_TARE":    0x15,
	"CMD_ACK_SET":      0x27,
	"CMD_ERROR":        0x15,
	"CMD_ACK_MASSA":    0x24,
}
