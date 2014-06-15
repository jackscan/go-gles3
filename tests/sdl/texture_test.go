// Copyright 2014 Mathias Fiedler
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

package main

import (
	"../../gl"
	"github.com/veandco/go-sdl2/sdl"
	"testing"
)

func TestTexture(t *testing.T) {
	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		panic(sdl.GetError())
	}
	defer sdl.Quit()

	window := sdl.CreateWindow("test", 0, 0, 320, 200, sdl.WINDOW_HIDDEN|sdl.WINDOW_OPENGL)
	defer window.Destroy()

	sdl.GL_SetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_ES)
	sdl.GL_SetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 2)
	sdl.GL_SetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 0)

	glctx := sdl.GL_CreateContext(window)
	defer sdl.GL_DeleteContext(glctx)

	tex := gl.CreateTexture()
	gl.BindTexture(gl.TEXTURE_2D, tex)
	if !gl.IsTexture(tex) {
		t.Error(tex, "is not a texture")
	}
}
