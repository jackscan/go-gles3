package gl

// #include <GLES2/gl2.h>
// #cgo LDFLAGS: -lGLESv2
import "C"

import (
	"unsafe"
)

type DataType C.GLenum
type InfoType C.GLenum
type DrawMode C.GLenum
type Error C.GLenum

const (
	ES_VERSION_2_0 = C.GL_ES_VERSION_2_0

	DEPTH_BUFFER_BIT   = C.GL_DEPTH_BUFFER_BIT
	STENCIL_BUFFER_BIT = C.GL_STENCIL_BUFFER_BIT
	COLOR_BUFFER_BIT   = C.GL_COLOR_BUFFER_BIT

	FALSE = C.GL_FALSE
	TRUE  = C.GL_TRUE

	ZERO                = C.GL_ZERO
	ONE                 = C.GL_ONE
	SRC_COLOR           = C.GL_SRC_COLOR
	ONE_MINUS_SRC_COLOR = C.GL_ONE_MINUS_SRC_COLOR
	SRC_ALPHA           = C.GL_SRC_ALPHA
	ONE_MINUS_SRC_ALPHA = C.GL_ONE_MINUS_SRC_ALPHA
	DST_ALPHA           = C.GL_DST_ALPHA
	ONE_MINUS_DST_ALPHA = C.GL_ONE_MINUS_DST_ALPHA

	DST_COLOR           = C.GL_DST_COLOR
	ONE_MINUS_DST_COLOR = C.GL_ONE_MINUS_DST_COLOR
	SRC_ALPHA_SATURATE  = C.GL_SRC_ALPHA_SATURATE

	FUNC_ADD             = C.GL_FUNC_ADD
	BLEND_EQUATION       = C.GL_BLEND_EQUATION
	BLEND_EQUATION_RGB   = C.GL_BLEND_EQUATION_RGB
	BLEND_EQUATION_ALPHA = C.GL_BLEND_EQUATION_ALPHA

	FUNC_SUBTRACT         = C.GL_FUNC_SUBTRACT
	FUNC_REVERSE_SUBTRACT = C.GL_FUNC_REVERSE_SUBTRACT

	BLEND_DST_RGB            = C.GL_BLEND_DST_RGB
	BLEND_SRC_RGB            = C.GL_BLEND_SRC_RGB
	BLEND_DST_ALPHA          = C.GL_BLEND_DST_ALPHA
	BLEND_SRC_ALPHA          = C.GL_BLEND_SRC_ALPHA
	CONSTANT_COLOR           = C.GL_CONSTANT_COLOR
	ONE_MINUS_CONSTANT_COLOR = C.GL_ONE_MINUS_CONSTANT_COLOR
	CONSTANT_ALPHA           = C.GL_CONSTANT_ALPHA
	ONE_MINUS_CONSTANT_ALPHA = C.GL_ONE_MINUS_CONSTANT_ALPHA
	BLEND_COLOR              = C.GL_BLEND_COLOR

	CURRENT_VERTEX_ATTRIB = C.GL_CURRENT_VERTEX_ATTRIB

	FRONT          = C.GL_FRONT
	BACK           = C.GL_BACK
	FRONT_AND_BACK = C.GL_FRONT_AND_BACK

	CULL_FACE                = C.GL_CULL_FACE
	BLEND                    = C.GL_BLEND
	DITHER                   = C.GL_DITHER
	STENCIL_TEST             = C.GL_STENCIL_TEST
	DEPTH_TEST               = C.GL_DEPTH_TEST
	SCISSOR_TEST             = C.GL_SCISSOR_TEST
	POLYGON_OFFSET_FILL      = C.GL_POLYGON_OFFSET_FILL
	SAMPLE_ALPHA_TO_COVERAGE = C.GL_SAMPLE_ALPHA_TO_COVERAGE
	SAMPLE_COVERAGE          = C.GL_SAMPLE_COVERAGE

	CW  = C.GL_CW
	CCW = C.GL_CCW

	LINE_WIDTH                   = C.GL_LINE_WIDTH
	ALIASED_POINT_SIZE_RANGE     = C.GL_ALIASED_POINT_SIZE_RANGE
	ALIASED_LINE_WIDTH_RANGE     = C.GL_ALIASED_LINE_WIDTH_RANGE
	CULL_FACE_MODE               = C.GL_CULL_FACE_MODE
	FRONT_FACE                   = C.GL_FRONT_FACE
	DEPTH_RANGE                  = C.GL_DEPTH_RANGE
	DEPTH_WRITEMASK              = C.GL_DEPTH_WRITEMASK
	DEPTH_CLEAR_VALUE            = C.GL_DEPTH_CLEAR_VALUE
	DEPTH_FUNC                   = C.GL_DEPTH_FUNC
	STENCIL_CLEAR_VALUE          = C.GL_STENCIL_CLEAR_VALUE
	STENCIL_FUNC                 = C.GL_STENCIL_FUNC
	STENCIL_FAIL                 = C.GL_STENCIL_FAIL
	STENCIL_PASS_DEPTH_FAIL      = C.GL_STENCIL_PASS_DEPTH_FAIL
	STENCIL_PASS_DEPTH_PASS      = C.GL_STENCIL_PASS_DEPTH_PASS
	STENCIL_REF                  = C.GL_STENCIL_REF
	STENCIL_VALUE_MASK           = C.GL_STENCIL_VALUE_MASK
	STENCIL_WRITEMASK            = C.GL_STENCIL_WRITEMASK
	STENCIL_BACK_FUNC            = C.GL_STENCIL_BACK_FUNC
	STENCIL_BACK_FAIL            = C.GL_STENCIL_BACK_FAIL
	STENCIL_BACK_PASS_DEPTH_FAIL = C.GL_STENCIL_BACK_PASS_DEPTH_FAIL
	STENCIL_BACK_PASS_DEPTH_PASS = C.GL_STENCIL_BACK_PASS_DEPTH_PASS
	STENCIL_BACK_REF             = C.GL_STENCIL_BACK_REF
	STENCIL_BACK_VALUE_MASK      = C.GL_STENCIL_BACK_VALUE_MASK
	STENCIL_BACK_WRITEMASK       = C.GL_STENCIL_BACK_WRITEMASK
	VIEWPORT                     = C.GL_VIEWPORT
	SCISSOR_BOX                  = C.GL_SCISSOR_BOX
	COLOR_CLEAR_VALUE            = C.GL_COLOR_CLEAR_VALUE
	COLOR_WRITEMASK              = C.GL_COLOR_WRITEMASK
	UNPACK_ALIGNMENT             = C.GL_UNPACK_ALIGNMENT
	PACK_ALIGNMENT               = C.GL_PACK_ALIGNMENT
	MAX_TEXTURE_SIZE             = C.GL_MAX_TEXTURE_SIZE
	MAX_VIEWPORT_DIMS            = C.GL_MAX_VIEWPORT_DIMS
	SUBPIXEL_BITS                = C.GL_SUBPIXEL_BITS
	RED_BITS                     = C.GL_RED_BITS
	GREEN_BITS                   = C.GL_GREEN_BITS
	BLUE_BITS                    = C.GL_BLUE_BITS
	ALPHA_BITS                   = C.GL_ALPHA_BITS
	DEPTH_BITS                   = C.GL_DEPTH_BITS
	STENCIL_BITS                 = C.GL_STENCIL_BITS
	POLYGON_OFFSET_UNITS         = C.GL_POLYGON_OFFSET_UNITS
	POLYGON_OFFSET_FACTOR        = C.GL_POLYGON_OFFSET_FACTOR
	TEXTURE_BINDING_2D           = C.GL_TEXTURE_BINDING_2D
	SAMPLE_BUFFERS               = C.GL_SAMPLE_BUFFERS
	SAMPLES                      = C.GL_SAMPLES
	SAMPLE_COVERAGE_VALUE        = C.GL_SAMPLE_COVERAGE_VALUE
	SAMPLE_COVERAGE_INVERT       = C.GL_SAMPLE_COVERAGE_INVERT

	NUM_COMPRESSED_TEXTURE_FORMATS = C.GL_NUM_COMPRESSED_TEXTURE_FORMATS
	COMPRESSED_TEXTURE_FORMATS     = C.GL_COMPRESSED_TEXTURE_FORMATS

	DONT_CARE = C.GL_DONT_CARE
	FASTEST   = C.GL_FASTEST
	NICEST    = C.GL_NICEST

	GENERATE_MIPMAP_HINT = C.GL_GENERATE_MIPMAP_HINT

	FRAGMENT_SHADER                  = C.GL_FRAGMENT_SHADER
	VERTEX_SHADER                    = C.GL_VERTEX_SHADER
	MAX_VERTEX_ATTRIBS               = C.GL_MAX_VERTEX_ATTRIBS
	MAX_VERTEX_UNIFORM_VECTORS       = C.GL_MAX_VERTEX_UNIFORM_VECTORS
	MAX_VARYING_VECTORS              = C.GL_MAX_VARYING_VECTORS
	MAX_COMBINED_TEXTURE_IMAGE_UNITS = C.GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS
	MAX_VERTEX_TEXTURE_IMAGE_UNITS   = C.GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS
	MAX_TEXTURE_IMAGE_UNITS          = C.GL_MAX_TEXTURE_IMAGE_UNITS
	MAX_FRAGMENT_UNIFORM_VECTORS     = C.GL_MAX_FRAGMENT_UNIFORM_VECTORS
	SHADER_TYPE                      = C.GL_SHADER_TYPE
	DELETE_STATUS                    = C.GL_DELETE_STATUS
	LINK_STATUS                      = C.GL_LINK_STATUS
	VALIDATE_STATUS                  = C.GL_VALIDATE_STATUS
	ATTACHED_SHADERS                 = C.GL_ATTACHED_SHADERS
	ACTIVE_UNIFORMS                  = C.GL_ACTIVE_UNIFORMS
	ACTIVE_UNIFORM_MAX_LENGTH        = C.GL_ACTIVE_UNIFORM_MAX_LENGTH
	ACTIVE_ATTRIBUTES                = C.GL_ACTIVE_ATTRIBUTES
	ACTIVE_ATTRIBUTE_MAX_LENGTH      = C.GL_ACTIVE_ATTRIBUTE_MAX_LENGTH
	SHADING_LANGUAGE_VERSION         = C.GL_SHADING_LANGUAGE_VERSION
	CURRENT_PROGRAM                  = C.GL_CURRENT_PROGRAM

	NEVER    = C.GL_NEVER
	LESS     = C.GL_LESS
	EQUAL    = C.GL_EQUAL
	LEQUAL   = C.GL_LEQUAL
	GREATER  = C.GL_GREATER
	NOTEQUAL = C.GL_NOTEQUAL
	GEQUAL   = C.GL_GEQUAL
	ALWAYS   = C.GL_ALWAYS

	KEEP      = C.GL_KEEP
	REPLACE   = C.GL_REPLACE
	INCR      = C.GL_INCR
	DECR      = C.GL_DECR
	INVERT    = C.GL_INVERT
	INCR_WRAP = C.GL_INCR_WRAP
	DECR_WRAP = C.GL_DECR_WRAP

	VERTEX_ATTRIB_ARRAY_ENABLED        = C.GL_VERTEX_ATTRIB_ARRAY_ENABLED
	VERTEX_ATTRIB_ARRAY_SIZE           = C.GL_VERTEX_ATTRIB_ARRAY_SIZE
	VERTEX_ATTRIB_ARRAY_STRIDE         = C.GL_VERTEX_ATTRIB_ARRAY_STRIDE
	VERTEX_ATTRIB_ARRAY_TYPE           = C.GL_VERTEX_ATTRIB_ARRAY_TYPE
	VERTEX_ATTRIB_ARRAY_NORMALIZED     = C.GL_VERTEX_ATTRIB_ARRAY_NORMALIZED
	VERTEX_ATTRIB_ARRAY_POINTER        = C.GL_VERTEX_ATTRIB_ARRAY_POINTER
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING = C.GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING

	IMPLEMENTATION_COLOR_READ_TYPE   = C.GL_IMPLEMENTATION_COLOR_READ_TYPE
	IMPLEMENTATION_COLOR_READ_FORMAT = C.GL_IMPLEMENTATION_COLOR_READ_FORMAT

	COMPILE_STATUS       = C.GL_COMPILE_STATUS
	INFO_LOG_LENGTH      = C.GL_INFO_LOG_LENGTH
	SHADER_SOURCE_LENGTH = C.GL_SHADER_SOURCE_LENGTH
	SHADER_COMPILER      = C.GL_SHADER_COMPILER

	SHADER_BINARY_FORMATS     = C.GL_SHADER_BINARY_FORMATS
	NUM_SHADER_BINARY_FORMATS = C.GL_NUM_SHADER_BINARY_FORMATS

	LOW_FLOAT    = C.GL_LOW_FLOAT
	MEDIUM_FLOAT = C.GL_MEDIUM_FLOAT
	HIGH_FLOAT   = C.GL_HIGH_FLOAT
	LOW_INT      = C.GL_LOW_INT
	MEDIUM_INT   = C.GL_MEDIUM_INT
	HIGH_INT     = C.GL_HIGH_INT

	FRAMEBUFFER  = C.GL_FRAMEBUFFER
	RENDERBUFFER = C.GL_RENDERBUFFER

	RGBA4             = C.GL_RGBA4
	RGB5_A1           = C.GL_RGB5_A1
	RGB565            = C.GL_RGB565
	DEPTH_COMPONENT16 = C.GL_DEPTH_COMPONENT16
	STENCIL_INDEX8    = C.GL_STENCIL_INDEX8

	RENDERBUFFER_WIDTH           = C.GL_RENDERBUFFER_WIDTH
	RENDERBUFFER_HEIGHT          = C.GL_RENDERBUFFER_HEIGHT
	RENDERBUFFER_INTERNAL_FORMAT = C.GL_RENDERBUFFER_INTERNAL_FORMAT
	RENDERBUFFER_RED_SIZE        = C.GL_RENDERBUFFER_RED_SIZE
	RENDERBUFFER_GREEN_SIZE      = C.GL_RENDERBUFFER_GREEN_SIZE
	RENDERBUFFER_BLUE_SIZE       = C.GL_RENDERBUFFER_BLUE_SIZE
	RENDERBUFFER_ALPHA_SIZE      = C.GL_RENDERBUFFER_ALPHA_SIZE
	RENDERBUFFER_DEPTH_SIZE      = C.GL_RENDERBUFFER_DEPTH_SIZE
	RENDERBUFFER_STENCIL_SIZE    = C.GL_RENDERBUFFER_STENCIL_SIZE

	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE           = C.GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME           = C.GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL         = C.GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = C.GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE

	COLOR_ATTACHMENT0  = C.GL_COLOR_ATTACHMENT0
	DEPTH_ATTACHMENT   = C.GL_DEPTH_ATTACHMENT
	STENCIL_ATTACHMENT = C.GL_STENCIL_ATTACHMENT

	NONE = C.GL_NONE

	FRAMEBUFFER_COMPLETE                      = C.GL_FRAMEBUFFER_COMPLETE
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT         = C.GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT = C.GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS         = C.GL_FRAMEBUFFER_INCOMPLETE_DIMENSIONS
	FRAMEBUFFER_UNSUPPORTED                   = C.GL_FRAMEBUFFER_UNSUPPORTED

	FRAMEBUFFER_BINDING   = C.GL_FRAMEBUFFER_BINDING
	RENDERBUFFER_BINDING  = C.GL_RENDERBUFFER_BINDING
	MAX_RENDERBUFFER_SIZE = C.GL_MAX_RENDERBUFFER_SIZE

	INVALID_FRAMEBUFFER_OPERATION = C.GL_INVALID_FRAMEBUFFER_OPERATION

	FLOAT_VEC2   = C.GL_FLOAT_VEC2
	FLOAT_VEC3   = C.GL_FLOAT_VEC3
	FLOAT_VEC4   = C.GL_FLOAT_VEC4
	INT_VEC2     = C.GL_INT_VEC2
	INT_VEC3     = C.GL_INT_VEC3
	INT_VEC4     = C.GL_INT_VEC4
	BOOL         = C.GL_BOOL
	BOOL_VEC2    = C.GL_BOOL_VEC2
	BOOL_VEC3    = C.GL_BOOL_VEC3
	BOOL_VEC4    = C.GL_BOOL_VEC4
	FLOAT_MAT2   = C.GL_FLOAT_MAT2
	FLOAT_MAT3   = C.GL_FLOAT_MAT3
	FLOAT_MAT4   = C.GL_FLOAT_MAT4
	SAMPLER_2D   = C.GL_SAMPLER_2D
	SAMPLER_CUBE = C.GL_SAMPLER_CUBE

	VENDOR     InfoType = C.GL_VENDOR
	RENDERER            = C.GL_RENDERER
	VERSION             = C.GL_VERSION
	EXTENSIONS          = C.GL_EXTENSIONS

	POINTS         DrawMode = C.GL_POINTS
	LINES                   = C.GL_LINES
	LINE_LOOP               = C.GL_LINE_LOOP
	LINE_STRIP              = C.GL_LINE_STRIP
	TRIANGLES               = C.GL_TRIANGLES
	TRIANGLE_STRIP          = C.GL_TRIANGLE_STRIP
	TRIANGLE_FAN            = C.GL_TRIANGLE_FAN

	BYTE           DataType = C.GL_BYTE
	UNSIGNED_BYTE           = C.GL_UNSIGNED_BYTE
	SHORT                   = C.GL_SHORT
	UNSIGNED_SHORT          = C.GL_UNSIGNED_SHORT
	INT                     = C.GL_INT
	UNSIGNED_INT            = C.GL_UNSIGNED_INT
	FLOAT                   = C.GL_FLOAT
	FIXED                   = C.GL_FIXED

	UNSIGNED_SHORT_4_4_4_4 = C.GL_UNSIGNED_SHORT_4_4_4_4
	UNSIGNED_SHORT_5_5_5_1 = C.GL_UNSIGNED_SHORT_5_5_5_1
	UNSIGNED_SHORT_5_6_5   = C.GL_UNSIGNED_SHORT_5_6_5

	NO_ERROR          Error = C.GL_NO_ERROR
	INVALID_ENUM            = C.GL_INVALID_ENUM
	INVALID_VALUE           = C.GL_INVALID_VALUE
	INVALID_OPERATION       = C.GL_INVALID_OPERATION
	OUT_OF_MEMORY           = C.GL_OUT_OF_MEMORY
)

func GetError() Error {
	return Error(C.glGetError())
}

func GetString(name InfoType) string {
	return C.GoString((*C.char)(unsafe.Pointer(C.glGetString(C.GLenum(name)))))
}

func Clear(mask int) {
	C.glClear(C.GLbitfield(mask))
}

func ClearColor(red, green, blue, alpha float32) {
	C.glClearColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
}

func ClearDepthf(depth float32) {
	C.glClearDepthf(C.GLclampf(depth))
}

func ClearStencil(s int32) {
	C.glClearStencil(C.GLint(s))
}

func DrawArrays(mode DrawMode, first, count int) {
	C.glDrawArrays(C.GLenum(mode), C.GLint(first), C.GLsizei(count))
}

func DrawElementsShort(mode DrawMode, count int, indices []uint16) {
	C.glDrawElements(C.GLenum(mode), C.GLsizei(count), C.GLenum(C.GL_UNSIGNED_SHORT), unsafe.Pointer(&indices[0]))
}

func DrawElementsByte(mode DrawMode, count int, indices []uint8) {
	C.glDrawElements(C.GLenum(mode), C.GLsizei(count), C.GLenum(C.GL_UNSIGNED_BYTE), unsafe.Pointer(&indices[0]))
}

// func BindFramebuffer(target int, framebuffer uint) {
// 	C.glBindFramebuffer(GLenum target, GLuint framebuffer)
// }

// func BindRenderbuffer(target int, renderbuffer uint) {
// 	C.glBindRenderbuffer(GLenum target, GLuint renderbuffer)
// }

// func BlendColor(red, green, blue, alpha float32) {
// 	C.glBlendColor(GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha)
// }

// func BlendEquation(mode int) {
// 	C.glBlendEquation( GLenum mode )
// }

// func BlendEquationSeparate(modeRGB int, modeAlpha int) {
// 	C.glBlendEquationSeparate(GLenum modeRGB, GLenum modeAlpha)
// }

// func BlendFunc(sfactor int, dfactor int) {
// 	C.glBlendFunc(GLenum sfactor, GLenum dfactor)
// }

// func BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha int) {
// 	C.glBlendFuncSeparate(GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha)
// }

// func CheckFramebufferStatus(target int) int {
// 	return C.glCheckFramebufferStatus(target)
// }

// func ColorMask(red, green, blue, alpha bool) {
// 	C.glColorMask(GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha)
// }

// func CullFace(mode int) {
// 	C.glCullFace(GLenum mode)
// }

// func DeleteFramebuffers(GLsizei n, framebuffers []uint) {
// 	C.glDeleteFramebuffers(GLsizei n, const GLuint* framebuffers)
// }

// func DeleteRenderbuffers(GLsizei n, renderbuffers []uint) {
// 	C.glDeleteRenderbuffers(GLsizei n, const GLuint* renderbuffers)
// }

// func DepthFunc(func int) {
// 	C.glDepthFunc(GLenum func)
// }

// func DepthMask(GLboolean flag) {
// 	C.glDepthMask(GLboolean flag)
// }

// func DepthRangef(GLclampf zNear, zFar float32) {
// 	C.glDepthRangef(GLclampf zNear, GLclampf zFar)
// }

// func Disable(cap int) {
// 	C.glDisable(GLenum cap)
// }

// func Enable(cap int) {
// 	C.glEnable(GLenum cap)
// }

// func Finish(void) {
// 	C.glFinish(void)
// }

// func Flush(void) {
// 	C.glFlush(void)
// }

// func FramebufferRenderbuffer(target int, attachment int, renderbuffertarget int, renderbuffer uint) {
// 	C.glFramebufferRenderbuffer(GLenum target, GLenum attachment, GLenum renderbuffertarget, GLuint renderbuffer)
// }

// func FramebufferTexture2D(target int, attachment int, textarget int, texture uint, level int) {
// 	C.glFramebufferTexture2D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level)
// }

// func FrontFace(mode int) {
// 	C.glFrontFace(GLenum mode)
// }

// func GenFramebuffers(GLsizei n) []uint {
// 	C.glGenFramebuffers(GLsizei n, GLuint* framebuffers)
// }

// func GenRenderbuffers(GLsizei n) []uint {
// 	C.glGenRenderbuffers(GLsizei n, GLuint* renderbuffers)
// }

// func GetBooleanv(pname int, GLboolean* params) {
// 	C.glGetBooleanv(GLenum pname, GLboolean* params)
// }

// func GetFloatv(pname int, GLfloat* params) {
// 	C.glGetFloatv(GLenum pname, GLfloat* params)
// }

// func GetFramebufferAttachmentParameteriv(target int, attachment int, pname int,  params int) {
// 	C.glGetFramebufferAttachmentParameteriv(GLenum target, GLenum attachment, GLenum pname, GLint* params)
// }

// func GetIntegerv(pname int,  params int) {
// 	C.glGetIntegerv(GLenum pname, GLint* params)
// }

// func GetRenderbufferParameteriv(target int, pname int,  params int) {
// 	C.glGetRenderbufferParameteriv(GLenum target, GLenum pname, GLint* params)
// }

// func GetVertexAttribfv(index uint, pname int, GLfloat* params) {
// 	C.glGetVertexAttribfv(GLuint index, GLenum pname, GLfloat* params)
// }

// func GetVertexAttribiv(index uint, pname int,  params int) {
// 	C.glGetVertexAttribiv(GLuint index, GLenum pname, GLint* params)
// }

// func GetVertexAttribPointerv(index uint, pname int, GLvoid** pointer) {
// 	C.glGetVertexAttribPointerv(GLuint index, GLenum pname, GLvoid** pointer)
// }

// func Hint(target int, mode int) {
// 	C.glHint(GLenum target, GLenum mode)
// }

// GL_APICALL GLboolean    GL_APIENTRY glIsBuffer (GLuint buffer);
// GL_APICALL GLboolean    GL_APIENTRY glIsEnabled (GLenum cap);
// GL_APICALL GLboolean    GL_APIENTRY glIsFramebuffer (GLuint framebuffer);
// GL_APICALL GLboolean    GL_APIENTRY glIsProgram (GLuint program);
// GL_APICALL GLboolean    GL_APIENTRY glIsRenderbuffer (GLuint renderbuffer);
// GL_APICALL GLboolean    GL_APIENTRY glIsShader (GLuint shader);
// GL_APICALL GLboolean    GL_APIENTRY glIsTexture (GLuint texture);
// func LineWidth(GLfloat width) {
// 	C.glLineWidth(GLfloat width)
// }

// func PixelStorei(pname int, param int) {
// 	C.glPixelStorei(GLenum pname, GLint param)
// }

// func PolygonOffset(GLfloat factor, GLfloat units) {
// 	C.glPolygonOffset(GLfloat factor, GLfloat units)
// }

// func ReadPixels(GLint x, y int, GLsizei width, GLsizei height, format int, type int, GLvoid* pixels) {
// 	C.glReadPixels(GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels)
// }

// func RenderbufferStorage(target int, internalformat int, GLsizei width, GLsizei height) {
// 	C.glRenderbufferStorage(GLenum target, GLenum internalformat, GLsizei width, GLsizei height)
// }

// func SampleCoverage(value float32, GLboolean invert) {
// 	C.glSampleCoverage(GLclampf value, GLboolean invert)
// }

// func Scissor(GLint x, y int, GLsizei width, GLsizei height) {
// 	C.glScissor(GLint x, GLint y, GLsizei width, GLsizei height)
// }

// func StencilFunc(func int, ref int, mask uint) {
// 	C.glStencilFunc(GLenum func, ref int, mask uint)
// }

// func StencilFuncSeparate(face int, func int, ref int, mask uint) {
// 	C.glStencilFuncSeparate(GLenum face, GLenum func, ref int, mask uint)
// }

// func StencilMask(mask uint) {
// 	C.glStencilMask(GLuint mask)
// }

// func StencilMaskSeparate(face int, mask uint) {
// 	C.glStencilMaskSeparate(GLenum face, GLuint mask)
// }

// func StencilOp(fail int, zfail int, zpass int) {
// 	C.glStencilOp(GLenum fail, GLenum zfail, GLenum zpass)
// }

// func StencilOpSeparate(face int, fail int, zfail int, zpass int) {
// 	C.glStencilOpSeparate(GLenum face, GLenum fail, GLenum zfail, GLenum zpass)
// }

// func Viewport(GLint x, y int, GLsizei width, GLsizei height) {
// 	C.glViewport(GLint x, GLint y, GLsizei width, GLsizei height)
// }
