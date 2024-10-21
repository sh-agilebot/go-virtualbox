package virtualbox

import (
	"syscall"
)

var sysProcAttr = &syscall.SysProcAttr{
	HideWindow: true,
}
