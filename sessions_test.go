package steel

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSessionsList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.List(context.Background(), SessionListParams{})
	if err != nil {
		t.Fatalf("client.Sessions.List: %v", err)
	}
}

func TestSessionsCreate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	_, err := client.Sessions.Files.Upload(context.Background(), "id", FileUploadParams2{})
	if err != nil {
		t.Fatalf("client.Sessions.Files.Upload: %v", err)
	}
}

func TestSessionsFilesDeleteAll(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	client := NewClientWithBaseURL("test-key", srv.URL)
	err := client.Sessions.Files.Delete(context.Background(), "id", "id")
	if err != nil {
		t.Fatalf("client.Sessions.Files.Delete: %v", err)
	}
}
