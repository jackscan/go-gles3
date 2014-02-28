package gl

// #include <GLES2/gl2.h>
import "C"

type Uniform interface {
	// Location() C.GLint
}

type uniformBase struct {
	location C.GLint
}

// func (u uniformBase) Location() C.GLint {
// 	return u.location
// }

type Uniform1f struct {
	uniformBase
}

type Uniform2f struct {
	uniformBase
}

type Uniform3f struct {
	uniformBase
}

type Uniform4f struct {
	uniformBase
}

type Uniform1i struct {
	uniformBase
}
type Uniform2i struct {
	uniformBase
}
type Uniform3i struct {
	uniformBase
}
type Uniform4i struct {
	uniformBase
}

type UniformMatrix2f struct {
	uniformBase
}
type UniformMatrix3f struct {
	uniformBase
}
type UniformMatrix4f struct {
	uniformBase
}

func (u Uniform1f) Set(x float32) {
	C.glUniform1f(u.location, C.GLfloat(x))
}

func (u Uniform1f) Setv(data []float32) {
	C.glUniform1fv(u.location, C.GLsizei(len(data)), (*C.GLfloat)(&data[0]))
}

func (u Uniform2f) Set(x, y float32) {
	C.glUniform2f(u.location, C.GLfloat(x), C.GLfloat(y))
}

func (u Uniform2f) Setv(data []float32) {
	C.glUniform2fv(u.location, C.GLsizei(len(data)/2), (*C.GLfloat)(&data[0]))
}

func (u Uniform1i) Set(x int32) {
	C.glUniform1i(u.location, C.GLint(x))
}

func (u Uniform1i) Setv(data []int32) {
	C.glUniform1iv(u.location, C.GLsizei(len(data)), (*C.GLint)(&data[0]))
}

func (u UniformMatrix4f) Set(data [16]float32) {
	C.glUniformMatrix4fv(u.location, 1, C.GLboolean(C.GL_FALSE), (*C.GLfloat)(&data[0]))
}

func (u UniformMatrix4f) Setv(data []float32) {
	C.glUniformMatrix4fv(u.location, C.GLsizei(len(data)/16), C.GLboolean(C.GL_FALSE), (*C.GLfloat)(&data[0]))
}

// func Uniform1f(location int, GLfloat x) {
// 	C.glUniform1f(GLint location, GLfloat x)
// }

// func Uniform1fv(location int, GLsizei count, const GLfloat* v) {
// 	C.glUniform1fv(GLint location, GLsizei count, const GLfloat* v)
// }

// func Uniform1i(GLint location, x int) {
// 	C.glUniform1i(GLint location, GLint x)
// }

// func Uniform1iv(GLint location, GLsizei count, const  v int) {
// 	C.glUniform1iv(GLint location, GLsizei count, const GLint* v)
// }

// func Uniform2f(location int, GLfloat x, GLfloat y) {
// 	C.glUniform2f(GLint location, GLfloat x, GLfloat y)
// }

// func Uniform2fv(location int, GLsizei count, const GLfloat* v) {
// 	C.glUniform2fv(GLint location, GLsizei count, const GLfloat* v)
// }

// func Uniform2i(GLint location, GLint x, y int) {
// 	C.glUniform2i(GLint location, GLint x, GLint y)
// }

// func Uniform2iv(GLint location, GLsizei count, const  v int) {
// 	C.glUniform2iv(GLint location, GLsizei count, const GLint* v)
// }

// func Uniform3f(location int, GLfloat x, GLfloat y, GLfloat z) {
// 	C.glUniform3f(GLint location, GLfloat x, GLfloat y, GLfloat z)
// }

// func Uniform3fv(location int, GLsizei count, const GLfloat* v) {
// 	C.glUniform3fv(GLint location, GLsizei count, const GLfloat* v)
// }

// func Uniform3i(GLint location, GLint x, GLint y, z int) {
// 	C.glUniform3i(GLint location, GLint x, GLint y, GLint z)
// }

// func Uniform3iv(GLint location, GLsizei count, const  v int) {
// 	C.glUniform3iv(GLint location, GLsizei count, const GLint* v)
// }

// func Uniform4f(location int, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
// 	C.glUniform4f(GLint location, GLfloat x, GLfloat y, GLfloat z, GLfloat w)
// }

// func Uniform4fv(location int, GLsizei count, const GLfloat* v) {
// 	C.glUniform4fv(GLint location, GLsizei count, const GLfloat* v)
// }

// func Uniform4i(GLint location, GLint x, GLint y, GLint z, w int) {
// 	C.glUniform4i(GLint location, GLint x, GLint y, GLint z, GLint w)
// }

// func Uniform4iv(GLint location, GLsizei count, const  v int) {
// 	C.glUniform4iv(GLint location, GLsizei count, const GLint* v)
// }

// func UniformMatrix2fv(location int, GLsizei count, GLboolean transpose, const GLfloat* value) {
// 	C.glUniformMatrix2fv(GLint location, GLsizei count, GLboolean transpose, const GLfloat* value)
// }

// func UniformMatrix3fv(location int, GLsizei count, GLboolean transpose, const GLfloat* value) {
// 	C.glUniformMatrix3fv(GLint location, GLsizei count, GLboolean transpose, const GLfloat* value)
// }
