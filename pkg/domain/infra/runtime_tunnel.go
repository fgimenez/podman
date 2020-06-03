// +build !ABISupport

package infra

import (
	"context"
	"fmt"

	"github.com/containers/libpod/pkg/bindings"
	"github.com/containers/libpod/pkg/domain/entities"
	"github.com/containers/libpod/pkg/domain/infra/tunnel"
)

func NewContainerEngine(facts *entities.PodmanConfig) (entities.ContainerEngine, error) {
	switch facts.EngineMode {
	case entities.ABIMode:
		return nil, fmt.Errorf("direct runtime not supported")
	case entities.TunnelMode:
		ctx, err := bindings.NewConnectionWithIdentity(context.Background(), facts.Uri, facts.PassPhrase, facts.Identities...)
		return &tunnel.ContainerEngine{ClientCxt: ctx}, err
	}
	return nil, fmt.Errorf("runtime mode '%v' is not supported", facts.EngineMode)
}

// NewImageEngine factory provides a libpod runtime for image-related operations
func NewImageEngine(facts *entities.PodmanConfig) (entities.ImageEngine, error) {
	switch facts.EngineMode {
	case entities.ABIMode:
		return nil, fmt.Errorf("direct image runtime not supported")
	case entities.TunnelMode:
		ctx, err := bindings.NewConnectionWithIdentity(context.Background(), facts.Uri, facts.PassPhrase, facts.Identities...)
		return &tunnel.ImageEngine{ClientCxt: ctx}, err
	}
	return nil, fmt.Errorf("runtime mode '%v' is not supported", facts.EngineMode)
}
