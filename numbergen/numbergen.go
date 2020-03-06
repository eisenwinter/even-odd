package numbergen

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type defaultGen struct {
	Value int64
}

//NumberGen creates a even or a odd number
type NumberGen interface {
	Even() int64
	Odd() int64
}

func (n *defaultGen) Even() int64 {
	return (int64(rand.Int31()) * 2)
}

func (n *defaultGen) Odd() int64 {
	return n.Even() - 1
}

//CreateNumberGen creates a new NumberGen
func CreateNumberGen() NumberGen {
	return &defaultGen{}
}
