// Gomodules is desabled by default in GOPATH. To use gomodules add `GO111MODULE=on` before every go run.
// Or `export GO111MODULE=on` for current terminal session.
// Download all vendors and build with `GO111MODULE=on go build`.
// Make copy of vendors to build without go modules: `GO111MODULE=on go mod vendor`

package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
)

func main() {
	arguments, _ := docopt.ParseDoc("example of usage")
	fmt.Println(arguments)
}
