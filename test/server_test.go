package main_test


import (
	"database/sql"
	"testing"
	"github.com/msalbrain/authSphere/internals/server"
)




func TestDbSetup(t *testing.T) {
	var Db *sql.DB

	t.Parallel()

	t.Run("Initalize db", func(t *testing.T) {

		_, db, err  := server.InitDatabase("sqlite3", "testdatabase.db")

		if err != nil {
			t.Errorf("had problems connecting to db %s", err)
		}

		Db = db
		Db.Ping()
	})


	t.Run("migrate db", func(t *testing.T) {

		err  := server.RunMigrateScripts(Db)

		if err != nil {
			t.Errorf("Had problems migrating db %s \n", err)
		}
	})

}
