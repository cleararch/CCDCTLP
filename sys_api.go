package main

import (
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
