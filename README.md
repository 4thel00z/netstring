# netstring

## Motivation

In case of unreliable `io` you want a transmission envelope which ensures that the payload arrives completely.

## Installation

```bash
go get -u github.com/4thel00z/netstring
```

## Usage

```go
package main

import (
	"github.com/4thel00z/netstring/v1/pkg/netstring"
	"log"
	"strings"
)

func main() {
	ns, err := netstring.FromReader(strings.NewReader("4:Test,"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(ns)
}
```


## License

This project is licensed under the GPL-3 license.
