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
	{0xffffffffffffffff, ""},
	{0xecbe3bd72fbc950d, "hello world 1"},
	{0xecbc39d72bb8950d, "hello world 2"},
	{0xecbe39d72bb8950d, "hello world 3"},
	{0x38370f452a65abfb, "the quick brown fox jumps over the lazy dog"},
	{0x7b3e2b252867bf9f, "The quick brown fox jumps over the lazy dog"},
	{0xfb3eab252867bfbf, "The quick brown fox jumps over the lazy dog!"},
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
