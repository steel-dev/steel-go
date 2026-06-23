# API reference

## credentials

- `GET /v1/credentials` — `client.Credentials.List(context.Background(), params)` → `CredentialListResponse`
- `POST /v1/credentials` — `client.Credentials.Create(context.Background(), params)` → `CredentialCreateResponse`
- `PUT /v1/credentials` — `client.Credentials.Update(context.Background(), params)` → `CredentialUpdateResponse`
- `DELETE /v1/credentials` — `client.Credentials.Delete(context.Background(), params)` → `CredentialDeleteResponse`

## extensions

- `GET /v1/extensions` — `client.Extensions.List(context.Background())` → `ExtensionListResponse`
- `POST /v1/extensions` — `client.Extensions.Upload(context.Background(), params)` → `ExtensionUploadResponse`
- `DELETE /v1/extensions` — `client.Extensions.DeleteAll(context.Background())` → `ExtensionDeleteAllResponse`
- `GET /v1/extensions/{extensionId}` — `client.Extensions.Download(context.Background(), params)`
- `PUT /v1/extensions/{extensionId}` — `client.Extensions.Update(context.Background(), params)` → `ExtensionUpdateResponse`
- `DELETE /v1/extensions/{extensionId}` — `client.Extensions.Delete(context.Background(), params)` → `ExtensionDeleteResponse`

## files

- `GET /v1/files` — `client.Files.List(context.Background())` → `Fileslist`
- `POST /v1/files` — `client.Files.Upload(context.Background(), params)` → `File`
- `GET /v1/files/{path}` — `client.Files.Download(context.Background(), params)`
- `DELETE /v1/files/{path}` — `client.Files.Delete(context.Background(), params)`

## Client methods

- `POST /v1/pdf` — `client.Pdf(context.Background(), params)` → `PdfResponse`
- `POST /v1/scrape` — `client.Scrape(context.Background(), params)` → `ScrapeResponse`
- `POST /v1/screenshot` — `client.Screenshot(context.Background(), params)` → `ScreenshotResponse`

## profiles

- `GET /v1/profiles` — `client.Profiles.List(context.Background(), params)` → `ProfileListResponse`
- `POST /v1/profiles` — `client.Profiles.Create(context.Background(), params)` → `ProfileCreateResponse`
- `GET /v1/profiles/{id}` — `client.Profiles.Get(context.Background(), params)` → `ProfileGetResponse`
- `PATCH /v1/profiles/{id}` — `client.Profiles.Update(context.Background(), params)` → `ProfileUpdateResponse`

## sessions

- `GET /v1/sessions` — `client.Sessions.List(context.Background(), params)` → `Sessionslist` _(paginated)_
- `POST /v1/sessions` — `client.Sessions.Create(context.Background(), params)` → `Session`
- `POST /v1/sessions/release` — `client.Sessions.ReleaseAll(context.Background(), params)` → `SessionReleaseAllResponse`
- `GET /v1/sessions/{id}` — `client.Sessions.Retrieve(context.Background(), params)` → `Session`
- `GET /v1/sessions/{id}/agent-traces` — `client.Sessions.AgentTraces(context.Background(), params)` → `SessionAgentTracesResponse`
- `GET /v1/sessions/{id}/context` — `client.Sessions.Context(context.Background(), params)` → `SessionContext`
- `GET /v1/sessions/{id}/cost` — `client.Sessions.Cost(context.Background(), params)` → `SessionCostResponse`
- `GET /v1/sessions/{id}/events` — `client.Sessions.Events(context.Background(), params)` → `SessionEventsResponse`
- `GET /v1/sessions/{id}/live-details` — `client.Sessions.LiveDetails(context.Background(), params)` → `SessionLiveDetailsResponse`
- `POST /v1/sessions/{id}/release` — `client.Sessions.Release(context.Background(), params)` → `SessionReleaseResponse`
- `POST /v1/sessions/{sessionId}/computer` — `client.Sessions.Computer(context.Background(), params)` → `SessionComputerResponse`

## sessions.captchas

- `POST /v1/sessions/{sessionId}/captchas/solve` — `client.Sessions.Captchas.Solve(context.Background(), params)` → `CaptchaSolveResponse`
- `POST /v1/sessions/{sessionId}/captchas/solve-image` — `client.Sessions.Captchas.SolveImage(context.Background(), params)` → `CaptchaSolveImageResponse`
- `GET /v1/sessions/{sessionId}/captchas/status` — `client.Sessions.Captchas.Status(context.Background(), params)` → `CaptchaStatusResponse`

## sessions.files

- `GET /v1/sessions/{sessionId}/files` — `client.Sessions.Files.List(context.Background(), params)` → `Fileslist`
- `POST /v1/sessions/{sessionId}/files` — `client.Sessions.Files.Upload(context.Background(), params)` → `File`
- `DELETE /v1/sessions/{sessionId}/files` — `client.Sessions.Files.DeleteAll(context.Background(), params)`
- `GET /v1/sessions/{sessionId}/files.zip` — `client.Sessions.Files.DownloadArchive(context.Background(), params)`
- `GET /v1/sessions/{sessionId}/files/{path}` — `client.Sessions.Files.Download(context.Background(), params)`
- `DELETE /v1/sessions/{sessionId}/files/{path}` — `client.Sessions.Files.Delete(context.Background(), params)`

