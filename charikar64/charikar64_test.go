package charikar64_test

import (
	"hash"
	"testing"
	"github.com/chmduquesne/simhash/charikar64"
)

var _ = hash.Hash64(charikar64.New())

var golden = []struct {
	out       uint64
	in        string
}{
	{0x185d611da9009af5, ""},
	{0xfcbf3bdfafbc9f0d, "hello world 1"},
	{0xecbd39d7abb8970d, "hello world 2"},
	{0xfcbf39dfabb8978d, "hello world 3"},
	{0x383f0f4d2a65abfb, "the quick brown fox jumps over the lazy dog"},
	{0x7b3e2b252867bfff, "The quick brown fox jumps over the lazy dog"},
	{0x7b3e29252865bfbf, "The quick brown fox jumps over the lazy dog!"},
}

func TestGolden(t *testing.T) {
	for _, g := range(golden) {
		c := charikar64.New()
		c.Write([]byte(g.in))
		if got := c.Sum64(); got != g.out {
			t.Errorf("%q got 0x%x, want 0x%x", g.in, got, g.out)
		}
	}
}
