package handler

import (
	U "github.com/neurons-platform/gotools/utils"
	M "neurons_robot/models"
)

func checkShellPermission(shellCmd, erp string) bool {
	for _, shellPermission := range M.DefaultShellPermissions.ShellPermissionList {
		if shellPermission.ShellName == shellCmd {
			c, err := U.Contain(erp, shellPermission.UserList)
			if err != nil || !c {
				return false
			}
		}
	}
	return true
}
