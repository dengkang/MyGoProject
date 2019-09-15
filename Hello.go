package main

import (
	"fmt"
	"time"
	)

func test (x int,y int) int{
	return x+y
}

func test_xc (x int){
	fmt.Println(x)
}

type People struct {
	name string
	age int
}

func (people *People) setXY(name string,age int)  {
	people.age = age
	people.name = name
}

func main() {


	name := "zhangsan"
	x1 := 1
	y1 :=  2
	fmt.Println("hello world "+name)

	z := test(x1,y1)
	fmt.Print(z)
	p := People{"zsf",100}
	fmt.Println(p.age)
	for i:=1;i<10;i++{
		go test_xc(i)
	}
	time.Sleep(2)
}
