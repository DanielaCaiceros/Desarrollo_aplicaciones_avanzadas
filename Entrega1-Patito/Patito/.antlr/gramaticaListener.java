// Generated from /Users/danielacaiceros/Dev/HappyFaces/Desarrollo_aplicaciones_avanzadas/Entrega1-Patito/Patito/gramatica.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link gramaticaParser}.
 */
public interface gramaticaListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#programa}.
	 * @param ctx the parse tree
	 */
	void enterPrograma(gramaticaParser.ProgramaContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#programa}.
	 * @param ctx the parse tree
	 */
	void exitPrograma(gramaticaParser.ProgramaContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#varsop}.
	 * @param ctx the parse tree
	 */
	void enterVarsop(gramaticaParser.VarsopContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#varsop}.
	 * @param ctx the parse tree
	 */
	void exitVarsop(gramaticaParser.VarsopContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#vars}.
	 * @param ctx the parse tree
	 */
	void enterVars(gramaticaParser.VarsContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#vars}.
	 * @param ctx the parse tree
	 */
	void exitVars(gramaticaParser.VarsContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#funcsop}.
	 * @param ctx the parse tree
	 */
	void enterFuncsop(gramaticaParser.FuncsopContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#funcsop}.
	 * @param ctx the parse tree
	 */
	void exitFuncsop(gramaticaParser.FuncsopContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#cuerpo}.
	 * @param ctx the parse tree
	 */
	void enterCuerpo(gramaticaParser.CuerpoContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#cuerpo}.
	 * @param ctx the parse tree
	 */
	void exitCuerpo(gramaticaParser.CuerpoContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#estatutos}.
	 * @param ctx the parse tree
	 */
	void enterEstatutos(gramaticaParser.EstatutosContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#estatutos}.
	 * @param ctx the parse tree
	 */
	void exitEstatutos(gramaticaParser.EstatutosContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#idop}.
	 * @param ctx the parse tree
	 */
	void enterIdop(gramaticaParser.IdopContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#idop}.
	 * @param ctx the parse tree
	 */
	void exitIdop(gramaticaParser.IdopContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#tipo}.
	 * @param ctx the parse tree
	 */
	void enterTipo(gramaticaParser.TipoContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#tipo}.
	 * @param ctx the parse tree
	 */
	void exitTipo(gramaticaParser.TipoContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#estatuto}.
	 * @param ctx the parse tree
	 */
	void enterEstatuto(gramaticaParser.EstatutoContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#estatuto}.
	 * @param ctx the parse tree
	 */
	void exitEstatuto(gramaticaParser.EstatutoContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#retorno}.
	 * @param ctx the parse tree
	 */
	void enterRetorno(gramaticaParser.RetornoContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#retorno}.
	 * @param ctx the parse tree
	 */
	void exitRetorno(gramaticaParser.RetornoContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#imprime}.
	 * @param ctx the parse tree
	 */
	void enterImprime(gramaticaParser.ImprimeContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#imprime}.
	 * @param ctx the parse tree
	 */
	void exitImprime(gramaticaParser.ImprimeContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#explet}.
	 * @param ctx the parse tree
	 */
	void enterExplet(gramaticaParser.ExpletContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#explet}.
	 * @param ctx the parse tree
	 */
	void exitExplet(gramaticaParser.ExpletContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#letreros}.
	 * @param ctx the parse tree
	 */
	void enterLetreros(gramaticaParser.LetrerosContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#letreros}.
	 * @param ctx the parse tree
	 */
	void exitLetreros(gramaticaParser.LetrerosContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#expresiones}.
	 * @param ctx the parse tree
	 */
	void enterExpresiones(gramaticaParser.ExpresionesContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#expresiones}.
	 * @param ctx the parse tree
	 */
	void exitExpresiones(gramaticaParser.ExpresionesContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#asigna}.
	 * @param ctx the parse tree
	 */
	void enterAsigna(gramaticaParser.AsignaContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#asigna}.
	 * @param ctx the parse tree
	 */
	void exitAsigna(gramaticaParser.AsignaContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#ciclo}.
	 * @param ctx the parse tree
	 */
	void enterCiclo(gramaticaParser.CicloContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#ciclo}.
	 * @param ctx the parse tree
	 */
	void exitCiclo(gramaticaParser.CicloContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#condicion}.
	 * @param ctx the parse tree
	 */
	void enterCondicion(gramaticaParser.CondicionContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#condicion}.
	 * @param ctx the parse tree
	 */
	void exitCondicion(gramaticaParser.CondicionContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#sinoop}.
	 * @param ctx the parse tree
	 */
	void enterSinoop(gramaticaParser.SinoopContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#sinoop}.
	 * @param ctx the parse tree
	 */
	void exitSinoop(gramaticaParser.SinoopContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterExpresion(gramaticaParser.ExpresionContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitExpresion(gramaticaParser.ExpresionContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#opc}.
	 * @param ctx the parse tree
	 */
	void enterOpc(gramaticaParser.OpcContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#opc}.
	 * @param ctx the parse tree
	 */
	void exitOpc(gramaticaParser.OpcContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#cte}.
	 * @param ctx the parse tree
	 */
	void enterCte(gramaticaParser.CteContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#cte}.
	 * @param ctx the parse tree
	 */
	void exitCte(gramaticaParser.CteContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#exp}.
	 * @param ctx the parse tree
	 */
	void enterExp(gramaticaParser.ExpContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#exp}.
	 * @param ctx the parse tree
	 */
	void exitExp(gramaticaParser.ExpContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#exopc}.
	 * @param ctx the parse tree
	 */
	void enterExopc(gramaticaParser.ExopcContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#exopc}.
	 * @param ctx the parse tree
	 */
	void exitExopc(gramaticaParser.ExopcContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#termino}.
	 * @param ctx the parse tree
	 */
	void enterTermino(gramaticaParser.TerminoContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#termino}.
	 * @param ctx the parse tree
	 */
	void exitTermino(gramaticaParser.TerminoContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#teropc}.
	 * @param ctx the parse tree
	 */
	void enterTeropc(gramaticaParser.TeropcContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#teropc}.
	 * @param ctx the parse tree
	 */
	void exitTeropc(gramaticaParser.TeropcContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#factor}.
	 * @param ctx the parse tree
	 */
	void enterFactor(gramaticaParser.FactorContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#factor}.
	 * @param ctx the parse tree
	 */
	void exitFactor(gramaticaParser.FactorContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#funcs}.
	 * @param ctx the parse tree
	 */
	void enterFuncs(gramaticaParser.FuncsContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#funcs}.
	 * @param ctx the parse tree
	 */
	void exitFuncs(gramaticaParser.FuncsContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#funcsopc}.
	 * @param ctx the parse tree
	 */
	void enterFuncsopc(gramaticaParser.FuncsopcContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#funcsopc}.
	 * @param ctx the parse tree
	 */
	void exitFuncsopc(gramaticaParser.FuncsopcContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#funcr}.
	 * @param ctx the parse tree
	 */
	void enterFuncr(gramaticaParser.FuncrContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#funcr}.
	 * @param ctx the parse tree
	 */
	void exitFuncr(gramaticaParser.FuncrContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#varsdec}.
	 * @param ctx the parse tree
	 */
	void enterVarsdec(gramaticaParser.VarsdecContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#varsdec}.
	 * @param ctx the parse tree
	 */
	void exitVarsdec(gramaticaParser.VarsdecContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#llamada}.
	 * @param ctx the parse tree
	 */
	void enterLlamada(gramaticaParser.LlamadaContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#llamada}.
	 * @param ctx the parse tree
	 */
	void exitLlamada(gramaticaParser.LlamadaContext ctx);
	/**
	 * Enter a parse tree produced by {@link gramaticaParser#llamadaexp}.
	 * @param ctx the parse tree
	 */
	void enterLlamadaexp(gramaticaParser.LlamadaexpContext ctx);
	/**
	 * Exit a parse tree produced by {@link gramaticaParser#llamadaexp}.
	 * @param ctx the parse tree
	 */
	void exitLlamadaexp(gramaticaParser.LlamadaexpContext ctx);
}