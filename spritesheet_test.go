package texturepacker

import (
	"errors"
	"image"
	"os"
	"reflect"
	"testing"
)

func TestSheetFromFile(t *testing.T) {
	path := "testdata/TestSheet.json"
	sheet, err := SheetFromFile(path, FormatJSONHash{})
	if err != nil {
		t.Errorf("could not read %s: %s", path, err)
	}
	want := &SpriteSheet{
		Sprites: map[string]*Sprite{
			"test_sprite1": {
				Frame:            image.Rect(0, 0, 15, 20),
				Rotated:          false,
				Trimmed:          false,
				SpriteSourceSize: image.Rect(0, 0, 15, 20),
				SourceSize:       image.Pt(15, 20),
				Pivot:            FloatPoint{0.5, 0.5},
			},
			"test_sprite2": {
				Frame:            image.Rect(18, 5, 43, 37),
				Rotated:          false,
				Trimmed:          true,
				SpriteSourceSize: image.Rect(30, 34, 62, 74),
				SourceSize:       image.Pt(60, 100),
				Pivot:            FloatPoint{0.5, 0.5},
			},
		},
		Meta: &Metadata{
			App:         "https://www.codeandweb.com/texturepacker",
			Version:     "1.0",
			Image:       "TestSheet.png",
			Format:      "RGBA8888",
			Size:        image.Pt(400, 200),
			Scale:       "1",
			SmartUpdate: "$TexturePacker:SmartUpdate:02448c254061b96dab05420adaf254e2:f8d8f202601d38ac1f1f09505bd108b5:0d78c8d373c64903321e3878035421d8$",
		},
	}
	if !reflect.DeepEqual(sheet.Meta, want.Meta) {
		t.Errorf("sheet.Meta was: %v, want: %v", sheet.Meta, want.Meta)
	}
	if len(sheet.Sprites) != len(want.Sprites) {
		t.Errorf("len(sheet.Sprites) was: %d, want: %d", len(sheet.Sprites), len(want.Sprites))
	}
	for name, spriteWant := range want.Sprites {
		if !reflect.DeepEqual(sheet.Sprites[name], spriteWant) {
			t.Errorf("sheet.Sprites[%q] was: %v, want: %v", name, sheet.Sprites[name], spriteWant)
		}
	}
}

func TestSheetFromFileNotExisting(t *testing.T) {
	_, err := SheetFromFile("testdata/TestSheetNotExisting.json", FormatJSONHash{})
	if err == nil {
		t.Errorf("no error, but wanted error")
	}
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("wanted ErrNotExist, but wasn't")
	}
}

func TestSheetFromFileInvalidFormat(t *testing.T) {
	_, err := SheetFromFile("testdata/InvalidSheet.json", FormatJSONHash{})
	if err == nil {
		t.Errorf("no error, but wanted error")
	}
}
