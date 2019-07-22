grammar Types;
import Words;

type: ARRAY | STRING | NUM | BOOL;

ARRAY: SQ (STRING | NUM | BOOL);
SQ : '[]';
WS : [ \t]+ -> skip ;