package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

// Administrator user of separation of powers
var (
	MountLabels       = "/etc/deepin-security/label/Labels"
	enahnceStatus     = "/etc/deepin-security/state/Enable"
	systemUserProfile = "/etc/passwd"
	configureScript   = "/etc/deepin-security/script/security_enhance.sh"
)

// 获取三权管理员用户是否存在
func IsAdministratorsExist() bool {
	var ret bool
	var administrators = []struct {
		admin string
		exist bool
	}{
		{"sysadm", false},
		{"secadm", false},
		{"audadm", false},
	}

	data, err := ioutil.ReadFile(systemUserProfile)
	if err != nil {
		return ret
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		for i, items := range administrators {
			if strings.Contains(line, items.admin) {
				administrators[i].exist = true
			}
		}
	}

	ret = administrators[0].exist && administrators[1].exist && administrators[2].exist
	return ret
}

// 执行等保三级配置脚本
func RunConfigureScript(enable bool, deleteadm bool, sysadmpasswd string, secadmpasswd string, audadmpasswd string) error {
	var status string
	var delete string
	if enable {
		status = "enable"
		ModifyEnhanceStatus("opening")
	} else {
		status = "disable"
		ModifyEnhanceStatus("closing")
	}

	if deleteadm {
		delete = "true"
	} else {
		delete = "false"
	}

	sysadmpasswd = strings.Replace(sysadmpasswd, "$", "\\$\\", -1)
	secadmpasswd = strings.Replace(secadmpasswd, "$", "\\$\\", -1)
	audadmpasswd = strings.Replace(audadmpasswd, "$", "\\$\\", -1)

	cmdline := fmt.Sprintf("%s %s %s %s %s %s", configureScript, status, delete, sysadmpasswd, secadmpasswd, audadmpasswd)

	cmd := exec.Command("/bin/bash", "-c", cmdline)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, string(out))
	}
	return errorHandler(string(out))
}

// 错误码处理
func errorHandler(rval string) error {
	var ret error
	switch strings.Replace(rval, "\n", "", -1) {
	case "0":
		ret = nil
	case "-100":
		ret = fmt.Errorf("Configuration / restore authentication failed :%s\n", rval)
	case "-200":
		ret = fmt.Errorf("Configuration / restore of autonomous access control failed :%s\n", rval)
	case "-300":
		ret = fmt.Errorf("Failed to configure / restore tags and enforce access control :%s\n", rval)
	case "-400":
		ret = fmt.Errorf("Configuration / restore security audit failed :%s\n", rval)
	case "-500":
		ret = fmt.Errorf("Failed to configure / restore data integrity :%s\n", rval)
	case "-600":
		ret = fmt.Errorf("Failed to configure / restore data confidentiality :%s\n", rval)
	case "-700":
		ret = fmt.Errorf("Failed to configure / restore network security :%s\n", rval)
	case "-800":
		ret = fmt.Errorf("Configuration / restore run security failed :%s\n", rval)
	case "-900":
		ret = fmt.Errorf("Failed to configure / restore resource utilization :%s\n", rval)
	case "-1000":
		ret = fmt.Errorf("Failed to configure / restore user login access control :%s\n", rval)
	case "-1100":
		ret = fmt.Errorf("Failed to configure / restore trusted metrics :%s\n", rval)
	default:
		ret = fmt.Errorf("%s\n", rval)
	}
	return ret
}

// 获取等保三级状态
func GetEnhanceStatus() (error, string) {
	var status string = "close"
	var statusList = []string{"open", "close", "opening", "closing"}
	data, err := ioutil.ReadFile(enahnceStatus)
	if err != nil {
		return err, status
	}
	if !isInList(string(data), statusList) {
		return fmt.Errorf("Invalid error status\n"), status
	}
	return err, string(data)
}

// 修改等保三级状态
func ModifyEnhanceStatus(i_state string) error {
	f, err := os.OpenFile(enahnceStatus, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0700)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("Failed to modify the third-level reinforcement state\n")
	}

	var content []byte = []byte(i_state)
	_, err = f.Write(content)
	return err
}

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
