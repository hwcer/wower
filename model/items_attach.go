package model

import "github.com/hwcer/cosgo/logger"

var Attach = att{}

const (
	AttachTypeTicket = "ts"  //门票上次结算时间
	AttachTypeExpire = "ttl" //通用过期时间
)

func init() {
	Attach.Register(AttachTypeTicket)
	Attach.Register(AttachTypeExpire)
}

type att struct {
	dict map[string]struct{}
}

func (a att) Has(k string) bool {
	_, has := a.dict[k]
	return has
}

func (a att) Register(k string) {
	if a.dict == nil {
		a.dict = make(map[string]struct{})
	}
	if _, ok := a.dict[k]; ok {
		logger.Panic("already registered: " + k)
	}

	a.dict[k] = struct{}{}
}
