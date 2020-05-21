// Code generated by "make api"; DO NOT EDIT.
package scopes

import (
	"context"
	"fmt"

	"github.com/hashicorp/watchtower/api"
)

func (s Organization) CreateProject(ctx context.Context, project *Project) (*Project, *api.Error, error) {
	if s.Client == nil {
		return nil, nil, fmt.Errorf("nil client in CreateProject request")
	}
	if s.Id == "" {

		// Assume the client has been configured with organization already and
		// move on

	} else {
		// If it's explicitly set here, override anything that might be in the
		// client

		ctx = context.WithValue(ctx, "org", s.Id)

	}

	req, err := s.Client.NewRequest(ctx, "POST", "projects", project)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating CreateProject request: %w", err)
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during CreateProject call: %w", err)
	}

	target := new(Project)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding CreateProject repsonse: %w", err)
	}

	target.Client = s.Client.Clone()
	target.Client.SetProject(target.Id)

	return target, apiErr, nil
}