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

package gl

// #include <GLES3/gl3.h>
import "C"
import (
	"unsafe"
)

type VertexAttrib interface {
	Index() C.GLuint
	Enable()
	Disable()
}

type Attrib struct {
	index C.GLuint
}

type FloatAttrib struct {
	Attrib
}

type Vec2Attrib struct {
	FloatAttrib
}

type Vec3Attrib struct {
	FloatAttrib
}

type Vec4Attrib struct {
	FloatAttrib
}

func (attrib Attrib) Index() C.GLuint {
	return attrib.index
}

func (a FloatAttrib) Set(x float32) {
	C.glVertexAttrib1f(a.index, C.GLfloat(x))
}

func (a Vec2Attrib) Set(x, y float32) {
	C.glVertexAttrib2f(a.index, C.GLfloat(x), C.GLfloat(y))
}

func (a Vec3Attrib) Set(x, y, z float32) {
	C.glVertexAttrib3f(a.index, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func (a Vec4Attrib) Set(x, y, z, w float32) {
	C.glVertexAttrib4f(a.index, C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

func (a FloatAttrib) Setv(values []float32) {
	C.glVertexAttrib1fv(a.index, (*C.GLfloat)(&values[0]))
}

func (a Vec2Attrib) Setv(values []float32) {
	C.glVertexAttrib2fv(a.index, (*C.GLfloat)(&values[0]))
}

func (a Vec3Attrib) Setv(values []float32) {
	C.glVertexAttrib3fv(a.index, (*C.GLfloat)(&values[0]))
}

func (a Vec4Attrib) Setv(values []float32) {
	C.glVertexAttrib4fv(a.index, (*C.GLfloat)(&values[0]))
}

func (attrib FloatAttrib) Pointerf(size int, stride int, data []float32) {
	C.glVertexAttribPointer(attrib.index, C.GLint(size), C.GL_FLOAT, C.GL_FALSE, C.GLsizei(stride), unsafe.Pointer(&data[0]))
}

func (attrib Attrib) Pointer(size int, stride int, offset uintptr) {
	C.glVertexAttribPointer(attrib.index, C.GLint(size), C.GL_FLOAT, C.GL_FALSE, C.GLsizei(stride), unsafe.Pointer(offset))
}

// func (attrib VertexAttrib) Pointer(size int, datatype DataType, normalized bool, stride int, offset uintptr) {
// 	nflag := C.GLboolean(C.GL_FALSE)
// 	if normalized {
// 		nflag = C.GLboolean(C.GL_TRUE)
// 	}
// 	C.glVertexAttribPointer(C.GLuint(attrib), C.GLint(size), C.GLenum(datatype), nflag, C.GLsizei(stride), unsafe.Pointer(offset))
// }

// func (attrib VertexAttrib) Pointerf(size int, datatype DataType, stride int, data []float32) {
// 	C.glVertexAttribPointer(C.GLuint(attrib), C.GLint(size), C.GLenum(datatype), C.GL_FALSE, C.GLsizei(stride), unsafe.Pointer(&data[0]))
// }

func (attrib Attrib) Divisor(divisor uint) {
	C.glVertexAttribDivisor(attrib.index, C.GLuint(divisor))
}

func (attrib Attrib) Enable() {
	C.glEnableVertexAttribArray(attrib.index)
}

func (attrib Attrib) Disable() {
	C.glDisableVertexAttribArray(attrib.index)
}
