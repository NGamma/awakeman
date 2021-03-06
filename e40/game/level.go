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

import "github.com/DrJosh9000/awakengine"

const tileSize = 16

func (*Game) Level() (*awakengine.Level, error) {
	bm, sz, err := awakengine.ImageAsMap("level2blockmap")
	if err != nil {
		return nil, err
	}
	return &awakengine.Level{
		Doodads:  nil,
		MapSize:  sz,
		BlockMap: bm,
		BlockInfos: []awakengine.TileInfo{
			{Name: "black"}, // 0
			{Name: "prism", Blocking: true},
			{Name: "shelf_d_1", Blocking: true},
			{Name: "shelf_d_2", Blocking: true},
			{Name: "shelf_d_3", Blocking: true},
			{Name: "shelf_d_4", Blocking: true}, // 5
			{Name: "shelf_r_1", Blocking: true},
			{Name: "shelf_r_2", Blocking: true},
			{Name: "shelf_r_3", Blocking: true},
			{Name: "column", Blocking: true},
			{Name: "stack", Blocking: true}, // 10
			{Name: "desk", Blocking: true},
			{Name: "wall_h", Blocking: true},
			{Name: "wall_v", Blocking: true},
			{Name: "wall_ul", Blocking: true},
			{Name: "wall_ur", Blocking: true}, // 15
			{Name: "wall_dr", Blocking: true},
			{Name: "wall_dl", Blocking: true},
			{Name: "wall_+", Blocking: true},
			{Name: "wall_-|", Blocking: true},
			{Name: "wall_|-", Blocking: true}, // 20
			{Name: "wall_T'", Blocking: true},
			{Name: "wall_T", Blocking: true},
			{Name: "wall_hr", Blocking: true},
			{Name: "wall_hl", Blocking: true},
			{Name: "wall_vu", Blocking: true}, // 25
			{Name: "wall_vd", Blocking: true},
			{Name: "exit_door", Blocking: true},
			{Name: "exit_left", Blocking: true},
			{Name: "exit_right", Blocking: true},
		},
		BlocksetKey: "libraryblocks",
		TileSize:    tileSize,
		BlockHeight: 32,
		Obstacles:   precomputedObstacles,
		Paths:       precomputedPaths,
	}, nil
}
