package main

import "fmt"

func main()  {
	s:=[]int{2,23,0,-24,3,-2}
	c:=make(chan int)
	go sum(s[:len(s)/2],c)
	go sum(s[len(s)/2:],c)
	x,y:= <-c,<-c //first <-c will have sum of the last three index values i.e recent value
	fmt.Println(x,y,x+y)
}

func sum(arr []int,c chan int)  {
	sum:=0
	for i:=0;i<len(arr);i++ {
		sum+=arr[i]
	}
	c<-sum
}