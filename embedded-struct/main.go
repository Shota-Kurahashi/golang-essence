package main

import "fmt"

type Attr struct {
	Name string
	Age  int
}

type Teacher struct {
	Attr
	Subject string
}

type Student struct {
	Attr
	Score int
}

func main() {
	teacher := Teacher{Attr: Attr{Name: "Mr. Wang", Age: 30}, Subject: "Math"}

	student := Student{Attr: Attr{Name: "Tom", Age: 18}, Score: 90}

	fmt.Println(teacher.Name, teacher.Age, teacher.Subject, teacher.Attr)
	fmt.Println(student.Name, student.Age, student.Score)
}
