package abi

import (
	"context"

	"github.com/khulnasoft-lab/podman/v4/pkg/autoupdate"
	"github.com/khulnasoft-lab/podman/v4/pkg/domain/entities"
)

func (ic *ContainerEngine) AutoUpdate(ctx context.Context, options entities.AutoUpdateOptions) ([]*entities.AutoUpdateReport, []error) {
	return autoupdate.AutoUpdate(ctx, ic.Libpod, options)
}
