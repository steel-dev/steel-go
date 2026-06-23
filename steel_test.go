package steel

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPdf(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/pdf" {
			t.Errorf("path = %s, want /v1/pdf", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Pdf(context.Background(), ClientPdfParams{})
	if err != nil {
		t.Fatalf("client.Pdf: %v", err)
	}
}

func TestScrape(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/scrape" {
			t.Errorf("path = %s, want /v1/scrape", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Scrape(context.Background(), ClientScrapeParams{})
	if err != nil {
		t.Fatalf("client.Scrape: %v", err)
	}
}

func TestScreenshot(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/screenshot" {
			t.Errorf("path = %s, want /v1/screenshot", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Screenshot(context.Background(), ClientScreenshotParams{})
	if err != nil {
		t.Fatalf("client.Screenshot: %v", err)
	}
}
