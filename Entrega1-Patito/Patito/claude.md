# Patito — Compilador (Contexto del Proyecto)

> Proyecto individual del curso **TC3002B – Desarrollo de aplicaciones avanzadas de Ciencias Computacionales** (Módulo Compiladores), Tec de Monterrey, Gpo 503. Autora: Daniela Caiceros Flores (A00837181). Profesora: Elda Quiroga.

Compilador para el mini-lenguaje **Patito** (imperativo, procedural, clásico). Se construye de forma **incremental por entregas numeradas**, agregando una fase del compilador en cada una. Cada entrega construye estrictamente sobre la anterior.

---

## 1. Stack y entorno

- **Lenguaje de implementación:** Go 1.22.3 (en macOS)
- **Generador de scanner/parser:** ANTLR4 (algoritmo ALL(*), top-down)
- **Runtime ANTLR Go:** `github.com/antlr4-go/antlr/v4` v4.13.1
- **Gramática:** `gramatica.g4` (`grammar gramatica;`), lexer + parser combinados
- **Módulo Go:** `patito`, paquete del parser generado: `patito/parser`
- **Directorio raíz del proyecto:** `Entrega1-Patito/Patito/`

### Generar el parser desde la gramática
```bash
java -jar antlr-4.13.2-complete.jar \
  -Dlanguage=Go \
  -package parser \
  -o parser \
  gramatica.g4
```
Esto produce `parser/gramatica_lexer.go` y `parser/gramatica_parser.go` (más bases del Listener), integrados vía el módulo `patito/parser`.

### Compilar y ejecutar
```bash
go run . <archivo.ptto>
```
El programa lee el archivo, parsea desde la regla `programa`, recorre el árbol con el `PatitoListener` e imprime la fila de cuádruplos.

---

## 2. Estructura de archivos

| Archivo | Rol |
|---|---|
| `gramatica.g4` | Gramática ANTLR4 (lexer + parser) |
| `parser/` | Código Go generado por ANTLR (no editar a mano) |
| `main.go` | Entry point: lexer → parser → walk del listener → imprime cuádruplos |
| `listener.go` | `PatitoListener`: puntos neurálgicos, generación de cuádruplos, validación semántica |
| `cubo_semantico.go` | Cubo semántico (tipos resultantes de operadores) |
| `cuadruplos.go` | Pila genérica + estructura de Cuádruplos (Agregar, Len, Backpatch, Imprimir) |
| `direcciones.go` | Administrador de memoria virtual (rangos segmentados por scope y tipo) + helpers `segmentoDe`/`tipoDeDireccion` — Entrega 4 |
| `directorio.go` | Directorio de funciones (`FuncionInfo`) y tablas de variables (`Variable`) — Entrega 4 |
| `maquina_virtual.go` | Máquina Virtual: memoria de ejecución segmentada, registros de activación, intérprete de cuádruplos — Entrega 5 |
| `*.ptto` | Casos de prueba (incl. `test_vm_*.ptto` para la Máquina Virtual) |

---

## 3. El lenguaje Patito

### 3.1 Tipos
Tipos **declarables**: `entero`, `flotante`, `boolean`, `string`.
Tipo especial: `nula` (solo para funciones sin valor de retorno).

> Keyword canónico (implementado en `gramatica.g4`, `listener.go` y `direcciones.go`): **`boolean`**. Internamente, las relacionales producen un `entero` con valor `1` (verdadero) / `0` (falso); ese entero es asignable a una variable `boolean` (ver `compatibles` en `listener.go`). La Máquina Virtual evalúa la condición de `GotoF` como "verdadero si ≠ 0".

### 3.2 Reglas de tipos
- **Sin conversiones implícitas en asignación.** El tipo de la expresión debe coincidir exactamente con el de la variable.
- **Widening implícito solo en aritmética:** `entero + flotante → flotante`.
- **Strings:** solo `+` (concatenación) en aritmética; solo `==` y `!=` en relacionales (registrado en `cubo_semantico.go`).
- **Relacionales** (`>`, `<`, `==`, `!=`) devuelven `entero` (`1`/`0`), asignable a `boolean`. `>` y `<` solo para `entero`/`flotante`.
- No se permiten comparaciones entre tipos diferentes (el cubo no las registra).

### 3.3 Construcciones del lenguaje
- Declaración de variables: `vars : x , y : entero ;`
- Asignación: `x = expresion ;`
- Condicional: `si ( expr ) { ... } sino { ... } ;`
- Ciclo: `mientras ( expr ) haz { ... } ;`
- Impresión: `escribe ( expr | "letrero" , ... ) ;`
- Funciones: `nula | tipo  nombre ( params ) { vars : ... { cuerpo } } ;`
- Retorno: `regresa expresion ;` (obligatorio en funciones con tipo; prohibido en `nula`)
- Llamada: `nombre ( arg , ... ) ;` — una función con tipo puede usarse como operando dentro de una expresión.

### 3.4 Notas léxicas/sintácticas (ANTLR4)
- **Las keywords deben definirse ANTES de la regla `ID`** en el lexer (si no, `inicio` se reconocería como `ID`).
- **`CTE_FLOAT` antes de `CTE_ENT`** (longest match correcto para floats).
- Convención de nombres ANTLR4: **reglas del parser en minúsculas, tokens/lexer en MAYÚSCULAS**.
- `[` y `]` (`LBRACKET`/`RBRACKET`) son **tokens reales** en Patito, no metasintaxis.
- `vars :` se hizo parte **obligatoria** de `programa` (y `varsdec` opcional explícito en funciones) para resolver una ambigüedad entre `varsop` y `funcsop` vacíos consecutivos.

---

## 4. Arquitectura semántica y de código intermedio

### 4.1 Cubo semántico (`cubo_semantico.go`)
Mapa con clave `"tipoIzq,tipoDer,operador"` → tipo resultante. Inicializa aritméticos (`+ - * /`) y relacionales (`> < != ==`). Devuelve error para combinaciones inválidas.

### 4.2 Directorio de funciones y tablas de variables (`directorio.go`)
- **Una tabla de variables por scope de función.** En `listener.go`, `funcionActual *FuncionInfo` rastrea el scope: `nil` = global; distinto de `nil` = dentro de esa función.
- Estructura `FuncionInfo`: `{ nombre, tipoRetorno, dirInicio, dirRetorno, parametros []*Variable, variables map[string]*Variable }`. `dirRetorno` es una dirección **global** donde la función deja su valor de retorno (`-1` si es `nula`).
- Estructura `Variable`: `{ nombre, tipo, direccion }` (dirección virtual).
- **Scope:** `buscarVariable` consulta primero la tabla local (si `funcionActual != nil`), luego la global. Shadowing permitido entre scopes; unicidad obligatoria dentro del mismo scope.

### 4.3 Cuádruplos (`cuadruplos.go`)
Formato `(operador, operando_izq, operando_der, resultado)`. Fila secuencial con acceso por índice. Pilas usadas: **operandos, tipos, operadores, saltos (backpatch), retornos (inicio de ciclo)**.

Tipos de cuádruplo generados (operandos = **direcciones virtuales**, salvo el nombre de función en `ERA`/`GOSUB`):
`(op, dirIzq, dirDer, dirTemp)` aritmético/relacional · `(=, dirVal, _, dirVar)` · `(GotoF, dirCond, _, idxDest)` · `(GOTO, _, _, idxDest)` · `(PRINT, dirVal, _, _)` · `(ERA, func, _, _)` · `(PARAM, dirArg, _, dirParam)` · `(GOSUB, func, _, idxInicio)` · `(RETURN, dirVal, _, dirRetorno)` · `(ENDPROC, _, _, _)`.

> El `PARAM` apunta a la **dirección virtual local del parámetro** (no a `paramN`), de modo que la Máquina Virtual lo copia directo al registro de activación. El `RETURN` lleva la dirección virtual del valor y, en el cuarto campo, la **dirección global de retorno** de la función. Las funciones `nula` terminan con `ENDPROC`; las funciones con tipo deben usar `regresa` (que emite `RETURN`).

### 4.4 Algoritmos de traducción (en `listener.go`)
- **Expresiones aritméticas/relacionales:** en `Enter` de `exopc`/`teropc`/`opc` se empuja el operador; en `Exit` se extraen 2 operandos+tipos, se consulta el cubo y se emite el cuádruplo con temporal `tN`. La precedencia la garantiza la gramática (`* /` antes que `+ -`).
- **Temporales:** notación `tN` con contador incremental. Las pilas manejan la precedencia implícitamente, **sin fondos falsos explícitos**.
- **Asignación (`ExitAsigna`):** valida compatibilidad de tipos y emite `(=, valor, _, var)`.
- **Condicional (`si/sino`):** `ExitExpresion` emite `GotoF` pendiente; `EnterSinoop` emite `GOTO` y hace backpatch del `GotoF` al inicio del `sino`; `ExitSinoop` hace el backpatch final. (Técnica de **backpatch** con `pilaSaltos`.)
- **Ciclo (`mientras`):** `EnterCiclo` guarda el índice de inicio de la condición en `pilaRetornos`; `ExitCiclo` emite `GOTO` de regreso y backpatch del `GotoF` al final.
- **Impresión:** expresiones → `PRINT` en `ExitExpresion`; letreros → `PRINT` en `EnterLetreros` (en **Enter**, para preservar orden izquierda→derecha por recursión derecha).
- **Llamadas/funciones:** `EnterLlamada` emite `ERA func`; cada argumento → `PARAM dirArg, dirParam` (con validación de tipo contra la firma); `ExitLlamada` valida la aridad y emite `GOSUB func, idxInicio`. La declaración (`ExitFuncs`) cierra con `ENDPROC`. `dirInicio` se fija en `EnterFuncs` (params/vars no generan cuádruplos). Si la función tiene tipo, `ExitLlamada` copia `dirRetorno` a un temporal y lo deja en la pila de operandos, permitiendo usar la llamada como operando.
- **Retorno (`ExitRetorno`):** valida que esté dentro de una función con tipo y que la expresión sea compatible con `tipoRetorno`; emite `(RETURN, dirVal, _, dirRetorno)`.

> ANTLR4 garantiza el pareo Enter/Exit, por eso el rastreo de scope con `funcionActual` es confiable sin resets manuales.

### 4.5 Memoria virtual (`direcciones.go`) — Entrega 4
Direcciones virtuales con **rangos segmentados por scope y tipo**: **1000 direcciones por tipo por segmento**. El generador de cuádruplos usa direcciones virtuales en lugar de nombres. El tipo "constante" omite `boolean`.

Layout exacto (constantes en `direcciones.go`):

| Ámbito | entero | flotante | boolean | string |
|---|---|---|---|---|
| global | 1000 | 2000 | 3000 | 4000 |
| local | 5000 | 6000 | 7000 | 8000 |
| temporal | 9000 | 10000 | 11000 | 12000 |
| constante | 13000 | 14000 | — | 15000 |

`AdministradorMemoria` (`direcciones.go`):
- `Asignar(ambito, tipo) int` — entrega la siguiente dirección libre del segmento (`base + offset`).
- `Constante(valor, tipo) int` — asigna o **reutiliza** la dirección de un literal (deduplicación).
- `ResetLocal()` — reinicia contadores local+temporal al terminar una función (reutiliza el espacio).
- `segmentoDe(dir)` / `tipoDeDireccion(dir)` — funciones puras que, **solo a partir del número**, deducen ámbito y tipo de una dirección. Son la base de la indexación en ejecución.

---

## 5. Máquina Virtual y Memoria de Ejecución (`maquina_virtual.go`) — Entrega 5

La Máquina Virtual (VM) **consume la fila de cuádruplos y resuelve las direcciones virtuales en tiempo de ejecución**. Se invoca al final de `main.go` tras imprimir el código intermedio.

### 5.1 Estructuras de la Memoria de Ejecución

| Estructura | Rol | Métodos de acceso |
|---|---|---|
| `Memoria` (`celdas map[int]interface{}`) | Bloque de celdas indexado por **dirección virtual**. Aloja `int`, `float64`, `string`, `bool`. | `Guardar(dir, valor)`, `Leer(dir)` (aborta si se lee una celda no inicializada) |
| `RegistroActivacion` (`memoria *Memoria`, `regreso int`) | Memoria **local + temporal** de UNA invocación + el índice de cuádruplo al cual volver. | se crea con `NuevoRegistroActivacion()` |
| `MaquinaVirtual` | `global` y `constantes` (memorias compartidas), `pila []*RegistroActivacion`, `pendiente` (RA en construcción entre `ERA` y `GOSUB`), `ip` (apuntador de instrucción) | `Ejecutar()`, `leer(dir)`, `escribir(dir,val)`, `memoriaDe(dir)` |

Se usa **un mapa disperso** (`map[int]interface{}`) en vez de arreglos contiguos: el espacio virtual es grande (16 000 direcciones) pero un programa real usa pocas; el mapa solo materializa las celdas tocadas.

### 5.2 Versión gráfica del mapa de memoria

```
                MÁQUINA VIRTUAL
   ip ─► [ fila de cuádruplos (solo lectura) ]

   COMPARTIDAS                         POR LLAMADA (pila de RA)
   ┌──────────────────┐               ┌───────────────────────────┐
   │ global           │               │ pila[tope] ──► RA actual   │
   │  1000 entero     │               │   ┌─────────────────────┐ │
   │  2000 flotante   │   local/temp  │   │ memoria (local+temp)│ │
   │  3000 boolean    │ ◄──índices───►│   │  5000.. local       │ │
   │  4000 string     │    9000..     │   │  9000.. temporal    │ │
   ├──────────────────┤               │   │ regreso = idx cuad. │ │
   │ constante        │               │   └─────────────────────┘ │
   │ 13000/14000/15000│               │ pila[..]  RA anterior     │
   └──────────────────┘               │ pila[0]   RA del main     │
                                       │ pendiente ─► RA en ERA    │
                                       └───────────────────────────┘
```

El RA en `pila[0]` es el del **programa principal** (aloja sus temporales). Cada `GOSUB` apila un RA; cada `ENDPROC` lo desapila y restaura `ip = regreso`.

### 5.3 Cómo las Direcciones Virtuales indexan la memoria

La dirección virtual **es el índice y a la vez la etiqueta de tipo y ámbito**. `memoriaDe(dir)` despacha sin consultar tablas:

```
segmentoDe(dir):  1000–4999 → global  |  5000–8999 → local
                  9000–12999 → temporal | 13000–15999 → constante
```

- `global`/`constante` → memoria **compartida** de la VM.
- `local`/`temporal` → memoria del **RA en el tope de la pila**.

`tipoDeDireccion(dir)` calcula el tipo por el desplazamiento dentro del ámbito (`(dir−base)/1000` ⇒ 0=entero, 1=flotante, 2=boolean, 3=string). Con esto la VM hace *widening* entero→flotante al asignar y precarga cada constante con su tipo correcto. **Ningún cuádruplo carga nombres de variable: el número lo dice todo.**

### 5.4 Intérprete: opcodes soportados (`Ejecutar`)

`=` · `+ - * /` (entero/flotante con widening; `+` también concatena strings) · `> < == !=` (→ `1`/`0`; `==`/`!=` también para strings) · `GOTO` · `GotoF` (salta si la condición es `0`) · `PRINT` · `ERA` (crea RA `pendiente`) · `PARAM` (copia arg→parámetro en el RA pendiente) · `GOSUB` (apila RA, guarda `regreso`, salta a `idxInicio`) · `RETURN` (escribe el valor en la dirección global de retorno —leído antes de desapilar— y regresa al llamador) · `ENDPROC` (desapila RA, restaura `ip`).

**Cobertura: la VM interpreta el 100% de los opcodes que emite el compilador.**

---

## 6. Estado por entregas

- **Entrega 0:** Expresiones regulares (tokens) y CFG formal equivalente a los diagramas de sintaxis.
- **Entrega 1:** Scanner + parser con ANTLR4 (`gramatica.g4`, integración Go, `go run . <archivo>`). Test plan con 7 casos (`test1`–`test7`), incluye errores léxico y sintáctico.
- **Entrega 2:** Diseño de análisis semántico (cubo semántico, directorio de funciones, tablas de variables, 7 puntos neurálgicos documentados). Extensión del lenguaje con `boolean` y `string`.
- **Entrega 3:** Generación de cuádruplos (pilas, fila de cuádruplos, traducción de aritmética/relacional, asignación, `si/sino`, `mientras`, `escribe`, llamadas).
- **Entrega 4:** Memoria virtual (`direcciones.go`) con direcciones segmentadas; generador de cuádruplos actualizado a direcciones virtuales; mecánica de funciones (`ERA`, `PARAM`, `GOSUB`, `ENDPROC`); `directorio.go`, `listener.go`, `main.go` actualizados.
- **Entrega 5 (actual, completada):** Máquina Virtual (`maquina_virtual.go`): memoria de ejecución segmentada (`Memoria`), registros de activación (`RegistroActivacion`), intérprete de **todos** los opcodes. `PARAM` apunta a la dirección virtual del parámetro. Cubo extendido con strings. **Sentencia `regresa` agregada a la gramática** (`REGRESA` + regla `retorno`) con su opcode `RETURN`, completando las funciones **con valor de retorno**. Casos `test_vm_*.ptto`.

### 6.1 Casos de prueba de la VM
- `test_vm_aritmetica.ptto` — precedencia, paréntesis, mezcla entero/flotante (widening) y `escribe`.
- `test_vm_condicion.ptto` — `si/sino` con relacionales (`>`, `==`).
- `test_vm_ciclo.ptto` — `mientras` sumando 1..5 (= 15).
- `test_vm_funcion.ptto` — función `nula` con parámetro (`ERA/PARAM/GOSUB/ENDPROC`), demuestra aislamiento del RA frente a la global.
- `test_vm_retorno.ptto` — funciones **con tipo**: `doble(8)=16`, `mayor(3,10)=10` (`regresa` dentro de `si/sino`) y llamada anidada como argumento `doble(doble(8))=32` (`RETURN`).

---

## 7. Convenciones de trabajo

- Cada entrega referencia y extiende la anterior (ej. `A00837181_entrega4` → Entrega 5).
- **Documentación en español**, con terminología propia de compiladores.
- Para tareas de implementación: código completo bienvenido. Para tareas conceptuales/de aprendizaje: se prefiere descubrimiento guiado.
- No hardcodear: respetar el layout de memoria virtual y los rangos definidos en `direcciones.go`.

## 8. Repositorio
`https://github.com/DanielaCaiceros/Desarrollo_aplicaciones_avanzadas.git`