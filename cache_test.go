package fontcache_test

import (
	"testing"

	"github.com/stephenwithav/fontcache"
)

var fc fontcache.Cache

func init() {
	fc = fontcache.New()
	fc.Init("./fonts/")
}

func TestGet(t *testing.T) {
	testcases := []struct {
		name, style string
	}{
		{"Luxi Sans", "Regular"},
		{"Luxi Mono", "Regular"},
	}

	for _, tt := range testcases {
		_, err := fc.Get(tt.name, tt.style)
		if err != nil {
			t.Errorf("error: %s", err)
		}
	}
}
