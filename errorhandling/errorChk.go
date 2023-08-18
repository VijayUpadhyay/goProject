package main

import "fmt"

var err = "err msg"

func main() {
	if errr := returnError(); errr != nil {
		fmt.Printf("failed with error: %+v", errr)
		return
	}

}

func returnError() error {
	return fmt.Errorf("Error is: %s", err)
}
