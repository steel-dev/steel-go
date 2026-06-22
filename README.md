# Steel API

Steel API client

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
	result, err := client.Sessions.Create(context.Background(), params)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
```

Authentication uses API key (`steel-api-key`). The example above reads the key from the `STEEL_API_KEY` environment variable.

## Pagination

List endpoints are paginated. Each response exposes the page items together with a cursor for fetching the next page.

## Errors

Every method returns an `error` as its final value; check it before using the result.

## Requirements

- Go 1.21 or later.

## API reference

See [api.md](./api.md) for the full list of resources and methods.

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md).

## License

Released under the MIT license. See [LICENSE](./LICENSE).
