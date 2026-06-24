# Steel API

Steel API client

This library provides convenient, typed access to the Steel API from Go.

The full API of this library can be found in [api.md](./api.md).

## Installation

```sh
go get github.com/steel-dev/steel-go
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/steel-dev/steel-go"
)

func main() {
	client := steel.NewClient(os.Getenv("STEEL_API_KEY"))
	result, err := client.Sessions.Create(context.Background(), steel.SessionCreateParams{})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
```

Request fields are wrapped in `param.Field[T]` so the client can distinguish an unset field from a zero value. Build params with the `steel.F` helper (or the typed shortcuts `steel.String`, `steel.Int`, `steel.Bool`), and `steel.Null[T]()` to send an explicit JSON `null`:

```go
params := steel.SomeParams{
	Name:    steel.F("value"),
	Enabled: steel.F(true),
}
```

Authentication uses API key (`steel-api-key`). The client reads the key from the `STEEL_API_KEY` environment variable when one is not passed explicitly.

## Pagination

List methods are paginated. The SDK fetches further pages automatically as you iterate.

```go
iter := client.Sessions.ListAutoPaging(context.Background(), steel.SessionListParams{})
for iter.Next() {
	item := iter.Current()
	fmt.Println(item)
}
if err := iter.Err(); err != nil {
	panic(err)
}
```

## Handling errors

Methods return an `error` as their final value. API errors can be inspected by asserting to `*APIError`:

```go
result, err := client.SomeResource.SomeMethod(context.Background())
if err != nil {
	var apiErr *steel.APIError
	if errors.As(err, &apiErr) {
		fmt.Println(apiErr.StatusCode)
		fmt.Println(apiErr.RequestID)
	}
	return err
}
```

## Retries

Certain errors are automatically retried twice by default with a short exponential backoff. Connection errors, 408, 409, 429, and 5xx responses are retried.

```go
// Configure the default for all requests:
client := steel.NewClient(apiKey, steel.WithMaxRetries(0)) // default is 2

// Or per-request:
client.SomeResource.SomeMethod(context.Background(), steel.WithRequestMaxRetries(5))
```

## Timeouts

Requests time out after 1 minute by default. Configure this with a timeout option:

```go
client := steel.NewClient(apiKey, steel.WithTimeout(20*time.Second))

// Or per-request:
client.SomeResource.SomeMethod(context.Background(), steel.WithRequestTimeout(5*time.Second))
```

## Per-request options

Pass `RequestOption` values to override behaviour for a single call:

```go
client.SomeResource.SomeMethod(
	context.Background(),
	steel.WithRequestHeader("X-Custom", "value"),
	steel.WithIdempotencyKey("my-key"),
)
```

Use `steel.WithResponseInto(&resp)` to access the raw `*http.Response`.

## Requirements

- Go 1.21 or later.

## API reference

See [api.md](./api.md) for the full list of resources and methods.

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md).

## License

Released under the MIT license. See [LICENSE](./LICENSE).
