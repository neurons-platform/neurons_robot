package models

import  (
	U "github.com/neurons-platform/gotools/utils"
)


type ShellPermission struct {
	ShellName string `yaml:"shellName"`
	UserList []string `yaml:"userList"`
}

type T struct {
	ShellPermissionList []ShellPermission `yaml:"ShellPermissionList"`
}

var DefaultShellPermissions = getShellPermissions()

func getShellPermissions() T {
	t := T{}
	str := U.ReadAllFile("./config/shellPermission.yaml")
	U.YamlStrToStruct(str,&t)
	return t

}
