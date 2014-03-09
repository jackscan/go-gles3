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

type Buffer C.GLuint
type BufferTarget C.GLenum
type BufferUsage C.GLenum

const (
	ARRAY_BUFFER_BINDING         = C.GL_ARRAY_BUFFER_BINDING
	ELEMENT_ARRAY_BUFFER_BINDING = C.GL_ELEMENT_ARRAY_BUFFER_BINDING

	BUFFER_SIZE  = C.GL_BUFFER_SIZE
	BUFFER_USAGE = C.GL_BUFFER_USAGE

	ARRAY_BUFFER         BufferTarget = C.GL_ARRAY_BUFFER
	ELEMENT_ARRAY_BUFFER              = C.GL_ELEMENT_ARRAY_BUFFER

	STREAM_DRAW  BufferUsage = C.GL_STREAM_DRAW
	STATIC_DRAW              = C.GL_STATIC_DRAW
	DYNAMIC_DRAW             = C.GL_DYNAMIC_DRAW
)

func GenBuffers(buffers []Buffer) {
	C.glGenBuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
}

func CreateBuffer() Buffer {
	buffer := Buffer(0)
	C.glGenBuffers(C.GLsizei(1), (*C.GLuint)(&buffer))
	return buffer
}

func BindBuffer(target BufferTarget, buffer Buffer) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(buffer))
}

func BufferDataf(target BufferTarget, data []float32, usage BufferUsage) {
	C.glBufferData(C.GLenum(target), C.GLsizeiptr(unsafe.Sizeof(data[0])*uintptr(len(data))), unsafe.Pointer(&data[0]), C.GLenum(usage))
}

func (b Buffer) Delete() {

	C.glDeleteBuffers(C.GLsizei(1), (*C.GLuint)(&b))
}

// func BufferSubData(target, offset int, data []interface{}) {
// 	C.glBufferSubData(GLenum target, GLintptr offset, GLsizeiptr size, const GLvoid* data)
// }

// func DeleteBuffers(GLsizei n, buffers []uint) {
// 	C.glDeleteBuffers(GLsizei n, const GLuint* buffers)
// }

// func GetBufferParameteriv(target int, pname int,  params int) {
// 	C.glGetBufferParameteriv(GLenum target, GLenum pname, GLint* params)
// }
