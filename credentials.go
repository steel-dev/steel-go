package steel

import (
	"context"
	"net/url"
)

type CredentialsResource struct {
	client *Client
}

// List all credential metadata
func (r *CredentialsResource) List(ctx context.Context, params CredentialListParams, opts ...RequestOption) (*CredentialListResponse, error) {
	path := "/v1/credentials"
	query := url.Values{}
	if params.ProjectID != nil {
		query.Set("projectId", *params.ProjectID)
	}
	if params.Namespace != nil {
		query.Set("namespace", *params.Namespace)
	}
	if params.Origin != nil {
		query.Set("origin", *params.Origin)
	}
	if encoded := query.Encode(); encoded != "" {
		path += "?" + encoded
	}
	var out CredentialListResponse
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Stores credentials
func (r *CredentialsResource) Create(ctx context.Context, body CredentialCreateParams, opts ...RequestOption) (*CredentialCreateResponse, error) {
	path := "/v1/credentials"
	var out CredentialCreateResponse
	if err := r.client.do(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Updates credentials
func (r *CredentialsResource) Update(ctx context.Context, body CredentialUpdateParams, opts ...RequestOption) (*CredentialCreateResponse, error) {
	path := "/v1/credentials"
	var out CredentialCreateResponse
	if err := r.client.do(ctx, "PUT", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Deletes credentials
func (r *CredentialsResource) Delete(ctx context.Context, body CredentialDeleteParams, opts ...RequestOption) (*CredentialDeleteResponse, error) {
	path := "/v1/credentials"
	var out CredentialDeleteResponse
	if err := r.client.do(ctx, "DELETE", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}
