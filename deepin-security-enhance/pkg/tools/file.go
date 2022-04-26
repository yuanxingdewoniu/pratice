package tools

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var Labels [][]string

// 获取标签值
func GetValue(filename string) (error, [][]string) {
	var err error
	var labels [][]string

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err, labels
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		var label []string
		items := strings.SplitN(line, ":", 2)
		parts := strings.SplitN(items[len(items)-1], ",", 2)
		for _, v := range parts {
			values := strings.SplitN(v, "=", 2)
			label = append(label, values[len(values)-1])
		}
		labels = append(labels, label)
	}

	return err, labels
}

// 修改标签值
func Modify(filename string, labels [][]string) error {
	info, err := os.Stat(filename)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = os.Remove(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		items := strings.SplitN(line, ":", 2)
		parts := strings.SplitN(items[len(items)-1], ",", 2)
		for j, v := range parts {
			values := strings.SplitN(v, "=", 2)
			for m, row := range labels {
				for n, column := range row {
					if strconv.Itoa(m+1) == items[0] && strconv.Itoa(n+1) == values[0] {
						if !strings.Contains(column, "\"") {
							column = "\"" + column + "\""
						}
						values[1] = column
					}
				}
				parts[j] = strings.Join(values, "=")
			}
		}
		items[1] = strings.Join(parts, ",")
		lines[i] = strings.Join(items, ":")
	}
	err = ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")), info.Mode())
	if err != nil {
		return err
	}

	return err
}

// 组合标签值
func GetLabel(device_type int) string {
	var label string
	var parts []string
	for i, value := range Labels[device_type-1] {
		if !strings.Contains(value, "\"") {
			value = "\"" + value + "\""
		}
		if value == "\"null\"" {
			continue
		}
		switch i {
		case 0:
			value = "rootcontext" + "=" + value
		case 1:
			value = "defcontext" + "=" + value
		default:
		}
		parts = append(parts, value)
	}
	label = strings.Join(parts, ",")
	return label
}
