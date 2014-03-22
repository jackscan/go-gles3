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
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type Shader C.GLuint
type ShaderType C.GLenum
type ShaderBinaryFormat C.GLint

const (
	VERTEX_SHADER   ShaderType = C.GL_VERTEX_SHADER
	FRAGMENT_SHADER            = C.GL_FRAGMENT_SHADER
)

type Program struct {
	id C.GLuint

	shadersValid bool
	shaders      []Shader
}

func CreateProgram() Program {
	return Program{id: C.glCreateProgram()}
}

func CreateShader(stype ShaderType) Shader {
	return Shader(C.glCreateShader(C.GLenum(stype)))
}

func (program *Program) Shaders() []Shader {
	if !program.shadersValid {
		nshaders := C.GLint(0)
		C.glGetProgramiv(program.id, C.GL_ATTACHED_SHADERS, &nshaders)
		program.shaders = make([]Shader, nshaders)
		C.glGetAttachedShaders(program.id, C.GLsizei(nshaders), nil, (*C.GLuint)(&program.shaders[0]))
		program.shadersValid = true
	}
	return program.shaders
}

func (program *Program) AttachShader(shader Shader) {
	C.glAttachShader(program.id, C.GLuint(shader))
	program.shadersValid = false
}

func (program *Program) BindAttribLocation(attrib VertexAttrib, name string) {

	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	C.glBindAttribLocation(program.id, attrib.Index(), (*C.GLchar)(cstr))
}

func (program *Program) Delete() {
	C.glDeleteProgram(program.id)
}

func (shader Shader) Delete() {
	C.glDeleteShader(C.GLuint(shader))
}

func (program *Program) DetachShader(shader Shader) {
	C.glDetachShader(program.id, C.GLuint(shader))
	program.shadersValid = false
}

// func (program *Program) GetActiveAttrib(index uint) (size int, atype int, name string) {

// 	nameBuf := make([]byte, MAX_ATTRIBUTE_LENGTH)
// 	sizeBuf := C.GLint(0)
// 	typeBuf := C.GLenum(0)

// 	C.glGetActiveAttrib(program.id, C.GLuint(index), C.GLsizei(len(nameBuf)), C.NULL, &sizeBuf, &typeBuf, (*C.GLchar)(&nameBuf[0]))

// 	size = int(sizeBuf)
// 	atype = int(typeBuf)
// 	name = string(nameBuf)
// }

// func (program *Program) GetActiveUniform(index uint) (size int, utype int, name string) {

// 	nameBuf := make([]byte, MAX_UNIFORM_LENGTH)
// 	sizeBuf := C.GLint(0)
// 	typeBuf := C.GLenum(0)

// 	C.glGetActiveUniform(program.id, C.GLuint(index), C.GLsizei(len(nameBuf)), C.NULL, &sizeBuf, &typeBuf, (*C.GLchar)(&nameBuf[0]))

// 	size = int(sizeBuf)
// 	utype = int(typeBuf)
// 	name = string(nameBuf)
// }

// func (program *Program) GetAttachedShaders(GLsizei maxcount, GLsizei* count) []uint {
// 	C.glGetAttachedShaders(program.id, GLsizei maxcount, GLsizei* count, GLuint* shaders)
// }

// func (program *Program) GetProgramiv(pname int,  params int) {
// 	C.glGetProgramiv(program.id, GLenum pname, GLint* params)
// }

// func (program *Program) InfoLog(GLsizei bufsize, GLsizei* length, GLchar* infolog) {
// 	C.glGetProgramInfoLog(program.id, GLsizei bufsize, GLsizei* length, GLchar* infolog)
// }

// func GetShaderiv(shader Shader, pname int, params int) {
// 	C.glGetShaderiv(C.GLuint(shader), GLenum pname, GLint* params)
// }

// func GetShaderInfoLog(shader Shader, GLsizei bufsize, GLsizei* length, GLchar* infolog) {
// 	C.glGetShaderInfoLog(C.GLuint(shader), GLsizei bufsize, GLsizei* length, GLchar* infolog)
// }

// func GetShaderPrecisionFormat(shadertype ShaderType, precisiontype int, GLint* range,  precision int) {
// 	C.glGetShaderPrecisionFormat(GLenum shadertype, GLenum precisiontype, GLint* range, GLint* precision)
// }

// func GetShaderSource(shader Shader, GLsizei bufsize, GLsizei* length, GLchar* source) {
// 	C.glGetShaderSource(C.GLuint(shader), GLsizei bufsize, GLsizei* length, GLchar* source)
// }

// func (program *Program) GetUniformfv(location int, GLfloat* params) {
// 	C.glGetUniformfv(program.id, GLint location, GLfloat* params)
// }

// func (program *Program) GetUniformiv(GLint location,  params int) {
// 	C.glGetUniformiv(program.id, GLint location, GLint* params)
// }

// GL_APICALL int          GL_APIENTRY glGetAttribLocation (GLuint program, const GLchar* name);
// GL_APICALL int          GL_APIENTRY glGetUniformLocation (GLuint program, const GLchar* name);

// func ShaderBinary(GLsizei n, shaders []uint, binaryformat int, const GLvoid* binary, GLsizei length) {
// 	C.glShaderBinary(GLsizei n, const GLuint* shaders, GLenum binaryformat, const GLvoid* binary, GLsizei length)
// }

func GetShaderBinaryFormats() (formats []ShaderBinaryFormat) {
	numFormats := C.GLint(0)
	C.glGetIntegerv(C.GL_NUM_SHADER_BINARY_FORMATS, &numFormats)
	if numFormats > 0 {
		formats = make([]ShaderBinaryFormat, numFormats)
		C.glGetIntegerv(C.GL_SHADER_BINARY_FORMATS, (*C.GLint)(&formats[0]))
	}
	return
}

func ShaderBinary(shaders []Shader, format ShaderBinaryFormat, binary []byte) {
	C.glShaderBinary(C.GLsizei(len(shaders)), (*C.GLuint)(&shaders[0]), C.GLenum(format), unsafe.Pointer(&binary[0]), C.GLsizei(len(binary)))
}

func (shader Shader) Source(src string) {
	cstr := (*C.GLchar)(C.CString(src))
	defer C.free(unsafe.Pointer(cstr))
	slen := C.GLint(len(src))
	C.glShaderSource(C.GLuint(shader), 1, &cstr, &slen)
}

func (shader Shader) Compile() {
	C.glCompileShader(C.GLuint(shader))

	status := C.GLint(0)
	C.glGetShaderiv(C.GLuint(shader), C.GL_COMPILE_STATUS, &status)

	if status != C.GL_TRUE {
		loglen := C.GLint(0)
		C.glGetShaderiv(C.GLuint(shader), C.GL_INFO_LOG_LENGTH, &loglen)
		log := (*C.GLchar)(C.malloc(C.size_t(loglen)))
		defer C.free(unsafe.Pointer(log))
		C.glGetShaderInfoLog(C.GLuint(shader), C.GLsizei(loglen), nil, log)
		panic(fmt.Errorf("Failed to compile shader: %s", C.GoString((*C.char)(log))))
	}
}

func (program *Program) Link() {
	C.glLinkProgram(program.id)

	status := C.GLint(0)
	C.glGetProgramiv(program.id, C.GL_LINK_STATUS, &status)

	if status != C.GL_TRUE {
		loglen := C.GLint(0)
		C.glGetProgramiv(program.id, C.GL_INFO_LOG_LENGTH, &loglen)
		log := (*C.GLchar)(C.malloc(C.size_t(loglen)))
		defer C.free(unsafe.Pointer(log))
		C.glGetProgramInfoLog(program.id, C.GLsizei(loglen), nil, log)
		panic(fmt.Errorf("Failed to link shader: %s", C.GoString((*C.char)(log))))
	}
}

func (program *Program) Use() {
	C.glUseProgram(program.id)
}

func (program *Program) Validate() {
	C.glValidateProgram(program.id)
}

func ReleaseShaderCompiler() {
	C.glReleaseShaderCompiler()
}

func (t ShaderType) String() string {
	switch t {
	case VERTEX_SHADER:
		return "VERTEX_SHADER"
	case FRAGMENT_SHADER:
		return "FRAGMENT_SHADER"
	default:
		return fmt.Sprintf("ShaderType(%#x)", t)
	}
}
