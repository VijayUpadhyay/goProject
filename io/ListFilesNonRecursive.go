package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
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
