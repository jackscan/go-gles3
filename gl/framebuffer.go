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
	"fmt"
)

type Framebuffer C.GLuint
type Renderbuffer C.GLuint
type FramebufferTarget C.GLenum
type RenderbufferTarget C.GLenum
type FramebufferStatus C.GLenum
type FramebufferAttachment C.GLenum
type FramebufferAttachmentParameter C.GLenum

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

	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE           FramebufferAttachmentParameter = C.GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME           FramebufferAttachmentParameter = C.GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL         FramebufferAttachmentParameter = C.GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE FramebufferAttachmentParameter = C.GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING        FramebufferAttachmentParameter = C.GL_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE        FramebufferAttachmentParameter = C.GL_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE
	FRAMEBUFFER_ATTACHMENT_RED_SIZE              FramebufferAttachmentParameter = C.GL_FRAMEBUFFER_ATTACHMENT_RED_SIZE
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE            FramebufferAttachmentParameter = C.GL_FRAMEBUFFER_ATTACHMENT_GREEN_SIZE
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE             FramebufferAttachmentParameter = C.GL_FRAMEBUFFER_ATTACHMENT_BLUE_SIZE
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE            FramebufferAttachmentParameter = C.GL_FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE            FramebufferAttachmentParameter = C.GL_FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE          FramebufferAttachmentParameter = C.GL_FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE

	NONE = C.GL_NONE

	FRAMEBUFFER_BINDING      = C.GL_FRAMEBUFFER_BINDING
	DRAW_FRAMEBUFFER_BINDING = C.GL_DRAW_FRAMEBUFFER_BINDING
	READ_FRAMEBUFFER_BINDING = C.GL_READ_FRAMEBUFFER_BINDING
	RENDERBUFFER_BINDING     = C.GL_RENDERBUFFER_BINDING

	MAX_RENDERBUFFER_SIZE = C.GL_MAX_RENDERBUFFER_SIZE

	FRAMEBUFFER      FramebufferTarget = C.GL_FRAMEBUFFER
	READ_FRAMEBUFFER FramebufferTarget = C.GL_READ_FRAMEBUFFER
	DRAW_FRAMEBUFFER FramebufferTarget = C.GL_DRAW_FRAMEBUFFER

	RENDERBUFFER RenderbufferTarget = C.GL_RENDERBUFFER

	COLOR_ATTACHMENT0  FramebufferAttachment = C.GL_COLOR_ATTACHMENT0
	DEPTH_ATTACHMENT   FramebufferAttachment = C.GL_DEPTH_ATTACHMENT
	STENCIL_ATTACHMENT FramebufferAttachment = C.GL_STENCIL_ATTACHMENT

	FRAMEBUFFER_COMPLETE                      FramebufferStatus = C.GL_FRAMEBUFFER_COMPLETE
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT         FramebufferStatus = C.GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT FramebufferStatus = C.GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS         FramebufferStatus = C.GL_FRAMEBUFFER_INCOMPLETE_DIMENSIONS
	FRAMEBUFFER_UNSUPPORTED                   FramebufferStatus = C.GL_FRAMEBUFFER_UNSUPPORTED
)

func GenFramebuffers(buffers []Framebuffer) {
	C.glGenFramebuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
}

func CreateFramebuffer() Framebuffer {
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

func GetDrawFramebufferBinding() Framebuffer {
	framebuffer := C.GLint(0)
	C.glGetIntegerv(C.GL_DRAW_FRAMEBUFFER_BINDING, &framebuffer)
	return Framebuffer(framebuffer)
}

func GetReadFramebufferBinding() Framebuffer {
	framebuffer := C.GLint(0)
	C.glGetIntegerv(C.GL_READ_FRAMEBUFFER_BINDING, &framebuffer)
	return Framebuffer(framebuffer)
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

func GetFramebufferAttachmentParameter(target FramebufferTarget, attachment FramebufferAttachment, pname FramebufferAttachmentParameter) int {
	param := C.GLint(0)
	C.glGetFramebufferAttachmentParameteriv(C.GLenum(target), C.GLenum(attachment), C.GLenum(pname), &param)
	return int(param)
}

func GenRenderbuffers(buffers []Renderbuffer) {
	C.glGenRenderbuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
}

func CreateRenderbuffer() Renderbuffer {
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
