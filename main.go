package main

import (
	"fmt"
	"github.com/hellojukay/proc/environ"
)

func main() {
	m, err := environ.GetEnv(1)
	if err != nil {
		panic(err)
	}
	for k, v := range m {
		fmt.Printf("%10s%30s\n", k, v)
	}
}
