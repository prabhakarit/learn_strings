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
	n := `Hello, 
		"world" with a backslash \`
	fmt.Println(n)
	// using string lenght and concatenation using + operator
	s = "Hello🐹, World"
	s1 := s[1:5]
	s2 := s[:1]
	s3 := s[2:]
	fmt.Println(s, len(s), s1, len(s1), s2, len(s2), s3, len(s3))
	// Using rune to assign unicode values to print into string
	s = "Hello, "
	var r rune = 127757
	r1 := '🐹'
	s = s + string(r) + string(r1)
	fmt.Println(s)
	// Using arrays with string
	// Array lenght is used to define the type as well
	// instead type Slice uses array underneath
	var vals [3]int
	vals[0] = 2
	vals[1] = 4
	vals[2] = 6
	fmt.Println(vals, vals[0], vals[1], vals[2])
}
