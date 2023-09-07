package virtualbox

import (
	"strings"
)

type ImportOption uint32

const (
	ImportToVDI ImportOption = iota
	KeepAllMacs
	KeepNatMacs
)

// ImportOV imports ova or ovf from the given path
func ImportOV(path string, name string, options []ImportOption) error {
	args := []string{"import", path,
		"--vsys", "0",
		"--vmname", name,
	}
	opts := make([]string, 0)
	for _, o := range options {
		switch o {
		case ImportToVDI:
			opts = append(opts, "importtovdi")
		case KeepAllMacs:
			opts = append(opts, "keepallmacs")
		case KeepNatMacs:
			opts = append(opts, "keepnatmacs")
		}
	}
	if len(opts) > 0 {
		args = append(args, "--options")
		args = append(args, strings.Join(opts, ","))
	}

	_, _, err := Manage().run(args...)
	return err
}
