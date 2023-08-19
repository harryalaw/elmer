package db

import (
	"strings"
	"time"
)

type Dir struct {
	path  string
	score int
	// time in seconds since Unix epoch that it was last accessed
	lastAccessed int64
}

func NewDir(path string) *Dir {
	return &Dir{
		path:         path,
		score:        1,
		lastAccessed: time.Now().Unix(),
	}
}

func (d *Dir) Use() string {
	d.score += 1
	d.lastAccessed = time.Now().Unix()
	return d.path
}

func (d *Dir) PathEnd() string {
	// assuming that all paths will be unix like and separated with
	// "/"
	// like with "/usr/bin/temp/whatever"
	parts := strings.Split(d.path, "/")

	return parts[len(parts)-1]
}

func (d *Dir) Rank() int {
	return d.score
}

func (d *Dir) GreaterThan(o *Dir) bool {
	return d.Rank() > o.Rank() ||
		(d.Rank() == o.Rank() && d.lastAccessed > o.lastAccessed)

}
