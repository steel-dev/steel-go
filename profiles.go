package steel

import (
	"context"
	"fmt"
	"github.com/steel-dev/steel-go/internal/apijson"
	"net/url"
	"time"
)

type ProfileService struct {
	client *Client
}

// List all profiles
func (r *ProfileService) List(ctx context.Context, params ProfileListParams, opts ...RequestOption) (*ProfileListResponse, error) {
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
func (r *ProfileService) Create(ctx context.Context, body ProfileCreateParams, opts ...RequestOption) (*ProfileCreateResponse, error) {
	path := "/v1/profiles"
	var out ProfileCreateResponse
	if err := r.client.doMultipart(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Get a profile
func (r *ProfileService) Get(ctx context.Context, id string, params ProfileGetParams, opts ...RequestOption) (*ProfileGetResponse, error) {
	path := fmt.Sprintf("/v1/profiles/%s", url.PathEscape(id))
	query := url.Values{}
	if params.ProjectID != nil {
		query.Set("projectId", *params.ProjectID)
	}
	if encoded := query.Encode(); encoded != "" {
		path += "?" + encoded
	}
	var out ProfileGetResponse
	if err := r.client.do(ctx, "GET", path, nil, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Update a profile
func (r *ProfileService) Update(ctx context.Context, id string, body ProfileUpdateParams, params ProfileUpdateQueryParams, opts ...RequestOption) (*ProfileUpdateResponse, error) {
	path := fmt.Sprintf("/v1/profiles/%s", url.PathEscape(id))
	query := url.Values{}
	if params.ProjectID != nil {
		query.Set("projectId", *params.ProjectID)
	}
	if encoded := query.Encode(); encoded != "" {
		path += "?" + encoded
	}
	var out ProfileUpdateResponse
	if err := r.client.doMultipart(ctx, "PATCH", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

type ProfileCreateResponse struct {
	// The date and time when the profile was created
	CreatedAt time.Time `json:"createdAt" api:"required"`
	// The credentials configuration associated with the profile
	CredentialsConfig interface{} `json:"credentialsConfig" api:"required"`
	// The dimensions associated with the profile
	Dimensions ProfileCreateResponseDimensions `json:"dimensions" api:"required"`
	// The extension IDs associated with the profile
	ExtensionIDs []string `json:"extensionIds" api:"required"`
	// The fingerprint associated with the profile
	Fingerprint ProfileCreateResponseFingerprint `json:"fingerprint" api:"required"`
	// The unique identifier for the profile
	ID string `json:"id" api:"required"`
	// The project ID associated with the profile
	ProjectID string `json:"projectId" api:"required"`
	// The last session ID associated with the profile
	SourceSessionID string `json:"sourceSessionId" api:"required"`
	// The status of the profile
	Status ProfileStatus `json:"status" api:"required"`
	// The date and time when the profile was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required"`
	// The proxy configuration associated with the profile
	UseProxyConfig interface{} `json:"useProxyConfig" api:"required"`
	// The user agent associated with the profile
	UserAgent string                    `json:"userAgent" api:"required"`
	JSON      profileCreateResponseJSON `json:"-"`
}

// profileCreateResponseJSON contains the JSON metadata for the struct [ProfileCreateResponse]
type profileCreateResponseJSON struct {
	CreatedAt         apijson.Field
	CredentialsConfig apijson.Field
	Dimensions        apijson.Field
	ExtensionIDs      apijson.Field
	Fingerprint       apijson.Field
	ID                apijson.Field
	ProjectID         apijson.Field
	SourceSessionID   apijson.Field
	Status            apijson.Field
	UpdatedAt         apijson.Field
	UseProxyConfig    apijson.Field
	UserAgent         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProfileCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateParams struct {
	// The dimensions associated with the profile
	Dimensions *ProfileCreateParamsDimensions `json:"dimensions,omitempty"`
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
	Count float64 `json:"count" api:"required"`
	// The list of profiles
	Profiles []ProfileListResponseProfile `json:"profiles" api:"required"`
	JSON     profileListResponseJSON      `json:"-"`
}

// profileListResponseJSON contains the JSON metadata for the struct [ProfileListResponse]
type profileListResponseJSON struct {
	Count       apijson.Field
	Profiles    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponse struct {
	// The date and time when the profile was created
	CreatedAt time.Time `json:"createdAt" api:"required"`
	// The credentials configuration associated with the profile
	CredentialsConfig interface{} `json:"credentialsConfig" api:"required"`
	// The dimensions associated with the profile
	Dimensions ProfileGetResponseDimensions `json:"dimensions" api:"required"`
	// The extension IDs associated with the profile
	ExtensionIDs []string `json:"extensionIds" api:"required"`
	// The fingerprint associated with the profile
	Fingerprint ProfileGetResponseFingerprint `json:"fingerprint" api:"required"`
	// The unique identifier for the profile
	ID string `json:"id" api:"required"`
	// The project ID associated with the profile
	ProjectID string `json:"projectId" api:"required"`
	// The last session ID associated with the profile
	SourceSessionID string `json:"sourceSessionId" api:"required"`
	// The status of the profile
	Status ProfileStatus `json:"status" api:"required"`
	// The date and time when the profile was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required"`
	// The proxy configuration associated with the profile
	UseProxyConfig interface{} `json:"useProxyConfig" api:"required"`
	// The user agent associated with the profile
	UserAgent string                 `json:"userAgent" api:"required"`
	JSON      profileGetResponseJSON `json:"-"`
}

// profileGetResponseJSON contains the JSON metadata for the struct [ProfileGetResponse]
type profileGetResponseJSON struct {
	CreatedAt         apijson.Field
	CredentialsConfig apijson.Field
	Dimensions        apijson.Field
	ExtensionIDs      apijson.Field
	Fingerprint       apijson.Field
	ID                apijson.Field
	ProjectID         apijson.Field
	SourceSessionID   apijson.Field
	Status            apijson.Field
	UpdatedAt         apijson.Field
	UseProxyConfig    apijson.Field
	UserAgent         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProfileGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateParams struct {
	// The dimensions associated with the profile
	Dimensions *ProfileUpdateParamsDimensions `json:"dimensions,omitempty"`
	// Project to create the profile in
	ProjectID *string `json:"projectId,omitempty"`
	// The proxy associated with the profile
	ProxyURL *string `json:"proxyUrl,omitempty"`
	// The user agent associated with the profile
	UserAgent *string `json:"userAgent,omitempty"`
	// The user data directory associated with the profile
	UserDataDir FileUpload `json:"userDataDir"`
}

type ProfileUpdateResponse struct {
	// The date and time when the profile was created
	CreatedAt time.Time `json:"createdAt" api:"required"`
	// The credentials configuration associated with the profile
	CredentialsConfig interface{} `json:"credentialsConfig" api:"required"`
	// The dimensions associated with the profile
	Dimensions ProfileUpdateResponseDimensions `json:"dimensions" api:"required"`
	// The extension IDs associated with the profile
	ExtensionIDs []string `json:"extensionIds" api:"required"`
	// The fingerprint associated with the profile
	Fingerprint ProfileUpdateResponseFingerprint `json:"fingerprint" api:"required"`
	// The unique identifier for the profile
	ID string `json:"id" api:"required"`
	// The project ID associated with the profile
	ProjectID string `json:"projectId" api:"required"`
	// The last session ID associated with the profile
	SourceSessionID string `json:"sourceSessionId" api:"required"`
	// The status of the profile
	Status ProfileStatus `json:"status" api:"required"`
	// The date and time when the profile was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required"`
	// The proxy configuration associated with the profile
	UseProxyConfig interface{} `json:"useProxyConfig" api:"required"`
	// The user agent associated with the profile
	UserAgent string                    `json:"userAgent" api:"required"`
	JSON      profileUpdateResponseJSON `json:"-"`
}

// profileUpdateResponseJSON contains the JSON metadata for the struct [ProfileUpdateResponse]
type profileUpdateResponseJSON struct {
	CreatedAt         apijson.Field
	CredentialsConfig apijson.Field
	Dimensions        apijson.Field
	ExtensionIDs      apijson.Field
	Fingerprint       apijson.Field
	ID                apijson.Field
	ProjectID         apijson.Field
	SourceSessionID   apijson.Field
	Status            apijson.Field
	UpdatedAt         apijson.Field
	UseProxyConfig    apijson.Field
	UserAgent         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProfileUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// ProfileCreateResponseDimensions The dimensions associated with the profile
type ProfileCreateResponseDimensions struct {
	Height float64                             `json:"height" api:"required"`
	Width  float64                             `json:"width" api:"required"`
	JSON   profileCreateResponseDimensionsJSON `json:"-"`
}

// profileCreateResponseDimensionsJSON contains the JSON metadata for the struct [ProfileCreateResponseDimensions]
type profileCreateResponseDimensionsJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCreateResponseDimensions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseDimensionsJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateResponseFingerprintFingerprintBattery struct {
	Charging        bool                                                   `json:"charging" api:"required"`
	ChargingTime    float64                                                `json:"chargingTime" api:"required"`
	DischargingTime float64                                                `json:"dischargingTime" api:"required"`
	Level           float64                                                `json:"level" api:"required"`
	JSON            profileCreateResponseFingerprintFingerprintBatteryJSON `json:"-"`
}

// profileCreateResponseFingerprintFingerprintBatteryJSON contains the JSON metadata for the struct [ProfileCreateResponseFingerprintFingerprintBattery]
type profileCreateResponseFingerprintFingerprintBatteryJSON struct {
	Charging        apijson.Field
	ChargingTime    apijson.Field
	DischargingTime apijson.Field
	Level           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProfileCreateResponseFingerprintFingerprintBattery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseFingerprintFingerprintBatteryJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateResponseFingerprintFingerprintMultimediaDevicesMicro struct {
	DeviceID string                                                                `json:"deviceId" api:"required"`
	GroupID  string                                                                `json:"groupId" api:"required"`
	Kind     string                                                                `json:"kind" api:"required"`
	Label    string                                                                `json:"label" api:"required"`
	JSON     profileCreateResponseFingerprintFingerprintMultimediaDevicesMicroJSON `json:"-"`
}

// profileCreateResponseFingerprintFingerprintMultimediaDevicesMicroJSON contains the JSON metadata for the struct [ProfileCreateResponseFingerprintFingerprintMultimediaDevicesMicro]
type profileCreateResponseFingerprintFingerprintMultimediaDevicesMicroJSON struct {
	DeviceID    apijson.Field
	GroupID     apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCreateResponseFingerprintFingerprintMultimediaDevicesMicro) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseFingerprintFingerprintMultimediaDevicesMicroJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateResponseFingerprintFingerprintMultimediaDevicesSpeaker struct {
	DeviceID string                                                                  `json:"deviceId" api:"required"`
	GroupID  string                                                                  `json:"groupId" api:"required"`
	Kind     string                                                                  `json:"kind" api:"required"`
	Label    string                                                                  `json:"label" api:"required"`
	JSON     profileCreateResponseFingerprintFingerprintMultimediaDevicesSpeakerJSON `json:"-"`
}

// profileCreateResponseFingerprintFingerprintMultimediaDevicesSpeakerJSON contains the JSON metadata for the struct [ProfileCreateResponseFingerprintFingerprintMultimediaDevicesSpeaker]
type profileCreateResponseFingerprintFingerprintMultimediaDevicesSpeakerJSON struct {
	DeviceID    apijson.Field
	GroupID     apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCreateResponseFingerprintFingerprintMultimediaDevicesSpeaker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseFingerprintFingerprintMultimediaDevicesSpeakerJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateResponseFingerprintFingerprintMultimediaDevicesWebcam struct {
	DeviceID string                                                                 `json:"deviceId" api:"required"`
	GroupID  string                                                                 `json:"groupId" api:"required"`
	Kind     string                                                                 `json:"kind" api:"required"`
	Label    string                                                                 `json:"label" api:"required"`
	JSON     profileCreateResponseFingerprintFingerprintMultimediaDevicesWebcamJSON `json:"-"`
}

// profileCreateResponseFingerprintFingerprintMultimediaDevicesWebcamJSON contains the JSON metadata for the struct [ProfileCreateResponseFingerprintFingerprintMultimediaDevicesWebcam]
type profileCreateResponseFingerprintFingerprintMultimediaDevicesWebcamJSON struct {
	DeviceID    apijson.Field
	GroupID     apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCreateResponseFingerprintFingerprintMultimediaDevicesWebcam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseFingerprintFingerprintMultimediaDevicesWebcamJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateResponseFingerprintFingerprintMultimediaDevices struct {
	Micros   []ProfileCreateResponseFingerprintFingerprintMultimediaDevicesMicro   `json:"micros" api:"required"`
	Speakers []ProfileCreateResponseFingerprintFingerprintMultimediaDevicesSpeaker `json:"speakers" api:"required"`
	Webcams  []ProfileCreateResponseFingerprintFingerprintMultimediaDevicesWebcam  `json:"webcams" api:"required"`
	JSON     profileCreateResponseFingerprintFingerprintMultimediaDevicesJSON      `json:"-"`
}

// profileCreateResponseFingerprintFingerprintMultimediaDevicesJSON contains the JSON metadata for the struct [ProfileCreateResponseFingerprintFingerprintMultimediaDevices]
type profileCreateResponseFingerprintFingerprintMultimediaDevicesJSON struct {
	Micros      apijson.Field
	Speakers    apijson.Field
	Webcams     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCreateResponseFingerprintFingerprintMultimediaDevices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseFingerprintFingerprintMultimediaDevicesJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateResponseFingerprintFingerprintNavigatorExtraProperties struct {
	GlobalPrivacyControl bool                                                                    `json:"globalPrivacyControl" api:"required"`
	InstalledApps        []string                                                                `json:"installedApps" api:"required"`
	PdfViewerEnabled     bool                                                                    `json:"pdfViewerEnabled" api:"required"`
	VendorFlavors        []string                                                                `json:"vendorFlavors" api:"required"`
	JSON                 profileCreateResponseFingerprintFingerprintNavigatorExtraPropertiesJSON `json:"-"`
}

// profileCreateResponseFingerprintFingerprintNavigatorExtraPropertiesJSON contains the JSON metadata for the struct [ProfileCreateResponseFingerprintFingerprintNavigatorExtraProperties]
type profileCreateResponseFingerprintFingerprintNavigatorExtraPropertiesJSON struct {
	GlobalPrivacyControl apijson.Field
	InstalledApps        apijson.Field
	PdfViewerEnabled     apijson.Field
	VendorFlavors        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProfileCreateResponseFingerprintFingerprintNavigatorExtraProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseFingerprintFingerprintNavigatorExtraPropertiesJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateResponseFingerprintFingerprintNavigatorUserAgentDataBrand struct {
	Brand   string                                                                     `json:"brand" api:"required"`
	Version string                                                                     `json:"version" api:"required"`
	JSON    profileCreateResponseFingerprintFingerprintNavigatorUserAgentDataBrandJSON `json:"-"`
}

// profileCreateResponseFingerprintFingerprintNavigatorUserAgentDataBrandJSON contains the JSON metadata for the struct [ProfileCreateResponseFingerprintFingerprintNavigatorUserAgentDataBrand]
type profileCreateResponseFingerprintFingerprintNavigatorUserAgentDataBrandJSON struct {
	Brand       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCreateResponseFingerprintFingerprintNavigatorUserAgentDataBrand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseFingerprintFingerprintNavigatorUserAgentDataBrandJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateResponseFingerprintFingerprintNavigatorUserAgentData struct {
	Brands   []ProfileCreateResponseFingerprintFingerprintNavigatorUserAgentDataBrand `json:"brands" api:"required"`
	Mobile   bool                                                                     `json:"mobile" api:"required"`
	Platform string                                                                   `json:"platform" api:"required"`
	JSON     profileCreateResponseFingerprintFingerprintNavigatorUserAgentDataJSON    `json:"-"`
}

// profileCreateResponseFingerprintFingerprintNavigatorUserAgentDataJSON contains the JSON metadata for the struct [ProfileCreateResponseFingerprintFingerprintNavigatorUserAgentData]
type profileCreateResponseFingerprintFingerprintNavigatorUserAgentDataJSON struct {
	Brands      apijson.Field
	Mobile      apijson.Field
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCreateResponseFingerprintFingerprintNavigatorUserAgentData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseFingerprintFingerprintNavigatorUserAgentDataJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateResponseFingerprintFingerprintNavigator struct {
	AppCodeName         string                                                              `json:"appCodeName" api:"required"`
	AppName             string                                                              `json:"appName" api:"required"`
	AppVersion          string                                                              `json:"appVersion" api:"required"`
	DeviceMemory        float64                                                             `json:"deviceMemory" api:"required"`
	DoNotTrack          string                                                              `json:"doNotTrack"`
	ExtraProperties     ProfileCreateResponseFingerprintFingerprintNavigatorExtraProperties `json:"extraProperties" api:"required"`
	HardwareConcurrency float64                                                             `json:"hardwareConcurrency" api:"required"`
	Language            string                                                              `json:"language" api:"required"`
	Languages           []string                                                            `json:"languages" api:"required"`
	MaxTouchPoints      float64                                                             `json:"maxTouchPoints" api:"required"`
	Oscpu               string                                                              `json:"oscpu" api:"required"`
	Platform            string                                                              `json:"platform" api:"required"`
	Product             string                                                              `json:"product" api:"required"`
	ProductSub          string                                                              `json:"productSub" api:"required"`
	UserAgent           string                                                              `json:"userAgent" api:"required"`
	UserAgentData       ProfileCreateResponseFingerprintFingerprintNavigatorUserAgentData   `json:"userAgentData" api:"required"`
	Vendor              string                                                              `json:"vendor" api:"required"`
	VendorSub           string                                                              `json:"vendorSub" api:"required"`
	Webdriver           bool                                                                `json:"webdriver" api:"required"`
	JSON                profileCreateResponseFingerprintFingerprintNavigatorJSON            `json:"-"`
}

// profileCreateResponseFingerprintFingerprintNavigatorJSON contains the JSON metadata for the struct [ProfileCreateResponseFingerprintFingerprintNavigator]
type profileCreateResponseFingerprintFingerprintNavigatorJSON struct {
	AppCodeName         apijson.Field
	AppName             apijson.Field
	AppVersion          apijson.Field
	DeviceMemory        apijson.Field
	DoNotTrack          apijson.Field
	ExtraProperties     apijson.Field
	HardwareConcurrency apijson.Field
	Language            apijson.Field
	Languages           apijson.Field
	MaxTouchPoints      apijson.Field
	Oscpu               apijson.Field
	Platform            apijson.Field
	Product             apijson.Field
	ProductSub          apijson.Field
	UserAgent           apijson.Field
	UserAgentData       apijson.Field
	Vendor              apijson.Field
	VendorSub           apijson.Field
	Webdriver           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProfileCreateResponseFingerprintFingerprintNavigator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseFingerprintFingerprintNavigatorJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateResponseFingerprintFingerprintPluginsDataPluginMimeType struct {
	Description   string                                                                   `json:"description" api:"required"`
	EnabledPlugin string                                                                   `json:"enabledPlugin" api:"required"`
	Suffixes      string                                                                   `json:"suffixes" api:"required"`
	Type          string                                                                   `json:"type" api:"required"`
	JSON          profileCreateResponseFingerprintFingerprintPluginsDataPluginMimeTypeJSON `json:"-"`
}

// profileCreateResponseFingerprintFingerprintPluginsDataPluginMimeTypeJSON contains the JSON metadata for the struct [ProfileCreateResponseFingerprintFingerprintPluginsDataPluginMimeType]
type profileCreateResponseFingerprintFingerprintPluginsDataPluginMimeTypeJSON struct {
	Description   apijson.Field
	EnabledPlugin apijson.Field
	Suffixes      apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProfileCreateResponseFingerprintFingerprintPluginsDataPluginMimeType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseFingerprintFingerprintPluginsDataPluginMimeTypeJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateResponseFingerprintFingerprintPluginsDataPlugin struct {
	Description string                                                                 `json:"description" api:"required"`
	Filename    string                                                                 `json:"filename" api:"required"`
	MimeTypes   []ProfileCreateResponseFingerprintFingerprintPluginsDataPluginMimeType `json:"mimeTypes" api:"required"`
	Name        string                                                                 `json:"name" api:"required"`
	JSON        profileCreateResponseFingerprintFingerprintPluginsDataPluginJSON       `json:"-"`
}

// profileCreateResponseFingerprintFingerprintPluginsDataPluginJSON contains the JSON metadata for the struct [ProfileCreateResponseFingerprintFingerprintPluginsDataPlugin]
type profileCreateResponseFingerprintFingerprintPluginsDataPluginJSON struct {
	Description apijson.Field
	Filename    apijson.Field
	MimeTypes   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCreateResponseFingerprintFingerprintPluginsDataPlugin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseFingerprintFingerprintPluginsDataPluginJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateResponseFingerprintFingerprintPluginsData struct {
	MimeTypes []string                                                       `json:"mimeTypes" api:"required"`
	Plugins   []ProfileCreateResponseFingerprintFingerprintPluginsDataPlugin `json:"plugins" api:"required"`
	JSON      profileCreateResponseFingerprintFingerprintPluginsDataJSON     `json:"-"`
}

// profileCreateResponseFingerprintFingerprintPluginsDataJSON contains the JSON metadata for the struct [ProfileCreateResponseFingerprintFingerprintPluginsData]
type profileCreateResponseFingerprintFingerprintPluginsDataJSON struct {
	MimeTypes   apijson.Field
	Plugins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCreateResponseFingerprintFingerprintPluginsData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseFingerprintFingerprintPluginsDataJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateResponseFingerprintFingerprintScreen struct {
	AvailHeight      float64                                               `json:"availHeight" api:"required"`
	AvailLeft        float64                                               `json:"availLeft" api:"required"`
	AvailTop         float64                                               `json:"availTop" api:"required"`
	AvailWidth       float64                                               `json:"availWidth" api:"required"`
	ClientHeight     float64                                               `json:"clientHeight" api:"required"`
	ClientWidth      float64                                               `json:"clientWidth" api:"required"`
	ColorDepth       float64                                               `json:"colorDepth" api:"required"`
	DevicePixelRatio float64                                               `json:"devicePixelRatio" api:"required"`
	HasHdr           bool                                                  `json:"hasHDR" api:"required"`
	Height           float64                                               `json:"height" api:"required"`
	InnerHeight      float64                                               `json:"innerHeight" api:"required"`
	InnerWidth       float64                                               `json:"innerWidth" api:"required"`
	OuterHeight      float64                                               `json:"outerHeight" api:"required"`
	OuterWidth       float64                                               `json:"outerWidth" api:"required"`
	PageXOffset      float64                                               `json:"pageXOffset" api:"required"`
	PageYOffset      float64                                               `json:"pageYOffset" api:"required"`
	PixelDepth       float64                                               `json:"pixelDepth" api:"required"`
	ScreenX          float64                                               `json:"screenX" api:"required"`
	Width            float64                                               `json:"width" api:"required"`
	JSON             profileCreateResponseFingerprintFingerprintScreenJSON `json:"-"`
}

// profileCreateResponseFingerprintFingerprintScreenJSON contains the JSON metadata for the struct [ProfileCreateResponseFingerprintFingerprintScreen]
type profileCreateResponseFingerprintFingerprintScreenJSON struct {
	AvailHeight      apijson.Field
	AvailLeft        apijson.Field
	AvailTop         apijson.Field
	AvailWidth       apijson.Field
	ClientHeight     apijson.Field
	ClientWidth      apijson.Field
	ColorDepth       apijson.Field
	DevicePixelRatio apijson.Field
	HasHdr           apijson.Field
	Height           apijson.Field
	InnerHeight      apijson.Field
	InnerWidth       apijson.Field
	OuterHeight      apijson.Field
	OuterWidth       apijson.Field
	PageXOffset      apijson.Field
	PageYOffset      apijson.Field
	PixelDepth       apijson.Field
	ScreenX          apijson.Field
	Width            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProfileCreateResponseFingerprintFingerprintScreen) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseFingerprintFingerprintScreenJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateResponseFingerprintFingerprintVideoCard struct {
	Renderer string                                                   `json:"renderer" api:"required"`
	Vendor   string                                                   `json:"vendor" api:"required"`
	JSON     profileCreateResponseFingerprintFingerprintVideoCardJSON `json:"-"`
}

// profileCreateResponseFingerprintFingerprintVideoCardJSON contains the JSON metadata for the struct [ProfileCreateResponseFingerprintFingerprintVideoCard]
type profileCreateResponseFingerprintFingerprintVideoCardJSON struct {
	Renderer    apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCreateResponseFingerprintFingerprintVideoCard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseFingerprintFingerprintVideoCardJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateResponseFingerprintFingerprint struct {
	AudioCodecs       map[string]string                                            `json:"audioCodecs" api:"required"`
	Battery           ProfileCreateResponseFingerprintFingerprintBattery           `json:"battery" api:"required"`
	Fonts             []string                                                     `json:"fonts" api:"required"`
	MockWebRtc        bool                                                         `json:"mockWebRTC" api:"required"`
	MultimediaDevices ProfileCreateResponseFingerprintFingerprintMultimediaDevices `json:"multimediaDevices" api:"required"`
	Navigator         ProfileCreateResponseFingerprintFingerprintNavigator         `json:"navigator" api:"required"`
	PluginsData       ProfileCreateResponseFingerprintFingerprintPluginsData       `json:"pluginsData" api:"required"`
	Screen            ProfileCreateResponseFingerprintFingerprintScreen            `json:"screen" api:"required"`
	Slim              bool                                                         `json:"slim" api:"required"`
	VideoCard         ProfileCreateResponseFingerprintFingerprintVideoCard         `json:"videoCard" api:"required"`
	VideoCodecs       map[string]string                                            `json:"videoCodecs" api:"required"`
	JSON              profileCreateResponseFingerprintFingerprintJSON              `json:"-"`
}

// profileCreateResponseFingerprintFingerprintJSON contains the JSON metadata for the struct [ProfileCreateResponseFingerprintFingerprint]
type profileCreateResponseFingerprintFingerprintJSON struct {
	AudioCodecs       apijson.Field
	Battery           apijson.Field
	Fonts             apijson.Field
	MockWebRtc        apijson.Field
	MultimediaDevices apijson.Field
	Navigator         apijson.Field
	PluginsData       apijson.Field
	Screen            apijson.Field
	Slim              apijson.Field
	VideoCard         apijson.Field
	VideoCodecs       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProfileCreateResponseFingerprintFingerprint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseFingerprintFingerprintJSON) RawJSON() string {
	return r.raw
}

type ProfileCreateResponseFingerprintHeaders struct {
	Accept                  string                                      `json:"accept"`
	AcceptEncoding          string                                      `json:"accept-encoding"`
	AcceptLanguage          string                                      `json:"accept-language"`
	Dnt                     string                                      `json:"dnt"`
	SecChUa                 string                                      `json:"sec-ch-ua"`
	SecChUaMobile           string                                      `json:"sec-ch-ua-mobile"`
	SecChUaPlatform         string                                      `json:"sec-ch-ua-platform"`
	SecFetchDest            string                                      `json:"sec-fetch-dest"`
	SecFetchMode            string                                      `json:"sec-fetch-mode"`
	SecFetchSite            string                                      `json:"sec-fetch-site"`
	SecFetchUser            string                                      `json:"sec-fetch-user"`
	UpgradeInsecureRequests string                                      `json:"upgrade-insecure-requests"`
	UserAgent               string                                      `json:"user-agent" api:"required"`
	JSON                    profileCreateResponseFingerprintHeadersJSON `json:"-"`
}

// profileCreateResponseFingerprintHeadersJSON contains the JSON metadata for the struct [ProfileCreateResponseFingerprintHeaders]
type profileCreateResponseFingerprintHeadersJSON struct {
	Accept                  apijson.Field
	AcceptEncoding          apijson.Field
	AcceptLanguage          apijson.Field
	Dnt                     apijson.Field
	SecChUa                 apijson.Field
	SecChUaMobile           apijson.Field
	SecChUaPlatform         apijson.Field
	SecFetchDest            apijson.Field
	SecFetchMode            apijson.Field
	SecFetchSite            apijson.Field
	SecFetchUser            apijson.Field
	UpgradeInsecureRequests apijson.Field
	UserAgent               apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProfileCreateResponseFingerprintHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseFingerprintHeadersJSON) RawJSON() string {
	return r.raw
}

// ProfileCreateResponseFingerprint The fingerprint associated with the profile
type ProfileCreateResponseFingerprint struct {
	Fingerprint ProfileCreateResponseFingerprintFingerprint `json:"fingerprint" api:"required"`
	Headers     ProfileCreateResponseFingerprintHeaders     `json:"headers" api:"required"`
	JSON        profileCreateResponseFingerprintJSON        `json:"-"`
}

// profileCreateResponseFingerprintJSON contains the JSON metadata for the struct [ProfileCreateResponseFingerprint]
type profileCreateResponseFingerprintJSON struct {
	Fingerprint apijson.Field
	Headers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCreateResponseFingerprint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCreateResponseFingerprintJSON) RawJSON() string {
	return r.raw
}

// ProfileCreateParamsDimensions The dimensions associated with the profile
type ProfileCreateParamsDimensions struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// ProfileListResponseProfileDimensions The dimensions associated with the profile
type ProfileListResponseProfileDimensions struct {
	Height float64                                  `json:"height" api:"required"`
	Width  float64                                  `json:"width" api:"required"`
	JSON   profileListResponseProfileDimensionsJSON `json:"-"`
}

// profileListResponseProfileDimensionsJSON contains the JSON metadata for the struct [ProfileListResponseProfileDimensions]
type profileListResponseProfileDimensionsJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileListResponseProfileDimensions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileDimensionsJSON) RawJSON() string {
	return r.raw
}

type ProfileListResponseProfileFingerprintFingerprintBattery struct {
	Charging        bool                                                        `json:"charging" api:"required"`
	ChargingTime    float64                                                     `json:"chargingTime" api:"required"`
	DischargingTime float64                                                     `json:"dischargingTime" api:"required"`
	Level           float64                                                     `json:"level" api:"required"`
	JSON            profileListResponseProfileFingerprintFingerprintBatteryJSON `json:"-"`
}

// profileListResponseProfileFingerprintFingerprintBatteryJSON contains the JSON metadata for the struct [ProfileListResponseProfileFingerprintFingerprintBattery]
type profileListResponseProfileFingerprintFingerprintBatteryJSON struct {
	Charging        apijson.Field
	ChargingTime    apijson.Field
	DischargingTime apijson.Field
	Level           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProfileListResponseProfileFingerprintFingerprintBattery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileFingerprintFingerprintBatteryJSON) RawJSON() string {
	return r.raw
}

type ProfileListResponseProfileFingerprintFingerprintMultimediaDevicesMicro struct {
	DeviceID string                                                                     `json:"deviceId" api:"required"`
	GroupID  string                                                                     `json:"groupId" api:"required"`
	Kind     string                                                                     `json:"kind" api:"required"`
	Label    string                                                                     `json:"label" api:"required"`
	JSON     profileListResponseProfileFingerprintFingerprintMultimediaDevicesMicroJSON `json:"-"`
}

// profileListResponseProfileFingerprintFingerprintMultimediaDevicesMicroJSON contains the JSON metadata for the struct [ProfileListResponseProfileFingerprintFingerprintMultimediaDevicesMicro]
type profileListResponseProfileFingerprintFingerprintMultimediaDevicesMicroJSON struct {
	DeviceID    apijson.Field
	GroupID     apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileListResponseProfileFingerprintFingerprintMultimediaDevicesMicro) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileFingerprintFingerprintMultimediaDevicesMicroJSON) RawJSON() string {
	return r.raw
}

type ProfileListResponseProfileFingerprintFingerprintMultimediaDevicesSpeaker struct {
	DeviceID string                                                                       `json:"deviceId" api:"required"`
	GroupID  string                                                                       `json:"groupId" api:"required"`
	Kind     string                                                                       `json:"kind" api:"required"`
	Label    string                                                                       `json:"label" api:"required"`
	JSON     profileListResponseProfileFingerprintFingerprintMultimediaDevicesSpeakerJSON `json:"-"`
}

// profileListResponseProfileFingerprintFingerprintMultimediaDevicesSpeakerJSON contains the JSON metadata for the struct [ProfileListResponseProfileFingerprintFingerprintMultimediaDevicesSpeaker]
type profileListResponseProfileFingerprintFingerprintMultimediaDevicesSpeakerJSON struct {
	DeviceID    apijson.Field
	GroupID     apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileListResponseProfileFingerprintFingerprintMultimediaDevicesSpeaker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileFingerprintFingerprintMultimediaDevicesSpeakerJSON) RawJSON() string {
	return r.raw
}

type ProfileListResponseProfileFingerprintFingerprintMultimediaDevicesWebcam struct {
	DeviceID string                                                                      `json:"deviceId" api:"required"`
	GroupID  string                                                                      `json:"groupId" api:"required"`
	Kind     string                                                                      `json:"kind" api:"required"`
	Label    string                                                                      `json:"label" api:"required"`
	JSON     profileListResponseProfileFingerprintFingerprintMultimediaDevicesWebcamJSON `json:"-"`
}

// profileListResponseProfileFingerprintFingerprintMultimediaDevicesWebcamJSON contains the JSON metadata for the struct [ProfileListResponseProfileFingerprintFingerprintMultimediaDevicesWebcam]
type profileListResponseProfileFingerprintFingerprintMultimediaDevicesWebcamJSON struct {
	DeviceID    apijson.Field
	GroupID     apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileListResponseProfileFingerprintFingerprintMultimediaDevicesWebcam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileFingerprintFingerprintMultimediaDevicesWebcamJSON) RawJSON() string {
	return r.raw
}

type ProfileListResponseProfileFingerprintFingerprintMultimediaDevices struct {
	Micros   []ProfileListResponseProfileFingerprintFingerprintMultimediaDevicesMicro   `json:"micros" api:"required"`
	Speakers []ProfileListResponseProfileFingerprintFingerprintMultimediaDevicesSpeaker `json:"speakers" api:"required"`
	Webcams  []ProfileListResponseProfileFingerprintFingerprintMultimediaDevicesWebcam  `json:"webcams" api:"required"`
	JSON     profileListResponseProfileFingerprintFingerprintMultimediaDevicesJSON      `json:"-"`
}

// profileListResponseProfileFingerprintFingerprintMultimediaDevicesJSON contains the JSON metadata for the struct [ProfileListResponseProfileFingerprintFingerprintMultimediaDevices]
type profileListResponseProfileFingerprintFingerprintMultimediaDevicesJSON struct {
	Micros      apijson.Field
	Speakers    apijson.Field
	Webcams     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileListResponseProfileFingerprintFingerprintMultimediaDevices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileFingerprintFingerprintMultimediaDevicesJSON) RawJSON() string {
	return r.raw
}

type ProfileListResponseProfileFingerprintFingerprintNavigatorExtraProperties struct {
	GlobalPrivacyControl bool                                                                         `json:"globalPrivacyControl" api:"required"`
	InstalledApps        []string                                                                     `json:"installedApps" api:"required"`
	PdfViewerEnabled     bool                                                                         `json:"pdfViewerEnabled" api:"required"`
	VendorFlavors        []string                                                                     `json:"vendorFlavors" api:"required"`
	JSON                 profileListResponseProfileFingerprintFingerprintNavigatorExtraPropertiesJSON `json:"-"`
}

// profileListResponseProfileFingerprintFingerprintNavigatorExtraPropertiesJSON contains the JSON metadata for the struct [ProfileListResponseProfileFingerprintFingerprintNavigatorExtraProperties]
type profileListResponseProfileFingerprintFingerprintNavigatorExtraPropertiesJSON struct {
	GlobalPrivacyControl apijson.Field
	InstalledApps        apijson.Field
	PdfViewerEnabled     apijson.Field
	VendorFlavors        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProfileListResponseProfileFingerprintFingerprintNavigatorExtraProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileFingerprintFingerprintNavigatorExtraPropertiesJSON) RawJSON() string {
	return r.raw
}

type ProfileListResponseProfileFingerprintFingerprintNavigatorUserAgentDataBrand struct {
	Brand   string                                                                          `json:"brand" api:"required"`
	Version string                                                                          `json:"version" api:"required"`
	JSON    profileListResponseProfileFingerprintFingerprintNavigatorUserAgentDataBrandJSON `json:"-"`
}

// profileListResponseProfileFingerprintFingerprintNavigatorUserAgentDataBrandJSON contains the JSON metadata for the struct [ProfileListResponseProfileFingerprintFingerprintNavigatorUserAgentDataBrand]
type profileListResponseProfileFingerprintFingerprintNavigatorUserAgentDataBrandJSON struct {
	Brand       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileListResponseProfileFingerprintFingerprintNavigatorUserAgentDataBrand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileFingerprintFingerprintNavigatorUserAgentDataBrandJSON) RawJSON() string {
	return r.raw
}

type ProfileListResponseProfileFingerprintFingerprintNavigatorUserAgentData struct {
	Brands   []ProfileListResponseProfileFingerprintFingerprintNavigatorUserAgentDataBrand `json:"brands" api:"required"`
	Mobile   bool                                                                          `json:"mobile" api:"required"`
	Platform string                                                                        `json:"platform" api:"required"`
	JSON     profileListResponseProfileFingerprintFingerprintNavigatorUserAgentDataJSON    `json:"-"`
}

// profileListResponseProfileFingerprintFingerprintNavigatorUserAgentDataJSON contains the JSON metadata for the struct [ProfileListResponseProfileFingerprintFingerprintNavigatorUserAgentData]
type profileListResponseProfileFingerprintFingerprintNavigatorUserAgentDataJSON struct {
	Brands      apijson.Field
	Mobile      apijson.Field
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileListResponseProfileFingerprintFingerprintNavigatorUserAgentData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileFingerprintFingerprintNavigatorUserAgentDataJSON) RawJSON() string {
	return r.raw
}

type ProfileListResponseProfileFingerprintFingerprintNavigator struct {
	AppCodeName         string                                                                   `json:"appCodeName" api:"required"`
	AppName             string                                                                   `json:"appName" api:"required"`
	AppVersion          string                                                                   `json:"appVersion" api:"required"`
	DeviceMemory        float64                                                                  `json:"deviceMemory" api:"required"`
	DoNotTrack          string                                                                   `json:"doNotTrack"`
	ExtraProperties     ProfileListResponseProfileFingerprintFingerprintNavigatorExtraProperties `json:"extraProperties" api:"required"`
	HardwareConcurrency float64                                                                  `json:"hardwareConcurrency" api:"required"`
	Language            string                                                                   `json:"language" api:"required"`
	Languages           []string                                                                 `json:"languages" api:"required"`
	MaxTouchPoints      float64                                                                  `json:"maxTouchPoints" api:"required"`
	Oscpu               string                                                                   `json:"oscpu" api:"required"`
	Platform            string                                                                   `json:"platform" api:"required"`
	Product             string                                                                   `json:"product" api:"required"`
	ProductSub          string                                                                   `json:"productSub" api:"required"`
	UserAgent           string                                                                   `json:"userAgent" api:"required"`
	UserAgentData       ProfileListResponseProfileFingerprintFingerprintNavigatorUserAgentData   `json:"userAgentData" api:"required"`
	Vendor              string                                                                   `json:"vendor" api:"required"`
	VendorSub           string                                                                   `json:"vendorSub" api:"required"`
	Webdriver           bool                                                                     `json:"webdriver" api:"required"`
	JSON                profileListResponseProfileFingerprintFingerprintNavigatorJSON            `json:"-"`
}

// profileListResponseProfileFingerprintFingerprintNavigatorJSON contains the JSON metadata for the struct [ProfileListResponseProfileFingerprintFingerprintNavigator]
type profileListResponseProfileFingerprintFingerprintNavigatorJSON struct {
	AppCodeName         apijson.Field
	AppName             apijson.Field
	AppVersion          apijson.Field
	DeviceMemory        apijson.Field
	DoNotTrack          apijson.Field
	ExtraProperties     apijson.Field
	HardwareConcurrency apijson.Field
	Language            apijson.Field
	Languages           apijson.Field
	MaxTouchPoints      apijson.Field
	Oscpu               apijson.Field
	Platform            apijson.Field
	Product             apijson.Field
	ProductSub          apijson.Field
	UserAgent           apijson.Field
	UserAgentData       apijson.Field
	Vendor              apijson.Field
	VendorSub           apijson.Field
	Webdriver           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProfileListResponseProfileFingerprintFingerprintNavigator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileFingerprintFingerprintNavigatorJSON) RawJSON() string {
	return r.raw
}

type ProfileListResponseProfileFingerprintFingerprintPluginsDataPluginMimeType struct {
	Description   string                                                                        `json:"description" api:"required"`
	EnabledPlugin string                                                                        `json:"enabledPlugin" api:"required"`
	Suffixes      string                                                                        `json:"suffixes" api:"required"`
	Type          string                                                                        `json:"type" api:"required"`
	JSON          profileListResponseProfileFingerprintFingerprintPluginsDataPluginMimeTypeJSON `json:"-"`
}

// profileListResponseProfileFingerprintFingerprintPluginsDataPluginMimeTypeJSON contains the JSON metadata for the struct [ProfileListResponseProfileFingerprintFingerprintPluginsDataPluginMimeType]
type profileListResponseProfileFingerprintFingerprintPluginsDataPluginMimeTypeJSON struct {
	Description   apijson.Field
	EnabledPlugin apijson.Field
	Suffixes      apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProfileListResponseProfileFingerprintFingerprintPluginsDataPluginMimeType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileFingerprintFingerprintPluginsDataPluginMimeTypeJSON) RawJSON() string {
	return r.raw
}

type ProfileListResponseProfileFingerprintFingerprintPluginsDataPlugin struct {
	Description string                                                                      `json:"description" api:"required"`
	Filename    string                                                                      `json:"filename" api:"required"`
	MimeTypes   []ProfileListResponseProfileFingerprintFingerprintPluginsDataPluginMimeType `json:"mimeTypes" api:"required"`
	Name        string                                                                      `json:"name" api:"required"`
	JSON        profileListResponseProfileFingerprintFingerprintPluginsDataPluginJSON       `json:"-"`
}

// profileListResponseProfileFingerprintFingerprintPluginsDataPluginJSON contains the JSON metadata for the struct [ProfileListResponseProfileFingerprintFingerprintPluginsDataPlugin]
type profileListResponseProfileFingerprintFingerprintPluginsDataPluginJSON struct {
	Description apijson.Field
	Filename    apijson.Field
	MimeTypes   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileListResponseProfileFingerprintFingerprintPluginsDataPlugin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileFingerprintFingerprintPluginsDataPluginJSON) RawJSON() string {
	return r.raw
}

type ProfileListResponseProfileFingerprintFingerprintPluginsData struct {
	MimeTypes []string                                                            `json:"mimeTypes" api:"required"`
	Plugins   []ProfileListResponseProfileFingerprintFingerprintPluginsDataPlugin `json:"plugins" api:"required"`
	JSON      profileListResponseProfileFingerprintFingerprintPluginsDataJSON     `json:"-"`
}

// profileListResponseProfileFingerprintFingerprintPluginsDataJSON contains the JSON metadata for the struct [ProfileListResponseProfileFingerprintFingerprintPluginsData]
type profileListResponseProfileFingerprintFingerprintPluginsDataJSON struct {
	MimeTypes   apijson.Field
	Plugins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileListResponseProfileFingerprintFingerprintPluginsData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileFingerprintFingerprintPluginsDataJSON) RawJSON() string {
	return r.raw
}

type ProfileListResponseProfileFingerprintFingerprintScreen struct {
	AvailHeight      float64                                                    `json:"availHeight" api:"required"`
	AvailLeft        float64                                                    `json:"availLeft" api:"required"`
	AvailTop         float64                                                    `json:"availTop" api:"required"`
	AvailWidth       float64                                                    `json:"availWidth" api:"required"`
	ClientHeight     float64                                                    `json:"clientHeight" api:"required"`
	ClientWidth      float64                                                    `json:"clientWidth" api:"required"`
	ColorDepth       float64                                                    `json:"colorDepth" api:"required"`
	DevicePixelRatio float64                                                    `json:"devicePixelRatio" api:"required"`
	HasHdr           bool                                                       `json:"hasHDR" api:"required"`
	Height           float64                                                    `json:"height" api:"required"`
	InnerHeight      float64                                                    `json:"innerHeight" api:"required"`
	InnerWidth       float64                                                    `json:"innerWidth" api:"required"`
	OuterHeight      float64                                                    `json:"outerHeight" api:"required"`
	OuterWidth       float64                                                    `json:"outerWidth" api:"required"`
	PageXOffset      float64                                                    `json:"pageXOffset" api:"required"`
	PageYOffset      float64                                                    `json:"pageYOffset" api:"required"`
	PixelDepth       float64                                                    `json:"pixelDepth" api:"required"`
	ScreenX          float64                                                    `json:"screenX" api:"required"`
	Width            float64                                                    `json:"width" api:"required"`
	JSON             profileListResponseProfileFingerprintFingerprintScreenJSON `json:"-"`
}

// profileListResponseProfileFingerprintFingerprintScreenJSON contains the JSON metadata for the struct [ProfileListResponseProfileFingerprintFingerprintScreen]
type profileListResponseProfileFingerprintFingerprintScreenJSON struct {
	AvailHeight      apijson.Field
	AvailLeft        apijson.Field
	AvailTop         apijson.Field
	AvailWidth       apijson.Field
	ClientHeight     apijson.Field
	ClientWidth      apijson.Field
	ColorDepth       apijson.Field
	DevicePixelRatio apijson.Field
	HasHdr           apijson.Field
	Height           apijson.Field
	InnerHeight      apijson.Field
	InnerWidth       apijson.Field
	OuterHeight      apijson.Field
	OuterWidth       apijson.Field
	PageXOffset      apijson.Field
	PageYOffset      apijson.Field
	PixelDepth       apijson.Field
	ScreenX          apijson.Field
	Width            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProfileListResponseProfileFingerprintFingerprintScreen) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileFingerprintFingerprintScreenJSON) RawJSON() string {
	return r.raw
}

type ProfileListResponseProfileFingerprintFingerprintVideoCard struct {
	Renderer string                                                        `json:"renderer" api:"required"`
	Vendor   string                                                        `json:"vendor" api:"required"`
	JSON     profileListResponseProfileFingerprintFingerprintVideoCardJSON `json:"-"`
}

// profileListResponseProfileFingerprintFingerprintVideoCardJSON contains the JSON metadata for the struct [ProfileListResponseProfileFingerprintFingerprintVideoCard]
type profileListResponseProfileFingerprintFingerprintVideoCardJSON struct {
	Renderer    apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileListResponseProfileFingerprintFingerprintVideoCard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileFingerprintFingerprintVideoCardJSON) RawJSON() string {
	return r.raw
}

type ProfileListResponseProfileFingerprintFingerprint struct {
	AudioCodecs       map[string]string                                                 `json:"audioCodecs" api:"required"`
	Battery           ProfileListResponseProfileFingerprintFingerprintBattery           `json:"battery" api:"required"`
	Fonts             []string                                                          `json:"fonts" api:"required"`
	MockWebRtc        bool                                                              `json:"mockWebRTC" api:"required"`
	MultimediaDevices ProfileListResponseProfileFingerprintFingerprintMultimediaDevices `json:"multimediaDevices" api:"required"`
	Navigator         ProfileListResponseProfileFingerprintFingerprintNavigator         `json:"navigator" api:"required"`
	PluginsData       ProfileListResponseProfileFingerprintFingerprintPluginsData       `json:"pluginsData" api:"required"`
	Screen            ProfileListResponseProfileFingerprintFingerprintScreen            `json:"screen" api:"required"`
	Slim              bool                                                              `json:"slim" api:"required"`
	VideoCard         ProfileListResponseProfileFingerprintFingerprintVideoCard         `json:"videoCard" api:"required"`
	VideoCodecs       map[string]string                                                 `json:"videoCodecs" api:"required"`
	JSON              profileListResponseProfileFingerprintFingerprintJSON              `json:"-"`
}

// profileListResponseProfileFingerprintFingerprintJSON contains the JSON metadata for the struct [ProfileListResponseProfileFingerprintFingerprint]
type profileListResponseProfileFingerprintFingerprintJSON struct {
	AudioCodecs       apijson.Field
	Battery           apijson.Field
	Fonts             apijson.Field
	MockWebRtc        apijson.Field
	MultimediaDevices apijson.Field
	Navigator         apijson.Field
	PluginsData       apijson.Field
	Screen            apijson.Field
	Slim              apijson.Field
	VideoCard         apijson.Field
	VideoCodecs       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProfileListResponseProfileFingerprintFingerprint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileFingerprintFingerprintJSON) RawJSON() string {
	return r.raw
}

type ProfileListResponseProfileFingerprintHeaders struct {
	Accept                  string                                           `json:"accept"`
	AcceptEncoding          string                                           `json:"accept-encoding"`
	AcceptLanguage          string                                           `json:"accept-language"`
	Dnt                     string                                           `json:"dnt"`
	SecChUa                 string                                           `json:"sec-ch-ua"`
	SecChUaMobile           string                                           `json:"sec-ch-ua-mobile"`
	SecChUaPlatform         string                                           `json:"sec-ch-ua-platform"`
	SecFetchDest            string                                           `json:"sec-fetch-dest"`
	SecFetchMode            string                                           `json:"sec-fetch-mode"`
	SecFetchSite            string                                           `json:"sec-fetch-site"`
	SecFetchUser            string                                           `json:"sec-fetch-user"`
	UpgradeInsecureRequests string                                           `json:"upgrade-insecure-requests"`
	UserAgent               string                                           `json:"user-agent" api:"required"`
	JSON                    profileListResponseProfileFingerprintHeadersJSON `json:"-"`
}

// profileListResponseProfileFingerprintHeadersJSON contains the JSON metadata for the struct [ProfileListResponseProfileFingerprintHeaders]
type profileListResponseProfileFingerprintHeadersJSON struct {
	Accept                  apijson.Field
	AcceptEncoding          apijson.Field
	AcceptLanguage          apijson.Field
	Dnt                     apijson.Field
	SecChUa                 apijson.Field
	SecChUaMobile           apijson.Field
	SecChUaPlatform         apijson.Field
	SecFetchDest            apijson.Field
	SecFetchMode            apijson.Field
	SecFetchSite            apijson.Field
	SecFetchUser            apijson.Field
	UpgradeInsecureRequests apijson.Field
	UserAgent               apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProfileListResponseProfileFingerprintHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileFingerprintHeadersJSON) RawJSON() string {
	return r.raw
}

// ProfileListResponseProfileFingerprint The fingerprint associated with the profile
type ProfileListResponseProfileFingerprint struct {
	Fingerprint ProfileListResponseProfileFingerprintFingerprint `json:"fingerprint" api:"required"`
	Headers     ProfileListResponseProfileFingerprintHeaders     `json:"headers" api:"required"`
	JSON        profileListResponseProfileFingerprintJSON        `json:"-"`
}

// profileListResponseProfileFingerprintJSON contains the JSON metadata for the struct [ProfileListResponseProfileFingerprint]
type profileListResponseProfileFingerprintJSON struct {
	Fingerprint apijson.Field
	Headers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileListResponseProfileFingerprint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileFingerprintJSON) RawJSON() string {
	return r.raw
}

// ProfileListResponseProfile The list of profiles
type ProfileListResponseProfile struct {
	// The date and time when the profile was created
	CreatedAt time.Time `json:"createdAt" api:"required"`
	// The credentials configuration associated with the profile
	CredentialsConfig interface{} `json:"credentialsConfig" api:"required"`
	// The dimensions associated with the profile
	Dimensions ProfileListResponseProfileDimensions `json:"dimensions" api:"required"`
	// The extension IDs associated with the profile
	ExtensionIDs []string `json:"extensionIds" api:"required"`
	// The fingerprint associated with the profile
	Fingerprint ProfileListResponseProfileFingerprint `json:"fingerprint" api:"required"`
	// The unique identifier for the profile
	ID string `json:"id" api:"required"`
	// The project ID associated with the profile
	ProjectID string `json:"projectId" api:"required"`
	// The last session ID associated with the profile
	SourceSessionID string `json:"sourceSessionId" api:"required"`
	// The status of the profile
	Status ProfileListResponseProfilesItemStatus `json:"status" api:"required"`
	// The date and time when the profile was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required"`
	// The proxy configuration associated with the profile
	UseProxyConfig interface{} `json:"useProxyConfig" api:"required"`
	// The user agent associated with the profile
	UserAgent string                         `json:"userAgent" api:"required"`
	JSON      profileListResponseProfileJSON `json:"-"`
}

// profileListResponseProfileJSON contains the JSON metadata for the struct [ProfileListResponseProfile]
type profileListResponseProfileJSON struct {
	CreatedAt         apijson.Field
	CredentialsConfig apijson.Field
	Dimensions        apijson.Field
	ExtensionIDs      apijson.Field
	Fingerprint       apijson.Field
	ID                apijson.Field
	ProjectID         apijson.Field
	SourceSessionID   apijson.Field
	Status            apijson.Field
	UpdatedAt         apijson.Field
	UseProxyConfig    apijson.Field
	UserAgent         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProfileListResponseProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileListResponseProfileJSON) RawJSON() string {
	return r.raw
}

// ProfileGetResponseDimensions The dimensions associated with the profile
type ProfileGetResponseDimensions struct {
	Height float64                          `json:"height" api:"required"`
	Width  float64                          `json:"width" api:"required"`
	JSON   profileGetResponseDimensionsJSON `json:"-"`
}

// profileGetResponseDimensionsJSON contains the JSON metadata for the struct [ProfileGetResponseDimensions]
type profileGetResponseDimensionsJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileGetResponseDimensions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseDimensionsJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseFingerprintFingerprintBattery struct {
	Charging        bool                                                `json:"charging" api:"required"`
	ChargingTime    float64                                             `json:"chargingTime" api:"required"`
	DischargingTime float64                                             `json:"dischargingTime" api:"required"`
	Level           float64                                             `json:"level" api:"required"`
	JSON            profileGetResponseFingerprintFingerprintBatteryJSON `json:"-"`
}

// profileGetResponseFingerprintFingerprintBatteryJSON contains the JSON metadata for the struct [ProfileGetResponseFingerprintFingerprintBattery]
type profileGetResponseFingerprintFingerprintBatteryJSON struct {
	Charging        apijson.Field
	ChargingTime    apijson.Field
	DischargingTime apijson.Field
	Level           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProfileGetResponseFingerprintFingerprintBattery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseFingerprintFingerprintBatteryJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseFingerprintFingerprintMultimediaDevicesMicro struct {
	DeviceID string                                                             `json:"deviceId" api:"required"`
	GroupID  string                                                             `json:"groupId" api:"required"`
	Kind     string                                                             `json:"kind" api:"required"`
	Label    string                                                             `json:"label" api:"required"`
	JSON     profileGetResponseFingerprintFingerprintMultimediaDevicesMicroJSON `json:"-"`
}

// profileGetResponseFingerprintFingerprintMultimediaDevicesMicroJSON contains the JSON metadata for the struct [ProfileGetResponseFingerprintFingerprintMultimediaDevicesMicro]
type profileGetResponseFingerprintFingerprintMultimediaDevicesMicroJSON struct {
	DeviceID    apijson.Field
	GroupID     apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileGetResponseFingerprintFingerprintMultimediaDevicesMicro) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseFingerprintFingerprintMultimediaDevicesMicroJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseFingerprintFingerprintMultimediaDevicesSpeaker struct {
	DeviceID string                                                               `json:"deviceId" api:"required"`
	GroupID  string                                                               `json:"groupId" api:"required"`
	Kind     string                                                               `json:"kind" api:"required"`
	Label    string                                                               `json:"label" api:"required"`
	JSON     profileGetResponseFingerprintFingerprintMultimediaDevicesSpeakerJSON `json:"-"`
}

// profileGetResponseFingerprintFingerprintMultimediaDevicesSpeakerJSON contains the JSON metadata for the struct [ProfileGetResponseFingerprintFingerprintMultimediaDevicesSpeaker]
type profileGetResponseFingerprintFingerprintMultimediaDevicesSpeakerJSON struct {
	DeviceID    apijson.Field
	GroupID     apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileGetResponseFingerprintFingerprintMultimediaDevicesSpeaker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseFingerprintFingerprintMultimediaDevicesSpeakerJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseFingerprintFingerprintMultimediaDevicesWebcam struct {
	DeviceID string                                                              `json:"deviceId" api:"required"`
	GroupID  string                                                              `json:"groupId" api:"required"`
	Kind     string                                                              `json:"kind" api:"required"`
	Label    string                                                              `json:"label" api:"required"`
	JSON     profileGetResponseFingerprintFingerprintMultimediaDevicesWebcamJSON `json:"-"`
}

// profileGetResponseFingerprintFingerprintMultimediaDevicesWebcamJSON contains the JSON metadata for the struct [ProfileGetResponseFingerprintFingerprintMultimediaDevicesWebcam]
type profileGetResponseFingerprintFingerprintMultimediaDevicesWebcamJSON struct {
	DeviceID    apijson.Field
	GroupID     apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileGetResponseFingerprintFingerprintMultimediaDevicesWebcam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseFingerprintFingerprintMultimediaDevicesWebcamJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseFingerprintFingerprintMultimediaDevices struct {
	Micros   []ProfileGetResponseFingerprintFingerprintMultimediaDevicesMicro   `json:"micros" api:"required"`
	Speakers []ProfileGetResponseFingerprintFingerprintMultimediaDevicesSpeaker `json:"speakers" api:"required"`
	Webcams  []ProfileGetResponseFingerprintFingerprintMultimediaDevicesWebcam  `json:"webcams" api:"required"`
	JSON     profileGetResponseFingerprintFingerprintMultimediaDevicesJSON      `json:"-"`
}

// profileGetResponseFingerprintFingerprintMultimediaDevicesJSON contains the JSON metadata for the struct [ProfileGetResponseFingerprintFingerprintMultimediaDevices]
type profileGetResponseFingerprintFingerprintMultimediaDevicesJSON struct {
	Micros      apijson.Field
	Speakers    apijson.Field
	Webcams     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileGetResponseFingerprintFingerprintMultimediaDevices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseFingerprintFingerprintMultimediaDevicesJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseFingerprintFingerprintNavigatorExtraProperties struct {
	GlobalPrivacyControl bool                                                                 `json:"globalPrivacyControl" api:"required"`
	InstalledApps        []string                                                             `json:"installedApps" api:"required"`
	PdfViewerEnabled     bool                                                                 `json:"pdfViewerEnabled" api:"required"`
	VendorFlavors        []string                                                             `json:"vendorFlavors" api:"required"`
	JSON                 profileGetResponseFingerprintFingerprintNavigatorExtraPropertiesJSON `json:"-"`
}

// profileGetResponseFingerprintFingerprintNavigatorExtraPropertiesJSON contains the JSON metadata for the struct [ProfileGetResponseFingerprintFingerprintNavigatorExtraProperties]
type profileGetResponseFingerprintFingerprintNavigatorExtraPropertiesJSON struct {
	GlobalPrivacyControl apijson.Field
	InstalledApps        apijson.Field
	PdfViewerEnabled     apijson.Field
	VendorFlavors        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProfileGetResponseFingerprintFingerprintNavigatorExtraProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseFingerprintFingerprintNavigatorExtraPropertiesJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseFingerprintFingerprintNavigatorUserAgentDataBrand struct {
	Brand   string                                                                  `json:"brand" api:"required"`
	Version string                                                                  `json:"version" api:"required"`
	JSON    profileGetResponseFingerprintFingerprintNavigatorUserAgentDataBrandJSON `json:"-"`
}

// profileGetResponseFingerprintFingerprintNavigatorUserAgentDataBrandJSON contains the JSON metadata for the struct [ProfileGetResponseFingerprintFingerprintNavigatorUserAgentDataBrand]
type profileGetResponseFingerprintFingerprintNavigatorUserAgentDataBrandJSON struct {
	Brand       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileGetResponseFingerprintFingerprintNavigatorUserAgentDataBrand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseFingerprintFingerprintNavigatorUserAgentDataBrandJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseFingerprintFingerprintNavigatorUserAgentData struct {
	Brands   []ProfileGetResponseFingerprintFingerprintNavigatorUserAgentDataBrand `json:"brands" api:"required"`
	Mobile   bool                                                                  `json:"mobile" api:"required"`
	Platform string                                                                `json:"platform" api:"required"`
	JSON     profileGetResponseFingerprintFingerprintNavigatorUserAgentDataJSON    `json:"-"`
}

// profileGetResponseFingerprintFingerprintNavigatorUserAgentDataJSON contains the JSON metadata for the struct [ProfileGetResponseFingerprintFingerprintNavigatorUserAgentData]
type profileGetResponseFingerprintFingerprintNavigatorUserAgentDataJSON struct {
	Brands      apijson.Field
	Mobile      apijson.Field
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileGetResponseFingerprintFingerprintNavigatorUserAgentData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseFingerprintFingerprintNavigatorUserAgentDataJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseFingerprintFingerprintNavigator struct {
	AppCodeName         string                                                           `json:"appCodeName" api:"required"`
	AppName             string                                                           `json:"appName" api:"required"`
	AppVersion          string                                                           `json:"appVersion" api:"required"`
	DeviceMemory        float64                                                          `json:"deviceMemory" api:"required"`
	DoNotTrack          string                                                           `json:"doNotTrack"`
	ExtraProperties     ProfileGetResponseFingerprintFingerprintNavigatorExtraProperties `json:"extraProperties" api:"required"`
	HardwareConcurrency float64                                                          `json:"hardwareConcurrency" api:"required"`
	Language            string                                                           `json:"language" api:"required"`
	Languages           []string                                                         `json:"languages" api:"required"`
	MaxTouchPoints      float64                                                          `json:"maxTouchPoints" api:"required"`
	Oscpu               string                                                           `json:"oscpu" api:"required"`
	Platform            string                                                           `json:"platform" api:"required"`
	Product             string                                                           `json:"product" api:"required"`
	ProductSub          string                                                           `json:"productSub" api:"required"`
	UserAgent           string                                                           `json:"userAgent" api:"required"`
	UserAgentData       ProfileGetResponseFingerprintFingerprintNavigatorUserAgentData   `json:"userAgentData" api:"required"`
	Vendor              string                                                           `json:"vendor" api:"required"`
	VendorSub           string                                                           `json:"vendorSub" api:"required"`
	Webdriver           bool                                                             `json:"webdriver" api:"required"`
	JSON                profileGetResponseFingerprintFingerprintNavigatorJSON            `json:"-"`
}

// profileGetResponseFingerprintFingerprintNavigatorJSON contains the JSON metadata for the struct [ProfileGetResponseFingerprintFingerprintNavigator]
type profileGetResponseFingerprintFingerprintNavigatorJSON struct {
	AppCodeName         apijson.Field
	AppName             apijson.Field
	AppVersion          apijson.Field
	DeviceMemory        apijson.Field
	DoNotTrack          apijson.Field
	ExtraProperties     apijson.Field
	HardwareConcurrency apijson.Field
	Language            apijson.Field
	Languages           apijson.Field
	MaxTouchPoints      apijson.Field
	Oscpu               apijson.Field
	Platform            apijson.Field
	Product             apijson.Field
	ProductSub          apijson.Field
	UserAgent           apijson.Field
	UserAgentData       apijson.Field
	Vendor              apijson.Field
	VendorSub           apijson.Field
	Webdriver           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProfileGetResponseFingerprintFingerprintNavigator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseFingerprintFingerprintNavigatorJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseFingerprintFingerprintPluginsDataPluginMimeType struct {
	Description   string                                                                `json:"description" api:"required"`
	EnabledPlugin string                                                                `json:"enabledPlugin" api:"required"`
	Suffixes      string                                                                `json:"suffixes" api:"required"`
	Type          string                                                                `json:"type" api:"required"`
	JSON          profileGetResponseFingerprintFingerprintPluginsDataPluginMimeTypeJSON `json:"-"`
}

// profileGetResponseFingerprintFingerprintPluginsDataPluginMimeTypeJSON contains the JSON metadata for the struct [ProfileGetResponseFingerprintFingerprintPluginsDataPluginMimeType]
type profileGetResponseFingerprintFingerprintPluginsDataPluginMimeTypeJSON struct {
	Description   apijson.Field
	EnabledPlugin apijson.Field
	Suffixes      apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProfileGetResponseFingerprintFingerprintPluginsDataPluginMimeType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseFingerprintFingerprintPluginsDataPluginMimeTypeJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseFingerprintFingerprintPluginsDataPlugin struct {
	Description string                                                              `json:"description" api:"required"`
	Filename    string                                                              `json:"filename" api:"required"`
	MimeTypes   []ProfileGetResponseFingerprintFingerprintPluginsDataPluginMimeType `json:"mimeTypes" api:"required"`
	Name        string                                                              `json:"name" api:"required"`
	JSON        profileGetResponseFingerprintFingerprintPluginsDataPluginJSON       `json:"-"`
}

// profileGetResponseFingerprintFingerprintPluginsDataPluginJSON contains the JSON metadata for the struct [ProfileGetResponseFingerprintFingerprintPluginsDataPlugin]
type profileGetResponseFingerprintFingerprintPluginsDataPluginJSON struct {
	Description apijson.Field
	Filename    apijson.Field
	MimeTypes   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileGetResponseFingerprintFingerprintPluginsDataPlugin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseFingerprintFingerprintPluginsDataPluginJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseFingerprintFingerprintPluginsData struct {
	MimeTypes []string                                                    `json:"mimeTypes" api:"required"`
	Plugins   []ProfileGetResponseFingerprintFingerprintPluginsDataPlugin `json:"plugins" api:"required"`
	JSON      profileGetResponseFingerprintFingerprintPluginsDataJSON     `json:"-"`
}

// profileGetResponseFingerprintFingerprintPluginsDataJSON contains the JSON metadata for the struct [ProfileGetResponseFingerprintFingerprintPluginsData]
type profileGetResponseFingerprintFingerprintPluginsDataJSON struct {
	MimeTypes   apijson.Field
	Plugins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileGetResponseFingerprintFingerprintPluginsData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseFingerprintFingerprintPluginsDataJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseFingerprintFingerprintScreen struct {
	AvailHeight      float64                                            `json:"availHeight" api:"required"`
	AvailLeft        float64                                            `json:"availLeft" api:"required"`
	AvailTop         float64                                            `json:"availTop" api:"required"`
	AvailWidth       float64                                            `json:"availWidth" api:"required"`
	ClientHeight     float64                                            `json:"clientHeight" api:"required"`
	ClientWidth      float64                                            `json:"clientWidth" api:"required"`
	ColorDepth       float64                                            `json:"colorDepth" api:"required"`
	DevicePixelRatio float64                                            `json:"devicePixelRatio" api:"required"`
	HasHdr           bool                                               `json:"hasHDR" api:"required"`
	Height           float64                                            `json:"height" api:"required"`
	InnerHeight      float64                                            `json:"innerHeight" api:"required"`
	InnerWidth       float64                                            `json:"innerWidth" api:"required"`
	OuterHeight      float64                                            `json:"outerHeight" api:"required"`
	OuterWidth       float64                                            `json:"outerWidth" api:"required"`
	PageXOffset      float64                                            `json:"pageXOffset" api:"required"`
	PageYOffset      float64                                            `json:"pageYOffset" api:"required"`
	PixelDepth       float64                                            `json:"pixelDepth" api:"required"`
	ScreenX          float64                                            `json:"screenX" api:"required"`
	Width            float64                                            `json:"width" api:"required"`
	JSON             profileGetResponseFingerprintFingerprintScreenJSON `json:"-"`
}

// profileGetResponseFingerprintFingerprintScreenJSON contains the JSON metadata for the struct [ProfileGetResponseFingerprintFingerprintScreen]
type profileGetResponseFingerprintFingerprintScreenJSON struct {
	AvailHeight      apijson.Field
	AvailLeft        apijson.Field
	AvailTop         apijson.Field
	AvailWidth       apijson.Field
	ClientHeight     apijson.Field
	ClientWidth      apijson.Field
	ColorDepth       apijson.Field
	DevicePixelRatio apijson.Field
	HasHdr           apijson.Field
	Height           apijson.Field
	InnerHeight      apijson.Field
	InnerWidth       apijson.Field
	OuterHeight      apijson.Field
	OuterWidth       apijson.Field
	PageXOffset      apijson.Field
	PageYOffset      apijson.Field
	PixelDepth       apijson.Field
	ScreenX          apijson.Field
	Width            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProfileGetResponseFingerprintFingerprintScreen) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseFingerprintFingerprintScreenJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseFingerprintFingerprintVideoCard struct {
	Renderer string                                                `json:"renderer" api:"required"`
	Vendor   string                                                `json:"vendor" api:"required"`
	JSON     profileGetResponseFingerprintFingerprintVideoCardJSON `json:"-"`
}

// profileGetResponseFingerprintFingerprintVideoCardJSON contains the JSON metadata for the struct [ProfileGetResponseFingerprintFingerprintVideoCard]
type profileGetResponseFingerprintFingerprintVideoCardJSON struct {
	Renderer    apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileGetResponseFingerprintFingerprintVideoCard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseFingerprintFingerprintVideoCardJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseFingerprintFingerprint struct {
	AudioCodecs       map[string]string                                         `json:"audioCodecs" api:"required"`
	Battery           ProfileGetResponseFingerprintFingerprintBattery           `json:"battery" api:"required"`
	Fonts             []string                                                  `json:"fonts" api:"required"`
	MockWebRtc        bool                                                      `json:"mockWebRTC" api:"required"`
	MultimediaDevices ProfileGetResponseFingerprintFingerprintMultimediaDevices `json:"multimediaDevices" api:"required"`
	Navigator         ProfileGetResponseFingerprintFingerprintNavigator         `json:"navigator" api:"required"`
	PluginsData       ProfileGetResponseFingerprintFingerprintPluginsData       `json:"pluginsData" api:"required"`
	Screen            ProfileGetResponseFingerprintFingerprintScreen            `json:"screen" api:"required"`
	Slim              bool                                                      `json:"slim" api:"required"`
	VideoCard         ProfileGetResponseFingerprintFingerprintVideoCard         `json:"videoCard" api:"required"`
	VideoCodecs       map[string]string                                         `json:"videoCodecs" api:"required"`
	JSON              profileGetResponseFingerprintFingerprintJSON              `json:"-"`
}

// profileGetResponseFingerprintFingerprintJSON contains the JSON metadata for the struct [ProfileGetResponseFingerprintFingerprint]
type profileGetResponseFingerprintFingerprintJSON struct {
	AudioCodecs       apijson.Field
	Battery           apijson.Field
	Fonts             apijson.Field
	MockWebRtc        apijson.Field
	MultimediaDevices apijson.Field
	Navigator         apijson.Field
	PluginsData       apijson.Field
	Screen            apijson.Field
	Slim              apijson.Field
	VideoCard         apijson.Field
	VideoCodecs       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProfileGetResponseFingerprintFingerprint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseFingerprintFingerprintJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseFingerprintHeaders struct {
	Accept                  string                                   `json:"accept"`
	AcceptEncoding          string                                   `json:"accept-encoding"`
	AcceptLanguage          string                                   `json:"accept-language"`
	Dnt                     string                                   `json:"dnt"`
	SecChUa                 string                                   `json:"sec-ch-ua"`
	SecChUaMobile           string                                   `json:"sec-ch-ua-mobile"`
	SecChUaPlatform         string                                   `json:"sec-ch-ua-platform"`
	SecFetchDest            string                                   `json:"sec-fetch-dest"`
	SecFetchMode            string                                   `json:"sec-fetch-mode"`
	SecFetchSite            string                                   `json:"sec-fetch-site"`
	SecFetchUser            string                                   `json:"sec-fetch-user"`
	UpgradeInsecureRequests string                                   `json:"upgrade-insecure-requests"`
	UserAgent               string                                   `json:"user-agent" api:"required"`
	JSON                    profileGetResponseFingerprintHeadersJSON `json:"-"`
}

// profileGetResponseFingerprintHeadersJSON contains the JSON metadata for the struct [ProfileGetResponseFingerprintHeaders]
type profileGetResponseFingerprintHeadersJSON struct {
	Accept                  apijson.Field
	AcceptEncoding          apijson.Field
	AcceptLanguage          apijson.Field
	Dnt                     apijson.Field
	SecChUa                 apijson.Field
	SecChUaMobile           apijson.Field
	SecChUaPlatform         apijson.Field
	SecFetchDest            apijson.Field
	SecFetchMode            apijson.Field
	SecFetchSite            apijson.Field
	SecFetchUser            apijson.Field
	UpgradeInsecureRequests apijson.Field
	UserAgent               apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProfileGetResponseFingerprintHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseFingerprintHeadersJSON) RawJSON() string {
	return r.raw
}

// ProfileGetResponseFingerprint The fingerprint associated with the profile
type ProfileGetResponseFingerprint struct {
	Fingerprint ProfileGetResponseFingerprintFingerprint `json:"fingerprint" api:"required"`
	Headers     ProfileGetResponseFingerprintHeaders     `json:"headers" api:"required"`
	JSON        profileGetResponseFingerprintJSON        `json:"-"`
}

// profileGetResponseFingerprintJSON contains the JSON metadata for the struct [ProfileGetResponseFingerprint]
type profileGetResponseFingerprintJSON struct {
	Fingerprint apijson.Field
	Headers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileGetResponseFingerprint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseFingerprintJSON) RawJSON() string {
	return r.raw
}

// ProfileUpdateParamsDimensions The dimensions associated with the profile
type ProfileUpdateParamsDimensions struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// ProfileUpdateResponseDimensions The dimensions associated with the profile
type ProfileUpdateResponseDimensions struct {
	Height float64                             `json:"height" api:"required"`
	Width  float64                             `json:"width" api:"required"`
	JSON   profileUpdateResponseDimensionsJSON `json:"-"`
}

// profileUpdateResponseDimensionsJSON contains the JSON metadata for the struct [ProfileUpdateResponseDimensions]
type profileUpdateResponseDimensionsJSON struct {
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileUpdateResponseDimensions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseDimensionsJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponseFingerprintFingerprintBattery struct {
	Charging        bool                                                   `json:"charging" api:"required"`
	ChargingTime    float64                                                `json:"chargingTime" api:"required"`
	DischargingTime float64                                                `json:"dischargingTime" api:"required"`
	Level           float64                                                `json:"level" api:"required"`
	JSON            profileUpdateResponseFingerprintFingerprintBatteryJSON `json:"-"`
}

// profileUpdateResponseFingerprintFingerprintBatteryJSON contains the JSON metadata for the struct [ProfileUpdateResponseFingerprintFingerprintBattery]
type profileUpdateResponseFingerprintFingerprintBatteryJSON struct {
	Charging        apijson.Field
	ChargingTime    apijson.Field
	DischargingTime apijson.Field
	Level           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProfileUpdateResponseFingerprintFingerprintBattery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseFingerprintFingerprintBatteryJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponseFingerprintFingerprintMultimediaDevicesMicro struct {
	DeviceID string                                                                `json:"deviceId" api:"required"`
	GroupID  string                                                                `json:"groupId" api:"required"`
	Kind     string                                                                `json:"kind" api:"required"`
	Label    string                                                                `json:"label" api:"required"`
	JSON     profileUpdateResponseFingerprintFingerprintMultimediaDevicesMicroJSON `json:"-"`
}

// profileUpdateResponseFingerprintFingerprintMultimediaDevicesMicroJSON contains the JSON metadata for the struct [ProfileUpdateResponseFingerprintFingerprintMultimediaDevicesMicro]
type profileUpdateResponseFingerprintFingerprintMultimediaDevicesMicroJSON struct {
	DeviceID    apijson.Field
	GroupID     apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileUpdateResponseFingerprintFingerprintMultimediaDevicesMicro) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseFingerprintFingerprintMultimediaDevicesMicroJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponseFingerprintFingerprintMultimediaDevicesSpeaker struct {
	DeviceID string                                                                  `json:"deviceId" api:"required"`
	GroupID  string                                                                  `json:"groupId" api:"required"`
	Kind     string                                                                  `json:"kind" api:"required"`
	Label    string                                                                  `json:"label" api:"required"`
	JSON     profileUpdateResponseFingerprintFingerprintMultimediaDevicesSpeakerJSON `json:"-"`
}

// profileUpdateResponseFingerprintFingerprintMultimediaDevicesSpeakerJSON contains the JSON metadata for the struct [ProfileUpdateResponseFingerprintFingerprintMultimediaDevicesSpeaker]
type profileUpdateResponseFingerprintFingerprintMultimediaDevicesSpeakerJSON struct {
	DeviceID    apijson.Field
	GroupID     apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileUpdateResponseFingerprintFingerprintMultimediaDevicesSpeaker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseFingerprintFingerprintMultimediaDevicesSpeakerJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponseFingerprintFingerprintMultimediaDevicesWebcam struct {
	DeviceID string                                                                 `json:"deviceId" api:"required"`
	GroupID  string                                                                 `json:"groupId" api:"required"`
	Kind     string                                                                 `json:"kind" api:"required"`
	Label    string                                                                 `json:"label" api:"required"`
	JSON     profileUpdateResponseFingerprintFingerprintMultimediaDevicesWebcamJSON `json:"-"`
}

// profileUpdateResponseFingerprintFingerprintMultimediaDevicesWebcamJSON contains the JSON metadata for the struct [ProfileUpdateResponseFingerprintFingerprintMultimediaDevicesWebcam]
type profileUpdateResponseFingerprintFingerprintMultimediaDevicesWebcamJSON struct {
	DeviceID    apijson.Field
	GroupID     apijson.Field
	Kind        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileUpdateResponseFingerprintFingerprintMultimediaDevicesWebcam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseFingerprintFingerprintMultimediaDevicesWebcamJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponseFingerprintFingerprintMultimediaDevices struct {
	Micros   []ProfileUpdateResponseFingerprintFingerprintMultimediaDevicesMicro   `json:"micros" api:"required"`
	Speakers []ProfileUpdateResponseFingerprintFingerprintMultimediaDevicesSpeaker `json:"speakers" api:"required"`
	Webcams  []ProfileUpdateResponseFingerprintFingerprintMultimediaDevicesWebcam  `json:"webcams" api:"required"`
	JSON     profileUpdateResponseFingerprintFingerprintMultimediaDevicesJSON      `json:"-"`
}

// profileUpdateResponseFingerprintFingerprintMultimediaDevicesJSON contains the JSON metadata for the struct [ProfileUpdateResponseFingerprintFingerprintMultimediaDevices]
type profileUpdateResponseFingerprintFingerprintMultimediaDevicesJSON struct {
	Micros      apijson.Field
	Speakers    apijson.Field
	Webcams     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileUpdateResponseFingerprintFingerprintMultimediaDevices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseFingerprintFingerprintMultimediaDevicesJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponseFingerprintFingerprintNavigatorExtraProperties struct {
	GlobalPrivacyControl bool                                                                    `json:"globalPrivacyControl" api:"required"`
	InstalledApps        []string                                                                `json:"installedApps" api:"required"`
	PdfViewerEnabled     bool                                                                    `json:"pdfViewerEnabled" api:"required"`
	VendorFlavors        []string                                                                `json:"vendorFlavors" api:"required"`
	JSON                 profileUpdateResponseFingerprintFingerprintNavigatorExtraPropertiesJSON `json:"-"`
}

// profileUpdateResponseFingerprintFingerprintNavigatorExtraPropertiesJSON contains the JSON metadata for the struct [ProfileUpdateResponseFingerprintFingerprintNavigatorExtraProperties]
type profileUpdateResponseFingerprintFingerprintNavigatorExtraPropertiesJSON struct {
	GlobalPrivacyControl apijson.Field
	InstalledApps        apijson.Field
	PdfViewerEnabled     apijson.Field
	VendorFlavors        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProfileUpdateResponseFingerprintFingerprintNavigatorExtraProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseFingerprintFingerprintNavigatorExtraPropertiesJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponseFingerprintFingerprintNavigatorUserAgentDataBrand struct {
	Brand   string                                                                     `json:"brand" api:"required"`
	Version string                                                                     `json:"version" api:"required"`
	JSON    profileUpdateResponseFingerprintFingerprintNavigatorUserAgentDataBrandJSON `json:"-"`
}

// profileUpdateResponseFingerprintFingerprintNavigatorUserAgentDataBrandJSON contains the JSON metadata for the struct [ProfileUpdateResponseFingerprintFingerprintNavigatorUserAgentDataBrand]
type profileUpdateResponseFingerprintFingerprintNavigatorUserAgentDataBrandJSON struct {
	Brand       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileUpdateResponseFingerprintFingerprintNavigatorUserAgentDataBrand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseFingerprintFingerprintNavigatorUserAgentDataBrandJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponseFingerprintFingerprintNavigatorUserAgentData struct {
	Brands   []ProfileUpdateResponseFingerprintFingerprintNavigatorUserAgentDataBrand `json:"brands" api:"required"`
	Mobile   bool                                                                     `json:"mobile" api:"required"`
	Platform string                                                                   `json:"platform" api:"required"`
	JSON     profileUpdateResponseFingerprintFingerprintNavigatorUserAgentDataJSON    `json:"-"`
}

// profileUpdateResponseFingerprintFingerprintNavigatorUserAgentDataJSON contains the JSON metadata for the struct [ProfileUpdateResponseFingerprintFingerprintNavigatorUserAgentData]
type profileUpdateResponseFingerprintFingerprintNavigatorUserAgentDataJSON struct {
	Brands      apijson.Field
	Mobile      apijson.Field
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileUpdateResponseFingerprintFingerprintNavigatorUserAgentData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseFingerprintFingerprintNavigatorUserAgentDataJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponseFingerprintFingerprintNavigator struct {
	AppCodeName         string                                                              `json:"appCodeName" api:"required"`
	AppName             string                                                              `json:"appName" api:"required"`
	AppVersion          string                                                              `json:"appVersion" api:"required"`
	DeviceMemory        float64                                                             `json:"deviceMemory" api:"required"`
	DoNotTrack          string                                                              `json:"doNotTrack"`
	ExtraProperties     ProfileUpdateResponseFingerprintFingerprintNavigatorExtraProperties `json:"extraProperties" api:"required"`
	HardwareConcurrency float64                                                             `json:"hardwareConcurrency" api:"required"`
	Language            string                                                              `json:"language" api:"required"`
	Languages           []string                                                            `json:"languages" api:"required"`
	MaxTouchPoints      float64                                                             `json:"maxTouchPoints" api:"required"`
	Oscpu               string                                                              `json:"oscpu" api:"required"`
	Platform            string                                                              `json:"platform" api:"required"`
	Product             string                                                              `json:"product" api:"required"`
	ProductSub          string                                                              `json:"productSub" api:"required"`
	UserAgent           string                                                              `json:"userAgent" api:"required"`
	UserAgentData       ProfileUpdateResponseFingerprintFingerprintNavigatorUserAgentData   `json:"userAgentData" api:"required"`
	Vendor              string                                                              `json:"vendor" api:"required"`
	VendorSub           string                                                              `json:"vendorSub" api:"required"`
	Webdriver           bool                                                                `json:"webdriver" api:"required"`
	JSON                profileUpdateResponseFingerprintFingerprintNavigatorJSON            `json:"-"`
}

// profileUpdateResponseFingerprintFingerprintNavigatorJSON contains the JSON metadata for the struct [ProfileUpdateResponseFingerprintFingerprintNavigator]
type profileUpdateResponseFingerprintFingerprintNavigatorJSON struct {
	AppCodeName         apijson.Field
	AppName             apijson.Field
	AppVersion          apijson.Field
	DeviceMemory        apijson.Field
	DoNotTrack          apijson.Field
	ExtraProperties     apijson.Field
	HardwareConcurrency apijson.Field
	Language            apijson.Field
	Languages           apijson.Field
	MaxTouchPoints      apijson.Field
	Oscpu               apijson.Field
	Platform            apijson.Field
	Product             apijson.Field
	ProductSub          apijson.Field
	UserAgent           apijson.Field
	UserAgentData       apijson.Field
	Vendor              apijson.Field
	VendorSub           apijson.Field
	Webdriver           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProfileUpdateResponseFingerprintFingerprintNavigator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseFingerprintFingerprintNavigatorJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponseFingerprintFingerprintPluginsDataPluginMimeType struct {
	Description   string                                                                   `json:"description" api:"required"`
	EnabledPlugin string                                                                   `json:"enabledPlugin" api:"required"`
	Suffixes      string                                                                   `json:"suffixes" api:"required"`
	Type          string                                                                   `json:"type" api:"required"`
	JSON          profileUpdateResponseFingerprintFingerprintPluginsDataPluginMimeTypeJSON `json:"-"`
}

// profileUpdateResponseFingerprintFingerprintPluginsDataPluginMimeTypeJSON contains the JSON metadata for the struct [ProfileUpdateResponseFingerprintFingerprintPluginsDataPluginMimeType]
type profileUpdateResponseFingerprintFingerprintPluginsDataPluginMimeTypeJSON struct {
	Description   apijson.Field
	EnabledPlugin apijson.Field
	Suffixes      apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProfileUpdateResponseFingerprintFingerprintPluginsDataPluginMimeType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseFingerprintFingerprintPluginsDataPluginMimeTypeJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponseFingerprintFingerprintPluginsDataPlugin struct {
	Description string                                                                 `json:"description" api:"required"`
	Filename    string                                                                 `json:"filename" api:"required"`
	MimeTypes   []ProfileUpdateResponseFingerprintFingerprintPluginsDataPluginMimeType `json:"mimeTypes" api:"required"`
	Name        string                                                                 `json:"name" api:"required"`
	JSON        profileUpdateResponseFingerprintFingerprintPluginsDataPluginJSON       `json:"-"`
}

// profileUpdateResponseFingerprintFingerprintPluginsDataPluginJSON contains the JSON metadata for the struct [ProfileUpdateResponseFingerprintFingerprintPluginsDataPlugin]
type profileUpdateResponseFingerprintFingerprintPluginsDataPluginJSON struct {
	Description apijson.Field
	Filename    apijson.Field
	MimeTypes   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileUpdateResponseFingerprintFingerprintPluginsDataPlugin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseFingerprintFingerprintPluginsDataPluginJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponseFingerprintFingerprintPluginsData struct {
	MimeTypes []string                                                       `json:"mimeTypes" api:"required"`
	Plugins   []ProfileUpdateResponseFingerprintFingerprintPluginsDataPlugin `json:"plugins" api:"required"`
	JSON      profileUpdateResponseFingerprintFingerprintPluginsDataJSON     `json:"-"`
}

// profileUpdateResponseFingerprintFingerprintPluginsDataJSON contains the JSON metadata for the struct [ProfileUpdateResponseFingerprintFingerprintPluginsData]
type profileUpdateResponseFingerprintFingerprintPluginsDataJSON struct {
	MimeTypes   apijson.Field
	Plugins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileUpdateResponseFingerprintFingerprintPluginsData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseFingerprintFingerprintPluginsDataJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponseFingerprintFingerprintScreen struct {
	AvailHeight      float64                                               `json:"availHeight" api:"required"`
	AvailLeft        float64                                               `json:"availLeft" api:"required"`
	AvailTop         float64                                               `json:"availTop" api:"required"`
	AvailWidth       float64                                               `json:"availWidth" api:"required"`
	ClientHeight     float64                                               `json:"clientHeight" api:"required"`
	ClientWidth      float64                                               `json:"clientWidth" api:"required"`
	ColorDepth       float64                                               `json:"colorDepth" api:"required"`
	DevicePixelRatio float64                                               `json:"devicePixelRatio" api:"required"`
	HasHdr           bool                                                  `json:"hasHDR" api:"required"`
	Height           float64                                               `json:"height" api:"required"`
	InnerHeight      float64                                               `json:"innerHeight" api:"required"`
	InnerWidth       float64                                               `json:"innerWidth" api:"required"`
	OuterHeight      float64                                               `json:"outerHeight" api:"required"`
	OuterWidth       float64                                               `json:"outerWidth" api:"required"`
	PageXOffset      float64                                               `json:"pageXOffset" api:"required"`
	PageYOffset      float64                                               `json:"pageYOffset" api:"required"`
	PixelDepth       float64                                               `json:"pixelDepth" api:"required"`
	ScreenX          float64                                               `json:"screenX" api:"required"`
	Width            float64                                               `json:"width" api:"required"`
	JSON             profileUpdateResponseFingerprintFingerprintScreenJSON `json:"-"`
}

// profileUpdateResponseFingerprintFingerprintScreenJSON contains the JSON metadata for the struct [ProfileUpdateResponseFingerprintFingerprintScreen]
type profileUpdateResponseFingerprintFingerprintScreenJSON struct {
	AvailHeight      apijson.Field
	AvailLeft        apijson.Field
	AvailTop         apijson.Field
	AvailWidth       apijson.Field
	ClientHeight     apijson.Field
	ClientWidth      apijson.Field
	ColorDepth       apijson.Field
	DevicePixelRatio apijson.Field
	HasHdr           apijson.Field
	Height           apijson.Field
	InnerHeight      apijson.Field
	InnerWidth       apijson.Field
	OuterHeight      apijson.Field
	OuterWidth       apijson.Field
	PageXOffset      apijson.Field
	PageYOffset      apijson.Field
	PixelDepth       apijson.Field
	ScreenX          apijson.Field
	Width            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProfileUpdateResponseFingerprintFingerprintScreen) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseFingerprintFingerprintScreenJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponseFingerprintFingerprintVideoCard struct {
	Renderer string                                                   `json:"renderer" api:"required"`
	Vendor   string                                                   `json:"vendor" api:"required"`
	JSON     profileUpdateResponseFingerprintFingerprintVideoCardJSON `json:"-"`
}

// profileUpdateResponseFingerprintFingerprintVideoCardJSON contains the JSON metadata for the struct [ProfileUpdateResponseFingerprintFingerprintVideoCard]
type profileUpdateResponseFingerprintFingerprintVideoCardJSON struct {
	Renderer    apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileUpdateResponseFingerprintFingerprintVideoCard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseFingerprintFingerprintVideoCardJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponseFingerprintFingerprint struct {
	AudioCodecs       map[string]string                                            `json:"audioCodecs" api:"required"`
	Battery           ProfileUpdateResponseFingerprintFingerprintBattery           `json:"battery" api:"required"`
	Fonts             []string                                                     `json:"fonts" api:"required"`
	MockWebRtc        bool                                                         `json:"mockWebRTC" api:"required"`
	MultimediaDevices ProfileUpdateResponseFingerprintFingerprintMultimediaDevices `json:"multimediaDevices" api:"required"`
	Navigator         ProfileUpdateResponseFingerprintFingerprintNavigator         `json:"navigator" api:"required"`
	PluginsData       ProfileUpdateResponseFingerprintFingerprintPluginsData       `json:"pluginsData" api:"required"`
	Screen            ProfileUpdateResponseFingerprintFingerprintScreen            `json:"screen" api:"required"`
	Slim              bool                                                         `json:"slim" api:"required"`
	VideoCard         ProfileUpdateResponseFingerprintFingerprintVideoCard         `json:"videoCard" api:"required"`
	VideoCodecs       map[string]string                                            `json:"videoCodecs" api:"required"`
	JSON              profileUpdateResponseFingerprintFingerprintJSON              `json:"-"`
}

// profileUpdateResponseFingerprintFingerprintJSON contains the JSON metadata for the struct [ProfileUpdateResponseFingerprintFingerprint]
type profileUpdateResponseFingerprintFingerprintJSON struct {
	AudioCodecs       apijson.Field
	Battery           apijson.Field
	Fonts             apijson.Field
	MockWebRtc        apijson.Field
	MultimediaDevices apijson.Field
	Navigator         apijson.Field
	PluginsData       apijson.Field
	Screen            apijson.Field
	Slim              apijson.Field
	VideoCard         apijson.Field
	VideoCodecs       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProfileUpdateResponseFingerprintFingerprint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseFingerprintFingerprintJSON) RawJSON() string {
	return r.raw
}

type ProfileUpdateResponseFingerprintHeaders struct {
	Accept                  string                                      `json:"accept"`
	AcceptEncoding          string                                      `json:"accept-encoding"`
	AcceptLanguage          string                                      `json:"accept-language"`
	Dnt                     string                                      `json:"dnt"`
	SecChUa                 string                                      `json:"sec-ch-ua"`
	SecChUaMobile           string                                      `json:"sec-ch-ua-mobile"`
	SecChUaPlatform         string                                      `json:"sec-ch-ua-platform"`
	SecFetchDest            string                                      `json:"sec-fetch-dest"`
	SecFetchMode            string                                      `json:"sec-fetch-mode"`
	SecFetchSite            string                                      `json:"sec-fetch-site"`
	SecFetchUser            string                                      `json:"sec-fetch-user"`
	UpgradeInsecureRequests string                                      `json:"upgrade-insecure-requests"`
	UserAgent               string                                      `json:"user-agent" api:"required"`
	JSON                    profileUpdateResponseFingerprintHeadersJSON `json:"-"`
}

// profileUpdateResponseFingerprintHeadersJSON contains the JSON metadata for the struct [ProfileUpdateResponseFingerprintHeaders]
type profileUpdateResponseFingerprintHeadersJSON struct {
	Accept                  apijson.Field
	AcceptEncoding          apijson.Field
	AcceptLanguage          apijson.Field
	Dnt                     apijson.Field
	SecChUa                 apijson.Field
	SecChUaMobile           apijson.Field
	SecChUaPlatform         apijson.Field
	SecFetchDest            apijson.Field
	SecFetchMode            apijson.Field
	SecFetchSite            apijson.Field
	SecFetchUser            apijson.Field
	UpgradeInsecureRequests apijson.Field
	UserAgent               apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ProfileUpdateResponseFingerprintHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseFingerprintHeadersJSON) RawJSON() string {
	return r.raw
}

// ProfileUpdateResponseFingerprint The fingerprint associated with the profile
type ProfileUpdateResponseFingerprint struct {
	Fingerprint ProfileUpdateResponseFingerprintFingerprint `json:"fingerprint" api:"required"`
	Headers     ProfileUpdateResponseFingerprintHeaders     `json:"headers" api:"required"`
	JSON        profileUpdateResponseFingerprintJSON        `json:"-"`
}

// profileUpdateResponseFingerprintJSON contains the JSON metadata for the struct [ProfileUpdateResponseFingerprint]
type profileUpdateResponseFingerprintJSON struct {
	Fingerprint apijson.Field
	Headers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileUpdateResponseFingerprint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileUpdateResponseFingerprintJSON) RawJSON() string {
	return r.raw
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
