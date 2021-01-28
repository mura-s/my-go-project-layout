package mylib

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		desc  string
		s     string
		times int
		out   string
	}{
		{"empty", "", 2, ""},
		{"2 x 2", "ab", 2, "abab"},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			assert.Equal(t, c.out, Repeat(c.s, c.times))
		})
	}
}

func TestUnique(t *testing.T) {
	cases := []struct {
		desc string
		in   []string
		out  []string
	}{
		{"no dup", []string{"c", "a", "b"}, []string{"a", "b", "c"}},
		{"dup", []string{"c", "a", "b", "c"}, []string{"a", "b", "c"}},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			actual := Unique(c.in)
			sort.Strings(actual)
			assert.Equal(t, c.out, actual)
		})
	}
}
