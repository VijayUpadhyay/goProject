<<<<<<< HEAD
package main

import (
	"fmt"
)

type IPAddr [4]byte

type IPAddrr struct {
	ipAddress byte
}

// TODO: Add a "String() string" method to IPAddr.

func (m IPAddr) String() string {

	n := len(m)
	s := string(m[:n])
	/*for i:=0;i<len(m);i++  {
		fmt.Println("m[i] is",m[i])
		p:=string(m[i])
		fmt.Println("p is ",p)
		s[i] = m[i]

		s = append(s,m[i])

	}*/
	fmt.Sprintf("Value is %v", s)
	//strings.Replace(string(m)," ",".",-1)
	return fmt.Sprintf("IP Address is %v\n", s)
}

func main() {
	var hosts = map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	//fmt.Println(hosts)
	fmt.Println(hosts["loopback"], hosts["googleDNS"])
	/*for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}*/
}
=======
package main

import (
	"fmt"
)

type IPAddr [4]byte

type IPAddrr struct {
	ipAddress byte
}

// TODO: Add a "String() string" method to IPAddr.

func (m IPAddr) String()  string{

	n := len(m)
	s := string(m[:n])
	/*for i:=0;i<len(m);i++  {
		fmt.Println("m[i] is",m[i])
		p:=string(m[i])
		fmt.Println("p is ",p)
		s[i] = m[i]

		s = append(s,m[i])

	}*/
	fmt.Sprintf("Value is %v",s)
	//strings.Replace(string(m)," ",".",-1)
	return fmt.Sprintf("IP Address is %v\n",s)
}

func main() {
	var hosts = map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	//fmt.Println(hosts)
	fmt.Println(hosts["loopback"],hosts["googleDNS"])
	/*for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}*/
}
>>>>>>> origin/master
