package cli

type ImagePullOptions struct {
	LogLevel string
}

type ImagePullReport struct{}

type ImageSaveOptions struct {
	LogLevel string
}

type PodmanEngine interface {
	Pull(rawImage string, opts ImagePullOptions) error
	Save(nameOrID string, tags []string, options ImageSaveOptions) error
}
