package main

import (
	"os/exec"
	"syscall"

	"github.com/sh-agilebot/go-virtualbox"
)

func open(args ...string) *exec.Cmd {
	argv := append([]string{"/c"}, "start")
	argv = append(argv, args...)
	virtualbox.Debug("Executing %v %v", "cmd", argv)
	cmd := exec.Command("cmd", argv...) // #nosec
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}
