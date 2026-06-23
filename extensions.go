package steel

import (
	"context"
	"fmt"
	"net/url"
)

type Extensions struct {
	client *Client
}

// List extensions
func (r *Extensions) List(ctx context.Context, opts ...RequestOption) (*ExtensionListResponse, error) {
	path := "/v1/extensions"
	var out ExtensionListResponse
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Upload extension
func (r *Extensions) Upload(ctx context.Context, body ExtensionUploadParams, opts ...RequestOption) (*ExtensionUploadResponse, error) {
	path := "/v1/extensions"
	var out ExtensionUploadResponse
	if err := r.client.doMultipart(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete all extensions
func (r *Extensions) DeleteAll(ctx context.Context, opts ...RequestOption) (*ExtensionDeleteAllResponse, error) {
	path := "/v1/extensions"
	var out ExtensionDeleteAllResponse
	if err := r.client.do(ctx, "DELETE", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Download extension
func (r *Extensions) Download(ctx context.Context, extensionID string, opts ...RequestOption) ([]byte, error) {
	path := fmt.Sprintf("/v1/extensions/%s", url.PathEscape(extensionID))
	data, err := r.client.doRaw(ctx, "GET", path, nil, opts...)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Update extension
func (r *Extensions) Update(ctx context.Context, extensionID string, body ExtensionUpdateParams, opts ...RequestOption) (*ExtensionUpdateResponse, error) {
	path := fmt.Sprintf("/v1/extensions/%s", url.PathEscape(extensionID))
	var out ExtensionUpdateResponse
	if err := r.client.doMultipart(ctx, "PUT", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete extension
func (r *Extensions) Delete(ctx context.Context, extensionID string, opts ...RequestOption) (*ExtensionDeleteResponse, error) {
	path := fmt.Sprintf("/v1/extensions/%s", url.PathEscape(extensionID))
	var out ExtensionDeleteResponse
	if err := r.client.do(ctx, "DELETE", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}
