package serve

import (
	"deepin-security-enhance/pkg/tools"
	"fmt"

	"pkg.deepin.io/lib/dbusutil"
)

// dbus服务信息
const (
	dbusName = "com.deepin.daemon.SecurityEnhance"  // @Name: 	名称
	dbusPath = "/com/deepin/daemon/SecurityEnhance" // @Path:	地址
	dbusIFC  = "com.deepin.daemon.SecurityEnhance"  // @IFC:	接口名
)

// dbus服务对象
type Object struct {
	methods *struct {
		Enable          func() `in:"enable,deleteadm,sysadmpasswd,secuadmpasswd,audadmpasswd"`
		Status          func() `out:"status"`
		SetLabel        func() `in:"labeltype,devicetype,label"`
		GetLabel        func() `in:"labeltype" out:"labels"`
		GetSEUserByName func() `in:"linuxuser" out:"seuser,level"`

		// usb 存储接口
		GetDeviceList     func() `out:"dev_list"` //??
		SetGlobalPermMode func() `in:"mode" out:"result"`
		GetGlobalPermMode func() `out:"mode"`
		AddWhiteList      func() `in:"info"`
		DeleteWhiteList   func() `in:"info"`
		ModifyWhiteList   func() `in:"info"`
		GetWhiteList      func() `out:"dev_list"` //??
	}

	signals *struct {
		Receipt struct {
			result bool
		}

		GlobalPermModeChanged struct {
			mode int
		}

		WhitelistChanged struct {
			mode int
			info string //设备的信息
		}

		DeviceAdded struct {
			info string //设备的信息
		}

		DeviceRemove struct {
			info string //设备的信息
		}
	}
}

// dbus 对象
type Service struct {
	conn   *dbusutil.Service
	Object *Object
}

// 实例化 dbus
var dbusSrv *Service

// 获取初始化的 dbus 对象,不存在就新建
func GetService() *Service {
	if dbusSrv != nil {
		return dbusSrv
	}
	var err error
	dbusSrv, err = newService()
	if err != nil {
		panic(err)
	}
	return dbusSrv
}

// 新建对象
func newService() (*Service, error) {
	var dbusService *Service
	srv, err := dbusutil.NewSystemService()
	if err != nil {
		return nil, fmt.Errorf("new system service is error:%s\n", err)
	}
	object := &Object{}
	dbusService = &Service{conn: srv, Object: object}
	return dbusService, nil
}

// 外部调用
func (srv *Service) Init() error {
	return srv.initDBus()
}

// 外调
func (srv *Service) initDBus() error {
	err := srv.conn.Export(dbusPath, GetService().Object)
	if err != nil {
		return err
	}
	return srv.conn.RequestName(dbusName)
}

// 获取 dbus对象 ifc名称
func (o *Object) GetInterfaceName() string {
	return dbusIFC
}

// 循环
func (srv *Service) Loop() {
	fmt.Println("dbus success")
	_, tools.Labels = tools.GetValue(tools.MountLabels)

	srv.conn.Wait()
}
