package main

import "fmt"

func main() {
	// standard format of initialising strings
	var s string
	s = "Hello, world!"
	fmt.Println(s)
	// using inference of type shortcut initialisation
	t := "Hello, world again!"
	fmt.Println(t)
	// escape sequences within double-quotes
	u := "Hello, \n\t\"world\" with a backslash \\"
	fmt.Println(u)
	// raw quotes usage does not require escape dequences

}
