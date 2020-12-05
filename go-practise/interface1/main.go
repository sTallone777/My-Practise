package main

import "fmt"

type People interface {
	ReturnName() string
}

type Student struct {
	Name string
}

func (s Student) ReturnName() string {
	return s.Name
}

func CheckPeople(test interface{}) {
	if x, ok := test.(People); ok {
		fmt.Printf("Student implements People")
		fmt.Printf(x.ReturnName())
	}
}

func main() {
	cbs := Student{Name: "kakakka"}

	var a People
	a = cbs

	name := a.ReturnName()
	fmt.Println(name)
	CheckPeople(cbs)
	fmt.Println(cbs.ReturnName())

}
