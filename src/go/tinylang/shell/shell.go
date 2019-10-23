package main

import (
	"flag"
	"io/ioutil"
	"log"
	"tiny-compiler/src/go/tinylang/env"
)

// -p C:\projects\tiny-compiler\src\ab\calculator.ab

func main() {

	p := flag.String("p", "", "path to script")
	r := flag.Bool("ir", false, "keep ir files")
	tr := flag.Bool("trace", false, "show internal logs")
	flag.Parse()

	if !*tr {
		log.SetOutput(ioutil.Discard)
	}

	if *p == "" {
		log.Printf(" please define path to file using -p flag ")
		return
	}

	env.Start(*p, *r)
}
