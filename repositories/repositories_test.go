package repositories

import (
	"github.com/iscfgibarra/applabs-data/connector"
	"github.com/iscfgibarra/applabs-data/drivers"
	"github.com/iscfgibarra/applabs-data/events"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// ExpenseType store expense type
type ExpenseType struct {
	Code        string    `json:"code"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

func TestGenericRepository(t *testing.T) {
	connector.InitSqlConnector(":memory:", drivers.SQLITE)

	sqlCmds := NewSqliteCommands()
	evtBus := events.NewDomainEventBus()
	repo := NewGenericRepository(connector.SqlCnn.Db, "main.ExpenseType", evtBus)

	err := repo.MigrateWithCmd(sqlCmds.GetTableDef())
	assert.NoError(t, err, "was ocurred a error when create expenses type table")

	expType := &ExpenseType{
		Code:        "CODE001",
		Description: "NEWEXPENSE1",
		CreatedAt:   time.Now(),
	}

	expType2 := &ExpenseType{
		Code:        "CODE002",
		Description: "NEWEXPENSE2",
		CreatedAt:   time.Now(),
	}

	expType3 := &ExpenseType{
		Code:        "CODE003",
		Description: "NEWEXPENSE3",
		CreatedAt:   time.Now(),
	}

	expType4 := &ExpenseType{
		Code:        "CODE004",
		Description: "NEWEXPENSE4",
		CreatedAt:   time.Now(),
	}

	err = repo.CreateWithCmd(sqlCmds.GetCreateExpenseType(),
		expType.Code,
		expType.Description,
		expType.CreatedAt)

	assert.NoError(t, err, "was ocurred a error when insert new expense type")

	err = repo.CreateWithCmd(sqlCmds.GetCreateExpenseType(),
		expType2.Code,
		expType2.Description,
		expType2.CreatedAt)

	assert.NoError(t, err, "was ocurred a error when insert new expense type")

	err = repo.CreateWithCmd(sqlCmds.GetCreateExpenseType(),
		expType3.Code,
		expType3.Description,
		expType3.CreatedAt)

	assert.NoError(t, err, "was ocurred a error when insert new expense type")

	err = repo.CreateWithCmd(sqlCmds.GetCreateExpenseType(),
		expType4.Code,
		expType4.Description,
		expType4.CreatedAt)

	assert.NoError(t, err, "was ocurred a error when insert new expense type")

	expTypeRow, err1 := repo.ByIdWithCmd("CODE001", sqlCmds.GetExpenseTypeByCode())

	assert.NoError(t, err1, "was ocurred a error when get CODE001 expense type")

	var createdAt string
	newExpType := &ExpenseType{}
	errScan := expTypeRow.Scan(&newExpType.Code, &newExpType.Description, &createdAt)

	assert.NoError(t, errScan, "was ocurred a error when scan new expense type")

	assert.Equal(t, "CODE001", newExpType.Code, "code not equal")
	assert.Equal(t, "NEWEXPENSE1", newExpType.Description, "Description not equal")

}
