package ccdctlp

import (
	"os/exec"
)

func mount_bind(dir string, bind_dir string) bool {
	// dir:From
	// bind_dir:To
	// 挂载目录到执行目录
	if exec.Command("mount", dir, "--bind", bind_dir).Run != nil {
		return false
	} else {
		return true
	}
}
