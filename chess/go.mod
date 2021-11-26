module chess

go 1.16

require (
	github.com/google/uuid v1.3.0
	github.com/johnjones4/GoChess/chess v0.0.0-00010101000000-000000000000
	github.com/mattn/go-sqlite3 v1.14.9
	golang.org/x/net v0.0.0-20211118161319-6a13c67c3ce4
)

replace github.com/johnjones4/GoChess/chess => ./
