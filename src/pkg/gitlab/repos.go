package gitlabclient

import (
	"errors"
	"log"
	"strconv"

	gitlab "gitlab.com/gitlab-org/api/client-go"
)

func (gl *GitLab) CreateRepository(input map[string]any) error {
	repoName, ok := input["name"].(string)
	if !ok {
		return errors.New("missing args: name (string)")
	}
	groupIDStr, ok := input["group"].(string)
	if !ok {
		return errors.New("missing args: group (int)")
	}
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		return errors.New("missing args: group (int)")
	}

	p := &gitlab.CreateProjectOptions{
		Name:                     gitlab.Ptr(repoName),
		NamespaceID:              gitlab.Ptr(groupID),
		MergeRequestsAccessLevel: gitlab.Ptr(gitlab.EnabledAccessControl),
		SnippetsAccessLevel:      gitlab.Ptr(gitlab.EnabledAccessControl),
		Visibility:               gitlab.Ptr(gitlab.InternalVisibility),
		InitializeWithReadme:     gitlab.Ptr(true),
	}

	project, res, err := gl.client.Projects.CreateProject(p)
	log.Println(res)
	if err != nil || project == nil {
		return err
	}

	return nil
}
