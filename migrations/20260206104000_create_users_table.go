package migrations

import "github.com/sllt/kite/pkg/kite/migration"

const createUsersTable = `CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id TEXT UNIQUE NOT NULL,
	nickname TEXT NOT NULL DEFAULT '',
	password TEXT NOT NULL,
	email TEXT NOT NULL,
	created_at DATETIME,
	updated_at DATETIME
);`

func createUsersTableMigration() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createUsersTable)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
