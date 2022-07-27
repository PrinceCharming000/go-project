package main

import (
	"fmt"
	"math"
	"time"
)

type Animal struct {
	Name string
	Age  uint
}

func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func main() {
	// fmt.Println("Hello World!")

	integer := 23
	fmt.Printf("%v\n", integer)

	placeholders := `foo "bar"`
	fmt.Printf("%v %s %q %#q\n", placeholders, placeholders, placeholders, placeholders)

	isLegume := map[string]bool{
		"peanut":    true,
		"dachshund": false,
	}
	fmt.Printf("%v %#v\n", isLegume, isLegume)

	person := struct {
		Name string
		Age  int
	}{"Kim", 22}
	fmt.Printf("%v %+v %#v\n", person, person, person)

	greats := [5]string{"Apple", "Banana", "Orange", "Strawberry", "Tomato"}
	fmt.Printf("%v %q\n", greats, greats)

	kGreats := greats[:3]
	fmt.Printf("%v %q %#v\n", kGreats, kGreats, kGreats)

	now := time.Unix(123456789, 0).UTC()
	fmt.Printf("%v %q\n", now, now)

	a, b := 3.0, 4.0
	h := math.Hypot(a, b)
	fmt.Print("The vector (", a, b, ") has length ", h, ".\n")

	fmt.Printf("The vector (%g %g) has length %g.\n", a, b, h)

	animal := Animal{
		Name: "Gorilla",
		Age:  2,
	}
	fmt.Println(animal)

	var name string
	var age int
	n, err := fmt.Sscanf("Kim is 22 years old", "%s is %d years old", &name, &age)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d: %s, %d\n", n, name, age)
}
