package steel

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExtensionsList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Extensions.List(context.Background())
	if err != nil {
		t.Fatalf("client.Extensions.List: %v", err)
	}
}

func TestExtensionsUpload(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Extensions.Upload(context.Background(), ExtensionUploadParams{})
	if err != nil {
		t.Fatalf("client.Extensions.Upload: %v", err)
	}
}

func TestExtensionsDeleteAll(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Extensions.DeleteAll(context.Background())
	if err != nil {
		t.Fatalf("client.Extensions.DeleteAll: %v", err)
	}
}

func TestExtensionsDownload(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Extensions.Download(context.Background(), "id")
	if err != nil {
		t.Fatalf("client.Extensions.Download: %v", err)
	}
}

func TestExtensionsUpdate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Extensions.Update(context.Background(), "id", ExtensionUpdateParams{})
	if err != nil {
		t.Fatalf("client.Extensions.Update: %v", err)
	}
}

func TestExtensionsDelete(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Extensions.Delete(context.Background(), "id")
	if err != nil {
		t.Fatalf("client.Extensions.Delete: %v", err)
	}
}
