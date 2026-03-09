// Go Program

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome := "Welcome to User Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza")

	// comma OK syntax or also called as error ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Printf("Thanks for rating our pizza %v", input)

}
