package main

import (
	"encoding/gob"
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/kyleconroy/archive/instagram"
	"github.com/kyleconroy/archive/twitter"
)

func main() {
	flag.Parse()

	var a interface{}
	var err error

	switch flag.Arg(0) {
	case "twitter":
		a, err = twitter.ParsePath("twitter/testdata/archive")
	case "instagram":
		a, err = instagram.ParsePath("instagram/testdata/archive")
	default:
		log.Fatalf("unknown archive type: %s", flag.Arg(0))
	}

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(filepath.Join(flag.Arg(0), "testdata", "golden.bin"))
	if err != nil {
		log.Fatal(err)
	}

	enc := gob.NewEncoder(f)
	if err := enc.Encode(a); err != nil {
		log.Fatal(err)
	}
}
