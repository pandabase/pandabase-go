# Pandabase Go

The official Pandabase Go client library.

## Requirements

- Go 1.19 or later

## Installation

Make sure your project is using Go Modules (it will have a `go.mod` file in its root if it already is):

```
go mod init
```

Then, reference `pandabase-go` in a Go program with import:

```go
import (
	"github.com/pandabase/pandabase-go/v1"
)
```

Run any of the normal `go` commands (`build`/`install`/`test`). The Go toolchain will resolve and fetch the `pandabase-go` module automatically.

Alternatively, you can also explicitly go get the package into a project:

```
go get -u github.com/pandabase/pandabase-go/v1
```
