package main

import (
	"log"

	"github.com/jbendotnet/pointless/v3/one"
	"github.com/jbendotnet/pointless/v3/two"

	onev4 "github.com/jbendotnet/pointless/v4/one"
	twov4 "github.com/jbendotnet/pointless/v4/two"

	three "github.com/jbendotnet/pointless2/three/v2"
)

func main() {
	one.Do(twov4.Get())
	onev4.Do(two.Get())
	log.Println(three.List())
}
