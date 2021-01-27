package mylib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		desc string
		a    int
		b    int
		out  int
	}{
		{"pos + pos", 1, 1, 2},
		{"neg + neg", -1, -1, -2},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			assert.Equal(t, c.out, Add(c.a, c.b))
		})
	}
}
