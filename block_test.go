package block

import (
	"strings"
	"testing"
)

func TestBlocked(t *testing.T) {
	var list = `
127.0.0.1	005.free-counter.co.uk
127.0.0.1	006.free-adult-counters.x-xtra.com
127.0.0.1	006.free-counter.co.uk
127.0.0.1	007.free-counter.co.uk
127.0.0.1	007.go2cloud.org
127.0.0.1	localhost
008.free-counter.co.uk
com
`

	b := new(Block)

	r := strings.NewReader(list)
	l := make(map[string]struct{})
	listRead(r, l)
	b.list = l

	tests := []struct {
		name    string
		blocked bool
	}{
		{"example.org.", false},
		{"localhost.", false},
		{"com.", false},

		{"005.free-counter.co.uk.", true},
		{"www.005.free-counter.co.uk.", true},
		{"008.free-counter.co.uk.", true},
		{"www.008.free-counter.co.uk.", true},
	}

	for _, test := range tests {
		got := b.blocked(test.name)
		if got != test.blocked {
			t.Errorf("Expected %s to be blocked", test.name)
		}
	}
}
