# Go 1.18+ Build Info

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

## Example

First build your package:
```bash
go build ./cmd/http
```

The build infomation available looks like this:
```sh
go	go1.23.0
path	example/cmd/http
mod	example	(devel)	
build	-buildmode=exe
build	-compiler=gc
build	DefaultGODEBUG=asmnctimerchan=1,gotypesalias=0,httplaxcontentlength=1,httpmuxgo121=1,httpservecontentkeepheaders=1,netedns0=0,panicnil=1,tls10server=1,tls3des=1,tlskyber=0,tlsrsakex=1,tlsunsafeekm=1,winreadlinkvolume=0,winsymlink=0,x509keypairleaf=0,x509negativeserial=1
build	CGO_ENABLED=1
build	CGO_CFLAGS=
build	CGO_CPPFLAGS=
build	CGO_CXXFLAGS=
build	CGO_LDFLAGS=
build	GOARCH=amd64
build	GOOS=linux
build	GOAMD64=v1
build	vcs=git
build	vcs.revision=27ec75a263361f2841449eb94b95df51e87fe6c7
build	vcs.time=2024-08-18T18:38:40Z
build	vcs.modified=true
```
