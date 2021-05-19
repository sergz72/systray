package systray

/*
#cgo linux pkg-config: gtk+-3.0 appindicator3-0.1 libnotify

void setInfo(char *title, char *text);
*/
import "C"

// SetTemplateIcon sets the systray icon as a template icon (on macOS), falling back
// to a regular icon on other platforms.
// templateIconBytes and iconBytes should be the content of .ico for windows and
// .ico/.jpg/.png for other platforms.
func SetTemplateIcon(templateIconBytes []byte, regularIconBytes []byte) {
	SetIcon(regularIconBytes)
}

// SetIcon sets the icon of a menu item. Only works on macOS and Windows.
// iconBytes should be the content of .ico/.jpg/.png
func (item *MenuItem) SetIcon(iconBytes []byte) {
}

// SetTemplateIcon sets the icon of a menu item as a template icon (on macOS). On Windows, it
// falls back to the regular icon bytes and on Linux it does nothing.
// templateIconBytes and regularIconBytes should be the content of .ico for windows and
// .ico/.jpg/.png for other platforms.
func (item *MenuItem) SetTemplateIcon(templateIconBytes []byte, regularIconBytes []byte) {
}

// SetInfo sets the systray info balloon
func SetInfo(appName string, text string, title string, timeout uint32, notificationType uint32) {
        // show a notification
	C.setInfo(C.CString(title), C.CString(text))
}
