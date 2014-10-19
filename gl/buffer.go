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

type Buffer C.GLuint
type BufferTarget C.GLenum
type BufferUsage C.GLenum
type BufferParameter C.GLenum

const (
	ARRAY_BUFFER_BINDING         = C.GL_ARRAY_BUFFER_BINDING
	ELEMENT_ARRAY_BUFFER_BINDING = C.GL_ELEMENT_ARRAY_BUFFER_BINDING

	ARRAY_BUFFER              BufferTarget = C.GL_ARRAY_BUFFER
	ELEMENT_ARRAY_BUFFER      BufferTarget = C.GL_ELEMENT_ARRAY_BUFFER
	TRANSFORM_FEEDBACK_BUFFER BufferTarget = C.GL_TRANSFORM_FEEDBACK_BUFFER
	PIXEL_UNPACK_BUFFER       BufferTarget = C.GL_PIXEL_UNPACK_BUFFER
	PIXEL_PACK_BUFFER         BufferTarget = C.GL_PIXEL_PACK_BUFFER
	COPY_WRITE_BUFFER         BufferTarget = C.GL_COPY_WRITE_BUFFER
	COPY_READ_BUFFER          BufferTarget = C.GL_COPY_READ_BUFFER
	UNIFORM_BUFFER            BufferTarget = C.GL_UNIFORM_BUFFER

	STREAM_DRAW  BufferUsage = C.GL_STREAM_DRAW
	STATIC_DRAW  BufferUsage = C.GL_STATIC_DRAW
	DYNAMIC_DRAW BufferUsage = C.GL_DYNAMIC_DRAW
	STREAM_READ  BufferUsage = C.GL_STREAM_READ
	STATIC_READ  BufferUsage = C.GL_STATIC_READ
	DYNAMIC_READ BufferUsage = C.GL_DYNAMIC_READ
	STREAM_COPY  BufferUsage = C.GL_STREAM_COPY
	STATIC_COPY  BufferUsage = C.GL_STATIC_COPY
	DYNAMIC_COPY BufferUsage = C.GL_DYNAMIC_COPY

	BUFFER_ACCESS_FLAGS BufferParameter = C.GL_BUFFER_ACCESS_FLAGS
	BUFFER_MAPPED       BufferParameter = C.GL_BUFFER_MAPPED
	BUFFER_MAP_LENGTH   BufferParameter = C.GL_BUFFER_MAP_LENGTH
	BUFFER_MAP_OFFSET   BufferParameter = C.GL_BUFFER_MAP_OFFSET
	BUFFER_SIZE         BufferParameter = C.GL_BUFFER_SIZE
	BUFFER_USAGE        BufferParameter = C.GL_BUFFER_USAGE
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

func BufferData(target BufferTarget, size int, usage BufferUsage) {
	C.glBufferData(C.GLenum(target), C.GLsizeiptr(size), nil, C.GLenum(usage))
}

func (b Buffer) Delete() {
	C.glDeleteBuffers(C.GLsizei(1), (*C.GLuint)(&b))
}

func DeleteBuffers(buffers []Buffer) {
	C.glDeleteBuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
}

func BufferSubData(target BufferTarget, offset int, data []float32) {
	C.glBufferSubData(C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(unsafe.Sizeof(data[0])*uintptr(len(data))), unsafe.Pointer(&data[0]))
}

func CopyBufferSubData(readTarget, writeTarget BufferTarget, readOffset, writeOffset, size int) {
	C.glCopyBufferSubData(C.GLenum(readTarget), C.GLenum(writeTarget), C.GLintptr(readOffset), C.GLintptr(writeOffset), C.GLsizeiptr(size))
}

func GetBufferParameter(target BufferTarget, param BufferParameter) int {
	value := C.GLint(0)
	C.glGetBufferParameteriv(C.GLenum(target), C.GLenum(param), &value)
	return int(value)
}

// func GetBufferParameteriv(target int, pname int,  params int) {
// 	C.glGetBufferParameteriv(GLenum target, GLenum pname, GLint* params)
// }
