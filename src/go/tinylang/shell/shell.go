package main

import (
	"flag"
	"log"
	"tiny-compiler/src/go/tinylang/env"
)

// -p C:\projects\tiny-compiler\src\ab\calculator.ab

func main(){
	p := flag.String("p", "", "path to script")
	r := flag.Bool("ir", false, "keep ir files")
	flag.Parse()

	if *p == "" {
		log.Printf(" please define path to file using -p flag ")
		return
	}


	env.Start(*p,*r)
}
