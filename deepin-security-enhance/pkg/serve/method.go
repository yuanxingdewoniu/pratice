package serve

// #cgo pkg-config: libselinux
// #include <selinux/selinux.h>
// #include <stdlib.h>
import "C"

import (
	"deepin-security-enhance/pkg/tools"
	"fmt"
	"os"
	"reflect"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil"
)

var (
	trustedProcess = []string{"/usr/bin/uos-usb-guarddeepin-defender",
		"/usr/sbin/deepin-defender-antiav", "/usr/bin/deepin-defender-daemonservice",
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
func (s *SecurityEnhance) Enable(sender dbus.Sender, enable bool, deleteadm bool, sysadmpasswd string, secadmpasswd string, audadmpasswd string) *dbus.Error {
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

	GetService().conn.Emit(s, "Receipt", result)

	return dbusutil.ToError(err)
}

// 获取等保三级状态
func (s *SecurityEnhance) Status() (string, *dbus.Error) {
	var state string
	var err error
	err, state = tools.GetEnhanceStatus()
	if err != nil {
		err = fmt.Errorf("Failed to get enhances status\n")
	}
	return state, dbusutil.ToError(err)
}

func (r *RemovableStorageDevice) IsServiceAvailable() (bool, *dbus.Error) {
	var err error
	var state bool

	state = true

	err = fmt.Errorf("GetDeviceList test \n ")

	return state, dbusutil.ToError(err)
}

func (r *RemovableStorageDevice) GetDeviceList() (string, *dbus.Error) {
	var err error
	var dev_list string
	dev_list = "1111111111"

	err = fmt.Errorf("GetDeviceList test \n ")

	return dev_list, dbusutil.ToError(err)
}

func (r *RemovableStorageDevice) GetWhiteList() (string, *dbus.Error) {

	var err error
	var dev_list string

	dev_list = "GetWhileList"

	err = fmt.Errorf("GetDeviceList test \n ")
	return dev_list, dbusutil.ToError(err)

}

func (r *RemovableStorageDevice) AddWhiteList(info string) *dbus.Error {

	var err error

	err = fmt.Errorf("GetDeviceList test \n ")

	return dbusutil.ToError(err)

}

func (r *RemovableStorageDevice) DeleteWhiteList(info string) *dbus.Error {
	var err error

	err = fmt.Errorf("GetDeviceList test \n ")

	return dbusutil.ToError(err)

}

func (r *RemovableStorageDevice) ModifyWhiteList(info string) *dbus.Error {
	var err error

	err = fmt.Errorf("ModifyWhiteList test \n ")

	return dbusutil.ToError(err)
}
