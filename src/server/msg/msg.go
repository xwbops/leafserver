package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

// Hello 消息结构体
type Hello struct {
	Name string
}

func init() {
	//注册一个JSON消息Hello
	Processor.Register(&Hello{})
}
