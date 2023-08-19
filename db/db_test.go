package db

import (
	"testing"
)

func TestFindDir(t *testing.T) {
	foo := Dir{path: "/test/foo", score: 1, lastAccessed: int64(1324)}
	bar := Dir{path: "/test/foo/bar", score: 10, lastAccessed: int64(4321)}
	betterFoo := Dir{path: "/test/better/foo", score: 100, lastAccessed: int64(2222)}
	baz := Dir{path: "/test/foo/baz", score: 10, lastAccessed: int64(1234)}

	testCases := []struct {
		dirs        []Dir
		expectedDir *Dir
		pathName    string
	}{
		{
			dirs:        []Dir{foo},
			expectedDir: &foo,
			pathName:    "foo",
		},
		{
			dirs:        []Dir{bar, foo},
			expectedDir: &foo,
			pathName:    "foo",
		},
		{
			dirs:        []Dir{foo, betterFoo},
			expectedDir: &betterFoo,
			pathName:    "foo",
		},
		{
			dirs:        []Dir{bar, baz},
			expectedDir: &bar,
			pathName:    "ba",
		},
		{
			dirs:        []Dir{},
			expectedDir: nil,
			pathName:    "anything",
		},
		{
			dirs:        []Dir{foo, bar, baz},
			expectedDir: nil,
			pathName:    "missing",
		},
	}

	for i, tt := range testCases {
		db := Db{dirs: tt.dirs}

		output := db.Find(tt.pathName)

		if !dirEquals(output, tt.expectedDir) {
			t.Errorf("Test case %d failed\nOutput not what was expected.\nGot: %+v\nExpected: %+v\n", i, output, tt.expectedDir)
		}
	}
}

func dirEquals(d *Dir, o *Dir) bool {
	if d == nil && o == nil {
		return true
	}
	return d != nil &&
		o != nil &&
		d.lastAccessed == o.lastAccessed &&
		d.score == o.score &&
		d.path == o.path
}
