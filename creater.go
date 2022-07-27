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
	// 安装依赖包用
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

func sys_bind(dir string) bool {
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
func Deb_uncompress(dir string, package_deb string) []string {
	// 解包deb
	var file_list []string
	err := os.MkdirAll(dir+"/package", 775)
	if err != nil {
		return file_list
	}
	os.Chdir(dir + "/package")
	if !Unpack_deb(package_deb, dir+"/package") {
		return file_list
	}
	filepath.Walk(dir+"/package", func(path string, info fs.FileInfo, err error) error {
		temp := strings.Replace(path, dir+"/package", "", -1)
		file_list = append(file_list, temp)
		return err
	})
	return file_list
}
func Create_package(dir string, package_deb string, config []string) bool {
	deb := Deb_uncompress(dir, package_deb) //解包deb
	if deb != nil {
		config = append(config, deb...)
	} else {
		return false
	}
	os.Chdir(dir)
	os.Mkdir(dir+"/root_sys", 775)
	if !Mount_bind("/", dir+"/root_sys") {
		return false
	}
	sys_bind(dir)
	filepath.Walk(dir+"/root_sys", func(path string, info fs.FileInfo, err error) error {
		// dir为虚拟根的上一层，/venu为虚拟根，/root_sys为root fs的一个映射
		default_judgment := strings.Replace(dir+path, dir+"/root_sys", "", -1)
		config_judgment := strings.Replace(default_judgment, dir, "", -1)
		for _, config_temp := range config {
			// config_judgment即为「实际路径」，config_temp即为需要排除的「实际路径」
			if strings.HasPrefix(config_judgment, config_temp) {
				// deb判断，实现创建一个除deb包、config外均链接root的虚拟根
				for _, deb_path := range deb {
					if strings.HasPrefix(config_judgment, deb_path) {
						temp, _ := os.Stat(path)
						if temp.IsDir() {
							os.MkdirAll(default_judgment, 775)
							return err
						}
						fmt.Println(os.Symlink(path, deb_path))
						fmt.Print("\n")
						return err
					}
				}
				fmt.Println("1")
				return err
			}
		}
		//若为文件夹，则创建，而非链接
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
