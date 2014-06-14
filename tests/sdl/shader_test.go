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

func TestBindVariables(t *testing.T) {
	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		panic(sdl.GetError())
	}

	window := sdl.CreateWindow("test", 0, 0, 320, 200, sdl.WINDOW_HIDDEN|sdl.WINDOW_OPENGL)

	sdl.GL_SetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_ES)
	sdl.GL_SetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 2)
	sdl.GL_SetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 0)

	glctx := sdl.GL_CreateContext(window)

	vs := gl.CreateShader(gl.VERTEX_SHADER)
	vs.Source(`
uniform mat4 u_projection;
uniform vec2 u_size;
attribute vec2 a_vertex;
attribute vec2 a_texcoord;
varying vec2 v_texcoord;

void main()
{
	v_texcoord = a_texcoord;
	gl_Position = u_projection * vec4(a_vertex * u_size, 0, 1);
}
`)
	vs.Compile()

	fs := gl.CreateShader(gl.FRAGMENT_SHADER)
	fs.Source(`
precision mediump float;

uniform sampler2D u_tex;
uniform vec4 u_color;

varying vec2 v_texcoord;

void main()
{
	gl_FragColor = texture2D(u_tex, v_texcoord).a * u_color;
}
`)
	fs.Compile()

	p := gl.CreateProgram()

	p.AttachShader(vs)
	p.AttachShader(fs)
	p.Link()

	var v struct {
		Proj      gl.UniformMatrix4f `glsl:"u_projection"`
		Size      gl.Uniform2f       `glsl:"u_size"`
		Tex       gl.Uniform1i       `glsl:"u_tex"`
		Color     gl.Uniform4f       `glsl:"u_color"`
		Vertices  gl.Vec2Attrib      `glsl:"a_vertex"`
		TexCoords gl.Vec2Attrib      `glsl:"a_texcoord"`
	}

	p.BindVariables(&v)

	uniforms := []gl.Uniform{v.Proj, v.Size, v.Tex, v.Color}
	cksum := 0
	for _, u := range uniforms {
		cksum |= 1 << uint(u.Location())
	}

	if cksum != 0xF {
		t.Errorf("unexpected uniform locations: %#x", cksum)
	}

	attributes := []gl.VertexAttrib{v.Vertices, v.TexCoords}
	cksum = 0
	for _, a := range attributes {
		cksum |= 1 << uint(a.Index())
	}

	if cksum != 0x3 {
		t.Errorf("unexpected attribute locations: %#x", cksum)
	}

	sdl.GL_DeleteContext(glctx)
	window.Destroy()
	sdl.Quit()
}
