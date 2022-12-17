package repositories

const (
	expenseTypeTableDefSqlite = `CREATE TABLE IF NOT EXISTS main.ExpenseType(
		Code TEXT PRIMARY KEY,
		Description TEXT NOT NULL,
		CreatedAt TEXT NOT NULL
	);`
	expenseTypeCreateInsertSqlite = `INSERT INTO main.ExpenseType(Code, Description, CreatedAt)
		VALUES (?, ?, ?);`
	selectExpenseTypeSqlite = `SELECT Code, Description, CreatedAt
		FROM main.ExpenseType `
	getExpenseTypeByCodeSqlite = selectExpenseTypeSqlite + `WHERE Code = ?`
	getExpenseTypePageSqlite   = selectExpenseTypeSqlite + `LIMIT ? OFFSET ?`
)

type SqliteCommands struct {
}

func NewSqliteCommands() *SqliteCommands {
	return &SqliteCommands{}
}

func (c *SqliteCommands) GetTableDef() string {
	return expenseTypeTableDefSqlite
}

func (c *SqliteCommands) GetExpenseTypeByCode() string {
	return getExpenseTypeByCodeSqlite
}

func (c *SqliteCommands) GetCreateExpenseType() string {
	return expenseTypeCreateInsertSqlite
}

func (c *SqliteCommands) GetExpenseTypePage() string {
	return getExpenseTypePageSqlite
}
