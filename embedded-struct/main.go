package main

import "fmt"

type Attr struct {
	Name string
	Age  int
}

type AttrEx struct {
	Name string
}

type Teacher struct {
	Attr
	AttrEx
	Subject string
}

type Student struct {
	Attr
	Score int
}

func (a Attr) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", a.Name, a.Age)
}

func (a AttrEx) String() string {
	return fmt.Sprintf("Name: %s", a.Name)
}

type Walker struct {
	Name string
}

func (w *Walker) Walk() {
	fmt.Printf("%s is walking\n", w.Name)
}

type Runner struct {
	Walker
}

func NewRunner(name string) *Runner {
	return &Runner{Walker: Walker{Name: name}}
}

func (r *Runner) Run() {
	fmt.Printf("%s is running\n", r.Name)
}

func main() {
	teacher := Teacher{Attr: Attr{Name: "Mr. Wang", Age: 30}, Subject: "Math", AttrEx: AttrEx{Name: "Mr. Wang"}}

	student := Student{Attr: Attr{Name: "Tom", Age: 18}, Score: 90}

	// 重複しているフィールド名がある場合は、フィールド名を指定する必要がある。

	//? fmt.Println(teacher.Name, teacher.Age, teacher.Subject, teacher.Attr)
	fmt.Println(teacher.Attr.Name, teacher.Age, teacher.Subject, teacher.Attr.String(), teacher.AttrEx.String())

	fmt.Println(student.Name, student.Age, student.Score)

	runner := NewRunner("Tom")
	// 普通だったら、runner.Walker.Walk()と書く必要があるが、埋め込みを使うと、以下のように書ける。
	runner.Walk()
	runner.Run()
}
