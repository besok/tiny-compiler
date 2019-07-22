grammar Var;
import Types;

var: 'var' ITEM type '=' ITEM ;  // TODO add right part of var after =
WS : [ \t]+ -> skip ;

