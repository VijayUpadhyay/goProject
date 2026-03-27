package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// ============================================================================
// Example 1: buffIO.go
// ============================================================================

func example1_buffio() {
	scanner := bufio.NewScanner(strings.NewReader(`one
two
three
four
`))
	var (
		text strings.Builder
		n    int
	)
	for scanner.Scan() {
		n++
		text.WriteString(fmt.Sprintf("%d. %s\n", n, scanner.Text()))
	}
	fmt.Print(text.String())
	//os.Stdout.Write(text.String())
	// Output:
	// 1. One
	// 2. Two
	// 3. Three
	// 4. Four
}

// ============================================================================
// Example 2: ListFilesNonRecursive.go
// ============================================================================

func example2_listfilesnonrecursive() {
	home, err := os.UserHomeDir()
	fmt.Println("Home: ", home)
	if err != nil {
		log.Fatal(err)
	}
	files, err := os.ReadDir("C:\\Users\\DELL\\eclipse\\jee-2022-12\\eclipse")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if !f.IsDir() {
			/*			fInfo, _ := f.Info()
						fmt.Println(fInfo.)*/
			fmt.Println(f.Name())
		}
	}
}

// ============================================================================
// Example 3: ReadData.go
// ============================================================================

func example3_readdata() {
	r := strings.NewReader("Vijay Upadhyay")
	b := make([]byte, 3)

	fmt.Println()

	for {
		n, err := r.Read(b)
		fmt.Printf("n: %v, err: %v, b: %v, ", n, err, b)
		fmt.Printf(" b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}

	}
}

// ============================================================================
// Example 4: ReadFile.go
// ============================================================================

func example4_readfile() {

	fmt.Println("Going to read file")
	data, err := ioutil.ReadFile("C:/Users/vupadhya/Documents/BuzzWords.txt")
	check(err)
	fmt.Println(string(data))

	fmt.Println("================================Open file and Read============================================")
	//fmt.Println(os.Environ())
	f, err2 := os.Open("C:/Users/vupadhya/Documents/Full_Forms.txt")
	check(err2)
	fmt.Printf("Data variable having value: %v", f)

	sliceEx := make([]byte, 4)
	n, err3 := f.Read(sliceEx)
	check(err3)
	/*for {
		n,err3:= f.Read(sliceEx)
		check(err3)
		fmt.Printf("n: %v, err: %v, sliceEx: %v, ",n,err,sliceEx)
		fmt.Printf(" sliceEx[:n] = %q\n",sliceEx[:n])
		if err3==io.EOF{
			break
		}
	}*/
	fmt.Println("Print Data")
	fmt.Printf("Number of bytes: %d and Data is: %s", n, string(sliceEx))

	fmt.Println("\n================================Implement Seek1============================================")
	o2, err4 := f.Seek(4, 0)
	check(err4)
	b3 := make([]byte, 2)
	n2, err5 := f.Read(b3)
	check(err5)
	fmt.Printf("\n%d bytes , Offset: %d, Data: %s", n2, o2, string(b3))
	fmt.Println("\n================================Implement Seek2============================================")
	o3, err6 := f.Seek(6, 0)
	check(err6)
	n4, err7 := f.Read(b3)
	check(err7)
	fmt.Printf("\n%d bytes, Offset: %d, Data is: %s", n4, o3, string(b3))

	fmt.Println("\n================================Implement Seek3============================================")
	o4, err8 := f.Seek(8, 0)
	check(err8)
	n5, err9 := io.ReadAtLeast(f, b3, 1)
	check(err9)
	fmt.Printf("%d bytes, Offset: %d, Data: %s", n5, o4, string(b3))

	fmt.Println("\n============Close the opened file==========")
	f.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// ============================================================================
// Example 5: ReadLineByLine.go
// ============================================================================

func example5_readlinebyline() {

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

// ============================================================================
// Example 6: WriteFile.go
// ============================================================================

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func example6_writefile() {
	fmt.Println("Open file")
	d1 := []byte("Hello Vijay!!\n\nHow are you?")
	err := ioutil.WriteFile("C:/Users/vupadhya/Documents/BuzzWordsfile.txt", d1, 0644)
	checkError(err)

	d2 := "Test Line"
	err = ioutil.WriteFile("C:/Users/vupadhya/Documents/BuzzWordsfile.txt", []byte(d2), 0644)
	checkError(err)

	d3 := "\nappending this line"
	f1, err := os.OpenFile("C:/Users/vupadhya/Documents/BuzzWordsfile.txt", os.O_APPEND, 0644)
	checkError(err)
	n_appendCheck, err := f1.WriteString(d3)
	fmt.Printf("Appending %d bytes in BuzzWordsfile.txt", n_appendCheck)

	fmt.Println("====================================2nd file=============================================")
	f, err := os.Create("C:/Users/vupadhya/Documents/BuzzWordsfile2.txt")
	checkError(err)
	defer f.Close()

	str := "New File"
	n2, err := f.Write([]byte(str))
	fmt.Printf("Wrote %d bytes", n2)
	n4, err := f.WriteString("\nCheck append")
	checkError(err)
	fmt.Printf("Appended %d bytes", n4)
	f.Sync()
	f1.Sync()

	//n3,err:= f.WriteString("add")
}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 6 examples from io")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_buffio ---")
	example1_buffio()

	fmt.Println("\n--- example2_listfilesnonrecursive ---")
	example2_listfilesnonrecursive()

	fmt.Println("\n--- example3_readdata ---")
	example3_readdata()

	fmt.Println("\n--- example4_readfile ---")
	example4_readfile()

	fmt.Println("\n--- example5_readlinebyline ---")
	example5_readlinebyline()

	fmt.Println("\n--- example6_writefile ---")
	example6_writefile()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
}
