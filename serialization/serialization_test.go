package serial

import (
	"os"
	"testing"

	"github.com/harryalaw/elmer/db"
)

func TestSerialization(t *testing.T) {
	filepath := "./elmer.db"

	foo := db.NewDir("/tmp/foo")
	bar := db.NewDir("/tmp/foo/bar")
	baz := db.NewDir("/tmp/baz")

	db := db.FromDirs([]db.Dir{*foo, *bar, *baz})

	err := WriteDb(db, filepath)

	if err != nil {
		t.Fatalf("Expected no writing error, got: %+v", err)
	}

	outDb, err := ImportDb(filepath)

	if err != nil {
		t.Fatalf("Expected no importing error, got: %+v", err)
	}

	if !outDb.Equals(db) {
		t.Fatalf("Retrieved DB didn't match original DB\nGot:%+v\nExpected:%+v\n", outDb, db)
	}

	os.Remove(filepath)
}
