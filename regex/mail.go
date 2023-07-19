package main

import (
	"fmt"
	"regexp"
)

func main() {
	mailId := "vij.upadyay92@gmail.com"
	// `^[a-z0-9._%+\-]+@gmail.com?` --> for 0 or more
	f, err := regexp.MatchString(`^[a-z0-9._%+\-]+@gmail.com$`, mailId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)
}
