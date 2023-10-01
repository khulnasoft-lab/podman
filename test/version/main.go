package main

import (
	"fmt"

	"github.com/khulnasoft-lab/podman/v4/version"
)

func main() {
	fmt.Print(version.Version.String())
}
