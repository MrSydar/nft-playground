package main

import "github.com/tjarratt/babble"

func initBabbler() babble.Babbler {
	babbler := babble.NewBabbler()
	babbler.Count = 1
	return babbler
}

var babbler babble.Babbler = initBabbler()
