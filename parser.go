package hive_sql_parser

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/ltfred/hive-sql-parser/antlr_gen"
)

type Statement interface {
	iStatement()
}

type Parser interface {
	Parse(sql string) Statement
}

func (*DropTableStmt) iStatement() {}

func (*CreateTableStmt) iStatement() {}

func newParser(sql string) *antlr_gen.HplsqlParser {
	stream := antlr.NewCommonTokenStream(antlr_gen.NewHplsqlLexer(antlr.NewInputStream(sql)), antlr.TokenDefaultChannel)
	return antlr_gen.NewHplsqlParser(stream)
}

type CreateTableStmt struct {
	TableName        string
	TableComment     string
	TableColumns     []Column
	PartitionColumns []Column
}

type createTableStmtParser struct {
}

func NewCreatTableStmtParser() Parser {
	return &createTableStmtParser{}
}

func (parser *createTableStmtParser) Parse(sql string) Statement {
	listener := createTableStmtListener{}
	antlr.ParseTreeWalkerDefault.Walk(&listener, newParser(sql).Create_table_stmt())

	return &CreateTableStmt{
		TableName:        listener.tableName,
		TableComment:     listener.tableComment,
		TableColumns:     listener.tableColumns,
		PartitionColumns: listener.partitionColumns,
	}
}

type DropTableStmt struct {
	TableName string
}

type dropTableStmtParser struct {
}

func NewDropTableStmtParser() Parser {
	return &dropTableStmtParser{}
}

func (parser *dropTableStmtParser) Parse(sql string) Statement {
	listener := dropTableStmtListener{}
	antlr.ParseTreeWalkerDefault.Walk(&listener, newParser(sql).Drop_stmt())

	return &DropTableStmt{
		TableName: listener.tableName,
	}
}

type AlterTableStmt struct {
	TableName string
}
