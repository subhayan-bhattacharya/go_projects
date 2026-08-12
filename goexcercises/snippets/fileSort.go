package snippets

import (
	"cmp"
	"slices"
	"strings"
)

func SortFiles(files []string) {
	slices.SortFunc(files, func(a, b string) int {
		fileA := strings.ToLower(a)
		fileB := strings.ToLower(b)
		return cmp.Compare(fileA, fileB)
	})
}
