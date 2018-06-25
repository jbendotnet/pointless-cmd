package main

import (
	// "github.com/jbendotnet/pointless/one"
	// "github.com/jbendotnet/pointless/two"
	"github.com/jbendotnet/pointless/v3/one"
	"github.com/jbendotnet/pointless/v3/two"

	onev4 "github.com/jbendotnet/pointless/v4/one"
	twov4 "github.com/jbendotnet/pointless/v4/two"
)

func main() {
	one.Do(two.Get())
	onev4.Do(twov4.Get())
	// one.Do(twov3.Get())
}
