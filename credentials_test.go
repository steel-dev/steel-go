package steel

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCredentialsList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Credentials.List(context.Background(), CredentialListParams{})
	if err != nil {
		t.Fatalf("client.Credentials.List: %v", err)
	}
}

func TestCredentialsCreate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Credentials.Create(context.Background(), CredentialCreateParams{})
	if err != nil {
		t.Fatalf("client.Credentials.Create: %v", err)
	}
}

func TestCredentialsUpdate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Credentials.Update(context.Background(), CredentialUpdateParams{})
	if err != nil {
		t.Fatalf("client.Credentials.Update: %v", err)
	}
}

func TestCredentialsDelete(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Credentials.Delete(context.Background(), CredentialDeleteParams{})
	if err != nil {
		t.Fatalf("client.Credentials.Delete: %v", err)
	}
}
