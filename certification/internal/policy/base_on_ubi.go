package policy

import (
	"os/exec"
	"strings"
)

type BasedOnUbiPolicy struct {
}

func (p BasedOnUbiPolicy) Validate(image string) (bool, error) {
	stdouterr, err := exec.Command("podman", "run", "-it", image, "cat", "/etc/os-release").CombinedOutput()
	if err != nil {
		return false, err
	}

	lines := strings.Split(string(stdouterr), "\n")

	var hasRHELID, hasRHELName bool
	for _, value := range lines {
		if strings.HasPrefix(value, `ID="rhel"`) {
			hasRHELID = true
		} else if strings.HasPrefix(value, `NAME="Red Hat Enterprise Linux"`) {
			hasRHELName = true
		}
	}
	if hasRHELID && hasRHELName {
		return true, nil
	}

	return false, nil
}

func (p BasedOnUbiPolicy) GetName() string {
	return "BasedOnUbi"
}

func (p BasedOnUbiPolicy) GetMetadata() Metadata {
	return Metadata{
		Description:      "Checking if the container's base image is based on UBI",
		Level:            "best",
		KnowledgeBaseURL: "https://connect.redhat.com/zones/containers/container-certification-policy-guide", // Placeholder
		PolicyURL:        "https://connect.redhat.com/zones/containers/container-certification-policy-guide",
	}
}

func (p BasedOnUbiPolicy) GetHelp() HelpText {
	return HelpText{
		Message:    "It is recommened that your image be based upon the Red Hat Universal Base Image (UBI)",
		Suggestion: "Change the FROM directive in your Dockerfile or Containerfile to FROM registry.access.redhat.com/ubi8/ubi",
	}
}
