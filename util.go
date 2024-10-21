package virtualbox

import (
	"context"
	"net"
	"os"
	"os/exec"
)

// ParseIPv4Mask parses IPv4 netmask written in IP form (e.g. 255.255.255.0).
// This function should really belong to the net package.
func ParseIPv4Mask(s string) net.IPMask {
	mask := net.ParseIP(s)
	if mask == nil {
		return nil
	}
	return net.IPv4Mask(mask[12], mask[13], mask[14], mask[15])
}

// Run is a helper method used to execute the commands using the configured
// VBoxManage path. The command should be omitted and only the arguments
// should be passed. It will return the stdout, stderr and error if one
// occured during command execution.
func Run(_ context.Context, args ...string) (string, string, error) {
	// TODO: Convert the function so you can pass in the context.
	return Manage().run(args...)
}

func isFileExists(name string) bool {
	if name == "" {
		return false
	}

	if fi, err := os.Stat(name); err == nil {
		return !fi.IsDir()
	}
	return false
}

// newExecCmd is a wrapper to exec.Command
func newExecCmd(name string, arg ...string) *exec.Cmd {
	cmd := exec.Command(name, arg...)
	cmd.SysProcAttr = sysProcAttr
	return cmd
}
