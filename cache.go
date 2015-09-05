package fontcache

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

type Cache map[string]Styles
type Styles map[string]*truetype.Font

func New() Cache {
	return make(Cache)
}

func (c Cache) Init(path string) {
	filepath.Walk(path, c.loadFont)
}

func (c Cache) Get(name, style string) (*truetype.Font, error) {
	if sm, ok := c[strings.ToLower(name)]; ok {
		if f, ok := sm[strings.ToLower(style)]; ok {
			return f, nil
		}
		return nil, errors.New(fmt.Sprintf("style %q not supported in font %q", style, name))
	}

	return nil, errors.New(fmt.Sprintf("font %q not found in cache", name))

}

func (c Cache) loadFont(path string, info os.FileInfo, err error) error {
	// process ttf files only
	if strings.ToLower(path[len(path)-4:]) != ".ttf" {
		return nil
	}

	ttfBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	font, err := freetype.ParseFont(ttfBytes)
	if err != nil {
		log.Fatal(err)
	}

	name, style := strings.ToLower(font.Name(truetype.NameIDFontFamily)), strings.ToLower(font.Name(truetype.NameIDFontSubfamily))

	if stylemap, ok := c[name]; !ok {
		stylemap := make(Styles)
		c[name] = stylemap
		stylemap[style] = font
	} else {
		stylemap[style] = font
	}

	return nil
}
