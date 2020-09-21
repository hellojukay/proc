package cmdline

import (
	"fmt"
	"io/ioutil"
)

/// CountThread count process threads
func CountThread(pid int) (int, error) {
	var path = fmt.Sprintf("/proc/%d/task/", pid)
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		return 0, err
	}
	return len(infos), nil
}
