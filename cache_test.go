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
		{"Luxi Serif", "Regular"},
	}

	for _, tt := range testcases {
		_, err := fc.Get(tt.name, tt.style)
		if err != nil {
			t.Errorf("error: %s", err)
		}
	}
}

func TestFallback(t *testing.T) {
	testcases := []struct {
		name, style string
	}{
		{"Luxi Sans", "Luxi Sans"},
		{"Luxi Monooooooo", "Regular"},
		{"Luxi Seriqzxf", "Regular"},
	}

	expectedName, expectedStyle := "Luxi Sans", "Regular"
	fc.SetFallbackFont(expectedName, expectedStyle)

	for _, tt := range testcases {
		font, err := fc.Get(tt.name, tt.style)
		if err != nil {
			t.Errorf("error: %s", err)
		} else {
			if got := font.Name(truetype.NameIDFontFamily); got != expectedName {
				t.Errorf("wanted familyName of %q, got %q", expectedName, got)
			}
			if got := font.Name(truetype.NameIDFontSubfamily); got != expectedStyle {
				t.Errorf("wanted subfamilyName of %q, got %q", expectedStyle, got)
			}
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
