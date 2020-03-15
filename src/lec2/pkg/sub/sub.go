// Package sub is demonstration package
// import it as github.com/polis-mail-ru-golang-1/examples/lec2/pkg/sub
package sub

// Sub demonstaits exporting of struct
type Sub struct {
	hidden int
	Shown  int
}

func (s Sub) privateGet() int {
	return s.Shown
}

// PublicGet returns hidden field
func (s Sub) PublicGet() int {
	return s.hidden
}

const (
	hiddenConst = iota
	// ShownConst is example with value 1
	ShownConst
)
