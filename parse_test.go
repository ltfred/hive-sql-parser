package hive_sql_parser

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ParseTestSuit struct {
	suite.Suite
}

func TestParse(t *testing.T) {
	suite.Run(t, new(ParseTestSuit))
}

func (s *ParseTestSuit) TestParseCreateTableStmt() {
	// no partition
	parser := NewCreatTableStmtParser()
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
	s.Equal("test1", tableStmt.TableName)
	s.Equal("3434", tableStmt.TableComment)
	s.Equal(7, len(tableStmt.TableColumns))
	s.Equal(0, len(tableStmt.PartitionColumns))
	s.Equal("c1c1", tableStmt.TableColumns[0].Comment)
	s.Equal(true, tableStmt.TableColumns[1].NotNull)
	s.Equal(255, tableStmt.TableColumns[6].Length)

	// with partition
	stmt = parser.Parse(`create table test1
	(
	   c1 int COMMENT 'c1c1',
	   c2 string not null,
	   c3 boolean,
	   c4 double,
	   c5 date COMMENT 'date',
	   c6 timestamp,
	   c7 varchar(255)
	) comment '3434' partitioned by (c8 string comment 'c8 partition')`)
	tableStmt = stmt.(*CreateTableStmt)
	s.Equal("test1", tableStmt.TableName)
	s.Equal("3434", tableStmt.TableComment)
	s.Equal(7, len(tableStmt.TableColumns))
	s.Equal(1, len(tableStmt.PartitionColumns))
	s.Equal("c8 partition", tableStmt.PartitionColumns[0].Comment)
	s.Equal("string", tableStmt.PartitionColumns[0].Type)

	stmt = parser.Parse("create table test1\n\t(\n\t   c1 int COMMENT 'c1c1'," +
		"\n\t   `type` string not null\n\t) comment '3434'")
	tableStmt = stmt.(*CreateTableStmt)
	s.Equal("test1", tableStmt.TableName)
	s.Equal("3434", tableStmt.TableComment)
	s.Equal(2, len(tableStmt.TableColumns))
	s.Equal("type", tableStmt.TableColumns[1].Name)
}

func (s *ParseTestSuit) TestParseDropTableStmt() {
	parser := NewDropTableStmtParser()
	stmt := parser.Parse(`drop table test1`)
	dropTableStmt := stmt.(*DropTableStmt)
	s.Equal("test1", dropTableStmt.TableName)
}
