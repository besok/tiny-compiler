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
IFELSE:'ifFalse';
GOTO:'goto';
NUMBER: [0-9]+;
STRING_RAW: '"' ( '\\"' | . )*? '"' ;
ITEM : [a-z][a-zA-Z0-9]* ;
NUM_SIGN: '+' | '-' | '*' | '/' | '%';
BOOL_SIGN:  '==' | '!=' | '>' | '<';
BOOL_PL:  '&&' | '||';
SYS_FUNC: 'Add' | 'Len' | 'Output' | 'Input' ;
ARRAY:']';
INNER_VAR:'_s' | '_n' | '_b' | '_p';

function: line FUN ITEM arg* retTp statement*;
arg: line ARG ITEM EQ ARRAY? type;
retTp: line RET_TP EQ (type | VOID);
statement:initVar | initArrVar;

initVar:line ITEM EQ INIT ARRAY? type;
initArrVar: initArrArg* initVar;
initArrArg: line internalVar EQ INIT_ARR (NUMBER | STRING_RAW | TRUE | FALSE);
internalVar:INNER_VAR NUMBER;


type: STR | BOOL | NUM;
line: NUMBER ':';