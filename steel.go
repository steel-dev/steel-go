// Package steel provides a client for the API.
//
// Steel API client
package steel

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"mime/multipart"
	"net"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const defaultBaseURL = "https://api.steel.dev"

type Client struct {
	apiKey        string
	baseURL       string
	httpClient    *http.Client
	maxRetries    int
	defaultHeader http.Header
	defaultQuery  url.Values

	Credentials *CredentialService
	Extensions  *ExtensionService
	Files       *FileService
	Profiles    *ProfileService
	Sessions    *SessionService
}

type Option func(*Client)

func WithMaxRetries(maxRetries int) Option {
	return func(c *Client) { c.maxRetries = maxRetries }
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) { c.httpClient.Timeout = timeout }
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) { c.httpClient = httpClient }
}

func WithHeader(key, value string) Option {
	return func(c *Client) { c.defaultHeader.Set(key, value) }
}

func WithQuery(key, value string) Option {
	return func(c *Client) { c.defaultQuery.Set(key, value) }
}

type RequestOption func(*requestConfig)

type requestConfig struct {
	header         http.Header
	query          url.Values
	maxRetries     *int
	timeout        *time.Duration
	idempotencyKey string
	responseInto   **http.Response
}

func (c *Client) newRequestConfig(opts []RequestOption) *requestConfig {
	cfg := &requestConfig{header: http.Header{}, query: url.Values{}}
	for key, values := range c.defaultHeader {
		cfg.header[key] = append([]string(nil), values...)
	}
	for key, values := range c.defaultQuery {
		cfg.query[key] = append([]string(nil), values...)
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func (cfg *requestConfig) retries(def int) int {
	if cfg.maxRetries != nil {
		return *cfg.maxRetries
	}
	return def
}

func (cfg *requestConfig) idemKey(method string) string {
	if cfg.idempotencyKey != "" {
		return cfg.idempotencyKey
	}
	return newRetryKey(method)
}

func (cfg *requestConfig) apply(req *http.Request) {
	for key, values := range cfg.header {
		for _, value := range values {
			req.Header.Set(key, value)
		}
	}
}

func (cfg *requestConfig) urlFor(baseURL, path string) string {
	baseURL = strings.TrimSuffix(baseURL, "/")
	if len(cfg.query) == 0 {
		return baseURL + path
	}
	sep := "?"
	if strings.Contains(path, "?") {
		sep = "&"
	}
	return baseURL + path + sep + cfg.query.Encode()
}

func (cfg *requestConfig) capture(resp *http.Response, data []byte) {
	if cfg.responseInto != nil {
		resp.Body.Close()
		resp.Body = io.NopCloser(bytes.NewReader(data))
		*cfg.responseInto = resp
	} else {
		resp.Body.Close()
	}
}

func WithRequestHeader(key, value string) RequestOption {
	return func(c *requestConfig) { c.header.Set(key, value) }
}

func WithRequestQuery(key, value string) RequestOption {
	return func(c *requestConfig) { c.query.Set(key, value) }
}

func WithRequestMaxRetries(maxRetries int) RequestOption {
	return func(c *requestConfig) { c.maxRetries = &maxRetries }
}

func WithRequestTimeout(timeout time.Duration) RequestOption {
	return func(c *requestConfig) { c.timeout = &timeout }
}

func WithIdempotencyKey(key string) RequestOption {
	return func(c *requestConfig) { c.idempotencyKey = key }
}

func WithResponseInto(dst **http.Response) RequestOption {
	return func(c *requestConfig) { c.responseInto = dst }
}

func NewClient(apiKey string, opts ...Option) *Client {
	baseURL := defaultBaseURL
	if v := os.Getenv("STEEL_BASE_URL"); v != "" {
		baseURL = v
	}
	return newClient(apiKey, baseURL, opts)
}

func NewClientWithBaseURL(apiKey, baseURL string, opts ...Option) *Client {
	return newClient(apiKey, baseURL, opts)
}

func newClient(apiKey, baseURL string, opts []Option) *Client {
	if apiKey == "" {
		apiKey = os.Getenv("STEEL_API_KEY")
	}
	c := &Client{
		apiKey:        apiKey,
		baseURL:       baseURL,
		httpClient:    &http.Client{Timeout: 60 * time.Second},
		maxRetries:    2,
		defaultHeader: http.Header{},
		defaultQuery:  url.Values{},
	}
	c.Credentials = &CredentialService{client: c}
	c.Extensions = &ExtensionService{client: c}
	c.Files = &FileService{client: c}
	c.Profiles = &ProfileService{client: c}
	c.Sessions = &SessionService{client: c, Captchas: &SessionServiceCaptchas{client: c}, Files: &SessionServiceFiles{client: c}}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Convert webpage to PDF
func (c *Client) Pdf(ctx context.Context, body ClientPdfParams, opts ...RequestOption) (*PdfResponse, error) {
	path := "/v1/pdf"
	var out PdfResponse
	if err := c.do(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Scrape webpage content
func (c *Client) Scrape(ctx context.Context, body ClientScrapeParams, opts ...RequestOption) (*ScrapeResponse, error) {
	path := "/v1/scrape"
	var out ScrapeResponse
	if err := c.do(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Capture webpage screenshot
func (c *Client) Screenshot(ctx context.Context, body ClientScreenshotParams, opts ...RequestOption) (*ScreenshotResponse, error) {
	path := "/v1/screenshot"
	var out ScreenshotResponse
	if err := c.do(ctx, "POST", path, body, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

type APIError struct {
	StatusCode int
	Message    string
	Header     http.Header
	RequestID  string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("api error: status %d: %s", e.StatusCode, e.Message)
}

type BadRequestError struct{ *APIError }
type AuthenticationError struct{ *APIError }
type PermissionDeniedError struct{ *APIError }
type NotFoundError struct{ *APIError }
type ConflictError struct{ *APIError }
type UnprocessableEntityError struct{ *APIError }
type RateLimitError struct{ *APIError }
type InternalServerError struct{ *APIError }

func (e *BadRequestError) Unwrap() error          { return e.APIError }
func (e *AuthenticationError) Unwrap() error      { return e.APIError }
func (e *PermissionDeniedError) Unwrap() error    { return e.APIError }
func (e *NotFoundError) Unwrap() error            { return e.APIError }
func (e *ConflictError) Unwrap() error            { return e.APIError }
func (e *UnprocessableEntityError) Unwrap() error { return e.APIError }
func (e *RateLimitError) Unwrap() error           { return e.APIError }
func (e *InternalServerError) Unwrap() error      { return e.APIError }

type APIConnectionError struct{ Err error }

func (e *APIConnectionError) Error() string {
	return fmt.Sprintf("connection error: %v", e.Err)
}

func (e *APIConnectionError) Unwrap() error { return e.Err }

type APITimeoutError struct{ Err error }

func (e *APITimeoutError) Error() string { return "request timed out" }

func (e *APITimeoutError) Unwrap() error { return e.Err }

func newConnectionError(err error) error {
	var netErr net.Error
	if errors.Is(err, context.DeadlineExceeded) || (errors.As(err, &netErr) && netErr.Timeout()) {
		return &APITimeoutError{Err: err}
	}
	return &APIConnectionError{Err: err}
}

func extractErrorMessage(body []byte) (string, bool) {
	var m map[string]json.RawMessage
	if json.Unmarshal(body, &m) != nil {
		return "", false
	}
	pick := func(raw json.RawMessage) (string, bool) {
		var str string
		if json.Unmarshal(raw, &str) == nil && str != "" {
			return str, true
		}
		return "", false
	}
	for _, k := range []string{"message", "detail", "error_message", "display_message", "title"} {
		if raw, ok := m[k]; ok {
			if str, ok := pick(raw); ok {
				return str, true
			}
		}
	}
	if raw, ok := m["error"]; ok {
		var inner map[string]json.RawMessage
		if json.Unmarshal(raw, &inner) == nil {
			if mr, ok := inner["message"]; ok {
				if str, ok := pick(mr); ok {
					return str, true
				}
			}
		} else if str, ok := pick(raw); ok {
			return str, true
		}
	}
	if raw, ok := m["errors"]; ok {
		var arr []map[string]json.RawMessage
		if json.Unmarshal(raw, &arr) == nil && len(arr) > 0 {
			if mr, ok := arr[0]["message"]; ok {
				if str, ok := pick(mr); ok {
					return str, true
				}
			}
		}
	}
	return "", false
}

func newAPIError(statusCode int, body []byte, header http.Header) error {
	message := string(body)
	if m, ok := extractErrorMessage(body); ok {
		message = m
	}
	base := &APIError{StatusCode: statusCode, Message: message, Header: header, RequestID: header.Get("X-Request-Id")}
	switch {
	case statusCode == 400:
		return &BadRequestError{base}
	case statusCode == 401:
		return &AuthenticationError{base}
	case statusCode == 403:
		return &PermissionDeniedError{base}
	case statusCode == 404:
		return &NotFoundError{base}
	case statusCode == 409:
		return &ConflictError{base}
	case statusCode == 422:
		return &UnprocessableEntityError{base}
	case statusCode == 429:
		return &RateLimitError{base}
	case statusCode >= 500 && statusCode < 600:
		return &InternalServerError{base}
	default:
		return base
	}
}

func shouldRetryStatus(statusCode int) bool {
	switch statusCode {
	case 408, 409, 429:
		return true
	}
	return statusCode >= 500
}

func shouldRetryResponse(resp *http.Response) bool {
	switch resp.Header.Get("x-should-retry") {
	case "true":
		return true
	case "false":
		return false
	}
	return shouldRetryStatus(resp.StatusCode)
}

func parseRetryAfter(resp *http.Response) time.Duration {
	if v := resp.Header.Get("retry-after-ms"); v != "" {
		if ms, err := strconv.Atoi(strings.TrimSpace(v)); err == nil {
			return time.Duration(ms) * time.Millisecond
		}
	}
	if v := resp.Header.Get("retry-after"); v != "" {
		if secs, err := strconv.Atoi(strings.TrimSpace(v)); err == nil {
			return time.Duration(secs) * time.Second
		}
	}
	return 0
}

func backoffDuration(attempt int, retryAfter time.Duration) time.Duration {
	if retryAfter > 0 {
		if retryAfter > 60*time.Second {
			return 60 * time.Second
		}
		return retryAfter
	}
	exp := 0.5 * math.Pow(2, float64(attempt-1))
	if exp > 8 {
		exp = 8
	}
	jitter := 0.5 + rand.Float64()*0.5
	return time.Duration(exp * jitter * float64(time.Second))
}

func sleepCtx(ctx context.Context, d time.Duration) {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
	case <-t.C:
	}
}

func newRetryKey(method string) string {
	if method == "GET" {
		return ""
	}
	const hexd = "0123456789abcdef"
	b := make([]byte, 32)
	for i := range b {
		b[i] = hexd[rand.Intn(16)]
	}
	return string(b)
}

func setRuntimeHeaders(req *http.Request, idempotencyKey string) {
	req.Header.Set("user-agent", "steel-go/0.1.0")
	req.Header.Set("x-steel-lang", "go")
	req.Header.Set("x-steel-package-version", "0.1.0")
	if idempotencyKey != "" {
		req.Header.Set("idempotency-key", idempotencyKey)
	}
}

func (c *Client) do(ctx context.Context, method, path string, body, out interface{}, opts ...RequestOption) error {
	cfg := c.newRequestConfig(opts)
	if cfg.timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *cfg.timeout)
		defer cancel()
	}
	maxRetries := cfg.retries(c.maxRetries)
	var bodyBytes []byte
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyBytes = buf
	}

	idemKey := cfg.idemKey(method)

	for attempt := 0; ; attempt++ {
		var reader io.Reader
		if bodyBytes != nil {
			reader = bytes.NewReader(bodyBytes)
		}
		req, err := http.NewRequestWithContext(ctx, method, cfg.urlFor(c.baseURL, path), reader)
		if err != nil {
			return err
		}
		if bodyBytes != nil {
			req.Header.Set("content-type", "application/json")
		}
		req.Header.Set("accept", "application/json")
		req.Header.Set("steel-api-key", c.apiKey)
		setRuntimeHeaders(req, idemKey)
		cfg.apply(req)
		resp, err := c.httpClient.Do(req)
		if err != nil {
			if ctx.Err() == nil && attempt < maxRetries {
				sleepCtx(ctx, backoffDuration(attempt+1, 0))
				continue
			}
			return newConnectionError(err)
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			data, err := io.ReadAll(resp.Body)
			cfg.capture(resp, data)
			if err != nil {
				return err
			}
			if out != nil && len(data) > 0 {
				if err := json.Unmarshal(data, out); err != nil {
					return err
				}
			}
			return nil
		}

		if shouldRetryResponse(resp) && attempt < maxRetries {
			wait := backoffDuration(attempt+1, parseRetryAfter(resp))
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
			sleepCtx(ctx, wait)
			continue
		}

		data, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return newAPIError(resp.StatusCode, data, resp.Header)
	}
}

func (c *Client) doRaw(ctx context.Context, method, path string, body interface{}, opts ...RequestOption) ([]byte, error) {
	cfg := c.newRequestConfig(opts)
	if cfg.timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *cfg.timeout)
		defer cancel()
	}
	maxRetries := cfg.retries(c.maxRetries)
	var bodyBytes []byte
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyBytes = buf
	}

	idemKey := cfg.idemKey(method)

	for attempt := 0; ; attempt++ {
		var reader io.Reader
		if bodyBytes != nil {
			reader = bytes.NewReader(bodyBytes)
		}
		req, err := http.NewRequestWithContext(ctx, method, cfg.urlFor(c.baseURL, path), reader)
		if err != nil {
			return nil, err
		}
		if bodyBytes != nil {
			req.Header.Set("content-type", "application/json")
		}
		req.Header.Set("steel-api-key", c.apiKey)
		setRuntimeHeaders(req, idemKey)
		cfg.apply(req)
		resp, err := c.httpClient.Do(req)
		if err != nil {
			if ctx.Err() == nil && attempt < maxRetries {
				sleepCtx(ctx, backoffDuration(attempt+1, 0))
				continue
			}
			return nil, newConnectionError(err)
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			data, err := io.ReadAll(resp.Body)
			cfg.capture(resp, data)
			return data, err
		}

		if shouldRetryResponse(resp) && attempt < maxRetries {
			wait := backoffDuration(attempt+1, parseRetryAfter(resp))
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
			sleepCtx(ctx, wait)
			continue
		}

		data, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, newAPIError(resp.StatusCode, data, resp.Header)
	}
}

type FileUpload struct {
	Name        string
	Content     []byte
	ContentType string
}

var mpQuoteEscaper = strings.NewReplacer("\\", "\\\\", "\"", "\\\"")

func writeFilePart(w *multipart.Writer, name string, file FileUpload) error {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, mpQuoteEscaper.Replace(name), mpQuoteEscaper.Replace(file.Name)))
	contentType := file.ContentType
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	h.Set("Content-Type", contentType)
	fw, err := w.CreatePart(h)
	if err != nil {
		return err
	}
	_, err = fw.Write(file.Content)
	return err
}

func isFileUnion(t reflect.Type) bool {
	if t.Kind() != reflect.Struct || t.NumField() == 0 {
		return false
	}
	fileType := reflect.TypeOf(FileUpload{})
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i).Type
		for ft.Kind() == reflect.Ptr {
			ft = ft.Elem()
		}
		isFile := ft == fileType
		isFileArray := ft.Kind() == reflect.Slice && ft.Elem() == fileType
		if !isFile && !isFileArray {
			return false
		}
	}
	return true
}

func writeMultipartField(w *multipart.Writer, name string, fv reflect.Value) error {
	for fv.Kind() == reflect.Ptr {
		if fv.IsNil() {
			return nil
		}
		fv = fv.Elem()
	}
	if file, ok := fv.Interface().(FileUpload); ok {
		return writeFilePart(w, name, file)
	}
	if files, ok := fv.Interface().([]FileUpload); ok {
		for _, file := range files {
			if err := writeFilePart(w, name+"[]", file); err != nil {
				return err
			}
		}
		return nil
	}
	if fv.Kind() == reflect.Struct && isFileUnion(fv.Type()) {
		for i := 0; i < fv.NumField(); i++ {
			if err := writeMultipartField(w, name, fv.Field(i)); err != nil {
				return err
			}
		}
		return nil
	}
	if fv.Kind() == reflect.String {
		return w.WriteField(name, fv.String())
	}
	b, err := json.Marshal(fv.Interface())
	if err != nil {
		return err
	}
	return w.WriteField(name, string(b))
}

func buildMultipart(body interface{}) (*bytes.Buffer, string, error) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)
	v := reflect.ValueOf(body)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := field.Tag.Get("json")
		if idx := strings.Index(name, ","); idx >= 0 {
			name = name[:idx]
		}
		if name == "-" {
			continue
		}
		if name == "" {
			name = field.Name
		}
		fv := v.Field(i)
		if fv.Kind() == reflect.Ptr {
			if fv.IsNil() {
				continue
			}
			fv = fv.Elem()
		}
		if err := writeMultipartField(w, name, fv); err != nil {
			return nil, "", err
		}
	}
	if err := w.Close(); err != nil {
		return nil, "", err
	}
	return buf, w.FormDataContentType(), nil
}

func (c *Client) doMultipart(ctx context.Context, method, path string, body, out interface{}, opts ...RequestOption) error {
	cfg := c.newRequestConfig(opts)
	if cfg.timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *cfg.timeout)
		defer cancel()
	}
	maxRetries := cfg.retries(c.maxRetries)
	buf, contentType, err := buildMultipart(body)
	if err != nil {
		return err
	}

	idemKey := cfg.idemKey(method)

	for attempt := 0; ; attempt++ {
		req, err := http.NewRequestWithContext(ctx, method, cfg.urlFor(c.baseURL, path), bytes.NewReader(buf.Bytes()))
		if err != nil {
			return err
		}
		req.Header.Set("content-type", contentType)
		req.Header.Set("accept", "application/json")
		req.Header.Set("steel-api-key", c.apiKey)
		setRuntimeHeaders(req, idemKey)
		cfg.apply(req)
		resp, err := c.httpClient.Do(req)
		if err != nil {
			if ctx.Err() == nil && attempt < maxRetries {
				sleepCtx(ctx, backoffDuration(attempt+1, 0))
				continue
			}
			return newConnectionError(err)
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			data, err := io.ReadAll(resp.Body)
			cfg.capture(resp, data)
			if err != nil {
				return err
			}
			if out != nil && len(data) > 0 {
				if err := json.Unmarshal(data, out); err != nil {
					return err
				}
			}
			return nil
		}

		if shouldRetryResponse(resp) && attempt < maxRetries {
			wait := backoffDuration(attempt+1, parseRetryAfter(resp))
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
			sleepCtx(ctx, wait)
			continue
		}

		data, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return newAPIError(resp.StatusCode, data, resp.Header)
	}
}
