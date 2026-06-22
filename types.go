package steel

import (
	"encoding/json"
	"time"
)

type ProfileStatus string

const (
	ProfileStatusUploading ProfileStatus = "UPLOADING"
	ProfileStatusReady     ProfileStatus = "READY"
	ProfileStatusFailed    ProfileStatus = "FAILED"
)

type ScrapeRequestFormatItem string

const (
	ScrapeRequestFormatItemHTML        ScrapeRequestFormatItem = "html"
	ScrapeRequestFormatItemReadability ScrapeRequestFormatItem = "readability"
	ScrapeRequestFormatItemCleanedHTML ScrapeRequestFormatItem = "cleaned_html"
	ScrapeRequestFormatItemMarkdown    ScrapeRequestFormatItem = "markdown"
)

type SessionCostResponseCurrency string

const (
	SessionCostResponseCurrencyUsd SessionCostResponseCurrency = "usd"
)

type SessionCostResponseUnit string

const (
	SessionCostResponseUnitDollar SessionCostResponseUnit = "dollar"
)

type SessionResponseProxySource string

const (
	SessionResponseProxySourceSteel    SessionResponseProxySource = "steel"
	SessionResponseProxySourceExternal SessionResponseProxySource = "external"
)

type SessionResponseRegion string

const (
	SessionResponseRegionLax         SessionResponseRegion = "lax"
	SessionResponseRegionOrd         SessionResponseRegion = "ord"
	SessionResponseRegionIad         SessionResponseRegion = "iad"
	SessionResponseRegionScl         SessionResponseRegion = "scl"
	SessionResponseRegionFra         SessionResponseRegion = "fra"
	SessionResponseRegionNrt         SessionResponseRegion = "nrt"
	SessionResponseRegionUsEast      SessionResponseRegion = "us-east"
	SessionResponseRegionUsWest      SessionResponseRegion = "us-west"
	SessionResponseRegionUsCentral   SessionResponseRegion = "us-central"
	SessionResponseRegionEuWest      SessionResponseRegion = "eu-west"
	SessionResponseRegionEuCentral   SessionResponseRegion = "eu-central"
	SessionResponseRegionApNortheast SessionResponseRegion = "ap-northeast"
	SessionResponseRegionApSoutheast SessionResponseRegion = "ap-southeast"
	SessionResponseRegionSaEast      SessionResponseRegion = "sa-east"
)

type SessionResponseReleaseReason string

const (
	SessionResponseReleaseReasonUserRequested     SessionResponseReleaseReason = "user_requested"
	SessionResponseReleaseReasonTimeout           SessionResponseReleaseReason = "timeout"
	SessionResponseReleaseReasonInactivityTimeout SessionResponseReleaseReason = "inactivity_timeout"
	SessionResponseReleaseReasonCreationTimeout   SessionResponseReleaseReason = "creation_timeout"
	SessionResponseReleaseReasonStartupFailed     SessionResponseReleaseReason = "startup_failed"
	SessionResponseReleaseReasonBrowserClosed     SessionResponseReleaseReason = "browser_closed"
	SessionResponseReleaseReasonBrowserCrashed    SessionResponseReleaseReason = "browser_crashed"
)

type SessionResponseStatus string

const (
	SessionResponseStatusLive     SessionResponseStatus = "live"
	SessionResponseStatusReleased SessionResponseStatus = "released"
	SessionResponseStatusFailed   SessionResponseStatus = "failed"
)

type ComputerActionRequestVariant0Action string

const (
	ComputerActionRequestVariant0ActionMoveMouse ComputerActionRequestVariant0Action = "move_mouse"
)

type ComputerActionRequestVariant1Action string

const (
	ComputerActionRequestVariant1ActionClickMouse ComputerActionRequestVariant1Action = "click_mouse"
)

type ComputerActionRequestVariant1Button string

const (
	ComputerActionRequestVariant1ButtonLeft    ComputerActionRequestVariant1Button = "left"
	ComputerActionRequestVariant1ButtonRight   ComputerActionRequestVariant1Button = "right"
	ComputerActionRequestVariant1ButtonMiddle  ComputerActionRequestVariant1Button = "middle"
	ComputerActionRequestVariant1ButtonBack    ComputerActionRequestVariant1Button = "back"
	ComputerActionRequestVariant1ButtonForward ComputerActionRequestVariant1Button = "forward"
)

type ComputerActionRequestVariant1ClickType string

const (
	ComputerActionRequestVariant1ClickTypeDown  ComputerActionRequestVariant1ClickType = "down"
	ComputerActionRequestVariant1ClickTypeUp    ComputerActionRequestVariant1ClickType = "up"
	ComputerActionRequestVariant1ClickTypeClick ComputerActionRequestVariant1ClickType = "click"
)

type ComputerActionRequestVariant2Action string

const (
	ComputerActionRequestVariant2ActionDragMouse ComputerActionRequestVariant2Action = "drag_mouse"
)

type ComputerActionRequestVariant3Action string

const (
	ComputerActionRequestVariant3ActionScroll ComputerActionRequestVariant3Action = "scroll"
)

type ComputerActionRequestVariant4Action string

const (
	ComputerActionRequestVariant4ActionPressKey ComputerActionRequestVariant4Action = "press_key"
)

type ComputerActionRequestVariant5Action string

const (
	ComputerActionRequestVariant5ActionTypeText ComputerActionRequestVariant5Action = "type_text"
)

type ComputerActionRequestVariant6Action string

const (
	ComputerActionRequestVariant6ActionWait ComputerActionRequestVariant6Action = "wait"
)

type ComputerActionRequestVariant7Action string

const (
	ComputerActionRequestVariant7ActionTakeScreenshot ComputerActionRequestVariant7Action = "take_screenshot"
)

type ComputerActionRequestVariant8Action string

const (
	ComputerActionRequestVariant8ActionGetCursorPosition ComputerActionRequestVariant8Action = "get_cursor_position"
)

type CreateSessionRequestDeviceConfigDevice string

const (
	CreateSessionRequestDeviceConfigDeviceDesktop CreateSessionRequestDeviceConfigDevice = "desktop"
	CreateSessionRequestDeviceConfigDeviceMobile  CreateSessionRequestDeviceConfigDevice = "mobile"
)

type CreateSessionRequestSessionContextCookiesItemPriority string

const (
	CreateSessionRequestSessionContextCookiesItemPriorityLow    CreateSessionRequestSessionContextCookiesItemPriority = "Low"
	CreateSessionRequestSessionContextCookiesItemPriorityMedium CreateSessionRequestSessionContextCookiesItemPriority = "Medium"
	CreateSessionRequestSessionContextCookiesItemPriorityHigh   CreateSessionRequestSessionContextCookiesItemPriority = "High"
)

type CreateSessionRequestSessionContextCookiesItemSameSite string

const (
	CreateSessionRequestSessionContextCookiesItemSameSiteStrict CreateSessionRequestSessionContextCookiesItemSameSite = "Strict"
	CreateSessionRequestSessionContextCookiesItemSameSiteLax    CreateSessionRequestSessionContextCookiesItemSameSite = "Lax"
	CreateSessionRequestSessionContextCookiesItemSameSiteNone   CreateSessionRequestSessionContextCookiesItemSameSite = "None"
)

type CreateSessionRequestSessionContextCookiesItemSourceScheme string

const (
	CreateSessionRequestSessionContextCookiesItemSourceSchemeUnset     CreateSessionRequestSessionContextCookiesItemSourceScheme = "Unset"
	CreateSessionRequestSessionContextCookiesItemSourceSchemeNonSecure CreateSessionRequestSessionContextCookiesItemSourceScheme = "NonSecure"
	CreateSessionRequestSessionContextCookiesItemSourceSchemeSecure    CreateSessionRequestSessionContextCookiesItemSourceScheme = "Secure"
)

type ProfileListResponseProfilesItemStatus string

const (
	ProfileListResponseProfilesItemStatusUploading ProfileListResponseProfilesItemStatus = "UPLOADING"
	ProfileListResponseProfilesItemStatusReady     ProfileListResponseProfilesItemStatus = "READY"
	ProfileListResponseProfilesItemStatusFailed    ProfileListResponseProfilesItemStatus = "FAILED"
)

type SessionContextResponseCookiesItemPriority string

const (
	SessionContextResponseCookiesItemPriorityLow    SessionContextResponseCookiesItemPriority = "Low"
	SessionContextResponseCookiesItemPriorityMedium SessionContextResponseCookiesItemPriority = "Medium"
	SessionContextResponseCookiesItemPriorityHigh   SessionContextResponseCookiesItemPriority = "High"
)

type SessionContextResponseCookiesItemSameSite string

const (
	SessionContextResponseCookiesItemSameSiteStrict SessionContextResponseCookiesItemSameSite = "Strict"
	SessionContextResponseCookiesItemSameSiteLax    SessionContextResponseCookiesItemSameSite = "Lax"
	SessionContextResponseCookiesItemSameSiteNone   SessionContextResponseCookiesItemSameSite = "None"
)

type SessionContextResponseCookiesItemSourceScheme string

const (
	SessionContextResponseCookiesItemSourceSchemeUnset     SessionContextResponseCookiesItemSourceScheme = "Unset"
	SessionContextResponseCookiesItemSourceSchemeNonSecure SessionContextResponseCookiesItemSourceScheme = "NonSecure"
	SessionContextResponseCookiesItemSourceSchemeSecure    SessionContextResponseCookiesItemSourceScheme = "Secure"
)

type SessionCostResponseUsageBrowserUnit string

const (
	SessionCostResponseUsageBrowserUnitMinute SessionCostResponseUsageBrowserUnit = "minute"
	SessionCostResponseUsageBrowserUnitSecond SessionCostResponseUsageBrowserUnit = "second"
)

type SessionResponseDeviceConfigDevice string

const (
	SessionResponseDeviceConfigDeviceDesktop SessionResponseDeviceConfigDevice = "desktop"
	SessionResponseDeviceConfigDeviceMobile  SessionResponseDeviceConfigDevice = "mobile"
)

type SessionsResponseSessionsItemDeviceConfigDevice string

const (
	SessionsResponseSessionsItemDeviceConfigDeviceDesktop SessionsResponseSessionsItemDeviceConfigDevice = "desktop"
	SessionsResponseSessionsItemDeviceConfigDeviceMobile  SessionsResponseSessionsItemDeviceConfigDevice = "mobile"
)

type SessionsResponseSessionsItemProxySource string

const (
	SessionsResponseSessionsItemProxySourceSteel    SessionsResponseSessionsItemProxySource = "steel"
	SessionsResponseSessionsItemProxySourceExternal SessionsResponseSessionsItemProxySource = "external"
)

type SessionsResponseSessionsItemRegion string

const (
	SessionsResponseSessionsItemRegionLax         SessionsResponseSessionsItemRegion = "lax"
	SessionsResponseSessionsItemRegionOrd         SessionsResponseSessionsItemRegion = "ord"
	SessionsResponseSessionsItemRegionIad         SessionsResponseSessionsItemRegion = "iad"
	SessionsResponseSessionsItemRegionScl         SessionsResponseSessionsItemRegion = "scl"
	SessionsResponseSessionsItemRegionFra         SessionsResponseSessionsItemRegion = "fra"
	SessionsResponseSessionsItemRegionNrt         SessionsResponseSessionsItemRegion = "nrt"
	SessionsResponseSessionsItemRegionUsEast      SessionsResponseSessionsItemRegion = "us-east"
	SessionsResponseSessionsItemRegionUsWest      SessionsResponseSessionsItemRegion = "us-west"
	SessionsResponseSessionsItemRegionUsCentral   SessionsResponseSessionsItemRegion = "us-central"
	SessionsResponseSessionsItemRegionEuWest      SessionsResponseSessionsItemRegion = "eu-west"
	SessionsResponseSessionsItemRegionEuCentral   SessionsResponseSessionsItemRegion = "eu-central"
	SessionsResponseSessionsItemRegionApNortheast SessionsResponseSessionsItemRegion = "ap-northeast"
	SessionsResponseSessionsItemRegionApSoutheast SessionsResponseSessionsItemRegion = "ap-southeast"
	SessionsResponseSessionsItemRegionSaEast      SessionsResponseSessionsItemRegion = "sa-east"
)

type SessionsResponseSessionsItemReleaseReason string

const (
	SessionsResponseSessionsItemReleaseReasonUserRequested     SessionsResponseSessionsItemReleaseReason = "user_requested"
	SessionsResponseSessionsItemReleaseReasonTimeout           SessionsResponseSessionsItemReleaseReason = "timeout"
	SessionsResponseSessionsItemReleaseReasonInactivityTimeout SessionsResponseSessionsItemReleaseReason = "inactivity_timeout"
	SessionsResponseSessionsItemReleaseReasonCreationTimeout   SessionsResponseSessionsItemReleaseReason = "creation_timeout"
	SessionsResponseSessionsItemReleaseReasonStartupFailed     SessionsResponseSessionsItemReleaseReason = "startup_failed"
	SessionsResponseSessionsItemReleaseReasonBrowserClosed     SessionsResponseSessionsItemReleaseReason = "browser_closed"
	SessionsResponseSessionsItemReleaseReasonBrowserCrashed    SessionsResponseSessionsItemReleaseReason = "browser_crashed"
)

type SessionsResponseSessionsItemStatus string

const (
	SessionsResponseSessionsItemStatusLive     SessionsResponseSessionsItemStatus = "live"
	SessionsResponseSessionsItemStatusReleased SessionsResponseSessionsItemStatus = "released"
	SessionsResponseSessionsItemStatusFailed   SessionsResponseSessionsItemStatus = "failed"
)

type GetSessionsStatus string

const (
	GetSessionsStatusLive     GetSessionsStatus = "live"
	GetSessionsStatusReleased GetSessionsStatus = "released"
	GetSessionsStatusFailed   GetSessionsStatus = "failed"
)

type CaptchaSolveResponse struct {
	// Response message
	Message *string `json:"message,omitempty"`
	// Whether the action was successful
	Success bool `json:"success"`
}

type CaptchaState struct {
	// Timestamp when the state was created
	Created *float64 `json:"created,omitempty"`
	// Whether a captcha is currently being solved
	IsSolvingCaptcha bool `json:"isSolvingCaptcha"`
	// Timestamp when the state was last updated
	LastUpdated *float64 `json:"lastUpdated,omitempty"`
	// The page ID where the captcha is located
	PageID string `json:"pageId"`
	// Array of captcha tasks
	Tasks []interface{} `json:"tasks"`
	// The URL where the captcha is located
	URL string `json:"url"`
}

type SessionComputerResponse struct {
	// Base64 encoded screenshot if requested
	Base64Image *string `json:"base64_image,omitempty"`
	// Error message if action failed
	Error *string `json:"error,omitempty"`
	// Output message from the action
	Output *string `json:"output,omitempty"`
	// System information
	System *string `json:"system,omitempty"`
}

// SessionCreateParams Request body schema for creating a new browser session.
type SessionCreateParams struct {
	// Block ads in the browser session. Default is false.
	BlockAds *bool `json:"blockAds,omitempty"`
	// PEM-encoded root CA certificates to trust in this session. THIS IS CURRENTLY AN EXPERIMENTAL FEATURE.
	CaCertificates *[]string `json:"caCertificates,omitempty"`
	// Number of sessions to create concurrently (check your plan limit)
	Concurrency *int64 `json:"concurrency,omitempty"`
	// Configuration for session credentials
	Credentials *Credentials `json:"credentials,omitempty"`
	// Configuration for the debug URL and session viewer. Controls interaction capabilities, cursor visibility, and other debug-related settings.
	DebugConfig *DebugConfig `json:"debugConfig,omitempty"`
	// Device configuration for the session. Specify 'mobile' for mobile device fingerprints and configurations.
	DeviceConfig *DeviceConfig `json:"deviceConfig,omitempty"`
	// Viewport and browser window dimensions for the session. Mobile sessions require dimensions of at least 508x1074; smaller mobile dimensions are rejected with a 400 response.
	Dimensions *Dimensions `json:"dimensions,omitempty"`
	// Enable experimental features for the session.
	ExperimentalFeatures *[]string `json:"experimentalFeatures,omitempty"`
	// Array of extension IDs to install in the session. Use ['all_ext'] to install all uploaded extensions.
	ExtensionIDs *[]string `json:"extensionIds,omitempty"`
	// Launch the browser in fullscreen mode, covering the full screen with no Chrome UI. Default is false.
	Fullscreen *bool `json:"fullscreen,omitempty"`
	// Enable headless browser mode (disable Headful mode)
	Headless *bool `json:"headless,omitempty"`
	// Inactivity timeout in milliseconds. When set, the session is released if no CDP command or remote input is received for this duration, even if `timeout` has not yet elapsed. Note that `timeout` remains the hard cap on session lifetime: if `inactivityTimeout` is greater than or equal to the effective `timeout`, it has no effect since `timeout` always elapses first. Omit to disable.
	InactivityTimeout *int64 `json:"inactivityTimeout,omitempty"`
	// Enable Selenium mode for the browser session (default is false). Use this when you plan to connect to the browser session via Selenium.
	IsSelenium *bool `json:"isSelenium,omitempty"`
	// The namespace the session should be created against. Defaults to "default".
	Namespace *string `json:"namespace,omitempty"`
	// Enable bandwidth optimizations. Passing true enables all flags (except hosts/patterns). Object allows granular control.
	OptimizeBandwidth *SessionCreateParamsOptimizeBandwidth `json:"optimizeBandwidth,omitempty"`
	// This flag will persist the user profile for the session.
	PersistProfile *bool `json:"persistProfile,omitempty"`
	// This flag will set the profile for the session.
	ProfileID *string `json:"profileId,omitempty"`
	// The project to create the session in. When provided, the session namespace is resolved from the project.
	ProjectID *string `json:"projectId,omitempty"`
	// Custom proxy URL for the browser session. Overrides useProxy, disabling Steel-provided proxies in favor of your specified proxy. Format: http(s)://username:password@hostname:port
	ProxyURL *string `json:"proxyUrl,omitempty"`
	// The desired region for the session. Available: us-east, us-west, us-central, eu-west, eu-central, ap-northeast, ap-southeast, sa-east. Legacy codes (iad, lax, ord) are also accepted.
	Region *interface{} `json:"region,omitempty"`
	// Session context data to be used in the created session. Sessions will start with an empty context by default.
	SessionContext *SessionContext2 `json:"sessionContext,omitempty"`
	// Optional custom UUID for the session
	SessionID *string `json:"sessionId,omitempty"`
	// Enable automatic captcha solving. Default is false.
	SolveCaptcha *bool `json:"solveCaptcha,omitempty"`
	// Stealth configuration for the session
	StealthConfig *StealthConfig `json:"stealthConfig,omitempty"`
	// Session timeout duration in milliseconds. Default is 300000 (5 minutes).
	Timeout  *int64       `json:"timeout,omitempty"`
	UseProxy *interface{} `json:"useProxy,omitempty"`
	// Custom user agent string for the browser session
	UserAgent *string `json:"userAgent,omitempty"`
}

type Credential struct {
	// Date and time the credential was created
	CreatedAt time.Time `json:"createdAt"`
	// Label for the credential
	Label *string `json:"label,omitempty"`
	// The namespace the credential is stored against. Defaults to "default".
	Namespace *string `json:"namespace,omitempty"`
	// Website origin the credential is for
	Origin *string `json:"origin,omitempty"`
	// Date and time the credential was last updated
	UpdatedAt time.Time `json:"updatedAt"`
	// Value for the credential
	Value map[string]string `json:"value"`
}

type CredentialCreateParams struct {
	// Label for the credential
	Label *string `json:"label,omitempty"`
	// The namespace the credential is stored against. Defaults to "default".
	Namespace *string `json:"namespace,omitempty"`
	// Website origin the credential is for
	Origin *string `json:"origin,omitempty"`
	// Project to store the credential in.
	ProjectID *string `json:"projectId,omitempty"`
	// Value for the credential
	Value map[string]string `json:"value"`
}

type CredentialListResponse struct {
	Credentials []Credential2 `json:"credentials"`
}

type CredentialCreateResponse struct {
	// Date and time the credential was created
	CreatedAt time.Time `json:"createdAt"`
	// Label for the credential
	Label *string `json:"label,omitempty"`
	// The namespace the credential is stored against. Defaults to "default".
	Namespace *string `json:"namespace,omitempty"`
	// Website origin the credential is for
	Origin *string `json:"origin,omitempty"`
	// Date and time the credential was last updated
	UpdatedAt time.Time `json:"updatedAt"`
}

type CredentialDeleteParams struct {
	// The namespace the credential is stored against. Defaults to "default".
	Namespace *string `json:"namespace,omitempty"`
	// Website origin the credential is for
	Origin string `json:"origin"`
	// Project to delete the credential from.
	ProjectID *string `json:"projectId,omitempty"`
}

type CredentialUpdateParams struct {
	// Label for the credential
	Label *string `json:"label,omitempty"`
	// The namespace the credential is stored against. Defaults to "default".
	Namespace *string `json:"namespace,omitempty"`
	// Website origin the credential is for
	Origin *string `json:"origin,omitempty"`
	// Project to update the credential in.
	ProjectID *string `json:"projectId,omitempty"`
	// Value for the credential
	Value *map[string]string `json:"value,omitempty"`
}

type ExtensionDeleteAllResponse struct {
	Message string `json:"message"`
}

type DetailsResponse struct {
	Plan string `json:"plan"`
}

// ErrorResponse An error response from the API
type ErrorResponse struct {
	Context    *[]Context `json:"context,omitempty"`
	Error      *string    `json:"error,omitempty"`
	LinkToDocs *string    `json:"linkToDocs,omitempty"`
	Message    string     `json:"message"`
}

type Extension struct {
	// Creation timestamp
	CreatedAt string `json:"createdAt"`
	// Unique extension identifier (e.g., ext_12345)
	ID string `json:"id"`
	// Extension name
	Name string `json:"name"`
	// Last update timestamp
	UpdatedAt string `json:"updatedAt"`
}

// ExtensionListResponse Response containing a list of extensions for the organization
type ExtensionListResponse struct {
	// Total number of extensions
	Count float64 `json:"count"`
	// List of extensions for the organization
	Extensions []Extension2 `json:"extensions"`
}

type ExtensionUpdateResponse struct {
	// Creation timestamp
	CreatedAt string `json:"createdAt"`
	// Unique extension identifier (e.g., ext_12345)
	ID string `json:"id"`
	// Extension name
	Name string `json:"name"`
	// Last update timestamp
	UpdatedAt string `json:"updatedAt"`
}

type File struct {
	// Timestamp when the file was created
	LastModified time.Time `json:"lastModified"`
	// Path to the file in the storage system
	Path string `json:"path"`
	// Size of the file in bytes
	Size float64 `json:"size"`
}

type FileUploadParams struct {
	// The file to upload (binary) or URL string to download from
	File FileUpload `json:"file"`
	// Path to the file in the storage system
	Path *string `json:"path,omitempty"`
}

type SessionAgentLogsResponse struct {
	Events  []map[string]interface{} `json:"events"`
	HasMore bool                     `json:"hasMore"`
	Total   float64                  `json:"total"`
}

type SessionAgentTracesResponse struct {
	Events  []map[string]interface{} `json:"events"`
	HasMore bool                     `json:"hasMore"`
	Total   float64                  `json:"total"`
}

type SessionLogsResponse struct {
	Events  []Event `json:"events"`
	HasMore bool    `json:"hasMore"`
	Total   float64 `json:"total"`
}

// WellknownListResponse A JSON Web Key Set (JWKS) containing a list of public keys.
type WellknownListResponse struct {
	// A list of JSON Web Keys.
	Keys []Key `json:"keys"`
}

type Fileslist struct {
	// Array of files for the current page
	Data []Data `json:"data"`
}

type ClientPdfParams struct {
	// Delay before generating the PDF (in milliseconds)
	Delay *float64 `json:"delay,omitempty"`
	// Project to execute the PDF generation in.
	ProjectID *string `json:"projectId,omitempty"`
	// The desired region for the action to be performed in
	Region *interface{} `json:"region,omitempty"`
	// URL of the webpage to convert to PDF
	URL string `json:"url"`
	// Use a Steel-provided residential proxy for generating the PDF
	UseProxy *bool `json:"useProxy,omitempty"`
}

type PdfResponse struct {
	// URL where the PDF is hosted
	URL string `json:"url"`
}

type ProfileCreateResponse struct {
	// The date and time when the profile was created
	CreatedAt time.Time `json:"createdAt"`
	// The credentials configuration associated with the profile
	CredentialsConfig interface{} `json:"credentialsConfig"`
	// The dimensions associated with the profile
	Dimensions *Dimensions2 `json:"dimensions,omitempty"`
	// The extension IDs associated with the profile
	ExtensionIDs *[]string `json:"extensionIds,omitempty"`
	// The fingerprint associated with the profile
	Fingerprint *Fingerprint `json:"fingerprint,omitempty"`
	// The unique identifier for the profile
	ID string `json:"id"`
	// The project ID associated with the profile
	ProjectID *string `json:"projectId,omitempty"`
	// The last session ID associated with the profile
	SourceSessionID *string `json:"sourceSessionId,omitempty"`
	// The status of the profile
	Status ProfileStatus `json:"status"`
	// The date and time when the profile was last updated
	UpdatedAt time.Time `json:"updatedAt"`
	// The proxy configuration associated with the profile
	UseProxyConfig interface{} `json:"useProxyConfig"`
	// The user agent associated with the profile
	UserAgent *string `json:"userAgent,omitempty"`
}

type ProfileCreateParams struct {
	// The dimensions associated with the profile
	Dimensions *Dimensions3 `json:"dimensions,omitempty"`
	// Project to create the profile in
	ProjectID *string `json:"projectId,omitempty"`
	// The proxy associated with the profile
	ProxyURL *string `json:"proxyUrl,omitempty"`
	// The user agent associated with the profile
	UserAgent *string `json:"userAgent,omitempty"`
	// The user data directory associated with the profile
	UserDataDir FileUpload `json:"userDataDir"`
}

type ProfileListResponse struct {
	// The total number of profiles
	Count float64 `json:"count"`
	// The list of profiles
	Profiles []Profile2 `json:"profiles"`
}

// SessionReleaseResponse Response for releasing a single session.
type SessionReleaseResponse struct {
	// Details about the outcome of the release operation
	Message string `json:"message"`
	// Indicates if the session was successfully released
	Success bool `json:"success"`
}

// SessionReleaseAllResponse Response for releasing multiple sessions.
type SessionReleaseAllResponse struct {
	// Details about the outcome of the release operation
	Message string `json:"message"`
	// Indicates if the sessions were successfully released
	Success bool `json:"success"`
}

type ClientScrapeParams struct {
	// Delay before scraping (in milliseconds)
	Delay *float64 `json:"delay,omitempty"`
	// Desired format(s) for the scraped content. Default is `html`.
	Format *[]ScrapeRequestFormatItem `json:"format,omitempty"`
	// Include a PDF in the response
	Pdf *bool `json:"pdf,omitempty"`
	// Project to execute the scrape in.
	ProjectID *string `json:"projectId,omitempty"`
	// The desired region for the action to be performed in
	Region *interface{} `json:"region,omitempty"`
	// Include a screenshot in the response
	Screenshot *bool `json:"screenshot,omitempty"`
	// URL of the webpage to scrape
	URL string `json:"url"`
	// Use a Steel-provided residential proxy for the scrape
	UseProxy *bool `json:"useProxy,omitempty"`
}

// ScrapeResponse Response from a successful scrape request
type ScrapeResponse struct {
	Content    Content     `json:"content"`
	Links      []Link      `json:"links"`
	Metadata   Metadata    `json:"metadata"`
	Pdf        *Pdf        `json:"pdf,omitempty"`
	Screenshot *Screenshot `json:"screenshot,omitempty"`
}

type ClientScreenshotParams struct {
	// Delay before capturing the screenshot (in milliseconds)
	Delay *float64 `json:"delay,omitempty"`
	// Capture the full page screenshot. Default is `false`.
	FullPage *bool `json:"fullPage,omitempty"`
	// Project to execute the screenshot in.
	ProjectID *string `json:"projectId,omitempty"`
	// The desired region for the action to be performed in
	Region *interface{} `json:"region,omitempty"`
	// URL of the webpage to capture
	URL string `json:"url"`
	// Use a Steel-provided residential proxy for capturing the screenshot
	UseProxy *bool `json:"useProxy,omitempty"`
}

type ScreenshotResponse struct {
	// URL where the screenshot is hosted
	URL string `json:"url"`
}

// SessionContext Session context data returned from a browser session.
type SessionContext struct {
	// Cookies to initialize in the session
	Cookies *[]Cookie `json:"cookies,omitempty"`
	// Domain-specific indexedDB items to initialize in the session
	IndexedDB *map[string][]IndexedDB `json:"indexedDB,omitempty"`
	// Domain-specific localStorage items to initialize in the session
	LocalStorage *map[string]map[string]string `json:"localStorage,omitempty"`
	// Domain-specific sessionStorage items to initialize in the session
	SessionStorage *map[string]map[string]string `json:"sessionStorage,omitempty"`
}

// SessionCostResponse Session cost breakdown in US dollars
type SessionCostResponse struct {
	// Currency used for the cost values
	Currency SessionCostResponseCurrency `json:"currency"`
	// Session ID
	ID string `json:"id"`
	// Exact total session cost in US dollars (e.g. 0.031905), with no rounding to whole cents. Rounded only to micro-dollar precision (6 decimal places) for precise billing passthrough.
	TotalCost float64 `json:"totalCost"`
	// Cost unit. Values are expressed in US dollars.
	Unit SessionCostResponseUnit `json:"unit"`
	// Billable usage inputs used to calculate the costs
	Usage Usage `json:"usage"`
}

type SessionLiveDetailsResponse struct {
	Pages                      []Page `json:"pages"`
	SessionViewerFullscreenURL string `json:"sessionViewerFullscreenUrl"`
	SessionViewerURL           string `json:"sessionViewerUrl"`
	WsURL                      string `json:"wsUrl"`
}

// Session Represents the data structure for a browser session, including its configuration and status.
type Session struct {
	// Timestamp when the session started
	CreatedAt time.Time `json:"createdAt"`
	// Amount of credits consumed by the session
	CreditsUsed int64 `json:"creditsUsed"`
	// Configuration for the debug URL and session viewer. Controls interaction capabilities and cursor visibility.
	DebugConfig *DebugConfig2 `json:"debugConfig,omitempty"`
	// URL for debugging the session
	DebugURL string `json:"debugUrl"`
	// Device configuration for the session
	DeviceConfig *DeviceConfig2 `json:"deviceConfig,omitempty"`
	// Viewport and browser window dimensions for the session
	Dimensions Dimensions4 `json:"dimensions"`
	// Duration of the session in milliseconds
	Duration int64 `json:"duration"`
	// Number of events processed in the session
	EventCount int64 `json:"eventCount"`
	// Launch the browser in fullscreen mode, covering the full screen with no Chrome UI.
	Fullscreen *bool `json:"fullscreen,omitempty"`
	// Indicates if the session is headless or headful
	Headless *bool `json:"headless,omitempty"`
	// Unique identifier for the session
	ID string `json:"id"`
	// Inactivity timeout in milliseconds, if one was set when the session was created
	InactivityTimeout *int64 `json:"inactivityTimeout,omitempty"`
	// Indicates if Selenium is used in the session
	IsSelenium *bool `json:"isSelenium,omitempty"`
	// Bandwidth optimizations that were applied to the session.
	OptimizeBandwidth OptimizeBandwidth `json:"optimizeBandwidth"`
	// This flag will persist the profile for the session.
	PersistProfile *bool `json:"persistProfile,omitempty"`
	// The ID of the profile associated with the session
	ProfileID *string `json:"profileId,omitempty"`
	// The project associated with the session
	ProjectID *string `json:"projectId,omitempty"`
	// Amount of data transmitted through the proxy
	ProxyBytesUsed int64 `json:"proxyBytesUsed"`
	// Source of the proxy used for the session
	ProxySource *SessionResponseProxySource `json:"proxySource,omitempty"`
	// The region where the session was created.
	Region *SessionResponseRegion `json:"region,omitempty"`
	// Why the session reached a terminal state. Null while the session is live, or when the reason is unknown (e.g. sessions created before this was tracked). One of: user_requested (released via the API/SDK), timeout (hard `timeout` elapsed), inactivity_timeout (no activity for the configured window), creation_timeout (never started in time), startup_failed (could not be dispatched), browser_closed (the browser or agent closed itself — not a crash), browser_crashed (the browser crashed or its machine became unresponsive).
	ReleaseReason *SessionResponseReleaseReason `json:"releaseReason,omitempty"`
	// URL to view session details
	SessionViewerURL string `json:"sessionViewerUrl"`
	// Indicates if captcha solving is enabled
	SolveCaptcha *bool `json:"solveCaptcha,omitempty"`
	// Status of the session
	Status SessionResponseStatus `json:"status"`
	// Stealth configuration for the session
	StealthConfig *StealthConfig2 `json:"stealthConfig,omitempty"`
	// Session timeout duration in milliseconds
	Timeout int64 `json:"timeout"`
	// User agent string used in the session
	UserAgent *string `json:"userAgent,omitempty"`
	// URL for the session's WebSocket connection
	WebsocketURL string `json:"websocketUrl"`
}

// Sessionslist Response containing a list of browser sessions with pagination details.
type Sessionslist struct {
	// Cursor for the next page of results. Null if no more pages.
	NextCursor *string `json:"nextCursor,omitempty"`
	// List of browser sessions
	Sessions []Session2 `json:"sessions"`
	// Total number of sessions matching the query. Only included for filtered queries (e.g. status=live).
	TotalCount *int64 `json:"totalCount,omitempty"`
}

// CaptchaSolveParams If no fields are provided, all detected captchas will be solved
type CaptchaSolveParams struct {
	// The page ID where the captcha is located
	PageID *string `json:"pageId,omitempty"`
	// The task ID of the specific captcha to solve
	TaskID *string `json:"taskId,omitempty"`
	// The URL where the captcha is located
	URL *string `json:"url,omitempty"`
}

type CaptchaSolveImageParams struct {
	// XPath to the captcha image element
	ImageXPath string `json:"imageXPath"`
	// XPath to the captcha input element
	InputXPath string `json:"inputXPath"`
	// URL where the captcha is located. Defaults to the current page URL
	URL *string `json:"url,omitempty"`
}

type CredentialDeleteResponse struct {
	Success bool `json:"success"`
}

type ExtensionUploadParams struct {
	// Extension .zip/.crx file
	File *FileUpload `json:"file,omitempty"`
	// Extension URL
	URL *string `json:"url,omitempty"`
}

type MoveMouse struct {
	Action ComputerActionRequestVariant0Action `json:"action"`
	// X and Y coordinates [x, y]
	Coordinates []float64 `json:"coordinates"`
	// Keys to hold while moving
	HoldKeys *[]string `json:"hold_keys,omitempty"`
	// Whether to take a screenshot after the action
	Screenshot *bool `json:"screenshot,omitempty"`
}

type ClickMouse struct {
	Action ComputerActionRequestVariant1Action `json:"action"`
	// Mouse button to click. Defaults to 'left'
	Button *ComputerActionRequestVariant1Button `json:"button,omitempty"`
	// Type of click (down, up, or click). Defaults to 'click'
	ClickType *ComputerActionRequestVariant1ClickType `json:"click_type,omitempty"`
	// X and Y coordinates [x, y]
	Coordinates *[]float64 `json:"coordinates,omitempty"`
	// Keys to hold while clicking
	HoldKeys *[]string `json:"hold_keys,omitempty"`
	// Number of clicks. Defaults to 1
	NumClicks *float64 `json:"num_clicks,omitempty"`
	// Whether to take a screenshot after the action
	Screenshot *bool `json:"screenshot,omitempty"`
}

type DragMouse struct {
	Action ComputerActionRequestVariant2Action `json:"action"`
	// Keys to hold while dragging
	HoldKeys *[]string `json:"hold_keys,omitempty"`
	// Array of [x, y] coordinate pairs
	Path [][]float64 `json:"path"`
	// Whether to take a screenshot after the action
	Screenshot *bool `json:"screenshot,omitempty"`
}

type Scroll struct {
	Action ComputerActionRequestVariant3Action `json:"action"`
	// X and Y coordinates [x, y]
	Coordinates *[]float64 `json:"coordinates,omitempty"`
	// Horizontal scroll amount. Defaults to 0
	DeltaX *float64 `json:"delta_x,omitempty"`
	// Vertical scroll amount. Defaults to 0
	DeltaY *float64 `json:"delta_y,omitempty"`
	// Keys to hold while scrolling
	HoldKeys *[]string `json:"hold_keys,omitempty"`
	// Whether to take a screenshot after the action
	Screenshot *bool `json:"screenshot,omitempty"`
}

type PressKey struct {
	Action ComputerActionRequestVariant4Action `json:"action"`
	// Duration to hold keys in seconds
	Duration *float64 `json:"duration,omitempty"`
	// Keys to press
	Keys []string `json:"keys"`
	// Whether to take a screenshot after the action
	Screenshot *bool `json:"screenshot,omitempty"`
}

type TypeText struct {
	Action ComputerActionRequestVariant5Action `json:"action"`
	// Keys to hold while typing
	HoldKeys *[]string `json:"hold_keys,omitempty"`
	// Whether to take a screenshot after the action
	Screenshot *bool `json:"screenshot,omitempty"`
	// Text to type
	Text string `json:"text"`
}

type Wait struct {
	Action ComputerActionRequestVariant6Action `json:"action"`
	// Duration to wait in seconds
	Duration float64 `json:"duration"`
	// Whether to take a screenshot after the action
	Screenshot *bool `json:"screenshot,omitempty"`
}

type TakeScreenshot struct {
	Action ComputerActionRequestVariant7Action `json:"action"`
}

type GetCursorPosition struct {
	Action ComputerActionRequestVariant8Action `json:"action"`
}

// Credentials Configuration for session credentials
type Credentials struct {
	AutoSubmit  *bool `json:"autoSubmit,omitempty"`
	BlurFields  *bool `json:"blurFields,omitempty"`
	ExactOrigin *bool `json:"exactOrigin,omitempty"`
}

// DebugConfig Configuration for the debug URL and session viewer. Controls interaction capabilities, cursor visibility, and other debug-related settings.
type DebugConfig struct {
	// Allow interaction with the browser session via the debug URL viewer. When false, the session viewer will be view-only. Default is true.
	Interactive *bool `json:"interactive,omitempty"`
	// Show the OS-level mouse cursor in the WebRTC stream (headful mode only). When false, the system cursor will not be rendered in the stream. Default is true.
	SystemCursor *bool `json:"systemCursor,omitempty"`
}

// DeviceConfig Device configuration for the session. Specify 'mobile' for mobile device fingerprints and configurations.
type DeviceConfig struct {
	Device *CreateSessionRequestDeviceConfigDevice `json:"device,omitempty"`
}

// Dimensions Viewport and browser window dimensions for the session. Mobile sessions require dimensions of at least 508x1074; smaller mobile dimensions are rejected with a 400 response.
type Dimensions struct {
	// Height of the session
	Height int64 `json:"height"`
	// Width of the session
	Width int64 `json:"width"`
}

// SessionContextCookiePartitionKey The partition key of the cookie
type SessionContextCookiePartitionKey struct {
	// Indicates if the cookie has any ancestors that are cross-site to the topLevelSite.
	HasCrossSiteAncestor bool `json:"hasCrossSiteAncestor"`
	// The site of the top-level URL the browser was visiting at the start of the request to the endpoint that set the cookie.
	TopLevelSite string `json:"topLevelSite"`
}

// SessionContextCookie Cookies to initialize in the session
type SessionContextCookie struct {
	// The domain of the cookie
	Domain *string `json:"domain,omitempty"`
	// The expiration date of the cookie
	Expires *float64 `json:"expires,omitempty"`
	// Whether the cookie is HTTP only
	HTTPOnly *bool `json:"httpOnly,omitempty"`
	// The name of the cookie
	Name string `json:"name"`
	// The partition key of the cookie
	PartitionKey *SessionContextCookiePartitionKey `json:"partitionKey,omitempty"`
	// The path of the cookie
	Path *string `json:"path,omitempty"`
	// The priority of the cookie
	Priority *CreateSessionRequestSessionContextCookiesItemPriority `json:"priority,omitempty"`
	// Whether the cookie is a same party cookie
	SameParty *bool `json:"sameParty,omitempty"`
	// The same site attribute of the cookie
	SameSite *CreateSessionRequestSessionContextCookiesItemSameSite `json:"sameSite,omitempty"`
	// Whether the cookie is secure
	Secure *bool `json:"secure,omitempty"`
	// Whether the cookie is a session cookie
	Session *bool `json:"session,omitempty"`
	// The size of the cookie
	Size *float64 `json:"size,omitempty"`
	// The source port of the cookie
	SourcePort *float64 `json:"sourcePort,omitempty"`
	// The source scheme of the cookie
	SourceScheme *CreateSessionRequestSessionContextCookiesItemSourceScheme `json:"sourceScheme,omitempty"`
	// The URL of the cookie
	URL *string `json:"url,omitempty"`
	// The value of the cookie
	Value string `json:"value"`
}

type SessionContextIndexedDBDataRecordBlobFile struct {
	BlobNumber   float64    `json:"blobNumber"`
	Filename     *string    `json:"filename,omitempty"`
	LastModified *time.Time `json:"lastModified,omitempty"`
	MimeType     string     `json:"mimeType"`
	Path         *string    `json:"path,omitempty"`
	Size         float64    `json:"size"`
}

type SessionContextIndexedDBDataRecord struct {
	BlobFiles *[]SessionContextIndexedDBDataRecordBlobFile `json:"blobFiles,omitempty"`
	Key       interface{}                                  `json:"key"`
	Value     interface{}                                  `json:"value"`
}

type SessionContextIndexedDBData struct {
	ID      float64                             `json:"id"`
	Name    string                              `json:"name"`
	Records []SessionContextIndexedDBDataRecord `json:"records"`
}

// SessionContextIndexedDB Domain-specific indexedDB items to initialize in the session
type SessionContextIndexedDB struct {
	Data []SessionContextIndexedDBData `json:"data"`
	ID   float64                       `json:"id"`
	Name string                        `json:"name"`
}

// SessionContext2 Session context data to be used in the created session. Sessions will start with an empty context by default.
type SessionContext2 struct {
	// Cookies to initialize in the session
	Cookies *[]SessionContextCookie `json:"cookies,omitempty"`
	// Domain-specific indexedDB items to initialize in the session
	IndexedDB *map[string][]SessionContextIndexedDB `json:"indexedDB,omitempty"`
	// Domain-specific localStorage items to initialize in the session
	LocalStorage *map[string]map[string]string `json:"localStorage,omitempty"`
	// Domain-specific sessionStorage items to initialize in the session
	SessionStorage *map[string]map[string]string `json:"sessionStorage,omitempty"`
}

// StealthConfig Stealth configuration for the session
type StealthConfig struct {
	// When true, captchas will be automatically solved when detected. When false, use the solve endpoints to manually initiate solving.
	AutoCaptchaSolving *bool `json:"autoCaptchaSolving,omitempty"`
	// This flag will make the browser act more human-like by moving the mouse in a more natural way.
	HumanizeInteractions *bool `json:"humanizeInteractions,omitempty"`
	// This flag will skip the fingerprint generation for the session.
	SkipFingerprintInjection *bool `json:"skipFingerprintInjection,omitempty"`
}

type Credential2 struct {
	// Date and time the credential was created
	CreatedAt time.Time `json:"createdAt"`
	// Label for the credential
	Label *string `json:"label,omitempty"`
	// The namespace the credential is stored against. Defaults to "default".
	Namespace *string `json:"namespace,omitempty"`
	// Website origin the credential is for
	Origin *string `json:"origin,omitempty"`
	// Date and time the credential was last updated
	UpdatedAt time.Time `json:"updatedAt"`
}

type Context struct {
	Keyword string      `json:"keyword"`
	Message string      `json:"message"`
	Params  interface{} `json:"params"`
}

// Extension2 List of extensions for the organization
type Extension2 struct {
	// Creation timestamp
	CreatedAt string `json:"createdAt"`
	// Unique extension identifier (e.g., ext_12345)
	ID string `json:"id"`
	// Extension name
	Name string `json:"name"`
	// Last update timestamp
	UpdatedAt string `json:"updatedAt"`
}

type Event struct {
	ID        string    `json:"id"`
	Log       string    `json:"log"`
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
}

// Key A list of JSON Web Keys.
type Key struct {
	// The specific cryptographic algorithm used with the key (e.g., 'RS256').
	Alg *string `json:"alg,omitempty"`
	// The exponent for the RSA public key.
	E *string `json:"e,omitempty"`
	// The unique identifier for the key.
	Kid *string `json:"kid,omitempty"`
	// The key type, identifying the algorithm family (e.g., 'RSA').
	Kty *string `json:"kty,omitempty"`
	// The modulus for the RSA public key.
	N *string `json:"n,omitempty"`
	// The intended use of the public key (e.g., 'sig' for signature).
	Use *string `json:"use,omitempty"`
}

// Data Array of files for the current page
type Data struct {
	// Timestamp when the file was created
	LastModified time.Time `json:"lastModified"`
	// Path to the file in the storage system
	Path string `json:"path"`
	// Size of the file in bytes
	Size float64 `json:"size"`
}

// Dimensions2 The dimensions associated with the profile
type Dimensions2 struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

type FingerprintFingerprintBattery struct {
	Charging        *bool    `json:"charging,omitempty"`
	ChargingTime    *float64 `json:"chargingTime,omitempty"`
	DischargingTime *float64 `json:"dischargingTime,omitempty"`
	Level           *float64 `json:"level,omitempty"`
}

type FingerprintFingerprintMultimediaDevicesMicro struct {
	DeviceID *string `json:"deviceId,omitempty"`
	GroupID  *string `json:"groupId,omitempty"`
	Kind     *string `json:"kind,omitempty"`
	Label    *string `json:"label,omitempty"`
}

type FingerprintFingerprintMultimediaDevicesSpeaker struct {
	DeviceID *string `json:"deviceId,omitempty"`
	GroupID  *string `json:"groupId,omitempty"`
	Kind     *string `json:"kind,omitempty"`
	Label    *string `json:"label,omitempty"`
}

type FingerprintFingerprintMultimediaDevicesWebcam struct {
	DeviceID *string `json:"deviceId,omitempty"`
	GroupID  *string `json:"groupId,omitempty"`
	Kind     *string `json:"kind,omitempty"`
	Label    *string `json:"label,omitempty"`
}

type FingerprintFingerprintMultimediaDevices struct {
	Micros   []FingerprintFingerprintMultimediaDevicesMicro   `json:"micros"`
	Speakers []FingerprintFingerprintMultimediaDevicesSpeaker `json:"speakers"`
	Webcams  []FingerprintFingerprintMultimediaDevicesWebcam  `json:"webcams"`
}

type FingerprintFingerprintNavigatorExtraProperties struct {
	GlobalPrivacyControl *bool    `json:"globalPrivacyControl,omitempty"`
	InstalledApps        []string `json:"installedApps"`
	PdfViewerEnabled     *bool    `json:"pdfViewerEnabled,omitempty"`
	VendorFlavors        []string `json:"vendorFlavors"`
}

type FingerprintFingerprintNavigatorUserAgentDataBrand struct {
	Brand   *string `json:"brand,omitempty"`
	Version *string `json:"version,omitempty"`
}

type FingerprintFingerprintNavigatorUserAgentData struct {
	Brands   []FingerprintFingerprintNavigatorUserAgentDataBrand `json:"brands"`
	Mobile   *bool                                               `json:"mobile,omitempty"`
	Platform *string                                             `json:"platform,omitempty"`
}

type FingerprintFingerprintNavigator struct {
	AppCodeName         *string                                        `json:"appCodeName,omitempty"`
	AppName             *string                                        `json:"appName,omitempty"`
	AppVersion          *string                                        `json:"appVersion,omitempty"`
	DeviceMemory        *float64                                       `json:"deviceMemory,omitempty"`
	DoNotTrack          *string                                        `json:"doNotTrack,omitempty"`
	ExtraProperties     FingerprintFingerprintNavigatorExtraProperties `json:"extraProperties"`
	HardwareConcurrency *float64                                       `json:"hardwareConcurrency,omitempty"`
	Language            *string                                        `json:"language,omitempty"`
	Languages           []string                                       `json:"languages"`
	MaxTouchPoints      *float64                                       `json:"maxTouchPoints,omitempty"`
	Oscpu               *string                                        `json:"oscpu,omitempty"`
	Platform            *string                                        `json:"platform,omitempty"`
	Product             *string                                        `json:"product,omitempty"`
	ProductSub          *string                                        `json:"productSub,omitempty"`
	UserAgent           *string                                        `json:"userAgent,omitempty"`
	UserAgentData       FingerprintFingerprintNavigatorUserAgentData   `json:"userAgentData"`
	Vendor              *string                                        `json:"vendor,omitempty"`
	VendorSub           *string                                        `json:"vendorSub,omitempty"`
	Webdriver           *bool                                          `json:"webdriver,omitempty"`
}

type FingerprintFingerprintPluginsDataPluginMimeType struct {
	Description   *string `json:"description,omitempty"`
	EnabledPlugin *string `json:"enabledPlugin,omitempty"`
	Suffixes      *string `json:"suffixes,omitempty"`
	Type          *string `json:"type,omitempty"`
}

type FingerprintFingerprintPluginsDataPlugin struct {
	Description *string                                           `json:"description,omitempty"`
	Filename    *string                                           `json:"filename,omitempty"`
	MimeTypes   []FingerprintFingerprintPluginsDataPluginMimeType `json:"mimeTypes"`
	Name        *string                                           `json:"name,omitempty"`
}

type FingerprintFingerprintPluginsData struct {
	MimeTypes []string                                  `json:"mimeTypes"`
	Plugins   []FingerprintFingerprintPluginsDataPlugin `json:"plugins"`
}

type FingerprintFingerprintScreen struct {
	AvailHeight      *float64 `json:"availHeight,omitempty"`
	AvailLeft        *float64 `json:"availLeft,omitempty"`
	AvailTop         *float64 `json:"availTop,omitempty"`
	AvailWidth       *float64 `json:"availWidth,omitempty"`
	ClientHeight     *float64 `json:"clientHeight,omitempty"`
	ClientWidth      *float64 `json:"clientWidth,omitempty"`
	ColorDepth       *float64 `json:"colorDepth,omitempty"`
	DevicePixelRatio *float64 `json:"devicePixelRatio,omitempty"`
	HasHdr           *bool    `json:"hasHDR,omitempty"`
	Height           *float64 `json:"height,omitempty"`
	InnerHeight      *float64 `json:"innerHeight,omitempty"`
	InnerWidth       *float64 `json:"innerWidth,omitempty"`
	OuterHeight      *float64 `json:"outerHeight,omitempty"`
	OuterWidth       *float64 `json:"outerWidth,omitempty"`
	PageXOffset      *float64 `json:"pageXOffset,omitempty"`
	PageYOffset      *float64 `json:"pageYOffset,omitempty"`
	PixelDepth       *float64 `json:"pixelDepth,omitempty"`
	ScreenX          *float64 `json:"screenX,omitempty"`
	Width            *float64 `json:"width,omitempty"`
}

type FingerprintFingerprintVideoCard struct {
	Renderer *string `json:"renderer,omitempty"`
	Vendor   *string `json:"vendor,omitempty"`
}

type FingerprintFingerprint struct {
	AudioCodecs       map[string]string                       `json:"audioCodecs"`
	Battery           FingerprintFingerprintBattery           `json:"battery"`
	Fonts             []string                                `json:"fonts"`
	MockWebRtc        *bool                                   `json:"mockWebRTC,omitempty"`
	MultimediaDevices FingerprintFingerprintMultimediaDevices `json:"multimediaDevices"`
	Navigator         FingerprintFingerprintNavigator         `json:"navigator"`
	PluginsData       FingerprintFingerprintPluginsData       `json:"pluginsData"`
	Screen            FingerprintFingerprintScreen            `json:"screen"`
	Slim              *bool                                   `json:"slim,omitempty"`
	VideoCard         FingerprintFingerprintVideoCard         `json:"videoCard"`
	VideoCodecs       map[string]string                       `json:"videoCodecs"`
}

type FingerprintHeaders struct {
	Accept                  *string `json:"accept,omitempty"`
	AcceptEncoding          *string `json:"accept-encoding,omitempty"`
	AcceptLanguage          *string `json:"accept-language,omitempty"`
	Dnt                     *string `json:"dnt,omitempty"`
	SecChUa                 *string `json:"sec-ch-ua,omitempty"`
	SecChUaMobile           *string `json:"sec-ch-ua-mobile,omitempty"`
	SecChUaPlatform         *string `json:"sec-ch-ua-platform,omitempty"`
	SecFetchDest            *string `json:"sec-fetch-dest,omitempty"`
	SecFetchMode            *string `json:"sec-fetch-mode,omitempty"`
	SecFetchSite            *string `json:"sec-fetch-site,omitempty"`
	SecFetchUser            *string `json:"sec-fetch-user,omitempty"`
	UpgradeInsecureRequests *string `json:"upgrade-insecure-requests,omitempty"`
	UserAgent               *string `json:"user-agent,omitempty"`
}

// Fingerprint The fingerprint associated with the profile
type Fingerprint struct {
	Fingerprint FingerprintFingerprint `json:"fingerprint"`
	Headers     FingerprintHeaders     `json:"headers"`
}

// Dimensions3 The dimensions associated with the profile
type Dimensions3 struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// ProfileDimensions The dimensions associated with the profile
type ProfileDimensions struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

type ProfileFingerprintFingerprintBattery struct {
	Charging        *bool    `json:"charging,omitempty"`
	ChargingTime    *float64 `json:"chargingTime,omitempty"`
	DischargingTime *float64 `json:"dischargingTime,omitempty"`
	Level           *float64 `json:"level,omitempty"`
}

type ProfileFingerprintFingerprintMultimediaDevicesMicro struct {
	DeviceID *string `json:"deviceId,omitempty"`
	GroupID  *string `json:"groupId,omitempty"`
	Kind     *string `json:"kind,omitempty"`
	Label    *string `json:"label,omitempty"`
}

type ProfileFingerprintFingerprintMultimediaDevicesSpeaker struct {
	DeviceID *string `json:"deviceId,omitempty"`
	GroupID  *string `json:"groupId,omitempty"`
	Kind     *string `json:"kind,omitempty"`
	Label    *string `json:"label,omitempty"`
}

type ProfileFingerprintFingerprintMultimediaDevicesWebcam struct {
	DeviceID *string `json:"deviceId,omitempty"`
	GroupID  *string `json:"groupId,omitempty"`
	Kind     *string `json:"kind,omitempty"`
	Label    *string `json:"label,omitempty"`
}

type ProfileFingerprintFingerprintMultimediaDevices struct {
	Micros   []ProfileFingerprintFingerprintMultimediaDevicesMicro   `json:"micros"`
	Speakers []ProfileFingerprintFingerprintMultimediaDevicesSpeaker `json:"speakers"`
	Webcams  []ProfileFingerprintFingerprintMultimediaDevicesWebcam  `json:"webcams"`
}

type ProfileFingerprintFingerprintNavigatorExtraProperties struct {
	GlobalPrivacyControl *bool    `json:"globalPrivacyControl,omitempty"`
	InstalledApps        []string `json:"installedApps"`
	PdfViewerEnabled     *bool    `json:"pdfViewerEnabled,omitempty"`
	VendorFlavors        []string `json:"vendorFlavors"`
}

type ProfileFingerprintFingerprintNavigatorUserAgentDataBrand struct {
	Brand   *string `json:"brand,omitempty"`
	Version *string `json:"version,omitempty"`
}

type ProfileFingerprintFingerprintNavigatorUserAgentData struct {
	Brands   []ProfileFingerprintFingerprintNavigatorUserAgentDataBrand `json:"brands"`
	Mobile   *bool                                                      `json:"mobile,omitempty"`
	Platform *string                                                    `json:"platform,omitempty"`
}

type ProfileFingerprintFingerprintNavigator struct {
	AppCodeName         *string                                               `json:"appCodeName,omitempty"`
	AppName             *string                                               `json:"appName,omitempty"`
	AppVersion          *string                                               `json:"appVersion,omitempty"`
	DeviceMemory        *float64                                              `json:"deviceMemory,omitempty"`
	DoNotTrack          *string                                               `json:"doNotTrack,omitempty"`
	ExtraProperties     ProfileFingerprintFingerprintNavigatorExtraProperties `json:"extraProperties"`
	HardwareConcurrency *float64                                              `json:"hardwareConcurrency,omitempty"`
	Language            *string                                               `json:"language,omitempty"`
	Languages           []string                                              `json:"languages"`
	MaxTouchPoints      *float64                                              `json:"maxTouchPoints,omitempty"`
	Oscpu               *string                                               `json:"oscpu,omitempty"`
	Platform            *string                                               `json:"platform,omitempty"`
	Product             *string                                               `json:"product,omitempty"`
	ProductSub          *string                                               `json:"productSub,omitempty"`
	UserAgent           *string                                               `json:"userAgent,omitempty"`
	UserAgentData       ProfileFingerprintFingerprintNavigatorUserAgentData   `json:"userAgentData"`
	Vendor              *string                                               `json:"vendor,omitempty"`
	VendorSub           *string                                               `json:"vendorSub,omitempty"`
	Webdriver           *bool                                                 `json:"webdriver,omitempty"`
}

type ProfileFingerprintFingerprintPluginsDataPluginMimeType struct {
	Description   *string `json:"description,omitempty"`
	EnabledPlugin *string `json:"enabledPlugin,omitempty"`
	Suffixes      *string `json:"suffixes,omitempty"`
	Type          *string `json:"type,omitempty"`
}

type ProfileFingerprintFingerprintPluginsDataPlugin struct {
	Description *string                                                  `json:"description,omitempty"`
	Filename    *string                                                  `json:"filename,omitempty"`
	MimeTypes   []ProfileFingerprintFingerprintPluginsDataPluginMimeType `json:"mimeTypes"`
	Name        *string                                                  `json:"name,omitempty"`
}

type ProfileFingerprintFingerprintPluginsData struct {
	MimeTypes []string                                         `json:"mimeTypes"`
	Plugins   []ProfileFingerprintFingerprintPluginsDataPlugin `json:"plugins"`
}

type ProfileFingerprintFingerprintScreen struct {
	AvailHeight      *float64 `json:"availHeight,omitempty"`
	AvailLeft        *float64 `json:"availLeft,omitempty"`
	AvailTop         *float64 `json:"availTop,omitempty"`
	AvailWidth       *float64 `json:"availWidth,omitempty"`
	ClientHeight     *float64 `json:"clientHeight,omitempty"`
	ClientWidth      *float64 `json:"clientWidth,omitempty"`
	ColorDepth       *float64 `json:"colorDepth,omitempty"`
	DevicePixelRatio *float64 `json:"devicePixelRatio,omitempty"`
	HasHdr           *bool    `json:"hasHDR,omitempty"`
	Height           *float64 `json:"height,omitempty"`
	InnerHeight      *float64 `json:"innerHeight,omitempty"`
	InnerWidth       *float64 `json:"innerWidth,omitempty"`
	OuterHeight      *float64 `json:"outerHeight,omitempty"`
	OuterWidth       *float64 `json:"outerWidth,omitempty"`
	PageXOffset      *float64 `json:"pageXOffset,omitempty"`
	PageYOffset      *float64 `json:"pageYOffset,omitempty"`
	PixelDepth       *float64 `json:"pixelDepth,omitempty"`
	ScreenX          *float64 `json:"screenX,omitempty"`
	Width            *float64 `json:"width,omitempty"`
}

type ProfileFingerprintFingerprintVideoCard struct {
	Renderer *string `json:"renderer,omitempty"`
	Vendor   *string `json:"vendor,omitempty"`
}

type ProfileFingerprintFingerprint struct {
	AudioCodecs       map[string]string                              `json:"audioCodecs"`
	Battery           ProfileFingerprintFingerprintBattery           `json:"battery"`
	Fonts             []string                                       `json:"fonts"`
	MockWebRtc        *bool                                          `json:"mockWebRTC,omitempty"`
	MultimediaDevices ProfileFingerprintFingerprintMultimediaDevices `json:"multimediaDevices"`
	Navigator         ProfileFingerprintFingerprintNavigator         `json:"navigator"`
	PluginsData       ProfileFingerprintFingerprintPluginsData       `json:"pluginsData"`
	Screen            ProfileFingerprintFingerprintScreen            `json:"screen"`
	Slim              *bool                                          `json:"slim,omitempty"`
	VideoCard         ProfileFingerprintFingerprintVideoCard         `json:"videoCard"`
	VideoCodecs       map[string]string                              `json:"videoCodecs"`
}

type ProfileFingerprintHeaders struct {
	Accept                  *string `json:"accept,omitempty"`
	AcceptEncoding          *string `json:"accept-encoding,omitempty"`
	AcceptLanguage          *string `json:"accept-language,omitempty"`
	Dnt                     *string `json:"dnt,omitempty"`
	SecChUa                 *string `json:"sec-ch-ua,omitempty"`
	SecChUaMobile           *string `json:"sec-ch-ua-mobile,omitempty"`
	SecChUaPlatform         *string `json:"sec-ch-ua-platform,omitempty"`
	SecFetchDest            *string `json:"sec-fetch-dest,omitempty"`
	SecFetchMode            *string `json:"sec-fetch-mode,omitempty"`
	SecFetchSite            *string `json:"sec-fetch-site,omitempty"`
	SecFetchUser            *string `json:"sec-fetch-user,omitempty"`
	UpgradeInsecureRequests *string `json:"upgrade-insecure-requests,omitempty"`
	UserAgent               *string `json:"user-agent,omitempty"`
}

// ProfileFingerprint The fingerprint associated with the profile
type ProfileFingerprint struct {
	Fingerprint ProfileFingerprintFingerprint `json:"fingerprint"`
	Headers     ProfileFingerprintHeaders     `json:"headers"`
}

// Profile2 The list of profiles
type Profile2 struct {
	// The date and time when the profile was created
	CreatedAt time.Time `json:"createdAt"`
	// The credentials configuration associated with the profile
	CredentialsConfig interface{} `json:"credentialsConfig"`
	// The dimensions associated with the profile
	Dimensions *ProfileDimensions `json:"dimensions,omitempty"`
	// The extension IDs associated with the profile
	ExtensionIDs *[]string `json:"extensionIds,omitempty"`
	// The fingerprint associated with the profile
	Fingerprint *ProfileFingerprint `json:"fingerprint,omitempty"`
	// The unique identifier for the profile
	ID string `json:"id"`
	// The project ID associated with the profile
	ProjectID *string `json:"projectId,omitempty"`
	// The last session ID associated with the profile
	SourceSessionID *string `json:"sourceSessionId,omitempty"`
	// The status of the profile
	Status ProfileListResponseProfilesItemStatus `json:"status"`
	// The date and time when the profile was last updated
	UpdatedAt time.Time `json:"updatedAt"`
	// The proxy configuration associated with the profile
	UseProxyConfig interface{} `json:"useProxyConfig"`
	// The user agent associated with the profile
	UserAgent *string `json:"userAgent,omitempty"`
}

type Content struct {
	// Cleaned HTML content of the webpage
	CleanedHTML *string `json:"cleaned_html,omitempty"`
	// Raw HTML content of the webpage
	HTML *string `json:"html,omitempty"`
	// Webpage content converted to Markdown
	Markdown *string `json:"markdown,omitempty"`
	// Webpage content in Readability format
	Readability *map[string]interface{} `json:"readability,omitempty"`
}

type Link struct {
	// Text content of the link
	Text string `json:"text"`
	// URL of the link
	URL string `json:"url"`
}

type Metadata struct {
	// Author of the article content
	ArticleAuthor *string `json:"articleAuthor,omitempty"`
	// Author of the webpage content
	Author *string `json:"author,omitempty"`
	// Canonical URL of the webpage
	Canonical *string `json:"canonical,omitempty"`
	// Description of the webpage
	Description *string `json:"description,omitempty"`
	// Favicon URL of the website
	Favicon *string `json:"favicon,omitempty"`
	// JSON-LD structured data from the webpage
	JSONLd *interface{} `json:"jsonLd,omitempty"`
	// Keywords associated with the webpage
	Keywords *string `json:"keywords,omitempty"`
	// Detected language of the webpage
	Language *string `json:"language,omitempty"`
	// Last modification time of the content
	ModifiedTime *string `json:"modifiedTime,omitempty"`
	// Open Graph description
	OgDescription *string `json:"ogDescription,omitempty"`
	// Open Graph image URL
	OgImage *string `json:"ogImage,omitempty"`
	// Open Graph site name
	OgSiteName *string `json:"ogSiteName,omitempty"`
	// Open Graph title
	OgTitle *string `json:"ogTitle,omitempty"`
	// Open Graph URL
	OgURL *string `json:"ogUrl,omitempty"`
	// Publication time of the content
	PublishedTime *string `json:"publishedTime,omitempty"`
	// HTTP status code of the response
	StatusCode int64 `json:"statusCode"`
	// Timestamp when the scrape was performed
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Title of the webpage
	Title *string `json:"title,omitempty"`
	// Source URL of the scraped page
	URLSource *string `json:"urlSource,omitempty"`
}

type Pdf struct {
	// URL of the generated PDF
	URL string `json:"url"`
}

type Screenshot struct {
	// URL of the screenshot image
	URL string `json:"url"`
}

// CookiePartitionKey The partition key of the cookie
type CookiePartitionKey struct {
	// Indicates if the cookie has any ancestors that are cross-site to the topLevelSite.
	HasCrossSiteAncestor bool `json:"hasCrossSiteAncestor"`
	// The site of the top-level URL the browser was visiting at the start of the request to the endpoint that set the cookie.
	TopLevelSite string `json:"topLevelSite"`
}

// Cookie Cookies to initialize in the session
type Cookie struct {
	// The domain of the cookie
	Domain *string `json:"domain,omitempty"`
	// The expiration date of the cookie
	Expires *float64 `json:"expires,omitempty"`
	// Whether the cookie is HTTP only
	HTTPOnly *bool `json:"httpOnly,omitempty"`
	// The name of the cookie
	Name string `json:"name"`
	// The partition key of the cookie
	PartitionKey *CookiePartitionKey `json:"partitionKey,omitempty"`
	// The path of the cookie
	Path *string `json:"path,omitempty"`
	// The priority of the cookie
	Priority *SessionContextResponseCookiesItemPriority `json:"priority,omitempty"`
	// Whether the cookie is a same party cookie
	SameParty *bool `json:"sameParty,omitempty"`
	// The same site attribute of the cookie
	SameSite *SessionContextResponseCookiesItemSameSite `json:"sameSite,omitempty"`
	// Whether the cookie is secure
	Secure *bool `json:"secure,omitempty"`
	// Whether the cookie is a session cookie
	Session *bool `json:"session,omitempty"`
	// The size of the cookie
	Size *float64 `json:"size,omitempty"`
	// The source port of the cookie
	SourcePort *float64 `json:"sourcePort,omitempty"`
	// The source scheme of the cookie
	SourceScheme *SessionContextResponseCookiesItemSourceScheme `json:"sourceScheme,omitempty"`
	// The URL of the cookie
	URL *string `json:"url,omitempty"`
	// The value of the cookie
	Value string `json:"value"`
}

type IndexedDBDataRecordBlobFile struct {
	BlobNumber   float64    `json:"blobNumber"`
	Filename     *string    `json:"filename,omitempty"`
	LastModified *time.Time `json:"lastModified,omitempty"`
	MimeType     string     `json:"mimeType"`
	Path         *string    `json:"path,omitempty"`
	Size         float64    `json:"size"`
}

type IndexedDBDataRecord struct {
	BlobFiles *[]IndexedDBDataRecordBlobFile `json:"blobFiles,omitempty"`
	Key       interface{}                    `json:"key"`
	Value     interface{}                    `json:"value"`
}

type IndexedDBData struct {
	ID      float64               `json:"id"`
	Name    string                `json:"name"`
	Records []IndexedDBDataRecord `json:"records"`
}

// IndexedDB Domain-specific indexedDB items to initialize in the session
type IndexedDB struct {
	Data []IndexedDBData `json:"data"`
	ID   float64         `json:"id"`
	Name string          `json:"name"`
}

type UsageBrowser struct {
	// Billing unit for browser usage
	Unit SessionCostResponseUsageBrowserUnit `json:"unit"`
	// Billable browser usage value
	Value float64 `json:"value"`
}

type UsageCaptcha struct {
	// Billable captcha solves
	Solves int64 `json:"solves"`
}

type UsageProxy struct {
	// Billable Steel proxy bytes for the session
	Bytes int64 `json:"bytes"`
}

// Usage Billable usage inputs used to calculate the costs
type Usage struct {
	Browser UsageBrowser `json:"browser"`
	Captcha UsageCaptcha `json:"captcha"`
	Proxy   UsageProxy   `json:"proxy"`
}

type Page struct {
	Favicon                    *string `json:"favicon,omitempty"`
	ID                         string  `json:"id"`
	SessionViewerFullscreenURL string  `json:"sessionViewerFullscreenUrl"`
	SessionViewerURL           string  `json:"sessionViewerUrl"`
	Title                      string  `json:"title"`
	URL                        string  `json:"url"`
}

// DebugConfig2 Configuration for the debug URL and session viewer. Controls interaction capabilities and cursor visibility.
type DebugConfig2 struct {
	// Whether interaction is allowed via the debug URL viewer. When false, the session viewer is view-only.
	Interactive *bool `json:"interactive,omitempty"`
	// Whether the OS-level mouse cursor is shown in the WebRTC stream (headful mode only).
	SystemCursor *bool `json:"systemCursor,omitempty"`
}

// DeviceConfig2 Device configuration for the session
type DeviceConfig2 struct {
	Device *SessionResponseDeviceConfigDevice `json:"device,omitempty"`
}

// Dimensions4 Viewport and browser window dimensions for the session
type Dimensions4 struct {
	// Height of the browser window
	Height int64 `json:"height"`
	// Width of the browser window
	Width int64 `json:"width"`
}

// OptimizeBandwidth Bandwidth optimizations that were applied to the session.
type OptimizeBandwidth struct {
	BlockHosts       *[]string `json:"blockHosts,omitempty"`
	BlockImages      *bool     `json:"blockImages,omitempty"`
	BlockMedia       *bool     `json:"blockMedia,omitempty"`
	BlockStylesheets *bool     `json:"blockStylesheets,omitempty"`
	BlockURLPatterns *[]string `json:"blockUrlPatterns,omitempty"`
}

// StealthConfig2 Stealth configuration for the session
type StealthConfig2 struct {
	// When true, captchas will be automatically solved when detected. When false, use the solve endpoints to manually initiate solving.
	AutoCaptchaSolving *bool `json:"autoCaptchaSolving,omitempty"`
	// This flag will make the browser act more human-like by moving the mouse in a more natural way
	HumanizeInteractions *bool `json:"humanizeInteractions,omitempty"`
	// This flag will skip the fingerprint generation for the session.
	SkipFingerprintInjection *bool `json:"skipFingerprintInjection,omitempty"`
}

// SessionDebugConfig Configuration for the debug URL and session viewer. Controls interaction capabilities and cursor visibility.
type SessionDebugConfig struct {
	// Whether interaction is allowed via the debug URL viewer. When false, the session viewer is view-only.
	Interactive *bool `json:"interactive,omitempty"`
	// Whether the OS-level mouse cursor is shown in the WebRTC stream (headful mode only).
	SystemCursor *bool `json:"systemCursor,omitempty"`
}

// SessionDeviceConfig Device configuration for the session
type SessionDeviceConfig struct {
	Device *SessionsResponseSessionsItemDeviceConfigDevice `json:"device,omitempty"`
}

// SessionDimensions Viewport and browser window dimensions for the session
type SessionDimensions struct {
	// Height of the browser window
	Height int64 `json:"height"`
	// Width of the browser window
	Width int64 `json:"width"`
}

// SessionOptimizeBandwidth Bandwidth optimizations that were applied to the session.
type SessionOptimizeBandwidth struct {
	BlockHosts       *[]string `json:"blockHosts,omitempty"`
	BlockImages      *bool     `json:"blockImages,omitempty"`
	BlockMedia       *bool     `json:"blockMedia,omitempty"`
	BlockStylesheets *bool     `json:"blockStylesheets,omitempty"`
	BlockURLPatterns *[]string `json:"blockUrlPatterns,omitempty"`
}

// SessionStealthConfig Stealth configuration for the session
type SessionStealthConfig struct {
	// When true, captchas will be automatically solved when detected. When false, use the solve endpoints to manually initiate solving.
	AutoCaptchaSolving *bool `json:"autoCaptchaSolving,omitempty"`
	// This flag will make the browser act more human-like by moving the mouse in a more natural way
	HumanizeInteractions *bool `json:"humanizeInteractions,omitempty"`
	// This flag will skip the fingerprint generation for the session.
	SkipFingerprintInjection *bool `json:"skipFingerprintInjection,omitempty"`
}

// Session2 List of browser sessions
type Session2 struct {
	// Timestamp when the session started
	CreatedAt time.Time `json:"createdAt"`
	// Amount of credits consumed by the session
	CreditsUsed int64 `json:"creditsUsed"`
	// Configuration for the debug URL and session viewer. Controls interaction capabilities and cursor visibility.
	DebugConfig *SessionDebugConfig `json:"debugConfig,omitempty"`
	// URL for debugging the session
	DebugURL string `json:"debugUrl"`
	// Device configuration for the session
	DeviceConfig *SessionDeviceConfig `json:"deviceConfig,omitempty"`
	// Viewport and browser window dimensions for the session
	Dimensions SessionDimensions `json:"dimensions"`
	// Duration of the session in milliseconds
	Duration int64 `json:"duration"`
	// Number of events processed in the session
	EventCount int64 `json:"eventCount"`
	// Launch the browser in fullscreen mode, covering the full screen with no Chrome UI.
	Fullscreen *bool `json:"fullscreen,omitempty"`
	// Indicates if the session is headless or headful
	Headless *bool `json:"headless,omitempty"`
	// Unique identifier for the session
	ID string `json:"id"`
	// Inactivity timeout in milliseconds, if one was set when the session was created
	InactivityTimeout *int64 `json:"inactivityTimeout,omitempty"`
	// Indicates if Selenium is used in the session
	IsSelenium *bool `json:"isSelenium,omitempty"`
	// Bandwidth optimizations that were applied to the session.
	OptimizeBandwidth SessionOptimizeBandwidth `json:"optimizeBandwidth"`
	// This flag will persist the profile for the session.
	PersistProfile *bool `json:"persistProfile,omitempty"`
	// The ID of the profile associated with the session
	ProfileID *string `json:"profileId,omitempty"`
	// The project associated with the session
	ProjectID *string `json:"projectId,omitempty"`
	// Amount of data transmitted through the proxy
	ProxyBytesUsed int64 `json:"proxyBytesUsed"`
	// Source of the proxy used for the session
	ProxySource *SessionsResponseSessionsItemProxySource `json:"proxySource,omitempty"`
	// The region where the session was created.
	Region *SessionsResponseSessionsItemRegion `json:"region,omitempty"`
	// Why the session reached a terminal state. Null while the session is live, or when the reason is unknown (e.g. sessions created before this was tracked). One of: user_requested (released via the API/SDK), timeout (hard `timeout` elapsed), inactivity_timeout (no activity for the configured window), creation_timeout (never started in time), startup_failed (could not be dispatched), browser_closed (the browser or agent closed itself — not a crash), browser_crashed (the browser crashed or its machine became unresponsive).
	ReleaseReason *SessionsResponseSessionsItemReleaseReason `json:"releaseReason,omitempty"`
	// URL to view session details
	SessionViewerURL string `json:"sessionViewerUrl"`
	// Indicates if captcha solving is enabled
	SolveCaptcha *bool `json:"solveCaptcha,omitempty"`
	// Status of the session
	Status SessionsResponseSessionsItemStatus `json:"status"`
	// Stealth configuration for the session
	StealthConfig *SessionStealthConfig `json:"stealthConfig,omitempty"`
	// Session timeout duration in milliseconds
	Timeout int64 `json:"timeout"`
	// User agent string used in the session
	UserAgent *string `json:"userAgent,omitempty"`
	// URL for the session's WebSocket connection
	WebsocketURL string `json:"websocketUrl"`
}

type CredentialListParams struct {
	ProjectID *string `json:"projectId,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Origin    *string `json:"origin,omitempty"`
}

type ProfileListParams struct {
	ProjectID *string `json:"projectId,omitempty"`
}

type ProfileGetParams struct {
	ProjectID *string `json:"projectId,omitempty"`
}

type ProfileUpdateQueryParams struct {
	ProjectID *string `json:"projectId,omitempty"`
}

type SessionListParams struct {
	CursorID  *string            `json:"cursorId,omitempty"`
	Limit     *int64             `json:"limit,omitempty"`
	Status    *GetSessionsStatus `json:"status,omitempty"`
	ProjectID *string            `json:"projectId,omitempty"`
}

type SessionReleaseAllQueryParams struct {
	ProjectID *string `json:"projectId,omitempty"`
}

type SessionAgentLogsParams struct {
	Namespace  *string    `json:"namespace,omitempty"`
	StartTime  *time.Time `json:"startTime,omitempty"`
	EndTime    *time.Time `json:"endTime,omitempty"`
	EventTypes *[]string  `json:"eventTypes,omitempty"`
	Limit      *int64     `json:"limit,omitempty"`
	Offset     *int64     `json:"offset,omitempty"`
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

type SessionLogsParams struct {
	Namespace  *string    `json:"namespace,omitempty"`
	StartTime  *time.Time `json:"startTime,omitempty"`
	EndTime    *time.Time `json:"endTime,omitempty"`
	EventTypes *[]string  `json:"eventTypes,omitempty"`
	Limit      *int64     `json:"limit,omitempty"`
	Offset     *int64     `json:"offset,omitempty"`
}

type CaptchaStatusResponse []map[string]interface{}

type SessionComputerParams interface{}

// SessionEventsResponse Events for a browser session
type SessionEventsResponse []interface{}

type SessionReleaseAllParams map[string]interface{}

type SessionCreateParamsOptimizeBandwidth struct {
	OfBool *bool                   `json:"-"`
	OfMap  *map[string]interface{} `json:"-"`
}

func (u SessionCreateParamsOptimizeBandwidth) MarshalJSON() ([]byte, error) {
	if u.OfBool != nil {
		return json.Marshal(u.OfBool)
	}
	if u.OfMap != nil {
		return json.Marshal(u.OfMap)
	}
	return []byte("null"), nil
}

func (u *SessionCreateParamsOptimizeBandwidth) UnmarshalJSON(data []byte) error {
	var tok byte
	for _, b := range data {
		if b == ' ' || b == '\t' || b == '\n' || b == '\r' {
			continue
		}
		tok = b
		break
	}
	switch {
	case tok == 't' || tok == 'f':
		u.OfBool = new(bool)
		return json.Unmarshal(data, u.OfBool)
	case tok == '{':
		u.OfMap = new(map[string]interface{})
		return json.Unmarshal(data, u.OfMap)
	}
	return nil
}
