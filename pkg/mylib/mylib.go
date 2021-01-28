package mylib

import (
	"strings"

	"github.com/mura-s/my-go-project-layout/pkg/internal/set"
)

const DefaultCount = 1

func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

func Unique(ss []string) []string {
	st := set.NewStringSet()
	for _, s := range ss {
		st[s] = struct{}{}
	}

	result := make([]string, 0, len(st))
	for k := range st {
		result = append(result, k)
	}
	return result
}
