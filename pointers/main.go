package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// Value receiver: gets a COPY, change is lost
func (u User) BirthdayCopy() {
	u.Age++
}

// Pointer receiver: gets the ADDRESS, change stays
func (u *User) Birthday() {
	u.Age++
}

// Gets a copy of the number
func addOneValue(n int) {
	n++
}

// Gets the address of the number
func addOnePointer(n *int) {
	*n++
}

// Returns a pointer to a new User
func newUser(name string) *User {
	return &User{Name: name, Age: 18}
}

func main() {
	fmt.Println("===== 1. Basics: & and * =====")

	x := 10
	p := &x // & = "address of"

	fmt.Println("x       =", x)
	fmt.Println("&x      =", &x) // address
	fmt.Println("p       =", p)  // same address
	fmt.Println("*p      =", *p) // * = "value at address"
	fmt.Printf("type p  = %T\n", p)

	fmt.Println()
	fmt.Println("===== 2. Change value through pointer =====")

	*p = 50
	fmt.Println("after *p = 50, x =", x) // x changed too

	fmt.Println()
	fmt.Println("===== 3. Functions: copy vs pointer =====")

	num := 5
	addOneValue(num)
	fmt.Println("after addOneValue:  ", num) // still 5

	addOnePointer(&num)
	fmt.Println("after addOnePointer:", num) // now 6

	fmt.Println()
	fmt.Println("===== 4. Struct methods: value vs pointer receiver =====")

	user := User{Name: "Ali", Age: 30}

	user.BirthdayCopy()
	fmt.Println("after BirthdayCopy:", user.Age) // still 30

	user.Birthday()                              // Go turns this into (&user).Birthday()
	fmt.Println("after Birthday:    ", user.Age) // now 31

	fmt.Println()
	fmt.Println("===== 5. Pointer to struct =====")

	up := &user
	up.Age = 40 // no need for (*up).Age, Go does it for you
	fmt.Println("user.Age =", user.Age)

	fmt.Println()
	fmt.Println("===== 6. Two pointers, same value =====")

	a := &user
	b := &user
	a.Name = "Sara"
	fmt.Println("b.Name =", b.Name) // Sara, because a and b point to the same user
	fmt.Println("a == b ?", a == b)

	copyUser := user // a real copy, not a pointer
	copyUser.Name = "Copy"
	fmt.Println("user.Name =", user.Name) // still Sara

	fmt.Println()
	fmt.Println("===== 7. new() and returning pointers =====")

	n := new(int) // pointer to a new int, value 0
	fmt.Println("*n =", *n)
	*n = 7
	fmt.Println("*n =", *n)

	u2 := newUser("Bilal")
	fmt.Printf("u2 = %+v (type %T)\n", *u2, u2)

	fmt.Println()
	fmt.Println("===== 8. nil pointer =====")

	var empty *User
	fmt.Println("empty == nil ?", empty == nil)

	if empty != nil {
		fmt.Println(empty.Name)
	} else {
		fmt.Println("empty is nil, so we do not use it")
	}

	// Remove the // below to see the crash:
	// fmt.Println(empty.Name) // panic: nil pointer dereference

	fmt.Println()
	fmt.Println("===== 9. Pointer to pointer =====")

	pp := &p // pointer to a pointer
	fmt.Println("**pp =", **pp)
	**pp = 99
	fmt.Println("x    =", x)
}
