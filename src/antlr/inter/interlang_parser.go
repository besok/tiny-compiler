// Code generated from C:/projects/tiny-compiler/src/antlr\InterLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // InterLang

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 31, 214,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 7, 2, 52, 10, 2, 12, 2, 14, 2, 55, 11,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 63, 10, 3, 12, 3, 14, 3, 66,
	11, 3, 3, 3, 3, 3, 7, 3, 70, 10, 3, 12, 3, 14, 3, 73, 11, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 86, 10, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 5, 6, 102, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 3, 9, 7, 9, 114, 10, 9, 12, 9, 14, 9, 117, 11, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 5, 19, 185, 10, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 24, 5, 24, 207, 10, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3,
	25, 3, 25, 2, 2, 26, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 2, 6, 4, 2, 15, 16, 22, 23, 4,
	2, 27, 27, 31, 31, 3, 2, 25, 26, 3, 2, 12, 14, 2, 209, 2, 53, 3, 2, 2,
	2, 4, 58, 3, 2, 2, 2, 6, 74, 3, 2, 2, 2, 8, 80, 3, 2, 2, 2, 10, 101, 3,
	2, 2, 2, 12, 103, 3, 2, 2, 2, 14, 109, 3, 2, 2, 2, 16, 115, 3, 2, 2, 2,
	18, 124, 3, 2, 2, 2, 20, 129, 3, 2, 2, 2, 22, 134, 3, 2, 2, 2, 24, 139,
	3, 2, 2, 2, 26, 144, 3, 2, 2, 2, 28, 150, 3, 2, 2, 2, 30, 157, 3, 2, 2,
	2, 32, 164, 3, 2, 2, 2, 34, 171, 3, 2, 2, 2, 36, 179, 3, 2, 2, 2, 38, 190,
	3, 2, 2, 2, 40, 194, 3, 2, 2, 2, 42, 197, 3, 2, 2, 2, 44, 200, 3, 2, 2,
	2, 46, 206, 3, 2, 2, 2, 48, 210, 3, 2, 2, 2, 50, 52, 5, 4, 3, 2, 51, 50,
	3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2,
	54, 56, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 56, 57, 7, 2, 2, 3, 57, 3, 3, 2,
	2, 2, 58, 59, 5, 48, 25, 2, 59, 60, 7, 6, 2, 2, 60, 64, 7, 31, 2, 2, 61,
	63, 5, 6, 4, 2, 62, 61, 3, 2, 2, 2, 63, 66, 3, 2, 2, 2, 64, 62, 3, 2, 2,
	2, 64, 65, 3, 2, 2, 2, 65, 67, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 67, 71,
	5, 8, 5, 2, 68, 70, 5, 10, 6, 2, 69, 68, 3, 2, 2, 2, 70, 73, 3, 2, 2, 2,
	71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 5, 3, 2, 2, 2, 73, 71, 3, 2,
	2, 2, 74, 75, 5, 48, 25, 2, 75, 76, 7, 7, 2, 2, 76, 77, 7, 31, 2, 2, 77,
	78, 7, 8, 2, 2, 78, 79, 5, 46, 24, 2, 79, 7, 3, 2, 2, 2, 80, 81, 5, 48,
	25, 2, 81, 82, 7, 9, 2, 2, 82, 85, 7, 8, 2, 2, 83, 86, 5, 46, 24, 2, 84,
	86, 7, 11, 2, 2, 85, 83, 3, 2, 2, 2, 85, 84, 3, 2, 2, 2, 86, 9, 3, 2, 2,
	2, 87, 102, 5, 16, 9, 2, 88, 102, 5, 18, 10, 2, 89, 102, 5, 22, 12, 2,
	90, 102, 5, 24, 13, 2, 91, 102, 5, 26, 14, 2, 92, 102, 5, 28, 15, 2, 93,
	102, 5, 30, 16, 2, 94, 102, 5, 32, 17, 2, 95, 102, 5, 34, 18, 2, 96, 102,
	5, 36, 19, 2, 97, 102, 5, 42, 22, 2, 98, 102, 5, 44, 23, 2, 99, 102, 5,
	20, 11, 2, 100, 102, 5, 38, 20, 2, 101, 87, 3, 2, 2, 2, 101, 88, 3, 2,
	2, 2, 101, 89, 3, 2, 2, 2, 101, 90, 3, 2, 2, 2, 101, 91, 3, 2, 2, 2, 101,
	92, 3, 2, 2, 2, 101, 93, 3, 2, 2, 2, 101, 94, 3, 2, 2, 2, 101, 95, 3, 2,
	2, 2, 101, 96, 3, 2, 2, 2, 101, 97, 3, 2, 2, 2, 101, 98, 3, 2, 2, 2, 101,
	99, 3, 2, 2, 2, 101, 100, 3, 2, 2, 2, 102, 11, 3, 2, 2, 2, 103, 104, 5,
	48, 25, 2, 104, 105, 5, 14, 8, 2, 105, 106, 7, 8, 2, 2, 106, 107, 7, 18,
	2, 2, 107, 108, 9, 2, 2, 2, 108, 13, 3, 2, 2, 2, 109, 110, 7, 29, 2, 2,
	110, 111, 7, 22, 2, 2, 111, 15, 3, 2, 2, 2, 112, 114, 5, 12, 7, 2, 113,
	112, 3, 2, 2, 2, 114, 117, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115, 116,
	3, 2, 2, 2, 116, 118, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 118, 119, 5, 48,
	25, 2, 119, 120, 7, 31, 2, 2, 120, 121, 7, 8, 2, 2, 121, 122, 7, 17, 2,
	2, 122, 123, 5, 46, 24, 2, 123, 17, 3, 2, 2, 2, 124, 125, 5, 48, 25, 2,
	125, 126, 7, 31, 2, 2, 126, 127, 7, 8, 2, 2, 127, 128, 5, 14, 8, 2, 128,
	19, 3, 2, 2, 2, 129, 130, 5, 48, 25, 2, 130, 131, 5, 14, 8, 2, 131, 132,
	7, 8, 2, 2, 132, 133, 5, 14, 8, 2, 133, 21, 3, 2, 2, 2, 134, 135, 5, 48,
	25, 2, 135, 136, 5, 14, 8, 2, 136, 137, 7, 8, 2, 2, 137, 138, 7, 31, 2,
	2, 138, 23, 3, 2, 2, 2, 139, 140, 5, 48, 25, 2, 140, 141, 5, 14, 8, 2,
	141, 142, 7, 8, 2, 2, 142, 143, 9, 2, 2, 2, 143, 25, 3, 2, 2, 2, 144, 145,
	5, 48, 25, 2, 145, 146, 5, 14, 8, 2, 146, 147, 7, 8, 2, 2, 147, 148, 7,
	30, 2, 2, 148, 149, 5, 14, 8, 2, 149, 27, 3, 2, 2, 2, 150, 151, 5, 48,
	25, 2, 151, 152, 5, 14, 8, 2, 152, 153, 7, 8, 2, 2, 153, 154, 7, 19, 2,
	2, 154, 155, 9, 3, 2, 2, 155, 156, 7, 22, 2, 2, 156, 29, 3, 2, 2, 2, 157,
	158, 5, 48, 25, 2, 158, 159, 5, 14, 8, 2, 159, 160, 7, 8, 2, 2, 160, 161,
	5, 14, 8, 2, 161, 162, 9, 4, 2, 2, 162, 163, 5, 14, 8, 2, 163, 31, 3, 2,
	2, 2, 164, 165, 5, 48, 25, 2, 165, 166, 5, 14, 8, 2, 166, 167, 7, 8, 2,
	2, 167, 168, 5, 14, 8, 2, 168, 169, 7, 24, 2, 2, 169, 170, 5, 14, 8, 2,
	170, 33, 3, 2, 2, 2, 171, 172, 5, 48, 25, 2, 172, 173, 5, 14, 8, 2, 173,
	174, 7, 8, 2, 2, 174, 175, 7, 31, 2, 2, 175, 176, 7, 3, 2, 2, 176, 177,
	5, 14, 8, 2, 177, 178, 7, 28, 2, 2, 178, 35, 3, 2, 2, 2, 179, 180, 5, 48,
	25, 2, 180, 181, 7, 31, 2, 2, 181, 184, 7, 3, 2, 2, 182, 185, 7, 22, 2,
	2, 183, 185, 5, 14, 8, 2, 184, 182, 3, 2, 2, 2, 184, 183, 3, 2, 2, 2, 185,
	186, 3, 2, 2, 2, 186, 187, 7, 28, 2, 2, 187, 188, 7, 8, 2, 2, 188, 189,
	5, 14, 8, 2, 189, 37, 3, 2, 2, 2, 190, 191, 5, 48, 25, 2, 191, 192, 7,
	10, 2, 2, 192, 193, 5, 14, 8, 2, 193, 39, 3, 2, 2, 2, 194, 195, 7, 21,
	2, 2, 195, 196, 7, 22, 2, 2, 196, 41, 3, 2, 2, 2, 197, 198, 5, 48, 25,
	2, 198, 199, 5, 40, 21, 2, 199, 43, 3, 2, 2, 2, 200, 201, 5, 48, 25, 2,
	201, 202, 7, 20, 2, 2, 202, 203, 5, 14, 8, 2, 203, 204, 5, 40, 21, 2, 204,
	45, 3, 2, 2, 2, 205, 207, 7, 28, 2, 2, 206, 205, 3, 2, 2, 2, 206, 207,
	3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 209, 9, 5, 2, 2, 209, 47, 3, 2,
	2, 2, 210, 211, 7, 22, 2, 2, 211, 212, 7, 4, 2, 2, 212, 49, 3, 2, 2, 2,
	10, 53, 64, 71, 85, 101, 115, 184, 206,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'['", "':'", "", "'function'", "'argument'", "'='", "'return_type'",
	"'return'", "'void'", "'string'", "'bool'", "'num'", "'true'", "'false'",
	"'init'", "'init_arr'", "'call'", "'ifFalse'", "'goto'", "", "", "", "",
	"", "", "']'", "", "'param'",
}
var symbolicNames = []string{
	"", "", "", "WS", "FUN", "ARG", "EQ", "RET_TP", "RET", "VOID", "STR", "BOOL",
	"NUM", "TRUE", "FALSE", "INIT", "INIT_ARR", "CALL", "IFFALSE", "GOTO",
	"NUMBER", "STRING_RAW", "NUM_SIGN", "BOOL_SIGN", "BOOL_PL", "SYS_FUNC",
	"ARRAY", "INNER_VAR", "PARAM", "ITEM",
}

var ruleNames = []string{
	"file", "function", "arg", "retTp", "statement", "internalArrArg", "internalVar",
	"newVar", "updVar", "initIntVar", "initItem", "initPrim", "initParam",
	"initCall", "initBoolOp", "initNumOp", "initArrEl", "extArrEl", "returnTp",
	"gotoIn", "gotoTp", "ifFalse", "typeTp", "line",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type InterLangParser struct {
	*antlr.BaseParser
}

func NewInterLangParser(input antlr.TokenStream) *InterLangParser {
	this := new(InterLangParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "InterLang.g4"

	return this
}

// InterLangParser tokens.
const (
	InterLangParserEOF        = antlr.TokenEOF
	InterLangParserT__0       = 1
	InterLangParserT__1       = 2
	InterLangParserWS         = 3
	InterLangParserFUN        = 4
	InterLangParserARG        = 5
	InterLangParserEQ         = 6
	InterLangParserRET_TP     = 7
	InterLangParserRET        = 8
	InterLangParserVOID       = 9
	InterLangParserSTR        = 10
	InterLangParserBOOL       = 11
	InterLangParserNUM        = 12
	InterLangParserTRUE       = 13
	InterLangParserFALSE      = 14
	InterLangParserINIT       = 15
	InterLangParserINIT_ARR   = 16
	InterLangParserCALL       = 17
	InterLangParserIFFALSE    = 18
	InterLangParserGOTO       = 19
	InterLangParserNUMBER     = 20
	InterLangParserSTRING_RAW = 21
	InterLangParserNUM_SIGN   = 22
	InterLangParserBOOL_SIGN  = 23
	InterLangParserBOOL_PL    = 24
	InterLangParserSYS_FUNC   = 25
	InterLangParserARRAY      = 26
	InterLangParserINNER_VAR  = 27
	InterLangParserPARAM      = 28
	InterLangParserITEM       = 29
)

// InterLangParser rules.
const (
	InterLangParserRULE_file           = 0
	InterLangParserRULE_function       = 1
	InterLangParserRULE_arg            = 2
	InterLangParserRULE_retTp          = 3
	InterLangParserRULE_statement      = 4
	InterLangParserRULE_internalArrArg = 5
	InterLangParserRULE_internalVar    = 6
	InterLangParserRULE_newVar         = 7
	InterLangParserRULE_updVar         = 8
	InterLangParserRULE_initIntVar     = 9
	InterLangParserRULE_initItem       = 10
	InterLangParserRULE_initPrim       = 11
	InterLangParserRULE_initParam      = 12
	InterLangParserRULE_initCall       = 13
	InterLangParserRULE_initBoolOp     = 14
	InterLangParserRULE_initNumOp      = 15
	InterLangParserRULE_initArrEl      = 16
	InterLangParserRULE_extArrEl       = 17
	InterLangParserRULE_returnTp       = 18
	InterLangParserRULE_gotoIn         = 19
	InterLangParserRULE_gotoTp         = 20
	InterLangParserRULE_ifFalse        = 21
	InterLangParserRULE_typeTp         = 22
	InterLangParserRULE_line           = 23
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(InterLangParserEOF, 0)
}

func (s *FileContext) AllFunction() []IFunctionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionContext)(nil)).Elem())
	var tst = make([]IFunctionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionContext)
		}
	}

	return tst
}

func (s *FileContext) Function(i int) IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *InterLangParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, InterLangParserRULE_file)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == InterLangParserNUMBER {
		{
			p.SetState(48)
			p.Function()
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(54)
		p.Match(InterLangParserEOF)
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *FunctionContext) FUN() antlr.TerminalNode {
	return s.GetToken(InterLangParserFUN, 0)
}

func (s *FunctionContext) ITEM() antlr.TerminalNode {
	return s.GetToken(InterLangParserITEM, 0)
}

func (s *FunctionContext) RetTp() IRetTpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRetTpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRetTpContext)
}

func (s *FunctionContext) AllArg() []IArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgContext)(nil)).Elem())
	var tst = make([]IArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgContext)
		}
	}

	return tst
}

func (s *FunctionContext) Arg(i int) IArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *FunctionContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *FunctionContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *InterLangParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, InterLangParserRULE_function)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Line()
	}
	{
		p.SetState(57)
		p.Match(InterLangParserFUN)
	}
	{
		p.SetState(58)
		p.Match(InterLangParserITEM)
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(59)
				p.Arg()
			}

		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	{
		p.SetState(65)
		p.RetTp()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(66)
				p.Statement()
			}

		}
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_arg
	return p
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ArgContext) ARG() antlr.TerminalNode {
	return s.GetToken(InterLangParserARG, 0)
}

func (s *ArgContext) ITEM() antlr.TerminalNode {
	return s.GetToken(InterLangParserITEM, 0)
}

func (s *ArgContext) EQ() antlr.TerminalNode {
	return s.GetToken(InterLangParserEQ, 0)
}

func (s *ArgContext) TypeTp() ITypeTpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTpContext)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterArg(s)
	}
}

func (s *ArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitArg(s)
	}
}

func (p *InterLangParser) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, InterLangParserRULE_arg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Line()
	}
	{
		p.SetState(73)
		p.Match(InterLangParserARG)
	}
	{
		p.SetState(74)
		p.Match(InterLangParserITEM)
	}
	{
		p.SetState(75)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(76)
		p.TypeTp()
	}

	return localctx
}

// IRetTpContext is an interface to support dynamic dispatch.
type IRetTpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRetTpContext differentiates from other interfaces.
	IsRetTpContext()
}

type RetTpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetTpContext() *RetTpContext {
	var p = new(RetTpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_retTp
	return p
}

func (*RetTpContext) IsRetTpContext() {}

func NewRetTpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetTpContext {
	var p = new(RetTpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_retTp

	return p
}

func (s *RetTpContext) GetParser() antlr.Parser { return s.parser }

func (s *RetTpContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *RetTpContext) RET_TP() antlr.TerminalNode {
	return s.GetToken(InterLangParserRET_TP, 0)
}

func (s *RetTpContext) EQ() antlr.TerminalNode {
	return s.GetToken(InterLangParserEQ, 0)
}

func (s *RetTpContext) TypeTp() ITypeTpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTpContext)
}

func (s *RetTpContext) VOID() antlr.TerminalNode {
	return s.GetToken(InterLangParserVOID, 0)
}

func (s *RetTpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetTpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetTpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterRetTp(s)
	}
}

func (s *RetTpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitRetTp(s)
	}
}

func (p *InterLangParser) RetTp() (localctx IRetTpContext) {
	localctx = NewRetTpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, InterLangParserRULE_retTp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Line()
	}
	{
		p.SetState(79)
		p.Match(InterLangParserRET_TP)
	}
	{
		p.SetState(80)
		p.Match(InterLangParserEQ)
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case InterLangParserSTR, InterLangParserBOOL, InterLangParserNUM, InterLangParserARRAY:
		{
			p.SetState(81)
			p.TypeTp()
		}

	case InterLangParserVOID:
		{
			p.SetState(82)
			p.Match(InterLangParserVOID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) NewVar() INewVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewVarContext)
}

func (s *StatementContext) UpdVar() IUpdVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdVarContext)
}

func (s *StatementContext) InitItem() IInitItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitItemContext)
}

func (s *StatementContext) InitPrim() IInitPrimContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitPrimContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitPrimContext)
}

func (s *StatementContext) InitParam() IInitParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitParamContext)
}

func (s *StatementContext) InitCall() IInitCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitCallContext)
}

func (s *StatementContext) InitBoolOp() IInitBoolOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitBoolOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitBoolOpContext)
}

func (s *StatementContext) InitNumOp() IInitNumOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitNumOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitNumOpContext)
}

func (s *StatementContext) InitArrEl() IInitArrElContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitArrElContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitArrElContext)
}

func (s *StatementContext) ExtArrEl() IExtArrElContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtArrElContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtArrElContext)
}

func (s *StatementContext) GotoTp() IGotoTpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGotoTpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGotoTpContext)
}

func (s *StatementContext) IfFalse() IIfFalseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfFalseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfFalseContext)
}

func (s *StatementContext) InitIntVar() IInitIntVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitIntVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitIntVarContext)
}

func (s *StatementContext) ReturnTp() IReturnTpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnTpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnTpContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *InterLangParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, InterLangParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
			p.NewVar()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)
			p.UpdVar()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(87)
			p.InitItem()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(88)
			p.InitPrim()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(89)
			p.InitParam()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(90)
			p.InitCall()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(91)
			p.InitBoolOp()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(92)
			p.InitNumOp()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(93)
			p.InitArrEl()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(94)
			p.ExtArrEl()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(95)
			p.GotoTp()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(96)
			p.IfFalse()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(97)
			p.InitIntVar()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(98)
			p.ReturnTp()
		}

	}

	return localctx
}

// IInternalArrArgContext is an interface to support dynamic dispatch.
type IInternalArrArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInternalArrArgContext differentiates from other interfaces.
	IsInternalArrArgContext()
}

type InternalArrArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInternalArrArgContext() *InternalArrArgContext {
	var p = new(InternalArrArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_internalArrArg
	return p
}

func (*InternalArrArgContext) IsInternalArrArgContext() {}

func NewInternalArrArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InternalArrArgContext {
	var p = new(InternalArrArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_internalArrArg

	return p
}

func (s *InternalArrArgContext) GetParser() antlr.Parser { return s.parser }

func (s *InternalArrArgContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *InternalArrArgContext) InternalVar() IInternalVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInternalVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInternalVarContext)
}

func (s *InternalArrArgContext) EQ() antlr.TerminalNode {
	return s.GetToken(InterLangParserEQ, 0)
}

func (s *InternalArrArgContext) INIT_ARR() antlr.TerminalNode {
	return s.GetToken(InterLangParserINIT_ARR, 0)
}

func (s *InternalArrArgContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(InterLangParserNUMBER, 0)
}

func (s *InternalArrArgContext) STRING_RAW() antlr.TerminalNode {
	return s.GetToken(InterLangParserSTRING_RAW, 0)
}

func (s *InternalArrArgContext) TRUE() antlr.TerminalNode {
	return s.GetToken(InterLangParserTRUE, 0)
}

func (s *InternalArrArgContext) FALSE() antlr.TerminalNode {
	return s.GetToken(InterLangParserFALSE, 0)
}

func (s *InternalArrArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InternalArrArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InternalArrArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterInternalArrArg(s)
	}
}

func (s *InternalArrArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitInternalArrArg(s)
	}
}

func (p *InterLangParser) InternalArrArg() (localctx IInternalArrArgContext) {
	localctx = NewInternalArrArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, InterLangParserRULE_internalArrArg)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Line()
	}
	{
		p.SetState(102)
		p.InternalVar()
	}
	{
		p.SetState(103)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(104)
		p.Match(InterLangParserINIT_ARR)
	}
	{
		p.SetState(105)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<InterLangParserTRUE)|(1<<InterLangParserFALSE)|(1<<InterLangParserNUMBER)|(1<<InterLangParserSTRING_RAW))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IInternalVarContext is an interface to support dynamic dispatch.
type IInternalVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInternalVarContext differentiates from other interfaces.
	IsInternalVarContext()
}

type InternalVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInternalVarContext() *InternalVarContext {
	var p = new(InternalVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_internalVar
	return p
}

func (*InternalVarContext) IsInternalVarContext() {}

func NewInternalVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InternalVarContext {
	var p = new(InternalVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_internalVar

	return p
}

func (s *InternalVarContext) GetParser() antlr.Parser { return s.parser }

func (s *InternalVarContext) INNER_VAR() antlr.TerminalNode {
	return s.GetToken(InterLangParserINNER_VAR, 0)
}

func (s *InternalVarContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(InterLangParserNUMBER, 0)
}

func (s *InternalVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InternalVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InternalVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterInternalVar(s)
	}
}

func (s *InternalVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitInternalVar(s)
	}
}

func (p *InterLangParser) InternalVar() (localctx IInternalVarContext) {
	localctx = NewInternalVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, InterLangParserRULE_internalVar)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(InterLangParserINNER_VAR)
	}
	{
		p.SetState(108)
		p.Match(InterLangParserNUMBER)
	}

	return localctx
}

// INewVarContext is an interface to support dynamic dispatch.
type INewVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNewVarContext differentiates from other interfaces.
	IsNewVarContext()
}

type NewVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNewVarContext() *NewVarContext {
	var p = new(NewVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_newVar
	return p
}

func (*NewVarContext) IsNewVarContext() {}

func NewNewVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NewVarContext {
	var p = new(NewVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_newVar

	return p
}

func (s *NewVarContext) GetParser() antlr.Parser { return s.parser }

func (s *NewVarContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *NewVarContext) ITEM() antlr.TerminalNode {
	return s.GetToken(InterLangParserITEM, 0)
}

func (s *NewVarContext) EQ() antlr.TerminalNode {
	return s.GetToken(InterLangParserEQ, 0)
}

func (s *NewVarContext) INIT() antlr.TerminalNode {
	return s.GetToken(InterLangParserINIT, 0)
}

func (s *NewVarContext) TypeTp() ITypeTpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTpContext)
}

func (s *NewVarContext) AllInternalArrArg() []IInternalArrArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInternalArrArgContext)(nil)).Elem())
	var tst = make([]IInternalArrArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInternalArrArgContext)
		}
	}

	return tst
}

func (s *NewVarContext) InternalArrArg(i int) IInternalArrArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInternalArrArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInternalArrArgContext)
}

func (s *NewVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NewVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterNewVar(s)
	}
}

func (s *NewVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitNewVar(s)
	}
}

func (p *InterLangParser) NewVar() (localctx INewVarContext) {
	localctx = NewNewVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, InterLangParserRULE_newVar)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(110)
				p.InternalArrArg()
			}

		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}
	{
		p.SetState(116)
		p.Line()
	}
	{
		p.SetState(117)
		p.Match(InterLangParserITEM)
	}
	{
		p.SetState(118)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(119)
		p.Match(InterLangParserINIT)
	}
	{
		p.SetState(120)
		p.TypeTp()
	}

	return localctx
}

// IUpdVarContext is an interface to support dynamic dispatch.
type IUpdVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdVarContext differentiates from other interfaces.
	IsUpdVarContext()
}

type UpdVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdVarContext() *UpdVarContext {
	var p = new(UpdVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_updVar
	return p
}

func (*UpdVarContext) IsUpdVarContext() {}

func NewUpdVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdVarContext {
	var p = new(UpdVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_updVar

	return p
}

func (s *UpdVarContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdVarContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *UpdVarContext) ITEM() antlr.TerminalNode {
	return s.GetToken(InterLangParserITEM, 0)
}

func (s *UpdVarContext) EQ() antlr.TerminalNode {
	return s.GetToken(InterLangParserEQ, 0)
}

func (s *UpdVarContext) InternalVar() IInternalVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInternalVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInternalVarContext)
}

func (s *UpdVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpdVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterUpdVar(s)
	}
}

func (s *UpdVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitUpdVar(s)
	}
}

func (p *InterLangParser) UpdVar() (localctx IUpdVarContext) {
	localctx = NewUpdVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, InterLangParserRULE_updVar)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Line()
	}
	{
		p.SetState(123)
		p.Match(InterLangParserITEM)
	}
	{
		p.SetState(124)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(125)
		p.InternalVar()
	}

	return localctx
}

// IInitIntVarContext is an interface to support dynamic dispatch.
type IInitIntVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitIntVarContext differentiates from other interfaces.
	IsInitIntVarContext()
}

type InitIntVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitIntVarContext() *InitIntVarContext {
	var p = new(InitIntVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_initIntVar
	return p
}

func (*InitIntVarContext) IsInitIntVarContext() {}

func NewInitIntVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitIntVarContext {
	var p = new(InitIntVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_initIntVar

	return p
}

func (s *InitIntVarContext) GetParser() antlr.Parser { return s.parser }

func (s *InitIntVarContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *InitIntVarContext) AllInternalVar() []IInternalVarContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInternalVarContext)(nil)).Elem())
	var tst = make([]IInternalVarContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInternalVarContext)
		}
	}

	return tst
}

func (s *InitIntVarContext) InternalVar(i int) IInternalVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInternalVarContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInternalVarContext)
}

func (s *InitIntVarContext) EQ() antlr.TerminalNode {
	return s.GetToken(InterLangParserEQ, 0)
}

func (s *InitIntVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitIntVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitIntVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterInitIntVar(s)
	}
}

func (s *InitIntVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitInitIntVar(s)
	}
}

func (p *InterLangParser) InitIntVar() (localctx IInitIntVarContext) {
	localctx = NewInitIntVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, InterLangParserRULE_initIntVar)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Line()
	}
	{
		p.SetState(128)
		p.InternalVar()
	}
	{
		p.SetState(129)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(130)
		p.InternalVar()
	}

	return localctx
}

// IInitItemContext is an interface to support dynamic dispatch.
type IInitItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitItemContext differentiates from other interfaces.
	IsInitItemContext()
}

type InitItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitItemContext() *InitItemContext {
	var p = new(InitItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_initItem
	return p
}

func (*InitItemContext) IsInitItemContext() {}

func NewInitItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitItemContext {
	var p = new(InitItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_initItem

	return p
}

func (s *InitItemContext) GetParser() antlr.Parser { return s.parser }

func (s *InitItemContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *InitItemContext) InternalVar() IInternalVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInternalVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInternalVarContext)
}

func (s *InitItemContext) EQ() antlr.TerminalNode {
	return s.GetToken(InterLangParserEQ, 0)
}

func (s *InitItemContext) ITEM() antlr.TerminalNode {
	return s.GetToken(InterLangParserITEM, 0)
}

func (s *InitItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterInitItem(s)
	}
}

func (s *InitItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitInitItem(s)
	}
}

func (p *InterLangParser) InitItem() (localctx IInitItemContext) {
	localctx = NewInitItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, InterLangParserRULE_initItem)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Line()
	}
	{
		p.SetState(133)
		p.InternalVar()
	}
	{
		p.SetState(134)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(135)
		p.Match(InterLangParserITEM)
	}

	return localctx
}

// IInitPrimContext is an interface to support dynamic dispatch.
type IInitPrimContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitPrimContext differentiates from other interfaces.
	IsInitPrimContext()
}

type InitPrimContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitPrimContext() *InitPrimContext {
	var p = new(InitPrimContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_initPrim
	return p
}

func (*InitPrimContext) IsInitPrimContext() {}

func NewInitPrimContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitPrimContext {
	var p = new(InitPrimContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_initPrim

	return p
}

func (s *InitPrimContext) GetParser() antlr.Parser { return s.parser }

func (s *InitPrimContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *InitPrimContext) InternalVar() IInternalVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInternalVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInternalVarContext)
}

func (s *InitPrimContext) EQ() antlr.TerminalNode {
	return s.GetToken(InterLangParserEQ, 0)
}

func (s *InitPrimContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(InterLangParserNUMBER, 0)
}

func (s *InitPrimContext) TRUE() antlr.TerminalNode {
	return s.GetToken(InterLangParserTRUE, 0)
}

func (s *InitPrimContext) FALSE() antlr.TerminalNode {
	return s.GetToken(InterLangParserFALSE, 0)
}

func (s *InitPrimContext) STRING_RAW() antlr.TerminalNode {
	return s.GetToken(InterLangParserSTRING_RAW, 0)
}

func (s *InitPrimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitPrimContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitPrimContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterInitPrim(s)
	}
}

func (s *InitPrimContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitInitPrim(s)
	}
}

func (p *InterLangParser) InitPrim() (localctx IInitPrimContext) {
	localctx = NewInitPrimContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, InterLangParserRULE_initPrim)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Line()
	}
	{
		p.SetState(138)
		p.InternalVar()
	}
	{
		p.SetState(139)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(140)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<InterLangParserTRUE)|(1<<InterLangParserFALSE)|(1<<InterLangParserNUMBER)|(1<<InterLangParserSTRING_RAW))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IInitParamContext is an interface to support dynamic dispatch.
type IInitParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitParamContext differentiates from other interfaces.
	IsInitParamContext()
}

type InitParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitParamContext() *InitParamContext {
	var p = new(InitParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_initParam
	return p
}

func (*InitParamContext) IsInitParamContext() {}

func NewInitParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitParamContext {
	var p = new(InitParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_initParam

	return p
}

func (s *InitParamContext) GetParser() antlr.Parser { return s.parser }

func (s *InitParamContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *InitParamContext) AllInternalVar() []IInternalVarContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInternalVarContext)(nil)).Elem())
	var tst = make([]IInternalVarContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInternalVarContext)
		}
	}

	return tst
}

func (s *InitParamContext) InternalVar(i int) IInternalVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInternalVarContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInternalVarContext)
}

func (s *InitParamContext) EQ() antlr.TerminalNode {
	return s.GetToken(InterLangParserEQ, 0)
}

func (s *InitParamContext) PARAM() antlr.TerminalNode {
	return s.GetToken(InterLangParserPARAM, 0)
}

func (s *InitParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterInitParam(s)
	}
}

func (s *InitParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitInitParam(s)
	}
}

func (p *InterLangParser) InitParam() (localctx IInitParamContext) {
	localctx = NewInitParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, InterLangParserRULE_initParam)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Line()
	}
	{
		p.SetState(143)
		p.InternalVar()
	}
	{
		p.SetState(144)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(145)
		p.Match(InterLangParserPARAM)
	}
	{
		p.SetState(146)
		p.InternalVar()
	}

	return localctx
}

// IInitCallContext is an interface to support dynamic dispatch.
type IInitCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitCallContext differentiates from other interfaces.
	IsInitCallContext()
}

type InitCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitCallContext() *InitCallContext {
	var p = new(InitCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_initCall
	return p
}

func (*InitCallContext) IsInitCallContext() {}

func NewInitCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitCallContext {
	var p = new(InitCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_initCall

	return p
}

func (s *InitCallContext) GetParser() antlr.Parser { return s.parser }

func (s *InitCallContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *InitCallContext) InternalVar() IInternalVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInternalVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInternalVarContext)
}

func (s *InitCallContext) EQ() antlr.TerminalNode {
	return s.GetToken(InterLangParserEQ, 0)
}

func (s *InitCallContext) CALL() antlr.TerminalNode {
	return s.GetToken(InterLangParserCALL, 0)
}

func (s *InitCallContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(InterLangParserNUMBER, 0)
}

func (s *InitCallContext) SYS_FUNC() antlr.TerminalNode {
	return s.GetToken(InterLangParserSYS_FUNC, 0)
}

func (s *InitCallContext) ITEM() antlr.TerminalNode {
	return s.GetToken(InterLangParserITEM, 0)
}

func (s *InitCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterInitCall(s)
	}
}

func (s *InitCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitInitCall(s)
	}
}

func (p *InterLangParser) InitCall() (localctx IInitCallContext) {
	localctx = NewInitCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, InterLangParserRULE_initCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Line()
	}
	{
		p.SetState(149)
		p.InternalVar()
	}
	{
		p.SetState(150)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(151)
		p.Match(InterLangParserCALL)
	}
	{
		p.SetState(152)
		_la = p.GetTokenStream().LA(1)

		if !(_la == InterLangParserSYS_FUNC || _la == InterLangParserITEM) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(153)
		p.Match(InterLangParserNUMBER)
	}

	return localctx
}

// IInitBoolOpContext is an interface to support dynamic dispatch.
type IInitBoolOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitBoolOpContext differentiates from other interfaces.
	IsInitBoolOpContext()
}

type InitBoolOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitBoolOpContext() *InitBoolOpContext {
	var p = new(InitBoolOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_initBoolOp
	return p
}

func (*InitBoolOpContext) IsInitBoolOpContext() {}

func NewInitBoolOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitBoolOpContext {
	var p = new(InitBoolOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_initBoolOp

	return p
}

func (s *InitBoolOpContext) GetParser() antlr.Parser { return s.parser }

func (s *InitBoolOpContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *InitBoolOpContext) AllInternalVar() []IInternalVarContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInternalVarContext)(nil)).Elem())
	var tst = make([]IInternalVarContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInternalVarContext)
		}
	}

	return tst
}

func (s *InitBoolOpContext) InternalVar(i int) IInternalVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInternalVarContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInternalVarContext)
}

func (s *InitBoolOpContext) EQ() antlr.TerminalNode {
	return s.GetToken(InterLangParserEQ, 0)
}

func (s *InitBoolOpContext) BOOL_SIGN() antlr.TerminalNode {
	return s.GetToken(InterLangParserBOOL_SIGN, 0)
}

func (s *InitBoolOpContext) BOOL_PL() antlr.TerminalNode {
	return s.GetToken(InterLangParserBOOL_PL, 0)
}

func (s *InitBoolOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitBoolOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitBoolOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterInitBoolOp(s)
	}
}

func (s *InitBoolOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitInitBoolOp(s)
	}
}

func (p *InterLangParser) InitBoolOp() (localctx IInitBoolOpContext) {
	localctx = NewInitBoolOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, InterLangParserRULE_initBoolOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Line()
	}
	{
		p.SetState(156)
		p.InternalVar()
	}
	{
		p.SetState(157)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(158)
		p.InternalVar()
	}
	{
		p.SetState(159)
		_la = p.GetTokenStream().LA(1)

		if !(_la == InterLangParserBOOL_SIGN || _la == InterLangParserBOOL_PL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(160)
		p.InternalVar()
	}

	return localctx
}

// IInitNumOpContext is an interface to support dynamic dispatch.
type IInitNumOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitNumOpContext differentiates from other interfaces.
	IsInitNumOpContext()
}

type InitNumOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitNumOpContext() *InitNumOpContext {
	var p = new(InitNumOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_initNumOp
	return p
}

func (*InitNumOpContext) IsInitNumOpContext() {}

func NewInitNumOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitNumOpContext {
	var p = new(InitNumOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_initNumOp

	return p
}

func (s *InitNumOpContext) GetParser() antlr.Parser { return s.parser }

func (s *InitNumOpContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *InitNumOpContext) AllInternalVar() []IInternalVarContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInternalVarContext)(nil)).Elem())
	var tst = make([]IInternalVarContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInternalVarContext)
		}
	}

	return tst
}

func (s *InitNumOpContext) InternalVar(i int) IInternalVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInternalVarContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInternalVarContext)
}

func (s *InitNumOpContext) EQ() antlr.TerminalNode {
	return s.GetToken(InterLangParserEQ, 0)
}

func (s *InitNumOpContext) NUM_SIGN() antlr.TerminalNode {
	return s.GetToken(InterLangParserNUM_SIGN, 0)
}

func (s *InitNumOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitNumOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitNumOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterInitNumOp(s)
	}
}

func (s *InitNumOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitInitNumOp(s)
	}
}

func (p *InterLangParser) InitNumOp() (localctx IInitNumOpContext) {
	localctx = NewInitNumOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, InterLangParserRULE_initNumOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Line()
	}
	{
		p.SetState(163)
		p.InternalVar()
	}
	{
		p.SetState(164)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(165)
		p.InternalVar()
	}
	{
		p.SetState(166)
		p.Match(InterLangParserNUM_SIGN)
	}
	{
		p.SetState(167)
		p.InternalVar()
	}

	return localctx
}

// IInitArrElContext is an interface to support dynamic dispatch.
type IInitArrElContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitArrElContext differentiates from other interfaces.
	IsInitArrElContext()
}

type InitArrElContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitArrElContext() *InitArrElContext {
	var p = new(InitArrElContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_initArrEl
	return p
}

func (*InitArrElContext) IsInitArrElContext() {}

func NewInitArrElContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitArrElContext {
	var p = new(InitArrElContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_initArrEl

	return p
}

func (s *InitArrElContext) GetParser() antlr.Parser { return s.parser }

func (s *InitArrElContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *InitArrElContext) AllInternalVar() []IInternalVarContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInternalVarContext)(nil)).Elem())
	var tst = make([]IInternalVarContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInternalVarContext)
		}
	}

	return tst
}

func (s *InitArrElContext) InternalVar(i int) IInternalVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInternalVarContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInternalVarContext)
}

func (s *InitArrElContext) EQ() antlr.TerminalNode {
	return s.GetToken(InterLangParserEQ, 0)
}

func (s *InitArrElContext) ITEM() antlr.TerminalNode {
	return s.GetToken(InterLangParserITEM, 0)
}

func (s *InitArrElContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(InterLangParserARRAY, 0)
}

func (s *InitArrElContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitArrElContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitArrElContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterInitArrEl(s)
	}
}

func (s *InitArrElContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitInitArrEl(s)
	}
}

func (p *InterLangParser) InitArrEl() (localctx IInitArrElContext) {
	localctx = NewInitArrElContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, InterLangParserRULE_initArrEl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Line()
	}
	{
		p.SetState(170)
		p.InternalVar()
	}
	{
		p.SetState(171)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(172)
		p.Match(InterLangParserITEM)
	}
	{
		p.SetState(173)
		p.Match(InterLangParserT__0)
	}
	{
		p.SetState(174)
		p.InternalVar()
	}
	{
		p.SetState(175)
		p.Match(InterLangParserARRAY)
	}

	return localctx
}

// IExtArrElContext is an interface to support dynamic dispatch.
type IExtArrElContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtArrElContext differentiates from other interfaces.
	IsExtArrElContext()
}

type ExtArrElContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtArrElContext() *ExtArrElContext {
	var p = new(ExtArrElContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_extArrEl
	return p
}

func (*ExtArrElContext) IsExtArrElContext() {}

func NewExtArrElContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtArrElContext {
	var p = new(ExtArrElContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_extArrEl

	return p
}

func (s *ExtArrElContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtArrElContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ExtArrElContext) ITEM() antlr.TerminalNode {
	return s.GetToken(InterLangParserITEM, 0)
}

func (s *ExtArrElContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(InterLangParserARRAY, 0)
}

func (s *ExtArrElContext) EQ() antlr.TerminalNode {
	return s.GetToken(InterLangParserEQ, 0)
}

func (s *ExtArrElContext) AllInternalVar() []IInternalVarContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInternalVarContext)(nil)).Elem())
	var tst = make([]IInternalVarContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInternalVarContext)
		}
	}

	return tst
}

func (s *ExtArrElContext) InternalVar(i int) IInternalVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInternalVarContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInternalVarContext)
}

func (s *ExtArrElContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(InterLangParserNUMBER, 0)
}

func (s *ExtArrElContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtArrElContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtArrElContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterExtArrEl(s)
	}
}

func (s *ExtArrElContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitExtArrEl(s)
	}
}

func (p *InterLangParser) ExtArrEl() (localctx IExtArrElContext) {
	localctx = NewExtArrElContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, InterLangParserRULE_extArrEl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Line()
	}
	{
		p.SetState(178)
		p.Match(InterLangParserITEM)
	}
	{
		p.SetState(179)
		p.Match(InterLangParserT__0)
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case InterLangParserNUMBER:
		{
			p.SetState(180)
			p.Match(InterLangParserNUMBER)
		}

	case InterLangParserINNER_VAR:
		{
			p.SetState(181)
			p.InternalVar()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(184)
		p.Match(InterLangParserARRAY)
	}
	{
		p.SetState(185)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(186)
		p.InternalVar()
	}

	return localctx
}

// IReturnTpContext is an interface to support dynamic dispatch.
type IReturnTpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnTpContext differentiates from other interfaces.
	IsReturnTpContext()
}

type ReturnTpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnTpContext() *ReturnTpContext {
	var p = new(ReturnTpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_returnTp
	return p
}

func (*ReturnTpContext) IsReturnTpContext() {}

func NewReturnTpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnTpContext {
	var p = new(ReturnTpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_returnTp

	return p
}

func (s *ReturnTpContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnTpContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ReturnTpContext) RET() antlr.TerminalNode {
	return s.GetToken(InterLangParserRET, 0)
}

func (s *ReturnTpContext) InternalVar() IInternalVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInternalVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInternalVarContext)
}

func (s *ReturnTpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnTpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnTpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterReturnTp(s)
	}
}

func (s *ReturnTpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitReturnTp(s)
	}
}

func (p *InterLangParser) ReturnTp() (localctx IReturnTpContext) {
	localctx = NewReturnTpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, InterLangParserRULE_returnTp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Line()
	}
	{
		p.SetState(189)
		p.Match(InterLangParserRET)
	}
	{
		p.SetState(190)
		p.InternalVar()
	}

	return localctx
}

// IGotoInContext is an interface to support dynamic dispatch.
type IGotoInContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGotoInContext differentiates from other interfaces.
	IsGotoInContext()
}

type GotoInContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGotoInContext() *GotoInContext {
	var p = new(GotoInContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_gotoIn
	return p
}

func (*GotoInContext) IsGotoInContext() {}

func NewGotoInContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GotoInContext {
	var p = new(GotoInContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_gotoIn

	return p
}

func (s *GotoInContext) GetParser() antlr.Parser { return s.parser }

func (s *GotoInContext) GOTO() antlr.TerminalNode {
	return s.GetToken(InterLangParserGOTO, 0)
}

func (s *GotoInContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(InterLangParserNUMBER, 0)
}

func (s *GotoInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GotoInContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GotoInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterGotoIn(s)
	}
}

func (s *GotoInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitGotoIn(s)
	}
}

func (p *InterLangParser) GotoIn() (localctx IGotoInContext) {
	localctx = NewGotoInContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, InterLangParserRULE_gotoIn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(InterLangParserGOTO)
	}
	{
		p.SetState(193)
		p.Match(InterLangParserNUMBER)
	}

	return localctx
}

// IGotoTpContext is an interface to support dynamic dispatch.
type IGotoTpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGotoTpContext differentiates from other interfaces.
	IsGotoTpContext()
}

type GotoTpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGotoTpContext() *GotoTpContext {
	var p = new(GotoTpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_gotoTp
	return p
}

func (*GotoTpContext) IsGotoTpContext() {}

func NewGotoTpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GotoTpContext {
	var p = new(GotoTpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_gotoTp

	return p
}

func (s *GotoTpContext) GetParser() antlr.Parser { return s.parser }

func (s *GotoTpContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *GotoTpContext) GotoIn() IGotoInContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGotoInContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGotoInContext)
}

func (s *GotoTpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GotoTpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GotoTpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterGotoTp(s)
	}
}

func (s *GotoTpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitGotoTp(s)
	}
}

func (p *InterLangParser) GotoTp() (localctx IGotoTpContext) {
	localctx = NewGotoTpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, InterLangParserRULE_gotoTp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Line()
	}
	{
		p.SetState(196)
		p.GotoIn()
	}

	return localctx
}

// IIfFalseContext is an interface to support dynamic dispatch.
type IIfFalseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfFalseContext differentiates from other interfaces.
	IsIfFalseContext()
}

type IfFalseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfFalseContext() *IfFalseContext {
	var p = new(IfFalseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_ifFalse
	return p
}

func (*IfFalseContext) IsIfFalseContext() {}

func NewIfFalseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfFalseContext {
	var p = new(IfFalseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_ifFalse

	return p
}

func (s *IfFalseContext) GetParser() antlr.Parser { return s.parser }

func (s *IfFalseContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *IfFalseContext) IFFALSE() antlr.TerminalNode {
	return s.GetToken(InterLangParserIFFALSE, 0)
}

func (s *IfFalseContext) InternalVar() IInternalVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInternalVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInternalVarContext)
}

func (s *IfFalseContext) GotoIn() IGotoInContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGotoInContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGotoInContext)
}

func (s *IfFalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfFalseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfFalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterIfFalse(s)
	}
}

func (s *IfFalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitIfFalse(s)
	}
}

func (p *InterLangParser) IfFalse() (localctx IIfFalseContext) {
	localctx = NewIfFalseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, InterLangParserRULE_ifFalse)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Line()
	}
	{
		p.SetState(199)
		p.Match(InterLangParserIFFALSE)
	}
	{
		p.SetState(200)
		p.InternalVar()
	}
	{
		p.SetState(201)
		p.GotoIn()
	}

	return localctx
}

// ITypeTpContext is an interface to support dynamic dispatch.
type ITypeTpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeTpContext differentiates from other interfaces.
	IsTypeTpContext()
}

type TypeTpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeTpContext() *TypeTpContext {
	var p = new(TypeTpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_typeTp
	return p
}

func (*TypeTpContext) IsTypeTpContext() {}

func NewTypeTpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeTpContext {
	var p = new(TypeTpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_typeTp

	return p
}

func (s *TypeTpContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeTpContext) STR() antlr.TerminalNode {
	return s.GetToken(InterLangParserSTR, 0)
}

func (s *TypeTpContext) BOOL() antlr.TerminalNode {
	return s.GetToken(InterLangParserBOOL, 0)
}

func (s *TypeTpContext) NUM() antlr.TerminalNode {
	return s.GetToken(InterLangParserNUM, 0)
}

func (s *TypeTpContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(InterLangParserARRAY, 0)
}

func (s *TypeTpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeTpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeTpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterTypeTp(s)
	}
}

func (s *TypeTpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitTypeTp(s)
	}
}

func (p *InterLangParser) TypeTp() (localctx ITypeTpContext) {
	localctx = NewTypeTpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, InterLangParserRULE_typeTp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == InterLangParserARRAY {
		{
			p.SetState(203)
			p.Match(InterLangParserARRAY)
		}

	}
	{
		p.SetState(206)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<InterLangParserSTR)|(1<<InterLangParserBOOL)|(1<<InterLangParserNUM))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(InterLangParserNUMBER, 0)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *InterLangParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, InterLangParserRULE_line)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(InterLangParserNUMBER)
	}
	{
		p.SetState(209)
		p.Match(InterLangParserT__1)
	}

	return localctx
}
