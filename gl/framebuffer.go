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
	"fmt"
)

type Framebuffer C.GLuint
type Renderbuffer C.GLuint
type FramebufferTarget C.GLenum
type RenderbufferTarget C.GLenum
type FramebufferStatus C.GLenum
type FramebufferAttachment C.GLenum
type InternalFormat C.GLenum

const (
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

	NONE = C.GL_NONE

	FRAMEBUFFER_BINDING   = C.GL_FRAMEBUFFER_BINDING
	RENDERBUFFER_BINDING  = C.GL_RENDERBUFFER_BINDING
	MAX_RENDERBUFFER_SIZE = C.GL_MAX_RENDERBUFFER_SIZE

	FRAMEBUFFER  FramebufferTarget  = C.GL_FRAMEBUFFER
	RENDERBUFFER RenderbufferTarget = C.GL_RENDERBUFFER

	COLOR_ATTACHMENT0  FramebufferAttachment = C.GL_COLOR_ATTACHMENT0
	DEPTH_ATTACHMENT                         = C.GL_DEPTH_ATTACHMENT
	STENCIL_ATTACHMENT                       = C.GL_STENCIL_ATTACHMENT

	FRAMEBUFFER_COMPLETE                      FramebufferStatus = C.GL_FRAMEBUFFER_COMPLETE
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                           = C.GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                   = C.GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS                           = C.GL_FRAMEBUFFER_INCOMPLETE_DIMENSIONS
	FRAMEBUFFER_UNSUPPORTED                                     = C.GL_FRAMEBUFFER_UNSUPPORTED

	RGBA4             InternalFormat = C.GL_RGBA4
	RGB5_A1                          = C.GL_RGB5_A1
	RGB565                           = C.GL_RGB565
	DEPTH_COMPONENT16                = C.GL_DEPTH_COMPONENT16
	STENCIL_INDEX8                   = C.GL_STENCIL_INDEX8
)

func GenFramebuffers(buffers []Framebuffer) {
	C.glGenFramebuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
}

func GenFramebuffer() Framebuffer {
	buffer := Framebuffer(0)
	C.glGenFramebuffers(C.GLsizei(1), (*C.GLuint)(&buffer))
	return buffer
}

func DeleteFramebuffers(framebuffers []Framebuffer) {
	C.glDeleteFramebuffers(C.GLsizei(len(framebuffers)), (*C.GLuint)(&framebuffers[0]))
}

func (f Framebuffer) Delete() {
	C.glDeleteFramebuffers(1, (*C.GLuint)(&f))
}

func BindFramebuffer(target FramebufferTarget, framebuffer Framebuffer) {
	C.glBindFramebuffer(C.GLenum(target), C.GLuint(framebuffer))
}

func CheckFramebufferStatus(target FramebufferTarget) FramebufferStatus {
	return FramebufferStatus(C.glCheckFramebufferStatus(C.GLenum(target)))
}

func FramebufferRenderbuffer(target FramebufferTarget, attachment FramebufferAttachment, renderbuffertarget RenderbufferTarget, renderbuffer Renderbuffer) {
	C.glFramebufferRenderbuffer(C.GLenum(target), C.GLenum(attachment), C.GLenum(renderbuffertarget), C.GLuint(renderbuffer))
}

func FramebufferTexture2D(target FramebufferTarget, attachment FramebufferAttachment, textarget TextureTarget, texture Texture, level int) {
	C.glFramebufferTexture2D(C.GLenum(target), C.GLenum(attachment), C.GLenum(textarget), C.GLuint(texture), C.GLint(level))
}

func GenRenderbuffers(buffers []Renderbuffer) {
	C.glGenRenderbuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
}

func GenRenderbuffer() Renderbuffer {
	buffer := Renderbuffer(0)
	C.glGenRenderbuffers(C.GLsizei(1), (*C.GLuint)(&buffer))
	return buffer
}

func DeleteRenderbuffers(renderbuffers []Renderbuffer) {
	C.glDeleteRenderbuffers(C.GLsizei(len(renderbuffers)), (*C.GLuint)(&renderbuffers[0]))
}

func (r Renderbuffer) Delete() {
	C.glDeleteRenderbuffers(1, (*C.GLuint)(&r))
}

func BindRenderbuffer(target RenderbufferTarget, renderbuffer Renderbuffer) {
	C.glBindRenderbuffer(C.GLenum(target), C.GLuint(renderbuffer))
}

func RenderbufferStorage(target RenderbufferTarget, internalformat InternalFormat, width, height int) {
	C.glRenderbufferStorage(C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height))
}

// func GetFramebufferAttachmentParameteriv(target int, attachment int, pname int,  params int) {
// 	C.glGetFramebufferAttachmentParameteriv(GLenum target, GLenum attachment, GLenum pname, GLint* params)
// }

func (s FramebufferStatus) String() string {
	switch s {
	case FRAMEBUFFER_COMPLETE:
		return "FRAMEBUFFER_COMPLETE"
	case FRAMEBUFFER_INCOMPLETE_ATTACHMENT:
		return "FRAMEBUFFER_INCOMPLETE_ATTACHMENT"
	case FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT:
		return "FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT"
	case FRAMEBUFFER_INCOMPLETE_DIMENSIONS:
		return "FRAMEBUFFER_INCOMPLETE_DIMENSIONS"
	case FRAMEBUFFER_UNSUPPORTED:
		return "FRAMEBUFFER_UNSUPPORTED"
	default:
		return fmt.Sprintf("FramebufferStatus(%#x)", s)
	}
}
