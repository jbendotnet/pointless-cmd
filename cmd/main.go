package main

import (
	"log"

	// "github.com/jbendotnet/pointless/one"
	// "github.com/jbendotnet/pointless/two"
	"github.com/jbendotnet/pointless/v3/one"
	"github.com/jbendotnet/pointless/v3/two"

	onev4 "github.com/jbendotnet/pointless/v4/one"
	twov4 "github.com/jbendotnet/pointless/v4/two"

	threev2 "github.com/jbendotnet/pointless2/three/v2"
)

func main() {
	one.Do(two.Get())
	onev4.Do(twov4.Get())
	// one.Do(twov3.Get())
	ls := threev2.List()
	log.Println(ls)
}
