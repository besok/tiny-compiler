package main

import (
	"flag"
	"tiny-compiler/src/go/tinylang/env"
)

// -p C:\projects\tiny-compiler\src\ab\calculator.ab

func main(){
	p := flag.String(string("p"), "", "path to script")
	flag.Parse()
	env.Start(*p)
}
