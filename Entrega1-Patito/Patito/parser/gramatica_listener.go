// Code generated from gramatica.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // gramatica
import "github.com/antlr4-go/antlr/v4"

// gramaticaListener is a complete listener for a parse tree produced by gramaticaParser.
type gramaticaListener interface {
	antlr.ParseTreeListener

	// EnterPrograma is called when entering the programa production.
	EnterPrograma(c *ProgramaContext)

	// EnterVarsop is called when entering the varsop production.
	EnterVarsop(c *VarsopContext)

	// EnterVars is called when entering the vars production.
	EnterVars(c *VarsContext)

	// EnterFuncsop is called when entering the funcsop production.
	EnterFuncsop(c *FuncsopContext)

	// EnterCuerpo is called when entering the cuerpo production.
	EnterCuerpo(c *CuerpoContext)

	// EnterEstatutos is called when entering the estatutos production.
	EnterEstatutos(c *EstatutosContext)

	// EnterIdop is called when entering the idop production.
	EnterIdop(c *IdopContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterEstatuto is called when entering the estatuto production.
	EnterEstatuto(c *EstatutoContext)

	// EnterImprime is called when entering the imprime production.
	EnterImprime(c *ImprimeContext)

	// EnterExplet is called when entering the explet production.
	EnterExplet(c *ExpletContext)

	// EnterLetreros is called when entering the letreros production.
	EnterLetreros(c *LetrerosContext)

	// EnterExpresiones is called when entering the expresiones production.
	EnterExpresiones(c *ExpresionesContext)

	// EnterAsigna is called when entering the asigna production.
	EnterAsigna(c *AsignaContext)

	// EnterCiclo is called when entering the ciclo production.
	EnterCiclo(c *CicloContext)

	// EnterCondicion is called when entering the condicion production.
	EnterCondicion(c *CondicionContext)

	// EnterSinoop is called when entering the sinoop production.
	EnterSinoop(c *SinoopContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// EnterOpc is called when entering the opc production.
	EnterOpc(c *OpcContext)

	// EnterCte is called when entering the cte production.
	EnterCte(c *CteContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterExopc is called when entering the exopc production.
	EnterExopc(c *ExopcContext)

	// EnterTermino is called when entering the termino production.
	EnterTermino(c *TerminoContext)

	// EnterTeropc is called when entering the teropc production.
	EnterTeropc(c *TeropcContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterFuncs is called when entering the funcs production.
	EnterFuncs(c *FuncsContext)

	// EnterFuncsopc is called when entering the funcsopc production.
	EnterFuncsopc(c *FuncsopcContext)

	// EnterFuncr is called when entering the funcr production.
	EnterFuncr(c *FuncrContext)

	// EnterVarsdec is called when entering the varsdec production.
	EnterVarsdec(c *VarsdecContext)

	// EnterLlamada is called when entering the llamada production.
	EnterLlamada(c *LlamadaContext)

	// EnterLlamadaexp is called when entering the llamadaexp production.
	EnterLlamadaexp(c *LlamadaexpContext)

	// ExitPrograma is called when exiting the programa production.
	ExitPrograma(c *ProgramaContext)

	// ExitVarsop is called when exiting the varsop production.
	ExitVarsop(c *VarsopContext)

	// ExitVars is called when exiting the vars production.
	ExitVars(c *VarsContext)

	// ExitFuncsop is called when exiting the funcsop production.
	ExitFuncsop(c *FuncsopContext)

	// ExitCuerpo is called when exiting the cuerpo production.
	ExitCuerpo(c *CuerpoContext)

	// ExitEstatutos is called when exiting the estatutos production.
	ExitEstatutos(c *EstatutosContext)

	// ExitIdop is called when exiting the idop production.
	ExitIdop(c *IdopContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitEstatuto is called when exiting the estatuto production.
	ExitEstatuto(c *EstatutoContext)

	// ExitImprime is called when exiting the imprime production.
	ExitImprime(c *ImprimeContext)

	// ExitExplet is called when exiting the explet production.
	ExitExplet(c *ExpletContext)

	// ExitLetreros is called when exiting the letreros production.
	ExitLetreros(c *LetrerosContext)

	// ExitExpresiones is called when exiting the expresiones production.
	ExitExpresiones(c *ExpresionesContext)

	// ExitAsigna is called when exiting the asigna production.
	ExitAsigna(c *AsignaContext)

	// ExitCiclo is called when exiting the ciclo production.
	ExitCiclo(c *CicloContext)

	// ExitCondicion is called when exiting the condicion production.
	ExitCondicion(c *CondicionContext)

	// ExitSinoop is called when exiting the sinoop production.
	ExitSinoop(c *SinoopContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)

	// ExitOpc is called when exiting the opc production.
	ExitOpc(c *OpcContext)

	// ExitCte is called when exiting the cte production.
	ExitCte(c *CteContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitExopc is called when exiting the exopc production.
	ExitExopc(c *ExopcContext)

	// ExitTermino is called when exiting the termino production.
	ExitTermino(c *TerminoContext)

	// ExitTeropc is called when exiting the teropc production.
	ExitTeropc(c *TeropcContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitFuncs is called when exiting the funcs production.
	ExitFuncs(c *FuncsContext)

	// ExitFuncsopc is called when exiting the funcsopc production.
	ExitFuncsopc(c *FuncsopcContext)

	// ExitFuncr is called when exiting the funcr production.
	ExitFuncr(c *FuncrContext)

	// ExitVarsdec is called when exiting the varsdec production.
	ExitVarsdec(c *VarsdecContext)

	// ExitLlamada is called when exiting the llamada production.
	ExitLlamada(c *LlamadaContext)

	// ExitLlamadaexp is called when exiting the llamadaexp production.
	ExitLlamadaexp(c *LlamadaexpContext)
}
