package formats

import "fmt"

func Main() {
	fmtPackage()
}

type F struct {
	Name string
	Age  int
}
type Fruit int

const (
	Apple Fruit = iota
	Orange
	Banana
)

func (i Fruit) String() string {
	switch i {
	case Apple:
		return "Apple"
	case Orange:
		return "Orange"
	case Banana:
		return "Banana"

	}

	return fmt.Sprintf("Fruit(%d)", i)

}

func (f *F) String() string {
	return fmt.Sprintf("NAME=%q AGE=%d", f.Name, f.Age)
}

func fmtPackage() {
	x := 1
	s := fmt.Sprint(x)

	fmt.Println(s)

	f := &F{
		Name: "foo",
		Age:  20,
	}

	fmt.Printf("%v\n", f)
	fmt.Printf("%+v\n", f)
	fmt.Printf("%#v\n", f)

	fmt.Printf("%v\n", Apple)

}
