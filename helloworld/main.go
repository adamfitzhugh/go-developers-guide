/*
Think of a package as a project/workspace. 
'main' is used to declare an executable package. 
Any other name is known as a reusable package
*/
package main

// Give my package access to all the functionality that is contained under "fmt" package
import "fmt"

// Declare functions. Tell Go to do things
func main() {
	fmt.Println("Hi there!")
}