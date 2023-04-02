package syntax

import "fmt"

type Person interface {
	fullName() string
}

type Employees struct {
	firstName string
	surname   string
}

func (e Employees) fullName() string {
	return fmt.Sprintf("%s %s", e.firstName, e.surname)
}

func personTask(p Person) {
	fmt.Println(p.fullName())
}

func Interface() {
	fmt.Println("Learn Go Interface")

	e := Employees{firstName: "John", surname: "Doe"}
	personTask(e)

	/* Interface */
	/*
		//--------------------------------------------------
		fmt.Println(person{"Bob", 20})
		fmt.Println(person{name: "Alice", age: 30})
		fmt.Println(person{name: "Fred"})
		fmt.Println(&person{name: "Ann", age: 40})
		fmt.Println(newPerson("Jon"))

		s := person{name: "Sean", age: 50}
		fmt.Println(s.name)

		sp := &s
		fmt.Println(sp.age)

		sp.age = 51
		fmt.Println(sp.age)
		//--------------------------------------------------
	*/
}
