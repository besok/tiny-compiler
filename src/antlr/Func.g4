grammar Func;

WS : [ \t\r\n]+ -> skip ;
WHILE: 'while';
RETURN: 'return';
IF: 'if';
FOR: 'for';
ELSE: 'else';
SYS_FUNC: 'Add' | 'Len' | 'Output' | 'Input' ;
STRING : 'string';
VOID : 'void';
NUM : 'num';
BOOL: 'bool';
NULL: 'null';
ITEM : [a-z][a-zA-Z0-9]* ;
TRUE: 'true';
FALSE: 'false';
BOOL_VAL: TRUE | FALSE;
NUMBER: [0-9]+;
STRING_RAW: '"' ( '\\"' | . )*? '"' ;
SQ : '[]';
COMMENT : '/*' .*? '*/' -> skip ;
ARRAY: SQ (STRING | NUM | BOOL);
NUM_SIGN: '+' | '-' | '*' | '/' | '%';
BOOL_SIGN:  '==' | '!=' | '>' | '<';

func: 'func' ITEM '(' args ')' ( type | VOID ) '{' commonBody return?'}' ;
args: ( ITEM type (',' ITEM type)*)? ;
return: RETURN (arrayElem | funcInvoc | ITEM | STRING_RAW | BOOL_VAL | NUMBER | expr | boolExpr) ;

funcInvoc: ( SYS_FUNC | ITEM ) '('argsInvoc')' ;
argsInvoc:( argInvoc (',' argInvoc)* )?;
argInvoc : NUMBER | STRING_RAW | BOOL_VAL | ITEM | funcInvoc | arrayElem;

updVar: (ITEM | arrayElem) '=' (expr | funcInvoc | BOOL_VAL | STRING_RAW | NUMBER | ITEM | arrayInit | arrayElem  );
newVar: 'var' ITEM type '=' (expr | BOOL_VAL | STRING_RAW | NUMBER | ITEM | arrayInit | funcInvoc );

type: ARRAY | STRING | NUM | BOOL;

arrayElem: ITEM'['( NUMBER | ITEM | funcInvoc | expr)']';
arrayInit: arrayInitVal | arrayInitEmp;
arrayInitEmp: '[' (NUMBER | '..' )']' (NUM | STRING | BOOL) ;
arrayInitVal: '{' arrayInitElems '}' ;
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
exprOp: funcInvoc | NUMBER | ITEM | STRING_RAW  | arrayElem ;

boolExpr:
        ((boolExprOp) BOOL_SIGN (boolExprOp | boolExpr))
        |
        ('(' (boolExprOp | boolExpr) BOOL_SIGN (boolExprOp | boolExpr)')')
        ;
boolExprOp:  expr | exprOp | BOOL_VAL | arrayElem ;


commonBody: (newVar | updVar | funcInvoc | ifElse | while | for)*;

ifElse: if elseIf* else?;
if : IF '('(BOOL_VAL | boolExpr | ITEM | funcInvoc )')' '{' commonBody?'}' ;
elseIf: ELSE IF '('(BOOL_VAL | boolExpr | ITEM | funcInvoc )')' '{'commonBody? '}';
else: ELSE '{'commonBody? '}';

while: WHILE '(' (BOOL_VAL | boolExpr | ITEM | funcInvoc )')' '{' commonBody? '}';

for: FOR '(' updVar ';' boolExpr ';' updVar ')' '{' commonBody'}';

file: func* EOF;