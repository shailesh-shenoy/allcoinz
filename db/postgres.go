package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"io/fs"

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

//go:embed migrations/*.sql
var migrationFS embed.FS

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

func (ds *DataStore) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return ds.db.BeginTx(ctx, opts)

}

func (ds *DataStore) migrate() (err error) {
	// Ensure the 'migrations' table exists so we don't duplicate migrations.
	if _, err := ds.db.Exec(`CREATE TABLE IF NOT EXISTS migrations (name TEXT PRIMARY KEY);`); err != nil {
		return fmt.Errorf("cannot create migrations table: %w", err)
	}

	// Read migration files from our embedded file system.
	names, err := fs.Glob(migrationFS, "migrations/*.sql")
	if err != nil {
		return err
	}

	// Loop over migration files and attempt to migrate each one
	for _, name := range names {
		if err := ds.migrateFile(name); err != nil {
			return fmt.Errorf("migration failed for file: %q with error: %w", name, err)
		}
	}

	return nil
}

func (ds *DataStore) migrateFile(name string) (err error) {
	tx, err := ds.db.Begin()
	if err != nil {
		return err
	}
	// Rollback the transaction before returning.
	// If there is an error, all SQL in migration file is rolled back.
	// If there is no error, the transaction was already committed
	// and the rollback is a NO-OP as there is nothing to rollback.
	// * This is the reason there are strange empty rollbacks when using ORMs.
	defer tx.Rollback()

	// Ensure that the migration has not already been run.
	// The below query returns n > 0 if migration was already run and inserted in the DB.
	var n int
	if err = tx.QueryRow(`SELECT COUNT(*) FROM migrations WHERE name = $1;`, name).Scan(&n); err != nil {
		return err
	}
	if n != 0 {
		// Migration was already run for this file, skip
		return nil
	}

	// Read and execute migration file
	if buf, err := fs.ReadFile(migrationFS, name); err != nil {
		return err
	} else if _, err = tx.Exec(string(buf)); err != nil {
		return err
	}

	if _, err = tx.Exec(`INSERT INTO migrations (name) VALUES($1)`, name); err != nil {
		return err
	}
	// Now that the file is migrated, commit the file
	// If there is an error in the commit, it gets returned
	return tx.Commit()
}
