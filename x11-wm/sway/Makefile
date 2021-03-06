# $FreeBSD$

PORTNAME=	sway
DISTVERSION=	1.5.1
CATEGORIES=	x11-wm

PATCH_SITES=	https://github.com/${GH_ACCOUNT}/${GH_PROJECT}/commit/
PATCHFILES+=	cd9d1038a822.patch:-p1 # https://github.com/swaywm/sway/pull/5090
PATCHFILES+=	f98ca3ab7ca1.patch:-p1 # https://github.com/swaywm/sway/pull/5090

MAINTAINER=	jbeich@FreeBSD.org
COMMENT=	i3-compatible Wayland compositor

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	evdev-proto>0:devel/evdev-proto \
		wayland-protocols>=1.14:graphics/wayland-protocols \
		wlroots>=0.12.0<0.13.0:x11-toolkits/wlroots
LIB_DEPENDS=	libjson-c.so:devel/json-c \
		libevdev.so:devel/libevdev \
		libpcre.so:devel/pcre \
		libwayland-egl.so:graphics/wayland \
		libwlroots.so:x11-toolkits/wlroots \
		libinput.so:x11/libinput \
		libxkbcommon.so:x11/libxkbcommon
RUN_DEPENDS=	swaybg:x11/swaybg

USES=		compiler:c11 gl gnome meson pkgconfig xorg
USE_GITHUB=	yes
USE_GL=		glesv2
USE_GNOME=	cairo pango
USE_XORG=	pixman
GH_ACCOUNT=	swaywm

OPTIONS_DEFINE=	MANPAGES PIXBUF SUID X11
OPTIONS_DEFAULT=MANPAGES PIXBUF SUID X11
OPTIONS_SUB=	yes

MANPAGES_BUILD_DEPENDS=	scdoc>=1.9.2:textproc/scdoc
MANPAGES_MESON_ENABLED=	man-pages

PIXBUF_USE=		GNOME=gdkpixbuf2
PIXBUF_MESON_ENABLED=	gdk-pixbuf

SUID_DESC=		setuid bit on "sway" binary (to claim DRM without seatd/console-kit-daemon)
SUID_PLIST_SUB=		MAYBE_SUID="@(,,4755) "
SUID_PLIST_SUB_OFF=	MAYBE_SUID=""

X11_USE=		XORG=xcb
X11_MESON_ENABLED=	xwayland

post-patch:
# Let @sample handle default files under etc/
	@${REINPLACE_CMD} -e '/config\.in/{ N; s/@BASENAME@/&.sample/; }' \
		${WRKSRC}/meson.build
# Extract (snapshot) version from the port instead of meson.build
	@${REINPLACE_CMD} -i .nogit -e 's/git.found()/false/' \
		-e '/project_version/s/@0@/${DISTVERSIONFULL}/' \
		${WRKSRC}/meson.build

.include <bsd.port.mk>
