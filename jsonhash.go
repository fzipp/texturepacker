// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package texturepacker

import (
	"encoding/json"
	"fmt"
	"image"
)

// FormatJSONHash implements the Format interface and is a parser for
// the TexturePacker output data format "JSON (Hash)".
type FormatJSONHash struct{}

// SpriteSheetFrom parses the given JSON data in the TexturePacker
// output data format "JSON (Hash)" and returns a sprite sheet.
func (f FormatJSONHash) SpriteSheetFrom(data []byte) (*SpriteSheet, error) {
	var sheet sheetJSON
	err := json.Unmarshal(data, &sheet)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal sprite sheet JSON: %w", err)
	}
	return toSpriteSheet(&sheet), nil
}

type sheetJSON struct {
	Frames map[string]*spriteJSON `json:"frames"`
	Meta   *metadataJSON          `json:"meta"`
}

type spriteJSON struct {
	Frame            intRectJSON    `json:"frame"`
	Rotated          bool           `json:"rotated"`
	Trimmed          bool           `json:"trimmed"`
	SpriteSourceSize intRectJSON    `json:"spriteSourceSize"`
	SourceSize       intSizeJSON    `json:"sourceSize"`
	Pivot            floatPointJSON `json:"pivot"`
}

type metadataJSON struct {
	App         string      `json:"app"`
	Version     string      `json:"version"`
	Image       string      `json:"image"`
	Format      string      `json:"format"`
	Size        intSizeJSON `json:"size"`
	Scale       string      `json:"scale"`
	SmartUpdate string      `json:"smartupdate"`
}

type intRectJSON struct {
	intPointJSON
	intSizeJSON
}

type intPointJSON struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type intSizeJSON struct {
	W int `json:"w"`
	H int `json:"h"`
}

type floatPointJSON struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func toSpriteSheet(s *sheetJSON) *SpriteSheet {
	return &SpriteSheet{
		Sprites: toSprites(s.Frames),
		Meta:    toMetadata(s.Meta),
	}
}

func toSprites(frames map[string]*spriteJSON) map[string]*Sprite {
	sprites := make(map[string]*Sprite, len(frames))
	for name, frame := range frames {
		sprites[name] = toSprite(frame)
	}
	return sprites
}

func toSprite(s *spriteJSON) *Sprite {
	return &Sprite{
		Frame:            toRectangle(s.Frame),
		Rotated:          s.Rotated,
		Trimmed:          s.Trimmed,
		SpriteSourceSize: toRectangle(s.SpriteSourceSize),
		SourceSize:       toPoint(s.SourceSize),
		Pivot:            toVec2(s.Pivot),
	}
}

func toRectangle(r intRectJSON) image.Rectangle {
	return image.Rect(
		r.X, r.Y,
		r.X+r.intSizeJSON.W, r.Y+r.intSizeJSON.H,
	)
}

func toPoint(s intSizeJSON) image.Point {
	return image.Pt(s.W, s.H)
}

func toVec2(p floatPointJSON) FloatPoint {
	return FloatPoint(p)
}

func toMetadata(meta *metadataJSON) *Metadata {
	return &Metadata{
		App:         meta.App,
		Version:     meta.Version,
		Image:       meta.Image,
		Format:      meta.Format,
		Size:        toPoint(meta.Size),
		Scale:       meta.Scale,
		SmartUpdate: meta.SmartUpdate,
	}
}
