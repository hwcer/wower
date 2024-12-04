package options

import (
	"github.com/hwcer/cosgo/uuid"
	"github.com/hwcer/updater"
	"github.com/hwcer/wower/model"
	"github.com/hwcer/wower/share"
)

func init() {
	updater.Config.IMax = share.Config.ITypes.GetIMax
	updater.Config.IType = share.Config.ITypes.GetIType
	updater.Config.ParseId = ParseId
}

func ParseId(u *updater.Updater, oid string) (int32, error) {
	if i, _, err := uuid.Split(oid, model.BaseSize, 1); err != nil {
		return 0, err
	} else {
		return int32(i), nil
	}
}
