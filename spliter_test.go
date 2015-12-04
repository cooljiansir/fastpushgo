package main

import (
	"github.com/cooljiansir/fastpush"
	"os"
	"bufio"
	"fmt"
	"path/filepath"
)

func main(){
	filestr := os.Args[1]
	file,err := os.Open(filestr)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for{
		blks := fastpush.Split(r,4*1024,1)
		if len(blks) == 0{
			break
		}
		fmt.Printf("----\nhash: %x\noffset: %d\nlength: %d\n-----\n",blks[0].Hash(),blks[0].Offset(),blks[0].Length())
	}
}
