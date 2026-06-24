# API reference

## credentials

- `GET /v1/credentials` — `client.Credentials.List(context.Background(), params)` → `CredentialListResponse`
- `POST /v1/credentials` — `client.Credentials.Create(context.Background(), body)` → `CredentialCreateResponse`
- `PUT /v1/credentials` — `client.Credentials.Update(context.Background(), body)` → `CredentialUpdateResponse`
- `DELETE /v1/credentials` — `client.Credentials.Delete(context.Background(), body)` → `CredentialDeleteResponse`

## extensions

- `GET /v1/extensions` — `client.Extensions.List(context.Background())` → `ExtensionListResponse`
- `POST /v1/extensions` — `client.Extensions.Upload(context.Background(), body)` → `ExtensionUploadResponse`
- `DELETE /v1/extensions` — `client.Extensions.DeleteAll(context.Background())` → `ExtensionDeleteAllResponse`
- `GET /v1/extensions/{extensionId}` — `client.Extensions.Download(context.Background(), extensionID)`
- `PUT /v1/extensions/{extensionId}` — `client.Extensions.Update(context.Background(), extensionID, body)` → `ExtensionUpdateResponse`
- `DELETE /v1/extensions/{extensionId}` — `client.Extensions.Delete(context.Background(), extensionID)` → `ExtensionDeleteResponse`

## files

- `GET /v1/files` — `client.Files.List(context.Background())` → `FileList`
- `POST /v1/files` — `client.Files.Upload(context.Background(), body)` → `File`
- `GET /v1/files/{path}` — `client.Files.Download(context.Background(), path)`
- `DELETE /v1/files/{path}` — `client.Files.Delete(context.Background(), path)`

## Client methods

- `POST /v1/pdf` — `client.Pdf(context.Background(), body)` → `PdfResponse`
- `POST /v1/scrape` — `client.Scrape(context.Background(), body)` → `ScrapeResponse`
- `POST /v1/screenshot` — `client.Screenshot(context.Background(), body)` → `ScreenshotResponse`

## profiles

- `GET /v1/profiles` — `client.Profiles.List(context.Background(), params)` → `ProfileListResponse`
- `POST /v1/profiles` — `client.Profiles.Create(context.Background(), body)` → `ProfileCreateResponse`
- `GET /v1/profiles/{id}` — `client.Profiles.Get(context.Background(), id, params)` → `ProfileGetResponse`
- `PATCH /v1/profiles/{id}` — `client.Profiles.Update(context.Background(), id, body, params)` → `ProfileUpdateResponse`

## sessions

- `GET /v1/sessions` — `client.Sessions.List(context.Background(), params)` → `SessionList` _(paginated)_
- `POST /v1/sessions` — `client.Sessions.Create(context.Background(), body)` → `Session`
- `POST /v1/sessions/release` — `client.Sessions.ReleaseAll(context.Background(), body, params)` → `SessionReleaseAllResponse`
- `GET /v1/sessions/{id}` — `client.Sessions.Retrieve(context.Background(), id)` → `Session`
- `GET /v1/sessions/{id}/agent-traces` — `client.Sessions.AgentTraces(context.Background(), id, params)` → `SessionAgentTracesResponse`
- `GET /v1/sessions/{id}/context` — `client.Sessions.Context(context.Background(), id)` → `SessionContext`
- `GET /v1/sessions/{id}/cost` — `client.Sessions.Cost(context.Background(), id)` → `SessionCostResponse`
- `GET /v1/sessions/{id}/events` — `client.Sessions.Events(context.Background(), id, params)` → `SessionEventsResponse`
- `GET /v1/sessions/{id}/live-details` — `client.Sessions.LiveDetails(context.Background(), id)` → `SessionLiveDetailsResponse`
- `POST /v1/sessions/{id}/release` — `client.Sessions.Release(context.Background(), id, body)` → `SessionReleaseResponse`
- `POST /v1/sessions/{sessionId}/computer` — `client.Sessions.Computer(context.Background(), sessionID, body)` → `SessionComputerResponse`

## sessions.captchas

- `POST /v1/sessions/{sessionId}/captchas/solve` — `client.Sessions.Captchas.Solve(context.Background(), sessionID, body)` → `CaptchaSolveResponse`
- `POST /v1/sessions/{sessionId}/captchas/solve-image` — `client.Sessions.Captchas.SolveImage(context.Background(), sessionID, body)` → `CaptchaSolveImageResponse`
- `GET /v1/sessions/{sessionId}/captchas/status` — `client.Sessions.Captchas.Status(context.Background(), sessionID)` → `CaptchaStatusResponse`

## sessions.files

- `GET /v1/sessions/{sessionId}/files` — `client.Sessions.Files.List(context.Background(), sessionID)` → `FileList`
- `POST /v1/sessions/{sessionId}/files` — `client.Sessions.Files.Upload(context.Background(), sessionID, body)` → `File`
- `DELETE /v1/sessions/{sessionId}/files` — `client.Sessions.Files.DeleteAll(context.Background(), sessionID)`
- `GET /v1/sessions/{sessionId}/files.zip` — `client.Sessions.Files.DownloadArchive(context.Background(), sessionID)`
- `GET /v1/sessions/{sessionId}/files/{path}` — `client.Sessions.Files.Download(context.Background(), sessionID, path)`
- `DELETE /v1/sessions/{sessionId}/files/{path}` — `client.Sessions.Files.Delete(context.Background(), sessionID, path)`

