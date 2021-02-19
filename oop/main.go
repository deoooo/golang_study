package main

import "fmt"

type Animal interface {
	Say()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Say() {
	fmt.Printf("person saying!!! \n")
}

func (p Person) getNameAndAge() (string, int) {
	return p.Name, p.Age
}

type Student struct {
	Person
	Speciality string
}

func (s Student) getSpeciality() string {
	return s.Speciality
}

func (s Student) Say() {
	fmt.Printf("student saying!!! \n")
}

func main() {
	person := Person{
		Name: "Deo",
		Age:  20,
	}

	name, age := person.getNameAndAge()
	fmt.Printf("person name:%s age:%d \n", name, age)

	var a Animal = person
	a.Say()

	student := Student{
		Person: Person{
			Name: "Deo Student",
			Age:  18,
		},
		Speciality: "English",
	}

	student.Say()

	sName, sAge := student.getNameAndAge()
	fmt.Printf("student name:%s age:%d speciality:%s \n", sName, sAge, student.getSpeciality())
}
