package share

import (
	"github.com/hwcer/cosgo/logger"
)

func init() {
	logger.SetPathTrim("src")
	logger.SetCallDepth(4)
}
