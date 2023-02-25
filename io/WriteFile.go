package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func checkError(err error)  {
	if err!= nil{
		panic(err)
	}
}
func main()  {
	fmt.Println("Open file")
	d1:=[]byte("Hello Vijay!!\n\nHow are you?")
	err:= ioutil.WriteFile("C:/Users/vupadhya/Documents/BuzzWordsfile.txt",d1,0644)
	checkError(err)

	d2:="Test Line"
	err= ioutil.WriteFile("C:/Users/vupadhya/Documents/BuzzWordsfile.txt",[]byte(d2),0644)
	checkError(err)

	d3:= "\nappending this line"
	f1,err:=os.OpenFile("C:/Users/vupadhya/Documents/BuzzWordsfile.txt",os.O_APPEND,0644)
	checkError(err)
	n_appendCheck,err:= f1.WriteString(d3)
	fmt.Printf("Appending %d bytes in BuzzWordsfile.txt",n_appendCheck)

	fmt.Println("====================================2nd file=============================================")
	f,err := os.Create("C:/Users/vupadhya/Documents/BuzzWordsfile2.txt")
	checkError(err)
	defer f.Close()

	str:="New File"
	n2,err:=f.Write([]byte(str))
	fmt.Printf("Wrote %d bytes",n2)
	n4,err:=f.WriteString("\nCheck append")
	checkError(err)
	fmt.Printf("Appended %d bytes",n4)
	f.Sync()
	f1.Sync()

	//n3,err:= f.WriteString("add")
}