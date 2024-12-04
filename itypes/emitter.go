package itypes

import "github.com/hwcer/cosgo/logger"

var Emitter = struct {
	GetConfig func(int32) EmitterConfig
}{
	GetConfig: getEmitterConfig,
}

type EmitterConfig interface {
	GetDaily() int32
	GetRecord() int32
	GetEvents() int32
	GetUpdate() int32
	GetReplace() int32
}

func getEmitterConfig(i int32) EmitterConfig {
	logger.Alert("请设置 itypes.Record.GetConfig")
	return nil
}
