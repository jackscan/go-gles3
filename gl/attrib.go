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

// #include <GLES2/gl2.h>
import "C"
import (
	"unsafe"
)

type VertexAttrib interface {
	Index() C.GLuint
	Enable()
	Disable()
}

type FloatAttrib struct {
	index C.GLuint
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

func (attrib FloatAttrib) Index() C.GLuint {
	return attrib.index
}

// func VertexAttrib1f(indx uint, GLfloat x) {
// 	C.glVertexAttrib1f(GLuint indx, GLfloat x)
// }

// func VertexAttrib1fv(indx uint, const GLfloat* values) {
// 	C.glVertexAttrib1fv(GLuint indx, const GLfloat* values)
// }

// func VertexAttrib2f(indx uint, GLfloat x, GLfloat y) {
// 	C.glVertexAttrib2f(GLuint indx, GLfloat x, GLfloat y)
// }

// func VertexAttrib2fv(indx uint, const GLfloat* values) {
// 	C.glVertexAttrib2fv(GLuint indx, const GLfloat* values)
// }

// func VertexAttrib3f(indx uint, GLfloat x, GLfloat y, GLfloat z) {
// 	C.glVertexAttrib3f(GLuint indx, GLfloat x, GLfloat y, GLfloat z)
// }

// func VertexAttrib3fv(indx uint, const GLfloat* values) {
// 	C.glVertexAttrib3fv(GLuint indx, const GLfloat* values)
// }

// func VertexAttrib4f(indx uint, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
// 	C.glVertexAttrib4f(GLuint indx, GLfloat x, GLfloat y, GLfloat z, GLfloat w)
// }

// func VertexAttrib4fv(indx uint, const GLfloat* values) {
// 	C.glVertexAttrib4fv(GLuint indx, const GLfloat* values)
// }

func (attrib FloatAttrib) Pointerf(size int, stride int, data []float32) {
	C.glVertexAttribPointer(attrib.index, C.GLint(size), C.GL_FLOAT, C.GL_FALSE, C.GLsizei(stride), unsafe.Pointer(&data[0]))
}

func (attrib FloatAttrib) Pointer(size int, stride int, offset uintptr) {
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

func (attrib FloatAttrib) Enable() {
	C.glEnableVertexAttribArray(attrib.index)
}

func (attrib FloatAttrib) Disable() {
	C.glDisableVertexAttribArray(attrib.index)
}
