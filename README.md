# Jikan-go

[![Go Reference](https://pkg.go.dev/badge/github.com/djeddi-yacine/jikan-go.svg)](https://pkg.go.dev/github.com/djeddi-yacine/jikan-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/djeddi-yacine/jikan-go)](https://goreportcard.com/report/github.com/djeddi-yacine/jikan-go)

A Go wrapper for the [Jikan v4 API](https://jikan.moe/).

Documentation can be found
at [go.dev](https://pkg.go.dev/github.com/djeddi-yacine/jikan-go).

### Installation

To install/update: `go install github.com/djeddi-yacine/jikan-go@latest`

To import: `import "github.com/djeddi-yacine/jikan-go"` and use as `jikan`

### Usage

##### Basic Example

```go
package main

import (
	"fmt"
	"github.com/djeddi-yacine/jikan-go"
)

func main() {
	// Get anime
	anime, err := jikan.GetAnimeById(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(anime.Data.Title)
}
```

```
Output:

Cowboy Bebop
```

##### Search Query Example

Please refer to the
official [Jikan API documentation](https://docs.api.jikan.moe/) for
clarification on query parameters.

```go
package main

import (
	"fmt"
	"github.com/djeddi-yacine/jikan-go"
	"net/url"
)

func main() {
	// Setup query
	query := url.Values{}
	query.Set("q", "Cowboy Bebop")
	query.Set("type", "tv")

	// Search anime
	search, err := jikan.GetAnimeSearch(query)
	if err != nil {
		panic(err)
	}
	fmt.Println(search.Data[0].Score)
}
```

```
Output:

8.79
```

##### Troubleshooting

If it is necessary to modify the http client (eg. modify timeout), you can
access the client via `jikan.Client`.

Example: setting the client timeout from 60 seconds (default) to 10 seconds.

```go
jikan.Client.Timeout = time.Second * 10
```
