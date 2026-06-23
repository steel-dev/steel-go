package steel

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProfilesList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Profiles.List(context.Background(), ProfileListParams{})
	if err != nil {
		t.Fatalf("client.Profiles.List: %v", err)
	}
}

func TestProfilesCreate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Profiles.Create(context.Background(), ProfileCreateParams{})
	if err != nil {
		t.Fatalf("client.Profiles.Create: %v", err)
	}
}

func TestProfilesGet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Profiles.Get(context.Background(), "id", ProfileGetParams{})
	if err != nil {
		t.Fatalf("client.Profiles.Get: %v", err)
	}
}

func TestProfilesUpdate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Profiles.Update(context.Background(), "id", ProfileUpdateParams{}, ProfileUpdateQueryParams{})
	if err != nil {
		t.Fatalf("client.Profiles.Update: %v", err)
	}
}
