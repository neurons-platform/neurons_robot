package models

import (
	CH "github.com/neurons-platform/gotools/cache"
)

type User struct {
	UserName    string
	Email       string
	Erp         string
	Mobile      string
	Token       string
	Password    string
	Permissions map[string]*Permission
}

var Users, _ = CH.New(100)

func init() {
	var jingminglang = User{UserName: "jingmignlang", Permissions: DefaultPermisssons}
	var huangxiaoming = User{UserName: "huangxiaoming", Permissions: DefaultPermisssons}
	jingminglang.Permissions["sx"].SetPermission(true)
	jingminglang.Permissions["sx"].SetPermission(true)
	jingminglang.Permissions["pj"].SetPermission(true)
	jingminglang.Permissions["sw"].SetPermission(true)
	jingminglang.Permissions["set_wx_token"].SetPermission(true)
	jingminglang.Permissions["get_wx_token"].SetPermission(true)
	jingminglang.Permissions["gcmd"].SetPermission(true)
	jingminglang.Permissions["clean"].SetPermission(true)

	huangxiaoming.Permissions["sx"].SetPermission(true)
	huangxiaoming.Permissions["pj"].SetPermission(true)
	huangxiaoming.Permissions["sw"].SetPermission(true)
	huangxiaoming.Permissions["set_wx_token"].SetPermission(true)
	huangxiaoming.Permissions["get_wx_token"].SetPermission(true)
	huangxiaoming.Permissions["gcmd"].SetPermission(true)
	huangxiaoming.Permissions["clean"].SetPermission(true)
	Users.Add("jingminglang", jingminglang)
	Users.Add("huangxiaoming", huangxiaoming)
}

func GetUser(erp string) User {
	u, ok := Users.Get(erp)
	if ok {
		return u.(User)
	}
	return User{UserName: erp, Permissions: DefaultPermisssons}
}
