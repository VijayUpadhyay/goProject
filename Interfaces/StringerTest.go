<<<<<<< HEAD
package main

import "fmt"

type Person struct {
	firstName string
	age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%v is %d years", p.firstName, p.age)
}

func main() {
	v := Person{"vijay", 10}
	fmt.Println("Test")
	fmt.Println(v)

}
=======
package main

import "fmt"

type Person struct {
	firstName string
	age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%v is %d years", p.firstName, p.age)
}

func main() {
	v := Person{"vijay", 10}
	fmt.Println("Test")
	fmt.Println(v)

}
>>>>>>> origin/master
