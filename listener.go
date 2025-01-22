package hive_sql_parser

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/ltfred/hive-sql-parser/antlr_gen"
	"strconv"
	"strings"
)

type NullInt struct {
	Valid bool
	Value int
}

type Column struct {
	Name      string
	Type      string
	NotNull   bool
	Length    int
	Precision int
	Scale     NullInt
	Comment   string
}

type createTableStmtListener struct {
	antlr_gen.BaseHplsqlListener
	tableName        string
	tableComment     string
	tableColumns     []Column
	partitionColumns []Column
}

func (l *createTableStmtListener) EnterTable_name(ctx *antlr_gen.Table_nameContext) {
	l.tableName = ctx.GetText()
}

func (l *createTableStmtListener) EnterColumn_name(ctx *antlr_gen.Column_nameContext) {
	s := strings.TrimRight(strings.TrimLeft(ctx.GetText(), "`"), "`")
	l.tableColumns = append(l.tableColumns, Column{Name: s})
}

func (l *createTableStmtListener) EnterDtype(ctx *antlr_gen.DtypeContext) {
	l.tableColumns[len(l.tableColumns)-1].Type = ctx.GetText()
}

func (l *createTableStmtListener) EnterDtype_len(ctx *antlr_gen.Dtype_lenContext) {
	ls := strings.TrimLeft(strings.TrimRight(ctx.GetText(), ")"), "(")
	ll := strings.Split(ls, ",")
	switch len(ll) {
	case 1:
		length, _ := strconv.ParseInt(ll[0], 10, 64)
		l.tableColumns[len(l.tableColumns)-1].Length = int(length)
	case 2:
		precision, _ := strconv.ParseInt(ll[0], 10, 64)
		l.tableColumns[len(l.tableColumns)-1].Precision = int(precision)
		scale, _ := strconv.ParseInt(ll[1], 10, 64)
		l.tableColumns[len(l.tableColumns)-1].Scale = NullInt{Valid: true, Value: int(scale)}
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

func (l *createTableStmtListener) EnterCreate_table_options_hive_comment(ctx *antlr_gen.Create_table_options_hive_commentContext) {
	l.tableComment = strings.TrimRight(ctx.GetText()[8:], "'")
}

func (l *createTableStmtListener) EnterPartition_column_name(ctx *antlr_gen.Partition_column_nameContext) {
	l.partitionColumns = append(l.partitionColumns, Column{Name: ctx.GetText()})
}

func (l *createTableStmtListener) EnterPartition_dtype(ctx *antlr_gen.Partition_dtypeContext) {
	l.partitionColumns[len(l.partitionColumns)-1].Type = ctx.GetText()
}

func (l *createTableStmtListener) EnterPartition_dtype_len(ctx *antlr_gen.Partition_dtype_lenContext) {
	node, ok := ctx.GetChildren()[1].(*antlr.TerminalNodeImpl)
	if ok {
		length, _ := strconv.ParseInt(node.String(), 10, 64)
		l.partitionColumns[len(l.partitionColumns)-1].Length = int(length)
	}
}

func (l *createTableStmtListener) EnterCreate_table_hive_partition_column_comment(ctx *antlr_gen.Create_table_hive_partition_column_commentContext) {
	l.partitionColumns[len(l.partitionColumns)-1].Comment = strings.TrimRight(ctx.GetText()[8:], "'")
}

type dropTableStmtListener struct {
	antlr_gen.BaseHplsqlListener
	tableName string
}

func (l *dropTableStmtListener) EnterTable_name(ctx *antlr_gen.Table_nameContext) {
	l.tableName = ctx.GetText()
}
