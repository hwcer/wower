package itypes

import (
	"github.com/hwcer/updater"
	"github.com/hwcer/wower/model"
	"github.com/hwcer/wower/options"
)

var Active = NewIType(options.ITypeActive)
var Config = NewIType(options.ITypeConfig) //后台配置的活动

func init() {
	Active.SetCreator(activeCreator)
	Config.SetCreator(activeCreator)
}

func activeCreator(u *updater.Updater, iid int32, val int64) (any, error) {
	var err error
	i := &model.Active{}
	i.Model.Init(u, iid)
	i.OID, err = Active.ObjectId(u, iid)
	i.Update = u.Time.Unix()
	return i, err
}
