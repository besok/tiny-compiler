grammar Func;

WS : [ \t\r\n]+ -> skip ;
NUMBER: [0-9]+;
STRING_RAW: '"' ( '\\"' | . )*? '"' ;
BOOL_VAL: TRUE | FALSE;
STRING : 'string';
VOID : 'void';
NUM : 'num';
BOOL: 'bool';
NULL: 'null';
TRUE: 'true';
FALSE: 'false';
ITEM : [a-z][a-zA-Z0-9]* ;
SQ : '[]';
COMMENT : '/*' .*? '*/' -> skip ;
ARRAY: SQ (STRING | NUM | BOOL);
NUM_SIGN: '+' | '-' | '*' | '/' | '%';

func: 'func' ITEM '(' args ')' ( type | VOID ) '{' '}' ;
args: ( ITEM type (',' ITEM type)*)? ;

funcInvoc: ITEM '('argsInvoc')' ;
argsInvoc:( argInvoc (',' argInvoc)* )?;
argInvoc : NUMBER | STRING_RAW | BOOL_VAL | ITEM | funcInvoc;

updVar: ITEM '=' (BOOL_VAL | STRING_RAW | NUMBER | ITEM | arrayInit | funcInvoc | expr);
newVar: 'var' ITEM type '=' (BOOL_VAL | STRING_RAW | NUMBER | ITEM | arrayInit | funcInvoc | expr);

type: ARRAY | STRING | NUM | BOOL;

arrayInit: '{' arrayInitElems '}' ;
arrayInitElems:
    (BOOL_VAL (',' BOOL_VAL)* )
    |
    (STRING_RAW (',' STRING_RAW)* )
    |
    (NUMBER (',' NUMBER)* );

expr:
       ((exprOp) NUM_SIGN (exprOp | expr))
       |
       ('(' (exprOp | expr) NUM_SIGN (exprOp | expr)')')
       ;
exprOp: funcInvoc | NUMBER | ITEM | STRING_RAW  ;



