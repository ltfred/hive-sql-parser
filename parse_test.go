package main

import "testing"

func TestParseCreateTableStmt(t *testing.T) {
	parser := CreateTableStmt{}
	stmt := parser.Parse(`create table test1
	(
	   c1 int COMMENT 'c1c1',
	   c2 string not null,
	   c3 boolean,
	   c4 double,
	   c5 date COMMENT 'date',
	   c6 timestamp,
	   c7 varchar(255)
	) comment '3434'`)
	tableStmt := stmt.(*CreateTableStmt)
	t.Log(tableStmt)
}
