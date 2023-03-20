// functions

package main

import "fmt"

func main() {
	simple() // callig function that means main is calling '()'
	add(2, 4)

	a, b := returningSingleValue()
	fmt.Println(a, b)

	a, b = returnMultipleValues()
	fmt.Println(a, b)

	e, f := mul(10, " srasm")
	fmt.Println(e, f)

	c := sub(10, 30)
	fmt.Println(c)

}

func sub(f, g int) int {
	c := f - g

	return c

}

/*func main(){
n, g := sub(10, 20)
fmt.Println("check this in main how it is gng") // i********imp***
fmt.Println(n, g)
}

func sub(c, d int) (int, int) {
fmt.Println("check this in sub func how it is working", c-d)
return c,d

} */

func mul(a int, b string) (int, string) {
	fmt.Println(" enter values to perform mul", a, b)
	c := 5
	d := " string"

	return c, d

}

func simple() {
	fmt.Println("simple function is called in main")
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
}

func add(a int, b int) { //when you use multiple param u use fmt.println
	fmt.Println("so adding this value", a+b)

}

func returningSingleValue() (int, string) {
	return 20, "siri"
}

func returnMultipleValues() (s int, n string) {
	s = 20
	n = "raw"
	return
}
