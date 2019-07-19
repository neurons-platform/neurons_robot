package models

type Permission struct {
	PermissonCmd  string
	PermissonName string
	//是否有操作权限
	HasPermission bool
}

func (p *Permission) SetPermission(b bool) {
	p.HasPermission = b
}

var DefaultPermisssons = map[string]*Permission{
	"help":         &Permission{"help", "显示帮助", true},
	"cd":           &Permission{"cd", "清理磁盘", true},
	"pj":           &Permission{"pj", "流量平均分配", false},
	"sw":           &Permission{"sw", "切换集群", false},
	"sx":           &Permission{"sx", "上线集群", false},
	"last":         &Permission{"last", "查看监控报警", true},
	"set_wx_token": &Permission{"set_wx_token", "设置微信token", false},
	"get_wx_token": &Permission{"get_wx_token", "查看微信token", false},
	"gcmd":         &Permission{"gcmd", "获取清理命令", false},
	"clean":        &Permission{"clean", "清理图标缓存", false},
}
