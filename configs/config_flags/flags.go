package config_flags

import (
	"makex/pkg/orchestrators"
	"makex/pkg/secrets"
)

type FlagsStruct struct {
	Subfolder       *string
	Orchestrator    *[]orchestrators.OrchestratorEnum
	Secrets         *secrets.SecretEnum
	Target          *string
	MultipleTargets *bool
	Environment     *string
}

var Flags = &FlagsStruct{
	Subfolder:       new(string),
	Orchestrator:    new([]orchestrators.OrchestratorEnum),
	Secrets:         new(secrets.SecretEnum),
	Target:          new(string),
	MultipleTargets: new(bool),
	Environment:     new(string),
}
