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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 31, 213,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 7, 2, 52, 10, 2, 12, 2, 14, 2, 55, 11,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 63, 10, 3, 12, 3, 14, 3, 66,
	11, 3, 3, 3, 3, 3, 7, 3, 70, 10, 3, 12, 3, 14, 3, 73, 11, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 5, 4, 80, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 5, 5, 89, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 105, 10, 6, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9,
	121, 10, 9, 3, 9, 3, 9, 3, 10, 7, 10, 126, 10, 10, 12, 10, 14, 10, 129,
	11, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 188, 10, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3,
	25, 3, 25, 2, 2, 26, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 2, 5, 4, 2, 15, 16, 22, 23, 4,
	2, 27, 27, 31, 31, 3, 2, 12, 14, 2, 209, 2, 53, 3, 2, 2, 2, 4, 58, 3, 2,
	2, 2, 6, 74, 3, 2, 2, 2, 8, 83, 3, 2, 2, 2, 10, 104, 3, 2, 2, 2, 12, 106,
	3, 2, 2, 2, 14, 112, 3, 2, 2, 2, 16, 115, 3, 2, 2, 2, 18, 127, 3, 2, 2,
	2, 20, 132, 3, 2, 2, 2, 22, 137, 3, 2, 2, 2, 24, 142, 3, 2, 2, 2, 26, 147,
	3, 2, 2, 2, 28, 153, 3, 2, 2, 2, 30, 160, 3, 2, 2, 2, 32, 167, 3, 2, 2,
	2, 34, 174, 3, 2, 2, 2, 36, 182, 3, 2, 2, 2, 38, 193, 3, 2, 2, 2, 40, 196,
	3, 2, 2, 2, 42, 199, 3, 2, 2, 2, 44, 202, 3, 2, 2, 2, 46, 207, 3, 2, 2,
	2, 48, 209, 3, 2, 2, 2, 50, 52, 5, 4, 3, 2, 51, 50, 3, 2, 2, 2, 52, 55,
	3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 56, 3, 2, 2, 2,
	55, 53, 3, 2, 2, 2, 56, 57, 7, 2, 2, 3, 57, 3, 3, 2, 2, 2, 58, 59, 5, 48,
	25, 2, 59, 60, 7, 6, 2, 2, 60, 64, 7, 31, 2, 2, 61, 63, 5, 6, 4, 2, 62,
	61, 3, 2, 2, 2, 63, 66, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2,
	2, 65, 67, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 67, 71, 5, 8, 5, 2, 68, 70,
	5, 10, 6, 2, 69, 68, 3, 2, 2, 2, 70, 73, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2,
	71, 72, 3, 2, 2, 2, 72, 5, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 74, 75, 5, 48,
	25, 2, 75, 76, 7, 7, 2, 2, 76, 77, 7, 31, 2, 2, 77, 79, 7, 8, 2, 2, 78,
	80, 7, 28, 2, 2, 79, 78, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 81, 3, 2,
	2, 2, 81, 82, 5, 46, 24, 2, 82, 7, 3, 2, 2, 2, 83, 84, 5, 48, 25, 2, 84,
	85, 7, 9, 2, 2, 85, 88, 7, 8, 2, 2, 86, 89, 5, 46, 24, 2, 87, 89, 7, 11,
	2, 2, 88, 86, 3, 2, 2, 2, 88, 87, 3, 2, 2, 2, 89, 9, 3, 2, 2, 2, 90, 105,
	5, 16, 9, 2, 91, 105, 5, 20, 11, 2, 92, 105, 5, 18, 10, 2, 93, 105, 5,
	22, 12, 2, 94, 105, 5, 24, 13, 2, 95, 105, 5, 26, 14, 2, 96, 105, 5, 28,
	15, 2, 97, 105, 5, 30, 16, 2, 98, 105, 5, 32, 17, 2, 99, 105, 5, 34, 18,
	2, 100, 105, 5, 36, 19, 2, 101, 105, 5, 42, 22, 2, 102, 105, 5, 44, 23,
	2, 103, 105, 5, 38, 20, 2, 104, 90, 3, 2, 2, 2, 104, 91, 3, 2, 2, 2, 104,
	92, 3, 2, 2, 2, 104, 93, 3, 2, 2, 2, 104, 94, 3, 2, 2, 2, 104, 95, 3, 2,
	2, 2, 104, 96, 3, 2, 2, 2, 104, 97, 3, 2, 2, 2, 104, 98, 3, 2, 2, 2, 104,
	99, 3, 2, 2, 2, 104, 100, 3, 2, 2, 2, 104, 101, 3, 2, 2, 2, 104, 102, 3,
	2, 2, 2, 104, 103, 3, 2, 2, 2, 105, 11, 3, 2, 2, 2, 106, 107, 5, 48, 25,
	2, 107, 108, 5, 14, 8, 2, 108, 109, 7, 8, 2, 2, 109, 110, 7, 18, 2, 2,
	110, 111, 9, 2, 2, 2, 111, 13, 3, 2, 2, 2, 112, 113, 7, 29, 2, 2, 113,
	114, 7, 22, 2, 2, 114, 15, 3, 2, 2, 2, 115, 116, 5, 48, 25, 2, 116, 117,
	7, 31, 2, 2, 117, 118, 7, 8, 2, 2, 118, 120, 7, 17, 2, 2, 119, 121, 7,
	28, 2, 2, 120, 119, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 122, 3, 2, 2,
	2, 122, 123, 5, 46, 24, 2, 123, 17, 3, 2, 2, 2, 124, 126, 5, 12, 7, 2,
	125, 124, 3, 2, 2, 2, 126, 129, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 127,
	128, 3, 2, 2, 2, 128, 130, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 130, 131,
	5, 16, 9, 2, 131, 19, 3, 2, 2, 2, 132, 133, 5, 48, 25, 2, 133, 134, 7,
	31, 2, 2, 134, 135, 7, 8, 2, 2, 135, 136, 5, 14, 8, 2, 136, 21, 3, 2, 2,
	2, 137, 138, 5, 48, 25, 2, 138, 139, 5, 14, 8, 2, 139, 140, 7, 8, 2, 2,
	140, 141, 7, 31, 2, 2, 141, 23, 3, 2, 2, 2, 142, 143, 5, 48, 25, 2, 143,
	144, 5, 14, 8, 2, 144, 145, 7, 8, 2, 2, 145, 146, 9, 2, 2, 2, 146, 25,
	3, 2, 2, 2, 147, 148, 5, 48, 25, 2, 148, 149, 5, 14, 8, 2, 149, 150, 7,
	8, 2, 2, 150, 151, 7, 30, 2, 2, 151, 152, 5, 14, 8, 2, 152, 27, 3, 2, 2,
	2, 153, 154, 5, 48, 25, 2, 154, 155, 5, 14, 8, 2, 155, 156, 7, 8, 2, 2,
	156, 157, 7, 19, 2, 2, 157, 158, 9, 3, 2, 2, 158, 159, 7, 22, 2, 2, 159,
	29, 3, 2, 2, 2, 160, 161, 5, 48, 25, 2, 161, 162, 5, 14, 8, 2, 162, 163,
	7, 8, 2, 2, 163, 164, 5, 14, 8, 2, 164, 165, 7, 25, 2, 2, 165, 166, 5,
	14, 8, 2, 166, 31, 3, 2, 2, 2, 167, 168, 5, 48, 25, 2, 168, 169, 5, 14,
	8, 2, 169, 170, 7, 8, 2, 2, 170, 171, 5, 14, 8, 2, 171, 172, 7, 24, 2,
	2, 172, 173, 5, 14, 8, 2, 173, 33, 3, 2, 2, 2, 174, 175, 5, 48, 25, 2,
	175, 176, 5, 14, 8, 2, 176, 177, 7, 8, 2, 2, 177, 178, 7, 31, 2, 2, 178,
	179, 7, 3, 2, 2, 179, 180, 5, 14, 8, 2, 180, 181, 7, 28, 2, 2, 181, 35,
	3, 2, 2, 2, 182, 183, 5, 48, 25, 2, 183, 184, 7, 31, 2, 2, 184, 187, 7,
	3, 2, 2, 185, 188, 7, 22, 2, 2, 186, 188, 5, 14, 8, 2, 187, 185, 3, 2,
	2, 2, 187, 186, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 190, 7, 28, 2, 2,
	190, 191, 7, 8, 2, 2, 191, 192, 5, 14, 8, 2, 192, 37, 3, 2, 2, 2, 193,
	194, 5, 48, 25, 2, 194, 195, 5, 14, 8, 2, 195, 39, 3, 2, 2, 2, 196, 197,
	7, 21, 2, 2, 197, 198, 7, 22, 2, 2, 198, 41, 3, 2, 2, 2, 199, 200, 5, 48,
	25, 2, 200, 201, 5, 40, 21, 2, 201, 43, 3, 2, 2, 2, 202, 203, 5, 48, 25,
	2, 203, 204, 7, 20, 2, 2, 204, 205, 5, 14, 8, 2, 205, 206, 5, 40, 21, 2,
	206, 45, 3, 2, 2, 2, 207, 208, 9, 4, 2, 2, 208, 47, 3, 2, 2, 2, 209, 210,
	7, 22, 2, 2, 210, 211, 7, 4, 2, 2, 211, 49, 3, 2, 2, 2, 11, 53, 64, 71,
	79, 88, 104, 120, 127, 187,
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
	"newVar", "newArrVar", "updVar", "initItem", "initPrim", "initParam", "initCall",
	"initBoolOp", "initNumOp", "initArrEl", "extArrEl", "return", "gotoIn",
	"goto", "ifFalse", "type", "line",
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
	InterLangParserRULE_newArrVar      = 8
	InterLangParserRULE_updVar         = 9
	InterLangParserRULE_initItem       = 10
	InterLangParserRULE_initPrim       = 11
	InterLangParserRULE_initParam      = 12
	InterLangParserRULE_initCall       = 13
	InterLangParserRULE_initBoolOp     = 14
	InterLangParserRULE_initNumOp      = 15
	InterLangParserRULE_initArrEl      = 16
	InterLangParserRULE_extArrEl       = 17
	InterLangParserRULE_return         = 18
	InterLangParserRULE_gotoIn         = 19
	InterLangParserRULE_goto           = 20
	InterLangParserRULE_ifFalse        = 21
	InterLangParserRULE_type           = 22
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

func (s *FileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitFile(s)

	default:
		return t.VisitChildren(s)
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

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
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

func (s *ArgContext) Type() ITypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ArgContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(InterLangParserARRAY, 0)
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

func (s *ArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InterLangParser) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, InterLangParserRULE_arg)
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
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == InterLangParserARRAY {
		{
			p.SetState(76)
			p.Match(InterLangParserARRAY)
		}

	}
	{
		p.SetState(79)
		p.Type()
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

func (s *RetTpContext) Type() ITypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
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

func (s *RetTpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitRetTp(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(81)
		p.Line()
	}
	{
		p.SetState(82)
		p.Match(InterLangParserRET_TP)
	}
	{
		p.SetState(83)
		p.Match(InterLangParserEQ)
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case InterLangParserSTR, InterLangParserBOOL, InterLangParserNUM:
		{
			p.SetState(84)
			p.Type()
		}

	case InterLangParserVOID:
		{
			p.SetState(85)
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

func (s *StatementContext) NewArrVar() INewArrVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewArrVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewArrVarContext)
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

func (s *StatementContext) Goto() IGotoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGotoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGotoContext)
}

func (s *StatementContext) IfFalse() IIfFalseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfFalseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfFalseContext)
}

func (s *StatementContext) Return() IReturnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnContext)
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

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
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

	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.NewVar()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.UpdVar()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.NewArrVar()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)
			p.InitItem()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(92)
			p.InitPrim()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(93)
			p.InitParam()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(94)
			p.InitCall()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(95)
			p.InitBoolOp()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(96)
			p.InitNumOp()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(97)
			p.InitArrEl()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(98)
			p.ExtArrEl()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(99)
			p.Goto()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(100)
			p.IfFalse()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(101)
			p.Return()
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

func (s *InternalArrArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitInternalArrArg(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(104)
		p.Line()
	}
	{
		p.SetState(105)
		p.InternalVar()
	}
	{
		p.SetState(106)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(107)
		p.Match(InterLangParserINIT_ARR)
	}
	{
		p.SetState(108)
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

func (s *InternalVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitInternalVar(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(110)
		p.Match(InterLangParserINNER_VAR)
	}
	{
		p.SetState(111)
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

func (s *NewVarContext) Type() ITypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *NewVarContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(InterLangParserARRAY, 0)
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

func (s *NewVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitNewVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InterLangParser) NewVar() (localctx INewVarContext) {
	localctx = NewNewVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, InterLangParserRULE_newVar)
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
		p.SetState(113)
		p.Line()
	}
	{
		p.SetState(114)
		p.Match(InterLangParserITEM)
	}
	{
		p.SetState(115)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(116)
		p.Match(InterLangParserINIT)
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == InterLangParserARRAY {
		{
			p.SetState(117)
			p.Match(InterLangParserARRAY)
		}

	}
	{
		p.SetState(120)
		p.Type()
	}

	return localctx
}

// INewArrVarContext is an interface to support dynamic dispatch.
type INewArrVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNewArrVarContext differentiates from other interfaces.
	IsNewArrVarContext()
}

type NewArrVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNewArrVarContext() *NewArrVarContext {
	var p = new(NewArrVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_newArrVar
	return p
}

func (*NewArrVarContext) IsNewArrVarContext() {}

func NewNewArrVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NewArrVarContext {
	var p = new(NewArrVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_newArrVar

	return p
}

func (s *NewArrVarContext) GetParser() antlr.Parser { return s.parser }

func (s *NewArrVarContext) NewVar() INewVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewVarContext)
}

func (s *NewArrVarContext) AllInternalArrArg() []IInternalArrArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInternalArrArgContext)(nil)).Elem())
	var tst = make([]IInternalArrArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInternalArrArgContext)
		}
	}

	return tst
}

func (s *NewArrVarContext) InternalArrArg(i int) IInternalArrArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInternalArrArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInternalArrArgContext)
}

func (s *NewArrVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewArrVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NewArrVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterNewArrVar(s)
	}
}

func (s *NewArrVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitNewArrVar(s)
	}
}

func (s *NewArrVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitNewArrVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InterLangParser) NewArrVar() (localctx INewArrVarContext) {
	localctx = NewNewArrVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, InterLangParserRULE_newArrVar)

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
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(122)
				p.InternalArrArg()
			}

		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}
	{
		p.SetState(128)
		p.NewVar()
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

func (s *UpdVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitUpdVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InterLangParser) UpdVar() (localctx IUpdVarContext) {
	localctx = NewUpdVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, InterLangParserRULE_updVar)

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
		p.SetState(130)
		p.Line()
	}
	{
		p.SetState(131)
		p.Match(InterLangParserITEM)
	}
	{
		p.SetState(132)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(133)
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

func (s *InitItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitInitItem(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(135)
		p.Line()
	}
	{
		p.SetState(136)
		p.InternalVar()
	}
	{
		p.SetState(137)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(138)
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

func (s *InitPrimContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitInitPrim(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(140)
		p.Line()
	}
	{
		p.SetState(141)
		p.InternalVar()
	}
	{
		p.SetState(142)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(143)
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

func (s *InitParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitInitParam(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(145)
		p.Line()
	}
	{
		p.SetState(146)
		p.InternalVar()
	}
	{
		p.SetState(147)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(148)
		p.Match(InterLangParserPARAM)
	}
	{
		p.SetState(149)
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

func (s *InitCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitInitCall(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(151)
		p.Line()
	}
	{
		p.SetState(152)
		p.InternalVar()
	}
	{
		p.SetState(153)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(154)
		p.Match(InterLangParserCALL)
	}
	{
		p.SetState(155)
		_la = p.GetTokenStream().LA(1)

		if !(_la == InterLangParserSYS_FUNC || _la == InterLangParserITEM) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(156)
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

func (s *InitBoolOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitInitBoolOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InterLangParser) InitBoolOp() (localctx IInitBoolOpContext) {
	localctx = NewInitBoolOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, InterLangParserRULE_initBoolOp)

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
		p.SetState(158)
		p.Line()
	}
	{
		p.SetState(159)
		p.InternalVar()
	}
	{
		p.SetState(160)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(161)
		p.InternalVar()
	}
	{
		p.SetState(162)
		p.Match(InterLangParserBOOL_SIGN)
	}
	{
		p.SetState(163)
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

func (s *InitNumOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitInitNumOp(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(165)
		p.Line()
	}
	{
		p.SetState(166)
		p.InternalVar()
	}
	{
		p.SetState(167)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(168)
		p.InternalVar()
	}
	{
		p.SetState(169)
		p.Match(InterLangParserNUM_SIGN)
	}
	{
		p.SetState(170)
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

func (s *InitArrElContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitInitArrEl(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(172)
		p.Line()
	}
	{
		p.SetState(173)
		p.InternalVar()
	}
	{
		p.SetState(174)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(175)
		p.Match(InterLangParserITEM)
	}
	{
		p.SetState(176)
		p.Match(InterLangParserT__0)
	}
	{
		p.SetState(177)
		p.InternalVar()
	}
	{
		p.SetState(178)
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

func (s *ExtArrElContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitExtArrEl(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(180)
		p.Line()
	}
	{
		p.SetState(181)
		p.Match(InterLangParserITEM)
	}
	{
		p.SetState(182)
		p.Match(InterLangParserT__0)
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case InterLangParserNUMBER:
		{
			p.SetState(183)
			p.Match(InterLangParserNUMBER)
		}

	case InterLangParserINNER_VAR:
		{
			p.SetState(184)
			p.InternalVar()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(187)
		p.Match(InterLangParserARRAY)
	}
	{
		p.SetState(188)
		p.Match(InterLangParserEQ)
	}
	{
		p.SetState(189)
		p.InternalVar()
	}

	return localctx
}

// IReturnContext is an interface to support dynamic dispatch.
type IReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnContext differentiates from other interfaces.
	IsReturnContext()
}

type ReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnContext() *ReturnContext {
	var p = new(ReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_return
	return p
}

func (*ReturnContext) IsReturnContext() {}

func NewReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnContext {
	var p = new(ReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_return

	return p
}

func (s *ReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ReturnContext) InternalVar() IInternalVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInternalVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInternalVarContext)
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterReturn(s)
	}
}

func (s *ReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitReturn(s)
	}
}

func (s *ReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InterLangParser) Return() (localctx IReturnContext) {
	localctx = NewReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, InterLangParserRULE_return)

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
		p.SetState(191)
		p.Line()
	}
	{
		p.SetState(192)
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

func (s *GotoInContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitGotoIn(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(194)
		p.Match(InterLangParserGOTO)
	}
	{
		p.SetState(195)
		p.Match(InterLangParserNUMBER)
	}

	return localctx
}

// IGotoContext is an interface to support dynamic dispatch.
type IGotoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGotoContext differentiates from other interfaces.
	IsGotoContext()
}

type GotoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGotoContext() *GotoContext {
	var p = new(GotoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_goto
	return p
}

func (*GotoContext) IsGotoContext() {}

func NewGotoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GotoContext {
	var p = new(GotoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_goto

	return p
}

func (s *GotoContext) GetParser() antlr.Parser { return s.parser }

func (s *GotoContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *GotoContext) GotoIn() IGotoInContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGotoInContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGotoInContext)
}

func (s *GotoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GotoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GotoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterGoto(s)
	}
}

func (s *GotoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitGoto(s)
	}
}

func (s *GotoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitGoto(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InterLangParser) Goto() (localctx IGotoContext) {
	localctx = NewGotoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, InterLangParserRULE_goto)

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
		p.SetState(197)
		p.Line()
	}
	{
		p.SetState(198)
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

func (s *IfFalseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitIfFalse(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(200)
		p.Line()
	}
	{
		p.SetState(201)
		p.Match(InterLangParserIFFALSE)
	}
	{
		p.SetState(202)
		p.InternalVar()
	}
	{
		p.SetState(203)
		p.GotoIn()
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InterLangParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterLangParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) STR() antlr.TerminalNode {
	return s.GetToken(InterLangParserSTR, 0)
}

func (s *TypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(InterLangParserBOOL, 0)
}

func (s *TypeContext) NUM() antlr.TerminalNode {
	return s.GetToken(InterLangParserNUM, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterLangListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *InterLangParser) Type() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, InterLangParserRULE_type)
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
		p.SetState(205)
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

func (s *LineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case InterLangVisitor:
		return t.VisitLine(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(207)
		p.Match(InterLangParserNUMBER)
	}
	{
		p.SetState(208)
		p.Match(InterLangParserT__1)
	}

	return localctx
}
