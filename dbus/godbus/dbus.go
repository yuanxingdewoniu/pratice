package  main

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"log"
	"time"
)


// 导出对象，
type Obj struct {
}
func (o *Obj) GetString() (string, *dbus.Error) {
	return "object", nil
}

//goland:noinspection GoUnreachableCode
func main()  {



	//调用方法
	sessionBus, err := dbus.SessionBus()
	obj :=sessionBus.Object("org.freedesktop.DBus","/org/freedesktop/DBus")
	var ret string
	//obj.Call 方法的第一个参数 method 的值是接口名 + "." + 方法名
	err = obj.Call("org.freedesktop.DBus.GetNameOwner", 0, "com.deepin.dde.daemon.Dock").Store(&ret)
	fmt.Println("err = ", err)


	// 获取属性 获取单个属性是调用对象的 org.freedesktop.DBus.Properties
	// 接口的 Get 方法，此方法传入参数是接口名和属性名，返回值是属性值，类型为 dbus.Variant 。
	var Var_ret dbus.Variant
	err = obj.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.DBus", "Features").Store(&Var_ret)
	v := Var_ret.Value().([]string)
	fmt.Println(v)


	// 设置属性

	propVal := dbus.MakeVariant([]string{"abc"})
	err = obj.Call("org.freedesktop.DBus.Properties.Set" , 0, "org.freedesktop.DBus", "Features", propVal).Err
	fmt.Println("org.freedesktop.DBus.Properties.Set", err)

	// 监控信号 调用 AddMatch 方法添加监控规则，如果不需要监控了，就调用 RemoveMatch 方法去掉监控规则。
	// go 的 dbus 库提供了方便的方法：=AddMatchSignal= 和 =RemoveMatchSignal=。
	err = sessionBus.BusObject().AddMatchSignal("org.freedesktop.DBus", "NameOwnerChanged",
		dbus.WithMatchObjectPath("/org/freedesktop/DBus")).Err

	signalCh := make(chan  *dbus.Signal, 10)
	sessionBus.Signal(signalCh)
	go func() {
		for {
			select {
			case sig := <-signalCh:
			//	log.Printf("sig: %#v\n", sig)  //输出dbus 信息
				if sig.Path == "/org/freedesktop/DBus" &&
					sig.Name == "org.freedesktop.DBus.NameOwnerChanged" {
					var name string
					var oldOwner string
					var newOwner string
					err = dbus.Store(sig.Body, &name, &oldOwner, &newOwner)
				//	log.Printf("%s %s %s\n", name, oldOwner, newOwner)
				}
			}
		}
	}()

	//time.Sleep(100*time.Second)
	//err = sessionBus.BusObject().RemoveMatchSignal("org.freedesktop.DBus", "NameOwnerChanged",
	//	dbus.WithMatchObjectPath("/org/freedesktop/Bus")).Err
	//sessionBus.RemoveSignal(signalCh)

	//将调用 sessionBus.Signal 方法注册 signalCh 通道，以此来通过该通道接收信号，
	//go func () 中的代码是循环接收 signalCh 通道传来的数据并处理。如果不需要再从这个通道接收信号了，
	//可以调用 sessionBus.RemoveSignal 取消注册 signalCh 通道，这样信号就不会再被发送到 signalCh 通道中了。
	//在处理信号时，可以先验证信号（ *dbus.Signal 类型）的 Path 和 Name 字段，不必验证 Sender 字段。
	//处理信号的 Body 字段，它是 []interface{} 类型，就可以用 dbus.Store 方法了。


	//导出对象
	Owner_sessionBus, err := dbus.SessionBus()
	obj2 := &Obj{}
	err = Owner_sessionBus.Export(obj2, "/p1/p2/p3", "p1.p2.p3")
	_, err = Owner_sessionBus.RequestName("p1.p2.p3", 0)
	log.Print("names:", Owner_sessionBus.Names())
	select {
	}


	for {
		err = Owner_sessionBus.Emit("/p1/p2/p3", "p1.p2.p3.Signal1", "arg1", 2)
		time.Sleep(2*time.Second)
	}


	//
	//sessionBus, err := dbus.SessionBus()
	//dockObj := dock.NewDock(sessionBus)
	//ok, err := dockObj.RequestDock(0, "/usr/share/applications/deepin-editor.desktop", -1)
	//log.Println("ok:", ok)







}
