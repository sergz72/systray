package systray

/*
#cgo darwin CFLAGS: -DDARWIN -x objective-c -fobjc-arc
#cgo darwin LDFLAGS: -framework Cocoa -framework WebKit

void showPopover(void *popover);
void setObject(void *obj);
#include "systray.h"
*/
import "C"

import (
	"unsafe"
        "github.com/martinlindhe/notify"
)

// SetTemplateIcon sets the systray icon as a template icon (on Mac), falling back
// to a regular icon on other platforms.
// templateIconBytes and regularIconBytes should be the content of .ico for windows and
// .ico/.jpg/.png for other platforms.
func SetTemplateIcon(templateIconBytes []byte, regularIconBytes []byte) {
	cstr := (*C.char)(unsafe.Pointer(&templateIconBytes[0]))
	C.setIcon(cstr, (C.int)(len(templateIconBytes)), true)
}

// SetIcon sets the icon of a menu item. Only works on macOS and Windows.
// iconBytes should be the content of .ico/.jpg/.png
func (item *MenuItem) SetIcon(iconBytes []byte) {
	cstr := (*C.char)(unsafe.Pointer(&iconBytes[0]))
	C.setMenuItemIcon(cstr, (C.int)(len(iconBytes)), C.int(item.id), false)
}

// SetTemplateIcon sets the icon of a menu item as a template icon (on macOS). On Windows, it
// falls back to the regular icon bytes and on Linux it does nothing.
// templateIconBytes and regularIconBytes should be the content of .ico for windows and
// .ico/.jpg/.png for other platforms.
func (item *MenuItem) SetTemplateIcon(templateIconBytes []byte, regularIconBytes []byte) {
	cstr := (*C.char)(unsafe.Pointer(&templateIconBytes[0]))
	C.setMenuItemIcon(cstr, (C.int)(len(templateIconBytes)), C.int(item.id), true)
}

// SetInfo sets the systray info balloon
func SetInfo(appName string, text string, title string, timeout uint32, notificationType uint32) {
        // show a notification
	notify.Notify(appName, title, text, "")
}

func ShowPopover(popover unsafe.Pointer) {
	C.showPopover(popover)
}

func SetObject(obj unsafe.Pointer) {
	C.setObject(obj)
}
