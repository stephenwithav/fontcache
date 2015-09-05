package fontcache_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/golang/freetype/truetype"
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

func Example() {
	fontPath := "./fonts"

	fc := fontcache.New()
	fc.Init(fontPath)

	// Retrieve font by name for use in a program.
	font, err := fc.Get("Luxi Sans", "Regular")
	if err != nil {
		log.Fatalf("%s", err)
	}

	func(f *truetype.Font) {
		fmt.Printf("Do something with %s", f.Name(truetype.NameIDFontFamily))
	}(font)
	// Output: Do something with Luxi Sans
}
