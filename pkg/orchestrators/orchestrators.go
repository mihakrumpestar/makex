package orchestrators

import (
	"fmt"
	"slices"
)

// -- Enum -- //

type OrchestratorEnum struct {
	slug string
}

func (o OrchestratorEnum) String() string {
	return o.slug
}

var (
	DockerBuild   = OrchestratorEnum{"docker-build"}
	DockerCompose = OrchestratorEnum{"docker-compose"}
	DockerSwarm   = OrchestratorEnum{"docker-swarm"}
)

func FromStringArray(inputs []string) ([]OrchestratorEnum, error) {
	// The slice to store OrchestratorEnums
	var orchestrators []OrchestratorEnum

	for _, input := range inputs {
		switch input {
		case DockerBuild.slug:
			orchestrators = append(orchestrators, DockerBuild)
		case DockerCompose.slug:
			orchestrators = append(orchestrators, DockerCompose)
		case DockerSwarm.slug:
			if slices.Contains(inputs, DockerCompose.slug) {
				return nil, fmt.Errorf("can't use docker-swarm and docker-compose at the same time")
			}
			orchestrators = append(orchestrators, DockerSwarm)
		default:
			return nil, fmt.Errorf("orchestrator '%s' is not valid", inputs)
		}
	}
	if len(orchestrators) == 0 {
		return nil, fmt.Errorf("no valid orchestrator found in input: %s", inputs)
	}

	return orchestrators, nil
}

// -- Factory -- //

type OrchestratorActions interface {
	build()
	deploy()
	destroy()
}
