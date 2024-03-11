package webgl

/* ClearBufferMask */
const (
	DEPTH_BUFFER_BIT   GLenum = 0x00000100
	STENCIL_BUFFER_BIT GLenum = 0x00000400
	COLOR_BUFFER_BIT   GLenum = 0x00004000
)

/* Shaders */
const (
	FRAGMENT_SHADER                  GLenum = 0x8B30
	VERTEX_SHADER                    GLenum = 0x8B31
	MAX_VERTEX_ATTRIBS               GLenum = 0x8869
	MAX_VERTEX_UNIFORM_VECTORS       GLenum = 0x8DFB
	MAX_VARYING_VECTORS              GLenum = 0x8DFC
	MAX_COMBINED_TEXTURE_IMAGE_UNITS GLenum = 0x8B4D
	MAX_VERTEX_TEXTURE_IMAGE_UNITS   GLenum = 0x8B4C
	MAX_TEXTURE_IMAGE_UNITS          GLenum = 0x8872
	MAX_FRAGMENT_UNIFORM_VECTORS     GLenum = 0x8DFD
	SHADER_TYPE                      GLenum = 0x8B4F
	DELETE_STATUS                    GLenum = 0x8B80
	LINK_STATUS                      GLenum = 0x8B82
	VALIDATE_STATUS                  GLenum = 0x8B83
	ATTACHED_SHADERS                 GLenum = 0x8B85
	ACTIVE_UNIFORMS                  GLenum = 0x8B86
	ACTIVE_ATTRIBUTES                GLenum = 0x8B89
	SHADING_LANGUAGE_VERSION         GLenum = 0x8B8C
	CURRENT_PROGRAM                  GLenum = 0x8B8D
)

/* Buffer Objects */
const (
	ARRAY_BUFFER                 GLenum = 0x8892
	ELEMENT_ARRAY_BUFFER         GLenum = 0x8893
	ARRAY_BUFFER_BINDING         GLenum = 0x8894
	ELEMENT_ARRAY_BUFFER_BINDING GLenum = 0x8895
	STREAM_DRAW                  GLenum = 0x88E0
	STATIC_DRAW                  GLenum = 0x88E4
	DYNAMIC_DRAW                 GLenum = 0x88E8
	BUFFER_SIZE                  GLenum = 0x8764
	BUFFER_USAGE                 GLenum = 0x8765
	CURRENT_VERTEX_ATTRIB        GLenum = 0x8626
)

/* DataType */
const (
	BYTE           GLenum = 0x1400
	UNSIGNED_BYTE  GLenum = 0x1401
	SHORT          GLenum = 0x1402
	UNSIGNED_SHORT GLenum = 0x1403
	INT            GLenum = 0x1404
	UNSIGNED_INT   GLenum = 0x1405
	FLOAT          GLenum = 0x1406
)

/* TEXTURE_2D */
const (
	CULL_FACE                GLenum = 0x0B44
	BLEND                    GLenum = 0x0BE2
	DITHER                   GLenum = 0x0BD0
	STENCIL_TEST             GLenum = 0x0B90
	DEPTH_TEST               GLenum = 0x0B71
	SCISSOR_TEST             GLenum = 0x0C11
	POLYGON_OFFSET_FILL      GLenum = 0x8037
	SAMPLE_ALPHA_TO_COVERAGE GLenum = 0x809E
	SAMPLE_COVERAGE          GLenum = 0x80A0
)

/* StencilFunction */
const (
	NEVER    GLenum = 0x0200
	LESS     GLenum = 0x0201
	EQUAL    GLenum = 0x0202
	LEQUAL   GLenum = 0x0203
	GREATER  GLenum = 0x0204
	NOTEQUAL GLenum = 0x0205
	GEQUAL   GLenum = 0x0206
	ALWAYS   GLenum = 0x0207
)

/* BeginMode */
const (
	POINTS         GLenum = 0x0000
	LINES          GLenum = 0x0001
	LINE_LOOP      GLenum = 0x0002
	LINE_STRIP     GLenum = 0x0003
	TRIANGLES      GLenum = 0x0004
	TRIANGLE_STRIP GLenum = 0x0005
	TRIANGLE_FAN   GLenum = 0x0006
)
