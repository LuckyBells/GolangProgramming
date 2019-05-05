package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	s, sep := "",""
	for index, arg := range os.Args{
		s += sep +strconv.Itoa(index)+":"+arg
		sep = " "
	}
	fmt.Println(s)
}
