package tinyindexdb

import (
	"github.com/gosimple/slug"
)

func sanitiseString(input string) string {
	return slug.Make(input)
}
