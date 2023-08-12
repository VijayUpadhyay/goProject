package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	var reSlash = regexp.MustCompile(`surveys\/sections\/([a-zA-Z_]+)\/fields`)
	match = reSlash.MatchString("surveys/sections/cogs/fields")
	fmt.Println(match)
	//match, _ = regexp.MatchString("/surveys/sections/[a-zA-Z]/fields", "/surveys/sections/balancesheet/fields")

	reSlash = regexp.MustCompile(`users\/([a-zA-Z0-9-]+)\/roles`)
	match = reSlash.MatchString("users/fbd3036f-0f1c-4e98-b71c-d4cd61213f90/roles")
	fmt.Println(match)
}
