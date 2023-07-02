package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"io"
)

func main()  {

	fmt.Println("Going to read file")
	data,err:=ioutil.ReadFile("C:/Users/vupadhya/Documents/BuzzWords.txt")
	check(err)
	fmt.Println(string(data))

	fmt.Println("================================Open file and Read============================================")
	//fmt.Println(os.Environ())
	f,err2:=os.Open("C:/Users/vupadhya/Documents/Full_Forms.txt")
	check(err2)
	fmt.Printf("Data variable having value: %v",f)

	sliceEx:=make([]byte,4)
	n,err3:= f.Read(sliceEx)
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
	fmt.Printf("Number of bytes: %d and Data is: %s",n,string(sliceEx))

	fmt.Println("\n================================Implement Seek1============================================")
	o2,err4:=f.Seek(4,0)
	check(err4)
	b3:= make([]byte,2)
	n2,err5:=f.Read(b3)
	check(err5)
	fmt.Printf("\n%d bytes , Offset: %d, Data: %s",n2,o2,string(b3))
	fmt.Println("\n================================Implement Seek2============================================")
	o3,err6:=f.Seek(6,0)
	check(err6)
	n4,err7:= f.Read(b3)
	check(err7)
	fmt.Printf("\n%d bytes, Offset: %d, Data is: %s",n4,o3,string(b3))

	fmt.Println("\n================================Implement Seek3============================================")
	o4,err8:=f.Seek(8,0)
	check(err8)
	n5,err9:=io.ReadAtLeast(f,b3,1)
	check(err9)
	fmt.Printf("%d bytes, Offset: %d, Data: %s",n5,o4,string(b3))

	fmt.Println("\n============Close the opened file==========")
	f.Close()
}

func check(e error)  {
	if e!= nil {
		panic(e)
	}
}
