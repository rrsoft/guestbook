package data

type CommandType int

const (
	Text            CommandType = 1
	StoredProcedure             = 4
	TableDirect                 = 512
)
