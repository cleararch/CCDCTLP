package main

import (
	"fmt"
)

func main() {
	a := Unpack_deb("/home/jack/下载/hello.deb", "/home/jack/下载/")
	fmt.Println(a)
}
