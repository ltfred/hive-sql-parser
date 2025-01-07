package main

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/ltfred/hive-sql-parser/antlr_gen"
	"strconv"
	"strings"
)

type Column struct {
	Name    string
	Type    string
	NotNull bool
	Length  int
	Comment string
}

type createTableStmtListener struct {
	antlr_gen.BaseHplsqlListener
	tableName    string
	tableComment string
	tableColumns []Column
}

func (l *createTableStmtListener) EnterTable_name(ctx *antlr_gen.Table_nameContext) {
	l.tableName = ctx.GetText()
}

func (l *createTableStmtListener) EnterColumn_name(ctx *antlr_gen.Column_nameContext) {
	l.tableColumns = append(l.tableColumns, Column{Name: ctx.GetText()})
}

func (l *createTableStmtListener) EnterDtype(ctx *antlr_gen.DtypeContext) {
	l.tableColumns[len(l.tableColumns)-1].Type = ctx.GetText()
}

func (l *createTableStmtListener) EnterDtype_len(ctx *antlr_gen.Dtype_lenContext) {
	node, ok := ctx.GetChildren()[1].(*antlr.TerminalNodeImpl)
	if ok {
		length, _ := strconv.ParseInt(node.String(), 10, 64)
		l.tableColumns[len(l.tableColumns)-1].Length = int(length)
	}

}

func (l *createTableStmtListener) EnterDtype_attr(ctx *antlr_gen.Dtype_attrContext) {
	if strings.Contains(strings.ToLower(ctx.GetText()), "notnull") {
		l.tableColumns[len(l.tableColumns)-1].NotNull = true
	}
}

func (l *createTableStmtListener) EnterCreate_table_column_comment(ctx *antlr_gen.Create_table_column_commentContext) {
	l.tableColumns[len(l.tableColumns)-1].Comment = strings.TrimRight(ctx.GetText()[8:], "'")
}

func (l *createTableStmtListener) EnterCreate_table_options(ctx *antlr_gen.Create_table_optionsContext) {
	l.tableComment = strings.TrimRight(ctx.GetText()[8:], "'")
}
