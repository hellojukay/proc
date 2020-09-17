package cmdline

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CmdLine linux process command line
type CmdLine struct {
	Exe     string
	Cwd     string
	CmdLine string
}

func GetCmdLine(pid int) (CmdLine, error) {
	var cmd CmdLine
	cwd, err := readCwd(pid)
	if err != nil {
		return cmd, err
	}
	exe, err := readExe(pid)
	if err != nil {
		return cmd, err
	}
	cmdline, err := readCmdLine(pid)
	if err != nil {
		return cmd, err
	}
	cmd.CmdLine = cmdline
	cmd.Exe = exe
	cmd.Cwd = cwd
	return cmd, nil
}

func readExe(pid int) (string, error) {
	return os.Readlink(fmt.Sprintf("/proc/%d/exe", pid))
}

func readCwd(pid int) (string, error) {
	return os.Readlink(fmt.Sprintf("/proc/%d/cwd", pid))
}

func readCmdLine(pid int) (string, error) {
	fh, err := os.Open(fmt.Sprintf("/proc/%d/cmdline", pid))
	if err != nil {
		return "", err
	}
	defer fh.Close()
	var reader = bufio.NewReader(fh)
	var lines []string
	for {
		line, err := reader.ReadBytes(byte(0))
		if err != nil {
			break
		}
		lines = append(lines, string(line))
	}
	return strings.Join(lines, " "), nil
}

func PrintCmdLine(c CmdLine) {
	fmt.Printf("%-10s%s\n", "cmdline", c.CmdLine)
	fmt.Printf("%-10s%s\n", "exe", c.Exe)
	fmt.Printf("%-10s%s\n", "workspace", c.Cwd)
}
