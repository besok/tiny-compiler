package main

import (
	"flag"
	"log"
	"tiny-compiler/src/go/tinylang/env"
)

// -p C:\projects\tiny-compiler\src\ab\calculator.ab

func main(){
	p := flag.String(string("p"), "", "path to script")
	flag.Parse()

	if *p == "" {
		log.Printf(" please define path to file using -p flag ")
		return
	}

	env.Start(*p)
}
