// Code generated from Hplsql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Hplsql

import "github.com/antlr4-go/antlr/v4"

// HplsqlListener is a complete listener for a parse tree produced by HplsqlParser.
type HplsqlListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBegin_end_block is called when entering the begin_end_block production.
	EnterBegin_end_block(c *Begin_end_blockContext)

	// EnterSingle_block_stmt is called when entering the single_block_stmt production.
	EnterSingle_block_stmt(c *Single_block_stmtContext)

	// EnterBlock_end is called when entering the block_end production.
	EnterBlock_end(c *Block_endContext)

	// EnterProc_block is called when entering the proc_block production.
	EnterProc_block(c *Proc_blockContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterSemicolon_stmt is called when entering the semicolon_stmt production.
	EnterSemicolon_stmt(c *Semicolon_stmtContext)

	// EnterException_block is called when entering the exception_block production.
	EnterException_block(c *Exception_blockContext)

	// EnterException_block_item is called when entering the exception_block_item production.
	EnterException_block_item(c *Exception_block_itemContext)

	// EnterNull_stmt is called when entering the null_stmt production.
	EnterNull_stmt(c *Null_stmtContext)

	// EnterExpr_stmt is called when entering the expr_stmt production.
	EnterExpr_stmt(c *Expr_stmtContext)

	// EnterAssignment_stmt is called when entering the assignment_stmt production.
	EnterAssignment_stmt(c *Assignment_stmtContext)

	// EnterAssignment_stmt_item is called when entering the assignment_stmt_item production.
	EnterAssignment_stmt_item(c *Assignment_stmt_itemContext)

	// EnterAssignment_stmt_single_item is called when entering the assignment_stmt_single_item production.
	EnterAssignment_stmt_single_item(c *Assignment_stmt_single_itemContext)

	// EnterAssignment_stmt_collection_item is called when entering the assignment_stmt_collection_item production.
	EnterAssignment_stmt_collection_item(c *Assignment_stmt_collection_itemContext)

	// EnterAssignment_stmt_multiple_item is called when entering the assignment_stmt_multiple_item production.
	EnterAssignment_stmt_multiple_item(c *Assignment_stmt_multiple_itemContext)

	// EnterAssignment_stmt_select_item is called when entering the assignment_stmt_select_item production.
	EnterAssignment_stmt_select_item(c *Assignment_stmt_select_itemContext)

	// EnterAllocate_cursor_stmt is called when entering the allocate_cursor_stmt production.
	EnterAllocate_cursor_stmt(c *Allocate_cursor_stmtContext)

	// EnterAssociate_locator_stmt is called when entering the associate_locator_stmt production.
	EnterAssociate_locator_stmt(c *Associate_locator_stmtContext)

	// EnterBegin_transaction_stmt is called when entering the begin_transaction_stmt production.
	EnterBegin_transaction_stmt(c *Begin_transaction_stmtContext)

	// EnterBreak_stmt is called when entering the break_stmt production.
	EnterBreak_stmt(c *Break_stmtContext)

	// EnterCall_stmt is called when entering the call_stmt production.
	EnterCall_stmt(c *Call_stmtContext)

	// EnterDeclare_stmt is called when entering the declare_stmt production.
	EnterDeclare_stmt(c *Declare_stmtContext)

	// EnterDeclare_block is called when entering the declare_block production.
	EnterDeclare_block(c *Declare_blockContext)

	// EnterDeclare_block_inplace is called when entering the declare_block_inplace production.
	EnterDeclare_block_inplace(c *Declare_block_inplaceContext)

	// EnterDeclare_stmt_item is called when entering the declare_stmt_item production.
	EnterDeclare_stmt_item(c *Declare_stmt_itemContext)

	// EnterDeclare_var_item is called when entering the declare_var_item production.
	EnterDeclare_var_item(c *Declare_var_itemContext)

	// EnterDeclare_condition_item is called when entering the declare_condition_item production.
	EnterDeclare_condition_item(c *Declare_condition_itemContext)

	// EnterDeclare_cursor_item is called when entering the declare_cursor_item production.
	EnterDeclare_cursor_item(c *Declare_cursor_itemContext)

	// EnterCursor_with_return is called when entering the cursor_with_return production.
	EnterCursor_with_return(c *Cursor_with_returnContext)

	// EnterCursor_without_return is called when entering the cursor_without_return production.
	EnterCursor_without_return(c *Cursor_without_returnContext)

	// EnterDeclare_handler_item is called when entering the declare_handler_item production.
	EnterDeclare_handler_item(c *Declare_handler_itemContext)

	// EnterDeclare_temporary_table_item is called when entering the declare_temporary_table_item production.
	EnterDeclare_temporary_table_item(c *Declare_temporary_table_itemContext)

	// EnterCreate_table_stmt is called when entering the create_table_stmt production.
	EnterCreate_table_stmt(c *Create_table_stmtContext)

	// EnterCreate_local_temp_table_stmt is called when entering the create_local_temp_table_stmt production.
	EnterCreate_local_temp_table_stmt(c *Create_local_temp_table_stmtContext)

	// EnterCreate_table_definition is called when entering the create_table_definition production.
	EnterCreate_table_definition(c *Create_table_definitionContext)

	// EnterCreate_table_column_comment is called when entering the create_table_column_comment production.
	EnterCreate_table_column_comment(c *Create_table_column_commentContext)

	// EnterCreate_table_columns is called when entering the create_table_columns production.
	EnterCreate_table_columns(c *Create_table_columnsContext)

	// EnterCreate_table_columns_item is called when entering the create_table_columns_item production.
	EnterCreate_table_columns_item(c *Create_table_columns_itemContext)

	// EnterCreate_table_type_stmt is called when entering the create_table_type_stmt production.
	EnterCreate_table_type_stmt(c *Create_table_type_stmtContext)

	// EnterTbl_type is called when entering the tbl_type production.
	EnterTbl_type(c *Tbl_typeContext)

	// EnterSql_type is called when entering the sql_type production.
	EnterSql_type(c *Sql_typeContext)

	// EnterColumn_name is called when entering the column_name production.
	EnterColumn_name(c *Column_nameContext)

	// EnterCreate_table_column_inline_cons is called when entering the create_table_column_inline_cons production.
	EnterCreate_table_column_inline_cons(c *Create_table_column_inline_consContext)

	// EnterCreate_table_column_cons is called when entering the create_table_column_cons production.
	EnterCreate_table_column_cons(c *Create_table_column_consContext)

	// EnterCreate_table_fk_action is called when entering the create_table_fk_action production.
	EnterCreate_table_fk_action(c *Create_table_fk_actionContext)

	// EnterCreate_table_preoptions is called when entering the create_table_preoptions production.
	EnterCreate_table_preoptions(c *Create_table_preoptionsContext)

	// EnterCreate_table_preoptions_item is called when entering the create_table_preoptions_item production.
	EnterCreate_table_preoptions_item(c *Create_table_preoptions_itemContext)

	// EnterCreate_table_preoptions_td_item is called when entering the create_table_preoptions_td_item production.
	EnterCreate_table_preoptions_td_item(c *Create_table_preoptions_td_itemContext)

	// EnterCreate_table_options is called when entering the create_table_options production.
	EnterCreate_table_options(c *Create_table_optionsContext)

	// EnterCreate_table_options_item is called when entering the create_table_options_item production.
	EnterCreate_table_options_item(c *Create_table_options_itemContext)

	// EnterCreate_table_options_ora_item is called when entering the create_table_options_ora_item production.
	EnterCreate_table_options_ora_item(c *Create_table_options_ora_itemContext)

	// EnterCreate_table_options_db2_item is called when entering the create_table_options_db2_item production.
	EnterCreate_table_options_db2_item(c *Create_table_options_db2_itemContext)

	// EnterCreate_table_options_td_item is called when entering the create_table_options_td_item production.
	EnterCreate_table_options_td_item(c *Create_table_options_td_itemContext)

	// EnterCreate_table_options_hive_item is called when entering the create_table_options_hive_item production.
	EnterCreate_table_options_hive_item(c *Create_table_options_hive_itemContext)

	// EnterCreate_table_hive_row_format is called when entering the create_table_hive_row_format production.
	EnterCreate_table_hive_row_format(c *Create_table_hive_row_formatContext)

	// EnterCreate_table_hive_row_format_fields is called when entering the create_table_hive_row_format_fields production.
	EnterCreate_table_hive_row_format_fields(c *Create_table_hive_row_format_fieldsContext)

	// EnterCreate_table_options_mssql_item is called when entering the create_table_options_mssql_item production.
	EnterCreate_table_options_mssql_item(c *Create_table_options_mssql_itemContext)

	// EnterCreate_table_options_mysql_item is called when entering the create_table_options_mysql_item production.
	EnterCreate_table_options_mysql_item(c *Create_table_options_mysql_itemContext)

	// EnterAlter_table_stmt is called when entering the alter_table_stmt production.
	EnterAlter_table_stmt(c *Alter_table_stmtContext)

	// EnterAlter_table_item is called when entering the alter_table_item production.
	EnterAlter_table_item(c *Alter_table_itemContext)

	// EnterAlter_table_add_constraint is called when entering the alter_table_add_constraint production.
	EnterAlter_table_add_constraint(c *Alter_table_add_constraintContext)

	// EnterAlter_table_add_constraint_item is called when entering the alter_table_add_constraint_item production.
	EnterAlter_table_add_constraint_item(c *Alter_table_add_constraint_itemContext)

	// EnterDtype is called when entering the dtype production.
	EnterDtype(c *DtypeContext)

	// EnterDtype_len is called when entering the dtype_len production.
	EnterDtype_len(c *Dtype_lenContext)

	// EnterDtype_attr is called when entering the dtype_attr production.
	EnterDtype_attr(c *Dtype_attrContext)

	// EnterDtype_default is called when entering the dtype_default production.
	EnterDtype_default(c *Dtype_defaultContext)

	// EnterCreate_database_stmt is called when entering the create_database_stmt production.
	EnterCreate_database_stmt(c *Create_database_stmtContext)

	// EnterCreate_database_option is called when entering the create_database_option production.
	EnterCreate_database_option(c *Create_database_optionContext)

	// EnterCreate_function_stmt is called when entering the create_function_stmt production.
	EnterCreate_function_stmt(c *Create_function_stmtContext)

	// EnterCreate_function_return is called when entering the create_function_return production.
	EnterCreate_function_return(c *Create_function_returnContext)

	// EnterCreate_package_stmt is called when entering the create_package_stmt production.
	EnterCreate_package_stmt(c *Create_package_stmtContext)

	// EnterPackage_spec is called when entering the package_spec production.
	EnterPackage_spec(c *Package_specContext)

	// EnterPackage_spec_item is called when entering the package_spec_item production.
	EnterPackage_spec_item(c *Package_spec_itemContext)

	// EnterCreate_package_body_stmt is called when entering the create_package_body_stmt production.
	EnterCreate_package_body_stmt(c *Create_package_body_stmtContext)

	// EnterPackage_body is called when entering the package_body production.
	EnterPackage_body(c *Package_bodyContext)

	// EnterPackage_body_item is called when entering the package_body_item production.
	EnterPackage_body_item(c *Package_body_itemContext)

	// EnterCreate_procedure_stmt is called when entering the create_procedure_stmt production.
	EnterCreate_procedure_stmt(c *Create_procedure_stmtContext)

	// EnterCreate_routine_params is called when entering the create_routine_params production.
	EnterCreate_routine_params(c *Create_routine_paramsContext)

	// EnterCreate_routine_param_item is called when entering the create_routine_param_item production.
	EnterCreate_routine_param_item(c *Create_routine_param_itemContext)

	// EnterCreate_routine_options is called when entering the create_routine_options production.
	EnterCreate_routine_options(c *Create_routine_optionsContext)

	// EnterCreate_routine_option is called when entering the create_routine_option production.
	EnterCreate_routine_option(c *Create_routine_optionContext)

	// EnterDrop_stmt is called when entering the drop_stmt production.
	EnterDrop_stmt(c *Drop_stmtContext)

	// EnterEnd_transaction_stmt is called when entering the end_transaction_stmt production.
	EnterEnd_transaction_stmt(c *End_transaction_stmtContext)

	// EnterExec_stmt is called when entering the exec_stmt production.
	EnterExec_stmt(c *Exec_stmtContext)

	// EnterIf_stmt is called when entering the if_stmt production.
	EnterIf_stmt(c *If_stmtContext)

	// EnterIf_plsql_stmt is called when entering the if_plsql_stmt production.
	EnterIf_plsql_stmt(c *If_plsql_stmtContext)

	// EnterIf_tsql_stmt is called when entering the if_tsql_stmt production.
	EnterIf_tsql_stmt(c *If_tsql_stmtContext)

	// EnterIf_bteq_stmt is called when entering the if_bteq_stmt production.
	EnterIf_bteq_stmt(c *If_bteq_stmtContext)

	// EnterElseif_block is called when entering the elseif_block production.
	EnterElseif_block(c *Elseif_blockContext)

	// EnterElse_block is called when entering the else_block production.
	EnterElse_block(c *Else_blockContext)

	// EnterInclude_stmt is called when entering the include_stmt production.
	EnterInclude_stmt(c *Include_stmtContext)

	// EnterInsert_stmt is called when entering the insert_stmt production.
	EnterInsert_stmt(c *Insert_stmtContext)

	// EnterInsert_stmt_cols is called when entering the insert_stmt_cols production.
	EnterInsert_stmt_cols(c *Insert_stmt_colsContext)

	// EnterInsert_stmt_rows is called when entering the insert_stmt_rows production.
	EnterInsert_stmt_rows(c *Insert_stmt_rowsContext)

	// EnterInsert_stmt_row is called when entering the insert_stmt_row production.
	EnterInsert_stmt_row(c *Insert_stmt_rowContext)

	// EnterInsert_directory_stmt is called when entering the insert_directory_stmt production.
	EnterInsert_directory_stmt(c *Insert_directory_stmtContext)

	// EnterExit_stmt is called when entering the exit_stmt production.
	EnterExit_stmt(c *Exit_stmtContext)

	// EnterGet_diag_stmt is called when entering the get_diag_stmt production.
	EnterGet_diag_stmt(c *Get_diag_stmtContext)

	// EnterGet_diag_stmt_item is called when entering the get_diag_stmt_item production.
	EnterGet_diag_stmt_item(c *Get_diag_stmt_itemContext)

	// EnterGet_diag_stmt_exception_item is called when entering the get_diag_stmt_exception_item production.
	EnterGet_diag_stmt_exception_item(c *Get_diag_stmt_exception_itemContext)

	// EnterGet_diag_stmt_rowcount_item is called when entering the get_diag_stmt_rowcount_item production.
	EnterGet_diag_stmt_rowcount_item(c *Get_diag_stmt_rowcount_itemContext)

	// EnterGrant_stmt is called when entering the grant_stmt production.
	EnterGrant_stmt(c *Grant_stmtContext)

	// EnterGrant_stmt_item is called when entering the grant_stmt_item production.
	EnterGrant_stmt_item(c *Grant_stmt_itemContext)

	// EnterLeave_stmt is called when entering the leave_stmt production.
	EnterLeave_stmt(c *Leave_stmtContext)

	// EnterMap_object_stmt is called when entering the map_object_stmt production.
	EnterMap_object_stmt(c *Map_object_stmtContext)

	// EnterOpen_stmt is called when entering the open_stmt production.
	EnterOpen_stmt(c *Open_stmtContext)

	// EnterFetch_stmt is called when entering the fetch_stmt production.
	EnterFetch_stmt(c *Fetch_stmtContext)

	// EnterFetch_limit is called when entering the fetch_limit production.
	EnterFetch_limit(c *Fetch_limitContext)

	// EnterCollect_stats_stmt is called when entering the collect_stats_stmt production.
	EnterCollect_stats_stmt(c *Collect_stats_stmtContext)

	// EnterCollect_stats_clause is called when entering the collect_stats_clause production.
	EnterCollect_stats_clause(c *Collect_stats_clauseContext)

	// EnterClose_stmt is called when entering the close_stmt production.
	EnterClose_stmt(c *Close_stmtContext)

	// EnterCmp_stmt is called when entering the cmp_stmt production.
	EnterCmp_stmt(c *Cmp_stmtContext)

	// EnterCmp_source is called when entering the cmp_source production.
	EnterCmp_source(c *Cmp_sourceContext)

	// EnterCopy_from_local_stmt is called when entering the copy_from_local_stmt production.
	EnterCopy_from_local_stmt(c *Copy_from_local_stmtContext)

	// EnterCopy_stmt is called when entering the copy_stmt production.
	EnterCopy_stmt(c *Copy_stmtContext)

	// EnterCopy_source is called when entering the copy_source production.
	EnterCopy_source(c *Copy_sourceContext)

	// EnterCopy_target is called when entering the copy_target production.
	EnterCopy_target(c *Copy_targetContext)

	// EnterCopy_option is called when entering the copy_option production.
	EnterCopy_option(c *Copy_optionContext)

	// EnterCopy_file_option is called when entering the copy_file_option production.
	EnterCopy_file_option(c *Copy_file_optionContext)

	// EnterCommit_stmt is called when entering the commit_stmt production.
	EnterCommit_stmt(c *Commit_stmtContext)

	// EnterCreate_index_stmt is called when entering the create_index_stmt production.
	EnterCreate_index_stmt(c *Create_index_stmtContext)

	// EnterCreate_index_col is called when entering the create_index_col production.
	EnterCreate_index_col(c *Create_index_colContext)

	// EnterIndex_storage_clause is called when entering the index_storage_clause production.
	EnterIndex_storage_clause(c *Index_storage_clauseContext)

	// EnterIndex_mssql_storage_clause is called when entering the index_mssql_storage_clause production.
	EnterIndex_mssql_storage_clause(c *Index_mssql_storage_clauseContext)

	// EnterPrint_stmt is called when entering the print_stmt production.
	EnterPrint_stmt(c *Print_stmtContext)

	// EnterQuit_stmt is called when entering the quit_stmt production.
	EnterQuit_stmt(c *Quit_stmtContext)

	// EnterRaise_stmt is called when entering the raise_stmt production.
	EnterRaise_stmt(c *Raise_stmtContext)

	// EnterResignal_stmt is called when entering the resignal_stmt production.
	EnterResignal_stmt(c *Resignal_stmtContext)

	// EnterReturn_stmt is called when entering the return_stmt production.
	EnterReturn_stmt(c *Return_stmtContext)

	// EnterRollback_stmt is called when entering the rollback_stmt production.
	EnterRollback_stmt(c *Rollback_stmtContext)

	// EnterSet_session_option is called when entering the set_session_option production.
	EnterSet_session_option(c *Set_session_optionContext)

	// EnterSet_current_schema_option is called when entering the set_current_schema_option production.
	EnterSet_current_schema_option(c *Set_current_schema_optionContext)

	// EnterSet_mssql_session_option is called when entering the set_mssql_session_option production.
	EnterSet_mssql_session_option(c *Set_mssql_session_optionContext)

	// EnterSet_teradata_session_option is called when entering the set_teradata_session_option production.
	EnterSet_teradata_session_option(c *Set_teradata_session_optionContext)

	// EnterSignal_stmt is called when entering the signal_stmt production.
	EnterSignal_stmt(c *Signal_stmtContext)

	// EnterSummary_stmt is called when entering the summary_stmt production.
	EnterSummary_stmt(c *Summary_stmtContext)

	// EnterTruncate_stmt is called when entering the truncate_stmt production.
	EnterTruncate_stmt(c *Truncate_stmtContext)

	// EnterUse_stmt is called when entering the use_stmt production.
	EnterUse_stmt(c *Use_stmtContext)

	// EnterValues_into_stmt is called when entering the values_into_stmt production.
	EnterValues_into_stmt(c *Values_into_stmtContext)

	// EnterWhile_stmt is called when entering the while_stmt production.
	EnterWhile_stmt(c *While_stmtContext)

	// EnterUnconditional_loop_stmt is called when entering the unconditional_loop_stmt production.
	EnterUnconditional_loop_stmt(c *Unconditional_loop_stmtContext)

	// EnterFor_cursor_stmt is called when entering the for_cursor_stmt production.
	EnterFor_cursor_stmt(c *For_cursor_stmtContext)

	// EnterFor_range_stmt is called when entering the for_range_stmt production.
	EnterFor_range_stmt(c *For_range_stmtContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterUsing_clause is called when entering the using_clause production.
	EnterUsing_clause(c *Using_clauseContext)

	// EnterSelect_stmt is called when entering the select_stmt production.
	EnterSelect_stmt(c *Select_stmtContext)

	// EnterCte_select_stmt is called when entering the cte_select_stmt production.
	EnterCte_select_stmt(c *Cte_select_stmtContext)

	// EnterCte_select_stmt_item is called when entering the cte_select_stmt_item production.
	EnterCte_select_stmt_item(c *Cte_select_stmt_itemContext)

	// EnterCte_select_cols is called when entering the cte_select_cols production.
	EnterCte_select_cols(c *Cte_select_colsContext)

	// EnterFullselect_stmt is called when entering the fullselect_stmt production.
	EnterFullselect_stmt(c *Fullselect_stmtContext)

	// EnterFullselect_stmt_item is called when entering the fullselect_stmt_item production.
	EnterFullselect_stmt_item(c *Fullselect_stmt_itemContext)

	// EnterFullselect_set_clause is called when entering the fullselect_set_clause production.
	EnterFullselect_set_clause(c *Fullselect_set_clauseContext)

	// EnterSubselect_stmt is called when entering the subselect_stmt production.
	EnterSubselect_stmt(c *Subselect_stmtContext)

	// EnterSelect_list is called when entering the select_list production.
	EnterSelect_list(c *Select_listContext)

	// EnterSelect_list_set is called when entering the select_list_set production.
	EnterSelect_list_set(c *Select_list_setContext)

	// EnterSelect_list_limit is called when entering the select_list_limit production.
	EnterSelect_list_limit(c *Select_list_limitContext)

	// EnterSelect_list_item is called when entering the select_list_item production.
	EnterSelect_list_item(c *Select_list_itemContext)

	// EnterSelect_list_alias is called when entering the select_list_alias production.
	EnterSelect_list_alias(c *Select_list_aliasContext)

	// EnterSelect_list_asterisk is called when entering the select_list_asterisk production.
	EnterSelect_list_asterisk(c *Select_list_asteriskContext)

	// EnterTable_row is called when entering the table_row production.
	EnterTable_row(c *Table_rowContext)

	// EnterInto_clause is called when entering the into_clause production.
	EnterInto_clause(c *Into_clauseContext)

	// EnterBulk_collect_clause is called when entering the bulk_collect_clause production.
	EnterBulk_collect_clause(c *Bulk_collect_clauseContext)

	// EnterFrom_clause is called when entering the from_clause production.
	EnterFrom_clause(c *From_clauseContext)

	// EnterFrom_table_clause is called when entering the from_table_clause production.
	EnterFrom_table_clause(c *From_table_clauseContext)

	// EnterFrom_table_name_clause is called when entering the from_table_name_clause production.
	EnterFrom_table_name_clause(c *From_table_name_clauseContext)

	// EnterFrom_subselect_clause is called when entering the from_subselect_clause production.
	EnterFrom_subselect_clause(c *From_subselect_clauseContext)

	// EnterFrom_join_clause is called when entering the from_join_clause production.
	EnterFrom_join_clause(c *From_join_clauseContext)

	// EnterFrom_join_type_clause is called when entering the from_join_type_clause production.
	EnterFrom_join_type_clause(c *From_join_type_clauseContext)

	// EnterFrom_table_values_clause is called when entering the from_table_values_clause production.
	EnterFrom_table_values_clause(c *From_table_values_clauseContext)

	// EnterFrom_table_values_row is called when entering the from_table_values_row production.
	EnterFrom_table_values_row(c *From_table_values_rowContext)

	// EnterFrom_alias_clause is called when entering the from_alias_clause production.
	EnterFrom_alias_clause(c *From_alias_clauseContext)

	// EnterTable_name is called when entering the table_name production.
	EnterTable_name(c *Table_nameContext)

	// EnterWhere_clause is called when entering the where_clause production.
	EnterWhere_clause(c *Where_clauseContext)

	// EnterGroup_by_clause is called when entering the group_by_clause production.
	EnterGroup_by_clause(c *Group_by_clauseContext)

	// EnterHaving_clause is called when entering the having_clause production.
	EnterHaving_clause(c *Having_clauseContext)

	// EnterQualify_clause is called when entering the qualify_clause production.
	EnterQualify_clause(c *Qualify_clauseContext)

	// EnterOrder_by_clause is called when entering the order_by_clause production.
	EnterOrder_by_clause(c *Order_by_clauseContext)

	// EnterSelect_options is called when entering the select_options production.
	EnterSelect_options(c *Select_optionsContext)

	// EnterSelect_options_item is called when entering the select_options_item production.
	EnterSelect_options_item(c *Select_options_itemContext)

	// EnterUpdate_stmt is called when entering the update_stmt production.
	EnterUpdate_stmt(c *Update_stmtContext)

	// EnterUpdate_assignment is called when entering the update_assignment production.
	EnterUpdate_assignment(c *Update_assignmentContext)

	// EnterUpdate_table is called when entering the update_table production.
	EnterUpdate_table(c *Update_tableContext)

	// EnterUpdate_upsert is called when entering the update_upsert production.
	EnterUpdate_upsert(c *Update_upsertContext)

	// EnterMerge_stmt is called when entering the merge_stmt production.
	EnterMerge_stmt(c *Merge_stmtContext)

	// EnterMerge_table is called when entering the merge_table production.
	EnterMerge_table(c *Merge_tableContext)

	// EnterMerge_condition is called when entering the merge_condition production.
	EnterMerge_condition(c *Merge_conditionContext)

	// EnterMerge_action is called when entering the merge_action production.
	EnterMerge_action(c *Merge_actionContext)

	// EnterDelete_stmt is called when entering the delete_stmt production.
	EnterDelete_stmt(c *Delete_stmtContext)

	// EnterDelete_alias is called when entering the delete_alias production.
	EnterDelete_alias(c *Delete_aliasContext)

	// EnterDescribe_stmt is called when entering the describe_stmt production.
	EnterDescribe_stmt(c *Describe_stmtContext)

	// EnterBool_expr is called when entering the bool_expr production.
	EnterBool_expr(c *Bool_exprContext)

	// EnterBool_expr_atom is called when entering the bool_expr_atom production.
	EnterBool_expr_atom(c *Bool_expr_atomContext)

	// EnterBool_expr_unary is called when entering the bool_expr_unary production.
	EnterBool_expr_unary(c *Bool_expr_unaryContext)

	// EnterBool_expr_single_in is called when entering the bool_expr_single_in production.
	EnterBool_expr_single_in(c *Bool_expr_single_inContext)

	// EnterBool_expr_multi_in is called when entering the bool_expr_multi_in production.
	EnterBool_expr_multi_in(c *Bool_expr_multi_inContext)

	// EnterBool_expr_binary is called when entering the bool_expr_binary production.
	EnterBool_expr_binary(c *Bool_expr_binaryContext)

	// EnterBool_expr_logical_operator is called when entering the bool_expr_logical_operator production.
	EnterBool_expr_logical_operator(c *Bool_expr_logical_operatorContext)

	// EnterBool_expr_binary_operator is called when entering the bool_expr_binary_operator production.
	EnterBool_expr_binary_operator(c *Bool_expr_binary_operatorContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExpr_atom is called when entering the expr_atom production.
	EnterExpr_atom(c *Expr_atomContext)

	// EnterExpr_interval is called when entering the expr_interval production.
	EnterExpr_interval(c *Expr_intervalContext)

	// EnterInterval_item is called when entering the interval_item production.
	EnterInterval_item(c *Interval_itemContext)

	// EnterExpr_concat is called when entering the expr_concat production.
	EnterExpr_concat(c *Expr_concatContext)

	// EnterExpr_concat_item is called when entering the expr_concat_item production.
	EnterExpr_concat_item(c *Expr_concat_itemContext)

	// EnterExpr_case is called when entering the expr_case production.
	EnterExpr_case(c *Expr_caseContext)

	// EnterExpr_case_simple is called when entering the expr_case_simple production.
	EnterExpr_case_simple(c *Expr_case_simpleContext)

	// EnterExpr_case_searched is called when entering the expr_case_searched production.
	EnterExpr_case_searched(c *Expr_case_searchedContext)

	// EnterExpr_cursor_attribute is called when entering the expr_cursor_attribute production.
	EnterExpr_cursor_attribute(c *Expr_cursor_attributeContext)

	// EnterExpr_agg_window_func is called when entering the expr_agg_window_func production.
	EnterExpr_agg_window_func(c *Expr_agg_window_funcContext)

	// EnterExpr_func_all_distinct is called when entering the expr_func_all_distinct production.
	EnterExpr_func_all_distinct(c *Expr_func_all_distinctContext)

	// EnterExpr_func_over_clause is called when entering the expr_func_over_clause production.
	EnterExpr_func_over_clause(c *Expr_func_over_clauseContext)

	// EnterExpr_func_partition_by_clause is called when entering the expr_func_partition_by_clause production.
	EnterExpr_func_partition_by_clause(c *Expr_func_partition_by_clauseContext)

	// EnterExpr_spec_func is called when entering the expr_spec_func production.
	EnterExpr_spec_func(c *Expr_spec_funcContext)

	// EnterExpr_func is called when entering the expr_func production.
	EnterExpr_func(c *Expr_funcContext)

	// EnterExpr_dot is called when entering the expr_dot production.
	EnterExpr_dot(c *Expr_dotContext)

	// EnterExpr_dot_method_call is called when entering the expr_dot_method_call production.
	EnterExpr_dot_method_call(c *Expr_dot_method_callContext)

	// EnterExpr_dot_property_access is called when entering the expr_dot_property_access production.
	EnterExpr_dot_property_access(c *Expr_dot_property_accessContext)

	// EnterExpr_func_params is called when entering the expr_func_params production.
	EnterExpr_func_params(c *Expr_func_paramsContext)

	// EnterFunc_param is called when entering the func_param production.
	EnterFunc_param(c *Func_paramContext)

	// EnterExpr_select is called when entering the expr_select production.
	EnterExpr_select(c *Expr_selectContext)

	// EnterExpr_file is called when entering the expr_file production.
	EnterExpr_file(c *Expr_fileContext)

	// EnterHive is called when entering the hive production.
	EnterHive(c *HiveContext)

	// EnterHive_item is called when entering the hive_item production.
	EnterHive_item(c *Hive_itemContext)

	// EnterHost is called when entering the host production.
	EnterHost(c *HostContext)

	// EnterHost_cmd is called when entering the host_cmd production.
	EnterHost_cmd(c *Host_cmdContext)

	// EnterHost_stmt is called when entering the host_stmt production.
	EnterHost_stmt(c *Host_stmtContext)

	// EnterFile_name is called when entering the file_name production.
	EnterFile_name(c *File_nameContext)

	// EnterDate_literal is called when entering the date_literal production.
	EnterDate_literal(c *Date_literalContext)

	// EnterTimestamp_literal is called when entering the timestamp_literal production.
	EnterTimestamp_literal(c *Timestamp_literalContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterQident is called when entering the qident production.
	EnterQident(c *QidentContext)

	// EnterSingle_quotedString is called when entering the single_quotedString production.
	EnterSingle_quotedString(c *Single_quotedStringContext)

	// EnterDouble_quotedString is called when entering the double_quotedString production.
	EnterDouble_quotedString(c *Double_quotedStringContext)

	// EnterInt_number is called when entering the int_number production.
	EnterInt_number(c *Int_numberContext)

	// EnterDec_number is called when entering the dec_number production.
	EnterDec_number(c *Dec_numberContext)

	// EnterBool_literal is called when entering the bool_literal production.
	EnterBool_literal(c *Bool_literalContext)

	// EnterNull_const is called when entering the null_const production.
	EnterNull_const(c *Null_constContext)

	// EnterNon_reserved_words is called when entering the non_reserved_words production.
	EnterNon_reserved_words(c *Non_reserved_wordsContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBegin_end_block is called when exiting the begin_end_block production.
	ExitBegin_end_block(c *Begin_end_blockContext)

	// ExitSingle_block_stmt is called when exiting the single_block_stmt production.
	ExitSingle_block_stmt(c *Single_block_stmtContext)

	// ExitBlock_end is called when exiting the block_end production.
	ExitBlock_end(c *Block_endContext)

	// ExitProc_block is called when exiting the proc_block production.
	ExitProc_block(c *Proc_blockContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitSemicolon_stmt is called when exiting the semicolon_stmt production.
	ExitSemicolon_stmt(c *Semicolon_stmtContext)

	// ExitException_block is called when exiting the exception_block production.
	ExitException_block(c *Exception_blockContext)

	// ExitException_block_item is called when exiting the exception_block_item production.
	ExitException_block_item(c *Exception_block_itemContext)

	// ExitNull_stmt is called when exiting the null_stmt production.
	ExitNull_stmt(c *Null_stmtContext)

	// ExitExpr_stmt is called when exiting the expr_stmt production.
	ExitExpr_stmt(c *Expr_stmtContext)

	// ExitAssignment_stmt is called when exiting the assignment_stmt production.
	ExitAssignment_stmt(c *Assignment_stmtContext)

	// ExitAssignment_stmt_item is called when exiting the assignment_stmt_item production.
	ExitAssignment_stmt_item(c *Assignment_stmt_itemContext)

	// ExitAssignment_stmt_single_item is called when exiting the assignment_stmt_single_item production.
	ExitAssignment_stmt_single_item(c *Assignment_stmt_single_itemContext)

	// ExitAssignment_stmt_collection_item is called when exiting the assignment_stmt_collection_item production.
	ExitAssignment_stmt_collection_item(c *Assignment_stmt_collection_itemContext)

	// ExitAssignment_stmt_multiple_item is called when exiting the assignment_stmt_multiple_item production.
	ExitAssignment_stmt_multiple_item(c *Assignment_stmt_multiple_itemContext)

	// ExitAssignment_stmt_select_item is called when exiting the assignment_stmt_select_item production.
	ExitAssignment_stmt_select_item(c *Assignment_stmt_select_itemContext)

	// ExitAllocate_cursor_stmt is called when exiting the allocate_cursor_stmt production.
	ExitAllocate_cursor_stmt(c *Allocate_cursor_stmtContext)

	// ExitAssociate_locator_stmt is called when exiting the associate_locator_stmt production.
	ExitAssociate_locator_stmt(c *Associate_locator_stmtContext)

	// ExitBegin_transaction_stmt is called when exiting the begin_transaction_stmt production.
	ExitBegin_transaction_stmt(c *Begin_transaction_stmtContext)

	// ExitBreak_stmt is called when exiting the break_stmt production.
	ExitBreak_stmt(c *Break_stmtContext)

	// ExitCall_stmt is called when exiting the call_stmt production.
	ExitCall_stmt(c *Call_stmtContext)

	// ExitDeclare_stmt is called when exiting the declare_stmt production.
	ExitDeclare_stmt(c *Declare_stmtContext)

	// ExitDeclare_block is called when exiting the declare_block production.
	ExitDeclare_block(c *Declare_blockContext)

	// ExitDeclare_block_inplace is called when exiting the declare_block_inplace production.
	ExitDeclare_block_inplace(c *Declare_block_inplaceContext)

	// ExitDeclare_stmt_item is called when exiting the declare_stmt_item production.
	ExitDeclare_stmt_item(c *Declare_stmt_itemContext)

	// ExitDeclare_var_item is called when exiting the declare_var_item production.
	ExitDeclare_var_item(c *Declare_var_itemContext)

	// ExitDeclare_condition_item is called when exiting the declare_condition_item production.
	ExitDeclare_condition_item(c *Declare_condition_itemContext)

	// ExitDeclare_cursor_item is called when exiting the declare_cursor_item production.
	ExitDeclare_cursor_item(c *Declare_cursor_itemContext)

	// ExitCursor_with_return is called when exiting the cursor_with_return production.
	ExitCursor_with_return(c *Cursor_with_returnContext)

	// ExitCursor_without_return is called when exiting the cursor_without_return production.
	ExitCursor_without_return(c *Cursor_without_returnContext)

	// ExitDeclare_handler_item is called when exiting the declare_handler_item production.
	ExitDeclare_handler_item(c *Declare_handler_itemContext)

	// ExitDeclare_temporary_table_item is called when exiting the declare_temporary_table_item production.
	ExitDeclare_temporary_table_item(c *Declare_temporary_table_itemContext)

	// ExitCreate_table_stmt is called when exiting the create_table_stmt production.
	ExitCreate_table_stmt(c *Create_table_stmtContext)

	// ExitCreate_local_temp_table_stmt is called when exiting the create_local_temp_table_stmt production.
	ExitCreate_local_temp_table_stmt(c *Create_local_temp_table_stmtContext)

	// ExitCreate_table_definition is called when exiting the create_table_definition production.
	ExitCreate_table_definition(c *Create_table_definitionContext)

	// ExitCreate_table_column_comment is called when exiting the create_table_column_comment production.
	ExitCreate_table_column_comment(c *Create_table_column_commentContext)

	// ExitCreate_table_columns is called when exiting the create_table_columns production.
	ExitCreate_table_columns(c *Create_table_columnsContext)

	// ExitCreate_table_columns_item is called when exiting the create_table_columns_item production.
	ExitCreate_table_columns_item(c *Create_table_columns_itemContext)

	// ExitCreate_table_type_stmt is called when exiting the create_table_type_stmt production.
	ExitCreate_table_type_stmt(c *Create_table_type_stmtContext)

	// ExitTbl_type is called when exiting the tbl_type production.
	ExitTbl_type(c *Tbl_typeContext)

	// ExitSql_type is called when exiting the sql_type production.
	ExitSql_type(c *Sql_typeContext)

	// ExitColumn_name is called when exiting the column_name production.
	ExitColumn_name(c *Column_nameContext)

	// ExitCreate_table_column_inline_cons is called when exiting the create_table_column_inline_cons production.
	ExitCreate_table_column_inline_cons(c *Create_table_column_inline_consContext)

	// ExitCreate_table_column_cons is called when exiting the create_table_column_cons production.
	ExitCreate_table_column_cons(c *Create_table_column_consContext)

	// ExitCreate_table_fk_action is called when exiting the create_table_fk_action production.
	ExitCreate_table_fk_action(c *Create_table_fk_actionContext)

	// ExitCreate_table_preoptions is called when exiting the create_table_preoptions production.
	ExitCreate_table_preoptions(c *Create_table_preoptionsContext)

	// ExitCreate_table_preoptions_item is called when exiting the create_table_preoptions_item production.
	ExitCreate_table_preoptions_item(c *Create_table_preoptions_itemContext)

	// ExitCreate_table_preoptions_td_item is called when exiting the create_table_preoptions_td_item production.
	ExitCreate_table_preoptions_td_item(c *Create_table_preoptions_td_itemContext)

	// ExitCreate_table_options is called when exiting the create_table_options production.
	ExitCreate_table_options(c *Create_table_optionsContext)

	// ExitCreate_table_options_item is called when exiting the create_table_options_item production.
	ExitCreate_table_options_item(c *Create_table_options_itemContext)

	// ExitCreate_table_options_ora_item is called when exiting the create_table_options_ora_item production.
	ExitCreate_table_options_ora_item(c *Create_table_options_ora_itemContext)

	// ExitCreate_table_options_db2_item is called when exiting the create_table_options_db2_item production.
	ExitCreate_table_options_db2_item(c *Create_table_options_db2_itemContext)

	// ExitCreate_table_options_td_item is called when exiting the create_table_options_td_item production.
	ExitCreate_table_options_td_item(c *Create_table_options_td_itemContext)

	// ExitCreate_table_options_hive_item is called when exiting the create_table_options_hive_item production.
	ExitCreate_table_options_hive_item(c *Create_table_options_hive_itemContext)

	// ExitCreate_table_hive_row_format is called when exiting the create_table_hive_row_format production.
	ExitCreate_table_hive_row_format(c *Create_table_hive_row_formatContext)

	// ExitCreate_table_hive_row_format_fields is called when exiting the create_table_hive_row_format_fields production.
	ExitCreate_table_hive_row_format_fields(c *Create_table_hive_row_format_fieldsContext)

	// ExitCreate_table_options_mssql_item is called when exiting the create_table_options_mssql_item production.
	ExitCreate_table_options_mssql_item(c *Create_table_options_mssql_itemContext)

	// ExitCreate_table_options_mysql_item is called when exiting the create_table_options_mysql_item production.
	ExitCreate_table_options_mysql_item(c *Create_table_options_mysql_itemContext)

	// ExitAlter_table_stmt is called when exiting the alter_table_stmt production.
	ExitAlter_table_stmt(c *Alter_table_stmtContext)

	// ExitAlter_table_item is called when exiting the alter_table_item production.
	ExitAlter_table_item(c *Alter_table_itemContext)

	// ExitAlter_table_add_constraint is called when exiting the alter_table_add_constraint production.
	ExitAlter_table_add_constraint(c *Alter_table_add_constraintContext)

	// ExitAlter_table_add_constraint_item is called when exiting the alter_table_add_constraint_item production.
	ExitAlter_table_add_constraint_item(c *Alter_table_add_constraint_itemContext)

	// ExitDtype is called when exiting the dtype production.
	ExitDtype(c *DtypeContext)

	// ExitDtype_len is called when exiting the dtype_len production.
	ExitDtype_len(c *Dtype_lenContext)

	// ExitDtype_attr is called when exiting the dtype_attr production.
	ExitDtype_attr(c *Dtype_attrContext)

	// ExitDtype_default is called when exiting the dtype_default production.
	ExitDtype_default(c *Dtype_defaultContext)

	// ExitCreate_database_stmt is called when exiting the create_database_stmt production.
	ExitCreate_database_stmt(c *Create_database_stmtContext)

	// ExitCreate_database_option is called when exiting the create_database_option production.
	ExitCreate_database_option(c *Create_database_optionContext)

	// ExitCreate_function_stmt is called when exiting the create_function_stmt production.
	ExitCreate_function_stmt(c *Create_function_stmtContext)

	// ExitCreate_function_return is called when exiting the create_function_return production.
	ExitCreate_function_return(c *Create_function_returnContext)

	// ExitCreate_package_stmt is called when exiting the create_package_stmt production.
	ExitCreate_package_stmt(c *Create_package_stmtContext)

	// ExitPackage_spec is called when exiting the package_spec production.
	ExitPackage_spec(c *Package_specContext)

	// ExitPackage_spec_item is called when exiting the package_spec_item production.
	ExitPackage_spec_item(c *Package_spec_itemContext)

	// ExitCreate_package_body_stmt is called when exiting the create_package_body_stmt production.
	ExitCreate_package_body_stmt(c *Create_package_body_stmtContext)

	// ExitPackage_body is called when exiting the package_body production.
	ExitPackage_body(c *Package_bodyContext)

	// ExitPackage_body_item is called when exiting the package_body_item production.
	ExitPackage_body_item(c *Package_body_itemContext)

	// ExitCreate_procedure_stmt is called when exiting the create_procedure_stmt production.
	ExitCreate_procedure_stmt(c *Create_procedure_stmtContext)

	// ExitCreate_routine_params is called when exiting the create_routine_params production.
	ExitCreate_routine_params(c *Create_routine_paramsContext)

	// ExitCreate_routine_param_item is called when exiting the create_routine_param_item production.
	ExitCreate_routine_param_item(c *Create_routine_param_itemContext)

	// ExitCreate_routine_options is called when exiting the create_routine_options production.
	ExitCreate_routine_options(c *Create_routine_optionsContext)

	// ExitCreate_routine_option is called when exiting the create_routine_option production.
	ExitCreate_routine_option(c *Create_routine_optionContext)

	// ExitDrop_stmt is called when exiting the drop_stmt production.
	ExitDrop_stmt(c *Drop_stmtContext)

	// ExitEnd_transaction_stmt is called when exiting the end_transaction_stmt production.
	ExitEnd_transaction_stmt(c *End_transaction_stmtContext)

	// ExitExec_stmt is called when exiting the exec_stmt production.
	ExitExec_stmt(c *Exec_stmtContext)

	// ExitIf_stmt is called when exiting the if_stmt production.
	ExitIf_stmt(c *If_stmtContext)

	// ExitIf_plsql_stmt is called when exiting the if_plsql_stmt production.
	ExitIf_plsql_stmt(c *If_plsql_stmtContext)

	// ExitIf_tsql_stmt is called when exiting the if_tsql_stmt production.
	ExitIf_tsql_stmt(c *If_tsql_stmtContext)

	// ExitIf_bteq_stmt is called when exiting the if_bteq_stmt production.
	ExitIf_bteq_stmt(c *If_bteq_stmtContext)

	// ExitElseif_block is called when exiting the elseif_block production.
	ExitElseif_block(c *Elseif_blockContext)

	// ExitElse_block is called when exiting the else_block production.
	ExitElse_block(c *Else_blockContext)

	// ExitInclude_stmt is called when exiting the include_stmt production.
	ExitInclude_stmt(c *Include_stmtContext)

	// ExitInsert_stmt is called when exiting the insert_stmt production.
	ExitInsert_stmt(c *Insert_stmtContext)

	// ExitInsert_stmt_cols is called when exiting the insert_stmt_cols production.
	ExitInsert_stmt_cols(c *Insert_stmt_colsContext)

	// ExitInsert_stmt_rows is called when exiting the insert_stmt_rows production.
	ExitInsert_stmt_rows(c *Insert_stmt_rowsContext)

	// ExitInsert_stmt_row is called when exiting the insert_stmt_row production.
	ExitInsert_stmt_row(c *Insert_stmt_rowContext)

	// ExitInsert_directory_stmt is called when exiting the insert_directory_stmt production.
	ExitInsert_directory_stmt(c *Insert_directory_stmtContext)

	// ExitExit_stmt is called when exiting the exit_stmt production.
	ExitExit_stmt(c *Exit_stmtContext)

	// ExitGet_diag_stmt is called when exiting the get_diag_stmt production.
	ExitGet_diag_stmt(c *Get_diag_stmtContext)

	// ExitGet_diag_stmt_item is called when exiting the get_diag_stmt_item production.
	ExitGet_diag_stmt_item(c *Get_diag_stmt_itemContext)

	// ExitGet_diag_stmt_exception_item is called when exiting the get_diag_stmt_exception_item production.
	ExitGet_diag_stmt_exception_item(c *Get_diag_stmt_exception_itemContext)

	// ExitGet_diag_stmt_rowcount_item is called when exiting the get_diag_stmt_rowcount_item production.
	ExitGet_diag_stmt_rowcount_item(c *Get_diag_stmt_rowcount_itemContext)

	// ExitGrant_stmt is called when exiting the grant_stmt production.
	ExitGrant_stmt(c *Grant_stmtContext)

	// ExitGrant_stmt_item is called when exiting the grant_stmt_item production.
	ExitGrant_stmt_item(c *Grant_stmt_itemContext)

	// ExitLeave_stmt is called when exiting the leave_stmt production.
	ExitLeave_stmt(c *Leave_stmtContext)

	// ExitMap_object_stmt is called when exiting the map_object_stmt production.
	ExitMap_object_stmt(c *Map_object_stmtContext)

	// ExitOpen_stmt is called when exiting the open_stmt production.
	ExitOpen_stmt(c *Open_stmtContext)

	// ExitFetch_stmt is called when exiting the fetch_stmt production.
	ExitFetch_stmt(c *Fetch_stmtContext)

	// ExitFetch_limit is called when exiting the fetch_limit production.
	ExitFetch_limit(c *Fetch_limitContext)

	// ExitCollect_stats_stmt is called when exiting the collect_stats_stmt production.
	ExitCollect_stats_stmt(c *Collect_stats_stmtContext)

	// ExitCollect_stats_clause is called when exiting the collect_stats_clause production.
	ExitCollect_stats_clause(c *Collect_stats_clauseContext)

	// ExitClose_stmt is called when exiting the close_stmt production.
	ExitClose_stmt(c *Close_stmtContext)

	// ExitCmp_stmt is called when exiting the cmp_stmt production.
	ExitCmp_stmt(c *Cmp_stmtContext)

	// ExitCmp_source is called when exiting the cmp_source production.
	ExitCmp_source(c *Cmp_sourceContext)

	// ExitCopy_from_local_stmt is called when exiting the copy_from_local_stmt production.
	ExitCopy_from_local_stmt(c *Copy_from_local_stmtContext)

	// ExitCopy_stmt is called when exiting the copy_stmt production.
	ExitCopy_stmt(c *Copy_stmtContext)

	// ExitCopy_source is called when exiting the copy_source production.
	ExitCopy_source(c *Copy_sourceContext)

	// ExitCopy_target is called when exiting the copy_target production.
	ExitCopy_target(c *Copy_targetContext)

	// ExitCopy_option is called when exiting the copy_option production.
	ExitCopy_option(c *Copy_optionContext)

	// ExitCopy_file_option is called when exiting the copy_file_option production.
	ExitCopy_file_option(c *Copy_file_optionContext)

	// ExitCommit_stmt is called when exiting the commit_stmt production.
	ExitCommit_stmt(c *Commit_stmtContext)

	// ExitCreate_index_stmt is called when exiting the create_index_stmt production.
	ExitCreate_index_stmt(c *Create_index_stmtContext)

	// ExitCreate_index_col is called when exiting the create_index_col production.
	ExitCreate_index_col(c *Create_index_colContext)

	// ExitIndex_storage_clause is called when exiting the index_storage_clause production.
	ExitIndex_storage_clause(c *Index_storage_clauseContext)

	// ExitIndex_mssql_storage_clause is called when exiting the index_mssql_storage_clause production.
	ExitIndex_mssql_storage_clause(c *Index_mssql_storage_clauseContext)

	// ExitPrint_stmt is called when exiting the print_stmt production.
	ExitPrint_stmt(c *Print_stmtContext)

	// ExitQuit_stmt is called when exiting the quit_stmt production.
	ExitQuit_stmt(c *Quit_stmtContext)

	// ExitRaise_stmt is called when exiting the raise_stmt production.
	ExitRaise_stmt(c *Raise_stmtContext)

	// ExitResignal_stmt is called when exiting the resignal_stmt production.
	ExitResignal_stmt(c *Resignal_stmtContext)

	// ExitReturn_stmt is called when exiting the return_stmt production.
	ExitReturn_stmt(c *Return_stmtContext)

	// ExitRollback_stmt is called when exiting the rollback_stmt production.
	ExitRollback_stmt(c *Rollback_stmtContext)

	// ExitSet_session_option is called when exiting the set_session_option production.
	ExitSet_session_option(c *Set_session_optionContext)

	// ExitSet_current_schema_option is called when exiting the set_current_schema_option production.
	ExitSet_current_schema_option(c *Set_current_schema_optionContext)

	// ExitSet_mssql_session_option is called when exiting the set_mssql_session_option production.
	ExitSet_mssql_session_option(c *Set_mssql_session_optionContext)

	// ExitSet_teradata_session_option is called when exiting the set_teradata_session_option production.
	ExitSet_teradata_session_option(c *Set_teradata_session_optionContext)

	// ExitSignal_stmt is called when exiting the signal_stmt production.
	ExitSignal_stmt(c *Signal_stmtContext)

	// ExitSummary_stmt is called when exiting the summary_stmt production.
	ExitSummary_stmt(c *Summary_stmtContext)

	// ExitTruncate_stmt is called when exiting the truncate_stmt production.
	ExitTruncate_stmt(c *Truncate_stmtContext)

	// ExitUse_stmt is called when exiting the use_stmt production.
	ExitUse_stmt(c *Use_stmtContext)

	// ExitValues_into_stmt is called when exiting the values_into_stmt production.
	ExitValues_into_stmt(c *Values_into_stmtContext)

	// ExitWhile_stmt is called when exiting the while_stmt production.
	ExitWhile_stmt(c *While_stmtContext)

	// ExitUnconditional_loop_stmt is called when exiting the unconditional_loop_stmt production.
	ExitUnconditional_loop_stmt(c *Unconditional_loop_stmtContext)

	// ExitFor_cursor_stmt is called when exiting the for_cursor_stmt production.
	ExitFor_cursor_stmt(c *For_cursor_stmtContext)

	// ExitFor_range_stmt is called when exiting the for_range_stmt production.
	ExitFor_range_stmt(c *For_range_stmtContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitUsing_clause is called when exiting the using_clause production.
	ExitUsing_clause(c *Using_clauseContext)

	// ExitSelect_stmt is called when exiting the select_stmt production.
	ExitSelect_stmt(c *Select_stmtContext)

	// ExitCte_select_stmt is called when exiting the cte_select_stmt production.
	ExitCte_select_stmt(c *Cte_select_stmtContext)

	// ExitCte_select_stmt_item is called when exiting the cte_select_stmt_item production.
	ExitCte_select_stmt_item(c *Cte_select_stmt_itemContext)

	// ExitCte_select_cols is called when exiting the cte_select_cols production.
	ExitCte_select_cols(c *Cte_select_colsContext)

	// ExitFullselect_stmt is called when exiting the fullselect_stmt production.
	ExitFullselect_stmt(c *Fullselect_stmtContext)

	// ExitFullselect_stmt_item is called when exiting the fullselect_stmt_item production.
	ExitFullselect_stmt_item(c *Fullselect_stmt_itemContext)

	// ExitFullselect_set_clause is called when exiting the fullselect_set_clause production.
	ExitFullselect_set_clause(c *Fullselect_set_clauseContext)

	// ExitSubselect_stmt is called when exiting the subselect_stmt production.
	ExitSubselect_stmt(c *Subselect_stmtContext)

	// ExitSelect_list is called when exiting the select_list production.
	ExitSelect_list(c *Select_listContext)

	// ExitSelect_list_set is called when exiting the select_list_set production.
	ExitSelect_list_set(c *Select_list_setContext)

	// ExitSelect_list_limit is called when exiting the select_list_limit production.
	ExitSelect_list_limit(c *Select_list_limitContext)

	// ExitSelect_list_item is called when exiting the select_list_item production.
	ExitSelect_list_item(c *Select_list_itemContext)

	// ExitSelect_list_alias is called when exiting the select_list_alias production.
	ExitSelect_list_alias(c *Select_list_aliasContext)

	// ExitSelect_list_asterisk is called when exiting the select_list_asterisk production.
	ExitSelect_list_asterisk(c *Select_list_asteriskContext)

	// ExitTable_row is called when exiting the table_row production.
	ExitTable_row(c *Table_rowContext)

	// ExitInto_clause is called when exiting the into_clause production.
	ExitInto_clause(c *Into_clauseContext)

	// ExitBulk_collect_clause is called when exiting the bulk_collect_clause production.
	ExitBulk_collect_clause(c *Bulk_collect_clauseContext)

	// ExitFrom_clause is called when exiting the from_clause production.
	ExitFrom_clause(c *From_clauseContext)

	// ExitFrom_table_clause is called when exiting the from_table_clause production.
	ExitFrom_table_clause(c *From_table_clauseContext)

	// ExitFrom_table_name_clause is called when exiting the from_table_name_clause production.
	ExitFrom_table_name_clause(c *From_table_name_clauseContext)

	// ExitFrom_subselect_clause is called when exiting the from_subselect_clause production.
	ExitFrom_subselect_clause(c *From_subselect_clauseContext)

	// ExitFrom_join_clause is called when exiting the from_join_clause production.
	ExitFrom_join_clause(c *From_join_clauseContext)

	// ExitFrom_join_type_clause is called when exiting the from_join_type_clause production.
	ExitFrom_join_type_clause(c *From_join_type_clauseContext)

	// ExitFrom_table_values_clause is called when exiting the from_table_values_clause production.
	ExitFrom_table_values_clause(c *From_table_values_clauseContext)

	// ExitFrom_table_values_row is called when exiting the from_table_values_row production.
	ExitFrom_table_values_row(c *From_table_values_rowContext)

	// ExitFrom_alias_clause is called when exiting the from_alias_clause production.
	ExitFrom_alias_clause(c *From_alias_clauseContext)

	// ExitTable_name is called when exiting the table_name production.
	ExitTable_name(c *Table_nameContext)

	// ExitWhere_clause is called when exiting the where_clause production.
	ExitWhere_clause(c *Where_clauseContext)

	// ExitGroup_by_clause is called when exiting the group_by_clause production.
	ExitGroup_by_clause(c *Group_by_clauseContext)

	// ExitHaving_clause is called when exiting the having_clause production.
	ExitHaving_clause(c *Having_clauseContext)

	// ExitQualify_clause is called when exiting the qualify_clause production.
	ExitQualify_clause(c *Qualify_clauseContext)

	// ExitOrder_by_clause is called when exiting the order_by_clause production.
	ExitOrder_by_clause(c *Order_by_clauseContext)

	// ExitSelect_options is called when exiting the select_options production.
	ExitSelect_options(c *Select_optionsContext)

	// ExitSelect_options_item is called when exiting the select_options_item production.
	ExitSelect_options_item(c *Select_options_itemContext)

	// ExitUpdate_stmt is called when exiting the update_stmt production.
	ExitUpdate_stmt(c *Update_stmtContext)

	// ExitUpdate_assignment is called when exiting the update_assignment production.
	ExitUpdate_assignment(c *Update_assignmentContext)

	// ExitUpdate_table is called when exiting the update_table production.
	ExitUpdate_table(c *Update_tableContext)

	// ExitUpdate_upsert is called when exiting the update_upsert production.
	ExitUpdate_upsert(c *Update_upsertContext)

	// ExitMerge_stmt is called when exiting the merge_stmt production.
	ExitMerge_stmt(c *Merge_stmtContext)

	// ExitMerge_table is called when exiting the merge_table production.
	ExitMerge_table(c *Merge_tableContext)

	// ExitMerge_condition is called when exiting the merge_condition production.
	ExitMerge_condition(c *Merge_conditionContext)

	// ExitMerge_action is called when exiting the merge_action production.
	ExitMerge_action(c *Merge_actionContext)

	// ExitDelete_stmt is called when exiting the delete_stmt production.
	ExitDelete_stmt(c *Delete_stmtContext)

	// ExitDelete_alias is called when exiting the delete_alias production.
	ExitDelete_alias(c *Delete_aliasContext)

	// ExitDescribe_stmt is called when exiting the describe_stmt production.
	ExitDescribe_stmt(c *Describe_stmtContext)

	// ExitBool_expr is called when exiting the bool_expr production.
	ExitBool_expr(c *Bool_exprContext)

	// ExitBool_expr_atom is called when exiting the bool_expr_atom production.
	ExitBool_expr_atom(c *Bool_expr_atomContext)

	// ExitBool_expr_unary is called when exiting the bool_expr_unary production.
	ExitBool_expr_unary(c *Bool_expr_unaryContext)

	// ExitBool_expr_single_in is called when exiting the bool_expr_single_in production.
	ExitBool_expr_single_in(c *Bool_expr_single_inContext)

	// ExitBool_expr_multi_in is called when exiting the bool_expr_multi_in production.
	ExitBool_expr_multi_in(c *Bool_expr_multi_inContext)

	// ExitBool_expr_binary is called when exiting the bool_expr_binary production.
	ExitBool_expr_binary(c *Bool_expr_binaryContext)

	// ExitBool_expr_logical_operator is called when exiting the bool_expr_logical_operator production.
	ExitBool_expr_logical_operator(c *Bool_expr_logical_operatorContext)

	// ExitBool_expr_binary_operator is called when exiting the bool_expr_binary_operator production.
	ExitBool_expr_binary_operator(c *Bool_expr_binary_operatorContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExpr_atom is called when exiting the expr_atom production.
	ExitExpr_atom(c *Expr_atomContext)

	// ExitExpr_interval is called when exiting the expr_interval production.
	ExitExpr_interval(c *Expr_intervalContext)

	// ExitInterval_item is called when exiting the interval_item production.
	ExitInterval_item(c *Interval_itemContext)

	// ExitExpr_concat is called when exiting the expr_concat production.
	ExitExpr_concat(c *Expr_concatContext)

	// ExitExpr_concat_item is called when exiting the expr_concat_item production.
	ExitExpr_concat_item(c *Expr_concat_itemContext)

	// ExitExpr_case is called when exiting the expr_case production.
	ExitExpr_case(c *Expr_caseContext)

	// ExitExpr_case_simple is called when exiting the expr_case_simple production.
	ExitExpr_case_simple(c *Expr_case_simpleContext)

	// ExitExpr_case_searched is called when exiting the expr_case_searched production.
	ExitExpr_case_searched(c *Expr_case_searchedContext)

	// ExitExpr_cursor_attribute is called when exiting the expr_cursor_attribute production.
	ExitExpr_cursor_attribute(c *Expr_cursor_attributeContext)

	// ExitExpr_agg_window_func is called when exiting the expr_agg_window_func production.
	ExitExpr_agg_window_func(c *Expr_agg_window_funcContext)

	// ExitExpr_func_all_distinct is called when exiting the expr_func_all_distinct production.
	ExitExpr_func_all_distinct(c *Expr_func_all_distinctContext)

	// ExitExpr_func_over_clause is called when exiting the expr_func_over_clause production.
	ExitExpr_func_over_clause(c *Expr_func_over_clauseContext)

	// ExitExpr_func_partition_by_clause is called when exiting the expr_func_partition_by_clause production.
	ExitExpr_func_partition_by_clause(c *Expr_func_partition_by_clauseContext)

	// ExitExpr_spec_func is called when exiting the expr_spec_func production.
	ExitExpr_spec_func(c *Expr_spec_funcContext)

	// ExitExpr_func is called when exiting the expr_func production.
	ExitExpr_func(c *Expr_funcContext)

	// ExitExpr_dot is called when exiting the expr_dot production.
	ExitExpr_dot(c *Expr_dotContext)

	// ExitExpr_dot_method_call is called when exiting the expr_dot_method_call production.
	ExitExpr_dot_method_call(c *Expr_dot_method_callContext)

	// ExitExpr_dot_property_access is called when exiting the expr_dot_property_access production.
	ExitExpr_dot_property_access(c *Expr_dot_property_accessContext)

	// ExitExpr_func_params is called when exiting the expr_func_params production.
	ExitExpr_func_params(c *Expr_func_paramsContext)

	// ExitFunc_param is called when exiting the func_param production.
	ExitFunc_param(c *Func_paramContext)

	// ExitExpr_select is called when exiting the expr_select production.
	ExitExpr_select(c *Expr_selectContext)

	// ExitExpr_file is called when exiting the expr_file production.
	ExitExpr_file(c *Expr_fileContext)

	// ExitHive is called when exiting the hive production.
	ExitHive(c *HiveContext)

	// ExitHive_item is called when exiting the hive_item production.
	ExitHive_item(c *Hive_itemContext)

	// ExitHost is called when exiting the host production.
	ExitHost(c *HostContext)

	// ExitHost_cmd is called when exiting the host_cmd production.
	ExitHost_cmd(c *Host_cmdContext)

	// ExitHost_stmt is called when exiting the host_stmt production.
	ExitHost_stmt(c *Host_stmtContext)

	// ExitFile_name is called when exiting the file_name production.
	ExitFile_name(c *File_nameContext)

	// ExitDate_literal is called when exiting the date_literal production.
	ExitDate_literal(c *Date_literalContext)

	// ExitTimestamp_literal is called when exiting the timestamp_literal production.
	ExitTimestamp_literal(c *Timestamp_literalContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitQident is called when exiting the qident production.
	ExitQident(c *QidentContext)

	// ExitSingle_quotedString is called when exiting the single_quotedString production.
	ExitSingle_quotedString(c *Single_quotedStringContext)

	// ExitDouble_quotedString is called when exiting the double_quotedString production.
	ExitDouble_quotedString(c *Double_quotedStringContext)

	// ExitInt_number is called when exiting the int_number production.
	ExitInt_number(c *Int_numberContext)

	// ExitDec_number is called when exiting the dec_number production.
	ExitDec_number(c *Dec_numberContext)

	// ExitBool_literal is called when exiting the bool_literal production.
	ExitBool_literal(c *Bool_literalContext)

	// ExitNull_const is called when exiting the null_const production.
	ExitNull_const(c *Null_constContext)

	// ExitNon_reserved_words is called when exiting the non_reserved_words production.
	ExitNon_reserved_words(c *Non_reserved_wordsContext)
}
