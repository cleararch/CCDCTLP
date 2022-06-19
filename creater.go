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
func deb_create_package(dir string, package_deb string) bool {
	err := os.MkdirAll(dir)
	if err != nil {
		return false
	}
	os.Chdir(dir)
}
