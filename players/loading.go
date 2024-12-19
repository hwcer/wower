package players

import (
	"github.com/hwcer/cosgo/logger"
	"github.com/hwcer/cosgo/times"
	"github.com/hwcer/wower/model"
	"github.com/hwcer/wower/players/options"
	"github.com/hwcer/wower/players/player"
)

// loading 初始加载用户到内存
func loading() (err error) {
	if options.Options.MemoryInstall == 0 {
		return nil
	}
	var rows []*model.Role
	now := times.Unix()
	lastTime := now - 7*86400
	tx := model.DB.Select("_id", "name").Order("update", -1)
	tx = tx.Where("update > ?", lastTime)
	tx = tx.Limit(options.Options.MemoryInstall).Find(&rows)
	if tx.Error != nil {
		return tx.Error
	}
	var p *player.Player

	for _, r := range rows {
		p = player.New(r.Uid)
		if err = p.Loading(true); err == nil {
			ps.Store(r.Uid, p)
			p.KeepAlive(now)
			logger.Debug("预加载用户: [%v] %v", p.Uid(), r.Name)
		}
	}
	logger.Trace("累计预加载用户数量:%v\n", len(rows))
	return
}
