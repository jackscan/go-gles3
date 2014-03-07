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

type Framebuffer C.GLuint
type Renderbuffer C.GLuint
type FramebufferTarget C.GLenum
type RenderbufferTarget C.GLenum
type FramebufferStatus C.GLenum

const (
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

	FRAMEBUFFER_BINDING   = C.GL_FRAMEBUFFER_BINDING
	RENDERBUFFER_BINDING  = C.GL_RENDERBUFFER_BINDING
	MAX_RENDERBUFFER_SIZE = C.GL_MAX_RENDERBUFFER_SIZE

	FRAMEBUFFER  FramebufferTarget  = C.GL_FRAMEBUFFER
	RENDERBUFFER RenderbufferTarget = C.GL_RENDERBUFFER

	FRAMEBUFFER_COMPLETE                      FramebufferStatus = C.GL_FRAMEBUFFER_COMPLETE
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                           = C.GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                   = C.GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS                           = C.GL_FRAMEBUFFER_INCOMPLETE_DIMENSIONS
	FRAMEBUFFER_UNSUPPORTED                                     = C.GL_FRAMEBUFFER_UNSUPPORTED
)

func GenFramebuffers(buffers []Framebuffer) {
	C.glGenFramebuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
}

func GenFramebuffer() Framebuffer {
	buffer := Framebuffer(0)
	C.glGenFramebuffers(C.GLsizei(1), (*C.GLuint)(&buffer))
	return buffer
}

func BindFramebuffer(target FramebufferTarget, framebuffer Framebuffer) {
	C.glBindFramebuffer(C.GLenum(target), C.GLuint(framebuffer))
}

func CheckFramebufferStatus(target FramebufferTarget) FramebufferStatus {
	return FramebufferStatus(C.glCheckFramebufferStatus(C.GLenum(target)))
}

func GenRenderbuffers(buffers []Renderbuffer) {
	C.glGenRenderbuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
}

func GenRenderbuffer() Renderbuffer {
	buffer := Renderbuffer(0)
	C.glGenRenderbuffers(C.GLsizei(1), (*C.GLuint)(&buffer))
	return buffer
}

func BindRenderbuffer(target RenderbufferTarget, renderbuffer Renderbuffer) {
	C.glBindRenderbuffer(C.GLenum(target), C.GLuint(renderbuffer))
}

// func BindFramebuffer(target int, framebuffer uint) {
// 	C.glBindFramebuffer(GLenum target, GLuint framebuffer)
// }

// func DeleteFramebuffers(GLsizei n, framebuffers []uint) {
// 	C.glDeleteFramebuffers(GLsizei n, const GLuint* framebuffers)
// }

// func DeleteRenderbuffers(GLsizei n, renderbuffers []uint) {
// 	C.glDeleteRenderbuffers(GLsizei n, const GLuint* renderbuffers)
// }

// func FramebufferRenderbuffer(target int, attachment int, renderbuffertarget int, renderbuffer uint) {
// 	C.glFramebufferRenderbuffer(GLenum target, GLenum attachment, GLenum renderbuffertarget, GLuint renderbuffer)
// }

// func FramebufferTexture2D(target int, attachment int, textarget int, texture uint, level int) {
// 	C.glFramebufferTexture2D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level)
// }

// func RenderbufferStorage(target int, internalformat int, GLsizei width, GLsizei height) {
// 	C.glRenderbufferStorage(GLenum target, GLenum internalformat, GLsizei width, GLsizei height)
// }

// func GetFramebufferAttachmentParameteriv(target int, attachment int, pname int,  params int) {
// 	C.glGetFramebufferAttachmentParameteriv(GLenum target, GLenum attachment, GLenum pname, GLint* params)
// }
