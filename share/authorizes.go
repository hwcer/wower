package share

import (
	"github.com/hwcer/wower/options"
	"path"
	"strings"
)

// 接口权限设置

const (
	AuthorizesTypeNone      int8 = iota //不需要登录
	AuthorizesTypeOAuth                 //需要认证
	AuthorizesTypeCharacter             //需要选择角色
)

var Authorizes = authorizes{}

type authorizes map[string]int8

func init() {
	s := map[string]int8{
		"/ping":        AuthorizesTypeNone,
		"/heart":       AuthorizesTypeNone,
		"/login":       AuthorizesTypeNone,
		"/roles":       AuthorizesTypeOAuth,
		"/role/create": AuthorizesTypeOAuth,
		"/role/select": AuthorizesTypeOAuth,
	}
	for k, v := range s {
		Authorizes.Set(options.ServiceTypeGame, k, v)
	}
}

func (auth authorizes) Set(servicePath, serviceMethod string, i int8) {
	r := path.Join(servicePath, serviceMethod)
	r = strings.ToLower(r)
	if !strings.HasPrefix(r, "/") {
		r = "/" + r
	}
	auth[r] = i
}

func (auth authorizes) Get(s string) int8 {
	s = strings.ToLower(s)
	if !strings.HasPrefix(s, "/") {
		s = "/" + s
	}
	if v, ok := auth[s]; !ok {
		return AuthorizesTypeCharacter
	} else {
		return v
	}
}

func (auth authorizes) Range(f func(s string, i int8)) {
	for k, v := range auth {
		f(k, v)
	}
}
