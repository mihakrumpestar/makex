package docker_compose

import (
	"fmt"
	"makex/internal/helpers"
	"path"
)

type DockerCompose struct {
	subPath string
	target  string
}

func (db *DockerCompose) Build() error {
	return fmt.Errorf("Build function is not callable for DockerCompose orchestrator")
}

func (dc *DockerCompose) Deploy() error {
	err := helpers.ExecStreaming("sudo", "-E", "docker", "compose", "-f", fmt.Sprintf("%s/docker-compose.yml", path.Join(dc.subPath, dc.target)), "-p", dc.target, "up", "-d", "--force-recreate", "-t", "30")
	if err != nil {
		return fmt.Errorf("running Docker compose: %s\n", err)
	} else {
		fmt.Println("Docker compose target deployed")
	}

	return nil
}

func (dc *DockerCompose) Destroy() error {
	err := helpers.ExecStreaming("sudo", "-E", "docker", "compose", "-f", fmt.Sprintf("%s/docker-compose.yml", path.Join(dc.subPath, dc.target)), "-p", dc.target, "down", "-t", "30")
	if err != nil {
		return fmt.Errorf("running Docker compose: %s\n", err)
	} else {
		fmt.Println("Docker compose target down")
	}

	return nil
}
