// Code generated by "./generator ./com.deepin.daemon.powermanager"; DO NOT EDIT.

package powermanager

import (
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type PowerManager interface {
	powerManager // interface com.deepin.daemon.PowerManager
	proxy.Object
}

type objectPowerManager struct {
	interfacePowerManager // interface com.deepin.daemon.PowerManager
	proxy.ImplObject
}

func NewPowerManager(conn *dbus.Conn) PowerManager {
	obj := new(objectPowerManager)
	obj.ImplObject.Init_(conn, "com.deepin.daemon.PowerManager", "/com/deepin/daemon/PowerManager")
	return obj
}

type powerManager interface {
	GoCanShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CanShutdown(flags dbus.Flags) (bool, error)
	GoCanReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CanReboot(flags dbus.Flags) (bool, error)
	GoCanSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CanSuspend(flags dbus.Flags) (bool, error)
	GoCanHibernate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CanHibernate(flags dbus.Flags) (bool, error)
}

type interfacePowerManager struct{}

func (v *interfacePowerManager) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfacePowerManager) GetInterfaceName_() string {
	return "com.deepin.daemon.PowerManager"
}

// method CanShutdown

func (v *interfacePowerManager) GoCanShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanShutdown", flags, ch)
}

func (*interfacePowerManager) StoreCanShutdown(call *dbus.Call) (can bool, err error) {
	err = call.Store(&can)
	return
}

func (v *interfacePowerManager) CanShutdown(flags dbus.Flags) (bool, error) {
	return v.StoreCanShutdown(
		<-v.GoCanShutdown(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanReboot

func (v *interfacePowerManager) GoCanReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanReboot", flags, ch)
}

func (*interfacePowerManager) StoreCanReboot(call *dbus.Call) (can bool, err error) {
	err = call.Store(&can)
	return
}

func (v *interfacePowerManager) CanReboot(flags dbus.Flags) (bool, error) {
	return v.StoreCanReboot(
		<-v.GoCanReboot(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanSuspend

func (v *interfacePowerManager) GoCanSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanSuspend", flags, ch)
}

func (*interfacePowerManager) StoreCanSuspend(call *dbus.Call) (can bool, err error) {
	err = call.Store(&can)
	return
}

func (v *interfacePowerManager) CanSuspend(flags dbus.Flags) (bool, error) {
	return v.StoreCanSuspend(
		<-v.GoCanSuspend(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanHibernate

func (v *interfacePowerManager) GoCanHibernate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanHibernate", flags, ch)
}

func (*interfacePowerManager) StoreCanHibernate(call *dbus.Call) (can bool, err error) {
	err = call.Store(&can)
	return
}

func (v *interfacePowerManager) CanHibernate(flags dbus.Flags) (bool, error) {
	return v.StoreCanHibernate(
		<-v.GoCanHibernate(flags, make(chan *dbus.Call, 1)).Done)
}