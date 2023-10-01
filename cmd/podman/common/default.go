package common

import (
	"github.com/khulnasoft-lab/podman/v4/cmd/podman/registry"
)

var (
	// Pull in configured json library
	json = registry.JSONLibrary()
)
