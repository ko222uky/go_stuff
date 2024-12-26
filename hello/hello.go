package hello // package to which this file belongs
/*
To import this package, I needed to specify in this local package's mod file the following:

replace example/hello => ../hello


*/

import "fmt" // import statement for fmt package, which is used for formatted I/O


func Hello() {
	// exportable functions start with a capital letter
	fmt.Println("Hello, World!")
}