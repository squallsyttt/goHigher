package codec

import "io"

type Header struct {
	ServiceMethod string // format "Service.Method"
	Seq           uint64 // sequence number chosen by client
	Error         string
}

// Codec 对消息体进行编码解码的接口
// 抽象出接口 是为了实现不同的codec实例
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

// NewCodecFunc  这是codec的构造函数
type NewCodecFunc func(io.ReadWriteCloser) Codec
type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implemented
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	//make 只能用于map slice chan 三个类型
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
	NewCodecFuncMap[JsonType] = NewGobCodec
}
