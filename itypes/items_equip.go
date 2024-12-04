package itypes

import (
	"github.com/hwcer/wower/options"
)

var Equip = NewItemsIType(options.ITypeEquip)

func init() {
	Equip.SetStacked(false)
	//ITypeEquip.SetAttach(itemsEquipAttach)
}

//func itemsEquipAttach(u *updater.Updater, item *Item) (r any, err error) {
//	return
//}
