// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package texturepacker reads sprite sheets created and exported as JSON by
// TexturePacker: https://www.codeandweb.com/texturepacker
package texturepacker

import (
	"fmt"
	"image"
	"io/ioutil"
)

// A SpriteSheet is a collection of sprites packed into a single image.
// Each sprite has a name by which it can be looked up and a rectangular area
// within the sheet image. The sheet image file is referenced by its file name
// in the metadata structure.
type SpriteSheet struct {
	Sprites map[string]*Sprite
	Meta    *Metadata
}

// A Sprite is a rectangular area within a sprite sheet image.
type Sprite struct {
	Frame            image.Rectangle
	Rotated          bool
	Trimmed          bool
	SpriteSourceSize image.Rectangle
	SourceSize       image.Point
	Pivot            FloatPoint
}

// A FloatPoint is an X, Y coordinate pair.
type FloatPoint struct{ X, Y float64 }

// The Metadata structure holds information about the sprite sheet image file
// and the application with which the sprite sheet was created.
type Metadata struct {
	App         string
	Version     string
	Image       string
	Format      string
	Size        image.Point
	Scale       string
	SmartUpdate string
}

// SheetFromFile loads and parses sprite sheet information from a file with the
// specified format.
func SheetFromFile(path string, f Format) (*SpriteSheet, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read sprite sheet file: %w", err)
	}
	return SheetFromData(data, f)
}

// SheetFromData parses sprite sheet information from data with the
// specified format.
func SheetFromData(data []byte, f Format) (*SpriteSheet, error) {
	return f.SpriteSheetFrom(data)
}
