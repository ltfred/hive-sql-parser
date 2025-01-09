## Usage

```go
import (
    "github.com/ltfred/hive-sql-parser"
)
```

### ParseCreateTableStatement

```go
parser := hive_sql_parser.NewCreatTableStmtParser()
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
createStmt := stmt.(*hive_sql_parser.CreateTableStmt)
fmt.Println(createStmt)
// {test1 3434 [{c1 int false 0 c1c1} {c2 string true 0 } {c3 boolean false 0 } {c4 double false 0 } {c5 date false 0 date} {c6 timestamp false 0 } {c7 varchar false 255 }] []}

```
See [parse_test.g](https://github.com/ltfred/hive-sql-parser/blob/master/parse_test.go)o for more examples