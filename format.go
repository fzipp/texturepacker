// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package texturepacker

// Format represents a generic parser for a sprite sheet data format.
type Format interface {
	SpriteSheetFrom(data []byte) (*SpriteSheet, error)
}
