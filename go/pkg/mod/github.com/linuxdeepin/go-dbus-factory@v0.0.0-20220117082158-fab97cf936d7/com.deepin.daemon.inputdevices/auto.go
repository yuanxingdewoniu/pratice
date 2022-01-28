// Code generated by "./generator ./com.deepin.daemon.inputdevices"; DO NOT EDIT.

package inputdevices

import (
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Keyboard interface {
	keyboard // interface com.deepin.daemon.InputDevice.Keyboard
	proxy.Object
}

type objectKeyboard struct {
	interfaceKeyboard // interface com.deepin.daemon.InputDevice.Keyboard
	proxy.ImplObject
}

func NewKeyboard(conn *dbus.Conn) Keyboard {
	obj := new(objectKeyboard)
	obj.ImplObject.Init_(conn, "com.deepin.daemon.InputDevices", "/com/deepin/daemon/InputDevice/Keyboard")
	return obj
}

type keyboard interface {
	GoAddLayoutOption(flags dbus.Flags, ch chan *dbus.Call, option string) *dbus.Call
	AddLayoutOption(flags dbus.Flags, option string) error
	GoAddUserLayout(flags dbus.Flags, ch chan *dbus.Call, layout string) *dbus.Call
	AddUserLayout(flags dbus.Flags, layout string) error
	GoClearLayoutOption(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ClearLayoutOption(flags dbus.Flags) error
	GoDeleteLayoutOption(flags dbus.Flags, ch chan *dbus.Call, option string) *dbus.Call
	DeleteLayoutOption(flags dbus.Flags, option string) error
	GoDeleteUserLayout(flags dbus.Flags, ch chan *dbus.Call, layout string) *dbus.Call
	DeleteUserLayout(flags dbus.Flags, layout string) error
	GoGetLayoutDesc(flags dbus.Flags, ch chan *dbus.Call, layout string) *dbus.Call
	GetLayoutDesc(flags dbus.Flags, layout string) (string, error)
	GoLayoutList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	LayoutList(flags dbus.Flags) (map[string]string, error)
	GoReset(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Reset(flags dbus.Flags) error
	UserOptionList() proxy.PropStringArray
	RepeatEnabled() proxy.PropBool
	CapslockToggle() proxy.PropBool
	CursorBlink() proxy.PropInt32
	RepeatInterval() proxy.PropUint32
	RepeatDelay() proxy.PropUint32
	CurrentLayout() proxy.PropString
	UserLayoutList() proxy.PropStringArray
}

type interfaceKeyboard struct{}

func (v *interfaceKeyboard) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceKeyboard) GetInterfaceName_() string {
	return "com.deepin.daemon.InputDevice.Keyboard"
}

// method AddLayoutOption

func (v *interfaceKeyboard) GoAddLayoutOption(flags dbus.Flags, ch chan *dbus.Call, option string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddLayoutOption", flags, ch, option)
}

func (v *interfaceKeyboard) AddLayoutOption(flags dbus.Flags, option string) error {
	return (<-v.GoAddLayoutOption(flags, make(chan *dbus.Call, 1), option).Done).Err
}

// method AddUserLayout

func (v *interfaceKeyboard) GoAddUserLayout(flags dbus.Flags, ch chan *dbus.Call, layout string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddUserLayout", flags, ch, layout)
}

func (v *interfaceKeyboard) AddUserLayout(flags dbus.Flags, layout string) error {
	return (<-v.GoAddUserLayout(flags, make(chan *dbus.Call, 1), layout).Done).Err
}

// method ClearLayoutOption

func (v *interfaceKeyboard) GoClearLayoutOption(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClearLayoutOption", flags, ch)
}

func (v *interfaceKeyboard) ClearLayoutOption(flags dbus.Flags) error {
	return (<-v.GoClearLayoutOption(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method DeleteLayoutOption

func (v *interfaceKeyboard) GoDeleteLayoutOption(flags dbus.Flags, ch chan *dbus.Call, option string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteLayoutOption", flags, ch, option)
}

func (v *interfaceKeyboard) DeleteLayoutOption(flags dbus.Flags, option string) error {
	return (<-v.GoDeleteLayoutOption(flags, make(chan *dbus.Call, 1), option).Done).Err
}

// method DeleteUserLayout

func (v *interfaceKeyboard) GoDeleteUserLayout(flags dbus.Flags, ch chan *dbus.Call, layout string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteUserLayout", flags, ch, layout)
}

func (v *interfaceKeyboard) DeleteUserLayout(flags dbus.Flags, layout string) error {
	return (<-v.GoDeleteUserLayout(flags, make(chan *dbus.Call, 1), layout).Done).Err
}

// method GetLayoutDesc

func (v *interfaceKeyboard) GoGetLayoutDesc(flags dbus.Flags, ch chan *dbus.Call, layout string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetLayoutDesc", flags, ch, layout)
}

func (*interfaceKeyboard) StoreGetLayoutDesc(call *dbus.Call) (description string, err error) {
	err = call.Store(&description)
	return
}

func (v *interfaceKeyboard) GetLayoutDesc(flags dbus.Flags, layout string) (string, error) {
	return v.StoreGetLayoutDesc(
		<-v.GoGetLayoutDesc(flags, make(chan *dbus.Call, 1), layout).Done)
}

// method LayoutList

func (v *interfaceKeyboard) GoLayoutList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LayoutList", flags, ch)
}

func (*interfaceKeyboard) StoreLayoutList(call *dbus.Call) (layout_list map[string]string, err error) {
	err = call.Store(&layout_list)
	return
}

func (v *interfaceKeyboard) LayoutList(flags dbus.Flags) (map[string]string, error) {
	return v.StoreLayoutList(
		<-v.GoLayoutList(flags, make(chan *dbus.Call, 1)).Done)
}

// method Reset

func (v *interfaceKeyboard) GoReset(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reset", flags, ch)
}

func (v *interfaceKeyboard) Reset(flags dbus.Flags) error {
	return (<-v.GoReset(flags, make(chan *dbus.Call, 1)).Done).Err
}

// property UserOptionList as

func (v *interfaceKeyboard) UserOptionList() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "UserOptionList",
	}
}

// property RepeatEnabled b

func (v *interfaceKeyboard) RepeatEnabled() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "RepeatEnabled",
	}
}

// property CapslockToggle b

func (v *interfaceKeyboard) CapslockToggle() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "CapslockToggle",
	}
}

// property CursorBlink i

func (v *interfaceKeyboard) CursorBlink() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "CursorBlink",
	}
}

// property RepeatInterval u

func (v *interfaceKeyboard) RepeatInterval() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "RepeatInterval",
	}
}

// property RepeatDelay u

func (v *interfaceKeyboard) RepeatDelay() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "RepeatDelay",
	}
}

// property CurrentLayout s

func (v *interfaceKeyboard) CurrentLayout() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "CurrentLayout",
	}
}

// property UserLayoutList as

func (v *interfaceKeyboard) UserLayoutList() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "UserLayoutList",
	}
}

type TouchPad interface {
	touchPad // interface com.deepin.daemon.InputDevice.TouchPad
	proxy.Object
}

type objectTouchPad struct {
	interfaceTouchPad // interface com.deepin.daemon.InputDevice.TouchPad
	proxy.ImplObject
}

func NewTouchPad(conn *dbus.Conn) TouchPad {
	obj := new(objectTouchPad)
	obj.ImplObject.Init_(conn, "com.deepin.daemon.InputDevices", "/com/deepin/daemon/InputDevice/TouchPad")
	return obj
}

type touchPad interface {
	GoReset(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Reset(flags dbus.Flags) error
	EdgeScroll() proxy.PropBool
	PalmDetect() proxy.PropBool
	MotionAcceleration() proxy.PropDouble
	DeltaScroll() proxy.PropInt32
	DragThreshold() proxy.PropInt32
	LeftHanded() proxy.PropBool
	DisableIfTyping() proxy.PropBool
	NaturalScroll() proxy.PropBool
	HorizScroll() proxy.PropBool
	VertScroll() proxy.PropBool
	MotionThreshold() proxy.PropDouble
	DoubleClick() proxy.PropInt32
	DeviceList() proxy.PropString
	TPadEnable() proxy.PropBool
	PalmMinZ() proxy.PropInt32
	Exist() proxy.PropBool
	TapClick() proxy.PropBool
	MotionScaling() proxy.PropDouble
	PalmMinWidth() proxy.PropInt32
}

type interfaceTouchPad struct{}

func (v *interfaceTouchPad) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceTouchPad) GetInterfaceName_() string {
	return "com.deepin.daemon.InputDevice.TouchPad"
}

// method Reset

func (v *interfaceTouchPad) GoReset(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Reset", flags, ch)
}

func (v *interfaceTouchPad) Reset(flags dbus.Flags) error {
	return (<-v.GoReset(flags, make(chan *dbus.Call, 1)).Done).Err
}

// property EdgeScroll b

func (v *interfaceTouchPad) EdgeScroll() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "EdgeScroll",
	}
}

// property PalmDetect b

func (v *interfaceTouchPad) PalmDetect() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "PalmDetect",
	}
}

// property MotionAcceleration d

func (v *interfaceTouchPad) MotionAcceleration() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "MotionAcceleration",
	}
}

// property DeltaScroll i

func (v *interfaceTouchPad) DeltaScroll() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "DeltaScroll",
	}
}

// property DragThreshold i

func (v *interfaceTouchPad) DragThreshold() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "DragThreshold",
	}
}

// property LeftHanded b

func (v *interfaceTouchPad) LeftHanded() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "LeftHanded",
	}
}

// property DisableIfTyping b

func (v *interfaceTouchPad) DisableIfTyping() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "DisableIfTyping",
	}
}

// property NaturalScroll b

func (v *interfaceTouchPad) NaturalScroll() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "NaturalScroll",
	}
}

// property HorizScroll b

func (v *interfaceTouchPad) HorizScroll() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "HorizScroll",
	}
}

// property VertScroll b

func (v *interfaceTouchPad) VertScroll() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "VertScroll",
	}
}

// property MotionThreshold d

func (v *interfaceTouchPad) MotionThreshold() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "MotionThreshold",
	}
}

// property DoubleClick i

func (v *interfaceTouchPad) DoubleClick() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "DoubleClick",
	}
}

// property DeviceList s

func (v *interfaceTouchPad) DeviceList() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "DeviceList",
	}
}

// property TPadEnable b

func (v *interfaceTouchPad) TPadEnable() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "TPadEnable",
	}
}

// property PalmMinZ i

func (v *interfaceTouchPad) PalmMinZ() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "PalmMinZ",
	}
}

// property Exist b

func (v *interfaceTouchPad) Exist() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Exist",
	}
}

// property TapClick b

func (v *interfaceTouchPad) TapClick() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "TapClick",
	}
}

// property MotionScaling d

func (v *interfaceTouchPad) MotionScaling() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "MotionScaling",
	}
}

// property PalmMinWidth i

func (v *interfaceTouchPad) PalmMinWidth() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "PalmMinWidth",
	}
}