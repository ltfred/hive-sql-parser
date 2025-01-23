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

	stmt = parser.Parse(`create table test1
	(
	   c1 int COMMENT 'c1c1',
	   c2 string not null,
	   c3 boolean,
	   c4 double,
	   c5 date COMMENT 'date',
	   c6 timestamp,
	   c7 varchar(255),
	   c8 decimal(14, 2),
	   c9 decimal(10, 0)
	) comment '3434'`)
	tableStmt = stmt.(*CreateTableStmt)

	s.Equal("decimal", tableStmt.TableColumns[7].Type)
	s.Equal(14, tableStmt.TableColumns[7].Precision)
	s.Equal(2, tableStmt.TableColumns[7].Scale.Value)
	s.Equal(10, tableStmt.TableColumns[8].Precision)
	s.Equal(0, tableStmt.TableColumns[8].Scale.Value)
	s.Equal(true, tableStmt.TableColumns[8].Scale.Valid)

	stmt = parser.Parse(`create table dwd_air_iep_forecast_station_hour_detail (
    oid SMALLSERIAL NULL,
code SERIAL NULL,
name BIGSERIAL NULL,
aqi SMALLINT NULL, 
pm25 INTEGER NULL,
pm10 BIGINT NULL, 
so2 CHAR(10) NULL,
no2 VARCHAR(10) NULL,
co CHARACTER(10) NULL,
o3 CHARACTER VARYING(10) NULL,
aqi_main TEXT NULL,
temperature REAL NULL,
humidity DOUBLE PRECISION NULL,
wind_speed BOOLEAN NULL,
wind_direction TIME NULL,
pressure TIMETZ NULL,
create_at TIMESTAMP NULL,
tim TIMESTAMPTZ NULL,
cid TIMESTAMP WITHOUT TIME ZONE(6) NULL,
cid_name TIMESTAMP WITH TIME ZONE(6) NULL,
type DATE NULL,
forecast_time GEOGRAPHY NULL,
aqi_max GEOMETRY NULL, 
aqi_min1 JSON NULL,
aqi_min2 NUMERIC NULL,
aqi_min3 DECIMAL NULL,
)`)

	tableStmt = stmt.(*CreateTableStmt)
	s.T().Log(tableStmt)
}

func (s *ParseTestSuit) TestParseDropTableStmt() {
	parser := NewDropTableStmtParser()
	stmt := parser.Parse(`drop table test1`)
	dropTableStmt := stmt.(*DropTableStmt)
	s.Equal("test1", dropTableStmt.TableName)
}
