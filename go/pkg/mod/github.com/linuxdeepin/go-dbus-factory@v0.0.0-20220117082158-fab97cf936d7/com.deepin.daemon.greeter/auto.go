// Code generated by "./generator ./com.deepin.daemon.greeter"; DO NOT EDIT.

package greeter

import (
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Greeter interface {
	greeter // interface com.deepin.daemon.Greeter
	proxy.Object
}

type objectGreeter struct {
	interfaceGreeter // interface com.deepin.daemon.Greeter
	proxy.ImplObject
}

func NewGreeter(conn *dbus.Conn) Greeter {
	obj := new(objectGreeter)
	obj.ImplObject.Init_(conn, "com.deepin.daemon.Greeter", "/com/deepin/daemon/Greeter")
	return obj
}

type greeter interface {
	GoUpdateGreeterQtTheme(flags dbus.Flags, ch chan *dbus.Call, fd dbus.UnixFD) *dbus.Call
	UpdateGreeterQtTheme(flags dbus.Flags, fd dbus.UnixFD) error
}

type interfaceGreeter struct{}

func (v *interfaceGreeter) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceGreeter) GetInterfaceName_() string {
	return "com.deepin.daemon.Greeter"
}

// method UpdateGreeterQtTheme

func (v *interfaceGreeter) GoUpdateGreeterQtTheme(flags dbus.Flags, ch chan *dbus.Call, fd dbus.UnixFD) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UpdateGreeterQtTheme", flags, ch, fd)
}

func (v *interfaceGreeter) UpdateGreeterQtTheme(flags dbus.Flags, fd dbus.UnixFD) error {
	return (<-v.GoUpdateGreeterQtTheme(flags, make(chan *dbus.Call, 1), fd).Done).Err
}