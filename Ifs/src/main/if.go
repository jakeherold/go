package main

import (
	"fmt"
	//"os"
)

// func main(){
// 	fmt.Println("Script started... \n... \n...\n")
// 	x := 9000
// 	test(x)

// }

// func test(pwr int){
// 	// power:= 9000
// 	// var power int
// 	// power = 9000
// 	fmt.Println("It's over ", pwr)
// }

func main(){
	fmt.Println(getPower("9000", 10000, true))
}

func getPower(x string, y int, z bool) (string, int) {

	return x, y

}