package xrpc

const (
	HEADER_SIZE      = 36
	HEADER_MAGIC_NUM = 0x11451419
)

type Header struct {
	Id       uint16
	Version  uint16
	LogId    uint32
	Provider [16]byte // 服务提供者
	MagicNum uint32
	Reserved uint32 //保留
	BodyLen  uint32
}
