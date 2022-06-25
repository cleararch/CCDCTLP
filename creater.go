package ccdctlp

import (
	"os"
	"os/exec"
	"sys_api"
)

func unpack_deb(name string, dir string) bool {
	if exec.Command("dpkg", "-X", name, dir).Run() != nil {
		return false
	}
	if exec.Command("dpkg", "-e", name, dir).Run() != nil {
		return false
	}
}
func sys_bind(dir string) {
	//挂载/proc /tmp /run /sys
	if mount_bind("/proc", dir+"/proc") == false {
		return false
	}
	if mount_bind("/tmp", dir+"/tmp") == false {
		return false
	}
	if mount_bind("/run", dir+"/run") == false {
		return false
	}
	if mount_bind("/run", dir+"/sys") == false {
		return false
	}
	return true
}
func deb_create_package(dir string, package_deb string) bool {
	err := os.MkdirAll(dir)
	if err != nil {
		return false
	}
	os.Chdir(dir)
	if unpack_deb(package_deb, dir) == false {
		return false
	}
	os.Mkdir(dir + "/root_sys")
	if mount_bind("/", dir+"/root_sys") == false {
		return false
	}
	if sys_bind() == false {
		return false
	}
}
