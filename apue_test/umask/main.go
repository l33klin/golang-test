package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func main() {
	unix.Umask(0)
	_, err := os.Create("foo.tmp")
	if err != nil {
		fmt.Println("Create Error")
	}
	unix.Umask(unix.S_IRGRP | unix.S_IWGRP | unix.S_IROTH | unix.S_IWOTH)
	_, err2 := os.Create("bar.tmp")

	if err2 != nil {
		fmt.Println("Create Error")
	}

	unix.Umask(0022) // for test docker default umask
	_, err3 := os.Create("docker_default.tmp")

	if err3 != nil {
		fmt.Println("Create Error")
	}
}
