grammar TinyLang;

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

file: funcInit* EOF;

funcInit: 'func' ITEM '(' funcArgs ')' ( variableType | VOID ) '{' statementBody funcReturn?'}' ;
funcArgs: ( ITEM variableType (',' ITEM variableType)*)? ;
funcReturn: RETURN (arrayElem | funcInvoc | ITEM | STRING_RAW | BOOL_VAL | NUMBER | expr | boolExpr) ;

funcInvoc: ( SYS_FUNC | ITEM ) '('funcArgsInvoc')' ;
funcArgsInvoc:( funcInvocArgs (',' funcInvocArgs)* )?;
funcInvocArgs : NUMBER | STRING_RAW | BOOL_VAL | ITEM | funcInvoc | arrayElem;

updVariable: (ITEM | arrayElem) '=' (expr | funcInvoc | BOOL_VAL | STRING_RAW | NUMBER | ITEM | arrayInit | arrayElem  );
newVariable: 'var' ITEM variableType '=' (expr | BOOL_VAL | STRING_RAW | NUMBER | ITEM | arrayInit | funcInvoc );

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

boolExpr:
        ((boolExprOperand) BOOL_SIGN (boolExprOperand | boolExpr))
        |
        ('(' (boolExprOperand | boolExpr) BOOL_SIGN (boolExprOperand | boolExpr)')')
        ;
boolExprOperand:  expr | exprOperand | BOOL_VAL | arrayElem ;


statementBody: (newVariable | updVariable | funcInvoc | ifElseSt | whileSt | forSt)*;

ifElseSt: ifSt elseIfSt* elseSt?;
ifSt : IF '('(BOOL_VAL | boolExpr | ITEM | funcInvoc )')' '{' statementBody?'}' ;
elseIfSt: ELSE IF '('(BOOL_VAL | boolExpr | ITEM | funcInvoc )')' '{'statementBody? '}';
elseSt: ELSE '{'statementBody? '}';

whileSt: WHILE '(' (BOOL_VAL | boolExpr | ITEM | funcInvoc )')' '{' statementBody? '}';

forSt: FOR '(' updVariable ';' boolExpr ';' updVariable ')' '{' statementBody'}';

