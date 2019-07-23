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
ITEM : [a-z][a-zA-Z0-9]+ ;
SQ : '[]';
COMMENT : '/*' .*? '*/' -> skip ;
ARRAY: SQ (STRING | NUM | BOOL);


func: 'func' ITEM '(' args ')' ( type | VOID ) '{' '}' ;
args: ( ITEM type (',' ITEM type)*)? ;

funcInvoc: ITEM '('argsInvoc')' ;
argsInvoc:( ITEM (',' ITEM)* )?;

var: 'var' ITEM type '=' (BOOL_VAL | STRING_RAW | NUMBER | ITEM | arrayInit);

type: ARRAY | STRING | NUM | BOOL;


arrayInit: '{' arrayInitElems '}' ;
arrayInitElems
    :
    (BOOL_VAL (',' BOOL_VAL)* )
    |
    (STRING_RAW (',' STRING_RAW)* )
    |
    (NUMBER (',' NUMBER)* )
    ;