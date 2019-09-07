# Jikan Go

[![GoDoc](https://godoc.org/github.com/darenliang/jikan-go?status.svg)](https://godoc.org/github.com/darenliang/jikan-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/darenliang/jikan-go)](https://goreportcard.com/report/github.com/darenliang/jikan-go)

Go bindings for Jikan.

Since the [documentation](https://godoc.org/github.com/darenliang/jikan-go) is not very comprehensive, please refer to the official [Jikan API documentation](https://jikan.docs.apiary.io) for assistance on certain request parameters.

All json data is put into maps instead of structs, so consulting the official documentation of the REST API will show you the available data fields.

To visualize the response you can printout the map using `fmt.Println()` or you can use the `PrettyPrint()` function bundled in the package.

To install: `go get github.com/darenliang/jikan-go`

To import: `import "github.com/darenliang/jikan-go"` and use as `jikan`

Sample usage:
```go
package main

import (
	"fmt"
	"github.com/darenliang/jikan-go"
)

func main() {
	anime, _ := jikan.GetAnime(jikan.Anime{ID: 1})
	fmt.Println(anime["title"])
}
```
```
Output:

Cowboy Bebop
```
```go
package main

import (
	"fmt"
	"github.com/darenliang/jikan-go"
)

func main() {
	search, _ := jikan.GetSearch(jikan.Search{Type: "anime", Q: "FMAB", OrderBy: "score"})
	firstAnime := search["results"].([]interface{})[0].(map[string]interface{})
	fmt.Printf("Title: %v\nID: %v", firstAnime["title"], firstAnime["mal_id"])
}
```
```
Output:

Title: Fullmetal Alchemist: Brotherhood
ID: 5114
```
### Why use maps over structs for reponses?
Due to the multiple different response structures of the Jikan API, it is not very convenient to call different functions depending on the return type of the struct.

Ideally, using functions like `GetAnimeEpisodes` should get you the Jikan's episode json data. However, the problem is that adding many structures creates a few nuisances that are hard to go unnoticed.

1. You have to rely on the Go Docs for all struct fields since they are different from Jikan's normal key values.
2. You have to know the exact function name to call for a specific struct.
3. Whenever a new data field is added to the Jikan API, the corresponding struct that is supposed to hold the data field will ignore it unless it is updated.

### Why do I have to perform type assertions when I want to use the nested data?
Golang doesn't allow dynamic return types for functions other than using interfaces.

Returning interfaces doesn't allow you to use the interface value's underlying concrete value.

Since all structs returned are in the type `map[string]interface{}`, you can only readily use first level data.

For accessing lower levels of data, you must perform type assertions in advance.

To type assert a slice/array, append `.([]interface{})` before accessing the index.

To type assert a map, append `.(map[string]interface{})` before accessing the index.

### Why do I need to create new structs every time I use the API?
Due to the nature of Golang, it is [not possible to use optional parameters in functions](https://golang.org/doc/faq#overloading).

It starts becoming a real problem when Jikan API's "search" or "user" request is used. There is just no simple way of calling only the request parameters you need directly. Although there are such things as "variadic arguments", it does not provide the ability to only add request parameters you need.

One way around this problem is by using structs. We can take advantage of Golang's zero values, allowing us to not have to specify every single request parameter.

If there is a better way of addressing this problem, please let me know.
