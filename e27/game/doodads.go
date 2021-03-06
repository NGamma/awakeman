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
	"github.com/DrJosh9000/awakengine"
	"github.com/DrJosh9000/vec"
)

var (
	treesSheet = &awakengine.Sheet{
		Key:       "trees",
		Frames:    9,
		FrameSize: vec.I2{64, 64},
	}
	otherDoodadsSheet = &awakengine.Sheet{
		Key:       "doodads",
		Frames:    6,
		FrameSize: vec.I2{32, 32},
	}

	baseDoodads = map[string]*awakengine.BaseDoodad{
		// --------------------- TREES ---------------------
		"tree1": {
			Sheet:  treesSheet,
			Offset: vec.I2{30, 59},
			Frame:  0,
			UL:     vec.I2{25, 56},
			DR:     vec.I2{34, 63},
		},
		"tree2": {
			Sheet:  treesSheet,
			Offset: vec.I2{33, 61},
			Frame:  1,
			UL:     vec.I2{29, 55},
			DR:     vec.I2{38, 63},
		},
		"tree3": {
			Sheet:  treesSheet,
			Offset: vec.I2{9, 58},
			Frame:  2,
			UL:     vec.I2{3, 53},
			DR:     vec.I2{60, 62},
		},
		"tree4": {
			Sheet:  treesSheet,
			Offset: vec.I2{31, 60},
			Frame:  3,
			UL:     vec.I2{4, 53},
			DR:     vec.I2{59, 61},
		},
		"tree5": {
			Sheet:  treesSheet,
			Offset: vec.I2{31, 59},
			Frame:  4,
			UL:     vec.I2{19, 52},
			DR:     vec.I2{44, 63},
		},
		"tree6": {
			Sheet:  treesSheet,
			Offset: vec.I2{31, 56},
			Frame:  5,
			UL:     vec.I2{22, 51},
			DR:     vec.I2{34, 57},
		},
		"tree7": {
			Sheet:  treesSheet,
			Offset: vec.I2{32, 58},
			Frame:  6,
			UL:     vec.I2{19, 52},
			DR:     vec.I2{44, 63},
		},
		"tree8": {
			Sheet:  treesSheet,
			Offset: vec.I2{31, 59},
			Frame:  7,
			UL:     vec.I2{18, 56},
			DR:     vec.I2{41, 63},
		},
		"tree9": {
			Sheet:  treesSheet,
			Offset: vec.I2{7, 59},
			Frame:  8,
			UL:     vec.I2{3, 53},
			DR:     vec.I2{63, 61},
		},

		// --------------------- SMALL DOODADS ---------------------
		"bench_h": {
			Sheet:  otherDoodadsSheet,
			Offset: vec.I2{16, 24},
			Frame:  0,
			UL:     vec.I2{1, 15},
			DR:     vec.I2{30, 24},
		},
		/*	"W": {
			Sheet: otherDoodadsSheet,
				Offset:    vec.I2{16, 27},
			Frame:  1,
			UL: vec.I2{3, 23},
			DR: vec.I2{26, 29},
		},*/
		"puffyplant": {
			Sheet:  otherDoodadsSheet,
			Offset: vec.I2{16, 24},
			Frame:  2,
			UL:     vec.I2{6, 24},
			DR:     vec.I2{24, 27},
		},
		"evilplant": {
			Sheet:  otherDoodadsSheet,
			Offset: vec.I2{16, 24},
			Frame:  3,
			UL:     vec.I2{2, 22},
			DR:     vec.I2{30, 27},
		},
		"bench_v": {
			Sheet:  otherDoodadsSheet,
			Offset: vec.I2{8, 22},
			Frame:  4,
			UL:     vec.I2{4, 13},
			DR:     vec.I2{15, 30},
		},
		"alamore": {
			Sheet:  otherDoodadsSheet,
			Offset: vec.I2{16, 25},
			Frame:  5,
			UL:     vec.I2{8, 20},
			DR:     vec.I2{23, 28},
		},
	}
)
