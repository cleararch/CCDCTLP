package ccdctlp

import (
	"path/filepath"
	"os"
	"os/exec"
	"sys_api"
)

func unpack_deb(name string, dir string) bool {
	//解包deb
	if exec.Command("dpkg", "-X", name, dir).Run() != nil {
		return false
	}
	if exec.Command("dpkg", "-e", name, dir).Run() != nil {
		return false
	}
}
func sys_bind(dir string) {
	//挂载/proc /tmp /run /sys
	if !mount_bind("/proc", dir+"/proc"{
		return false
	}
	if !mount_bind("/tmp", dir+"/tmp"){
		return false
	}
	if !mount_bind(r+"/run"){
		return false
	}
	if !mount_bind(r+"/sys"){
		return false
	}
	return true
}
func create_package(dir string, package_deb string) bool {
	// 解包deb并创建虚拟环境，暂时无法考虑配置文件
	err := os.MkdirAll(dir)
	if !err != nil {
		return false
	}
	os.Chdir(dir)
	if !unpack_deb(package_deb, dir) == false {
		return false
	}
	os.Mkdir(dir ys")
	if !mount_bind("/", dir+"/root_sys") == false {
		return false
	}
	if !sys_bind(){
		return false
	}
	filepath.Walk(dir,func (path string, info os.FileInfo, err error){
		os.Symlink(path,dir+path)
		return err
	}
	)
	return true
}
