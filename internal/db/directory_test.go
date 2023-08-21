package db

import (
	"testing"
)

func TestUse(t *testing.T) {
	path := "/test/dir/foo"
	score := 10
	lastAccessed := int64(1692435000)

	d := Dir{
		path:         path,
		score:        score,
		lastAccessed: lastAccessed,
	}

	output := d.Use()

	if output != path {
		t.Fatalf("Output path was: %s, expected: %s", output, path)
	}

	if d.path != path {
		t.Fatalf("Path has changed, was: %s, expected: %s", d.path, path)
	}
	if d.score != score+1 {
		t.Fatalf("score was: %d, expected: %d", d.score, score+1)
	}

	if d.lastAccessed <= lastAccessed {
		t.Fatalf("lastAccessed was: %d, expected: %d", d.lastAccessed, lastAccessed)
	}
}
