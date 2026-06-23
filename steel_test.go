package steel

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPdf(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Screenshot(context.Background(), ClientScreenshotParams{})
	if err != nil {
		t.Fatalf("client.Screenshot: %v", err)
	}
}
