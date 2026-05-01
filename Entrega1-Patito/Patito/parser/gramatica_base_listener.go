// Code generated from gramatica.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // gramatica
import "github.com/antlr4-go/antlr/v4"

// BasegramaticaListener is a complete listener for a parse tree produced by gramaticaParser.
type BasegramaticaListener struct{}

var _ gramaticaListener = &BasegramaticaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasegramaticaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasegramaticaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasegramaticaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasegramaticaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrograma is called when production programa is entered.
func (s *BasegramaticaListener) EnterPrograma(ctx *ProgramaContext) {}

// ExitPrograma is called when production programa is exited.
func (s *BasegramaticaListener) ExitPrograma(ctx *ProgramaContext) {}

// EnterVarsop is called when production varsop is entered.
func (s *BasegramaticaListener) EnterVarsop(ctx *VarsopContext) {}

// ExitVarsop is called when production varsop is exited.
func (s *BasegramaticaListener) ExitVarsop(ctx *VarsopContext) {}

// EnterVars is called when production vars is entered.
func (s *BasegramaticaListener) EnterVars(ctx *VarsContext) {}

// ExitVars is called when production vars is exited.
func (s *BasegramaticaListener) ExitVars(ctx *VarsContext) {}

// EnterFuncsop is called when production funcsop is entered.
func (s *BasegramaticaListener) EnterFuncsop(ctx *FuncsopContext) {}

// ExitFuncsop is called when production funcsop is exited.
func (s *BasegramaticaListener) ExitFuncsop(ctx *FuncsopContext) {}

// EnterCuerpo is called when production cuerpo is entered.
func (s *BasegramaticaListener) EnterCuerpo(ctx *CuerpoContext) {}

// ExitCuerpo is called when production cuerpo is exited.
func (s *BasegramaticaListener) ExitCuerpo(ctx *CuerpoContext) {}

// EnterEstatutos is called when production estatutos is entered.
func (s *BasegramaticaListener) EnterEstatutos(ctx *EstatutosContext) {}

// ExitEstatutos is called when production estatutos is exited.
func (s *BasegramaticaListener) ExitEstatutos(ctx *EstatutosContext) {}

// EnterIdop is called when production idop is entered.
func (s *BasegramaticaListener) EnterIdop(ctx *IdopContext) {}

// ExitIdop is called when production idop is exited.
func (s *BasegramaticaListener) ExitIdop(ctx *IdopContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BasegramaticaListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BasegramaticaListener) ExitTipo(ctx *TipoContext) {}

// EnterEstatuto is called when production estatuto is entered.
func (s *BasegramaticaListener) EnterEstatuto(ctx *EstatutoContext) {}

// ExitEstatuto is called when production estatuto is exited.
func (s *BasegramaticaListener) ExitEstatuto(ctx *EstatutoContext) {}

// EnterImprime is called when production imprime is entered.
func (s *BasegramaticaListener) EnterImprime(ctx *ImprimeContext) {}

// ExitImprime is called when production imprime is exited.
func (s *BasegramaticaListener) ExitImprime(ctx *ImprimeContext) {}

// EnterExplet is called when production explet is entered.
func (s *BasegramaticaListener) EnterExplet(ctx *ExpletContext) {}

// ExitExplet is called when production explet is exited.
func (s *BasegramaticaListener) ExitExplet(ctx *ExpletContext) {}

// EnterLetreros is called when production letreros is entered.
func (s *BasegramaticaListener) EnterLetreros(ctx *LetrerosContext) {}

// ExitLetreros is called when production letreros is exited.
func (s *BasegramaticaListener) ExitLetreros(ctx *LetrerosContext) {}

// EnterExpresiones is called when production expresiones is entered.
func (s *BasegramaticaListener) EnterExpresiones(ctx *ExpresionesContext) {}

// ExitExpresiones is called when production expresiones is exited.
func (s *BasegramaticaListener) ExitExpresiones(ctx *ExpresionesContext) {}

// EnterAsigna is called when production asigna is entered.
func (s *BasegramaticaListener) EnterAsigna(ctx *AsignaContext) {}

// ExitAsigna is called when production asigna is exited.
func (s *BasegramaticaListener) ExitAsigna(ctx *AsignaContext) {}

// EnterCiclo is called when production ciclo is entered.
func (s *BasegramaticaListener) EnterCiclo(ctx *CicloContext) {}

// ExitCiclo is called when production ciclo is exited.
func (s *BasegramaticaListener) ExitCiclo(ctx *CicloContext) {}

// EnterCondicion is called when production condicion is entered.
func (s *BasegramaticaListener) EnterCondicion(ctx *CondicionContext) {}

// ExitCondicion is called when production condicion is exited.
func (s *BasegramaticaListener) ExitCondicion(ctx *CondicionContext) {}

// EnterSinoop is called when production sinoop is entered.
func (s *BasegramaticaListener) EnterSinoop(ctx *SinoopContext) {}

// ExitSinoop is called when production sinoop is exited.
func (s *BasegramaticaListener) ExitSinoop(ctx *SinoopContext) {}

// EnterExpresion is called when production expresion is entered.
func (s *BasegramaticaListener) EnterExpresion(ctx *ExpresionContext) {}

// ExitExpresion is called when production expresion is exited.
func (s *BasegramaticaListener) ExitExpresion(ctx *ExpresionContext) {}

// EnterOpc is called when production opc is entered.
func (s *BasegramaticaListener) EnterOpc(ctx *OpcContext) {}

// ExitOpc is called when production opc is exited.
func (s *BasegramaticaListener) ExitOpc(ctx *OpcContext) {}

// EnterCte is called when production cte is entered.
func (s *BasegramaticaListener) EnterCte(ctx *CteContext) {}

// ExitCte is called when production cte is exited.
func (s *BasegramaticaListener) ExitCte(ctx *CteContext) {}

// EnterExp is called when production exp is entered.
func (s *BasegramaticaListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BasegramaticaListener) ExitExp(ctx *ExpContext) {}

// EnterExopc is called when production exopc is entered.
func (s *BasegramaticaListener) EnterExopc(ctx *ExopcContext) {}

// ExitExopc is called when production exopc is exited.
func (s *BasegramaticaListener) ExitExopc(ctx *ExopcContext) {}

// EnterTermino is called when production termino is entered.
func (s *BasegramaticaListener) EnterTermino(ctx *TerminoContext) {}

// ExitTermino is called when production termino is exited.
func (s *BasegramaticaListener) ExitTermino(ctx *TerminoContext) {}

// EnterTeropc is called when production teropc is entered.
func (s *BasegramaticaListener) EnterTeropc(ctx *TeropcContext) {}

// ExitTeropc is called when production teropc is exited.
func (s *BasegramaticaListener) ExitTeropc(ctx *TeropcContext) {}

// EnterFactor is called when production factor is entered.
func (s *BasegramaticaListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BasegramaticaListener) ExitFactor(ctx *FactorContext) {}

// EnterFuncs is called when production funcs is entered.
func (s *BasegramaticaListener) EnterFuncs(ctx *FuncsContext) {}

// ExitFuncs is called when production funcs is exited.
func (s *BasegramaticaListener) ExitFuncs(ctx *FuncsContext) {}

// EnterFuncsopc is called when production funcsopc is entered.
func (s *BasegramaticaListener) EnterFuncsopc(ctx *FuncsopcContext) {}

// ExitFuncsopc is called when production funcsopc is exited.
func (s *BasegramaticaListener) ExitFuncsopc(ctx *FuncsopcContext) {}

// EnterFuncr is called when production funcr is entered.
func (s *BasegramaticaListener) EnterFuncr(ctx *FuncrContext) {}

// ExitFuncr is called when production funcr is exited.
func (s *BasegramaticaListener) ExitFuncr(ctx *FuncrContext) {}

// EnterVarsdec is called when production varsdec is entered.
func (s *BasegramaticaListener) EnterVarsdec(ctx *VarsdecContext) {}

// ExitVarsdec is called when production varsdec is exited.
func (s *BasegramaticaListener) ExitVarsdec(ctx *VarsdecContext) {}

// EnterLlamada is called when production llamada is entered.
func (s *BasegramaticaListener) EnterLlamada(ctx *LlamadaContext) {}

// ExitLlamada is called when production llamada is exited.
func (s *BasegramaticaListener) ExitLlamada(ctx *LlamadaContext) {}

// EnterLlamadaexp is called when production llamadaexp is entered.
func (s *BasegramaticaListener) EnterLlamadaexp(ctx *LlamadaexpContext) {}

// ExitLlamadaexp is called when production llamadaexp is exited.
func (s *BasegramaticaListener) ExitLlamadaexp(ctx *LlamadaexpContext) {}
