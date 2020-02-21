# texturepacker

Package texturepacker reads sprite sheets created and exported as JSON by
[TexturePacker](https://www.codeandweb.com/texturepacker).

Add it to a module as a dependency via:

```
go get github.com/fzipp/texturepacker
```

## Documentation

Package documentation is available [on pkg.go.dev](https://pkg.go.dev/github.com/fzipp/texturepacker?tab=doc).

## Example usage

Create a sprite sheet with TexturePacker and and save it in "JSON (Hash)" format.

![TexturePacker format selection dialog](doc/jsonhashformat.png?raw=true "TexturePacker format selection dialog")

```
package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/fzipp/texturepacker"
)

func main() {
	sheet, err := texturepacker.SheetFromFile("ExampleSheet.json", texturepacker.FormatJSONHash{})
	if err != nil {
		log.Fatal(err)
	}
	imageFile, err := os.Open(sheet.Meta.Image)
	defer imageFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(imageFile)
	sheetImage, ok := img.(image.RGBA)
	if !ok {
		log.Fatal("expected RGBA image")
	}
	for name, sprite := range sheet.Sprites {
		spriteImage := sheetImage.SubImage(sprite.Frame)
		// ...
	}
}
```

## License

This project is free and open source software licensed under the
[BSD 3-Clause License](LICENSE).
