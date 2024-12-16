package context

import (
	"github.com/hwcer/wower/options"
	"github.com/hwcer/wower/share"
	"path"
)

const (
	ServiceMethodDebug       = "/debug"
	ServiceMethodRoleRenewal = "/role/renewal" //续约
)

var Verify func(*Context) (Token, error) //登录验证

//func Start() error {
//	//return loadAlphaAccount()
//	return nil
//}

func MethodGrade(s string) int8 {
	//if options.Gate.Prefix != "" {
	//	routePrefix := registry.Join(options.Gate.Prefix)
	//	serviceMethod = strings.TrimPrefix(serviceMethod, routePrefix)
	//}
	s = share.TrimServiceMethod(s)
	s = path.Join(options.ServiceTypeGame, s)
	return share.Authorizes.Get(s)
}
