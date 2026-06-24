package steel

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/steel-dev/steel-go/internal/apijson"
	"github.com/steel-dev/steel-go/internal/param"
	"io"
	"net/url"
	"strconv"
	"time"
)

type SessionService struct {
	client   *Client
	Captchas *SessionServiceCaptchas
	Files    *SessionServiceFiles
}

// List all sessions
func (r *SessionService) List(ctx context.Context, params SessionListParams, opts ...RequestOption) (*CursorPage[SessionListSession], error) {
	nextCursor := func(items []SessionListSession, lastID string) string { return lastID }
	var fetch func(ctx context.Context, cursor string) (*CursorPage[SessionListSession], error)
	fetch = func(ctx context.Context, cursor string) (*CursorPage[SessionListSession], error) {
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
		var raw map[string]json.RawMessage
		if err := r.client.do(ctx, "GET", path, nil, &raw, opts...); err != nil {
			return nil, err
		}
		items, cursorVal, hasMore, derr := decodeCursorPage[SessionListSession](raw, "sessions", "nextCursor", "")
		if derr != nil {
			return nil, derr
		}
		out := CursorPage[SessionListSession]{Items: items, cursor: cursorVal, hasMore: hasMore, fetch: fetch, ctx: ctx, nextCursor: nextCursor}
		return &out, nil
	}
	return fetch(ctx, "")
}

func (r *SessionService) ListAutoPaging(ctx context.Context, params SessionListParams, opts ...RequestOption) *CursorPageAutoPager[SessionListSession] {
	page, err := r.List(ctx, params, opts...)
	return NewCursorPageAutoPager(page, err)
}

// Create a session
func (r *SessionService) Create(ctx context.Context, body SessionCreateParams, opts ...RequestOption) (*Session, error) {
	path := "/v1/sessions"
	var out Session
	if err := r.client.do(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Release all sessions
func (r *SessionService) ReleaseAll(ctx context.Context, body SessionReleaseAllParams, params SessionReleaseAllQueryParams, opts ...RequestOption) (*SessionReleaseAllResponse, error) {
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
func (r *SessionService) Retrieve(ctx context.Context, id string, opts ...RequestOption) (*Session, error) {
	path := fmt.Sprintf("/v1/sessions/%s", url.PathEscape(id))
	var out Session
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get session agent traces
func (r *SessionService) AgentTraces(ctx context.Context, id string, params SessionAgentTracesParams, opts ...RequestOption) (*SessionAgentTracesResponse, error) {
	path := fmt.Sprintf("/v1/sessions/%s/agent-traces", url.PathEscape(id))
	query := url.Values{}
	if params.Namespace != nil {
		query.Set("namespace", *params.Namespace)
	}
	if params.StartTime != nil {
		query.Set("startTime", (*params.StartTime).Format(time.RFC3339))
	}
	if params.EndTime != nil {
		query.Set("endTime", (*params.EndTime).Format(time.RFC3339))
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
func (r *SessionService) Context(ctx context.Context, id string, opts ...RequestOption) (*SessionContext, error) {
	path := fmt.Sprintf("/v1/sessions/%s/context", url.PathEscape(id))
	var out SessionContext
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get session cost
func (r *SessionService) Cost(ctx context.Context, id string, opts ...RequestOption) (*SessionCostResponse, error) {
	path := fmt.Sprintf("/v1/sessions/%s/cost", url.PathEscape(id))
	var out SessionCostResponse
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get recorded events
func (r *SessionService) Events(ctx context.Context, id string, params SessionEventsParams, opts ...RequestOption) (*SessionEventsResponse, error) {
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
func (r *SessionService) LiveDetails(ctx context.Context, id string, opts ...RequestOption) (*SessionLiveDetailsResponse, error) {
	path := fmt.Sprintf("/v1/sessions/%s/live-details", url.PathEscape(id))
	var out SessionLiveDetailsResponse
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Release a session
func (r *SessionService) Release(ctx context.Context, id string, body SessionReleaseParams, opts ...RequestOption) (*SessionReleaseResponse, error) {
	path := fmt.Sprintf("/v1/sessions/%s/release", url.PathEscape(id))
	var out SessionReleaseResponse
	if err := r.client.do(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Execute computer action
func (r *SessionService) Computer(ctx context.Context, sessionID string, body SessionComputerParams, opts ...RequestOption) (*SessionComputerResponse, error) {
	path := fmt.Sprintf("/v1/sessions/%s/computer", url.PathEscape(sessionID))
	var out SessionComputerResponse
	if err := r.client.do(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

type SessionServiceCaptchas struct {
	client *Client
}

// Solve captcha(s)
func (r *SessionServiceCaptchas) Solve(ctx context.Context, sessionID string, body CaptchaSolveParams, opts ...RequestOption) (*CaptchaSolveResponse, error) {
	path := fmt.Sprintf("/v1/sessions/%s/captchas/solve", url.PathEscape(sessionID))
	var out CaptchaSolveResponse
	if err := r.client.do(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Solve image captcha
func (r *SessionServiceCaptchas) SolveImage(ctx context.Context, sessionID string, body CaptchaSolveImageParams, opts ...RequestOption) (*CaptchaSolveImageResponse, error) {
	path := fmt.Sprintf("/v1/sessions/%s/captchas/solve-image", url.PathEscape(sessionID))
	var out CaptchaSolveImageResponse
	if err := r.client.do(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get captcha status
func (r *SessionServiceCaptchas) Status(ctx context.Context, sessionID string, opts ...RequestOption) (*CaptchaStatusResponse, error) {
	path := fmt.Sprintf("/v1/sessions/%s/captchas/status", url.PathEscape(sessionID))
	var out CaptchaStatusResponse
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

type SessionServiceFiles struct {
	client *Client
}

// List files
func (r *SessionServiceFiles) List(ctx context.Context, sessionID string, opts ...RequestOption) (*FileList, error) {
	path := fmt.Sprintf("/v1/sessions/%s/files", url.PathEscape(sessionID))
	var out FileList
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Upload a file
func (r *SessionServiceFiles) Upload(ctx context.Context, sessionID string, body SessionFileUploadParams, opts ...RequestOption) (*File, error) {
	path := fmt.Sprintf("/v1/sessions/%s/files", url.PathEscape(sessionID))
	var out File
	if err := r.client.doMultipart(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Delete all files
func (r *SessionServiceFiles) DeleteAll(ctx context.Context, sessionID string, opts ...RequestOption) error {
	path := fmt.Sprintf("/v1/sessions/%s/files", url.PathEscape(sessionID))
	return r.client.do(ctx, "DELETE", path, nil, nil, opts...)
}

// Download archive
func (r *SessionServiceFiles) DownloadArchive(ctx context.Context, sessionID string, opts ...RequestOption) (io.ReadCloser, error) {
	path := fmt.Sprintf("/v1/sessions/%s/files.zip", url.PathEscape(sessionID))
	stream, err := r.client.doRawStream(ctx, "GET", path, nil, opts...)
	if err != nil {
		return nil, err
	}
	return stream, nil
}

// Download a file
func (r *SessionServiceFiles) Download(ctx context.Context, sessionID string, path string, opts ...RequestOption) (io.ReadCloser, error) {
	reqPath := fmt.Sprintf("/v1/sessions/%s/files/%s", url.PathEscape(sessionID), encodePath(path))
	stream, err := r.client.doRawStream(ctx, "GET", reqPath, nil, opts...)
	if err != nil {
		return nil, err
	}
	return stream, nil
}

// Delete a file
func (r *SessionServiceFiles) Delete(ctx context.Context, sessionID string, path string, opts ...RequestOption) error {
	reqPath := fmt.Sprintf("/v1/sessions/%s/files/%s", url.PathEscape(sessionID), encodePath(path))
	return r.client.do(ctx, "DELETE", reqPath, nil, nil, opts...)
}

type CaptchaSolveResponse struct {
	// Response message
	Message string `json:"message"`
	// Whether the action was successful
	Success bool                     `json:"success" api:"required"`
	JSON    captchaSolveResponseJSON `json:"-"`
}

// captchaSolveResponseJSON contains the JSON metadata for the struct [CaptchaSolveResponse]
type captchaSolveResponseJSON struct {
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptchaSolveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captchaSolveResponseJSON) RawJSON() string {
	return r.raw
}

type SessionComputerResponse struct {
	// Base64 encoded screenshot if requested
	Base64Image string `json:"base64_image"`
	// Error message if action failed
	Error string `json:"error"`
	// Output message from the action
	Output string `json:"output"`
	// System information
	System string                      `json:"system"`
	JSON   sessionComputerResponseJSON `json:"-"`
}

// sessionComputerResponseJSON contains the JSON metadata for the struct [SessionComputerResponse]
type sessionComputerResponseJSON struct {
	Base64Image apijson.Field
	Error       apijson.Field
	Output      apijson.Field
	System      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionComputerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionComputerResponseJSON) RawJSON() string {
	return r.raw
}

// SessionCreateParams Request body schema for creating a new browser session.
type SessionCreateParams struct {
	// Block ads in the browser session. Default is false.
	BlockAds param.Field[bool] `json:"blockAds"`
	// PEM-encoded root CA certificates to trust in this session. THIS IS CURRENTLY AN EXPERIMENTAL FEATURE.
	CaCertificates param.Field[[]string] `json:"caCertificates"`
	// Number of sessions to create concurrently (check your plan limit)
	Concurrency param.Field[int64] `json:"concurrency"`
	// Configuration for session credentials
	Credentials param.Field[SessionCreateParamsCredentials] `json:"credentials"`
	// Configuration for the debug URL and session viewer. Controls interaction capabilities, cursor visibility, and other debug-related settings.
	DebugConfig param.Field[SessionCreateParamsDebugConfig] `json:"debugConfig"`
	// Device configuration for the session. Specify 'mobile' for mobile device fingerprints and configurations.
	DeviceConfig param.Field[SessionCreateParamsDeviceConfig] `json:"deviceConfig"`
	// Viewport and browser window dimensions for the session. Mobile sessions require dimensions of at least 508x1074; smaller mobile dimensions are rejected with a 400 response.
	Dimensions param.Field[SessionCreateParamsDimensions] `json:"dimensions"`
	// Enable experimental features for the session.
	ExperimentalFeatures param.Field[[]string] `json:"experimentalFeatures"`
	// Array of extension IDs to install in the session. Use ['all_ext'] to install all uploaded extensions.
	ExtensionIDs param.Field[[]string] `json:"extensionIds"`
	// Launch the browser in fullscreen mode, covering the full screen with no Chrome UI. Default is false.
	Fullscreen param.Field[bool] `json:"fullscreen"`
	// Enable headless browser mode (disable Headful mode)
	Headless param.Field[bool] `json:"headless"`
	// Inactivity timeout in milliseconds. When set, the session is released if no CDP command or remote input is received for this duration, even if `timeout` has not yet elapsed. Note that `timeout` remains the hard cap on session lifetime: if `inactivityTimeout` is greater than or equal to the effective `timeout`, it has no effect since `timeout` always elapses first. Omit to disable.
	InactivityTimeout param.Field[int64] `json:"inactivityTimeout"`
	// Enable Selenium mode for the browser session (default is false). Use this when you plan to connect to the browser session via Selenium.
	IsSelenium param.Field[bool] `json:"isSelenium"`
	// The namespace the session should be created against. Defaults to "default".
	Namespace param.Field[string] `json:"namespace"`
	// Enable bandwidth optimizations. Passing true enables all flags (except hosts/patterns). Object allows granular control.
	OptimizeBandwidth param.Field[SessionCreateParamsOptimizeBandwidth] `json:"optimizeBandwidth"`
	// This flag will persist the user profile for the session.
	PersistProfile param.Field[bool] `json:"persistProfile"`
	// This flag will set the profile for the session.
	ProfileID param.Field[string] `json:"profileId"`
	// The project to create the session in. When provided, the session namespace is resolved from the project.
	ProjectID param.Field[string] `json:"projectId"`
	// Custom proxy URL for the browser session. Overrides useProxy, disabling Steel-provided proxies in favor of your specified proxy. Format: http(s)://username:password@hostname:port
	ProxyURL param.Field[string] `json:"proxyUrl"`
	// The desired region for the session. Available: us-east, us-west, us-central, eu-west, eu-central, ap-northeast, ap-southeast, sa-east. Legacy codes (iad, lax, ord) are also accepted.
	Region param.Field[string] `json:"region"`
	// Session context data to be used in the created session. Sessions will start with an empty context by default.
	SessionContext param.Field[SessionCreateParamsSessionContext] `json:"sessionContext"`
	// Optional custom UUID for the session
	SessionID param.Field[string] `json:"sessionId"`
	// Enable automatic captcha solving. Default is false.
	SolveCaptcha param.Field[bool] `json:"solveCaptcha"`
	// Stealth configuration for the session
	StealthConfig param.Field[SessionCreateParamsStealthConfig] `json:"stealthConfig"`
	// Session timeout duration in milliseconds. Default is 300000 (5 minutes).
	Timeout  param.Field[int64]                       `json:"timeout"`
	UseProxy param.Field[SessionCreateParamsUseProxy] `json:"useProxy"`
	// Custom user agent string for the browser session
	UserAgent param.Field[string] `json:"userAgent"`
}

func (r SessionCreateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SessionAgentTracesResponse struct {
	Events  []map[string]interface{}       `json:"events" api:"required"`
	HasMore bool                           `json:"hasMore" api:"required"`
	Total   int64                          `json:"total" api:"required"`
	JSON    sessionAgentTracesResponseJSON `json:"-"`
}

// sessionAgentTracesResponseJSON contains the JSON metadata for the struct [SessionAgentTracesResponse]
type sessionAgentTracesResponseJSON struct {
	Events      apijson.Field
	HasMore     apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionAgentTracesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionAgentTracesResponseJSON) RawJSON() string {
	return r.raw
}

// SessionReleaseResponse Response for releasing a single session.
type SessionReleaseResponse struct {
	// Details about the outcome of the release operation
	Message string `json:"message" api:"required"`
	// Indicates if the session was successfully released
	Success bool                       `json:"success" api:"required"`
	JSON    sessionReleaseResponseJSON `json:"-"`
}

// sessionReleaseResponseJSON contains the JSON metadata for the struct [SessionReleaseResponse]
type sessionReleaseResponseJSON struct {
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionReleaseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionReleaseResponseJSON) RawJSON() string {
	return r.raw
}

// SessionReleaseAllResponse Response for releasing multiple sessions.
type SessionReleaseAllResponse struct {
	// Details about the outcome of the release operation
	Message string `json:"message" api:"required"`
	// Indicates if the sessions were successfully released
	Success bool                          `json:"success" api:"required"`
	JSON    sessionReleaseAllResponseJSON `json:"-"`
}

// sessionReleaseAllResponseJSON contains the JSON metadata for the struct [SessionReleaseAllResponse]
type sessionReleaseAllResponseJSON struct {
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionReleaseAllResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionReleaseAllResponseJSON) RawJSON() string {
	return r.raw
}

// SessionContext Session context data returned from a browser session.
type SessionContext struct {
	// Cookies to initialize in the session
	Cookies []SessionContextCookie `json:"cookies"`
	// Domain-specific indexedDB items to initialize in the session
	IndexedDB map[string][]SessionContextIndexedDB `json:"indexedDB"`
	// Domain-specific localStorage items to initialize in the session
	LocalStorage map[string]map[string]string `json:"localStorage"`
	// Domain-specific sessionStorage items to initialize in the session
	SessionStorage map[string]map[string]string `json:"sessionStorage"`
	JSON           sessionContextJSON           `json:"-"`
}

// sessionContextJSON contains the JSON metadata for the struct [SessionContext]
type sessionContextJSON struct {
	Cookies        apijson.Field
	IndexedDB      apijson.Field
	LocalStorage   apijson.Field
	SessionStorage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SessionContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionContextJSON) RawJSON() string {
	return r.raw
}

// SessionCostResponse Session cost breakdown in US dollars
type SessionCostResponse struct {
	// Currency used for the cost values
	Currency SessionCostResponseCurrency `json:"currency" api:"required"`
	// Session ID
	ID string `json:"id" api:"required"`
	// Exact total session cost in US dollars (e.g. 0.031905), with no rounding to whole cents. Rounded only to micro-dollar precision (6 decimal places) for precise billing passthrough.
	TotalCost float64 `json:"totalCost" api:"required"`
	// Cost unit. Values are expressed in US dollars.
	Unit SessionCostResponseUnit `json:"unit" api:"required"`
	// Billable usage inputs used to calculate the costs
	Usage SessionCostResponseUsage `json:"usage" api:"required"`
	JSON  sessionCostResponseJSON  `json:"-"`
}

// sessionCostResponseJSON contains the JSON metadata for the struct [SessionCostResponse]
type sessionCostResponseJSON struct {
	Currency    apijson.Field
	ID          apijson.Field
	TotalCost   apijson.Field
	Unit        apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionCostResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionCostResponseJSON) RawJSON() string {
	return r.raw
}

type SessionLiveDetailsResponse struct {
	Pages                      []SessionLiveDetailsResponsePage `json:"pages" api:"required"`
	SessionViewerFullscreenURL string                           `json:"sessionViewerFullscreenUrl" api:"required"`
	SessionViewerURL           string                           `json:"sessionViewerUrl" api:"required"`
	WsURL                      string                           `json:"wsUrl" api:"required"`
	JSON                       sessionLiveDetailsResponseJSON   `json:"-"`
}

// sessionLiveDetailsResponseJSON contains the JSON metadata for the struct [SessionLiveDetailsResponse]
type sessionLiveDetailsResponseJSON struct {
	Pages                      apijson.Field
	SessionViewerFullscreenURL apijson.Field
	SessionViewerURL           apijson.Field
	WsURL                      apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *SessionLiveDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionLiveDetailsResponseJSON) RawJSON() string {
	return r.raw
}

// Session Represents the data structure for a browser session, including its configuration and status.
type Session struct {
	// Timestamp when the session started
	CreatedAt time.Time `json:"createdAt" api:"required"`
	// Amount of credits consumed by the session
	CreditsUsed int64 `json:"creditsUsed" api:"required"`
	// Configuration for the debug URL and session viewer. Controls interaction capabilities and cursor visibility.
	DebugConfig SessionDebugConfig `json:"debugConfig"`
	// URL for debugging the session
	DebugURL string `json:"debugUrl" api:"required"`
	// Device configuration for the session
	DeviceConfig SessionDeviceConfig `json:"deviceConfig"`
	// Viewport and browser window dimensions for the session
	Dimensions SessionDimensions `json:"dimensions" api:"required"`
	// Duration of the session in milliseconds
	Duration int64 `json:"duration" api:"required"`
	// Number of events processed in the session
	EventCount int64 `json:"eventCount" api:"required"`
	// Launch the browser in fullscreen mode, covering the full screen with no Chrome UI.
	Fullscreen bool `json:"fullscreen"`
	// Indicates if the session is headless or headful
	Headless bool `json:"headless"`
	// Unique identifier for the session
	ID string `json:"id" api:"required"`
	// Inactivity timeout in milliseconds, if one was set when the session was created
	InactivityTimeout int64 `json:"inactivityTimeout"`
	// Indicates if Selenium is used in the session
	IsSelenium bool `json:"isSelenium"`
	// Bandwidth optimizations that were applied to the session.
	OptimizeBandwidth SessionOptimizeBandwidth `json:"optimizeBandwidth" api:"required"`
	// This flag will persist the profile for the session.
	PersistProfile bool `json:"persistProfile"`
	// The ID of the profile associated with the session
	ProfileID string `json:"profileId"`
	// The project associated with the session
	ProjectID string `json:"projectId"`
	// Amount of data transmitted through the proxy
	ProxyBytesUsed int64 `json:"proxyBytesUsed" api:"required"`
	// Source of the proxy used for the session
	ProxySource SessionResponseProxySource `json:"proxySource" api:"required"`
	// The region where the session was created.
	Region SessionResponseRegion `json:"region"`
	// Why the session reached a terminal state. Null while the session is live, or when the reason is unknown (e.g. sessions created before this was tracked). One of: user_requested (released via the API/SDK), timeout (hard `timeout` elapsed), inactivity_timeout (no activity for the configured window), creation_timeout (never started in time), startup_failed (could not be dispatched), browser_closed (the browser or agent closed itself — not a crash), browser_crashed (the browser crashed or its machine became unresponsive).
	ReleaseReason SessionResponseReleaseReason `json:"releaseReason"`
	// URL to view session details
	SessionViewerURL string `json:"sessionViewerUrl" api:"required"`
	// Indicates if captcha solving is enabled
	SolveCaptcha bool `json:"solveCaptcha"`
	// Status of the session
	Status SessionResponseStatus `json:"status" api:"required"`
	// Stealth configuration for the session
	StealthConfig SessionStealthConfig `json:"stealthConfig"`
	// Session timeout duration in milliseconds
	Timeout int64 `json:"timeout" api:"required"`
	// User agent string used in the session
	UserAgent string `json:"userAgent"`
	// URL for the session's WebSocket connection
	WebsocketURL string      `json:"websocketUrl" api:"required"`
	JSON         sessionJSON `json:"-"`
}

// sessionJSON contains the JSON metadata for the struct [Session]
type sessionJSON struct {
	CreatedAt         apijson.Field
	CreditsUsed       apijson.Field
	DebugConfig       apijson.Field
	DebugURL          apijson.Field
	DeviceConfig      apijson.Field
	Dimensions        apijson.Field
	Duration          apijson.Field
	EventCount        apijson.Field
	Fullscreen        apijson.Field
	Headless          apijson.Field
	ID                apijson.Field
	InactivityTimeout apijson.Field
	IsSelenium        apijson.Field
	OptimizeBandwidth apijson.Field
	PersistProfile    apijson.Field
	ProfileID         apijson.Field
	ProjectID         apijson.Field
	ProxyBytesUsed    apijson.Field
	ProxySource       apijson.Field
	Region            apijson.Field
	ReleaseReason     apijson.Field
	SessionViewerURL  apijson.Field
	SolveCaptcha      apijson.Field
	Status            apijson.Field
	StealthConfig     apijson.Field
	Timeout           apijson.Field
	UserAgent         apijson.Field
	WebsocketURL      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Session) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionJSON) RawJSON() string {
	return r.raw
}

// SessionList Response containing a list of browser sessions with pagination details.
type SessionList struct {
	// Cursor for the next page of results. Null if no more pages.
	NextCursor string `json:"nextCursor" api:"required"`
	// List of browser sessions
	Sessions []SessionListSession `json:"sessions" api:"required"`
	// Total number of sessions matching the query. Only included for filtered queries (e.g. status=live).
	TotalCount int64           `json:"totalCount"`
	JSON       sessionListJSON `json:"-"`
}

// sessionListJSON contains the JSON metadata for the struct [SessionList]
type sessionListJSON struct {
	NextCursor  apijson.Field
	Sessions    apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionListJSON) RawJSON() string {
	return r.raw
}

// CaptchaSolveParams If no fields are provided, all detected captchas will be solved
type CaptchaSolveParams struct {
	// The page ID where the captcha is located
	PageID param.Field[string] `json:"pageId"`
	// The task ID of the specific captcha to solve
	TaskID param.Field[string] `json:"taskId"`
	// The URL where the captcha is located
	URL param.Field[string] `json:"url"`
}

func (r CaptchaSolveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CaptchaSolveImageParams struct {
	// XPath to the captcha image element
	ImageXPath param.Field[string] `json:"imageXPath" api:"required"`
	// XPath to the captcha input element
	InputXPath param.Field[string] `json:"inputXPath" api:"required"`
	// URL where the captcha is located. Defaults to the current page URL
	URL param.Field[string] `json:"url"`
}

func (r CaptchaSolveImageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CaptchaSolveImageResponse struct {
	// Response message
	Message string `json:"message"`
	// Whether the action was successful
	Success bool                          `json:"success" api:"required"`
	JSON    captchaSolveImageResponseJSON `json:"-"`
}

// captchaSolveImageResponseJSON contains the JSON metadata for the struct [CaptchaSolveImageResponse]
type captchaSolveImageResponseJSON struct {
	Message     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptchaSolveImageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captchaSolveImageResponseJSON) RawJSON() string {
	return r.raw
}

type SessionFileUploadParams struct {
	// The file to upload (binary) or URL string to download from
	File FileUpload `json:"file"`
	// Path to the file in the storage system
	Path *string `json:"path,omitempty"`
}

type ComputerActionRequestMoveMouse struct {
	Action param.Field[ComputerActionRequestMoveMouseAction] `json:"action" api:"required"`
	// X and Y coordinates [x, y]
	Coordinates param.Field[[]float64] `json:"coordinates" api:"required"`
	// Keys to hold while moving
	HoldKeys param.Field[[]string] `json:"hold_keys"`
	// Whether to take a screenshot after the action
	Screenshot param.Field[bool] `json:"screenshot"`
}

func (r ComputerActionRequestMoveMouse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ComputerActionRequestClickMouse struct {
	Action param.Field[ComputerActionRequestClickMouseAction] `json:"action" api:"required"`
	// Mouse button to click. Defaults to 'left'
	Button param.Field[ComputerActionRequestClickMouseButton] `json:"button"`
	// Type of click (down, up, or click). Defaults to 'click'
	ClickType param.Field[ComputerActionRequestClickMouseClickType] `json:"click_type"`
	// X and Y coordinates [x, y]
	Coordinates param.Field[[]float64] `json:"coordinates"`
	// Keys to hold while clicking
	HoldKeys param.Field[[]string] `json:"hold_keys"`
	// Number of clicks. Defaults to 1
	NumClicks param.Field[float64] `json:"num_clicks"`
	// Whether to take a screenshot after the action
	Screenshot param.Field[bool] `json:"screenshot"`
}

func (r ComputerActionRequestClickMouse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ComputerActionRequestDragMouse struct {
	Action param.Field[ComputerActionRequestDragMouseAction] `json:"action" api:"required"`
	// Keys to hold while dragging
	HoldKeys param.Field[[]string] `json:"hold_keys"`
	// Array of [x, y] coordinate pairs
	Path param.Field[[][]float64] `json:"path" api:"required"`
	// Whether to take a screenshot after the action
	Screenshot param.Field[bool] `json:"screenshot"`
}

func (r ComputerActionRequestDragMouse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ComputerActionRequestScroll struct {
	Action param.Field[ComputerActionRequestScrollAction] `json:"action" api:"required"`
	// X and Y coordinates [x, y]
	Coordinates param.Field[[]float64] `json:"coordinates"`
	// Horizontal scroll amount. Defaults to 0
	DeltaX param.Field[float64] `json:"delta_x"`
	// Vertical scroll amount. Defaults to 0
	DeltaY param.Field[float64] `json:"delta_y"`
	// Keys to hold while scrolling
	HoldKeys param.Field[[]string] `json:"hold_keys"`
	// Whether to take a screenshot after the action
	Screenshot param.Field[bool] `json:"screenshot"`
}

func (r ComputerActionRequestScroll) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ComputerActionRequestPressKey struct {
	Action param.Field[ComputerActionRequestPressKeyAction] `json:"action" api:"required"`
	// Duration to hold keys in seconds
	Duration param.Field[float64] `json:"duration"`
	// Keys to press
	Keys param.Field[[]string] `json:"keys" api:"required"`
	// Whether to take a screenshot after the action
	Screenshot param.Field[bool] `json:"screenshot"`
}

func (r ComputerActionRequestPressKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ComputerActionRequestTypeText struct {
	Action param.Field[ComputerActionRequestTypeTextAction] `json:"action" api:"required"`
	// Keys to hold while typing
	HoldKeys param.Field[[]string] `json:"hold_keys"`
	// Whether to take a screenshot after the action
	Screenshot param.Field[bool] `json:"screenshot"`
	// Text to type
	Text param.Field[string] `json:"text" api:"required"`
}

func (r ComputerActionRequestTypeText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ComputerActionRequestWait struct {
	Action param.Field[ComputerActionRequestWaitAction] `json:"action" api:"required"`
	// Duration to wait in seconds
	Duration param.Field[float64] `json:"duration" api:"required"`
	// Whether to take a screenshot after the action
	Screenshot param.Field[bool] `json:"screenshot"`
}

func (r ComputerActionRequestWait) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ComputerActionRequestTakeScreenshot struct {
	Action param.Field[ComputerActionRequestTakeScreenshotAction] `json:"action" api:"required"`
}

func (r ComputerActionRequestTakeScreenshot) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ComputerActionRequestGetCursorPosition struct {
	Action param.Field[ComputerActionRequestGetCursorPositionAction] `json:"action" api:"required"`
}

func (r ComputerActionRequestGetCursorPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CaptchaStatusResponseItem struct {
	// Timestamp when the state was created
	Created float64 `json:"created"`
	// Whether a captcha is currently being solved
	IsSolvingCaptcha bool `json:"isSolvingCaptcha" api:"required"`
	// Timestamp when the state was last updated
	LastUpdated float64 `json:"lastUpdated"`
	// The page ID where the captcha is located
	PageID string `json:"pageId" api:"required"`
	// Array of captcha tasks
	Tasks []interface{} `json:"tasks" api:"required"`
	// The URL where the captcha is located
	URL  string                        `json:"url" api:"required"`
	JSON captchaStatusResponseItemJSON `json:"-"`
}

// captchaStatusResponseItemJSON contains the JSON metadata for the struct [CaptchaStatusResponseItem]
type captchaStatusResponseItemJSON struct {
	Created          apijson.Field
	IsSolvingCaptcha apijson.Field
	LastUpdated      apijson.Field
	PageID           apijson.Field
	Tasks            apijson.Field
	URL              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CaptchaStatusResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captchaStatusResponseItemJSON) RawJSON() string {
	return r.raw
}

// SessionCreateParamsCredentials Configuration for session credentials
type SessionCreateParamsCredentials struct {
	AutoSubmit  param.Field[bool] `json:"autoSubmit"`
	BlurFields  param.Field[bool] `json:"blurFields"`
	ExactOrigin param.Field[bool] `json:"exactOrigin"`
}

func (r SessionCreateParamsCredentials) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SessionCreateParamsDebugConfig Configuration for the debug URL and session viewer. Controls interaction capabilities, cursor visibility, and other debug-related settings.
type SessionCreateParamsDebugConfig struct {
	// Allow interaction with the browser session via the debug URL viewer. When false, the session viewer will be view-only. Default is true.
	Interactive param.Field[bool] `json:"interactive"`
	// Show the OS-level mouse cursor in the WebRTC stream (headful mode only). When false, the system cursor will not be rendered in the stream. Default is true.
	SystemCursor param.Field[bool] `json:"systemCursor"`
}

func (r SessionCreateParamsDebugConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SessionCreateParamsDeviceConfig Device configuration for the session. Specify 'mobile' for mobile device fingerprints and configurations.
type SessionCreateParamsDeviceConfig struct {
	Device param.Field[CreateSessionRequestDeviceConfigDevice] `json:"device"`
}

func (r SessionCreateParamsDeviceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SessionCreateParamsDimensions Viewport and browser window dimensions for the session. Mobile sessions require dimensions of at least 508x1074; smaller mobile dimensions are rejected with a 400 response.
type SessionCreateParamsDimensions struct {
	// Height of the session
	Height param.Field[int64] `json:"height" api:"required"`
	// Width of the session
	Width param.Field[int64] `json:"width" api:"required"`
}

func (r SessionCreateParamsDimensions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SessionCreateParamsOptimizeBandwidthUnionMember1 struct {
	BlockHosts       param.Field[[]string] `json:"blockHosts"`
	BlockImages      param.Field[bool]     `json:"blockImages"`
	BlockMedia       param.Field[bool]     `json:"blockMedia"`
	BlockStylesheets param.Field[bool]     `json:"blockStylesheets"`
	BlockURLPatterns param.Field[[]string] `json:"blockUrlPatterns"`
}

func (r SessionCreateParamsOptimizeBandwidthUnionMember1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SessionCreateParamsSessionContextCookiePartitionKey The partition key of the cookie
type SessionCreateParamsSessionContextCookiePartitionKey struct {
	// Indicates if the cookie has any ancestors that are cross-site to the topLevelSite.
	HasCrossSiteAncestor param.Field[bool] `json:"hasCrossSiteAncestor" api:"required"`
	// The site of the top-level URL the browser was visiting at the start of the request to the endpoint that set the cookie.
	TopLevelSite param.Field[string] `json:"topLevelSite" api:"required"`
}

func (r SessionCreateParamsSessionContextCookiePartitionKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SessionCreateParamsSessionContextCookie Cookies to initialize in the session
type SessionCreateParamsSessionContextCookie struct {
	// The domain of the cookie
	Domain param.Field[string] `json:"domain"`
	// The expiration date of the cookie
	Expires param.Field[float64] `json:"expires"`
	// Whether the cookie is HTTP only
	HTTPOnly param.Field[bool] `json:"httpOnly"`
	// The name of the cookie
	Name param.Field[string] `json:"name" api:"required"`
	// The partition key of the cookie
	PartitionKey param.Field[SessionCreateParamsSessionContextCookiePartitionKey] `json:"partitionKey"`
	// The path of the cookie
	Path param.Field[string] `json:"path"`
	// The priority of the cookie
	Priority param.Field[CreateSessionRequestSessionContextCookiesItemPriority] `json:"priority"`
	// Whether the cookie is a same party cookie
	SameParty param.Field[bool] `json:"sameParty"`
	// The same site attribute of the cookie
	SameSite param.Field[CreateSessionRequestSessionContextCookiesItemSameSite] `json:"sameSite"`
	// Whether the cookie is secure
	Secure param.Field[bool] `json:"secure"`
	// Whether the cookie is a session cookie
	Session param.Field[bool] `json:"session"`
	// The size of the cookie
	Size param.Field[int64] `json:"size"`
	// The source port of the cookie
	SourcePort param.Field[float64] `json:"sourcePort"`
	// The source scheme of the cookie
	SourceScheme param.Field[CreateSessionRequestSessionContextCookiesItemSourceScheme] `json:"sourceScheme"`
	// The URL of the cookie
	URL param.Field[string] `json:"url"`
	// The value of the cookie
	Value param.Field[string] `json:"value" api:"required"`
}

func (r SessionCreateParamsSessionContextCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SessionCreateParamsSessionContextIndexedDBDataRecordBlobFile struct {
	BlobNumber   param.Field[float64]   `json:"blobNumber" api:"required"`
	Filename     param.Field[string]    `json:"filename"`
	LastModified param.Field[time.Time] `json:"lastModified"`
	MimeType     param.Field[string]    `json:"mimeType" api:"required"`
	Path         param.Field[string]    `json:"path"`
	Size         param.Field[int64]     `json:"size" api:"required"`
}

func (r SessionCreateParamsSessionContextIndexedDBDataRecordBlobFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SessionCreateParamsSessionContextIndexedDBDataRecord struct {
	BlobFiles param.Field[[]SessionCreateParamsSessionContextIndexedDBDataRecordBlobFile] `json:"blobFiles"`
	Key       param.Field[interface{}]                                                    `json:"key" api:"required"`
	Value     param.Field[interface{}]                                                    `json:"value" api:"required"`
}

func (r SessionCreateParamsSessionContextIndexedDBDataRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SessionCreateParamsSessionContextIndexedDBData struct {
	ID      param.Field[float64]                                                `json:"id" api:"required"`
	Name    param.Field[string]                                                 `json:"name" api:"required"`
	Records param.Field[[]SessionCreateParamsSessionContextIndexedDBDataRecord] `json:"records" api:"required"`
}

func (r SessionCreateParamsSessionContextIndexedDBData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SessionCreateParamsSessionContextIndexedDB Domain-specific indexedDB items to initialize in the session
type SessionCreateParamsSessionContextIndexedDB struct {
	Data param.Field[[]SessionCreateParamsSessionContextIndexedDBData] `json:"data" api:"required"`
	ID   param.Field[float64]                                          `json:"id" api:"required"`
	Name param.Field[string]                                           `json:"name" api:"required"`
}

func (r SessionCreateParamsSessionContextIndexedDB) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SessionCreateParamsSessionContext Session context data to be used in the created session. Sessions will start with an empty context by default.
type SessionCreateParamsSessionContext struct {
	// Cookies to initialize in the session
	Cookies param.Field[[]SessionCreateParamsSessionContextCookie] `json:"cookies"`
	// Domain-specific indexedDB items to initialize in the session
	IndexedDB param.Field[map[string][]SessionCreateParamsSessionContextIndexedDB] `json:"indexedDB"`
	// Domain-specific localStorage items to initialize in the session
	LocalStorage param.Field[map[string]map[string]string] `json:"localStorage"`
	// Domain-specific sessionStorage items to initialize in the session
	SessionStorage param.Field[map[string]map[string]string] `json:"sessionStorage"`
}

func (r SessionCreateParamsSessionContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SessionCreateParamsStealthConfig Stealth configuration for the session
type SessionCreateParamsStealthConfig struct {
	// When true, captchas will be automatically solved when detected. When false, use the solve endpoints to manually initiate solving.
	AutoCaptchaSolving param.Field[bool] `json:"autoCaptchaSolving"`
	// This flag will make the browser act more human-like by moving the mouse in a more natural way.
	HumanizeInteractions param.Field[bool] `json:"humanizeInteractions"`
	// This flag will skip the fingerprint generation for the session.
	SkipFingerprintInjection param.Field[bool] `json:"skipFingerprintInjection"`
}

func (r SessionCreateParamsStealthConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SessionCreateParamsUseProxyGeolocationGeolocation Geographic location for the proxy
type SessionCreateParamsUseProxyGeolocationGeolocation struct {
	// City name (e.g., 'NEW_YORK', 'LOS_ANGELES')
	City param.Field[CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCity] `json:"city"`
	// Country code (e.g., 'US', 'GB', 'DE') - ISO 3166-1 alpha-2
	Country param.Field[CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationCountry] `json:"country" api:"required"`
	// State code (e.g., 'NY', 'CA') - US states only
	State param.Field[CreateSessionRequestUseProxyVariant0Variant1Variant0GeolocationState] `json:"state"`
}

func (r SessionCreateParamsUseProxyGeolocationGeolocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SessionCreateParamsUseProxyGeolocation struct {
	// Geographic location for the proxy
	Geolocation param.Field[SessionCreateParamsUseProxyGeolocationGeolocation] `json:"geolocation" api:"required"`
}

func (r SessionCreateParamsUseProxyGeolocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SessionCreateParamsUseProxyServer struct {
	// Proxy server URL
	Server param.Field[string] `json:"server" api:"required"`
}

func (r SessionCreateParamsUseProxyServer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SessionContextCookiePartitionKey The partition key of the cookie
type SessionContextCookiePartitionKey struct {
	// Indicates if the cookie has any ancestors that are cross-site to the topLevelSite.
	HasCrossSiteAncestor bool `json:"hasCrossSiteAncestor" api:"required"`
	// The site of the top-level URL the browser was visiting at the start of the request to the endpoint that set the cookie.
	TopLevelSite string                               `json:"topLevelSite" api:"required"`
	JSON         sessionContextCookiePartitionKeyJSON `json:"-"`
}

// sessionContextCookiePartitionKeyJSON contains the JSON metadata for the struct [SessionContextCookiePartitionKey]
type sessionContextCookiePartitionKeyJSON struct {
	HasCrossSiteAncestor apijson.Field
	TopLevelSite         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SessionContextCookiePartitionKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionContextCookiePartitionKeyJSON) RawJSON() string {
	return r.raw
}

// SessionContextCookie Cookies to initialize in the session
type SessionContextCookie struct {
	// The domain of the cookie
	Domain string `json:"domain"`
	// The expiration date of the cookie
	Expires float64 `json:"expires"`
	// Whether the cookie is HTTP only
	HTTPOnly bool `json:"httpOnly"`
	// The name of the cookie
	Name string `json:"name" api:"required"`
	// The partition key of the cookie
	PartitionKey SessionContextCookiePartitionKey `json:"partitionKey"`
	// The path of the cookie
	Path string `json:"path"`
	// The priority of the cookie
	Priority CreateSessionRequestSessionContextCookiesItemPriority `json:"priority"`
	// Whether the cookie is a same party cookie
	SameParty bool `json:"sameParty"`
	// The same site attribute of the cookie
	SameSite CreateSessionRequestSessionContextCookiesItemSameSite `json:"sameSite"`
	// Whether the cookie is secure
	Secure bool `json:"secure"`
	// Whether the cookie is a session cookie
	Session bool `json:"session"`
	// The size of the cookie
	Size int64 `json:"size"`
	// The source port of the cookie
	SourcePort float64 `json:"sourcePort"`
	// The source scheme of the cookie
	SourceScheme CreateSessionRequestSessionContextCookiesItemSourceScheme `json:"sourceScheme"`
	// The URL of the cookie
	URL string `json:"url"`
	// The value of the cookie
	Value string                   `json:"value" api:"required"`
	JSON  sessionContextCookieJSON `json:"-"`
}

// sessionContextCookieJSON contains the JSON metadata for the struct [SessionContextCookie]
type sessionContextCookieJSON struct {
	Domain       apijson.Field
	Expires      apijson.Field
	HTTPOnly     apijson.Field
	Name         apijson.Field
	PartitionKey apijson.Field
	Path         apijson.Field
	Priority     apijson.Field
	SameParty    apijson.Field
	SameSite     apijson.Field
	Secure       apijson.Field
	Session      apijson.Field
	Size         apijson.Field
	SourcePort   apijson.Field
	SourceScheme apijson.Field
	URL          apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SessionContextCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionContextCookieJSON) RawJSON() string {
	return r.raw
}

type SessionContextIndexedDBDataRecordBlobFile struct {
	BlobNumber   float64                                       `json:"blobNumber" api:"required"`
	Filename     string                                        `json:"filename"`
	LastModified time.Time                                     `json:"lastModified"`
	MimeType     string                                        `json:"mimeType" api:"required"`
	Path         string                                        `json:"path"`
	Size         int64                                         `json:"size" api:"required"`
	JSON         sessionContextIndexedDBDataRecordBlobFileJSON `json:"-"`
}

// sessionContextIndexedDBDataRecordBlobFileJSON contains the JSON metadata for the struct [SessionContextIndexedDBDataRecordBlobFile]
type sessionContextIndexedDBDataRecordBlobFileJSON struct {
	BlobNumber   apijson.Field
	Filename     apijson.Field
	LastModified apijson.Field
	MimeType     apijson.Field
	Path         apijson.Field
	Size         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SessionContextIndexedDBDataRecordBlobFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionContextIndexedDBDataRecordBlobFileJSON) RawJSON() string {
	return r.raw
}

type SessionContextIndexedDBDataRecord struct {
	BlobFiles []SessionContextIndexedDBDataRecordBlobFile `json:"blobFiles"`
	Key       interface{}                                 `json:"key" api:"required"`
	Value     interface{}                                 `json:"value" api:"required"`
	JSON      sessionContextIndexedDBDataRecordJSON       `json:"-"`
}

// sessionContextIndexedDBDataRecordJSON contains the JSON metadata for the struct [SessionContextIndexedDBDataRecord]
type sessionContextIndexedDBDataRecordJSON struct {
	BlobFiles   apijson.Field
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionContextIndexedDBDataRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionContextIndexedDBDataRecordJSON) RawJSON() string {
	return r.raw
}

type SessionContextIndexedDBData struct {
	ID      float64                             `json:"id" api:"required"`
	Name    string                              `json:"name" api:"required"`
	Records []SessionContextIndexedDBDataRecord `json:"records" api:"required"`
	JSON    sessionContextIndexedDBDataJSON     `json:"-"`
}

// sessionContextIndexedDBDataJSON contains the JSON metadata for the struct [SessionContextIndexedDBData]
type sessionContextIndexedDBDataJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Records     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionContextIndexedDBData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionContextIndexedDBDataJSON) RawJSON() string {
	return r.raw
}

// SessionContextIndexedDB Domain-specific indexedDB items to initialize in the session
type SessionContextIndexedDB struct {
	Data []SessionContextIndexedDBData `json:"data" api:"required"`
	ID   float64                       `json:"id" api:"required"`
	Name string                        `json:"name" api:"required"`
	JSON sessionContextIndexedDBJSON   `json:"-"`
}

// sessionContextIndexedDBJSON contains the JSON metadata for the struct [SessionContextIndexedDB]
type sessionContextIndexedDBJSON struct {
	Data        apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionContextIndexedDB) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionContextIndexedDBJSON) RawJSON() string {
	return r.raw
}

type SessionCostResponseUsageBrowser struct {
	// Billing unit for browser usage
	Unit SessionCostResponseUsageBrowserUnit `json:"unit" api:"required"`
	// Billable browser usage value
	Value float64                             `json:"value" api:"required"`
	JSON  sessionCostResponseUsageBrowserJSON `json:"-"`
}

// sessionCostResponseUsageBrowserJSON contains the JSON metadata for the struct [SessionCostResponseUsageBrowser]
type sessionCostResponseUsageBrowserJSON struct {
	Unit        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionCostResponseUsageBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionCostResponseUsageBrowserJSON) RawJSON() string {
	return r.raw
}

type SessionCostResponseUsageCaptcha struct {
	// Billable captcha solves
	Solves int64                               `json:"solves" api:"required"`
	JSON   sessionCostResponseUsageCaptchaJSON `json:"-"`
}

// sessionCostResponseUsageCaptchaJSON contains the JSON metadata for the struct [SessionCostResponseUsageCaptcha]
type sessionCostResponseUsageCaptchaJSON struct {
	Solves      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionCostResponseUsageCaptcha) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionCostResponseUsageCaptchaJSON) RawJSON() string {
	return r.raw
}

type SessionCostResponseUsageProxy struct {
	// Billable Steel proxy bytes for the session
	Bytes int64                             `json:"bytes" api:"required"`
	JSON  sessionCostResponseUsageProxyJSON `json:"-"`
}

// sessionCostResponseUsageProxyJSON contains the JSON metadata for the struct [SessionCostResponseUsageProxy]
type sessionCostResponseUsageProxyJSON struct {
	Bytes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionCostResponseUsageProxy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionCostResponseUsageProxyJSON) RawJSON() string {
	return r.raw
}

// SessionCostResponseUsage Billable usage inputs used to calculate the costs
type SessionCostResponseUsage struct {
	Browser SessionCostResponseUsageBrowser `json:"browser" api:"required"`
	Captcha SessionCostResponseUsageCaptcha `json:"captcha" api:"required"`
	Proxy   SessionCostResponseUsageProxy   `json:"proxy" api:"required"`
	JSON    sessionCostResponseUsageJSON    `json:"-"`
}

// sessionCostResponseUsageJSON contains the JSON metadata for the struct [SessionCostResponseUsage]
type sessionCostResponseUsageJSON struct {
	Browser     apijson.Field
	Captcha     apijson.Field
	Proxy       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionCostResponseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionCostResponseUsageJSON) RawJSON() string {
	return r.raw
}

type SessionLiveDetailsResponsePage struct {
	Favicon                    string                             `json:"favicon" api:"required"`
	ID                         string                             `json:"id" api:"required"`
	SessionViewerFullscreenURL string                             `json:"sessionViewerFullscreenUrl" api:"required"`
	SessionViewerURL           string                             `json:"sessionViewerUrl" api:"required"`
	Title                      string                             `json:"title" api:"required"`
	URL                        string                             `json:"url" api:"required"`
	JSON                       sessionLiveDetailsResponsePageJSON `json:"-"`
}

// sessionLiveDetailsResponsePageJSON contains the JSON metadata for the struct [SessionLiveDetailsResponsePage]
type sessionLiveDetailsResponsePageJSON struct {
	Favicon                    apijson.Field
	ID                         apijson.Field
	SessionViewerFullscreenURL apijson.Field
	SessionViewerURL           apijson.Field
	Title                      apijson.Field
	URL                        apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *SessionLiveDetailsResponsePage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionLiveDetailsResponsePageJSON) RawJSON() string {
	return r.raw
}

// SessionDebugConfig Configuration for the debug URL and session viewer. Controls interaction capabilities and cursor visibility.
type SessionDebugConfig struct {
	// Whether interaction is allowed via the debug URL viewer. When false, the session viewer is view-only.
	Interactive bool `json:"interactive"`
	// Whether the OS-level mouse cursor is shown in the WebRTC stream (headful mode only).
	SystemCursor bool                   `json:"systemCursor"`
	JSON         sessionDebugConfigJSON `json:"-"`
}

// sessionDebugConfigJSON contains the JSON metadata for the struct [SessionDebugConfig]
type sessionDebugConfigJSON struct {
	Interactive  apijson.Field
	SystemCursor apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SessionDebugConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionDebugConfigJSON) RawJSON() string {
	return r.raw
}

// SessionDeviceConfig Device configuration for the session
type SessionDeviceConfig struct {
	Device CreateSessionRequestDeviceConfigDevice `json:"device"`
	JSON   sessionDeviceConfigJSON                `json:"-"`
}

// sessionDeviceConfigJSON contains the JSON metadata for the struct [SessionDeviceConfig]
type sessionDeviceConfigJSON struct {
	Device      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionDeviceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionDeviceConfigJSON) RawJSON() string {
	return r.raw
}

// SessionDimensions Viewport and browser window dimensions for the session
type SessionDimensions struct {
	// Height of the browser window
	Height int64 `json:"height" api:"required"`
	// Width of the browser window
	Width int64                 `json:"width" api:"required"`
	JSON  sessionDimensionsJSON `json:"-"`
}

// sessionDimensionsJSON contains the JSON metadata for the struct [SessionDimensions]
type sessionDimensionsJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionDimensions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionDimensionsJSON) RawJSON() string {
	return r.raw
}

// SessionOptimizeBandwidth Bandwidth optimizations that were applied to the session.
type SessionOptimizeBandwidth struct {
	BlockHosts       []string                     `json:"blockHosts"`
	BlockImages      bool                         `json:"blockImages"`
	BlockMedia       bool                         `json:"blockMedia"`
	BlockStylesheets bool                         `json:"blockStylesheets"`
	BlockURLPatterns []string                     `json:"blockUrlPatterns"`
	JSON             sessionOptimizeBandwidthJSON `json:"-"`
}

// sessionOptimizeBandwidthJSON contains the JSON metadata for the struct [SessionOptimizeBandwidth]
type sessionOptimizeBandwidthJSON struct {
	BlockHosts       apijson.Field
	BlockImages      apijson.Field
	BlockMedia       apijson.Field
	BlockStylesheets apijson.Field
	BlockURLPatterns apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SessionOptimizeBandwidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionOptimizeBandwidthJSON) RawJSON() string {
	return r.raw
}

// SessionStealthConfig Stealth configuration for the session
type SessionStealthConfig struct {
	// When true, captchas will be automatically solved when detected. When false, use the solve endpoints to manually initiate solving.
	AutoCaptchaSolving bool `json:"autoCaptchaSolving"`
	// This flag will make the browser act more human-like by moving the mouse in a more natural way
	HumanizeInteractions bool `json:"humanizeInteractions"`
	// This flag will skip the fingerprint generation for the session.
	SkipFingerprintInjection bool                     `json:"skipFingerprintInjection"`
	JSON                     sessionStealthConfigJSON `json:"-"`
}

// sessionStealthConfigJSON contains the JSON metadata for the struct [SessionStealthConfig]
type sessionStealthConfigJSON struct {
	AutoCaptchaSolving       apijson.Field
	HumanizeInteractions     apijson.Field
	SkipFingerprintInjection apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *SessionStealthConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionStealthConfigJSON) RawJSON() string {
	return r.raw
}

// SessionListSessionDebugConfig Configuration for the debug URL and session viewer. Controls interaction capabilities and cursor visibility.
type SessionListSessionDebugConfig struct {
	// Whether interaction is allowed via the debug URL viewer. When false, the session viewer is view-only.
	Interactive bool `json:"interactive"`
	// Whether the OS-level mouse cursor is shown in the WebRTC stream (headful mode only).
	SystemCursor bool                              `json:"systemCursor"`
	JSON         sessionListSessionDebugConfigJSON `json:"-"`
}

// sessionListSessionDebugConfigJSON contains the JSON metadata for the struct [SessionListSessionDebugConfig]
type sessionListSessionDebugConfigJSON struct {
	Interactive  apijson.Field
	SystemCursor apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SessionListSessionDebugConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionListSessionDebugConfigJSON) RawJSON() string {
	return r.raw
}

// SessionListSessionDeviceConfig Device configuration for the session
type SessionListSessionDeviceConfig struct {
	Device CreateSessionRequestDeviceConfigDevice `json:"device"`
	JSON   sessionListSessionDeviceConfigJSON     `json:"-"`
}

// sessionListSessionDeviceConfigJSON contains the JSON metadata for the struct [SessionListSessionDeviceConfig]
type sessionListSessionDeviceConfigJSON struct {
	Device      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionListSessionDeviceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionListSessionDeviceConfigJSON) RawJSON() string {
	return r.raw
}

// SessionListSessionDimensions Viewport and browser window dimensions for the session
type SessionListSessionDimensions struct {
	// Height of the browser window
	Height int64 `json:"height" api:"required"`
	// Width of the browser window
	Width int64                            `json:"width" api:"required"`
	JSON  sessionListSessionDimensionsJSON `json:"-"`
}

// sessionListSessionDimensionsJSON contains the JSON metadata for the struct [SessionListSessionDimensions]
type sessionListSessionDimensionsJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionListSessionDimensions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionListSessionDimensionsJSON) RawJSON() string {
	return r.raw
}

// SessionListSessionOptimizeBandwidth Bandwidth optimizations that were applied to the session.
type SessionListSessionOptimizeBandwidth struct {
	BlockHosts       []string                                `json:"blockHosts"`
	BlockImages      bool                                    `json:"blockImages"`
	BlockMedia       bool                                    `json:"blockMedia"`
	BlockStylesheets bool                                    `json:"blockStylesheets"`
	BlockURLPatterns []string                                `json:"blockUrlPatterns"`
	JSON             sessionListSessionOptimizeBandwidthJSON `json:"-"`
}

// sessionListSessionOptimizeBandwidthJSON contains the JSON metadata for the struct [SessionListSessionOptimizeBandwidth]
type sessionListSessionOptimizeBandwidthJSON struct {
	BlockHosts       apijson.Field
	BlockImages      apijson.Field
	BlockMedia       apijson.Field
	BlockStylesheets apijson.Field
	BlockURLPatterns apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SessionListSessionOptimizeBandwidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionListSessionOptimizeBandwidthJSON) RawJSON() string {
	return r.raw
}

// SessionListSessionStealthConfig Stealth configuration for the session
type SessionListSessionStealthConfig struct {
	// When true, captchas will be automatically solved when detected. When false, use the solve endpoints to manually initiate solving.
	AutoCaptchaSolving bool `json:"autoCaptchaSolving"`
	// This flag will make the browser act more human-like by moving the mouse in a more natural way
	HumanizeInteractions bool `json:"humanizeInteractions"`
	// This flag will skip the fingerprint generation for the session.
	SkipFingerprintInjection bool                                `json:"skipFingerprintInjection"`
	JSON                     sessionListSessionStealthConfigJSON `json:"-"`
}

// sessionListSessionStealthConfigJSON contains the JSON metadata for the struct [SessionListSessionStealthConfig]
type sessionListSessionStealthConfigJSON struct {
	AutoCaptchaSolving       apijson.Field
	HumanizeInteractions     apijson.Field
	SkipFingerprintInjection apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *SessionListSessionStealthConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionListSessionStealthConfigJSON) RawJSON() string {
	return r.raw
}

// SessionListSession List of browser sessions
type SessionListSession struct {
	// Timestamp when the session started
	CreatedAt time.Time `json:"createdAt" api:"required"`
	// Amount of credits consumed by the session
	CreditsUsed int64 `json:"creditsUsed" api:"required"`
	// Configuration for the debug URL and session viewer. Controls interaction capabilities and cursor visibility.
	DebugConfig SessionListSessionDebugConfig `json:"debugConfig"`
	// URL for debugging the session
	DebugURL string `json:"debugUrl" api:"required"`
	// Device configuration for the session
	DeviceConfig SessionListSessionDeviceConfig `json:"deviceConfig"`
	// Viewport and browser window dimensions for the session
	Dimensions SessionListSessionDimensions `json:"dimensions" api:"required"`
	// Duration of the session in milliseconds
	Duration int64 `json:"duration" api:"required"`
	// Number of events processed in the session
	EventCount int64 `json:"eventCount" api:"required"`
	// Launch the browser in fullscreen mode, covering the full screen with no Chrome UI.
	Fullscreen bool `json:"fullscreen"`
	// Indicates if the session is headless or headful
	Headless bool `json:"headless"`
	// Unique identifier for the session
	ID string `json:"id" api:"required"`
	// Inactivity timeout in milliseconds, if one was set when the session was created
	InactivityTimeout int64 `json:"inactivityTimeout"`
	// Indicates if Selenium is used in the session
	IsSelenium bool `json:"isSelenium"`
	// Bandwidth optimizations that were applied to the session.
	OptimizeBandwidth SessionListSessionOptimizeBandwidth `json:"optimizeBandwidth" api:"required"`
	// This flag will persist the profile for the session.
	PersistProfile bool `json:"persistProfile"`
	// The ID of the profile associated with the session
	ProfileID string `json:"profileId"`
	// The project associated with the session
	ProjectID string `json:"projectId"`
	// Amount of data transmitted through the proxy
	ProxyBytesUsed int64 `json:"proxyBytesUsed" api:"required"`
	// Source of the proxy used for the session
	ProxySource SessionResponseProxySource `json:"proxySource" api:"required"`
	// The region where the session was created.
	Region SessionResponseRegion `json:"region"`
	// Why the session reached a terminal state. Null while the session is live, or when the reason is unknown (e.g. sessions created before this was tracked). One of: user_requested (released via the API/SDK), timeout (hard `timeout` elapsed), inactivity_timeout (no activity for the configured window), creation_timeout (never started in time), startup_failed (could not be dispatched), browser_closed (the browser or agent closed itself — not a crash), browser_crashed (the browser crashed or its machine became unresponsive).
	ReleaseReason SessionResponseReleaseReason `json:"releaseReason"`
	// URL to view session details
	SessionViewerURL string `json:"sessionViewerUrl" api:"required"`
	// Indicates if captcha solving is enabled
	SolveCaptcha bool `json:"solveCaptcha"`
	// Status of the session
	Status SessionResponseStatus `json:"status" api:"required"`
	// Stealth configuration for the session
	StealthConfig SessionListSessionStealthConfig `json:"stealthConfig"`
	// Session timeout duration in milliseconds
	Timeout int64 `json:"timeout" api:"required"`
	// User agent string used in the session
	UserAgent string `json:"userAgent"`
	// URL for the session's WebSocket connection
	WebsocketURL string                 `json:"websocketUrl" api:"required"`
	JSON         sessionListSessionJSON `json:"-"`
}

// sessionListSessionJSON contains the JSON metadata for the struct [SessionListSession]
type sessionListSessionJSON struct {
	CreatedAt         apijson.Field
	CreditsUsed       apijson.Field
	DebugConfig       apijson.Field
	DebugURL          apijson.Field
	DeviceConfig      apijson.Field
	Dimensions        apijson.Field
	Duration          apijson.Field
	EventCount        apijson.Field
	Fullscreen        apijson.Field
	Headless          apijson.Field
	ID                apijson.Field
	InactivityTimeout apijson.Field
	IsSelenium        apijson.Field
	OptimizeBandwidth apijson.Field
	PersistProfile    apijson.Field
	ProfileID         apijson.Field
	ProjectID         apijson.Field
	ProxyBytesUsed    apijson.Field
	ProxySource       apijson.Field
	Region            apijson.Field
	ReleaseReason     apijson.Field
	SessionViewerURL  apijson.Field
	SolveCaptcha      apijson.Field
	Status            apijson.Field
	StealthConfig     apijson.Field
	Timeout           apijson.Field
	UserAgent         apijson.Field
	WebsocketURL      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SessionListSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionListSessionJSON) RawJSON() string {
	return r.raw
}

type SessionListParams struct {
	CursorID  *string                `json:"cursorId,omitempty"`
	Limit     *int64                 `json:"limit,omitempty"`
	Status    *SessionResponseStatus `json:"status,omitempty"`
	ProjectID *string                `json:"projectId,omitempty"`
}

type SessionReleaseAllQueryParams struct {
	ProjectID *string `json:"projectId,omitempty"`
}

type SessionAgentTracesParams struct {
	Namespace  *string    `json:"namespace,omitempty"`
	StartTime  *time.Time `json:"startTime,omitempty"`
	EndTime    *time.Time `json:"endTime,omitempty"`
	EventTypes *[]string  `json:"eventTypes,omitempty"`
}

type SessionEventsParams struct {
	Compressed *bool   `json:"compressed,omitempty"`
	Limit      *int64  `json:"limit,omitempty"`
	Pointer    *string `json:"pointer,omitempty"`
}
