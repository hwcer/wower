package options

import (
	"github.com/hwcer/cosgo"
	"github.com/hwcer/cosrpc/xshare"
	"sync/atomic"
)

var initialize int32

const (
	ServiceTypeGate   = "gate"
	ServiceTypeGame   = "game"
	ServiceTypeChat   = "chat" //聊天
	ServiceTypeBattle = "battle"
	ServiceTypeRooms  = "rooms"  //游戏大厅
	ServiceTypeSocial = "social" //社交用户中心
)

func Initialize() error {
	if atomic.CompareAndSwapInt32(&initialize, 0, 1) {
		return cosgo.Config.Unmarshal(Options)
	}
	return nil
}

var Service = map[string]string{}

var Options = &struct {
	Data    string //静态数据地址
	Debug   bool
	Appid   string
	Master  string
	Secret  string //秘钥,必须8位
	Verify  int8   `json:"verify"` //平台验证方式,0-不验证，1-仅仅验证签名，2-严格模式
	Service map[string]string
	Game    *game
	Gate    *gate
	Rpcx    *xshare.Rpcx
}{
	Verify:  1,
	Service: Service,
	Game:    Game,
	Gate:    Gate,
	Rpcx:    xshare.Options,
}
