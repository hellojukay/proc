package environ

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// GetEnv Get process enviroment list
func GetEnv(pid int) (env map[string]string, err error) {
	file := fmt.Sprintf("/proc/%d/environ", pid)
	fh, err := os.Open(file)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("can not open file %s:%s", file, err))
	}
	env = make(map[string]string)
	var reader = bufio.NewReader(fh)
	for {
		line, err := reader.ReadSlice(byte(0))
		if err != nil {
			break
		}

		keyValue := strings.Split(string(line), "=")
		if len(keyValue) < 2 {
			env[keyValue[0]] = ""
		} else {
			env[keyValue[0]] = keyValue[1]
		}
	}
	return env, nil
}
