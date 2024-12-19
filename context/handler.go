package context

import (
	"github.com/hwcer/cosgo"
	"github.com/hwcer/cosgo/registry"
	"github.com/hwcer/cosgo/times"
	"github.com/hwcer/cosgo/values"
	"github.com/hwcer/wower/options"
	"github.com/hwcer/wower/players"
	"github.com/hwcer/wower/players/player"
	"github.com/hwcer/wower/share"
	"strings"
)

func caller(c *Context, node *registry.Node) any {
	var v interface{}
	if node.IsFunc() {
		m := node.Method().(func(*Context) interface{})
		v = m(c)
	} else if s, ok := node.Binder().(Caller); ok {
		v = s.Caller(node, c)
	} else {
		vs := node.Call(c)
		v = vs[0].Interface()
	}

	//直接返回二进制不做任何处理
	if r, ok := v.([]byte); ok {
		return r
	}

	r := Parse(v)
	r.Time = c.Time().UnixMilli()
	if r.Code == 0 && c.Player != nil {
		var err error
		if r.Cache, err = c.Player.Submit(); err != nil {
			r = Error(err)
		} else {
			//r.Notify = c.Player.Notify.Get()
		}
	}

	return r
}

// limits 检查并 lock AND reset
func verify(c *Context, handle func() error) (err error) {
	path := c.ServiceMethod()
	if strings.HasPrefix(path, ServiceMethodDebug) && !cosgo.Debug() {
		return values.Errorf(0, "unauthorized")
	}
	if !share.HasServiceMethod(path) {
		return handle() //内网通信启用玩家数据
	} else {
		//c.Binder = binder.New(binder.MIMEPROTOBUF) //外部通信
	}

	l := MethodGrade(path)
	if l == share.AuthorizesTypeNone {
		return handle()
	}
	if l == share.AuthorizesTypeOAuth {
		if guid := c.GetMetadata(options.ServiceMetadataGUID); guid == "" {
			return values.Errorf(0, "guid empty")
		} else {
			return handle()
		}
	}
	uid := c.Uid()
	if uid == 0 {
		return values.Errorf(0, "not select role")
	}

	err = players.Try(uid, func(p *player.Player) error {
		if p == nil {
			return share.ErrLogin
		}
		c.Player = p
		c.Player.KeepAlive(c.Unix())

		if update := c.Player.Role.Val("update"); update < times.Daily(0).Unix() && c.ServiceMethod() != ServiceMethodRoleRenewal {
			return share.ErrNeedResetSession
		}
		//尝试重新上线
		meta := c.Metadata()
		if c.Player.Status != player.StatusConnected {
			if e := players.Connect(p, meta); e != nil {
				return e
			}
		} else if session := meta[options.ServicePlayerSession]; session != p.Session {
			return share.ErrReplaced
		}

		return handle()
	})
	return
}
