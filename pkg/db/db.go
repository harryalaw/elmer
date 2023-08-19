package db

import "strings"

type Db struct {
	dirs []Dir
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
