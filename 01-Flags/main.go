package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var beers = map[string]string{
	"0214124": "Quilmes",
	"02224":   "Brahma",
}

func main() {
	beerscmd := flag.NewFlagSet("beers", flag.ExitOnError)
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatal("You must specifed a command beers")
		os.Exit(2)
	}

	switch flag.Arg(0) {
	case "beers":
		ID := beerscmd.String("id", "", "find my ID")
		beerscmd.Parse(os.Args[2:])

		if *ID != "" {
			fmt.Println(beers[*ID])
		} else {
			fmt.Println(beers)
		}
	}
}
