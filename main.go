package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hellojukay/proc/cmdline"
	"github.com/hellojukay/proc/environ"
	"github.com/hellojukay/proc/fd"
	"github.com/hellojukay/proc/network"
)

var (
	enableEnv  = false
	enableNet  = false
	enableFile = false
	enableCmd  = false
	pid        = 0
)

func init() {
	flag.BoolVar(&enableEnv, "e", false, "show process environment list")
	flag.BoolVar(&enableNet, "n", false, "show process network connection infomation")
	flag.BoolVar(&enableCmd, "c", false, "show process command line")
	flag.BoolVar(&enableFile, "f", false, "list process open files")
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
		fmt.Print("\n")
	}
	if enableFile {
		files, err := fd.ReadFd(pid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		for _, file := range files {
			fmt.Printf("%-5s%-s\n", file.Fd, file.Link)
		}
	}
	if enableNet {
		netInfos, err := network.ReadNetInfo(pid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		network.PrintNetInfo(netInfos)
		fmt.Print("\n")
	}
}
