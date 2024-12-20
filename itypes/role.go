package itypes

import (
	"github.com/hwcer/cosgo/logger"
	"github.com/hwcer/cosgo/uuid"
	"github.com/hwcer/updater"
	"github.com/hwcer/updater/operator"
	"github.com/hwcer/wower/options"
)

const (
	RoleModelPlug = "_model_role_plug"
)

var Role = &roleIType{IType: NewIType(options.ITypeRole)}

//func init() {
//	it := []updater.IType{Role, ItemsGroup, ItemsPacks}
//	//ROLE
//	if err := updater.Register(updater.ParserTypeDocument, updater.RAMTypeAlways, &model.Role{}, it...); err != nil {
//		logger.Panic(err)
//	}
//}

type roleIType struct {
	*IType
	Upgrade roleUpgrade
	Builder *uuid.Builder
}
type roleUpgrade interface {
	Verify(u *updater.Updater, exp int64) (newExp int64)          //获得经验时进行检查
	Submit(u *updater.Updater, lv int32, exp int64) (newLv int32) //判断升级，返回新的等级
}

// Listener 监听升级状态

func (this *roleIType) Listener(u *updater.Updater, op *operator.Operator) {
	if op.Type == operator.TypesAdd && (op.Key == "exp" || op.Key == "Exp") {
		if this.Upgrade == nil {
			logger.Alert("ITypes.Options.RoleUpgrade is nil")
			return
		}
		if exp := this.Upgrade.Verify(u, op.Value); exp > 0 {
			op.Value = exp
			_ = u.Events.LoadOrCreate(RoleModelPlug, this.NewMiddleware)
		} else {
			op.Type = operator.TypesDrop //最大等级不给经验
		}
	}
}

func (this *roleIType) NewMiddleware() updater.Middleware {
	return &RoleMiddleware{}
}

type RoleMiddleware struct {
}

func (this RoleMiddleware) Emit(u *updater.Updater, t updater.EventType) bool {
	if t == updater.OnPreSubmit {
		return this.upgrade(u)
	}
	return true
}

func (this RoleMiddleware) upgrade(u *updater.Updater) bool {
	lv := int32(u.Val("lv"))
	exp := u.Val("exp")
	if newLv := Role.Upgrade.Submit(u, lv, exp); newLv > 0 && newLv != lv {
		role := u.Handle(options.ITypeRole)
		role.Add("lv", newLv-lv)
	}
	return false
}
