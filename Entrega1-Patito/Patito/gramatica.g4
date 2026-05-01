grammar gramatica;

//Reglas parser
programa : PROGRAMA ID SEMICOLON 
           VARS DPUNTOS varsop
           funcsop 
           INICIO cuerpo FIN ;

varsop : vars varsop
       |
       ;

vars : idop DPUNTOS tipo SEMICOLON ;
funcsop : funcs funcsop
        |
        ;
cuerpo: LCORCHETE estatutos RCORCHETE ;
estatutos: estatuto estatutos
         |
         ;
idop: ID COMMA idop
     | ID ;
tipo: ENTERO 
    | FLOTANTE ;
estatuto: asigna
            | condicion
            | ciclo
            | llamada SEMICOLON
            | imprime
            | LCORCHETE estatutos RCORCHETE
            ;
imprime: ESCRIBE LPAR explet RPAR SEMICOLON ;
explet: expresiones
      | letreros
      ;
letreros: LETRERO COMMA letreros
         | LETRERO ;
expresiones: expresion COMMA expresiones
           | expresion ;
asigna: ID ASIG expresion SEMICOLON ;
ciclo: MIENTRAS LPAR expresion RPAR HAZ cuerpo SEMICOLON;
condicion: SI LPAR expresion RPAR cuerpo sinoop SEMICOLON ;
sinoop: SINO cuerpo
       |
       ;
expresion: exp opc;
opc: MAYOR exp
   | MENOR exp
   | NEQ exp
   | EQ exp
   |
   ;
cte: CTE_ENT
   | CTE_FLOAT
   ;
exp: termino exopc ;
exopc: MAS termino exopc
     | MENOS termino exopc
     |
     ;
termino: factor teropc ;
teropc: MULT factor teropc
       | DIV factor teropc
       |
       ;
factor: ID
      | cte
      | LPAR expresion RPAR
      | MAS factor
      | MENOS factor
      | llamada
      ;
funcs: funcsopc ID LPAR funcr RPAR LCORCHETE varsdec cuerpo RCORCHETE ;
funcsopc: NULA
        | tipo
        ;
funcr: ID DPUNTOS tipo COMMA funcr
     | ID DPUNTOS tipo
     |
     ;
varsdec : VARS DPUNTOS varsop
        |
        ;
llamada: ID LPAR llamadaexp RPAR ;
llamadaexp: expresion COMMA llamadaexp
           | expresion
           |
           ;

//Reglas lexer
SEMICOLON : ';' ;
COMMA : ',' ;
DPUNTOS : ':' ;
LCORCHETE : '{' ;
RCORCHETE : '}' ;
LPAR : '(' ;
RPAR : ')' ;
LBRACKET : '[' ;
RBRACKET : ']' ;
ASIG : '=' ;
MENOR : '<' ;
MAYOR : '>' ;   
MAS : '+' ;
MENOS : '-' ;
MULT : '*' ;
DIV : '/' ;
NEQ : '!=' ;
EQ : '==' ;
CTE_FLOAT : [0-9]+ '.' [0-9]+ ;
CTE_ENT : [0-9]+ ;
PROGRAMA : 'programa' ;
INICIO   : 'inicio' ;
FIN : 'fin';
VARS: 'vars';
ENTERO: 'entero';
FLOTANTE: 'flotante';
ESCRIBE: 'escribe';
MIENTRAS: 'mientras';
HAZ: 'haz';
SI: 'si';
SINO: 'sino';
NULA: 'nula';
LETRERO : '"' ~["]* '"' ;  
ID : [a-zA-Z_][a-zA-Z0-9_]* ;
WS : [ \t\r\n]+ -> skip ;
COMMENT_LINE  : '//' [^\n]* '\n'  -> skip ;
COMMENT_BLOCK : '/*' .*? '*/'     -> skip ;


