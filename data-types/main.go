package main

import "fmt"

// Struct: your own type with fields
type User struct {
	Name string
	Age  int
}

// Method on User, so User matches the Speaker interface
func (u User) Speak() string {
	return "Hi, I am " + u.Name
}

// Interface: any type with Speak() string
type Speaker interface {
	Speak() string
}

func main() {
	fmt.Println("===== 1. Basic types =====")

	var ok bool = true
	var city string = "Bradford"

	var i int = 42
	var i8 int8 = -128
	var i16 int16 = 32000
	var i32 int32 = -2000000
	var i64 int64 = 9000000000

	var u uint = 42
	var u8 uint8 = 255
	var u16 uint16 = 65000
	var u32 uint32 = 4000000000
	var u64 uint64 = 18000000000000000000

	var f32 float32 = 3.14
	var f64 float64 = 10.5

	var c64 complex64 = 1 + 2i
	var c128 complex128 = 3 + 4i

	var b byte = 'A' // same as uint8
	var r rune = 'ا' // same as int32, one Unicode character

	fmt.Printf("%-10T %v\n", ok, ok)
	fmt.Printf("%-10T %v\n", city, city)
	fmt.Printf("%-10T %v\n", i, i)
	fmt.Printf("%-10T %v\n", i8, i8)
	fmt.Printf("%-10T %v\n", i16, i16)
	fmt.Printf("%-10T %v\n", i32, i32)
	fmt.Printf("%-10T %v\n", i64, i64)
	fmt.Printf("%-10T %v\n", u, u)
	fmt.Printf("%-10T %v\n", u8, u8)
	fmt.Printf("%-10T %v\n", u16, u16)
	fmt.Printf("%-10T %v\n", u32, u32)
	fmt.Printf("%-10T %v\n", u64, u64)
	fmt.Printf("%-10T %v\n", f32, f32)
	fmt.Printf("%-10T %v\n", f64, f64)
	fmt.Printf("%-10T %v\n", c64, c64)
	fmt.Printf("%-10T %v\n", c128, c128)
	fmt.Printf("%-10T %v (as char: %c)\n", b, b, b)
	fmt.Printf("%-10T %v (as char: %c)\n", r, r, r)

	fmt.Println()
	fmt.Println("===== 2. Aggregate types =====")

	nums := [3]int{1, 2, 3}
	user := User{Name: "Ali", Age: 30}

	fmt.Printf("%-12T %v\n", nums, nums)
	fmt.Printf("%-12T %+v\n", user, user) // %+v shows field names

	fmt.Println()
	fmt.Println("===== 3. Reference types =====")

	cities := []string{"Leeds", "Paris"}
	cities = append(cities, "Tokyo")

	ages := map[string]int{"Ali": 30, "Sara": 25}

	p := &i  // pointer to i
	*p = 100 // change i through the pointer

	ch := make(chan string, 1)
	ch <- "hello"

	add := func(a, b int) int { return a + b }

	fmt.Printf("%-22T %v\n", cities, cities)
	fmt.Printf("%-22T %v\n", ages, ages)
	fmt.Printf("%-22T %v (value: %v, i is now %v)\n", p, p, *p, i)
	fmt.Printf("%-22T received: %v\n", ch, <-ch)
	fmt.Printf("%-22T add(2, 3) = %v\n", add, add(2, 3))

	fmt.Println()
	fmt.Println("===== 4. Interface =====")

	var s Speaker = user // User has Speak(), so it is a Speaker
	fmt.Printf("%-12T %v\n", s, s.Speak())

	var anything any
	anything = 42
	fmt.Printf("any holds: %-8T %v\n", anything, anything)
	anything = "text"
	fmt.Printf("any holds: %-8T %v\n", anything, anything)
	anything = user
	fmt.Printf("any holds: %-8T %v\n", anything, anything)

	fmt.Println()
	fmt.Println("===== 5. Zero values =====")

	var zInt int
	var zFloat float64
	var zString string
	var zBool bool
	var zPointer *int
	var zSlice []int
	var zMap map[string]int
	var zChan chan int
	var zFunc func()
	var zAny any
	var zUser User

	fmt.Printf("int:     %v\n", zInt)
	fmt.Printf("float64: %v\n", zFloat)
	fmt.Printf("string:  %q\n", zString) // %q shows "" so you can see it
	fmt.Printf("bool:    %v\n", zBool)
	fmt.Printf("pointer: %v\n", zPointer)
	fmt.Printf("slice:   %v (nil? %v)\n", zSlice, zSlice == nil)
	fmt.Printf("map:     %v (nil? %v)\n", zMap, zMap == nil)
	fmt.Printf("chan:    %v\n", zChan)
	fmt.Printf("func:    nil? %v\n", zFunc == nil)
	fmt.Printf("any:     %v\n", zAny)
	fmt.Printf("struct:  %+v\n", zUser)
}
