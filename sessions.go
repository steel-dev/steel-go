package steel

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type Sessions struct {
	client   *Client
	Captchas *SessionsCaptchas
	Files    *SessionsFiles
}

// List all sessions
func (r *Sessions) List(ctx context.Context, params SessionListParams, opts ...RequestOption) (*CursorPage[SessionslistSession], error) {
	nextCursor := func(items []SessionslistSession, lastID string) string { return lastID }
	var fetch func(ctx context.Context, cursor string) (*CursorPage[SessionslistSession], error)
	fetch = func(ctx context.Context, cursor string) (*CursorPage[SessionslistSession], error) {
		path := "/v1/sessions"
		query := url.Values{}
		if params.CursorID != nil {
			query.Set("cursorId", *params.CursorID)
		}
		if params.Limit != nil {
			query.Set("limit", strconv.FormatInt(*params.Limit, 10))
		}
		if params.Status != nil {
			query.Set("status", fmt.Sprintf("%v", *params.Status))
		}
		if params.ProjectID != nil {
			query.Set("projectId", *params.ProjectID)
		}
		if cursor != "" {
			query.Set("cursorId", cursor)
		}
		if encoded := query.Encode(); encoded != "" {
			path += "?" + encoded
		}
		var out CursorPage[SessionslistSession]
		if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
			return nil, err
		}
		out.fetch = fetch
		out.ctx = ctx
		out.nextCursor = nextCursor
		return &out, nil
	}
	return fetch(ctx, "")
}

func (r *Sessions) ListAutoPaging(ctx context.Context, params SessionListParams, opts ...RequestOption) *CursorPageAutoPager[SessionslistSession] {
	page, err := r.List(ctx, params, opts...)
	return NewCursorPageAutoPager(page, err)
}

// Create a session
func (r *Sessions) Create(ctx context.Context, body SessionCreateParams, opts ...RequestOption) (*Session, error) {
	path := "/v1/sessions"
	var out Session
	if err := r.client.do(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Release all sessions
func (r *Sessions) ReleaseAll(ctx context.Context, body SessionReleaseAllParams, params SessionReleaseAllQueryParams, opts ...RequestOption) (*SessionReleaseAllResponse, error) {
	path := "/v1/sessions/release"
	query := url.Values{}
	if params.ProjectID != nil {
		query.Set("projectId", *params.ProjectID)
	}
	if encoded := query.Encode(); encoded != "" {
		path += "?" + encoded
	}
	var out SessionReleaseAllResponse
	if err := r.client.do(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get session details
func (r *Sessions) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Session, error) {
	path := fmt.Sprintf("/v1/sessions/%s", url.PathEscape(id))
	var out Session
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get session agent traces
func (r *Sessions) AgentTraces(ctx context.Context, id string, params SessionAgentTracesParams, opts ...RequestOption) (*SessionAgentTracesResponse, error) {
	path := fmt.Sprintf("/v1/sessions/%s/agent-traces", url.PathEscape(id))
	query := url.Values{}
	if params.Namespace != nil {
		query.Set("namespace", *params.Namespace)
	}
	if params.StartTime != nil {
		query.Set("startTime", fmt.Sprintf("%v", *params.StartTime))
	}
	if params.EndTime != nil {
		query.Set("endTime", fmt.Sprintf("%v", *params.EndTime))
	}
	if params.EventTypes != nil {
		for _, __v := range *params.EventTypes {
			query.Add("eventTypes", __v)
		}
	}
	if encoded := query.Encode(); encoded != "" {
		path += "?" + encoded
	}
	var out SessionAgentTracesResponse
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get session context
func (r *Sessions) Context(ctx context.Context, id string, opts ...RequestOption) (*SessionContext, error) {
	path := fmt.Sprintf("/v1/sessions/%s/context", url.PathEscape(id))
	var out SessionContext
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get session cost
func (r *Sessions) Cost(ctx context.Context, id string, opts ...RequestOption) (*SessionCostResponse, error) {
	path := fmt.Sprintf("/v1/sessions/%s/cost", url.PathEscape(id))
	var out SessionCostResponse
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get recorded events
func (r *Sessions) Events(ctx context.Context, id string, params SessionEventsParams, opts ...RequestOption) (*SessionEventsResponse, error) {
	path := fmt.Sprintf("/v1/sessions/%s/events", url.PathEscape(id))
	query := url.Values{}
	if params.Compressed != nil {
		query.Set("compressed", strconv.FormatBool(*params.Compressed))
	}
	if params.Limit != nil {
		query.Set("limit", strconv.FormatInt(*params.Limit, 10))
	}
	if params.Pointer != nil {
		query.Set("pointer", *params.Pointer)
	}
	if encoded := query.Encode(); encoded != "" {
		path += "?" + encoded
	}
	var out SessionEventsResponse
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get live session details
func (r *Sessions) LiveDetails(ctx context.Context, id string, opts ...RequestOption) (*SessionLiveDetailsResponse, error) {
	path := fmt.Sprintf("/v1/sessions/%s/live-details", url.PathEscape(id))
	var out SessionLiveDetailsResponse
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Release a session
func (r *Sessions) Release(ctx context.Context, id string, body SessionReleaseParams, opts ...RequestOption) (*SessionReleaseResponse, error) {
	path := fmt.Sprintf("/v1/sessions/%s/release", url.PathEscape(id))
	var out SessionReleaseResponse
	if err := r.client.do(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Execute computer action
func (r *Sessions) Computer(ctx context.Context, sessionID string, body SessionComputerParams, opts ...RequestOption) (*SessionComputerResponse, error) {
	path := fmt.Sprintf("/v1/sessions/%s/computer", url.PathEscape(sessionID))
	var out SessionComputerResponse
	if err := r.client.do(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

type SessionsCaptchas struct {
	client *Client
}

// Solve captcha(s)
func (r *SessionsCaptchas) Solve(ctx context.Context, sessionID string, body CaptchaSolveParams, opts ...RequestOption) (*CaptchaSolveResponse, error) {
	path := fmt.Sprintf("/v1/sessions/%s/captchas/solve", url.PathEscape(sessionID))
	var out CaptchaSolveResponse
	if err := r.client.do(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Solve image captcha
func (r *SessionsCaptchas) SolveImage(ctx context.Context, sessionID string, body CaptchaSolveImageParams, opts ...RequestOption) (*CaptchaSolveImageResponse, error) {
	path := fmt.Sprintf("/v1/sessions/%s/captchas/solve-image", url.PathEscape(sessionID))
	var out CaptchaSolveImageResponse
	if err := r.client.do(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get captcha status
func (r *SessionsCaptchas) Status(ctx context.Context, sessionID string, opts ...RequestOption) (*CaptchaStatusResponse, error) {
	path := fmt.Sprintf("/v1/sessions/%s/captchas/status", url.PathEscape(sessionID))
	var out CaptchaStatusResponse
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

type SessionsFiles struct {
	client *Client
}

// List files
func (r *SessionsFiles) List(ctx context.Context, sessionID string, opts ...RequestOption) (*Fileslist, error) {
	path := fmt.Sprintf("/v1/sessions/%s/files", url.PathEscape(sessionID))
	var out Fileslist
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Upload a file
func (r *SessionsFiles) Upload(ctx context.Context, sessionID string, body FileUploadParams2, opts ...RequestOption) (*File, error) {
	path := fmt.Sprintf("/v1/sessions/%s/files", url.PathEscape(sessionID))
	var out File
	if err := r.client.doMultipart(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete all files
func (r *SessionsFiles) DeleteAll(ctx context.Context, sessionID string, opts ...RequestOption) error {
	path := fmt.Sprintf("/v1/sessions/%s/files", url.PathEscape(sessionID))
	return r.client.do(ctx, "DELETE", path, nil, nil, opts...)
}

// Download archive
func (r *SessionsFiles) DownloadArchive(ctx context.Context, sessionID string, opts ...RequestOption) ([]byte, error) {
	path := fmt.Sprintf("/v1/sessions/%s/files.zip", url.PathEscape(sessionID))
	data, err := r.client.doRaw(ctx, "GET", path, nil, opts...)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Download a file
func (r *SessionsFiles) Download(ctx context.Context, sessionID string, path string, opts ...RequestOption) ([]byte, error) {
	reqPath := fmt.Sprintf("/v1/sessions/%s/files/%s", url.PathEscape(sessionID), path)
	data, err := r.client.doRaw(ctx, "GET", reqPath, nil, opts...)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Delete a file
func (r *SessionsFiles) Delete(ctx context.Context, sessionID string, path string, opts ...RequestOption) error {
	reqPath := fmt.Sprintf("/v1/sessions/%s/files/%s", url.PathEscape(sessionID), path)
	return r.client.do(ctx, "DELETE", reqPath, nil, nil, opts...)
}
