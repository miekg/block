package block

import (
	"strings"
	"testing"
)

var blocked = `
127.0.0.1	005.free-counter.co.uk
127.0.0.1	006.free-adult-counters.x-xtra.com
127.0.0.1	006.free-counter.co.uk
127.0.0.1	007.free-counter.co.uk
127.0.0.1	007.go2cloud.org
127.0.0.1	localhost
`

func TestBlocked(t *testing.T) {
	r := strings.NewReader(blocked)
	l := make(map[string]struct{})
	l, err := ListRead(r, l)
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name    string
		blocked bool
	}{
		{"example.org.", false},
		{"localhost.", false},

		{"005.free-counter.co.uk.", true},
		{"www.005.free-counter.co.uk.", true},
	}
	t.Logf("%v", l)

	for _, test := range tests {
		got := Blocked(test.name, l)
		if got != test.blocked {
			t.Errorf("Expected %s to be blocked", test.name)
		}
	}
}
