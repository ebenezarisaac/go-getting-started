package main

import "fmt"

type Tester struct {
	name string
}

func nomain() {
	fmt.Println("Main starting...")
	var a *Tester
	a = &(Tester{name: "Ebenezar"})
	fmt.Println(a)
	fmt.Println(*a)

}

func forEach() {
	slice := []int{1, 2, 3}
	for i, v := range slice {
		fmt.Println(i, v)
	}

	ports := map[string]int{"http": 80, "https": 443}
	for k, v := range ports {
		fmt.Println(k, v)
	}
	for _, v := range ports {
		fmt.Println(v)
	}
}

func infiniteLoop() {
	var i int
	for {
		if i < 5 {
			break
		}
		fmt.Println(i)
		i++
	}
}

func forUntil() {
	var i int
	for i < 5 {
		fmt.Println(i)
		i++
		if i == 3 {
			continue
		}
	}
}

func pointerMath() {
	a := 1
	// b := 2
	fmt.Println(a)
	a = 3
	b := 4
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(addByPointers(&a, &b))
}
func addByPointers(a *int, b *int) int {
	return *a + *b
}

func add(a int, b int) int {
	return a + b
}
