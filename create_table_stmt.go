package main

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/ltfred/hive-sql-parser/antlr_gen"
)

type CreateTableStmt struct {
	TableName    string
	TableComment string
	TableColumns []Column
}

func (parser *CreateTableStmt) Parse(sql string) Statement {
	stream := antlr.NewCommonTokenStream(antlr_gen.NewHplsqlLexer(antlr.NewInputStream(sql)), antlr.TokenDefaultChannel)
	par := antlr_gen.NewHplsqlParser(stream)
	listener := createTableStmtListener{}
	antlr.ParseTreeWalkerDefault.Walk(&listener, par.Create_table_stmt())

	return &CreateTableStmt{
		TableName:    listener.tableName,
		TableComment: listener.tableComment,
		TableColumns: listener.tableColumns,
	}
}

func (*CreateTableStmt) iStatement() {}
