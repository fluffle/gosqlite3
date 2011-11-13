package sqlite3

import "bytes"
import "encoding/gob"
import "testing"

func TestGeneral(t *testing.T) {
	Initialize()
	defer Shutdown()
	t.Logf("Sqlite3 Version: %v\n", LibVersion())

	filename := ":memory:"
	db, e := Open(filename)
	fatalOnError(t, e, "opening %v", filename)

	defer db.Close()
	t.Logf("Database opened: %v [flags: %v]", db.Filename, int(db.Flags))
	t.Logf("Returning status: %v", e)
}

func TestBlob(t *testing.T) {
	Session("test.db", func(db *Database) {
		BAR.Drop(db)
		BAR.Create(db)

		buffer := new(bytes.Buffer)
		encoder := gob.NewEncoder(buffer)
		fatalOnError(t, encoder.Encode(TwoItems{"holy", "moly guacomole"}), "Encoding failed: buffer = %v", buffer)
		t.Logf("Encoded data: %v", buffer.Bytes())

		db.runQuery(t, "INSERT INTO bar values (?, ?)", 1, TwoItems{"holy moly", "guacomole"})
		db.stepThroughRows(t, BAR)
	})
}
