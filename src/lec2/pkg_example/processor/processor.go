// Package processor provides generating number multipled on coefficient
package processor

import (
	"math/rand"
)

const randN = 100

// Processor provides int generation
type Processor struct {
	coefficient int
}

// NewProcessor returns Processor instance
func NewProcessor(c int) Processor {
	return Processor{c}
}

// Number generates new number multipled on coefficient
func (p Processor) Number() int {
	return rand.Intn(randN) * p.coefficient
}
