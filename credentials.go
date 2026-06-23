package steel

import (
	"context"
	"github.com/steel-dev/steel-go/internal/apijson"
	"github.com/steel-dev/steel-go/internal/param"
	"net/url"
	"time"
)

type Credentials struct {
	client *Client
}

// List all credential metadata
func (r *Credentials) List(ctx context.Context, params CredentialListParams, opts ...RequestOption) (*CredentialListResponse, error) {
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
func (r *Credentials) Create(ctx context.Context, body CredentialCreateParams, opts ...RequestOption) (*CredentialCreateResponse, error) {
	path := "/v1/credentials"
	var out CredentialCreateResponse
	if err := r.client.do(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Updates credentials
func (r *Credentials) Update(ctx context.Context, body CredentialUpdateParams, opts ...RequestOption) (*CredentialUpdateResponse, error) {
	path := "/v1/credentials"
	var out CredentialUpdateResponse
	if err := r.client.do(ctx, "PUT", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Deletes credentials
func (r *Credentials) Delete(ctx context.Context, body CredentialDeleteParams, opts ...RequestOption) (*CredentialDeleteResponse, error) {
	path := "/v1/credentials"
	var out CredentialDeleteResponse
	if err := r.client.do(ctx, "DELETE", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

type CredentialCreateParams struct {
	// Label for the credential
	Label param.Field[string] `json:"label"`
	// The namespace the credential is stored against. Defaults to "default".
	Namespace param.Field[string] `json:"namespace"`
	// Website origin the credential is for
	Origin param.Field[string] `json:"origin"`
	// Project to store the credential in.
	ProjectID param.Field[string] `json:"projectId"`
	// Value for the credential
	Value param.Field[map[string]string] `json:"value" api:"required"`
}

func (r CredentialCreateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CredentialListResponse struct {
	Credentials []CredentialListResponseCredential `json:"credentials" api:"required"`
	JSON        credentialListResponseJSON         `json:"-"`
}

// credentialListResponseJSON contains the JSON metadata for the struct [CredentialListResponse]
type credentialListResponseJSON struct {
	Credentials apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CredentialListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r credentialListResponseJSON) RawJSON() string {
	return r.raw
}

type CredentialCreateResponse struct {
	// Date and time the credential was created
	CreatedAt time.Time `json:"createdAt" api:"required"`
	// Label for the credential
	Label string `json:"label"`
	// The namespace the credential is stored against. Defaults to "default".
	Namespace string `json:"namespace"`
	// Website origin the credential is for
	Origin string `json:"origin"`
	// Date and time the credential was last updated
	UpdatedAt time.Time                    `json:"updatedAt" api:"required"`
	JSON      credentialCreateResponseJSON `json:"-"`
}

// credentialCreateResponseJSON contains the JSON metadata for the struct [CredentialCreateResponse]
type credentialCreateResponseJSON struct {
	CreatedAt   apijson.Field
	Label       apijson.Field
	Namespace   apijson.Field
	Origin      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CredentialCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r credentialCreateResponseJSON) RawJSON() string {
	return r.raw
}

type CredentialDeleteParams struct {
	// The namespace the credential is stored against. Defaults to "default".
	Namespace param.Field[string] `json:"namespace"`
	// Website origin the credential is for
	Origin param.Field[string] `json:"origin" api:"required"`
	// Project to delete the credential from.
	ProjectID param.Field[string] `json:"projectId"`
}

func (r CredentialDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CredentialUpdateParams struct {
	// Label for the credential
	Label param.Field[string] `json:"label"`
	// The namespace the credential is stored against. Defaults to "default".
	Namespace param.Field[string] `json:"namespace"`
	// Website origin the credential is for
	Origin param.Field[string] `json:"origin"`
	// Project to update the credential in.
	ProjectID param.Field[string] `json:"projectId"`
	// Value for the credential
	Value param.Field[map[string]string] `json:"value"`
}

func (r CredentialUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CredentialDeleteResponse struct {
	Success bool                         `json:"success" api:"required"`
	JSON    credentialDeleteResponseJSON `json:"-"`
}

// credentialDeleteResponseJSON contains the JSON metadata for the struct [CredentialDeleteResponse]
type credentialDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CredentialDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r credentialDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CredentialUpdateResponse struct {
	// Date and time the credential was created
	CreatedAt time.Time `json:"createdAt" api:"required"`
	// Label for the credential
	Label string `json:"label"`
	// The namespace the credential is stored against. Defaults to "default".
	Namespace string `json:"namespace"`
	// Website origin the credential is for
	Origin string `json:"origin"`
	// Date and time the credential was last updated
	UpdatedAt time.Time                    `json:"updatedAt" api:"required"`
	JSON      credentialUpdateResponseJSON `json:"-"`
}

// credentialUpdateResponseJSON contains the JSON metadata for the struct [CredentialUpdateResponse]
type credentialUpdateResponseJSON struct {
	CreatedAt   apijson.Field
	Label       apijson.Field
	Namespace   apijson.Field
	Origin      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CredentialUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r credentialUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type CredentialListResponseCredential struct {
	// Date and time the credential was created
	CreatedAt time.Time `json:"createdAt" api:"required"`
	// Label for the credential
	Label string `json:"label"`
	// The namespace the credential is stored against. Defaults to "default".
	Namespace string `json:"namespace"`
	// Website origin the credential is for
	Origin string `json:"origin"`
	// Date and time the credential was last updated
	UpdatedAt time.Time                            `json:"updatedAt" api:"required"`
	JSON      credentialListResponseCredentialJSON `json:"-"`
}

// credentialListResponseCredentialJSON contains the JSON metadata for the struct [CredentialListResponseCredential]
type credentialListResponseCredentialJSON struct {
	CreatedAt   apijson.Field
	Label       apijson.Field
	Namespace   apijson.Field
	Origin      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CredentialListResponseCredential) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r credentialListResponseCredentialJSON) RawJSON() string {
	return r.raw
}

type CredentialListParams struct {
	ProjectID *string `json:"projectId,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Origin    *string `json:"origin,omitempty"`
}
