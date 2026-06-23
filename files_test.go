package steel

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFilesList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Files.List(context.Background())
	if err != nil {
		t.Fatalf("client.Files.List: %v", err)
	}
}

func TestFilesUpload(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Files.Upload(context.Background(), FileUploadParams{})
	if err != nil {
		t.Fatalf("client.Files.Upload: %v", err)
	}
}

func TestFilesDownload(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Files.Download(context.Background(), "id")
	if err != nil {
		t.Fatalf("client.Files.Download: %v", err)
	}
}

func TestFilesDelete(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	err := client.Files.Delete(context.Background(), "id")
	if err != nil {
		t.Fatalf("client.Files.Delete: %v", err)
	}
}
