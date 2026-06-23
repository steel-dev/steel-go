package steel

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCredentialsList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/v1/credentials" {
			t.Errorf("path = %s, want /v1/credentials", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
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
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/credentials" {
			t.Errorf("path = %s, want /v1/credentials", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
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
		if r.Method != "PUT" {
			t.Errorf("method = %s, want PUT", r.Method)
		}
		if r.URL.Path != "/v1/credentials" {
			t.Errorf("path = %s, want /v1/credentials", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
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
		if r.Method != "DELETE" {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		if r.URL.Path != "/v1/credentials" {
			t.Errorf("path = %s, want /v1/credentials", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Credentials.Delete(context.Background(), CredentialDeleteParams{})
	if err != nil {
		t.Fatalf("client.Credentials.Delete: %v", err)
	}
}
