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
	"image"
	"unsafe"
)

type TextureFormat C.GLenum
type TextureUnit C.GLenum
type Texture C.GLuint
type TextureTarget C.GLenum
type TextureParameter C.GLenum
type TextureParameterValue C.GLint

const (
	DEPTH_COMPONENT TextureFormat = C.GL_DEPTH_COMPONENT
	ALPHA           TextureFormat = C.GL_ALPHA
	RGB             TextureFormat = C.GL_RGB
	RGBA            TextureFormat = C.GL_RGBA
	LUMINANCE       TextureFormat = C.GL_LUMINANCE
	LUMINANCE_ALPHA TextureFormat = C.GL_LUMINANCE_ALPHA
	RED             TextureFormat = C.GL_RED
	RG              TextureFormat = C.GL_RG
	RED_INTEGER     TextureFormat = C.GL_RED_INTEGER
	RG_INTEGER      TextureFormat = C.GL_RG_INTEGER
	RGB_INTEGER     TextureFormat = C.GL_RGB_INTEGER
	RGBA_INTEGER    TextureFormat = C.GL_RGBA_INTEGER

	TEXTURE_MAG_FILTER TextureParameter = C.GL_TEXTURE_MAG_FILTER
	TEXTURE_MIN_FILTER TextureParameter = C.GL_TEXTURE_MIN_FILTER
	TEXTURE_WRAP_S     TextureParameter = C.GL_TEXTURE_WRAP_S
	TEXTURE_WRAP_T     TextureParameter = C.GL_TEXTURE_WRAP_T

	TEXTURE_2D                  TextureTarget = C.GL_TEXTURE_2D
	TEXTURE_CUBE_MAP            TextureTarget = C.GL_TEXTURE_CUBE_MAP
	TEXTURE_CUBE_MAP_POSITIVE_X TextureTarget = C.GL_TEXTURE_CUBE_MAP_POSITIVE_X
	TEXTURE_CUBE_MAP_NEGATIVE_X TextureTarget = C.GL_TEXTURE_CUBE_MAP_NEGATIVE_X
	TEXTURE_CUBE_MAP_POSITIVE_Y TextureTarget = C.GL_TEXTURE_CUBE_MAP_POSITIVE_Y
	TEXTURE_CUBE_MAP_NEGATIVE_Y TextureTarget = C.GL_TEXTURE_CUBE_MAP_NEGATIVE_Y
	TEXTURE_CUBE_MAP_POSITIVE_Z TextureTarget = C.GL_TEXTURE_CUBE_MAP_POSITIVE_Z
	TEXTURE_CUBE_MAP_NEGATIVE_Z TextureTarget = C.GL_TEXTURE_CUBE_MAP_NEGATIVE_Z

	MAX_CUBE_MAP_TEXTURE_SIZE = C.GL_MAX_CUBE_MAP_TEXTURE_SIZE

	TEXTURE0  TextureUnit = C.GL_TEXTURE0
	TEXTURE1  TextureUnit = C.GL_TEXTURE1
	TEXTURE2  TextureUnit = C.GL_TEXTURE2
	TEXTURE3  TextureUnit = C.GL_TEXTURE3
	TEXTURE4  TextureUnit = C.GL_TEXTURE4
	TEXTURE5  TextureUnit = C.GL_TEXTURE5
	TEXTURE6  TextureUnit = C.GL_TEXTURE6
	TEXTURE7  TextureUnit = C.GL_TEXTURE7
	TEXTURE8  TextureUnit = C.GL_TEXTURE8
	TEXTURE9  TextureUnit = C.GL_TEXTURE9
	TEXTURE10 TextureUnit = C.GL_TEXTURE10
	TEXTURE11 TextureUnit = C.GL_TEXTURE11
	TEXTURE12 TextureUnit = C.GL_TEXTURE12
	TEXTURE13 TextureUnit = C.GL_TEXTURE13
	TEXTURE14 TextureUnit = C.GL_TEXTURE14
	TEXTURE15 TextureUnit = C.GL_TEXTURE15
	TEXTURE16 TextureUnit = C.GL_TEXTURE16
	TEXTURE17 TextureUnit = C.GL_TEXTURE17
	TEXTURE18 TextureUnit = C.GL_TEXTURE18
	TEXTURE19 TextureUnit = C.GL_TEXTURE19
	TEXTURE20 TextureUnit = C.GL_TEXTURE20
	TEXTURE21 TextureUnit = C.GL_TEXTURE21
	TEXTURE22 TextureUnit = C.GL_TEXTURE22
	TEXTURE23 TextureUnit = C.GL_TEXTURE23
	TEXTURE24 TextureUnit = C.GL_TEXTURE24
	TEXTURE25 TextureUnit = C.GL_TEXTURE25
	TEXTURE26 TextureUnit = C.GL_TEXTURE26
	TEXTURE27 TextureUnit = C.GL_TEXTURE27
	TEXTURE28 TextureUnit = C.GL_TEXTURE28
	TEXTURE29 TextureUnit = C.GL_TEXTURE29
	TEXTURE30 TextureUnit = C.GL_TEXTURE30
	TEXTURE31 TextureUnit = C.GL_TEXTURE31

	NEAREST                TextureParameterValue = C.GL_NEAREST
	LINEAR                 TextureParameterValue = C.GL_LINEAR
	NEAREST_MIPMAP_NEAREST TextureParameterValue = C.GL_NEAREST_MIPMAP_NEAREST
	LINEAR_MIPMAP_NEAREST  TextureParameterValue = C.GL_LINEAR_MIPMAP_NEAREST
	NEAREST_MIPMAP_LINEAR  TextureParameterValue = C.GL_NEAREST_MIPMAP_LINEAR
	LINEAR_MIPMAP_LINEAR   TextureParameterValue = C.GL_LINEAR_MIPMAP_LINEAR
	REPEAT                 TextureParameterValue = C.GL_REPEAT
	CLAMP_TO_EDGE          TextureParameterValue = C.GL_CLAMP_TO_EDGE
	MIRRORED_REPEAT        TextureParameterValue = C.GL_MIRRORED_REPEAT
)

func ActiveTexture(texture TextureUnit) {
	C.glActiveTexture(C.GLenum(texture))
}

func GetActiveTexture() TextureUnit {
	texture := C.GLint(0)
	C.glGetIntegerv(C.GL_ACTIVE_TEXTURE, &texture)
	return TextureUnit(texture)
}

func BindTexture(target TextureTarget, texture Texture) {
	C.glBindTexture(C.GLenum(target), C.GLuint(texture))
}

// TODO:
// func CompressedTexImage2D(target TextureTarget, level int, internalformat TextureFormat, width, height int, data []interface{}) {
// 	C.glCompressedTexImage2D(C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), 0, unsafe.Sizeof(data[0])*len(data), data)
// }

// func CompressedTexSubImage2D(target TextureTarget, level, xoffset, yoffset, width, height, format TextureFormat, data []interface{}) {
// 	C.glCompressedTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), unsafe.Sizeof(data[0])*len(data), data)
// }

func CopyTexImage2D(target TextureTarget, level int, internalformat InternalFormat, x, y, width, height int) {
	C.glCopyTexImage2D(C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), 0)
}

func CopyTexSubImage2D(target TextureTarget, level, xoffset, yoffset, x, y, width, height int) {
	C.glCopyTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

func DeleteTextures(textures []Texture) {
	C.glDeleteTextures(C.GLsizei(len(textures)), (*C.GLuint)(&textures[0]))
}

func (t Texture) Delete() {
	C.glDeleteTextures(1, (*C.GLuint)(&t))
}

func CreateTexture() Texture {
	texture := Texture(0)
	C.glGenTextures(1, (*C.GLuint)(&texture))
	return texture
}

func GenTextures(textures []Texture) {
	C.glGenTextures(C.GLsizei(len(textures)), (*C.GLuint)(&textures[0]))
}

func IsTexture(texture Texture) bool {
	return C.glIsTexture(C.GLuint(texture)) == C.GL_TRUE
}

func GenerateMipmap(target TextureTarget) {
	C.glGenerateMipmap(C.GLenum(target))
}

func GetTexParameterf(target TextureTarget, pname TextureParameter) float32 {
	param := float32(0)
	C.glGetTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&param))
	return param
}

func GetTexParameteri(target TextureTarget, pname TextureParameter) int {
	param := C.GLint(0)
	C.glGetTexParameteriv(C.GLenum(target), C.GLenum(pname), &param)
	return int(param)
}

func TexImage2D(target TextureTarget, level int, internalformat InternalFormat, width, height int, format TextureFormat, datatype DataType, pixels []uint8) {
	if pixels != nil {
		C.glTexImage2D(C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), 0, C.GLenum(format), C.GLenum(datatype), unsafe.Pointer(&pixels[0]))
	} else {
		C.glTexImage2D(C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), 0, C.GLenum(format), C.GLenum(datatype), unsafe.Pointer(nil))
	}
}

func adjustUnpackAlignment(stride int) C.GLint {

	alignment := C.GLint(0)
	C.glGetIntegerv(C.GL_UNPACK_ALIGNMENT, &alignment)

	align := C.GLint(1)

	for align < alignment && stride%(int(align)*2) == 0 {
		align *= 2
	}

	// need smaller alignment
	if align < alignment {
		C.glPixelStorei(C.GL_UNPACK_ALIGNMENT, align)
		// return old alignment
		return alignment
	} else {
		return 0
	}
}

func restoreAlignment(alignment C.GLint) {
	if alignment > 0 {
		C.glPixelStorei(C.GL_UNPACK_ALIGNMENT, alignment)
	}
}

func TexImageRGBA(target TextureTarget, level int, img *image.NRGBA) {
	a := adjustUnpackAlignment(img.Stride)
	C.glTexImage2D(C.GLenum(target), C.GLint(level), C.GL_RGBA, C.GLsizei(img.Rect.Dx()), C.GLsizei(img.Rect.Dy()), 0, C.GL_RGBA, C.GL_UNSIGNED_BYTE, unsafe.Pointer(&img.Pix[0]))
	restoreAlignment(a)
}

func TexImageAlpha(target TextureTarget, level int, img *image.Alpha) {
	a := adjustUnpackAlignment(img.Stride)
	C.glTexImage2D(C.GLenum(target), C.GLint(level), C.GL_ALPHA, C.GLsizei(img.Rect.Dx()), C.GLsizei(img.Rect.Dy()), 0, C.GL_ALPHA, C.GL_UNSIGNED_BYTE, unsafe.Pointer(&img.Pix[0]))
	restoreAlignment(a)
}

func TexImageLuminance(target TextureTarget, level int, img *image.Gray) {
	a := adjustUnpackAlignment(img.Stride)
	C.glTexImage2D(C.GLenum(target), C.GLint(level), C.GL_LUMINANCE, C.GLsizei(img.Rect.Dx()), C.GLsizei(img.Rect.Dy()), 0, C.GL_LUMINANCE, C.GL_UNSIGNED_BYTE, unsafe.Pointer(&img.Pix[0]))
	restoreAlignment(a)
}

func TexParameter(target TextureTarget, pname TextureParameter, param TextureParameterValue) {
	C.glTexParameteri(C.GLenum(target), C.GLenum(pname), C.GLint(param))
}

func TexSubImage2D(target TextureTarget, level, xoffset, yoffset, width, height int, format TextureFormat, datatype DataType, pixels []uint8) {
	C.glTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(datatype), unsafe.Pointer(&pixels[0]))
}

func TexSubImageRGBA(target TextureTarget, level int, img *image.NRGBA) {
	a := adjustUnpackAlignment(img.Stride)
	C.glTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(img.Rect.Min.X), C.GLint(img.Rect.Min.Y), C.GLsizei(img.Rect.Dx()), C.GLsizei(img.Rect.Dy()), C.GL_RGBA, C.GL_UNSIGNED_BYTE, unsafe.Pointer(&img.Pix[0]))
	restoreAlignment(a)
}

func TexSubImageAlpha(target TextureTarget, level int, img *image.Alpha) {
	a := adjustUnpackAlignment(img.Stride)
	C.glTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(img.Rect.Min.X), C.GLint(img.Rect.Min.Y), C.GLsizei(img.Rect.Dx()), C.GLsizei(img.Rect.Dy()), C.GL_ALPHA, C.GL_UNSIGNED_BYTE, unsafe.Pointer(&img.Pix[0]))
	restoreAlignment(a)
}

func TexSubImageLuminance(target TextureTarget, level int, img *image.Gray) {
	a := adjustUnpackAlignment(img.Stride)
	C.glTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(img.Rect.Min.X), C.GLint(img.Rect.Min.Y), C.GLsizei(img.Rect.Dx()), C.GLsizei(img.Rect.Dy()), C.GL_LUMINANCE, C.GL_UNSIGNED_BYTE, unsafe.Pointer(&img.Pix[0]))
	restoreAlignment(a)
}
