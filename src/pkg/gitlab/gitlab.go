package gitlabclient

import (
	"os"

	gitlab "gitlab.com/gitlab-org/api/client-go"
)

type GitLab struct {
	client *gitlab.Client
}

func NewGitLab() (*GitLab, error) {
	gl, err := gitlab.NewClient(
		os.Getenv("GITLAB_TOKEN"),
		gitlab.WithBaseURL(os.Getenv("GITLAB_URL")),
	)
	if err != nil {
		return &GitLab{}, err
	}
	client := GitLab{
		client: gl,
	}
	return &client, nil
}
