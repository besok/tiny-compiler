grammar InterLang;

WS : [ \t\r\n]+ -> skip ;
FUN: 'function';
ARG: 'argument';
EQ:'=';
RET_TP:'return_type';
RET :'return';
VOID:'void';
STR: 'string';
BOOL: 'bool';
NUM: 'num';
TRUE: 'true';
FALSE: 'false';
INIT: 'init';
INIT_ARR:'init_arr';
CALL: 'call';
IFFALSE:'ifFalse';
GOTO:'goto';
NUMBER: [0-9]+;
STRING_RAW: '"' ( '\\"' | . )*? '"' ;
NUM_SIGN: '+' | '-' | '*' | '/' | '%';
BOOL_SIGN:  '==' | '!=' | '>' | '<';
BOOL_PL:  '&&' | '||';
SYS_FUNC: 'Add' | 'Len' | 'Output' | 'Input' ;
ARRAY:']';
INNER_VAR:'_s' | '_n' | '_b' | '_p';
PARAM: 'param';
ITEM : [a-z][a-zA-Z0-9]* ;

file: function* EOF;

function: line FUN ITEM arg* retTp statement*;
arg: line ARG ITEM EQ ARRAY? type;
retTp: line RET_TP EQ (ARRAY? type | VOID);
statement:
      newVar
    | updVar
    | newArrVar
    | initItem
    | initPrim
    | initParam
    | initCall
    | initBoolOp
    | initNumOp
    | initArrEl
    | extArrEl
    | goto
    | ifFalse
    | initIntVar
    | return;



internalArrArg: line internalVar EQ INIT_ARR (NUMBER | STRING_RAW | TRUE | FALSE);
internalVar:INNER_VAR NUMBER;

newVar:line ITEM EQ INIT ARRAY? type;
newArrVar: internalArrArg* newVar;
updVar: line ITEM EQ internalVar;
initIntVar: line internalVar EQ internalVar;
initItem: line internalVar EQ ITEM;
initPrim: line internalVar EQ (NUMBER | TRUE | FALSE | STRING_RAW);
initParam: line internalVar EQ PARAM internalVar;
initCall: line internalVar EQ CALL (SYS_FUNC | ITEM) NUMBER ;
initBoolOp: line internalVar EQ internalVar BOOL_SIGN internalVar ;
initNumOp: line internalVar EQ internalVar NUM_SIGN internalVar ;
initArrEl: line internalVar EQ ITEM'['internalVar']';
extArrEl: line ITEM'['(NUMBER | internalVar )']' EQ internalVar;

return: line RET internalVar;

gotoIn: GOTO NUMBER;
goto: line gotoIn;

ifFalse: line IFFALSE internalVar gotoIn;

type: STR | BOOL | NUM;
line: NUMBER ':';