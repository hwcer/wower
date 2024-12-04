package itypes

import (
	"github.com/hwcer/cosgo/logger"
	"github.com/hwcer/updater"
	"github.com/hwcer/wower/model"
	"github.com/hwcer/wower/options"
)

var Task = &taskIType{}

func init() {
	Task.IType = NewIType(options.ITypeTask)
	Task.SetCreator(taskCreator)
	Task.GetConfig = func(i int32) TaskConfig {
		logger.Alert("请设置 itypes.Task.GetConfig")
		return nil
	}
}

type TaskConfig interface {
	GetKey() int32
	GetArgs() []int32
	GetGoal() int32
	GetCondition() int32
}

func taskCreator(u *updater.Updater, iid int32, val int64) (any, error) {
	i := &model.Task{}
	i.Init(u, iid)
	i.OID, _ = Shop.ObjectId(u, iid)
	i.Value = int32(val)
	return i, nil
}

type taskIType struct {
	*IType
	GetConfig func(int32) TaskConfig
}
