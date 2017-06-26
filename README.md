# go-unleashed #

go-unleashed is a Go client library for accessing the [Unleashed API](https://apidocs.unleashedsoftware.com/).

**Build Status:** [![Build Status](https://travis-ci.org/zentby/go-unleashed.svg?branch=master)](https://travis-ci.org/zentby/go-unleashed)  

go-unleashed requires Go version 1.4 or greater.

## Usage ##

```go
import "github.com/zentby/go-unleashed/unleashed"
```

Construct a new Unleashed client, then use the various services on the client to
access different parts of the Unleashed API. For example:

```go
client := unleashed.NewClient("API_KEY", "API_SECRET")

// list first page of products in the account
products, _, err := client.Products.List(nil)
```


The services of a client divide the API into logical chunks and correspond to
the structure of the Unleashed API documentation at
https://apidocs.unleashedsoftware.com/.

## Code Style

The code of go-unleashed was inspired by [google/go-github](https://github.com/google/go-github).

## License ##

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE)
file.
