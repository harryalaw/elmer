package cmd

import (
	"testing"

	"github.com/harryalaw/elmer/internal/db"
)

func TestAddCommandExisting(t *testing.T) {
	dirB := db.NewDir("/foo/bar")
	database := db.FromDirs([]db.Dir{*dirB})

	updatedDb, err := add("/foo/bar", database)

	if err != nil {
		t.Fatalf("Err not nil, got: %+v\n", err)
	}

	bar := updatedDb.Find("bar")
	if bar.Rank() != 1 {
		t.Fatalf("Directory score not increased.\n Got: %+v\n Expected score to be 1", bar)
	}
}

func TestAddCommandNewDir(t *testing.T) {
	dirB := db.NewDir("/foo/bar")
	database := db.FromDirs([]db.Dir{*dirB})

	updatedDb, err := add("/foo/yargh", database)

	if err != nil {
		t.Fatalf("Err not nil, got: %+v\n", err)
	}

	bar := updatedDb.Find("bar")
	if bar.Rank() != 0 {
		t.Fatalf("Directory score increased incorrectly.\n Got: %+v\n Expected score to be 0", bar)
	}

	yargh := updatedDb.Find("yargh")
	if yargh.Rank() != 1 {
		t.Fatalf("Directory score not increased.\n Got: %+v\n Expected score to be 1", bar)
	}
}
