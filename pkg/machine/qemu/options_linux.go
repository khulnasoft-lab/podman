package qemu

import (
	"github.com/khulnasoft-lab/podman/v4/pkg/rootless"
	"github.com/khulnasoft-lab/podman/v4/pkg/util"
)

func getRuntimeDir() (string, error) {
	if !rootless.IsRootless() {
		return "/run", nil
	}
	return util.GetRuntimeDir()
}
