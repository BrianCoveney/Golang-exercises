package main

import (
	"fmt"
)

type Saiyan struct {
	Name string
	Power int
}


// (i)
// We use *X as a pointer to value of type X
func Super(sWithoutPointer Saiyan, sWithPointer *Saiyan)  {
	sWithoutPointer.Power += 10000
	sWithPointer.Power += 10000
}

// (ii)
// We can associate a method with a structure
func (s *Saiyan) SuperS() {
	s.Name = "Goku3"
	s.Power += 10000
}

// (iii) 'Constructors'
// Structures don’t have constructors. Instead, you create a function that returns an instance of the desired type
// (like a factory)
func NewSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name: name,
		Power: power,
	}
}

//(v) Composition
// Go supports composition, which is the act of including one structure into another.

type Person struct {
	Name string
}

func (p *Person) Introduce()  {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type SaiyanNew struct {
	*Person
	Power int
}


// (vi)
// Overloading
func (s *Saiyan) Introduce() {
	fmt.Printf("Hi, I'm %s. Ya!\n", s.Name)
}



func main() {
	// (i)
	goku1 := Saiyan{"Goku1", 9000}

	// We use the & operator to get the address of our value
	goku2 := &Saiyan{"Goku2", 9000}

	Super(goku1, goku2)

	fmt.Println("Goku1", goku1.Power) // 9000
	fmt.Println("Goku2", goku2.Power) // 19000

	// (ii)
	goku3 := &Saiyan{"Goku3", 9001}
	goku3.SuperS()
	fmt.Println(goku3.Name, goku3.Power) // 19001

	// (iii) 'Constructors' + (iv) 'New'
	// Despite the lack of constructors, Go does have a built-in new function which is used to allocate the memory required
	// by a type. The result of new(X) is the same as &X{}:
	goku4 := new(Saiyan)
	goku4.Name = "Goku4"
	goku4.Power = 9001
	fmt.Println(goku4.Name, goku4.Power)

	// vs

	goku5 := &Saiyan{
		Name: "Goku5",
		Power: 9001,
	}
	fmt.Println(goku5.Name, goku5.Power)

	goku6 := NewSaiyan("Goku6", 9001)
	fmt.Println(goku6.Name, goku6.Power)

	//(v)
	goku7 := &SaiyanNew {
		Person: &Person{"Goku7"},
		Power: 9001,
	}

	goku7.Introduce()

	// (vi)
	goku8 := &Saiyan{
		Name:"Goku8",
	}
	goku8.Introduce()

	goku7.Person.Introduce()


}
