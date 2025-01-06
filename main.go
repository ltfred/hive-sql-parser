package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/ltfred/hive-parse/parser"
)

func main() {
	charStream := antlr.NewInputStream(`create table test1
	(
	   c1 int COMMENT 'c1c1',
	   c2 string not null,
	   c3 boolean,
	   c4 double,
	   c5 date,
	   c6 timestamp,
	   c7 varchar(255)
	) comment '3434'`)

	lexer := parser.NewHplsqlLexer(charStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	par := parser.NewHplsqlParser(stream)

	ml := listener{}
	antlr.ParseTreeWalkerDefault.Walk(&ml, par.Create_table_stmt())
}

type listener struct{}

func (l listener) EnterEveryRule(ctx antlr.ParserRuleContext) {

	switch ctx.(type) {
	case *parser.Table_nameContext:
		fmt.Println("table name:", ctx.GetText())
	case *parser.Create_table_optionsContext:
		fmt.Println("table comment:", ctx.GetText())
	case *parser.Column_nameContext:
		fmt.Println("column:", ctx.GetText())
	case *parser.Create_table_column_commentContext:
		fmt.Println("comment:", ctx.GetText())
	case *parser.DtypeContext:
		fmt.Println("type:", ctx.GetText())
	case *parser.Dtype_attrContext:
		fmt.Println("attr:", ctx.GetText())
	case *parser.Dtype_lenContext:
		fmt.Println("len:", ctx.GetText())
	}
}
func (l listener) ExitEveryRule(ctx antlr.ParserRuleContext) {
}
func (l listener) VisitTerminal(node antlr.TerminalNode) {}
func (l listener) VisitErrorNode(node antlr.ErrorNode)   {}
