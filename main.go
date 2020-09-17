package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hellojukay/proc/cmdline"
	"github.com/hellojukay/proc/environ"
)

var enableEnv = false
var enableTcp = false
var enableFile = false
var enableCmd = false
var pid = 0

func init() {
	flag.BoolVar(&enableEnv, "e", false, "show process environment list")
	flag.BoolVar(&enableTcp, "t", false, "show process tcp infomation")
	flag.BoolVar(&enableCmd, "c", false, "show process command line")
	flag.IntVar(&pid, "p", 1, "process pid")
	flag.Parse()
}
func main() {
	if enableEnv {
		m, err := environ.GetEnv(pid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}
		environ.PrintEnv(m)
		fmt.Print("\n")
	}
	if enableCmd {
		c, err := cmdline.GetCmdLine(pid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}
		cmdline.PrintCmdLine(c)
	}
}
