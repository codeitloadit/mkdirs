package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var err error
	var path string
	var perm uint64
	var strPerm string = "755" // Default for mkdir

	args := os.Args[1:]

	switch len(args) {
	case 0:
		showUsage()
	case 1:
		path = args[0]
	case 2:
		path = args[0]
		strPerm = args[1]
		if len(strPerm) > 4 {
			showUsage()
		}
	default:
		fmt.Println("Unexpected arguments: ", args[2:])
		showUsage()
	}

	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			showUsage()
		}
	}

	perm, err = strconv.ParseUint(strPerm, 8, 32)
	fmt.Println(perm)
	if err != nil {
		fmt.Printf("perm [%v] must be an integer [0...7777]", args[1])
	}

	err = os.MkdirAll(path, os.FileMode(perm))
	if err != nil {
		panic(err.Error())
	}
}

func showUsage() {
	fmt.Println("usage: mkdirs path [perm]")
	os.Exit(0)
}
