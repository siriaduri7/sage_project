package main

import "fmt"

func main() {
	// function is called first class parameter as function is also called as variable

	//writing simple function below declaring in var abc func = func () {} /// a := func () {} ()
	/////////////////////////////(**********************************)////////////////////////////////////////////

	//other method of function declaring but its same as calling function type other way of calling but functionality is same)
	//closure function and verbose and pritf and sprintf is used in closure func//
	printsomething := func() {
		fmt.Println("something!")
	}

	printsomething()

	hi := func(a int, b int) int {
		fmt.Println("find execution flow")
		return a + b
	}

	//hi(10, 20)
	fmt.Printf("%v+ %v = %v ", 10, 20, hi(10, 20))

	/////////////////////////////////////////////////////////

	fmt.Println("//////////////////////////")
	c, f := func(a string, b int) (string, int) {
		return fmt.Sprintf("String : %s", a), b // a is like hello replaced in %s placeholder
	}("hello", 5)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", f)

}

/*	k := add(30, 5)
	fmt.Println(k)
}

/////////////////////////////(**********************************)////////////////////////////////////////////
func add(a, b int) (g int) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	g = a + b
	fmt.Println("after adding values its stored in g ")
	return g

}
*/

//call by value
