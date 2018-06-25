package main

import (
	"github.com/jbendotnet/pointless/one"
	"github.com/jbendotnet/pointless/two"
	// twov2 "github.com/jbendotnet/pointless/two/v2"
)

func main() {
	one.Do(two.Get())
	// one.Do(twov2.Get())
}
