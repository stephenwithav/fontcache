## fontcache

A simple font-caching utility enabled by [golang/freetype#19](https://github.com/golang/freetype/pull/19).

A simple example:

```go
import (
    "log"
    "github.com/stephenwithav/fontcache"
)

func main() {
    fontPath := "./fonts"

    fc := fontcache.New()
    fc.Init(fontPath)

    // Retrieve font by name for use in a program.
    font, err := fc.Get("Luxi Sans", "Regular")
    if err != nil {
	    log.Fatalf("%s", err)
    }

    DoStuff(font)
}
```
