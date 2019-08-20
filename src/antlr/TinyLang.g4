grammar TinyLang;

WS : [ \t\r\n]+ -> skip ;
WHILE: 'while';
RETURN: 'return';
BREAK:'break';
CONTINUE: 'continue';
IF: 'if';
FOR: 'for';
ELSE: 'else';
SYS_FUNC: 'Add' | 'Len' | 'Output' | 'Input' ;
STRING : 'string';
VOID : 'void';
NUM : 'num';
BOOL: 'bool';
NULL: 'null';
NUMBER: [0-9]+;
TRUE: 'true';
FALSE: 'false';
BOOL_VAL: TRUE | FALSE;
STRING_RAW: '"' ( '\\"' | . )*? '"' ;
ITEM : [a-z][a-zA-Z0-9]* ;
SQ : '[]';
COMMENT : '/*' .*? '*/' -> skip ;
ARRAY: SQ (STRING | NUM | BOOL);
NUM_SIGN: '+' | '-' | '*' | '/' | '%';
BOOL_SIGN:  '==' | '!=' | '>' | '<';
BOOL_PL:  '&&' | '||';

file: funcInit* EOF;

funcInit: 'func' ITEM '(' funcArgs ')' funcReturnType '{' statementBody funcReturn?'}' ;
funcArg: ITEM variableType ;
funcArgs: ( funcArg (',' funcArg)*)? ;
funcReturn: RETURN (arrayElem | funcInvoc | val | expr | boolExpr) ;
funcReturnType: variableType | VOID;

funcInvoc: ( SYS_FUNC | ITEM ) '('funcArgsInvoc')' ;
funcArgsInvoc:( funcInvocArgs (',' funcInvocArgs)* )?;
funcInvocArgs : val | funcInvoc | arrayElem; // <-

updVariable: (ITEM | arrayElem) '=' (expr | funcInvoc | val | arrayInit | arrayElem  );
newVariable: 'var' ITEM variableType '=' (expr | val | arrayInit | funcInvoc );

val: TRUE | FALSE | STRING_RAW | NUMBER | ITEM ;

variableType: ARRAY | STRING | NUM | BOOL;

arrayElem: ITEM'['( NUMBER | ITEM | funcInvoc | expr)']';
arrayInit: arrayInitValue | arrayInitEmpty;
arrayInitEmpty: '[' (NUMBER | '..' )']' (NUM | STRING | BOOL) ;
arrayInitValue: '{' arrayInitElems '}' ;
arrayInitElems:
    (BOOL_VAL (',' BOOL_VAL)* )
    |
    (STRING_RAW (',' STRING_RAW)* )
    |
    (NUMBER (',' NUMBER)* );

expr:
       ((exprOperand) NUM_SIGN (exprOperand | expr))
       |
       ('(' (exprOperand | expr) NUM_SIGN (exprOperand | expr)')')
       ;
exprOperand: funcInvoc | NUMBER | ITEM | STRING_RAW  | arrayElem ;

boolExpr: boolExprSingle (boolExprSingleExtra)*;
boolExprSingleExtra: BOOL_PL boolExprSingle;
boolExprSingle:
        ((boolExprOperand) BOOL_SIGN (boolExprOperand | boolExpr))
        |
        ('(' (boolExprOperand | boolExpr) BOOL_SIGN (boolExprOperand | boolExpr)')')
        ;
boolExprOperand:  expr | exprOperand | TRUE | FALSE | arrayElem ;
breakOrContinue: BREAK | CONTINUE;
statementBody: (newVariable | updVariable | funcInvoc | ifElseSt | whileSt | forSt | breakOrContinue)*;
ifElseSt: ifSt elseIfSt* elseSt?;
ifSt : IF '('(TRUE | FALSE | boolExpr | ITEM | funcInvoc )')' '{' statementBody?'}' ;
elseIfSt: ELSE IF '('(TRUE | FALSE | boolExpr | ITEM | funcInvoc )')' '{'statementBody? '}';
elseSt: ELSE '{'statementBody? '}';

whileSt: WHILE '(' (TRUE | FALSE | boolExpr | ITEM | funcInvoc )')' '{' statementBody? '}';

forSt: FOR '(' updVariable ';' boolExpr ';' updVariable ')' '{' statementBody'}';

