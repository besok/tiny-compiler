package parsing

import parser "tiny-compiler/src/antlr"

type PrintTinyListener struct {
	*parser.BaseTinyLangListener
}

func NewPrintListener() *PrintTinyListener {
	return new(PrintTinyListener)
}


