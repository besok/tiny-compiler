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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 27, 37, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 7, 2, 14,
	10, 2, 12, 2, 14, 2, 17, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	3, 25, 10, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 33, 10, 4, 3, 5,
	3, 5, 3, 5, 2, 2, 6, 2, 4, 6, 8, 2, 3, 3, 2, 10, 12, 2, 35, 2, 10, 3, 2,
	2, 2, 4, 20, 3, 2, 2, 2, 6, 28, 3, 2, 2, 2, 8, 34, 3, 2, 2, 2, 10, 11,
	7, 4, 2, 2, 11, 15, 7, 22, 2, 2, 12, 14, 5, 4, 3, 2, 13, 12, 3, 2, 2, 2,
	14, 17, 3, 2, 2, 2, 15, 13, 3, 2, 2, 2, 15, 16, 3, 2, 2, 2, 16, 18, 3,
	2, 2, 2, 17, 15, 3, 2, 2, 2, 18, 19, 7, 7, 2, 2, 19, 3, 3, 2, 2, 2, 20,
	21, 7, 5, 2, 2, 21, 22, 7, 22, 2, 2, 22, 24, 7, 6, 2, 2, 23, 25, 7, 5,
	2, 2, 24, 23, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 27,
	5, 6, 4, 2, 27, 5, 3, 2, 2, 2, 28, 29, 7, 7, 2, 2, 29, 32, 7, 6, 2, 2,
	30, 33, 5, 8, 5, 2, 31, 33, 7, 9, 2, 2, 32, 30, 3, 2, 2, 2, 32, 31, 3,
	2, 2, 2, 33, 7, 3, 2, 2, 2, 34, 35, 9, 2, 2, 2, 35, 9, 3, 2, 2, 2, 5, 15,
	24, 32,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'function'", "'argument'", "'='", "'return_type'", "'return'",
	"'void'", "'string'", "'bool'", "'num'", "'true'", "'false'", "'init'",
	"'init_arr'", "'call'", "'ifFalse'", "'goto'", "", "", "", "", "", "",
	"", "']'",
}
var symbolicNames = []string{
	"", "WS", "FUN", "ARG", "EQ", "RET_TP", "RET", "VOID", "STR", "BOOL", "NUM",
	"TRUE", "FALSE", "INIT", "INIT_ARR", "CALL", "IFELSE", "GOTO", "NUMBER",
	"STRING_RAW", "ITEM", "NUM_SIGN", "BOOL_SIGN", "BOOL_PL", "SYS_FUNC", "ARR",
}

var ruleNames = []string{
	"function", "arg", "retTp", "type",
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
	InterLangParserWS         = 1
	InterLangParserFUN        = 2
	InterLangParserARG        = 3
	InterLangParserEQ         = 4
	InterLangParserRET_TP     = 5
	InterLangParserRET        = 6
	InterLangParserVOID       = 7
	InterLangParserSTR        = 8
	InterLangParserBOOL       = 9
	InterLangParserNUM        = 10
	InterLangParserTRUE       = 11
	InterLangParserFALSE      = 12
	InterLangParserINIT       = 13
	InterLangParserINIT_ARR   = 14
	InterLangParserCALL       = 15
	InterLangParserIFELSE     = 16
	InterLangParserGOTO       = 17
	InterLangParserNUMBER     = 18
	InterLangParserSTRING_RAW = 19
	InterLangParserITEM       = 20
	InterLangParserNUM_SIGN   = 21
	InterLangParserBOOL_SIGN  = 22
	InterLangParserBOOL_PL    = 23
	InterLangParserSYS_FUNC   = 24
	InterLangParserARR        = 25
)

// InterLangParser rules.
const (
	InterLangParserRULE_function = 0
	InterLangParserRULE_arg      = 1
	InterLangParserRULE_retTp    = 2
	InterLangParserRULE_type     = 3
)

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

func (s *FunctionContext) FUN() antlr.TerminalNode {
	return s.GetToken(InterLangParserFUN, 0)
}

func (s *FunctionContext) ITEM() antlr.TerminalNode {
	return s.GetToken(InterLangParserITEM, 0)
}

func (s *FunctionContext) RET_TP() antlr.TerminalNode {
	return s.GetToken(InterLangParserRET_TP, 0)
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
	p.EnterRule(localctx, 0, InterLangParserRULE_function)
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
		p.SetState(8)
		p.Match(InterLangParserFUN)
	}
	{
		p.SetState(9)
		p.Match(InterLangParserITEM)
	}
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == InterLangParserARG {
		{
			p.SetState(10)
			p.Arg()
		}

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(16)
		p.Match(InterLangParserRET_TP)
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

func (s *ArgContext) AllARG() []antlr.TerminalNode {
	return s.GetTokens(InterLangParserARG)
}

func (s *ArgContext) ARG(i int) antlr.TerminalNode {
	return s.GetToken(InterLangParserARG, i)
}

func (s *ArgContext) ITEM() antlr.TerminalNode {
	return s.GetToken(InterLangParserITEM, 0)
}

func (s *ArgContext) EQ() antlr.TerminalNode {
	return s.GetToken(InterLangParserEQ, 0)
}

func (s *ArgContext) RetTp() IRetTpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRetTpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRetTpContext)
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
	p.EnterRule(localctx, 2, InterLangParserRULE_arg)
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
		p.SetState(18)
		p.Match(InterLangParserARG)
	}
	{
		p.SetState(19)
		p.Match(InterLangParserITEM)
	}
	{
		p.SetState(20)
		p.Match(InterLangParserEQ)
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == InterLangParserARG {
		{
			p.SetState(21)
			p.Match(InterLangParserARG)
		}

	}
	{
		p.SetState(24)
		p.RetTp()
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
	p.EnterRule(localctx, 4, InterLangParserRULE_retTp)

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
		p.SetState(26)
		p.Match(InterLangParserRET_TP)
	}
	{
		p.SetState(27)
		p.Match(InterLangParserEQ)
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case InterLangParserSTR, InterLangParserBOOL, InterLangParserNUM:
		{
			p.SetState(28)
			p.Type()
		}

	case InterLangParserVOID:
		{
			p.SetState(29)
			p.Match(InterLangParserVOID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 6, InterLangParserRULE_type)
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
		p.SetState(32)
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
