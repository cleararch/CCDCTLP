package main

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

func Unpack_deb(name string, dir string) bool {
	//解包deb
	if exec.Command("dpkg", "-X", name, dir).Run() != nil {
		return false
	}
	if exec.Command("dpkg", "-e", name, dir).Run() != nil {
		return false
	}
	return true
}
func Sys_bind(dir string) bool {
	//挂载/proc /tmp /run /sys
	if !Mount_bind("/proc", dir+"/proc") {
		return false
	}
	if !Mount_bind("/tmp", dir+"/tmp") {
		return false
	}
	if !Mount_bind("/run", dir+"/run") {
		return false
	}
	if !Mount_bind("/sys", dir+"/sys") {
		return false
	}
	return true
}
func Create_package(dir string, package_deb string) bool {
	// 解包deb并创建虚拟环境，暂时无法考虑配置文件
	err := os.MkdirAll(dir, 777)
	if err != nil {
		return false
	}
	os.Chdir(dir)
	if !Unpack_deb(package_deb, dir) {
		return false
	}
	os.Mkdir(dir+"/root_sys", 777)
	if !Mount_bind("/", dir+"/root_sys") {
		return false
	}
	if !Sys_bind(dir) {
		return false
	}
	filepath.Walk(dir+"/root_sys", func(path string, info fs.FileInfo, err error) error {
		os.Symlink(path, dir+path)
		return err
	})
	return true
}
