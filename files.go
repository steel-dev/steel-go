package steel

import (
	"context"
	"fmt"
)

type Files struct {
	client *Client
}

// List global files
func (r *Files) List(ctx context.Context, opts ...RequestOption) (*Fileslist, error) {
	path := "/v1/files"
	var out Fileslist
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Upload a global file
func (r *Files) Upload(ctx context.Context, body FileUploadParams, opts ...RequestOption) (*File, error) {
	path := "/v1/files"
	var out File
	if err := r.client.doMultipart(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Download a global file
func (r *Files) Download(ctx context.Context, path string, opts ...RequestOption) ([]byte, error) {
	reqPath := fmt.Sprintf("/v1/files/%s", path)
	data, err := r.client.doRaw(ctx, "GET", reqPath, nil, opts...)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Delete a global file
func (r *Files) Delete(ctx context.Context, path string, opts ...RequestOption) error {
	reqPath := fmt.Sprintf("/v1/files/%s", path)
	return r.client.do(ctx, "DELETE", reqPath, nil, nil, opts...)
}

type FileUploadParams struct {
	// The file to upload (binary) or URL string to download from
	File FileUpload `json:"file"`
	// Path to the file in the storage system
	Path *string `json:"path,omitempty"`
}
