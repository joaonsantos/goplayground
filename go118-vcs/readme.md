# Go 1.18 Build Info

## Quick Start

Since Go 1.18, Go embeds vcs information by default when building. 

You can do it like this:
```go
package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	// use new runtime/debug package
	bi, _ := debug.ReadBuildInfo()

	// just printing out the contents
	// (implements Stringer so it will look nice)
	fmt.Println(bi)
}
```

For more information check out https://go.dev/doc/go1.18#debug/buildinfo and https://go.dev/doc/go1.18#go-version


## Gotchas

Please not, One gotcha is if you have a main in ./cmd/http/main.go you need to build like this:
```bash
go build ./cmd/http
```

Result
```bash
go	go1.18.1
path	git.nfon.net/~joao.santos/go118buildinfo/cmd/http
mod	git.nfon.net/~joao.santos/go118buildinfo	(devel)	
build	-compiler=gc
build	CGO_ENABLED=1
build	CGO_CFLAGS=
build	CGO_CPPFLAGS=
build	CGO_CXXFLAGS=
build	CGO_LDFLAGS=
build	GOARCH=amd64
build	GOOS=darwin
build	GOAMD64=v1
build	vcs=git
build	vcs.revision=a71d3ad7d3f1833ff772f992e8f912e81098531b
build	vcs.time=2022-03-17T14:39:12Z
build	vcs.modified=true
```

If you send the whole path it wont show any vcs information

```bash
// go build or go run
go build ./cmd/http/main.go
```

Result
```bash
go	go1.18.1
path	command-line-arguments
build	-compiler=gc
build	CGO_ENABLED=1
build	CGO_CFLAGS=
build	CGO_CPPFLAGS=
build	CGO_CXXFLAGS=
build	CGO_LDFLAGS=
build	GOARCH=amd64
build	GOOS=darwin
build	GOAMD64=v1
```
