package main

import (
	"os"
	"os/exec"
)

func Mount_bind(dir string, bind_dir string) bool {
	// dir:From
	// bind_dir:To
	// 硬链接目录到执行目录
	if exec.Command("mount", dir, "--bind", bind_dir).Run != nil {
		return false
	} else {
		return true
	}
}

func APT_install(packet string, dir string) bool{
	os.Chdir(dir)
	cmd := exec.Command("apt", "install", packet)
	press_y := exec.Command("echo", "y")
	cmd.Stdin ,_ := press_y.StdoutPipe()
	cmd.Start()
	press_y.Run()
	err := cmd.Wait()
	if !err{
		return true
	} else{
		return false
	}
}
