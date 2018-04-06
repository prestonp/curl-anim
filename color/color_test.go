package color

import (
	"testing"
)

func TestColor(t *testing.T) {
	p := New()

	testcases := []struct {
		in  string
		out string
	}{
		{"hello", "\033[2J\033[H\033[0;31mhello\033[0m\n"},
		{"hello", "\033[2J\033[H\033[0;32mhello\033[0m\n"},
		{"hello", "\033[2J\033[H\033[0;33mhello\033[0m\n"},
		{"hello", "\033[2J\033[H\033[0;34mhello\033[0m\n"},
		{"hello", "\033[2J\033[H\033[0;35mhello\033[0m\n"},
		{"hello", "\033[2J\033[H\033[0;36mhello\033[0m\n"},
		{"hello", "\033[2J\033[H\033[0;37mhello\033[0m\n"},
		{"hello", "\033[2J\033[H\033[0;31mhello\033[0m\n"},
		{"hello", "\033[2J\033[H\033[0;32mhello\033[0m\n"},
	}

	for _, test := range testcases {
		result := p.Print(test.in)
		if test.out != result {
			t.Errorf("expected: %s, got :%s\n", test.out, result)
		}
	}
}
