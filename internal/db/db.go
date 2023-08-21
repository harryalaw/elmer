package db

import (
	"bytes"
	"encoding/gob"
	"strings"
)

type Db struct {
	dirs []Dir
}

func FromDirs(dirs []Dir) *Db {
	return &Db{dirs: dirs}
}

func (db *Db) AddDir(dir *Dir) *Db {
	db.dirs = append(db.dirs, *dir)
	return db
}

func (db *Db) Equals(o *Db) bool {
	if len(db.dirs) != len(o.dirs) {
		return false
	}

	for i, dir := range db.dirs {
		oDir := o.dirs[i]
		if oDir.path != dir.path || oDir.score != dir.score || oDir.lastAccessed != dir.lastAccessed {
			return false
		}
	}
	return true
}

func (db *Db) Use(dir *Dir) {
	updatedDirs := make([]Dir, len(db.dirs))
	for i, dbDir := range db.dirs {
		if dbDir.Equals(dir) {
			dbDir.Use()
		}
		updatedDirs[i] = dbDir
	}
	db.dirs = updatedDirs
}

func (db *Db) Find(path string) *Dir {
	filtered := make([]Dir, 0)

	for _, dir := range db.dirs {
		if end := dir.PathEnd(); strings.Contains(end, path) {
			filtered = append(filtered, dir)
		}
	}

	if len(filtered) == 0 {
		return nil
	}

	// sort descending filtered by score
	quicksort(&filtered)

	return &filtered[0]
}

// inmemory quicksort
func quicksort(dirs *[]Dir) {
	qs(dirs, 0, len(*dirs)-1)
}

func qs(dirs *[]Dir, lo, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(dirs, lo, hi)
	qs(dirs, lo, pivotIdx-1)
	qs(dirs, pivotIdx+1, hi)
}

func partition(dirs *[]Dir, lo, hi int) int {
	pivot := (*dirs)[hi]

	idx := lo - 1

	for i := lo; i < hi; i++ {
		if (*dirs)[i].GreaterThan(&pivot) {
			idx++
			tmp := (*dirs)[i]
			(*dirs)[i] = (*dirs)[idx]
			(*dirs)[idx] = tmp
		}
	}

	idx++
	(*dirs)[hi] = (*dirs)[idx]
	(*dirs)[idx] = pivot

	return idx
}

func (db *Db) GobEncode() ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	if err := encoder.Encode(len(db.dirs)); err != nil {
		return nil, err
	}

	for _, dir := range db.dirs {
		if err := encoder.Encode(dir.path); err != nil {
			return nil, err
		}
		if err := encoder.Encode(dir.score); err != nil {
			return nil, err
		}
		if err := encoder.Encode(dir.lastAccessed); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func (db *Db) GobDecode(data []byte) error {
	buf := bytes.NewReader(data)
	decoder := gob.NewDecoder(buf)

	// Decode Foo's bar slice length
	var dirLength int
	if err := decoder.Decode(&dirLength); err != nil {
		return err
	}

	// Decode each Baz element and populate Foo's bar slice
	for i := 0; i < dirLength; i++ {
		var dir Dir
		if err := decoder.Decode(&dir.path); err != nil {
			return err
		}
		if err := decoder.Decode(&dir.score); err != nil {
			return err
		}
		if err := decoder.Decode(&dir.lastAccessed); err != nil {
			return err
		}
		db.dirs = append(db.dirs, dir)
	}

	return nil
}
