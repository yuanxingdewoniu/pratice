package serve

import (
	"deepin-security-enhance/pkg/tools"
	"fmt"

	"pkg.deepin.io/lib/dbusutil"
)


// dbus后台服务执行程序
var (
	SecurityEnhanceDaemon = "deepin-security-enhance"
	RemovableStorageDeviceDaemon = "uos-usb-storage-daemon"
)


// dbus服务信息
const (
	SecurityEnhanceName = "com.deepin.daemon.SecurityEnhance"  // @Name: 	等保名称
	SecurityEnhancePath = "/com/deepin/daemon/SecurityEnhance" // @Path:	等保地址
	SecurityEnhanceIFC  = "com.deepin.daemon.SecurityEnhance"  // @IFC:	等保接口名

	RemovableStorageDeviceName = "com.deepin.daemon.RemovableStorageDevice"  // @Name: 	USB存储管控名称
	RemovableStorageDevicePath = "/com/deepin/daemon/RemovableStorageDevice" // @Path:	USB存储管控地址
	RemovableStorageDeviceIFC  = "com.deepin.daemon.RemovableStorageDevice"  // @IFC:	USB存储管控接口名
)

// dbus服务对象
type SecurityEnhance struct {
	methods *struct {
		Enable          func() `in:"enable,deleteadm,sysadmpasswd,secuadmpasswd,audadmpasswd"`
		Status          func() `out:"status"`
		SetLabel        func() `in:"labeltype,devicetype,label"`
		GetLabel        func() `in:"labeltype" out:"labels"`
		GetSEUserByName func() `in:"linuxuser" out:"seuser,level"`

	}

	signals *struct {
		Receipt struct {
			result bool
		}

	}
}


// dbus服务对象
type RemovableStorageDevice struct {
	methods *struct {

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
	conn                   *dbusutil.Service
	SecurityEnhance        *SecurityEnhance
	RemovableStorageDevice *RemovableStorageDevice
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
	securityenhance := &SecurityEnhance{}
    removablestoragedevice := &RemovableStorageDevice{}

	dbusService = &Service{conn: srv, SecurityEnhance:securityenhance,RemovableStorageDevice:removablestoragedevice }
	return dbusService, nil
}

// 外部调用 初始化注册指定服务对象
func (srv *Service) Init(daemonName string) error {
	switch daemonName {
	case SecurityEnhanceDaemon:
		return srv.initSecurityEnhanceDbus()
	case RemovableStorageDeviceDaemon:
		return srv.initRemovableStorageDeviceDbus()
	default:
		return fmt.Errorf("Failed to init deamon")
	}
}


// 初始化 等保管控的服务对象
func (srv *Service) initSecurityEnhanceDbus() error {
	err := srv.conn.Export(SecurityEnhancePath, GetService().SecurityEnhance)
	if err != nil {
		return err
	}
	return srv.conn.RequestName(SecurityEnhanceName)
}

// 初始化 注册USB设备管控的服务对象
func (srv *Service) initRemovableStorageDeviceDbus() error {
	err := srv.conn.Export(RemovableStorageDevicePath, GetService().RemovableStorageDevice)
	if err != nil {
		return err
	}
	return srv.conn.RequestName(RemovableStorageDeviceName)
}


// 获取 dbus对象 ifc名称
func (r *SecurityEnhance) GetInterfaceName() string {
	return SecurityEnhanceIFC
}

func (r *RemovableStorageDevice) GetInterfaceName() string {
	return RemovableStorageDeviceIFC
}





// 循环
func (srv *Service) Loop() {
	fmt.Println("dbus success")
	_, tools.Labels = tools.GetValue(tools.MountLabels)

	srv.conn.Wait()
}
