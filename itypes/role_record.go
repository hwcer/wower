package itype

import (
	"github.com/hwcer/wower/config"
)

var Record = NewIType(config.ITypeRecord)

//func init() {
//	if err := updater.Register(updater.ParserTypeValues, updater.RAMTypeAlways, &model.Record{}, Record); err != nil {
//		logger.Panic(err)
//	}
//}
