package main

import (
	// "github.com/jbendotnet/pointless/one"
	// "github.com/jbendotnet/pointless/two"
	onev3 "github.com/jbendotnet/pointless/v3/one"
	twov3 "github.com/jbendotnet/pointless/v3/two"
)

func main() {
	onev3.Do(twov3.Get())
	// one.Do(twov3.Get())
}
