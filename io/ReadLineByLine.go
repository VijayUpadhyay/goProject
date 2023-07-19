package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// os.Open() opens specific file in read-only mode and this return
	// a pointer of type os.
	file, err := os.Open("./sample.txt")
	v, _ := file.Stat()
	fmt.Printf("File Details: %+v\n", v)
	if err != nil {
		log.Fatalf("failed to open: %+v\n", err)
	}

	// The bufio.NewScanner() function is called in which the object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each new line using the bufio.Scanner.Scan() method.
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	// The method os.File.Close() is called on the os.File object to close the file
	if err = file.Close(); err != nil {
		return
	}
	// and then a loop iterates through and prints each of the slice values.
	for _, each_ln := range text {
		fmt.Println(each_ln)
	}
}
