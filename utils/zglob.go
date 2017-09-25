package utils

import (
	"github.com/k0kubun/pp"
	"github.com/mattn/go-zglob"
)

/*
	refs:
	- https://github.com/mattn/go-zglob
*/

func Find(rule string) {
	// rule := "./foo/b*/**/z*.txt"
	matches, err := zglob.Glob(rule)
	pp.Println(matches)
}
