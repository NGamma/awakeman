// Copyright 2016 Josh Deprez
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package game

import (
	"github.com/DrJosh9000/awakeman/common"
	"github.com/DrJosh9000/awakengine"
	"github.com/DrJosh9000/vec"
)

const (
	// One frame of animation for every (animationPeriod) frames rendered.
	// So animation FPS = 60 / animationPeriod.
	animPeriod = 3

	windowTitle = "Awakeman! #40: Escape from the Dark Library"
)

// Game implements awakengine.Game
type Game struct {
	pixelSize    int
	camSize      vec.I2
	levelPreview bool
}

// New creates a new Game.
func New(levelPreview bool) *Game {
	ps := 3
	cs := vec.I2{267, 150}
	if levelPreview {
		ps = 1
		//cs = vec.I2{1024, 512}
		cs = vec.I2{512, 256}
	}
	return &Game{
		pixelSize:    ps,
		camSize:      cs,
		levelPreview: levelPreview,
	}
}

func (*Game) BubbleKey() string { return "inv_bubble" }

// Font returns the default typeface.
func (*Game) Font() awakengine.Font {
	return common.MunroFont{
		Invert: true,
	}
}

func (*Game) Handle(t int, e awakengine.Event) {
	if e.Type == awakengine.EventMouseUp {
		goalAckMarker.begin(e.Pos, t)
		player.path = awakengine.Navigate(player.Pos(), e.Pos)
	}
}

// Player returns the player unit.
func (*Game) Player() awakengine.Unit {
	return player
}

// Objects provides all sprites in the level.
func (*Game) Objects() []awakengine.Object {
	return []awakengine.Object{
		&awakengine.SpriteObject{Sprite: player, Semiobject: player},
		&awakengine.SpriteObject{Sprite: goalAckMarker, Semiobject: goalAckMarker},
	}
}

// Viewport is the size of the window and the pixels in the window.
func (g *Game) Viewport() (cs vec.I2, ps, ap int, title string) {
	return g.camSize, g.pixelSize, animPeriod, windowTitle
}