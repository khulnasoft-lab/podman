package tunnel

import (
	"context"

	"github.com/khulnasoft-lab/podman/v4/libpod/define"
	"github.com/khulnasoft-lab/podman/v4/pkg/bindings/containers"
	"github.com/khulnasoft-lab/podman/v4/pkg/domain/entities"
)

func (ic *ContainerEngine) HealthCheckRun(ctx context.Context, nameOrID string, options entities.HealthCheckOptions) (*define.HealthCheckResults, error) {
	return containers.RunHealthCheck(ic.ClientCtx, nameOrID, nil)
}
