package steel

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSessionsList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/v1/sessions" {
			t.Errorf("path = %s, want /v1/sessions", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.List(context.Background(), SessionListParams{})
	if err != nil {
		t.Fatalf("client.Sessions.List: %v", err)
	}
}

func TestSessionsListAutoPaging(t *testing.T) {
	var calls int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if calls == 0 {
			_, _ = w.Write([]byte(`{"sessions":[{}],"nextCursor":"c2"}`))
		} else {
			_, _ = w.Write([]byte(`{"sessions":[{}]}`))
		}
		calls++
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	iter := client.Sessions.ListAutoPaging(context.Background(), SessionListParams{})
	count := 0
	for iter.Next() {
		_ = iter.Current()
		count++
		if count > 10 {
			t.Fatal("pagination did not terminate")
		}
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("AutoPaging error: %v", err)
	}
	if count != 2 {
		t.Errorf("items = %d, want 2", count)
	}
	if calls != 2 {
		t.Errorf("requests = %d, want 2", calls)
	}
}

func TestSessionsCreate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/sessions" {
			t.Errorf("path = %s, want /v1/sessions", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.Create(context.Background(), SessionCreateParams{})
	if err != nil {
		t.Fatalf("client.Sessions.Create: %v", err)
	}
}

func TestSessionsReleaseAll(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/sessions/release" {
			t.Errorf("path = %s, want /v1/sessions/release", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.ReleaseAll(context.Background(), SessionReleaseAllParams{}, SessionReleaseAllQueryParams{})
	if err != nil {
		t.Fatalf("client.Sessions.ReleaseAll: %v", err)
	}
}

func TestSessionsRetrieve(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/v1/sessions/id" {
			t.Errorf("path = %s, want /v1/sessions/id", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.Retrieve(context.Background(), "id")
	if err != nil {
		t.Fatalf("client.Sessions.Retrieve: %v", err)
	}
}

func TestSessionsAgentTraces(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/v1/sessions/id/agent-traces" {
			t.Errorf("path = %s, want /v1/sessions/id/agent-traces", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.AgentTraces(context.Background(), "id", SessionAgentTracesParams{})
	if err != nil {
		t.Fatalf("client.Sessions.AgentTraces: %v", err)
	}
}

func TestSessionsContext(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/v1/sessions/id/context" {
			t.Errorf("path = %s, want /v1/sessions/id/context", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.Context(context.Background(), "id")
	if err != nil {
		t.Fatalf("client.Sessions.Context: %v", err)
	}
}

func TestSessionsCost(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/v1/sessions/id/cost" {
			t.Errorf("path = %s, want /v1/sessions/id/cost", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.Cost(context.Background(), "id")
	if err != nil {
		t.Fatalf("client.Sessions.Cost: %v", err)
	}
}

func TestSessionsEvents(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/v1/sessions/id/events" {
			t.Errorf("path = %s, want /v1/sessions/id/events", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("[]"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.Events(context.Background(), "id", SessionEventsParams{})
	if err != nil {
		t.Fatalf("client.Sessions.Events: %v", err)
	}
}

func TestSessionsLiveDetails(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/v1/sessions/id/live-details" {
			t.Errorf("path = %s, want /v1/sessions/id/live-details", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.LiveDetails(context.Background(), "id")
	if err != nil {
		t.Fatalf("client.Sessions.LiveDetails: %v", err)
	}
}

func TestSessionsRelease(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/sessions/id/release" {
			t.Errorf("path = %s, want /v1/sessions/id/release", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.Release(context.Background(), "id", SessionReleaseParams{})
	if err != nil {
		t.Fatalf("client.Sessions.Release: %v", err)
	}
}

func TestSessionsComputer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/sessions/id/computer" {
			t.Errorf("path = %s, want /v1/sessions/id/computer", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.Computer(context.Background(), "id", SessionComputerParams{})
	if err != nil {
		t.Fatalf("client.Sessions.Computer: %v", err)
	}
}

func TestSessionsCaptchasSolve(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/sessions/id/captchas/solve" {
			t.Errorf("path = %s, want /v1/sessions/id/captchas/solve", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.Captchas.Solve(context.Background(), "id", CaptchaSolveParams{})
	if err != nil {
		t.Fatalf("client.Sessions.Captchas.Solve: %v", err)
	}
}

func TestSessionsCaptchasSolveImage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/sessions/id/captchas/solve-image" {
			t.Errorf("path = %s, want /v1/sessions/id/captchas/solve-image", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.Captchas.SolveImage(context.Background(), "id", CaptchaSolveImageParams{})
	if err != nil {
		t.Fatalf("client.Sessions.Captchas.SolveImage: %v", err)
	}
}

func TestSessionsCaptchasStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/v1/sessions/id/captchas/status" {
			t.Errorf("path = %s, want /v1/sessions/id/captchas/status", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("[]"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.Captchas.Status(context.Background(), "id")
	if err != nil {
		t.Fatalf("client.Sessions.Captchas.Status: %v", err)
	}
}

func TestSessionsFilesList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/v1/sessions/id/files" {
			t.Errorf("path = %s, want /v1/sessions/id/files", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.Files.List(context.Background(), "id")
	if err != nil {
		t.Fatalf("client.Sessions.Files.List: %v", err)
	}
}

func TestSessionsFilesUpload(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/sessions/id/files" {
			t.Errorf("path = %s, want /v1/sessions/id/files", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.Files.Upload(context.Background(), "id", SessionFileUploadParams{})
	if err != nil {
		t.Fatalf("client.Sessions.Files.Upload: %v", err)
	}
}

func TestSessionsFilesDeleteAll(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		if r.URL.Path != "/v1/sessions/id/files" {
			t.Errorf("path = %s, want /v1/sessions/id/files", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	err := client.Sessions.Files.DeleteAll(context.Background(), "id")
	if err != nil {
		t.Fatalf("client.Sessions.Files.DeleteAll: %v", err)
	}
}

func TestSessionsFilesDownloadArchive(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/v1/sessions/id/files.zip" {
			t.Errorf("path = %s, want /v1/sessions/id/files.zip", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.Files.DownloadArchive(context.Background(), "id")
	if err != nil {
		t.Fatalf("client.Sessions.Files.DownloadArchive: %v", err)
	}
}

func TestSessionsFilesDownload(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/v1/sessions/id/files/id" {
			t.Errorf("path = %s, want /v1/sessions/id/files/id", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.Files.Download(context.Background(), "id", "id")
	if err != nil {
		t.Fatalf("client.Sessions.Files.Download: %v", err)
	}
}

func TestSessionsFilesDelete(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		if r.URL.Path != "/v1/sessions/id/files/id" {
			t.Errorf("path = %s, want /v1/sessions/id/files/id", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	err := client.Sessions.Files.Delete(context.Background(), "id", "id")
	if err != nil {
		t.Fatalf("client.Sessions.Files.Delete: %v", err)
	}
}
