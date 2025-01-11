package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io"
	"time"
	"strings"
	"bufio"
)

func catfile(filename string, searchterm ... string) {
	/*
	This function will be used for goroutines.
	The goal is to pass it a string representing a filename.
	If that string has a .txt extension, we open the file, pass a pointer to the file to io.Copy().
	This will output the file's contents into the terminal
	*/
	// fmt.Printf("Search term: %s\n", searchterm[0]) // for testing
	// using variadic function to simulate keyword arguments	
	if len(searchterm) == 1 {
		// if theres exactly one searchterm, enter this block

		if filepath.Ext(filename) != ".txt" {
			return
		} else if filepath.Ext(filename) == ".txt" {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer file.Close() // This will execute when function scope returns
			// create scanner
			scanner := bufio.NewScanner(file)
			// read our file line by line
			for scanner.Scan() {
				if strings.Contains(scanner.Text(), searchterm[0]) {
					fmt.Printf("'%s' contains the search word '%s'.\n", filename, searchterm[0])
					return	
				}
			}
		}
	} else {

		if filepath.Ext(filename) != ".txt" {
			return
		} else if filepath.Ext(filename) == ".txt" {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer file.Close() // This will execute when function scope returns
		
			// Pass our filepointer to io.Copy, as well as the output
			_, err = io.Copy(os.Stdout, file)
			if err != nil {
				fmt.Println("\nError reading file:", err)
			}
		}
	}
}

func main() {
	// If we have no arguments, inform user.
	if len(os.Args) == 1  {
		fmt.Println("error: No arguments passed.\n")
		// note: the first element in the os.Args is the program name
		return
	}
	// the first argument is the program name, which we can either ignore or print for fun
	fmt.Printf("Program: %s\n", os.Args[0])

	// The rest of the os.Args list will contain our passed arguments
	for _, filename := range os.Args[2:] {
		// iterate through our list of filenames passed
		// only use the .txt extensions
		// Basic steps for our recursion are:
		// For when this function is modified to be recursive:
		// 1) if it is not a directory nor a .txt file, skip
		// 2) output .txt file to stdout
		// 3) if it is a directory, call recursive function
		go catfile(filename, os.Args[1]) // Providing exactly one search term (here, os.Args[1]) executes code as required in ch2 exercise 2.
	}

	// At this point, we should wait since the main() will end before all of our goroutines complete.
	time.Sleep(2 * time.Second) // wait for 2 seconds
}
