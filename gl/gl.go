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
// #cgo LDFLAGS: -lGLESv2
import "C"

import (
	"fmt"
	"image"
	"unsafe"
)

type Error C.GLenum
type DataType C.GLenum
type InternalFormat C.GLenum
type InfoType C.GLenum
type DrawMode C.GLenum
type BlendMode C.GLenum
type BlendFactor C.GLenum
type Capability C.GLenum
type StencilMode C.GLenum
type TestFunc C.GLenum
type StencilOperation C.GLenum
type Face C.GLenum
type Facing C.GLenum
type HintType C.GLenum
type HintValue C.GLenum
type Packing C.GLenum

const (
	ES_VERSION_2_0 = C.GL_ES_VERSION_2_0

	DEPTH_BUFFER_BIT   = C.GL_DEPTH_BUFFER_BIT
	STENCIL_BUFFER_BIT = C.GL_STENCIL_BUFFER_BIT
	COLOR_BUFFER_BIT   = C.GL_COLOR_BUFFER_BIT

	FALSE = C.GL_FALSE
	TRUE  = C.GL_TRUE

	CURRENT_VERTEX_ATTRIB = C.GL_CURRENT_VERTEX_ATTRIB

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

	LOW_FLOAT    = C.GL_LOW_FLOAT
	MEDIUM_FLOAT = C.GL_MEDIUM_FLOAT
	HIGH_FLOAT   = C.GL_HIGH_FLOAT
	LOW_INT      = C.GL_LOW_INT
	MEDIUM_INT   = C.GL_MEDIUM_INT
	HIGH_INT     = C.GL_HIGH_INT

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

	BLEND_EQUATION       = C.GL_BLEND_EQUATION
	BLEND_EQUATION_RGB   = C.GL_BLEND_EQUATION_RGB
	BLEND_EQUATION_ALPHA = C.GL_BLEND_EQUATION_ALPHA

	BLEND_DST_RGB   = C.GL_BLEND_DST_RGB
	BLEND_SRC_RGB   = C.GL_BLEND_SRC_RGB
	BLEND_DST_ALPHA = C.GL_BLEND_DST_ALPHA
	BLEND_SRC_ALPHA = C.GL_BLEND_SRC_ALPHA
	BLEND_COLOR     = C.GL_BLEND_COLOR

	VENDOR     InfoType = C.GL_VENDOR
	RENDERER            = C.GL_RENDERER
	VERSION             = C.GL_VERSION
	EXTENSIONS          = C.GL_EXTENSIONS

	GENERATE_MIPMAP_HINT HintType = C.GL_GENERATE_MIPMAP_HINT

	DONT_CARE HintValue = C.GL_DONT_CARE
	FASTEST   HintValue = C.GL_FASTEST
	NICEST    HintValue = C.GL_NICEST

	CULL_FACE                Capability = C.GL_CULL_FACE
	BLEND                    Capability = C.GL_BLEND
	DITHER                   Capability = C.GL_DITHER
	STENCIL_TEST             Capability = C.GL_STENCIL_TEST
	DEPTH_TEST               Capability = C.GL_DEPTH_TEST
	SCISSOR_TEST             Capability = C.GL_SCISSOR_TEST
	POLYGON_OFFSET_FILL      Capability = C.GL_POLYGON_OFFSET_FILL
	SAMPLE_ALPHA_TO_COVERAGE Capability = C.GL_SAMPLE_ALPHA_TO_COVERAGE
	SAMPLE_COVERAGE          Capability = C.GL_SAMPLE_COVERAGE

	UNPACK_ALIGNMENT Packing = C.GL_UNPACK_ALIGNMENT
	PACK_ALIGNMENT   Packing = C.GL_PACK_ALIGNMENT

	POINTS         DrawMode = C.GL_POINTS
	LINES          DrawMode = C.GL_LINES
	LINE_LOOP      DrawMode = C.GL_LINE_LOOP
	LINE_STRIP     DrawMode = C.GL_LINE_STRIP
	TRIANGLES      DrawMode = C.GL_TRIANGLES
	TRIANGLE_STRIP DrawMode = C.GL_TRIANGLE_STRIP
	TRIANGLE_FAN   DrawMode = C.GL_TRIANGLE_FAN

	FUNC_ADD              BlendMode = C.GL_FUNC_ADD
	FUNC_SUBTRACT         BlendMode = C.GL_FUNC_SUBTRACT
	FUNC_REVERSE_SUBTRACT BlendMode = C.GL_FUNC_REVERSE_SUBTRACT

	ZERO                     BlendFactor = C.GL_ZERO
	ONE                      BlendFactor = C.GL_ONE
	SRC_COLOR                BlendFactor = C.GL_SRC_COLOR
	ONE_MINUS_SRC_COLOR      BlendFactor = C.GL_ONE_MINUS_SRC_COLOR
	SRC_ALPHA                BlendFactor = C.GL_SRC_ALPHA
	ONE_MINUS_SRC_ALPHA      BlendFactor = C.GL_ONE_MINUS_SRC_ALPHA
	DST_ALPHA                BlendFactor = C.GL_DST_ALPHA
	ONE_MINUS_DST_ALPHA      BlendFactor = C.GL_ONE_MINUS_DST_ALPHA
	DST_COLOR                BlendFactor = C.GL_DST_COLOR
	ONE_MINUS_DST_COLOR      BlendFactor = C.GL_ONE_MINUS_DST_COLOR
	SRC_ALPHA_SATURATE       BlendFactor = C.GL_SRC_ALPHA_SATURATE
	CONSTANT_COLOR           BlendFactor = C.GL_CONSTANT_COLOR
	ONE_MINUS_CONSTANT_COLOR BlendFactor = C.GL_ONE_MINUS_CONSTANT_COLOR
	CONSTANT_ALPHA           BlendFactor = C.GL_CONSTANT_ALPHA
	ONE_MINUS_CONSTANT_ALPHA BlendFactor = C.GL_ONE_MINUS_CONSTANT_ALPHA

	NEVER    TestFunc = C.GL_NEVER
	LESS     TestFunc = C.GL_LESS
	EQUAL    TestFunc = C.GL_EQUAL
	LEQUAL   TestFunc = C.GL_LEQUAL
	GREATER  TestFunc = C.GL_GREATER
	NOTEQUAL TestFunc = C.GL_NOTEQUAL
	GEQUAL   TestFunc = C.GL_GEQUAL
	ALWAYS   TestFunc = C.GL_ALWAYS

	CW  Facing = C.GL_CW
	CCW Facing = C.GL_CCW

	FRONT          Face = C.GL_FRONT
	BACK           Face = C.GL_BACK
	FRONT_AND_BACK Face = C.GL_FRONT_AND_BACK

	RESET     StencilOperation = C.GL_ZERO
	KEEP      StencilOperation = C.GL_KEEP
	REPLACE   StencilOperation = C.GL_REPLACE
	INCR      StencilOperation = C.GL_INCR
	DECR      StencilOperation = C.GL_DECR
	INVERT    StencilOperation = C.GL_INVERT
	INCR_WRAP StencilOperation = C.GL_INCR_WRAP
	DECR_WRAP StencilOperation = C.GL_DECR_WRAP

	BYTE                           DataType = C.GL_BYTE
	UNSIGNED_BYTE                  DataType = C.GL_UNSIGNED_BYTE
	SHORT                          DataType = C.GL_SHORT
	UNSIGNED_SHORT                 DataType = C.GL_UNSIGNED_SHORT
	INT                            DataType = C.GL_INT
	UNSIGNED_INT                   DataType = C.GL_UNSIGNED_INT
	FLOAT                          DataType = C.GL_FLOAT
	FIXED                          DataType = C.GL_FIXED
	UNSIGNED_SHORT_4_4_4_4         DataType = C.GL_UNSIGNED_SHORT_4_4_4_4
	UNSIGNED_SHORT_5_5_5_1         DataType = C.GL_UNSIGNED_SHORT_5_5_5_1
	UNSIGNED_SHORT_5_6_5           DataType = C.GL_UNSIGNED_SHORT_5_6_5
	HALF_FLOAT                     DataType = C.GL_HALF_FLOAT
	UNSIGNED_INT_5_9_9_9_REV       DataType = C.GL_UNSIGNED_INT_5_9_9_9_REV
	UNSIGNED_INT_10F_11F_11F_REV   DataType = C.GL_UNSIGNED_INT_10F_11F_11F_REV
	UNSIGNED_INT_2_10_10_10_REV    DataType = C.GL_UNSIGNED_INT_2_10_10_10_REV
	UNSIGNED_INT_24_8              DataType = C.GL_UNSIGNED_INT_24_8
	FLOAT_32_UNSIGNED_INT_24_8_REV DataType = C.GL_FLOAT_32_UNSIGNED_INT_24_8_REV

	R8                 InternalFormat = C.GL_R8
	RG8                InternalFormat = C.GL_RG8
	RGB8               InternalFormat = C.GL_RGB8
	RGBA8              InternalFormat = C.GL_RGBA8
	SRGB               InternalFormat = C.GL_SRGB
	SRGB8              InternalFormat = C.GL_SRGB8
	SRGB8_ALPHA8       InternalFormat = C.GL_SRGB8_ALPHA8
	R8I                InternalFormat = C.GL_R8I
	RG8I               InternalFormat = C.GL_RG8I
	RGB8I              InternalFormat = C.GL_RGB8I
	RGBA8I             InternalFormat = C.GL_RGBA8I
	R8UI               InternalFormat = C.GL_R8UI
	RG8UI              InternalFormat = C.GL_RG8UI
	RGB8UI             InternalFormat = C.GL_RGB8UI
	RGBA8UI            InternalFormat = C.GL_RGBA8UI
	R16I               InternalFormat = C.GL_R16I
	RG16I              InternalFormat = C.GL_RG16I
	RGB16I             InternalFormat = C.GL_RGB16I
	RGBA16I            InternalFormat = C.GL_RGBA16I
	R16UI              InternalFormat = C.GL_R16UI
	RG16UI             InternalFormat = C.GL_RG16UI
	RGB16UI            InternalFormat = C.GL_RGB16UI
	RGBA16UI           InternalFormat = C.GL_RGBA16UI
	R16F               InternalFormat = C.GL_R16F
	RG16F              InternalFormat = C.GL_RG16F
	RGB16F             InternalFormat = C.GL_RGB16F
	RGBA16F            InternalFormat = C.GL_RGBA16F
	R32I               InternalFormat = C.GL_R32I
	RG32I              InternalFormat = C.GL_RG32I
	RGB32I             InternalFormat = C.GL_RGB32I
	RGBA32I            InternalFormat = C.GL_RGBA32I
	R32UI              InternalFormat = C.GL_R32UI
	RG32UI             InternalFormat = C.GL_RG32UI
	RGB32UI            InternalFormat = C.GL_RGB32UI
	RGBA32UI           InternalFormat = C.GL_RGBA32UI
	R32F               InternalFormat = C.GL_R32F
	RG32F              InternalFormat = C.GL_RG32F
	RGB32F             InternalFormat = C.GL_RGB32F
	RGBA32F            InternalFormat = C.GL_RGBA32F
	R11F_G11F_B10F     InternalFormat = C.GL_R11F_G11F_B10F
	RGBA4              InternalFormat = C.GL_RGBA4
	RGB565             InternalFormat = C.GL_RGB565
	RGB5_A1            InternalFormat = C.GL_RGB5_A1
	RGB9_E5            InternalFormat = C.GL_RGB9_E5
	RGB10_A2           InternalFormat = C.GL_RGB10_A2
	RGB10_A2UI         InternalFormat = C.GL_RGB10_A2UI
	DEPTH_COMPONENT16  InternalFormat = C.GL_DEPTH_COMPONENT16
	DEPTH_COMPONENT24  InternalFormat = C.GL_DEPTH_COMPONENT24
	DEPTH_COMPONENT32F InternalFormat = C.GL_DEPTH_COMPONENT32F
	DEPTH24_STENCIL8   InternalFormat = C.GL_DEPTH24_STENCIL8
	DEPTH32F_STENCIL8  InternalFormat = C.GL_DEPTH32F_STENCIL8
	STENCIL_INDEX8     InternalFormat = C.GL_STENCIL_INDEX8

	NO_ERROR                      Error = C.GL_NO_ERROR
	INVALID_ENUM                  Error = C.GL_INVALID_ENUM
	INVALID_VALUE                 Error = C.GL_INVALID_VALUE
	INVALID_OPERATION             Error = C.GL_INVALID_OPERATION
	OUT_OF_MEMORY                 Error = C.GL_OUT_OF_MEMORY
	INVALID_FRAMEBUFFER_OPERATION Error = C.GL_INVALID_FRAMEBUFFER_OPERATION
)

func GetError() Error {
	return Error(C.glGetError())
}

func GetString(name InfoType) string {
	return C.GoString((*C.char)(unsafe.Pointer(C.glGetString(C.GLenum(name)))))
}

func Hint(hint HintType, value HintValue) {
	C.glHint(C.GLenum(hint), C.GLenum(value))
}

func Finish() {
	C.glFinish()
}

func Flush() {
	C.glFlush()
}

func Enable(c Capability) {
	C.glEnable(C.GLenum(c))
}

func Disable(c Capability) {
	C.glDisable(C.GLenum(c))
}

func Scissor(x, y, width, height int) {
	C.glScissor(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

func Viewport(x, y, width, height int) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

func GetViewport() (int, int, int, int) {
	data := [4]C.GLint{0, 0, 0, 0}
	C.glGetIntegerv(C.GL_VIEWPORT, &data[0])
	return int(data[0]), int(data[1]), int(data[2]), int(data[3])
}

func Clear(mask int) {
	C.glClear(C.GLbitfield(mask))
}

func ClearColor(red, green, blue, alpha float32) {
	C.glClearColor(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
}

func ClearDepthf(depth float32) {
	C.glClearDepthf(C.GLfloat(depth))
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

func DrawArraysInstanced(mode DrawMode, first, count, instanceCount int) {
	C.glDrawArraysInstanced(C.GLenum(mode), C.GLint(first), C.GLsizei(count), C.GLsizei(instanceCount))
}

func DrawElementsInstancedShort(mode DrawMode, count int, indices []uint16, instanceCount int) {
	C.glDrawElementsInstanced(C.GLenum(mode), C.GLsizei(count), C.GL_UNSIGNED_SHORT, unsafe.Pointer(&indices[0]), C.GLsizei(instanceCount))
}

func BlendColor(red, green, blue, alpha float32) {
	C.glBlendColor(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
}

func BlendEquation(mode BlendMode) {
	C.glBlendEquation(C.GLenum(mode))
}

func BlendEquationSeparate(modeRGB, modeAlpha BlendMode) {
	C.glBlendEquationSeparate(C.GLenum(modeRGB), C.GLenum(modeAlpha))
}

func BlendFunc(sfactor, dfactor BlendFactor) {
	C.glBlendFunc(C.GLenum(sfactor), C.GLenum(dfactor))
}

func BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha BlendFactor) {
	C.glBlendFuncSeparate(C.GLenum(srcRGB), C.GLenum(dstRGB), C.GLenum(srcAlpha), C.GLenum(dstAlpha))
}

func ColorMask(red, green, blue, alpha bool) {
	r, g, b, a := C.GLboolean(C.GL_FALSE), C.GLboolean(C.GL_FALSE), C.GLboolean(C.GL_FALSE), C.GLboolean(C.GL_FALSE)
	if red {
		r = C.GLboolean(C.GL_TRUE)
	}
	if green {
		g = C.GLboolean(C.GL_TRUE)
	}
	if blue {
		b = C.GLboolean(C.GL_TRUE)
	}
	if alpha {
		a = C.GLboolean(C.GL_TRUE)
	}
	C.glColorMask(r, g, b, a)
}

func CullFace(mode Face) {
	C.glCullFace(C.GLenum(mode))
}

func FrontFace(mode Facing) {
	C.glFrontFace(C.GLenum(mode))
}

func DepthFunc(f TestFunc) {
	C.glDepthFunc(C.GLenum(f))
}

func DepthMask(flag bool) {
	b := C.GLboolean(C.GL_FALSE)
	if flag {
		b = C.GLboolean(C.GL_TRUE)
	}
	C.glDepthMask(b)
}

func DepthRangef(zNear, zFar float32) {
	C.glDepthRangef(C.GLfloat(zNear), C.GLfloat(zFar))
}

func StencilFunc(f TestFunc, ref int32, mask uint32) {
	C.glStencilFunc(C.GLenum(f), C.GLint(ref), C.GLuint(mask))
}

func StencilFuncSeparate(face Face, f TestFunc, ref int32, mask uint32) {
	C.glStencilFuncSeparate(C.GLenum(face), C.GLenum(f), C.GLint(ref), C.GLuint(mask))
}

func StencilMask(mask uint32) {
	C.glStencilMask(C.GLuint(mask))
}

func StencilMaskSeparate(face Face, mask uint32) {
	C.glStencilMaskSeparate(C.GLenum(face), C.GLuint(mask))
}

func StencilOp(fail, zfail, zpass StencilOperation) {
	C.glStencilOp(C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
}

func StencilOpSeparate(face Face, fail, zfail, zpass StencilOperation) {
	C.glStencilOpSeparate(C.GLenum(face), C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
}

func LineWidth(width float32) {
	C.glLineWidth(C.GLfloat(width))
}

func PolygonOffset(factor, units float32) {
	C.glPolygonOffset(C.GLfloat(factor), C.GLfloat(units))
}

func SampleCoverage(value float32, invert bool) {
	b := C.GLboolean(C.GL_FALSE)
	if invert {
		b = C.GLboolean(C.GL_TRUE)
	}

	C.glSampleCoverage(C.GLfloat(value), C.GLboolean(b))
}

func PixelStorei(pname Packing, param int) {
	C.glPixelStorei(C.GLenum(pname), C.GLint(param))
}

func ReadRGBA(image *image.NRGBA) {

	alignment := C.GLint(0)
	C.glGetIntegerv(C.GL_PACK_ALIGNMENT, &alignment)

	align := image.Stride % int(alignment) // align: 4 or 0

	// need smaller alignment
	if align > 0 {
		C.glPixelStorei(C.GL_PACK_ALIGNMENT, C.GLint(align))
	}

	C.glReadPixels(C.GLint(image.Rect.Min.X), C.GLint(image.Rect.Min.Y),
		C.GLsizei(image.Rect.Dx()), C.GLsizei(image.Rect.Dy()),
		C.GL_RGBA, C.GL_UNSIGNED_BYTE, unsafe.Pointer(&image.Pix[0]))

	// restore alignment
	if align > 0 {
		C.glPixelStorei(C.GL_PACK_ALIGNMENT, alignment)
	}
}

func ReadAlpha(image *image.Alpha) {

	alignment := C.GLint(0)
	C.glGetIntegerv(C.GL_PACK_ALIGNMENT, &alignment)

	align := C.GLint(1)

	for align < alignment && image.Stride%(int(align)*2) == 0 {
		align *= 2
	}

	// need smaller alignment
	if align < alignment {
		C.glPixelStorei(C.GL_PACK_ALIGNMENT, align)
	}

	C.glReadPixels(C.GLint(image.Rect.Min.X), C.GLint(image.Rect.Min.Y),
		C.GLsizei(image.Rect.Dx()), C.GLsizei(image.Rect.Dy()),
		C.GL_ALPHA, C.GL_UNSIGNED_BYTE, unsafe.Pointer(&image.Pix[0]))

	// restore alignment
	if align < alignment {
		C.glPixelStorei(C.GL_PACK_ALIGNMENT, alignment)
	}
}

func (err Error) String() string {

	switch err {
	case NO_ERROR:
		return "NO_ERROR"
	case INVALID_ENUM:
		return "INVALID_ENUM"
	case INVALID_VALUE:
		return "INVALID_VALUE"
	case INVALID_OPERATION:
		return "INVALID_OPERATION"
	case OUT_OF_MEMORY:
		return "OUT_OF_MEMORY"
	case INVALID_FRAMEBUFFER_OPERATION:
		return "INVALID_FRAMEBUFFER_OPERATION"
	default:
		return fmt.Sprintf("Error(%#x)", err)
	}
}

// func GetBooleanv(pname int, GLboolean* params) {
// 	C.glGetBooleanv(GLenum pname, GLboolean* params)
// }

// func GetFloatv(pname int, GLfloat* params) {
// 	C.glGetFloatv(GLenum pname, GLfloat* params)
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

// GL_APICALL GLboolean    GL_APIENTRY glIsBuffer (GLuint buffer);
// GL_APICALL GLboolean    GL_APIENTRY glIsEnabled (GLenum cap);
// GL_APICALL GLboolean    GL_APIENTRY glIsFramebuffer (GLuint framebuffer);
// GL_APICALL GLboolean    GL_APIENTRY glIsProgram (GLuint program);
// GL_APICALL GLboolean    GL_APIENTRY glIsRenderbuffer (GLuint renderbuffer);
// GL_APICALL GLboolean    GL_APIENTRY glIsShader (GLuint shader);
// GL_APICALL GLboolean    GL_APIENTRY glIsTexture (GLuint texture);
