<<<<<<< HEAD
package main

import "fmt"

type Ij interface {
	M()
}

func main() {
	var i Ij
	describe2(i)
	i.M()
}

func describe2(i Ij) {
	fmt.Printf("(%v, %T)\n", i, i)
}
=======
package main

import "fmt"

type Ij interface {
	M()
}

func main() {
	var i Ij
	describe2(i)
	i.M()
}



func describe2(i Ij) {
	fmt.Printf("(%v, %T)\n", i, i)
}
>>>>>>> origin/master
