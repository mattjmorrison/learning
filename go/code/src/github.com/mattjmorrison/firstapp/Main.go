package main

import (
	"fmt"
)

func main() {
	booleans()
	numbers()
	math()
	bitoperators()
	floatingpoint()
	complextypenumbers()
	texttype()

	fmt.Println("Constants")
	constantnaming()
	typedconstants()
	untypedconstants()
	enumeratedconstants()
	enumerationexpressions()
}

func constantnaming() {
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)
}

func typedconstants() {
	fmt.Println("Typed Constants")
	const a int = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)

	const x int = 42
	var y int = 27
	fmt.Printf("%v, %T\n", x + y, x + y)
}

func untypedconstants() {
	fmt.Println("Untyped Constants")
	const w = 42
	var z int16 = 27
	fmt.Printf("%v, %T\n", w + z, w + z)
}

func enumeratedconstants() {
	fmt.Println("Enumerated Constants")
	const a = iota
	fmt.Printf("%v, %T\n", a, a)

	const (
		w = iota
		x
		y
	)
	fmt.Printf("%v\n", w)
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)

	const (
		w2 = iota
	)
	fmt.Printf("%v\n", w2)

	// NOTE - is this sort of like an enum?
	const (
		_ = iota + 5 // ignore the initial value
		catSpecialist
		dogSpecialist
		snakeSpecialist
	)

	var specialistType int
	fmt.Printf("%v\n", catSpecialist)
	fmt.Printf("%v\n", specialistType == catSpecialist)

	const (
		_ = iota
		KB = 1 << (10 * iota)
		MB // the iota pattern continues down shifting bits for each one below
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	const (
		isAdmin = 1 << iota
		isHeadquarters
		canSeeFinancials

		canSeeAfrica
		canSeeAsia
		canSeeEurope
		canSeeNorthAmerica
		canSeeSouthAmerica
	)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin & roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters & roles == isHeadquarters)
}

func enumerationexpressions() {

}

func booleans() {
	fmt.Println("Booleans")
	var y bool = true
	fmt.Printf("%v, %T\n", y, y)
	var n bool = false
	fmt.Printf("%v, %T\n", n, n)

	a := 1 == 1
	b := 2 == 2
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)

	var x bool
	fmt.Printf("%v, %T\n", x, x)
}

func numbers() {
	var a int = 10  // can have system dependent minimum/maximum sizes?
	var b int8 = 10
	var c int16 = 10
	var d int32 = 10
	var e int64 = 10
	var f uint = 10
	var g uint8 = 10
	var h uint16 = 10
	var i uint32 = 10
	var j uint64 = 10  // in the youtube tutorial he said this didn't exist
	fmt.Println("Numbers")
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", d, d)
	fmt.Printf("%v %T\n", e, e)
	fmt.Printf("%v %T\n", f, f)
	fmt.Printf("%v %T\n", g, g)
	fmt.Printf("%v %T\n", h, h)
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", j, j)
}

func math() {
	a := 10
	b := 3
	fmt.Println("Math")
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	var x int = 10
	var y int8 = 3
	fmt.Println(x + int(y))
}

func bitoperators() {
	a := 10  // 1010
	b := 3   // 0011
	fmt.Println("Bitwise Operators")
	fmt.Println(a & b)  // 0010 == 2
	fmt.Println(a | b)  // 1011 == 11
	fmt.Println(a ^ b)  // 1001 == 9
	fmt.Println(a &^ b) // 0100 == 8??? no, 1000 would be 8, 
	                    // ok, so it's "and not" so for each bit it's "a and not b"
						// the youtube video had this one wrong

	fmt.Println("Bit Shifting")
	c := 8 // 2^3
	fmt.Println(c << 3) // 2^3 * 2^3 == 2^6
	fmt.Println(c >> 3) // 2^3 / 2^3 == 2^0
}

func floatingpoint() {
	fmt.Println("Floating Point Numbers")
	var n float64 = 3.14
	// also float32
	fmt.Printf("%v, %T\n", n, n)
	n = 13.7e72
	fmt.Printf("%v, %T\n", n, n)
	n = 2.1E14
	fmt.Printf("%v, %T\n", n, n)

	fmt.Println("Floating Point Math")
	a := 10.2
	b := 3.7
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
}

func complextypenumbers() {
	fmt.Println("Complex Numbers")
	var n complex64 = 1 + 2i
	// also complex128
	fmt.Printf("%v, %T\n", n, n)

	var x complex64 = 2i
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", real(x), real(x))
	fmt.Printf("%v, %T\n", imag(x), imag(x))
	var w complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", w, w)

	a := 1 + 2i
	b := 2 + 5.2i
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
}

func texttype() {
	fmt.Println("Text Types")
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)

	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", string(s[2]), string(s[2]))

	s2 := "this is also a string"
	fmt.Printf("%v, %T\n", s + s2, s + s2)

	b := []byte(s) // byte slice
	fmt.Printf("%v, %T\n", b, b)

	// rune - utf-32 characters instead of utf-8
	var r rune = 'a' // single quotes - rune is same thing as int32
	fmt.Printf("%v, %T\n", r, r)
}