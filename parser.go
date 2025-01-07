package main

type Statement interface {
	iStatement()
}

type Parser interface {
	Parse(sql string) Statement
}


