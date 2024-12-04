package itype

import (
	"github.com/hwcer/wower/config"
)

var Goods = NewIType(config.ITypeGoods)

//func init() {
//	if err := updater.Register(updater.ParserTypeValues, updater.RAMTypeAlways, &model.Goods{}, Goods); err != nil {
//		logger.Panic(err)
//	}
//}
