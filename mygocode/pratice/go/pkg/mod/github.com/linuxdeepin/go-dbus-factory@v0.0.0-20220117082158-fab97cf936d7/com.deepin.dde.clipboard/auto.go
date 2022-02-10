// Code generated by "./generator ./com.deepin.dde.clipboard"; DO NOT EDIT.

package clipboard

import (
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Clipboard interface {
	clipboard // interface com.deepin.dde.Clipboard
	proxy.Object
}

type objectClipboard struct {
	interfaceClipboard // interface com.deepin.dde.Clipboard
	proxy.ImplObject
}

func NewClipboard(conn *dbus.Conn) Clipboard {
	obj := new(objectClipboard)
	obj.ImplObject.Init_(conn, "com.deepin.dde.Clipboard", "/com/deepin/dde/Clipboard")
	return obj
}

type clipboard interface {
	GoToggle(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Toggle(flags dbus.Flags) error
	GoShow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Show(flags dbus.Flags) error
	GoHide(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Hide(flags dbus.Flags) error
}

type interfaceClipboard struct{}

func (v *interfaceClipboard) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceClipboard) GetInterfaceName_() string {
	return "com.deepin.dde.Clipboard"
}

// method Toggle

func (v *interfaceClipboard) GoToggle(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Toggle", flags, ch)
}

func (v *interfaceClipboard) Toggle(flags dbus.Flags) error {
	return (<-v.GoToggle(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Show

func (v *interfaceClipboard) GoShow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Show", flags, ch)
}

func (v *interfaceClipboard) Show(flags dbus.Flags) error {
	return (<-v.GoShow(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Hide

func (v *interfaceClipboard) GoHide(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Hide", flags, ch)
}

func (v *interfaceClipboard) Hide(flags dbus.Flags) error {
	return (<-v.GoHide(flags, make(chan *dbus.Call, 1)).Done).Err
}