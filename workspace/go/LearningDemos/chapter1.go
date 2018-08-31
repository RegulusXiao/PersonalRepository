package main

import (
	"fmt"
	"os"
)

func main()  {
	// first test, hello wrold
	fmt.Println("Hello, 世界")

	// for range, if we don't use index, may be error!Placeholder is charactor '_'
	for index, argv := range os.Args {
		fmt.Println(index, argv)
	}

	fmt.Println("test")
}