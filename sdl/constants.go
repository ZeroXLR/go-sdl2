package sdl

const (
	/* SDL.h */
	INIT_TIMER			= 0x00000001
	INIT_AUDIO			= 0x00000010
	INIT_VIDEO			= 0x00000020
	INIT_JOYSTICK			= 0x00000200
	INIT_HAPTIC			= 0x00001000
	INIT_GAMECONTROLLER		= 0x00002000      /**< turn on game controller also implicitly does JOYSTICK */
	INIT_NOPARACHUTE		= 0x00100000      /**< Don't catch fatal signals */
	INIT_EVERYTHING			= INIT_TIMER | INIT_AUDIO | INIT_VIDEO | INIT_JOYSTICK |
					  INIT_HAPTIC | INIT_GAMECONTROLLER

	WINDOWPOS_UNDEFINED_MASK	= 0x1FFF0000
	WINDOWPOS_UNDEFINED		= WINDOWPOS_UNDEFINED_MASK | 0

	/* SDL_blendmode.h */
	BLENDMODE_NONE			= 0x00000000     /**< No blending */
	BLENDMODE_BLEND			= 0x00000001    /**< dst = (src * A) + (dst * (1-A)) */
	BLENDMODE_ADD			= 0x00000002      /**< dst = (src * A) + dst */
	BLENDMODE_MOD			= 0x00000004       /**< dst = src * dst */

	/* SDL_error.h */
	ENOMEM				= 0x00000000
	EFREAD				= 0x00000001
	EFWRITE				= 0x00000002
	EFSEEK				= 0x00000003
	UNSUPPORTED			= 0x00000004
	LASTERROR			= 0x00000005

	/* SDL_events.h */
	FIRSTEVENT			= 0
	QUIT				= 0x100
	APP_TERMINATING			= 0x101
	APP_LOWMEMORY			= 0x102
	APP_WILLENTERBACKGROUND		= 0x103
	APP_DIDENTERBACKGROUND		= 0x104
	APP_WILLENTERFOREGROUND		= 0x105
	APP_DIDENTERFOREGROUND		= 0x106
	/* Window events */
	WINDOWEVENT			= 0x200
	SYSWMEVENT			= 0x201
	/* Keyboard events */
	KEYDOWN				= 0x300
	KEYUP				= 0x301
	TEXTEDITING			= 0x302
	TEXTINPUT			= 0x303
	/* Mouse events */
	MOUSEMOTION			= 0x400
	MOUSEBUTTONDOWN			= 0x401
	MOUSEBUTTONUP			= 0x402
	MOUSEWHEEL			= 0x403
	/* Joystick events */
	JOYAXISMOTION			= 0x600
	JOYBALLMOTION			= 0x601
	JOYHATMOTION			= 0x602
	JOYBUTTONDOWN			= 0x603
	JOYBUTTONUP			= 0x604
	JOYDEVICEADDED			= 0x605
	JOYDEVICEREMOVED		= 0x606
	/* Game controller events */
	CONTROLLERAXISMOTION		= 0x650
	CONTROLLERBUTTONDOWN		= 0x651
	CONTROLLERBUTTONUP		= 0x652
	CONTROLLERDEVICEADDED		= 0x653
	CONTROLLERDEVICEREMOVED		= 0x654
	CONTROLLERDEVICEREMAPPED	= 0x655
	/* Touch events */
	FINGERDOWN			= 0x700
	FINGERUP			= 0x701
	FINGERMOTION			= 0x702
	/* Gesture events */
	DOLLARGESTURE			= 0x800
	DOLLARRECORD			= 0x801
	MULTIGESTURE			= 0x802
	/* Clipboard events */
	CLIPBOARDUPDATE			= 0x900
	/* Drag and drop events */
	DROPFILE			= 0x1000

	USEREVENT			= 0x8000
	LASTEVENT			= 0xFFFF

	ADDEVENT			= 0x0
	PEEKEVENT			= 0x1
	GETEVENT			= 0x2

	/* SDL_render.h */
	RENDERER_SOFTWARE		= 0x00000001
	RENDERER_ACCELERATED		= 0x00000002
	RENDERER_PRESENTVSYNC		= 0x00000004
	RENDERER_TARGETTEXTURE		= 0x00000008

	TEXTUREACCESS_STATIC		= 0x00000001
	TEXTUREACCESS_STREAMING		= 0x00000002
	TEXTUREACCESS_TARGET		= 0x00000003

	TEXTUREMODULATE_NONE		= 0x00000000     /**< No modulation */
	TEXTUREMODULATE_COLOR		= 0x00000001    /**< srcC = srcC * color */
	TEXTUREMODULATE_ALPHA		= 0x00000002     /**< srcA = srcA * alpha */

	SDL_FLIP_NONE			= 0x00000000     /**< Do not flip */
	SDL_FLIP_HORIZONTAL		= 0x00000001    /**< flip horizontally */
	SDL_FLIP_VERTICAL		= 0x00000002     /**< flip vertically */

	/* SDL_video.h */
	WINDOW_FULLSCREEN		= 0x00000001         /**< fullscreen window */
	WINDOW_OPENGL			= 0x00000002             /**< window usable with OpenGL context */
	WINDOW_SHOWN			= 0x00000004              /**< window is visible */
	WINDOW_HIDDEN			= 0x00000008             /**< window is not visible */
	WINDOW_BORDERLESS		= 0x00000010         /**< no window decoration */
	WINDOW_RESIZABLE		= 0x00000020          /**< window can be resized */
	WINDOW_MINIMIZED		= 0x00000040          /**< window is minimized */
	WINDOW_MAXIMIZED		= 0x00000080          /**< window is maximized */
	WINDOW_INPUT_GRABBED		= 0x00000100      /**< window has grabbed input focus */
	WINDOW_INPUT_FOCUS		= 0x00000200        /**< window has input focus */
	WINDOW_MOUSE_FOCUS		= 0x00000400        /**< window has mouse focus */
	WINDOW_FULLSCREEN_DESKT		= WINDOW_FULLSCREEN | 0x00001000
	WINDOW_FOREIGN			= 0x00000800             /**< window not created by SDL */

	GL_RED_SIZE			= 0x00000000
	GL_GREEN_SIZE			= 0x00000001
	GL_BLUE_SIZE			= 0x00000002
	GL_ALPHA_SIZE			= 0x00000003
	GL_BUFFER_SIZE			= 0x00000004
	GL_DOUBLEBUFFER			= 0x00000005
	GL_DEPTH_SIZE			= 0x00000006
	GL_STENCIL_SIZE			= 0x00000007
	GL_ACCUM_RED_SIZE		= 0x00000008
	GL_ACCUM_GREEN_SIZE		= 0x00000009
	GL_ACCUM_BLUE_SIZE		= 0x0000000A
	GL_ACCUM_ALPHA_SIZE		= 0x0000000B
	GL_STEREO			= 0x0000000C
	GL_MULTISAMPLEBUFFERS		= 0x0000000D
	GL_MULTISAMPLESAMPLES		= 0x0000000E
	GL_ACCELERATED_VISUAL		= 0x0000000F
	GL_RETAINED_BACKING		= 0x00000010
	GL_CONTEXT_MAJOR_VERSION	= 0x00000011
	GL_CONTEXT_MINOR_VERSION	= 0x00000012
	GL_CONTEXT_EGL			= 0x00000013
	GL_CONTEXT_FLAGS		= 0x00000014
	GL_CONTEXT_PROFILE_MASK		= 0x00000015
	GL_SHARE_WITH_CURRENT_CONTEXT	= 0x00000016

	GL_CONTEXT_PROFILE_CORE           = 0x0001
	GL_CONTEXT_PROFILE_COMPATIBILITY  = 0x0002
	GL_CONTEXT_PROFILE_ES             = 0x0004

	GL_CONTEXT_DEBUG_FLAG              = 0x0001
	GL_CONTEXT_FORWARD_COMPATIBLE_FLAG = 0x0002
	GL_CONTEXT_ROBUST_ACCESS_FLAG      = 0x0004
	GL_CONTEXT_RESET_ISOLATION_FLAG    = 0x0008
)
