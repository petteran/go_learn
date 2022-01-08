package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	for i := 0; i < len(os.Args); i++ {
		sep := " "
		sep += strconv.Itoa(i) + sep + os.Args[i]
		fmt.Println(sep)

	}

}
