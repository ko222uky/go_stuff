package main

import ( 
	"fmt"
	"strings"
)
/*
TYPE DECLARATIONS
*/
// using type aliases, we define two types
type martian struct{} // type martian, a struct
type laser int // type laser with the same underling type 'int'


// Define an interface which can be assigned to variables with the talk() method
type talker interface {
    	talk() string
}

/*
METHOD DECLARATIONS
*/

// Always define methods outside of the main function.
// method for the martian type
func (m martian) talk() string {
	return "nack nack"
}
// method for the laser type
func (l laser) talk() string {
   	return strings.Repeat("pew ", int(l))
}

func shout(t talker) {
	// this function can be used with any variable that satisfies the talker interface
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}



func main() {
	// initialize variable t (which is type talker)
	var t talker
	t = laser(5)
	fmt.Println(t.talk())
	shout(t)

	t = martian{}
	fmt.Println(t.talk())
	shout(t)
}


