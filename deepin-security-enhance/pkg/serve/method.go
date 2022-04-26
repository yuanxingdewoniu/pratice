package serve

// #cgo pkg-config: libselinux
// #include <selinux/selinux.h>
// #include <stdlib.h>
import "C"

import (
	"deepin-security-enhance/pkg/tools"
	"fmt"
	"unsafe"
    
	"os" 
	"reflect"
	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil"
)


var (
	trustedProcess  = []string{"/usr/bin/uos-usb-guarddeepin-defender", 
	"/usr/sbin/deepin-defender-antiav","/usr/bin/deepin-defender-daemonservice",
	"/usr/bin/deepin-defender-datainterface",	
	"/usr/bin/deepin-MonitorNetFlow", 
	} // 可信应用进程
)  

// 判断是否在列表中
func isInList(item interface{}, list interface{}) bool {
        sVal := reflect.ValueOf(list)
        kind := sVal.Kind()
        if kind == reflect.Slice || kind == reflect.Array {
                for i := 0; i < sVal.Len(); i++ {
                        if sVal.Index(i).Interface() == item {
                                return true
                        }
                }
                return false
        }
        return false
}


// 检查是否为可信应用
func isTruestedProcess(sender string) bool {
        pid, err := GetService().conn.GetConnPID(sender)
        if err != nil {
                return false
        }

        cmdline := fmt.Sprintf("/proc/%d/exe", pid)
        processName, err := os.Readlink(cmdline)
        if err != nil {
                return false
        }

        return isInList(processName, trustedProcess)
}

// 等保三级开关
func (o *Object) Enable(sender dbus.Sender, enable bool, deleteadm bool, sysadmpasswd string, secadmpasswd string, audadmpasswd string) *dbus.Error {
	var err error
	var state string
	var result bool


    //限制只有可信进程才能调用接口
	if !isTruestedProcess(string(sender)) {
		err = fmt.Errorf("No permission to call this method %s", string(sender))
		goto end
	}
     
	if enable && tools.IsAdministratorsExist() && !deleteadm {
		err = fmt.Errorf("Three rights separation user already exists\n")
		goto end
	}

	err, state = tools.GetEnhanceStatus()
	if err != nil {
		err = fmt.Errorf("Failed to get enhances status\n")
		goto end
	}

	if (enable && (state == "open" || state == "opening")) || (!enable && (state == "close" || state == "closing")) {
		err = fmt.Errorf("The third level state of Level Protection has not changed, please check the input parameters\n")
		goto end
	}

	err = tools.RunConfigureScript(enable, deleteadm, sysadmpasswd, secadmpasswd, audadmpasswd)
	if err != nil {
		tools.ModifyEnhanceStatus(state)
		goto end
	}

	if enable {
		err = tools.ModifyEnhanceStatus("open")
	} else {
		err = tools.ModifyEnhanceStatus("close")
	}
end:
	if err != nil {
		result = false
	} else {
		result = true
	}

	GetService().conn.Emit(o, "Receipt", result)

	return dbusutil.ToError(err)
}

// 获取等保三级状态
func (o *Object) Status() (string, *dbus.Error) {
	var state string
	var err error
	err, state = tools.GetEnhanceStatus()
	if err != nil {
		err = fmt.Errorf("Failed to get enhances status\n")
	}
	return state, dbusutil.ToError(err)
}

// 设置对应标签类型、设备类型下标签值
func (o *Object) SetLabel(label_type int, device_type int, label string) *dbus.Error {
	var err error
	switch label_type {
	case 1, 2:
	default:
		err = fmt.Errorf("No such mount match type exists：%d", label_type)
		goto end
	}
	switch label_type {
	case 1, 2:
	default:
		err = fmt.Errorf("This device match type does not exist：%d", device_type)
		goto end
	}
	tools.Labels[device_type-1][label_type-1] = label
	err = tools.Modify(tools.MountLabels, tools.Labels)
end:
	return dbusutil.ToError(err)
}

// 获取设备类型下匹配两种挂载类型的标签值
func (o *Object) GetLabel(device_type int) (string, *dbus.Error) {
	var labels string
	var err error

	switch device_type {
	case 1, 2:
		labels = tools.GetLabel(device_type)
	default:
		err = fmt.Errorf("This device match type does not exist：%d", device_type)
	}
	return labels, dbusutil.ToError(err)
}

// 检索与给定 Linux 用户名关联的 SELinux 用户名和安全级别
func (o *Object) GetSEUserByName(linuxuser string) (string, string, *dbus.Error) {
	var err error
	var seuser, level string

	var c_seuser, c_level *C.char
	var c_linuxuser = C.CString(linuxuser)

	defer C.free(unsafe.Pointer(c_linuxuser))

	ret := C.getseuserbyname(c_linuxuser, &c_seuser, &c_level)
	if ret != C.int(0) {
		err = fmt.Errorf("Failed to getseuserbyname")
		goto end
	}

	seuser = C.GoString(c_seuser)
	level = C.GoString(c_level)

	C.free(unsafe.Pointer(c_seuser))
	C.free(unsafe.Pointer(c_level))
end:
	return seuser, level, dbusutil.ToError(err)
}
