package itypes

import (
	"github.com/hwcer/updater"
	"github.com/hwcer/wower/model"
	"github.com/hwcer/wower/options"
)

var Shop = NewIType(options.ITypeShop)

func init() {
	Shop.SetCreator(shopCreator)
	//if err := updater.Register(updater.ParserTypeCollection, updater.RAMTypeNone, &model.Shop{}, Shop); err != nil {
	//	logger.Panic(err)
	//}
}

func shopCreator(u *updater.Updater, iid int32, val int64) (any, error) {
	i := &model.Shop{}
	i.Init(u, iid)
	i.OID, _ = Shop.ObjectId(u, iid)
	i.Val = int32(val)
	return i, nil
}
