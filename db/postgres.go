package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DataStore struct {
	db         *sql.DB
	dataSource string
}

// Returns a new instance of DataStore associated with the given datasource string.
// * The actual DB instance is not initialized here
func NewDataStore(dsrc string) *DataStore {
	return &DataStore{
		dataSource: dsrc,
	}
}

// Opens a database connection, returns error if unsuccessful
func (ds *DataStore) Open() (err error) {

	if ds.db, err = sql.Open("postgres", ds.dataSource); err != nil {
		return fmt.Errorf("datasource required: %w", err)
	}

	if err = ds.db.Ping(); err != nil {
		return fmt.Errorf("unable to connect to postgres database: %s", err.Error())
	}

	if err = ds.migrate(); err != nil {
		return fmt.Errorf("migration error: %w", err)
	}

	return nil
}

func (ds *DataStore) migrate() (err error) {
	return nil
}
