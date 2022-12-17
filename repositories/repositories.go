package repositories

import (
	"database/sql"
	"github.com/iscfgibarra/applabs-data/abstractions"
	"github.com/iscfgibarra/applabs-data/events"
)

type GenericRepository struct {
	db        *sql.DB
	tableName string
	EventBus  abstractions.EventBus
}

func NewGenericRepository(db *sql.DB, tableName string, eventBus abstractions.EventBus) *GenericRepository {
	return &GenericRepository{
		db,
		tableName,
		eventBus,
	}
}

func (r *GenericRepository) MigrateWithCmd(migrateCmd string) error {
	_, err := r.db.Exec(migrateCmd)

	if err != nil {
		r.EventBus.Push("Error on migrate table "+r.tableName, err, events.ERROR)
	}

	return err
}

func (r *GenericRepository) CreateWithCmd(commandCreate string, args ...interface{}) error {
	tx, err := r.db.Begin()
	if err != nil {
		r.EventBus.Push("Error on init transaction Create ", err, events.ERROR)
		return err
	}

	stmt, err := tx.Prepare(commandCreate)
	if err != nil {
		r.EventBus.Push("Error on prepare query getCreate "+r.tableName, err, events.ERROR)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)

	if err != nil {
		r.EventBus.Push("Error on execute getCreate "+r.tableName, err, events.ERROR)
		tx.Rollback()
		return err
	}

	tx.Commit()
	r.EventBus.Push("Success create the table "+r.tableName, args, events.INFO)
	return nil
}

func (r *GenericRepository) ByIdWithCmd(id string, commandGetByid string) (*sql.Row, error) {
	stmt, err := r.db.Prepare(commandGetByid)
	if err != nil {
		r.EventBus.Push("Error on execute query getById for "+r.tableName, err, events.ERROR)
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRow(id), nil
}

func (r *GenericRepository) PageByIdWithCmd(id string, page int, size int, commandGetPageById string) (*sql.Rows, error) {
	stmt, err := r.db.Prepare(commandGetPageById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	offset := (page - 1) * size
	return stmt.Query(id, size, offset)
}

func (r *GenericRepository) PullEvents() events.Events {
	return r.EventBus.Pull()
}
