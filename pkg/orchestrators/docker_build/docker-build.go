package docker_build

import (
	"fmt"
	"makex/internal/helpers"
	"os"
	"path"
)

type DockerBuild struct {
	subPath string
	target  string
}

func (db *DockerBuild) Build() error {
	// sudo docker build -t "$MAIN_DOMAIN/$TARGET:latest" "./$TARGET/"

	TAG := fmt.Sprintf("%s/%s:latest", os.Getenv("MAIN_DOMAIN"), db.target)

	err := helpers.ExecStreaming("sudo", "-E", "docker", "build", "-t", TAG, path.Join(db.subPath, db.target))
	if err != nil {
		return fmt.Errorf("running Docker Build: %s\n", err)
	} else {
		fmt.Println("Docker build target build")
	}

	// sudo docker push "$MAIN_DOMAIN/$TARGET:latest"

	return nil
}

func (db *DockerBuild) Deploy() error {
	return fmt.Errorf("Deploy function is not callable for DockerBuild orchestrator")
}

func (db *DockerBuild) Destroy() error {
	return fmt.Errorf("Destroy function is not callable for DockerBuild orchestrator")
}
