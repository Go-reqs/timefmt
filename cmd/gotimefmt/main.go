package main

import (
	"fmt"
	"github.com/Go-reqs/timefmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("example:\n\t%s %s", os.Args[0], `"%y-%m-%d"`)
	}
	fmt.Println(timefmt.FromPosix(os.Args[1]))
}
