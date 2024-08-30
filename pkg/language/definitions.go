package language

type program struct {
	Functions []*functionDefinition `@@*`
}

type functionDefinition struct {
	Name           string          `"func" @Ident`
	Parameters     []FunctionParam `"(" (@@ ("," @@)?)* ")"`
	OptionalReturn []Type          `( @@ | "(" @@ ("," @@)* ")" )?`
	Statements     []Statement     `"{" (@@)* "}"`
}

type FunctionParam struct {
	Name string `@Ident`
	Type string `@Ident`
}

type Type struct {
	Name string `@Ident`
}

type Statement struct {
	FunctionCall FunctionCall ` @@`
}

type FunctionCall struct {
	FunctionName string  `@Ident`
	Arguments    []Value `@@ ("," @@)*`
}

type Value struct {
	String *string `@String`
}
