package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "../data/internet_20161227_00.txt"
	fin, err := os.Open(userFile)
	defer fin.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, _ := fin.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
