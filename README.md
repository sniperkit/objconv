objconv [![CircleCI](https://circleci.com/gh/segmentio/objconv.svg?style=shield)](https://circleci.com/gh/segmentio/objconv) [![Go Report Card](https://goreportcard.com/badge/github.com/segmentio/objconv)](https://goreportcard.com/report/github.com/segmentio/objconv) [![GoDoc](https://godoc.org/github.com/segmentio/objconv?status.svg)](https://godoc.org/github.com/segmentio/objconv)
=======

A Go package exposing encoder and decoders that support data streaming to and
from multiple formats.

The design of the package is inspired from the standard [image](https://golang.org/pkg/image/)
package, a high-level API is exposed and the program can load specific
implementations by importing subpackages of objconv.  
Each subpackage provides a parser and emitter in a specific format.

Installation
------------

```shell
go get github.com/segmentio/objconv
```

Encoder
-------

The package exposes a generic encoder API that let's the program serialize
native values into various formats.

Here's an example of how to serialize a structure to JSON:
```go
package main

import (
    "os"

    "github.com/segmentio/objconv"
    _ "github.com/segmentio/objconv/json" // load the JSON codec
)

func main() {
    // prints {"hello":"world"}
    //
    objconv.Encode(os.Stdout, "json", struct{
        Hello `objconv:"hello"`
    }{"world"})
}
```

To support multiple serialization formats the program has to import each
subpackage it's interested in, then it can select a different encoder name to
change format of the output.

Decoder
-------

Like the encoders, decoders can be identified by name and generate native Go
values from their serialized representation in different formats.

Here's an example of how to use a decoder:
```go
package main

import (
    "github.com/segmentio/objconv"
    _ "github.com/segmentio/objconv/json" // load the JSON codec
)

func main() {
    s := `{"hello":"world"}`
    v := struct{
       Hello `objconv:"hello"`
    }{}

    objconv.DecodeString(s, "json", &v)

    // v.Hello == "world"
    // ...
}
```

Streaming
---------

Mime Types
----------

The codecs registers themselves under multiple names, including the standard
mime-types associated with the serialization format they implement.  
For example the `objconv/json` package registers its encoder and decoder under
`text/json` and `application/json` on top of the simpler `json` name.  
This makes it easy to load encoders and decoders from an HTTP request's
Content-Type header for example.
