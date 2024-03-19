// use for loop using maps 
package main 
import "fmt"
// Main function 
func main() { 
	// using maps 
	mmap := map[int]string{ 
		22:"Geeks", 
		33:"GFG", 
		44:"GeeksforGeeks", 
	} 
	for key, value:= range mmap { 
	fmt.Println(key, value) 
	} 
	
} 
