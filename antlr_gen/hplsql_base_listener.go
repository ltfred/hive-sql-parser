// Code generated from Hplsql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlr_gen // Hplsql

import "github.com/antlr4-go/antlr/v4"

// BaseHplsqlListener is a complete listener for a parse tree produced by HplsqlParser.
type BaseHplsqlListener struct{}

var _ HplsqlListener = &BaseHplsqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHplsqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHplsqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHplsqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHplsqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseHplsqlListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseHplsqlListener) ExitProgram(ctx *ProgramContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseHplsqlListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseHplsqlListener) ExitBlock(ctx *BlockContext) {}

// EnterBegin_end_block is called when production begin_end_block is entered.
func (s *BaseHplsqlListener) EnterBegin_end_block(ctx *Begin_end_blockContext) {}

// ExitBegin_end_block is called when production begin_end_block is exited.
func (s *BaseHplsqlListener) ExitBegin_end_block(ctx *Begin_end_blockContext) {}

// EnterSingle_block_stmt is called when production single_block_stmt is entered.
func (s *BaseHplsqlListener) EnterSingle_block_stmt(ctx *Single_block_stmtContext) {}

// ExitSingle_block_stmt is called when production single_block_stmt is exited.
func (s *BaseHplsqlListener) ExitSingle_block_stmt(ctx *Single_block_stmtContext) {}

// EnterBlock_end is called when production block_end is entered.
func (s *BaseHplsqlListener) EnterBlock_end(ctx *Block_endContext) {}

// ExitBlock_end is called when production block_end is exited.
func (s *BaseHplsqlListener) ExitBlock_end(ctx *Block_endContext) {}

// EnterProc_block is called when production proc_block is entered.
func (s *BaseHplsqlListener) EnterProc_block(ctx *Proc_blockContext) {}

// ExitProc_block is called when production proc_block is exited.
func (s *BaseHplsqlListener) ExitProc_block(ctx *Proc_blockContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseHplsqlListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseHplsqlListener) ExitStmt(ctx *StmtContext) {}

// EnterSemicolon_stmt is called when production semicolon_stmt is entered.
func (s *BaseHplsqlListener) EnterSemicolon_stmt(ctx *Semicolon_stmtContext) {}

// ExitSemicolon_stmt is called when production semicolon_stmt is exited.
func (s *BaseHplsqlListener) ExitSemicolon_stmt(ctx *Semicolon_stmtContext) {}

// EnterException_block is called when production exception_block is entered.
func (s *BaseHplsqlListener) EnterException_block(ctx *Exception_blockContext) {}

// ExitException_block is called when production exception_block is exited.
func (s *BaseHplsqlListener) ExitException_block(ctx *Exception_blockContext) {}

// EnterException_block_item is called when production exception_block_item is entered.
func (s *BaseHplsqlListener) EnterException_block_item(ctx *Exception_block_itemContext) {}

// ExitException_block_item is called when production exception_block_item is exited.
func (s *BaseHplsqlListener) ExitException_block_item(ctx *Exception_block_itemContext) {}

// EnterNull_stmt is called when production null_stmt is entered.
func (s *BaseHplsqlListener) EnterNull_stmt(ctx *Null_stmtContext) {}

// ExitNull_stmt is called when production null_stmt is exited.
func (s *BaseHplsqlListener) ExitNull_stmt(ctx *Null_stmtContext) {}

// EnterExpr_stmt is called when production expr_stmt is entered.
func (s *BaseHplsqlListener) EnterExpr_stmt(ctx *Expr_stmtContext) {}

// ExitExpr_stmt is called when production expr_stmt is exited.
func (s *BaseHplsqlListener) ExitExpr_stmt(ctx *Expr_stmtContext) {}

// EnterAssignment_stmt is called when production assignment_stmt is entered.
func (s *BaseHplsqlListener) EnterAssignment_stmt(ctx *Assignment_stmtContext) {}

// ExitAssignment_stmt is called when production assignment_stmt is exited.
func (s *BaseHplsqlListener) ExitAssignment_stmt(ctx *Assignment_stmtContext) {}

// EnterAssignment_stmt_item is called when production assignment_stmt_item is entered.
func (s *BaseHplsqlListener) EnterAssignment_stmt_item(ctx *Assignment_stmt_itemContext) {}

// ExitAssignment_stmt_item is called when production assignment_stmt_item is exited.
func (s *BaseHplsqlListener) ExitAssignment_stmt_item(ctx *Assignment_stmt_itemContext) {}

// EnterAssignment_stmt_single_item is called when production assignment_stmt_single_item is entered.
func (s *BaseHplsqlListener) EnterAssignment_stmt_single_item(ctx *Assignment_stmt_single_itemContext) {
}

// ExitAssignment_stmt_single_item is called when production assignment_stmt_single_item is exited.
func (s *BaseHplsqlListener) ExitAssignment_stmt_single_item(ctx *Assignment_stmt_single_itemContext) {
}

// EnterAssignment_stmt_collection_item is called when production assignment_stmt_collection_item is entered.
func (s *BaseHplsqlListener) EnterAssignment_stmt_collection_item(ctx *Assignment_stmt_collection_itemContext) {
}

// ExitAssignment_stmt_collection_item is called when production assignment_stmt_collection_item is exited.
func (s *BaseHplsqlListener) ExitAssignment_stmt_collection_item(ctx *Assignment_stmt_collection_itemContext) {
}

// EnterAssignment_stmt_multiple_item is called when production assignment_stmt_multiple_item is entered.
func (s *BaseHplsqlListener) EnterAssignment_stmt_multiple_item(ctx *Assignment_stmt_multiple_itemContext) {
}

// ExitAssignment_stmt_multiple_item is called when production assignment_stmt_multiple_item is exited.
func (s *BaseHplsqlListener) ExitAssignment_stmt_multiple_item(ctx *Assignment_stmt_multiple_itemContext) {
}

// EnterAssignment_stmt_select_item is called when production assignment_stmt_select_item is entered.
func (s *BaseHplsqlListener) EnterAssignment_stmt_select_item(ctx *Assignment_stmt_select_itemContext) {
}

// ExitAssignment_stmt_select_item is called when production assignment_stmt_select_item is exited.
func (s *BaseHplsqlListener) ExitAssignment_stmt_select_item(ctx *Assignment_stmt_select_itemContext) {
}

// EnterAllocate_cursor_stmt is called when production allocate_cursor_stmt is entered.
func (s *BaseHplsqlListener) EnterAllocate_cursor_stmt(ctx *Allocate_cursor_stmtContext) {}

// ExitAllocate_cursor_stmt is called when production allocate_cursor_stmt is exited.
func (s *BaseHplsqlListener) ExitAllocate_cursor_stmt(ctx *Allocate_cursor_stmtContext) {}

// EnterAssociate_locator_stmt is called when production associate_locator_stmt is entered.
func (s *BaseHplsqlListener) EnterAssociate_locator_stmt(ctx *Associate_locator_stmtContext) {}

// ExitAssociate_locator_stmt is called when production associate_locator_stmt is exited.
func (s *BaseHplsqlListener) ExitAssociate_locator_stmt(ctx *Associate_locator_stmtContext) {}

// EnterBegin_transaction_stmt is called when production begin_transaction_stmt is entered.
func (s *BaseHplsqlListener) EnterBegin_transaction_stmt(ctx *Begin_transaction_stmtContext) {}

// ExitBegin_transaction_stmt is called when production begin_transaction_stmt is exited.
func (s *BaseHplsqlListener) ExitBegin_transaction_stmt(ctx *Begin_transaction_stmtContext) {}

// EnterBreak_stmt is called when production break_stmt is entered.
func (s *BaseHplsqlListener) EnterBreak_stmt(ctx *Break_stmtContext) {}

// ExitBreak_stmt is called when production break_stmt is exited.
func (s *BaseHplsqlListener) ExitBreak_stmt(ctx *Break_stmtContext) {}

// EnterCall_stmt is called when production call_stmt is entered.
func (s *BaseHplsqlListener) EnterCall_stmt(ctx *Call_stmtContext) {}

// ExitCall_stmt is called when production call_stmt is exited.
func (s *BaseHplsqlListener) ExitCall_stmt(ctx *Call_stmtContext) {}

// EnterDeclare_stmt is called when production declare_stmt is entered.
func (s *BaseHplsqlListener) EnterDeclare_stmt(ctx *Declare_stmtContext) {}

// ExitDeclare_stmt is called when production declare_stmt is exited.
func (s *BaseHplsqlListener) ExitDeclare_stmt(ctx *Declare_stmtContext) {}

// EnterDeclare_block is called when production declare_block is entered.
func (s *BaseHplsqlListener) EnterDeclare_block(ctx *Declare_blockContext) {}

// ExitDeclare_block is called when production declare_block is exited.
func (s *BaseHplsqlListener) ExitDeclare_block(ctx *Declare_blockContext) {}

// EnterDeclare_block_inplace is called when production declare_block_inplace is entered.
func (s *BaseHplsqlListener) EnterDeclare_block_inplace(ctx *Declare_block_inplaceContext) {}

// ExitDeclare_block_inplace is called when production declare_block_inplace is exited.
func (s *BaseHplsqlListener) ExitDeclare_block_inplace(ctx *Declare_block_inplaceContext) {}

// EnterDeclare_stmt_item is called when production declare_stmt_item is entered.
func (s *BaseHplsqlListener) EnterDeclare_stmt_item(ctx *Declare_stmt_itemContext) {}

// ExitDeclare_stmt_item is called when production declare_stmt_item is exited.
func (s *BaseHplsqlListener) ExitDeclare_stmt_item(ctx *Declare_stmt_itemContext) {}

// EnterDeclare_var_item is called when production declare_var_item is entered.
func (s *BaseHplsqlListener) EnterDeclare_var_item(ctx *Declare_var_itemContext) {}

// ExitDeclare_var_item is called when production declare_var_item is exited.
func (s *BaseHplsqlListener) ExitDeclare_var_item(ctx *Declare_var_itemContext) {}

// EnterDeclare_condition_item is called when production declare_condition_item is entered.
func (s *BaseHplsqlListener) EnterDeclare_condition_item(ctx *Declare_condition_itemContext) {}

// ExitDeclare_condition_item is called when production declare_condition_item is exited.
func (s *BaseHplsqlListener) ExitDeclare_condition_item(ctx *Declare_condition_itemContext) {}

// EnterDeclare_cursor_item is called when production declare_cursor_item is entered.
func (s *BaseHplsqlListener) EnterDeclare_cursor_item(ctx *Declare_cursor_itemContext) {}

// ExitDeclare_cursor_item is called when production declare_cursor_item is exited.
func (s *BaseHplsqlListener) ExitDeclare_cursor_item(ctx *Declare_cursor_itemContext) {}

// EnterCursor_with_return is called when production cursor_with_return is entered.
func (s *BaseHplsqlListener) EnterCursor_with_return(ctx *Cursor_with_returnContext) {}

// ExitCursor_with_return is called when production cursor_with_return is exited.
func (s *BaseHplsqlListener) ExitCursor_with_return(ctx *Cursor_with_returnContext) {}

// EnterCursor_without_return is called when production cursor_without_return is entered.
func (s *BaseHplsqlListener) EnterCursor_without_return(ctx *Cursor_without_returnContext) {}

// ExitCursor_without_return is called when production cursor_without_return is exited.
func (s *BaseHplsqlListener) ExitCursor_without_return(ctx *Cursor_without_returnContext) {}

// EnterDeclare_handler_item is called when production declare_handler_item is entered.
func (s *BaseHplsqlListener) EnterDeclare_handler_item(ctx *Declare_handler_itemContext) {}

// ExitDeclare_handler_item is called when production declare_handler_item is exited.
func (s *BaseHplsqlListener) ExitDeclare_handler_item(ctx *Declare_handler_itemContext) {}

// EnterDeclare_temporary_table_item is called when production declare_temporary_table_item is entered.
func (s *BaseHplsqlListener) EnterDeclare_temporary_table_item(ctx *Declare_temporary_table_itemContext) {
}

// ExitDeclare_temporary_table_item is called when production declare_temporary_table_item is exited.
func (s *BaseHplsqlListener) ExitDeclare_temporary_table_item(ctx *Declare_temporary_table_itemContext) {
}

// EnterCreate_table_stmt is called when production create_table_stmt is entered.
func (s *BaseHplsqlListener) EnterCreate_table_stmt(ctx *Create_table_stmtContext) {}

// ExitCreate_table_stmt is called when production create_table_stmt is exited.
func (s *BaseHplsqlListener) ExitCreate_table_stmt(ctx *Create_table_stmtContext) {}

// EnterCreate_local_temp_table_stmt is called when production create_local_temp_table_stmt is entered.
func (s *BaseHplsqlListener) EnterCreate_local_temp_table_stmt(ctx *Create_local_temp_table_stmtContext) {
}

// ExitCreate_local_temp_table_stmt is called when production create_local_temp_table_stmt is exited.
func (s *BaseHplsqlListener) ExitCreate_local_temp_table_stmt(ctx *Create_local_temp_table_stmtContext) {
}

// EnterCreate_table_definition is called when production create_table_definition is entered.
func (s *BaseHplsqlListener) EnterCreate_table_definition(ctx *Create_table_definitionContext) {}

// ExitCreate_table_definition is called when production create_table_definition is exited.
func (s *BaseHplsqlListener) ExitCreate_table_definition(ctx *Create_table_definitionContext) {}

// EnterCreate_table_column_comment is called when production create_table_column_comment is entered.
func (s *BaseHplsqlListener) EnterCreate_table_column_comment(ctx *Create_table_column_commentContext) {
}

// ExitCreate_table_column_comment is called when production create_table_column_comment is exited.
func (s *BaseHplsqlListener) ExitCreate_table_column_comment(ctx *Create_table_column_commentContext) {
}

// EnterCreate_table_columns is called when production create_table_columns is entered.
func (s *BaseHplsqlListener) EnterCreate_table_columns(ctx *Create_table_columnsContext) {}

// ExitCreate_table_columns is called when production create_table_columns is exited.
func (s *BaseHplsqlListener) ExitCreate_table_columns(ctx *Create_table_columnsContext) {}

// EnterCreate_table_columns_item is called when production create_table_columns_item is entered.
func (s *BaseHplsqlListener) EnterCreate_table_columns_item(ctx *Create_table_columns_itemContext) {}

// ExitCreate_table_columns_item is called when production create_table_columns_item is exited.
func (s *BaseHplsqlListener) ExitCreate_table_columns_item(ctx *Create_table_columns_itemContext) {}

// EnterCreate_table_type_stmt is called when production create_table_type_stmt is entered.
func (s *BaseHplsqlListener) EnterCreate_table_type_stmt(ctx *Create_table_type_stmtContext) {}

// ExitCreate_table_type_stmt is called when production create_table_type_stmt is exited.
func (s *BaseHplsqlListener) ExitCreate_table_type_stmt(ctx *Create_table_type_stmtContext) {}

// EnterTbl_type is called when production tbl_type is entered.
func (s *BaseHplsqlListener) EnterTbl_type(ctx *Tbl_typeContext) {}

// ExitTbl_type is called when production tbl_type is exited.
func (s *BaseHplsqlListener) ExitTbl_type(ctx *Tbl_typeContext) {}

// EnterSql_type is called when production sql_type is entered.
func (s *BaseHplsqlListener) EnterSql_type(ctx *Sql_typeContext) {}

// ExitSql_type is called when production sql_type is exited.
func (s *BaseHplsqlListener) ExitSql_type(ctx *Sql_typeContext) {}

// EnterColumn_name is called when production column_name is entered.
func (s *BaseHplsqlListener) EnterColumn_name(ctx *Column_nameContext) {}

// ExitColumn_name is called when production column_name is exited.
func (s *BaseHplsqlListener) ExitColumn_name(ctx *Column_nameContext) {}

// EnterCreate_table_column_inline_cons is called when production create_table_column_inline_cons is entered.
func (s *BaseHplsqlListener) EnterCreate_table_column_inline_cons(ctx *Create_table_column_inline_consContext) {
}

// ExitCreate_table_column_inline_cons is called when production create_table_column_inline_cons is exited.
func (s *BaseHplsqlListener) ExitCreate_table_column_inline_cons(ctx *Create_table_column_inline_consContext) {
}

// EnterCreate_table_column_cons is called when production create_table_column_cons is entered.
func (s *BaseHplsqlListener) EnterCreate_table_column_cons(ctx *Create_table_column_consContext) {}

// ExitCreate_table_column_cons is called when production create_table_column_cons is exited.
func (s *BaseHplsqlListener) ExitCreate_table_column_cons(ctx *Create_table_column_consContext) {}

// EnterCreate_table_fk_action is called when production create_table_fk_action is entered.
func (s *BaseHplsqlListener) EnterCreate_table_fk_action(ctx *Create_table_fk_actionContext) {}

// ExitCreate_table_fk_action is called when production create_table_fk_action is exited.
func (s *BaseHplsqlListener) ExitCreate_table_fk_action(ctx *Create_table_fk_actionContext) {}

// EnterCreate_table_preoptions is called when production create_table_preoptions is entered.
func (s *BaseHplsqlListener) EnterCreate_table_preoptions(ctx *Create_table_preoptionsContext) {}

// ExitCreate_table_preoptions is called when production create_table_preoptions is exited.
func (s *BaseHplsqlListener) ExitCreate_table_preoptions(ctx *Create_table_preoptionsContext) {}

// EnterCreate_table_preoptions_item is called when production create_table_preoptions_item is entered.
func (s *BaseHplsqlListener) EnterCreate_table_preoptions_item(ctx *Create_table_preoptions_itemContext) {
}

// ExitCreate_table_preoptions_item is called when production create_table_preoptions_item is exited.
func (s *BaseHplsqlListener) ExitCreate_table_preoptions_item(ctx *Create_table_preoptions_itemContext) {
}

// EnterCreate_table_preoptions_td_item is called when production create_table_preoptions_td_item is entered.
func (s *BaseHplsqlListener) EnterCreate_table_preoptions_td_item(ctx *Create_table_preoptions_td_itemContext) {
}

// ExitCreate_table_preoptions_td_item is called when production create_table_preoptions_td_item is exited.
func (s *BaseHplsqlListener) ExitCreate_table_preoptions_td_item(ctx *Create_table_preoptions_td_itemContext) {
}

// EnterCreate_table_options is called when production create_table_options is entered.
func (s *BaseHplsqlListener) EnterCreate_table_options(ctx *Create_table_optionsContext) {}

// ExitCreate_table_options is called when production create_table_options is exited.
func (s *BaseHplsqlListener) ExitCreate_table_options(ctx *Create_table_optionsContext) {}

// EnterCreate_table_options_item is called when production create_table_options_item is entered.
func (s *BaseHplsqlListener) EnterCreate_table_options_item(ctx *Create_table_options_itemContext) {}

// ExitCreate_table_options_item is called when production create_table_options_item is exited.
func (s *BaseHplsqlListener) ExitCreate_table_options_item(ctx *Create_table_options_itemContext) {}

// EnterCreate_table_options_ora_item is called when production create_table_options_ora_item is entered.
func (s *BaseHplsqlListener) EnterCreate_table_options_ora_item(ctx *Create_table_options_ora_itemContext) {
}

// ExitCreate_table_options_ora_item is called when production create_table_options_ora_item is exited.
func (s *BaseHplsqlListener) ExitCreate_table_options_ora_item(ctx *Create_table_options_ora_itemContext) {
}

// EnterCreate_table_options_db2_item is called when production create_table_options_db2_item is entered.
func (s *BaseHplsqlListener) EnterCreate_table_options_db2_item(ctx *Create_table_options_db2_itemContext) {
}

// ExitCreate_table_options_db2_item is called when production create_table_options_db2_item is exited.
func (s *BaseHplsqlListener) ExitCreate_table_options_db2_item(ctx *Create_table_options_db2_itemContext) {
}

// EnterCreate_table_options_td_item is called when production create_table_options_td_item is entered.
func (s *BaseHplsqlListener) EnterCreate_table_options_td_item(ctx *Create_table_options_td_itemContext) {
}

// ExitCreate_table_options_td_item is called when production create_table_options_td_item is exited.
func (s *BaseHplsqlListener) ExitCreate_table_options_td_item(ctx *Create_table_options_td_itemContext) {
}

// EnterPartition_column_name is called when production partition_column_name is entered.
func (s *BaseHplsqlListener) EnterPartition_column_name(ctx *Partition_column_nameContext) {}

// ExitPartition_column_name is called when production partition_column_name is exited.
func (s *BaseHplsqlListener) ExitPartition_column_name(ctx *Partition_column_nameContext) {}

// EnterPartition_dtype is called when production partition_dtype is entered.
func (s *BaseHplsqlListener) EnterPartition_dtype(ctx *Partition_dtypeContext) {}

// ExitPartition_dtype is called when production partition_dtype is exited.
func (s *BaseHplsqlListener) ExitPartition_dtype(ctx *Partition_dtypeContext) {}

// EnterPartition_dtype_len is called when production partition_dtype_len is entered.
func (s *BaseHplsqlListener) EnterPartition_dtype_len(ctx *Partition_dtype_lenContext) {}

// ExitPartition_dtype_len is called when production partition_dtype_len is exited.
func (s *BaseHplsqlListener) ExitPartition_dtype_len(ctx *Partition_dtype_lenContext) {}

// EnterCreate_table_hive_partition_columns_item is called when production create_table_hive_partition_columns_item is entered.
func (s *BaseHplsqlListener) EnterCreate_table_hive_partition_columns_item(ctx *Create_table_hive_partition_columns_itemContext) {
}

// ExitCreate_table_hive_partition_columns_item is called when production create_table_hive_partition_columns_item is exited.
func (s *BaseHplsqlListener) ExitCreate_table_hive_partition_columns_item(ctx *Create_table_hive_partition_columns_itemContext) {
}

// EnterCreate_table_hive_partition_column_comment is called when production create_table_hive_partition_column_comment is entered.
func (s *BaseHplsqlListener) EnterCreate_table_hive_partition_column_comment(ctx *Create_table_hive_partition_column_commentContext) {
}

// ExitCreate_table_hive_partition_column_comment is called when production create_table_hive_partition_column_comment is exited.
func (s *BaseHplsqlListener) ExitCreate_table_hive_partition_column_comment(ctx *Create_table_hive_partition_column_commentContext) {
}

// EnterCreate_table_hive_partition_columns is called when production create_table_hive_partition_columns is entered.
func (s *BaseHplsqlListener) EnterCreate_table_hive_partition_columns(ctx *Create_table_hive_partition_columnsContext) {
}

// ExitCreate_table_hive_partition_columns is called when production create_table_hive_partition_columns is exited.
func (s *BaseHplsqlListener) ExitCreate_table_hive_partition_columns(ctx *Create_table_hive_partition_columnsContext) {
}

// EnterCreate_table_hive_partitiion is called when production create_table_hive_partitiion is entered.
func (s *BaseHplsqlListener) EnterCreate_table_hive_partitiion(ctx *Create_table_hive_partitiionContext) {
}

// ExitCreate_table_hive_partitiion is called when production create_table_hive_partitiion is exited.
func (s *BaseHplsqlListener) ExitCreate_table_hive_partitiion(ctx *Create_table_hive_partitiionContext) {
}

// EnterCreate_table_options_hive_comment is called when production create_table_options_hive_comment is entered.
func (s *BaseHplsqlListener) EnterCreate_table_options_hive_comment(ctx *Create_table_options_hive_commentContext) {
}

// ExitCreate_table_options_hive_comment is called when production create_table_options_hive_comment is exited.
func (s *BaseHplsqlListener) ExitCreate_table_options_hive_comment(ctx *Create_table_options_hive_commentContext) {
}

// EnterCreate_table_options_hive_item is called when production create_table_options_hive_item is entered.
func (s *BaseHplsqlListener) EnterCreate_table_options_hive_item(ctx *Create_table_options_hive_itemContext) {
}

// ExitCreate_table_options_hive_item is called when production create_table_options_hive_item is exited.
func (s *BaseHplsqlListener) ExitCreate_table_options_hive_item(ctx *Create_table_options_hive_itemContext) {
}

// EnterCreate_table_hive_row_format is called when production create_table_hive_row_format is entered.
func (s *BaseHplsqlListener) EnterCreate_table_hive_row_format(ctx *Create_table_hive_row_formatContext) {
}

// ExitCreate_table_hive_row_format is called when production create_table_hive_row_format is exited.
func (s *BaseHplsqlListener) ExitCreate_table_hive_row_format(ctx *Create_table_hive_row_formatContext) {
}

// EnterCreate_table_hive_row_format_fields is called when production create_table_hive_row_format_fields is entered.
func (s *BaseHplsqlListener) EnterCreate_table_hive_row_format_fields(ctx *Create_table_hive_row_format_fieldsContext) {
}

// ExitCreate_table_hive_row_format_fields is called when production create_table_hive_row_format_fields is exited.
func (s *BaseHplsqlListener) ExitCreate_table_hive_row_format_fields(ctx *Create_table_hive_row_format_fieldsContext) {
}

// EnterCreate_table_options_mssql_item is called when production create_table_options_mssql_item is entered.
func (s *BaseHplsqlListener) EnterCreate_table_options_mssql_item(ctx *Create_table_options_mssql_itemContext) {
}

// ExitCreate_table_options_mssql_item is called when production create_table_options_mssql_item is exited.
func (s *BaseHplsqlListener) ExitCreate_table_options_mssql_item(ctx *Create_table_options_mssql_itemContext) {
}

// EnterCreate_table_options_mysql_item is called when production create_table_options_mysql_item is entered.
func (s *BaseHplsqlListener) EnterCreate_table_options_mysql_item(ctx *Create_table_options_mysql_itemContext) {
}

// ExitCreate_table_options_mysql_item is called when production create_table_options_mysql_item is exited.
func (s *BaseHplsqlListener) ExitCreate_table_options_mysql_item(ctx *Create_table_options_mysql_itemContext) {
}

// EnterAlter_table_stmt is called when production alter_table_stmt is entered.
func (s *BaseHplsqlListener) EnterAlter_table_stmt(ctx *Alter_table_stmtContext) {}

// ExitAlter_table_stmt is called when production alter_table_stmt is exited.
func (s *BaseHplsqlListener) ExitAlter_table_stmt(ctx *Alter_table_stmtContext) {}

// EnterAlter_table_item is called when production alter_table_item is entered.
func (s *BaseHplsqlListener) EnterAlter_table_item(ctx *Alter_table_itemContext) {}

// ExitAlter_table_item is called when production alter_table_item is exited.
func (s *BaseHplsqlListener) ExitAlter_table_item(ctx *Alter_table_itemContext) {}

// EnterAlter_table_add_constraint is called when production alter_table_add_constraint is entered.
func (s *BaseHplsqlListener) EnterAlter_table_add_constraint(ctx *Alter_table_add_constraintContext) {
}

// ExitAlter_table_add_constraint is called when production alter_table_add_constraint is exited.
func (s *BaseHplsqlListener) ExitAlter_table_add_constraint(ctx *Alter_table_add_constraintContext) {}

// EnterAlter_table_add_constraint_item is called when production alter_table_add_constraint_item is entered.
func (s *BaseHplsqlListener) EnterAlter_table_add_constraint_item(ctx *Alter_table_add_constraint_itemContext) {
}

// ExitAlter_table_add_constraint_item is called when production alter_table_add_constraint_item is exited.
func (s *BaseHplsqlListener) ExitAlter_table_add_constraint_item(ctx *Alter_table_add_constraint_itemContext) {
}

// EnterDtype is called when production dtype is entered.
func (s *BaseHplsqlListener) EnterDtype(ctx *DtypeContext) {}

// ExitDtype is called when production dtype is exited.
func (s *BaseHplsqlListener) ExitDtype(ctx *DtypeContext) {}

// EnterDtype_len is called when production dtype_len is entered.
func (s *BaseHplsqlListener) EnterDtype_len(ctx *Dtype_lenContext) {}

// ExitDtype_len is called when production dtype_len is exited.
func (s *BaseHplsqlListener) ExitDtype_len(ctx *Dtype_lenContext) {}

// EnterDtype_attr is called when production dtype_attr is entered.
func (s *BaseHplsqlListener) EnterDtype_attr(ctx *Dtype_attrContext) {}

// ExitDtype_attr is called when production dtype_attr is exited.
func (s *BaseHplsqlListener) ExitDtype_attr(ctx *Dtype_attrContext) {}

// EnterDtype_default is called when production dtype_default is entered.
func (s *BaseHplsqlListener) EnterDtype_default(ctx *Dtype_defaultContext) {}

// ExitDtype_default is called when production dtype_default is exited.
func (s *BaseHplsqlListener) ExitDtype_default(ctx *Dtype_defaultContext) {}

// EnterCreate_database_stmt is called when production create_database_stmt is entered.
func (s *BaseHplsqlListener) EnterCreate_database_stmt(ctx *Create_database_stmtContext) {}

// ExitCreate_database_stmt is called when production create_database_stmt is exited.
func (s *BaseHplsqlListener) ExitCreate_database_stmt(ctx *Create_database_stmtContext) {}

// EnterCreate_database_option is called when production create_database_option is entered.
func (s *BaseHplsqlListener) EnterCreate_database_option(ctx *Create_database_optionContext) {}

// ExitCreate_database_option is called when production create_database_option is exited.
func (s *BaseHplsqlListener) ExitCreate_database_option(ctx *Create_database_optionContext) {}

// EnterCreate_function_stmt is called when production create_function_stmt is entered.
func (s *BaseHplsqlListener) EnterCreate_function_stmt(ctx *Create_function_stmtContext) {}

// ExitCreate_function_stmt is called when production create_function_stmt is exited.
func (s *BaseHplsqlListener) ExitCreate_function_stmt(ctx *Create_function_stmtContext) {}

// EnterCreate_function_return is called when production create_function_return is entered.
func (s *BaseHplsqlListener) EnterCreate_function_return(ctx *Create_function_returnContext) {}

// ExitCreate_function_return is called when production create_function_return is exited.
func (s *BaseHplsqlListener) ExitCreate_function_return(ctx *Create_function_returnContext) {}

// EnterCreate_package_stmt is called when production create_package_stmt is entered.
func (s *BaseHplsqlListener) EnterCreate_package_stmt(ctx *Create_package_stmtContext) {}

// ExitCreate_package_stmt is called when production create_package_stmt is exited.
func (s *BaseHplsqlListener) ExitCreate_package_stmt(ctx *Create_package_stmtContext) {}

// EnterPackage_spec is called when production package_spec is entered.
func (s *BaseHplsqlListener) EnterPackage_spec(ctx *Package_specContext) {}

// ExitPackage_spec is called when production package_spec is exited.
func (s *BaseHplsqlListener) ExitPackage_spec(ctx *Package_specContext) {}

// EnterPackage_spec_item is called when production package_spec_item is entered.
func (s *BaseHplsqlListener) EnterPackage_spec_item(ctx *Package_spec_itemContext) {}

// ExitPackage_spec_item is called when production package_spec_item is exited.
func (s *BaseHplsqlListener) ExitPackage_spec_item(ctx *Package_spec_itemContext) {}

// EnterCreate_package_body_stmt is called when production create_package_body_stmt is entered.
func (s *BaseHplsqlListener) EnterCreate_package_body_stmt(ctx *Create_package_body_stmtContext) {}

// ExitCreate_package_body_stmt is called when production create_package_body_stmt is exited.
func (s *BaseHplsqlListener) ExitCreate_package_body_stmt(ctx *Create_package_body_stmtContext) {}

// EnterPackage_body is called when production package_body is entered.
func (s *BaseHplsqlListener) EnterPackage_body(ctx *Package_bodyContext) {}

// ExitPackage_body is called when production package_body is exited.
func (s *BaseHplsqlListener) ExitPackage_body(ctx *Package_bodyContext) {}

// EnterPackage_body_item is called when production package_body_item is entered.
func (s *BaseHplsqlListener) EnterPackage_body_item(ctx *Package_body_itemContext) {}

// ExitPackage_body_item is called when production package_body_item is exited.
func (s *BaseHplsqlListener) ExitPackage_body_item(ctx *Package_body_itemContext) {}

// EnterCreate_procedure_stmt is called when production create_procedure_stmt is entered.
func (s *BaseHplsqlListener) EnterCreate_procedure_stmt(ctx *Create_procedure_stmtContext) {}

// ExitCreate_procedure_stmt is called when production create_procedure_stmt is exited.
func (s *BaseHplsqlListener) ExitCreate_procedure_stmt(ctx *Create_procedure_stmtContext) {}

// EnterCreate_routine_params is called when production create_routine_params is entered.
func (s *BaseHplsqlListener) EnterCreate_routine_params(ctx *Create_routine_paramsContext) {}

// ExitCreate_routine_params is called when production create_routine_params is exited.
func (s *BaseHplsqlListener) ExitCreate_routine_params(ctx *Create_routine_paramsContext) {}

// EnterCreate_routine_param_item is called when production create_routine_param_item is entered.
func (s *BaseHplsqlListener) EnterCreate_routine_param_item(ctx *Create_routine_param_itemContext) {}

// ExitCreate_routine_param_item is called when production create_routine_param_item is exited.
func (s *BaseHplsqlListener) ExitCreate_routine_param_item(ctx *Create_routine_param_itemContext) {}

// EnterCreate_routine_options is called when production create_routine_options is entered.
func (s *BaseHplsqlListener) EnterCreate_routine_options(ctx *Create_routine_optionsContext) {}

// ExitCreate_routine_options is called when production create_routine_options is exited.
func (s *BaseHplsqlListener) ExitCreate_routine_options(ctx *Create_routine_optionsContext) {}

// EnterCreate_routine_option is called when production create_routine_option is entered.
func (s *BaseHplsqlListener) EnterCreate_routine_option(ctx *Create_routine_optionContext) {}

// ExitCreate_routine_option is called when production create_routine_option is exited.
func (s *BaseHplsqlListener) ExitCreate_routine_option(ctx *Create_routine_optionContext) {}

// EnterDrop_stmt is called when production drop_stmt is entered.
func (s *BaseHplsqlListener) EnterDrop_stmt(ctx *Drop_stmtContext) {}

// ExitDrop_stmt is called when production drop_stmt is exited.
func (s *BaseHplsqlListener) ExitDrop_stmt(ctx *Drop_stmtContext) {}

// EnterEnd_transaction_stmt is called when production end_transaction_stmt is entered.
func (s *BaseHplsqlListener) EnterEnd_transaction_stmt(ctx *End_transaction_stmtContext) {}

// ExitEnd_transaction_stmt is called when production end_transaction_stmt is exited.
func (s *BaseHplsqlListener) ExitEnd_transaction_stmt(ctx *End_transaction_stmtContext) {}

// EnterExec_stmt is called when production exec_stmt is entered.
func (s *BaseHplsqlListener) EnterExec_stmt(ctx *Exec_stmtContext) {}

// ExitExec_stmt is called when production exec_stmt is exited.
func (s *BaseHplsqlListener) ExitExec_stmt(ctx *Exec_stmtContext) {}

// EnterIf_stmt is called when production if_stmt is entered.
func (s *BaseHplsqlListener) EnterIf_stmt(ctx *If_stmtContext) {}

// ExitIf_stmt is called when production if_stmt is exited.
func (s *BaseHplsqlListener) ExitIf_stmt(ctx *If_stmtContext) {}

// EnterIf_plsql_stmt is called when production if_plsql_stmt is entered.
func (s *BaseHplsqlListener) EnterIf_plsql_stmt(ctx *If_plsql_stmtContext) {}

// ExitIf_plsql_stmt is called when production if_plsql_stmt is exited.
func (s *BaseHplsqlListener) ExitIf_plsql_stmt(ctx *If_plsql_stmtContext) {}

// EnterIf_tsql_stmt is called when production if_tsql_stmt is entered.
func (s *BaseHplsqlListener) EnterIf_tsql_stmt(ctx *If_tsql_stmtContext) {}

// ExitIf_tsql_stmt is called when production if_tsql_stmt is exited.
func (s *BaseHplsqlListener) ExitIf_tsql_stmt(ctx *If_tsql_stmtContext) {}

// EnterIf_bteq_stmt is called when production if_bteq_stmt is entered.
func (s *BaseHplsqlListener) EnterIf_bteq_stmt(ctx *If_bteq_stmtContext) {}

// ExitIf_bteq_stmt is called when production if_bteq_stmt is exited.
func (s *BaseHplsqlListener) ExitIf_bteq_stmt(ctx *If_bteq_stmtContext) {}

// EnterElseif_block is called when production elseif_block is entered.
func (s *BaseHplsqlListener) EnterElseif_block(ctx *Elseif_blockContext) {}

// ExitElseif_block is called when production elseif_block is exited.
func (s *BaseHplsqlListener) ExitElseif_block(ctx *Elseif_blockContext) {}

// EnterElse_block is called when production else_block is entered.
func (s *BaseHplsqlListener) EnterElse_block(ctx *Else_blockContext) {}

// ExitElse_block is called when production else_block is exited.
func (s *BaseHplsqlListener) ExitElse_block(ctx *Else_blockContext) {}

// EnterInclude_stmt is called when production include_stmt is entered.
func (s *BaseHplsqlListener) EnterInclude_stmt(ctx *Include_stmtContext) {}

// ExitInclude_stmt is called when production include_stmt is exited.
func (s *BaseHplsqlListener) ExitInclude_stmt(ctx *Include_stmtContext) {}

// EnterInsert_stmt is called when production insert_stmt is entered.
func (s *BaseHplsqlListener) EnterInsert_stmt(ctx *Insert_stmtContext) {}

// ExitInsert_stmt is called when production insert_stmt is exited.
func (s *BaseHplsqlListener) ExitInsert_stmt(ctx *Insert_stmtContext) {}

// EnterInsert_stmt_cols is called when production insert_stmt_cols is entered.
func (s *BaseHplsqlListener) EnterInsert_stmt_cols(ctx *Insert_stmt_colsContext) {}

// ExitInsert_stmt_cols is called when production insert_stmt_cols is exited.
func (s *BaseHplsqlListener) ExitInsert_stmt_cols(ctx *Insert_stmt_colsContext) {}

// EnterInsert_stmt_rows is called when production insert_stmt_rows is entered.
func (s *BaseHplsqlListener) EnterInsert_stmt_rows(ctx *Insert_stmt_rowsContext) {}

// ExitInsert_stmt_rows is called when production insert_stmt_rows is exited.
func (s *BaseHplsqlListener) ExitInsert_stmt_rows(ctx *Insert_stmt_rowsContext) {}

// EnterInsert_stmt_row is called when production insert_stmt_row is entered.
func (s *BaseHplsqlListener) EnterInsert_stmt_row(ctx *Insert_stmt_rowContext) {}

// ExitInsert_stmt_row is called when production insert_stmt_row is exited.
func (s *BaseHplsqlListener) ExitInsert_stmt_row(ctx *Insert_stmt_rowContext) {}

// EnterInsert_directory_stmt is called when production insert_directory_stmt is entered.
func (s *BaseHplsqlListener) EnterInsert_directory_stmt(ctx *Insert_directory_stmtContext) {}

// ExitInsert_directory_stmt is called when production insert_directory_stmt is exited.
func (s *BaseHplsqlListener) ExitInsert_directory_stmt(ctx *Insert_directory_stmtContext) {}

// EnterExit_stmt is called when production exit_stmt is entered.
func (s *BaseHplsqlListener) EnterExit_stmt(ctx *Exit_stmtContext) {}

// ExitExit_stmt is called when production exit_stmt is exited.
func (s *BaseHplsqlListener) ExitExit_stmt(ctx *Exit_stmtContext) {}

// EnterGet_diag_stmt is called when production get_diag_stmt is entered.
func (s *BaseHplsqlListener) EnterGet_diag_stmt(ctx *Get_diag_stmtContext) {}

// ExitGet_diag_stmt is called when production get_diag_stmt is exited.
func (s *BaseHplsqlListener) ExitGet_diag_stmt(ctx *Get_diag_stmtContext) {}

// EnterGet_diag_stmt_item is called when production get_diag_stmt_item is entered.
func (s *BaseHplsqlListener) EnterGet_diag_stmt_item(ctx *Get_diag_stmt_itemContext) {}

// ExitGet_diag_stmt_item is called when production get_diag_stmt_item is exited.
func (s *BaseHplsqlListener) ExitGet_diag_stmt_item(ctx *Get_diag_stmt_itemContext) {}

// EnterGet_diag_stmt_exception_item is called when production get_diag_stmt_exception_item is entered.
func (s *BaseHplsqlListener) EnterGet_diag_stmt_exception_item(ctx *Get_diag_stmt_exception_itemContext) {
}

// ExitGet_diag_stmt_exception_item is called when production get_diag_stmt_exception_item is exited.
func (s *BaseHplsqlListener) ExitGet_diag_stmt_exception_item(ctx *Get_diag_stmt_exception_itemContext) {
}

// EnterGet_diag_stmt_rowcount_item is called when production get_diag_stmt_rowcount_item is entered.
func (s *BaseHplsqlListener) EnterGet_diag_stmt_rowcount_item(ctx *Get_diag_stmt_rowcount_itemContext) {
}

// ExitGet_diag_stmt_rowcount_item is called when production get_diag_stmt_rowcount_item is exited.
func (s *BaseHplsqlListener) ExitGet_diag_stmt_rowcount_item(ctx *Get_diag_stmt_rowcount_itemContext) {
}

// EnterGrant_stmt is called when production grant_stmt is entered.
func (s *BaseHplsqlListener) EnterGrant_stmt(ctx *Grant_stmtContext) {}

// ExitGrant_stmt is called when production grant_stmt is exited.
func (s *BaseHplsqlListener) ExitGrant_stmt(ctx *Grant_stmtContext) {}

// EnterGrant_stmt_item is called when production grant_stmt_item is entered.
func (s *BaseHplsqlListener) EnterGrant_stmt_item(ctx *Grant_stmt_itemContext) {}

// ExitGrant_stmt_item is called when production grant_stmt_item is exited.
func (s *BaseHplsqlListener) ExitGrant_stmt_item(ctx *Grant_stmt_itemContext) {}

// EnterLeave_stmt is called when production leave_stmt is entered.
func (s *BaseHplsqlListener) EnterLeave_stmt(ctx *Leave_stmtContext) {}

// ExitLeave_stmt is called when production leave_stmt is exited.
func (s *BaseHplsqlListener) ExitLeave_stmt(ctx *Leave_stmtContext) {}

// EnterMap_object_stmt is called when production map_object_stmt is entered.
func (s *BaseHplsqlListener) EnterMap_object_stmt(ctx *Map_object_stmtContext) {}

// ExitMap_object_stmt is called when production map_object_stmt is exited.
func (s *BaseHplsqlListener) ExitMap_object_stmt(ctx *Map_object_stmtContext) {}

// EnterOpen_stmt is called when production open_stmt is entered.
func (s *BaseHplsqlListener) EnterOpen_stmt(ctx *Open_stmtContext) {}

// ExitOpen_stmt is called when production open_stmt is exited.
func (s *BaseHplsqlListener) ExitOpen_stmt(ctx *Open_stmtContext) {}

// EnterFetch_stmt is called when production fetch_stmt is entered.
func (s *BaseHplsqlListener) EnterFetch_stmt(ctx *Fetch_stmtContext) {}

// ExitFetch_stmt is called when production fetch_stmt is exited.
func (s *BaseHplsqlListener) ExitFetch_stmt(ctx *Fetch_stmtContext) {}

// EnterFetch_limit is called when production fetch_limit is entered.
func (s *BaseHplsqlListener) EnterFetch_limit(ctx *Fetch_limitContext) {}

// ExitFetch_limit is called when production fetch_limit is exited.
func (s *BaseHplsqlListener) ExitFetch_limit(ctx *Fetch_limitContext) {}

// EnterCollect_stats_stmt is called when production collect_stats_stmt is entered.
func (s *BaseHplsqlListener) EnterCollect_stats_stmt(ctx *Collect_stats_stmtContext) {}

// ExitCollect_stats_stmt is called when production collect_stats_stmt is exited.
func (s *BaseHplsqlListener) ExitCollect_stats_stmt(ctx *Collect_stats_stmtContext) {}

// EnterCollect_stats_clause is called when production collect_stats_clause is entered.
func (s *BaseHplsqlListener) EnterCollect_stats_clause(ctx *Collect_stats_clauseContext) {}

// ExitCollect_stats_clause is called when production collect_stats_clause is exited.
func (s *BaseHplsqlListener) ExitCollect_stats_clause(ctx *Collect_stats_clauseContext) {}

// EnterClose_stmt is called when production close_stmt is entered.
func (s *BaseHplsqlListener) EnterClose_stmt(ctx *Close_stmtContext) {}

// ExitClose_stmt is called when production close_stmt is exited.
func (s *BaseHplsqlListener) ExitClose_stmt(ctx *Close_stmtContext) {}

// EnterCmp_stmt is called when production cmp_stmt is entered.
func (s *BaseHplsqlListener) EnterCmp_stmt(ctx *Cmp_stmtContext) {}

// ExitCmp_stmt is called when production cmp_stmt is exited.
func (s *BaseHplsqlListener) ExitCmp_stmt(ctx *Cmp_stmtContext) {}

// EnterCmp_source is called when production cmp_source is entered.
func (s *BaseHplsqlListener) EnterCmp_source(ctx *Cmp_sourceContext) {}

// ExitCmp_source is called when production cmp_source is exited.
func (s *BaseHplsqlListener) ExitCmp_source(ctx *Cmp_sourceContext) {}

// EnterCopy_from_local_stmt is called when production copy_from_local_stmt is entered.
func (s *BaseHplsqlListener) EnterCopy_from_local_stmt(ctx *Copy_from_local_stmtContext) {}

// ExitCopy_from_local_stmt is called when production copy_from_local_stmt is exited.
func (s *BaseHplsqlListener) ExitCopy_from_local_stmt(ctx *Copy_from_local_stmtContext) {}

// EnterCopy_stmt is called when production copy_stmt is entered.
func (s *BaseHplsqlListener) EnterCopy_stmt(ctx *Copy_stmtContext) {}

// ExitCopy_stmt is called when production copy_stmt is exited.
func (s *BaseHplsqlListener) ExitCopy_stmt(ctx *Copy_stmtContext) {}

// EnterCopy_source is called when production copy_source is entered.
func (s *BaseHplsqlListener) EnterCopy_source(ctx *Copy_sourceContext) {}

// ExitCopy_source is called when production copy_source is exited.
func (s *BaseHplsqlListener) ExitCopy_source(ctx *Copy_sourceContext) {}

// EnterCopy_target is called when production copy_target is entered.
func (s *BaseHplsqlListener) EnterCopy_target(ctx *Copy_targetContext) {}

// ExitCopy_target is called when production copy_target is exited.
func (s *BaseHplsqlListener) ExitCopy_target(ctx *Copy_targetContext) {}

// EnterCopy_option is called when production copy_option is entered.
func (s *BaseHplsqlListener) EnterCopy_option(ctx *Copy_optionContext) {}

// ExitCopy_option is called when production copy_option is exited.
func (s *BaseHplsqlListener) ExitCopy_option(ctx *Copy_optionContext) {}

// EnterCopy_file_option is called when production copy_file_option is entered.
func (s *BaseHplsqlListener) EnterCopy_file_option(ctx *Copy_file_optionContext) {}

// ExitCopy_file_option is called when production copy_file_option is exited.
func (s *BaseHplsqlListener) ExitCopy_file_option(ctx *Copy_file_optionContext) {}

// EnterCommit_stmt is called when production commit_stmt is entered.
func (s *BaseHplsqlListener) EnterCommit_stmt(ctx *Commit_stmtContext) {}

// ExitCommit_stmt is called when production commit_stmt is exited.
func (s *BaseHplsqlListener) ExitCommit_stmt(ctx *Commit_stmtContext) {}

// EnterCreate_index_stmt is called when production create_index_stmt is entered.
func (s *BaseHplsqlListener) EnterCreate_index_stmt(ctx *Create_index_stmtContext) {}

// ExitCreate_index_stmt is called when production create_index_stmt is exited.
func (s *BaseHplsqlListener) ExitCreate_index_stmt(ctx *Create_index_stmtContext) {}

// EnterCreate_index_col is called when production create_index_col is entered.
func (s *BaseHplsqlListener) EnterCreate_index_col(ctx *Create_index_colContext) {}

// ExitCreate_index_col is called when production create_index_col is exited.
func (s *BaseHplsqlListener) ExitCreate_index_col(ctx *Create_index_colContext) {}

// EnterIndex_storage_clause is called when production index_storage_clause is entered.
func (s *BaseHplsqlListener) EnterIndex_storage_clause(ctx *Index_storage_clauseContext) {}

// ExitIndex_storage_clause is called when production index_storage_clause is exited.
func (s *BaseHplsqlListener) ExitIndex_storage_clause(ctx *Index_storage_clauseContext) {}

// EnterIndex_mssql_storage_clause is called when production index_mssql_storage_clause is entered.
func (s *BaseHplsqlListener) EnterIndex_mssql_storage_clause(ctx *Index_mssql_storage_clauseContext) {
}

// ExitIndex_mssql_storage_clause is called when production index_mssql_storage_clause is exited.
func (s *BaseHplsqlListener) ExitIndex_mssql_storage_clause(ctx *Index_mssql_storage_clauseContext) {}

// EnterPrint_stmt is called when production print_stmt is entered.
func (s *BaseHplsqlListener) EnterPrint_stmt(ctx *Print_stmtContext) {}

// ExitPrint_stmt is called when production print_stmt is exited.
func (s *BaseHplsqlListener) ExitPrint_stmt(ctx *Print_stmtContext) {}

// EnterQuit_stmt is called when production quit_stmt is entered.
func (s *BaseHplsqlListener) EnterQuit_stmt(ctx *Quit_stmtContext) {}

// ExitQuit_stmt is called when production quit_stmt is exited.
func (s *BaseHplsqlListener) ExitQuit_stmt(ctx *Quit_stmtContext) {}

// EnterRaise_stmt is called when production raise_stmt is entered.
func (s *BaseHplsqlListener) EnterRaise_stmt(ctx *Raise_stmtContext) {}

// ExitRaise_stmt is called when production raise_stmt is exited.
func (s *BaseHplsqlListener) ExitRaise_stmt(ctx *Raise_stmtContext) {}

// EnterResignal_stmt is called when production resignal_stmt is entered.
func (s *BaseHplsqlListener) EnterResignal_stmt(ctx *Resignal_stmtContext) {}

// ExitResignal_stmt is called when production resignal_stmt is exited.
func (s *BaseHplsqlListener) ExitResignal_stmt(ctx *Resignal_stmtContext) {}

// EnterReturn_stmt is called when production return_stmt is entered.
func (s *BaseHplsqlListener) EnterReturn_stmt(ctx *Return_stmtContext) {}

// ExitReturn_stmt is called when production return_stmt is exited.
func (s *BaseHplsqlListener) ExitReturn_stmt(ctx *Return_stmtContext) {}

// EnterRollback_stmt is called when production rollback_stmt is entered.
func (s *BaseHplsqlListener) EnterRollback_stmt(ctx *Rollback_stmtContext) {}

// ExitRollback_stmt is called when production rollback_stmt is exited.
func (s *BaseHplsqlListener) ExitRollback_stmt(ctx *Rollback_stmtContext) {}

// EnterSet_session_option is called when production set_session_option is entered.
func (s *BaseHplsqlListener) EnterSet_session_option(ctx *Set_session_optionContext) {}

// ExitSet_session_option is called when production set_session_option is exited.
func (s *BaseHplsqlListener) ExitSet_session_option(ctx *Set_session_optionContext) {}

// EnterSet_current_schema_option is called when production set_current_schema_option is entered.
func (s *BaseHplsqlListener) EnterSet_current_schema_option(ctx *Set_current_schema_optionContext) {}

// ExitSet_current_schema_option is called when production set_current_schema_option is exited.
func (s *BaseHplsqlListener) ExitSet_current_schema_option(ctx *Set_current_schema_optionContext) {}

// EnterSet_mssql_session_option is called when production set_mssql_session_option is entered.
func (s *BaseHplsqlListener) EnterSet_mssql_session_option(ctx *Set_mssql_session_optionContext) {}

// ExitSet_mssql_session_option is called when production set_mssql_session_option is exited.
func (s *BaseHplsqlListener) ExitSet_mssql_session_option(ctx *Set_mssql_session_optionContext) {}

// EnterSet_teradata_session_option is called when production set_teradata_session_option is entered.
func (s *BaseHplsqlListener) EnterSet_teradata_session_option(ctx *Set_teradata_session_optionContext) {
}

// ExitSet_teradata_session_option is called when production set_teradata_session_option is exited.
func (s *BaseHplsqlListener) ExitSet_teradata_session_option(ctx *Set_teradata_session_optionContext) {
}

// EnterSignal_stmt is called when production signal_stmt is entered.
func (s *BaseHplsqlListener) EnterSignal_stmt(ctx *Signal_stmtContext) {}

// ExitSignal_stmt is called when production signal_stmt is exited.
func (s *BaseHplsqlListener) ExitSignal_stmt(ctx *Signal_stmtContext) {}

// EnterSummary_stmt is called when production summary_stmt is entered.
func (s *BaseHplsqlListener) EnterSummary_stmt(ctx *Summary_stmtContext) {}

// ExitSummary_stmt is called when production summary_stmt is exited.
func (s *BaseHplsqlListener) ExitSummary_stmt(ctx *Summary_stmtContext) {}

// EnterTruncate_stmt is called when production truncate_stmt is entered.
func (s *BaseHplsqlListener) EnterTruncate_stmt(ctx *Truncate_stmtContext) {}

// ExitTruncate_stmt is called when production truncate_stmt is exited.
func (s *BaseHplsqlListener) ExitTruncate_stmt(ctx *Truncate_stmtContext) {}

// EnterUse_stmt is called when production use_stmt is entered.
func (s *BaseHplsqlListener) EnterUse_stmt(ctx *Use_stmtContext) {}

// ExitUse_stmt is called when production use_stmt is exited.
func (s *BaseHplsqlListener) ExitUse_stmt(ctx *Use_stmtContext) {}

// EnterValues_into_stmt is called when production values_into_stmt is entered.
func (s *BaseHplsqlListener) EnterValues_into_stmt(ctx *Values_into_stmtContext) {}

// ExitValues_into_stmt is called when production values_into_stmt is exited.
func (s *BaseHplsqlListener) ExitValues_into_stmt(ctx *Values_into_stmtContext) {}

// EnterWhile_stmt is called when production while_stmt is entered.
func (s *BaseHplsqlListener) EnterWhile_stmt(ctx *While_stmtContext) {}

// ExitWhile_stmt is called when production while_stmt is exited.
func (s *BaseHplsqlListener) ExitWhile_stmt(ctx *While_stmtContext) {}

// EnterUnconditional_loop_stmt is called when production unconditional_loop_stmt is entered.
func (s *BaseHplsqlListener) EnterUnconditional_loop_stmt(ctx *Unconditional_loop_stmtContext) {}

// ExitUnconditional_loop_stmt is called when production unconditional_loop_stmt is exited.
func (s *BaseHplsqlListener) ExitUnconditional_loop_stmt(ctx *Unconditional_loop_stmtContext) {}

// EnterFor_cursor_stmt is called when production for_cursor_stmt is entered.
func (s *BaseHplsqlListener) EnterFor_cursor_stmt(ctx *For_cursor_stmtContext) {}

// ExitFor_cursor_stmt is called when production for_cursor_stmt is exited.
func (s *BaseHplsqlListener) ExitFor_cursor_stmt(ctx *For_cursor_stmtContext) {}

// EnterFor_range_stmt is called when production for_range_stmt is entered.
func (s *BaseHplsqlListener) EnterFor_range_stmt(ctx *For_range_stmtContext) {}

// ExitFor_range_stmt is called when production for_range_stmt is exited.
func (s *BaseHplsqlListener) ExitFor_range_stmt(ctx *For_range_stmtContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseHplsqlListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseHplsqlListener) ExitLabel(ctx *LabelContext) {}

// EnterUsing_clause is called when production using_clause is entered.
func (s *BaseHplsqlListener) EnterUsing_clause(ctx *Using_clauseContext) {}

// ExitUsing_clause is called when production using_clause is exited.
func (s *BaseHplsqlListener) ExitUsing_clause(ctx *Using_clauseContext) {}

// EnterSelect_stmt is called when production select_stmt is entered.
func (s *BaseHplsqlListener) EnterSelect_stmt(ctx *Select_stmtContext) {}

// ExitSelect_stmt is called when production select_stmt is exited.
func (s *BaseHplsqlListener) ExitSelect_stmt(ctx *Select_stmtContext) {}

// EnterCte_select_stmt is called when production cte_select_stmt is entered.
func (s *BaseHplsqlListener) EnterCte_select_stmt(ctx *Cte_select_stmtContext) {}

// ExitCte_select_stmt is called when production cte_select_stmt is exited.
func (s *BaseHplsqlListener) ExitCte_select_stmt(ctx *Cte_select_stmtContext) {}

// EnterCte_select_stmt_item is called when production cte_select_stmt_item is entered.
func (s *BaseHplsqlListener) EnterCte_select_stmt_item(ctx *Cte_select_stmt_itemContext) {}

// ExitCte_select_stmt_item is called when production cte_select_stmt_item is exited.
func (s *BaseHplsqlListener) ExitCte_select_stmt_item(ctx *Cte_select_stmt_itemContext) {}

// EnterCte_select_cols is called when production cte_select_cols is entered.
func (s *BaseHplsqlListener) EnterCte_select_cols(ctx *Cte_select_colsContext) {}

// ExitCte_select_cols is called when production cte_select_cols is exited.
func (s *BaseHplsqlListener) ExitCte_select_cols(ctx *Cte_select_colsContext) {}

// EnterFullselect_stmt is called when production fullselect_stmt is entered.
func (s *BaseHplsqlListener) EnterFullselect_stmt(ctx *Fullselect_stmtContext) {}

// ExitFullselect_stmt is called when production fullselect_stmt is exited.
func (s *BaseHplsqlListener) ExitFullselect_stmt(ctx *Fullselect_stmtContext) {}

// EnterFullselect_stmt_item is called when production fullselect_stmt_item is entered.
func (s *BaseHplsqlListener) EnterFullselect_stmt_item(ctx *Fullselect_stmt_itemContext) {}

// ExitFullselect_stmt_item is called when production fullselect_stmt_item is exited.
func (s *BaseHplsqlListener) ExitFullselect_stmt_item(ctx *Fullselect_stmt_itemContext) {}

// EnterFullselect_set_clause is called when production fullselect_set_clause is entered.
func (s *BaseHplsqlListener) EnterFullselect_set_clause(ctx *Fullselect_set_clauseContext) {}

// ExitFullselect_set_clause is called when production fullselect_set_clause is exited.
func (s *BaseHplsqlListener) ExitFullselect_set_clause(ctx *Fullselect_set_clauseContext) {}

// EnterSubselect_stmt is called when production subselect_stmt is entered.
func (s *BaseHplsqlListener) EnterSubselect_stmt(ctx *Subselect_stmtContext) {}

// ExitSubselect_stmt is called when production subselect_stmt is exited.
func (s *BaseHplsqlListener) ExitSubselect_stmt(ctx *Subselect_stmtContext) {}

// EnterSelect_list is called when production select_list is entered.
func (s *BaseHplsqlListener) EnterSelect_list(ctx *Select_listContext) {}

// ExitSelect_list is called when production select_list is exited.
func (s *BaseHplsqlListener) ExitSelect_list(ctx *Select_listContext) {}

// EnterSelect_list_set is called when production select_list_set is entered.
func (s *BaseHplsqlListener) EnterSelect_list_set(ctx *Select_list_setContext) {}

// ExitSelect_list_set is called when production select_list_set is exited.
func (s *BaseHplsqlListener) ExitSelect_list_set(ctx *Select_list_setContext) {}

// EnterSelect_list_limit is called when production select_list_limit is entered.
func (s *BaseHplsqlListener) EnterSelect_list_limit(ctx *Select_list_limitContext) {}

// ExitSelect_list_limit is called when production select_list_limit is exited.
func (s *BaseHplsqlListener) ExitSelect_list_limit(ctx *Select_list_limitContext) {}

// EnterSelect_list_item is called when production select_list_item is entered.
func (s *BaseHplsqlListener) EnterSelect_list_item(ctx *Select_list_itemContext) {}

// ExitSelect_list_item is called when production select_list_item is exited.
func (s *BaseHplsqlListener) ExitSelect_list_item(ctx *Select_list_itemContext) {}

// EnterSelect_list_alias is called when production select_list_alias is entered.
func (s *BaseHplsqlListener) EnterSelect_list_alias(ctx *Select_list_aliasContext) {}

// ExitSelect_list_alias is called when production select_list_alias is exited.
func (s *BaseHplsqlListener) ExitSelect_list_alias(ctx *Select_list_aliasContext) {}

// EnterSelect_list_asterisk is called when production select_list_asterisk is entered.
func (s *BaseHplsqlListener) EnterSelect_list_asterisk(ctx *Select_list_asteriskContext) {}

// ExitSelect_list_asterisk is called when production select_list_asterisk is exited.
func (s *BaseHplsqlListener) ExitSelect_list_asterisk(ctx *Select_list_asteriskContext) {}

// EnterTable_row is called when production table_row is entered.
func (s *BaseHplsqlListener) EnterTable_row(ctx *Table_rowContext) {}

// ExitTable_row is called when production table_row is exited.
func (s *BaseHplsqlListener) ExitTable_row(ctx *Table_rowContext) {}

// EnterInto_clause is called when production into_clause is entered.
func (s *BaseHplsqlListener) EnterInto_clause(ctx *Into_clauseContext) {}

// ExitInto_clause is called when production into_clause is exited.
func (s *BaseHplsqlListener) ExitInto_clause(ctx *Into_clauseContext) {}

// EnterBulk_collect_clause is called when production bulk_collect_clause is entered.
func (s *BaseHplsqlListener) EnterBulk_collect_clause(ctx *Bulk_collect_clauseContext) {}

// ExitBulk_collect_clause is called when production bulk_collect_clause is exited.
func (s *BaseHplsqlListener) ExitBulk_collect_clause(ctx *Bulk_collect_clauseContext) {}

// EnterFrom_clause is called when production from_clause is entered.
func (s *BaseHplsqlListener) EnterFrom_clause(ctx *From_clauseContext) {}

// ExitFrom_clause is called when production from_clause is exited.
func (s *BaseHplsqlListener) ExitFrom_clause(ctx *From_clauseContext) {}

// EnterFrom_table_clause is called when production from_table_clause is entered.
func (s *BaseHplsqlListener) EnterFrom_table_clause(ctx *From_table_clauseContext) {}

// ExitFrom_table_clause is called when production from_table_clause is exited.
func (s *BaseHplsqlListener) ExitFrom_table_clause(ctx *From_table_clauseContext) {}

// EnterFrom_table_name_clause is called when production from_table_name_clause is entered.
func (s *BaseHplsqlListener) EnterFrom_table_name_clause(ctx *From_table_name_clauseContext) {}

// ExitFrom_table_name_clause is called when production from_table_name_clause is exited.
func (s *BaseHplsqlListener) ExitFrom_table_name_clause(ctx *From_table_name_clauseContext) {}

// EnterFrom_subselect_clause is called when production from_subselect_clause is entered.
func (s *BaseHplsqlListener) EnterFrom_subselect_clause(ctx *From_subselect_clauseContext) {}

// ExitFrom_subselect_clause is called when production from_subselect_clause is exited.
func (s *BaseHplsqlListener) ExitFrom_subselect_clause(ctx *From_subselect_clauseContext) {}

// EnterFrom_join_clause is called when production from_join_clause is entered.
func (s *BaseHplsqlListener) EnterFrom_join_clause(ctx *From_join_clauseContext) {}

// ExitFrom_join_clause is called when production from_join_clause is exited.
func (s *BaseHplsqlListener) ExitFrom_join_clause(ctx *From_join_clauseContext) {}

// EnterFrom_join_type_clause is called when production from_join_type_clause is entered.
func (s *BaseHplsqlListener) EnterFrom_join_type_clause(ctx *From_join_type_clauseContext) {}

// ExitFrom_join_type_clause is called when production from_join_type_clause is exited.
func (s *BaseHplsqlListener) ExitFrom_join_type_clause(ctx *From_join_type_clauseContext) {}

// EnterFrom_table_values_clause is called when production from_table_values_clause is entered.
func (s *BaseHplsqlListener) EnterFrom_table_values_clause(ctx *From_table_values_clauseContext) {}

// ExitFrom_table_values_clause is called when production from_table_values_clause is exited.
func (s *BaseHplsqlListener) ExitFrom_table_values_clause(ctx *From_table_values_clauseContext) {}

// EnterFrom_table_values_row is called when production from_table_values_row is entered.
func (s *BaseHplsqlListener) EnterFrom_table_values_row(ctx *From_table_values_rowContext) {}

// ExitFrom_table_values_row is called when production from_table_values_row is exited.
func (s *BaseHplsqlListener) ExitFrom_table_values_row(ctx *From_table_values_rowContext) {}

// EnterFrom_alias_clause is called when production from_alias_clause is entered.
func (s *BaseHplsqlListener) EnterFrom_alias_clause(ctx *From_alias_clauseContext) {}

// ExitFrom_alias_clause is called when production from_alias_clause is exited.
func (s *BaseHplsqlListener) ExitFrom_alias_clause(ctx *From_alias_clauseContext) {}

// EnterTable_name is called when production table_name is entered.
func (s *BaseHplsqlListener) EnterTable_name(ctx *Table_nameContext) {}

// ExitTable_name is called when production table_name is exited.
func (s *BaseHplsqlListener) ExitTable_name(ctx *Table_nameContext) {}

// EnterWhere_clause is called when production where_clause is entered.
func (s *BaseHplsqlListener) EnterWhere_clause(ctx *Where_clauseContext) {}

// ExitWhere_clause is called when production where_clause is exited.
func (s *BaseHplsqlListener) ExitWhere_clause(ctx *Where_clauseContext) {}

// EnterGroup_by_clause is called when production group_by_clause is entered.
func (s *BaseHplsqlListener) EnterGroup_by_clause(ctx *Group_by_clauseContext) {}

// ExitGroup_by_clause is called when production group_by_clause is exited.
func (s *BaseHplsqlListener) ExitGroup_by_clause(ctx *Group_by_clauseContext) {}

// EnterHaving_clause is called when production having_clause is entered.
func (s *BaseHplsqlListener) EnterHaving_clause(ctx *Having_clauseContext) {}

// ExitHaving_clause is called when production having_clause is exited.
func (s *BaseHplsqlListener) ExitHaving_clause(ctx *Having_clauseContext) {}

// EnterQualify_clause is called when production qualify_clause is entered.
func (s *BaseHplsqlListener) EnterQualify_clause(ctx *Qualify_clauseContext) {}

// ExitQualify_clause is called when production qualify_clause is exited.
func (s *BaseHplsqlListener) ExitQualify_clause(ctx *Qualify_clauseContext) {}

// EnterOrder_by_clause is called when production order_by_clause is entered.
func (s *BaseHplsqlListener) EnterOrder_by_clause(ctx *Order_by_clauseContext) {}

// ExitOrder_by_clause is called when production order_by_clause is exited.
func (s *BaseHplsqlListener) ExitOrder_by_clause(ctx *Order_by_clauseContext) {}

// EnterSelect_options is called when production select_options is entered.
func (s *BaseHplsqlListener) EnterSelect_options(ctx *Select_optionsContext) {}

// ExitSelect_options is called when production select_options is exited.
func (s *BaseHplsqlListener) ExitSelect_options(ctx *Select_optionsContext) {}

// EnterSelect_options_item is called when production select_options_item is entered.
func (s *BaseHplsqlListener) EnterSelect_options_item(ctx *Select_options_itemContext) {}

// ExitSelect_options_item is called when production select_options_item is exited.
func (s *BaseHplsqlListener) ExitSelect_options_item(ctx *Select_options_itemContext) {}

// EnterUpdate_stmt is called when production update_stmt is entered.
func (s *BaseHplsqlListener) EnterUpdate_stmt(ctx *Update_stmtContext) {}

// ExitUpdate_stmt is called when production update_stmt is exited.
func (s *BaseHplsqlListener) ExitUpdate_stmt(ctx *Update_stmtContext) {}

// EnterUpdate_assignment is called when production update_assignment is entered.
func (s *BaseHplsqlListener) EnterUpdate_assignment(ctx *Update_assignmentContext) {}

// ExitUpdate_assignment is called when production update_assignment is exited.
func (s *BaseHplsqlListener) ExitUpdate_assignment(ctx *Update_assignmentContext) {}

// EnterUpdate_table is called when production update_table is entered.
func (s *BaseHplsqlListener) EnterUpdate_table(ctx *Update_tableContext) {}

// ExitUpdate_table is called when production update_table is exited.
func (s *BaseHplsqlListener) ExitUpdate_table(ctx *Update_tableContext) {}

// EnterUpdate_upsert is called when production update_upsert is entered.
func (s *BaseHplsqlListener) EnterUpdate_upsert(ctx *Update_upsertContext) {}

// ExitUpdate_upsert is called when production update_upsert is exited.
func (s *BaseHplsqlListener) ExitUpdate_upsert(ctx *Update_upsertContext) {}

// EnterMerge_stmt is called when production merge_stmt is entered.
func (s *BaseHplsqlListener) EnterMerge_stmt(ctx *Merge_stmtContext) {}

// ExitMerge_stmt is called when production merge_stmt is exited.
func (s *BaseHplsqlListener) ExitMerge_stmt(ctx *Merge_stmtContext) {}

// EnterMerge_table is called when production merge_table is entered.
func (s *BaseHplsqlListener) EnterMerge_table(ctx *Merge_tableContext) {}

// ExitMerge_table is called when production merge_table is exited.
func (s *BaseHplsqlListener) ExitMerge_table(ctx *Merge_tableContext) {}

// EnterMerge_condition is called when production merge_condition is entered.
func (s *BaseHplsqlListener) EnterMerge_condition(ctx *Merge_conditionContext) {}

// ExitMerge_condition is called when production merge_condition is exited.
func (s *BaseHplsqlListener) ExitMerge_condition(ctx *Merge_conditionContext) {}

// EnterMerge_action is called when production merge_action is entered.
func (s *BaseHplsqlListener) EnterMerge_action(ctx *Merge_actionContext) {}

// ExitMerge_action is called when production merge_action is exited.
func (s *BaseHplsqlListener) ExitMerge_action(ctx *Merge_actionContext) {}

// EnterDelete_stmt is called when production delete_stmt is entered.
func (s *BaseHplsqlListener) EnterDelete_stmt(ctx *Delete_stmtContext) {}

// ExitDelete_stmt is called when production delete_stmt is exited.
func (s *BaseHplsqlListener) ExitDelete_stmt(ctx *Delete_stmtContext) {}

// EnterDelete_alias is called when production delete_alias is entered.
func (s *BaseHplsqlListener) EnterDelete_alias(ctx *Delete_aliasContext) {}

// ExitDelete_alias is called when production delete_alias is exited.
func (s *BaseHplsqlListener) ExitDelete_alias(ctx *Delete_aliasContext) {}

// EnterDescribe_stmt is called when production describe_stmt is entered.
func (s *BaseHplsqlListener) EnterDescribe_stmt(ctx *Describe_stmtContext) {}

// ExitDescribe_stmt is called when production describe_stmt is exited.
func (s *BaseHplsqlListener) ExitDescribe_stmt(ctx *Describe_stmtContext) {}

// EnterBool_expr is called when production bool_expr is entered.
func (s *BaseHplsqlListener) EnterBool_expr(ctx *Bool_exprContext) {}

// ExitBool_expr is called when production bool_expr is exited.
func (s *BaseHplsqlListener) ExitBool_expr(ctx *Bool_exprContext) {}

// EnterBool_expr_atom is called when production bool_expr_atom is entered.
func (s *BaseHplsqlListener) EnterBool_expr_atom(ctx *Bool_expr_atomContext) {}

// ExitBool_expr_atom is called when production bool_expr_atom is exited.
func (s *BaseHplsqlListener) ExitBool_expr_atom(ctx *Bool_expr_atomContext) {}

// EnterBool_expr_unary is called when production bool_expr_unary is entered.
func (s *BaseHplsqlListener) EnterBool_expr_unary(ctx *Bool_expr_unaryContext) {}

// ExitBool_expr_unary is called when production bool_expr_unary is exited.
func (s *BaseHplsqlListener) ExitBool_expr_unary(ctx *Bool_expr_unaryContext) {}

// EnterBool_expr_single_in is called when production bool_expr_single_in is entered.
func (s *BaseHplsqlListener) EnterBool_expr_single_in(ctx *Bool_expr_single_inContext) {}

// ExitBool_expr_single_in is called when production bool_expr_single_in is exited.
func (s *BaseHplsqlListener) ExitBool_expr_single_in(ctx *Bool_expr_single_inContext) {}

// EnterBool_expr_multi_in is called when production bool_expr_multi_in is entered.
func (s *BaseHplsqlListener) EnterBool_expr_multi_in(ctx *Bool_expr_multi_inContext) {}

// ExitBool_expr_multi_in is called when production bool_expr_multi_in is exited.
func (s *BaseHplsqlListener) ExitBool_expr_multi_in(ctx *Bool_expr_multi_inContext) {}

// EnterBool_expr_binary is called when production bool_expr_binary is entered.
func (s *BaseHplsqlListener) EnterBool_expr_binary(ctx *Bool_expr_binaryContext) {}

// ExitBool_expr_binary is called when production bool_expr_binary is exited.
func (s *BaseHplsqlListener) ExitBool_expr_binary(ctx *Bool_expr_binaryContext) {}

// EnterBool_expr_logical_operator is called when production bool_expr_logical_operator is entered.
func (s *BaseHplsqlListener) EnterBool_expr_logical_operator(ctx *Bool_expr_logical_operatorContext) {
}

// ExitBool_expr_logical_operator is called when production bool_expr_logical_operator is exited.
func (s *BaseHplsqlListener) ExitBool_expr_logical_operator(ctx *Bool_expr_logical_operatorContext) {}

// EnterBool_expr_binary_operator is called when production bool_expr_binary_operator is entered.
func (s *BaseHplsqlListener) EnterBool_expr_binary_operator(ctx *Bool_expr_binary_operatorContext) {}

// ExitBool_expr_binary_operator is called when production bool_expr_binary_operator is exited.
func (s *BaseHplsqlListener) ExitBool_expr_binary_operator(ctx *Bool_expr_binary_operatorContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseHplsqlListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseHplsqlListener) ExitExpr(ctx *ExprContext) {}

// EnterExpr_atom is called when production expr_atom is entered.
func (s *BaseHplsqlListener) EnterExpr_atom(ctx *Expr_atomContext) {}

// ExitExpr_atom is called when production expr_atom is exited.
func (s *BaseHplsqlListener) ExitExpr_atom(ctx *Expr_atomContext) {}

// EnterExpr_interval is called when production expr_interval is entered.
func (s *BaseHplsqlListener) EnterExpr_interval(ctx *Expr_intervalContext) {}

// ExitExpr_interval is called when production expr_interval is exited.
func (s *BaseHplsqlListener) ExitExpr_interval(ctx *Expr_intervalContext) {}

// EnterInterval_item is called when production interval_item is entered.
func (s *BaseHplsqlListener) EnterInterval_item(ctx *Interval_itemContext) {}

// ExitInterval_item is called when production interval_item is exited.
func (s *BaseHplsqlListener) ExitInterval_item(ctx *Interval_itemContext) {}

// EnterExpr_concat is called when production expr_concat is entered.
func (s *BaseHplsqlListener) EnterExpr_concat(ctx *Expr_concatContext) {}

// ExitExpr_concat is called when production expr_concat is exited.
func (s *BaseHplsqlListener) ExitExpr_concat(ctx *Expr_concatContext) {}

// EnterExpr_concat_item is called when production expr_concat_item is entered.
func (s *BaseHplsqlListener) EnterExpr_concat_item(ctx *Expr_concat_itemContext) {}

// ExitExpr_concat_item is called when production expr_concat_item is exited.
func (s *BaseHplsqlListener) ExitExpr_concat_item(ctx *Expr_concat_itemContext) {}

// EnterExpr_case is called when production expr_case is entered.
func (s *BaseHplsqlListener) EnterExpr_case(ctx *Expr_caseContext) {}

// ExitExpr_case is called when production expr_case is exited.
func (s *BaseHplsqlListener) ExitExpr_case(ctx *Expr_caseContext) {}

// EnterExpr_case_simple is called when production expr_case_simple is entered.
func (s *BaseHplsqlListener) EnterExpr_case_simple(ctx *Expr_case_simpleContext) {}

// ExitExpr_case_simple is called when production expr_case_simple is exited.
func (s *BaseHplsqlListener) ExitExpr_case_simple(ctx *Expr_case_simpleContext) {}

// EnterExpr_case_searched is called when production expr_case_searched is entered.
func (s *BaseHplsqlListener) EnterExpr_case_searched(ctx *Expr_case_searchedContext) {}

// ExitExpr_case_searched is called when production expr_case_searched is exited.
func (s *BaseHplsqlListener) ExitExpr_case_searched(ctx *Expr_case_searchedContext) {}

// EnterExpr_cursor_attribute is called when production expr_cursor_attribute is entered.
func (s *BaseHplsqlListener) EnterExpr_cursor_attribute(ctx *Expr_cursor_attributeContext) {}

// ExitExpr_cursor_attribute is called when production expr_cursor_attribute is exited.
func (s *BaseHplsqlListener) ExitExpr_cursor_attribute(ctx *Expr_cursor_attributeContext) {}

// EnterExpr_agg_window_func is called when production expr_agg_window_func is entered.
func (s *BaseHplsqlListener) EnterExpr_agg_window_func(ctx *Expr_agg_window_funcContext) {}

// ExitExpr_agg_window_func is called when production expr_agg_window_func is exited.
func (s *BaseHplsqlListener) ExitExpr_agg_window_func(ctx *Expr_agg_window_funcContext) {}

// EnterExpr_func_all_distinct is called when production expr_func_all_distinct is entered.
func (s *BaseHplsqlListener) EnterExpr_func_all_distinct(ctx *Expr_func_all_distinctContext) {}

// ExitExpr_func_all_distinct is called when production expr_func_all_distinct is exited.
func (s *BaseHplsqlListener) ExitExpr_func_all_distinct(ctx *Expr_func_all_distinctContext) {}

// EnterExpr_func_over_clause is called when production expr_func_over_clause is entered.
func (s *BaseHplsqlListener) EnterExpr_func_over_clause(ctx *Expr_func_over_clauseContext) {}

// ExitExpr_func_over_clause is called when production expr_func_over_clause is exited.
func (s *BaseHplsqlListener) ExitExpr_func_over_clause(ctx *Expr_func_over_clauseContext) {}

// EnterExpr_func_partition_by_clause is called when production expr_func_partition_by_clause is entered.
func (s *BaseHplsqlListener) EnterExpr_func_partition_by_clause(ctx *Expr_func_partition_by_clauseContext) {
}

// ExitExpr_func_partition_by_clause is called when production expr_func_partition_by_clause is exited.
func (s *BaseHplsqlListener) ExitExpr_func_partition_by_clause(ctx *Expr_func_partition_by_clauseContext) {
}

// EnterExpr_spec_func is called when production expr_spec_func is entered.
func (s *BaseHplsqlListener) EnterExpr_spec_func(ctx *Expr_spec_funcContext) {}

// ExitExpr_spec_func is called when production expr_spec_func is exited.
func (s *BaseHplsqlListener) ExitExpr_spec_func(ctx *Expr_spec_funcContext) {}

// EnterExpr_func is called when production expr_func is entered.
func (s *BaseHplsqlListener) EnterExpr_func(ctx *Expr_funcContext) {}

// ExitExpr_func is called when production expr_func is exited.
func (s *BaseHplsqlListener) ExitExpr_func(ctx *Expr_funcContext) {}

// EnterExpr_dot is called when production expr_dot is entered.
func (s *BaseHplsqlListener) EnterExpr_dot(ctx *Expr_dotContext) {}

// ExitExpr_dot is called when production expr_dot is exited.
func (s *BaseHplsqlListener) ExitExpr_dot(ctx *Expr_dotContext) {}

// EnterExpr_dot_method_call is called when production expr_dot_method_call is entered.
func (s *BaseHplsqlListener) EnterExpr_dot_method_call(ctx *Expr_dot_method_callContext) {}

// ExitExpr_dot_method_call is called when production expr_dot_method_call is exited.
func (s *BaseHplsqlListener) ExitExpr_dot_method_call(ctx *Expr_dot_method_callContext) {}

// EnterExpr_dot_property_access is called when production expr_dot_property_access is entered.
func (s *BaseHplsqlListener) EnterExpr_dot_property_access(ctx *Expr_dot_property_accessContext) {}

// ExitExpr_dot_property_access is called when production expr_dot_property_access is exited.
func (s *BaseHplsqlListener) ExitExpr_dot_property_access(ctx *Expr_dot_property_accessContext) {}

// EnterExpr_func_params is called when production expr_func_params is entered.
func (s *BaseHplsqlListener) EnterExpr_func_params(ctx *Expr_func_paramsContext) {}

// ExitExpr_func_params is called when production expr_func_params is exited.
func (s *BaseHplsqlListener) ExitExpr_func_params(ctx *Expr_func_paramsContext) {}

// EnterFunc_param is called when production func_param is entered.
func (s *BaseHplsqlListener) EnterFunc_param(ctx *Func_paramContext) {}

// ExitFunc_param is called when production func_param is exited.
func (s *BaseHplsqlListener) ExitFunc_param(ctx *Func_paramContext) {}

// EnterExpr_select is called when production expr_select is entered.
func (s *BaseHplsqlListener) EnterExpr_select(ctx *Expr_selectContext) {}

// ExitExpr_select is called when production expr_select is exited.
func (s *BaseHplsqlListener) ExitExpr_select(ctx *Expr_selectContext) {}

// EnterExpr_file is called when production expr_file is entered.
func (s *BaseHplsqlListener) EnterExpr_file(ctx *Expr_fileContext) {}

// ExitExpr_file is called when production expr_file is exited.
func (s *BaseHplsqlListener) ExitExpr_file(ctx *Expr_fileContext) {}

// EnterHive is called when production hive is entered.
func (s *BaseHplsqlListener) EnterHive(ctx *HiveContext) {}

// ExitHive is called when production hive is exited.
func (s *BaseHplsqlListener) ExitHive(ctx *HiveContext) {}

// EnterHive_item is called when production hive_item is entered.
func (s *BaseHplsqlListener) EnterHive_item(ctx *Hive_itemContext) {}

// ExitHive_item is called when production hive_item is exited.
func (s *BaseHplsqlListener) ExitHive_item(ctx *Hive_itemContext) {}

// EnterHost is called when production host is entered.
func (s *BaseHplsqlListener) EnterHost(ctx *HostContext) {}

// ExitHost is called when production host is exited.
func (s *BaseHplsqlListener) ExitHost(ctx *HostContext) {}

// EnterHost_cmd is called when production host_cmd is entered.
func (s *BaseHplsqlListener) EnterHost_cmd(ctx *Host_cmdContext) {}

// ExitHost_cmd is called when production host_cmd is exited.
func (s *BaseHplsqlListener) ExitHost_cmd(ctx *Host_cmdContext) {}

// EnterHost_stmt is called when production host_stmt is entered.
func (s *BaseHplsqlListener) EnterHost_stmt(ctx *Host_stmtContext) {}

// ExitHost_stmt is called when production host_stmt is exited.
func (s *BaseHplsqlListener) ExitHost_stmt(ctx *Host_stmtContext) {}

// EnterFile_name is called when production file_name is entered.
func (s *BaseHplsqlListener) EnterFile_name(ctx *File_nameContext) {}

// ExitFile_name is called when production file_name is exited.
func (s *BaseHplsqlListener) ExitFile_name(ctx *File_nameContext) {}

// EnterDate_literal is called when production date_literal is entered.
func (s *BaseHplsqlListener) EnterDate_literal(ctx *Date_literalContext) {}

// ExitDate_literal is called when production date_literal is exited.
func (s *BaseHplsqlListener) ExitDate_literal(ctx *Date_literalContext) {}

// EnterTimestamp_literal is called when production timestamp_literal is entered.
func (s *BaseHplsqlListener) EnterTimestamp_literal(ctx *Timestamp_literalContext) {}

// ExitTimestamp_literal is called when production timestamp_literal is exited.
func (s *BaseHplsqlListener) ExitTimestamp_literal(ctx *Timestamp_literalContext) {}

// EnterIdent is called when production ident is entered.
func (s *BaseHplsqlListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseHplsqlListener) ExitIdent(ctx *IdentContext) {}

// EnterQident is called when production qident is entered.
func (s *BaseHplsqlListener) EnterQident(ctx *QidentContext) {}

// ExitQident is called when production qident is exited.
func (s *BaseHplsqlListener) ExitQident(ctx *QidentContext) {}

// EnterSingle_quotedString is called when production single_quotedString is entered.
func (s *BaseHplsqlListener) EnterSingle_quotedString(ctx *Single_quotedStringContext) {}

// ExitSingle_quotedString is called when production single_quotedString is exited.
func (s *BaseHplsqlListener) ExitSingle_quotedString(ctx *Single_quotedStringContext) {}

// EnterDouble_quotedString is called when production double_quotedString is entered.
func (s *BaseHplsqlListener) EnterDouble_quotedString(ctx *Double_quotedStringContext) {}

// ExitDouble_quotedString is called when production double_quotedString is exited.
func (s *BaseHplsqlListener) ExitDouble_quotedString(ctx *Double_quotedStringContext) {}

// EnterInt_number is called when production int_number is entered.
func (s *BaseHplsqlListener) EnterInt_number(ctx *Int_numberContext) {}

// ExitInt_number is called when production int_number is exited.
func (s *BaseHplsqlListener) ExitInt_number(ctx *Int_numberContext) {}

// EnterDec_number is called when production dec_number is entered.
func (s *BaseHplsqlListener) EnterDec_number(ctx *Dec_numberContext) {}

// ExitDec_number is called when production dec_number is exited.
func (s *BaseHplsqlListener) ExitDec_number(ctx *Dec_numberContext) {}

// EnterBool_literal is called when production bool_literal is entered.
func (s *BaseHplsqlListener) EnterBool_literal(ctx *Bool_literalContext) {}

// ExitBool_literal is called when production bool_literal is exited.
func (s *BaseHplsqlListener) ExitBool_literal(ctx *Bool_literalContext) {}

// EnterNull_const is called when production null_const is entered.
func (s *BaseHplsqlListener) EnterNull_const(ctx *Null_constContext) {}

// ExitNull_const is called when production null_const is exited.
func (s *BaseHplsqlListener) ExitNull_const(ctx *Null_constContext) {}

// EnterNon_reserved_words is called when production non_reserved_words is entered.
func (s *BaseHplsqlListener) EnterNon_reserved_words(ctx *Non_reserved_wordsContext) {}

// ExitNon_reserved_words is called when production non_reserved_words is exited.
func (s *BaseHplsqlListener) ExitNon_reserved_words(ctx *Non_reserved_wordsContext) {}
