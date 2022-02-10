package main

// zhu
import (
	"github.com/godbus/dbus/v5"
	"github.com/snapcore/snapd/dbusutil"
)

type Obj struct {
	methods *struct {
		GetString func() `out:"result"`
	}
}

func (o *Obj) GetString() (string, *dbus.Error) {
	return "object", nil
}

func (o *Obj) GetInterfaceName() string {
	return "p1.p2.p3"
}

func main() {
	service, err := dbusutil.MockConnections()
	obj := &Obj{}
	err = service.Export("/p1/p2/p3", obj)
	err = service.RequestName("p1.p2.p3")
	service.Wait()
}
