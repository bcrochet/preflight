package shell

import (
	"os/exec"

	"github.com/komish/preflight/cli"
)

type PodmanCLIEngine struct{}

func (pe *PodmanCLIEngine) Pull(rawImage string, opts cli.ImagePullOptions) (*cli.ImagePullReport, error) {
	_, _ = exec.Command("podman", "pull", rawImage).CombinedOutput()
	return nil, nil
}

func (pe *PodmanCLIEngine) Save(nameOrID string, tags []string, options cli.ImageSaveOptions) error {
	return nil
}
