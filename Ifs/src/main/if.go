package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("Script started... \n... \n...\n")
	test()

}

func test(){
	fmt.Println("Test one: ")
	if len(os.Args) != 2 {
		os.Exit(6)
	}
	fmt.Println("It's over ", os.Args[1])
	//prefered response is, of course, 9000
}

