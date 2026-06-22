package steel

import (
	"context"
	"fmt"
	"net/url"
)

type Profiles struct {
	client *Client
}

// List all profiles
func (r *Profiles) List(ctx context.Context, params ProfileListParams, opts ...RequestOption) (*ProfileListResponse, error) {
	path := "/v1/profiles"
	query := url.Values{}
	if params.ProjectID != nil {
		query.Set("projectId", *params.ProjectID)
	}
	if encoded := query.Encode(); encoded != "" {
		path += "?" + encoded
	}
	var out ProfileListResponse
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Create a profile
func (r *Profiles) Create(ctx context.Context, body ProfileCreateParams, opts ...RequestOption) (*ProfileCreateResponse, error) {
	path := "/v1/profiles"
	var out ProfileCreateResponse
	if err := r.client.doMultipart(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get a profile
func (r *Profiles) Get(ctx context.Context, id string, params ProfileGetParams, opts ...RequestOption) (*ProfileCreateResponse, error) {
	path := fmt.Sprintf("/v1/profiles/%s", id)
	query := url.Values{}
	if params.ProjectID != nil {
		query.Set("projectId", *params.ProjectID)
	}
	if encoded := query.Encode(); encoded != "" {
		path += "?" + encoded
	}
	var out ProfileCreateResponse
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a profile
func (r *Profiles) Update(ctx context.Context, id string, body ProfileCreateParams, params ProfileUpdateQueryParams, opts ...RequestOption) (*ProfileCreateResponse, error) {
	path := fmt.Sprintf("/v1/profiles/%s", id)
	query := url.Values{}
	if params.ProjectID != nil {
		query.Set("projectId", *params.ProjectID)
	}
	if encoded := query.Encode(); encoded != "" {
		path += "?" + encoded
	}
	var out ProfileCreateResponse
	if err := r.client.doMultipart(ctx, "PATCH", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}
