package steel

import (
	"context"
	"fmt"
	"github.com/steel-dev/steel-go/internal/apijson"
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

type ExtensionDeleteAllResponse struct {
	Message string                         `json:"message" api:"required"`
	JSON    extensionDeleteAllResponseJSON `json:"-"`
}

// extensionDeleteAllResponseJSON contains the JSON metadata for the struct [ExtensionDeleteAllResponse]
type extensionDeleteAllResponseJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtensionDeleteAllResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extensionDeleteAllResponseJSON) RawJSON() string {
	return r.raw
}

// ExtensionListResponse Response containing a list of extensions for the organization
type ExtensionListResponse struct {
	// Total number of extensions
	Count float64 `json:"count" api:"required"`
	// List of extensions for the organization
	Extensions []ExtensionListResponseExtension `json:"extensions" api:"required"`
	JSON       extensionListResponseJSON        `json:"-"`
}

// extensionListResponseJSON contains the JSON metadata for the struct [ExtensionListResponse]
type extensionListResponseJSON struct {
	Count       apijson.Field
	Extensions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtensionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extensionListResponseJSON) RawJSON() string {
	return r.raw
}

type ExtensionUploadResponse struct {
	// Creation timestamp
	CreatedAt string `json:"createdAt" api:"required"`
	// Unique extension identifier (e.g., ext_12345)
	ID string `json:"id" api:"required"`
	// Extension name
	Name string `json:"name" api:"required"`
	// Last update timestamp
	UpdatedAt string                      `json:"updatedAt" api:"required"`
	JSON      extensionUploadResponseJSON `json:"-"`
}

// extensionUploadResponseJSON contains the JSON metadata for the struct [ExtensionUploadResponse]
type extensionUploadResponseJSON struct {
	CreatedAt   apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtensionUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extensionUploadResponseJSON) RawJSON() string {
	return r.raw
}

type ExtensionUploadParams struct {
	// Extension .zip/.crx file
	File *FileUpload `json:"file,omitempty"`
	// Extension URL
	URL *string `json:"url,omitempty"`
}

type ExtensionUpdateParams struct {
	// Extension .zip/.crx file
	File *FileUpload `json:"file,omitempty"`
	// Extension URL
	URL *string `json:"url,omitempty"`
}

type ExtensionUpdateResponse struct {
	// Creation timestamp
	CreatedAt string `json:"createdAt" api:"required"`
	// Unique extension identifier (e.g., ext_12345)
	ID string `json:"id" api:"required"`
	// Extension name
	Name string `json:"name" api:"required"`
	// Last update timestamp
	UpdatedAt string                      `json:"updatedAt" api:"required"`
	JSON      extensionUpdateResponseJSON `json:"-"`
}

// extensionUpdateResponseJSON contains the JSON metadata for the struct [ExtensionUpdateResponse]
type extensionUpdateResponseJSON struct {
	CreatedAt   apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtensionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extensionUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ExtensionDeleteResponse struct {
	Message string                      `json:"message" api:"required"`
	JSON    extensionDeleteResponseJSON `json:"-"`
}

// extensionDeleteResponseJSON contains the JSON metadata for the struct [ExtensionDeleteResponse]
type extensionDeleteResponseJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtensionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extensionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// ExtensionListResponseExtension List of extensions for the organization
type ExtensionListResponseExtension struct {
	// Creation timestamp
	CreatedAt string `json:"createdAt" api:"required"`
	// Unique extension identifier (e.g., ext_12345)
	ID string `json:"id" api:"required"`
	// Extension name
	Name string `json:"name" api:"required"`
	// Last update timestamp
	UpdatedAt string                             `json:"updatedAt" api:"required"`
	JSON      extensionListResponseExtensionJSON `json:"-"`
}

// extensionListResponseExtensionJSON contains the JSON metadata for the struct [ExtensionListResponseExtension]
type extensionListResponseExtensionJSON struct {
	CreatedAt   apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtensionListResponseExtension) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extensionListResponseExtensionJSON) RawJSON() string {
	return r.raw
}
