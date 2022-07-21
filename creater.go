package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Mount_bind(dir string, bind_dir string) bool {
	// dir:From
	// bind_dir:To
	// 硬链接目录到执行目录
	a := exec.Command("mount", dir, "-o", "bind", bind_dir).Run()
	if a != nil {
		fmt.Println(a)
		return false
	} else {
		return true
	}
}
func APT_install(packet string, dir string) bool {
	os.Chdir(dir)
	cmd := exec.Command("apt", "install", packet)
	press_y := exec.Command("echo", "y")
	cmd.Stdin, _ = press_y.StdoutPipe()
	cmd.Start()
	press_y.Run()
	err := cmd.Wait()
	if err == nil {
		return true
	} else {
		return false
	}
}
func Unpack_deb(name string, dir string) bool {
	//解包deb
	if exec.Command("dpkg", "-X", name, dir).Run() != nil {
		fmt.Println("0")
		return false
	}
	if exec.Command("dpkg", "-e", name, dir).Run() != nil {
		return false
	}
	return true
}

// func Sys_bind(dir string) bool {
// 	//挂载/proc /tmp /run /sys
// 	if !Mount_bind("/proc", dir+"/proc") {
// 		return false
// 	}
// 	if !Mount_bind("/tmp", dir+"/tmp") {
// 		return false
// 	}
// 	if !Mount_bind("/run", dir+"/run") {
// 		return false
// 	}
// 	if !Mount_bind("/sys", dir+"/sys") {
// 		return false
// 	}
// 	return true
// }
func Create_package(dir string, package_deb string, config []string) bool {
	// 解包deb并创建虚拟环境，暂时无法考虑配置文件
	err := os.MkdirAll(dir+"/package", 775)
	if err != nil {
		return false
	}
	os.Chdir(dir + "/package")
	if !Unpack_deb(package_deb, dir+"/package") {
		return false
	}
	os.Chdir(dir)
	os.Mkdir(dir+"/root_sys", 775)
	if !Mount_bind("/", dir+"/root_sys") {
		return false
	}
	err = os.MkdirAll(dir+"/venu", 775)
	os.Chdir(dir + "/venu")
	filepath.Walk(dir+"/root_sys", func(path string, info fs.FileInfo, err error) error {
		default_judgment := strings.Replace(dir+"/venu"+path, dir+"/root_sys", "", -1)
		config_judgment := strings.Trim(default_judgment, dir+"/venu/")
		for _, config_temp := range config {
			if strings.HasPrefix(config_judgment, config_temp) {
				return err
			}
		}
		temp, _ := os.Stat(path)
		if temp.IsDir() {
			os.MkdirAll(default_judgment, 775)
			return err
		}
		fmt.Println(os.Symlink(path, default_judgment))
		fmt.Print("\n")
		return err
	})
	return true
}
