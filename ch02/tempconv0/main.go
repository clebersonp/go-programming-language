package main

import (
	formatter "fmt" // alternative for import name
)

// Type declarations define a new named type that has the same underlying type as an existing type
// syntax: type name underlying-type
// type MyType int

// Celsius and Fahrenheit are not the same type even though both have the same underlying type, float64
// name type need conversion to other type of the same underlying type: like Celsius(t) or Fahrenheit(t) is required
// to convert from a float64
// A conversion from one type to another is allowed if both have the same underlying type, or if both are unnamed pointer
// types that point to variables of the same underlying type.

type Celsius float64
type Fahrenheit float64

// change the named type behavior with methods and satisfying interfaces type
func (c Celsius) String() string {
	return formatter.Sprintf("%gC", c) // nC -> 0C
}

func (f Fahrenheit) String() string {
	return formatter.Sprintf("%gF", f) // nF -> 0F
}

const (
	AbsoluteZeroC Celsius    = -273.15
	FreezingC     Celsius    = 0
	BoilingC      Celsius    = 100
	FreezingF     Fahrenheit = 32
	BoilingF      Fahrenheit = 212
)

func main() {
	formatter.Printf("%gC\n", BoilingC-FreezingC) // arithmetic operators work the same for Celsius
	// fmt.Printf("%g\n", BoilingC-FreezingF) // arithmetic operators don't work to different types, even though they are the same underlying type. Needs an explicit conversion
	formatter.Printf("%gC\n", BoilingC-Celsius(FreezingF)) // An explicit conversion from one named type to another, both the same underlying type
	boilingF := CToF(BoilingC)
	formatter.Printf("%gF\n", boilingF-CToF(FreezingC)) // both named type are the same type and the same underlying type
	//fmt.Printf("%g\n", boilingF-FreezingC)       // don't work, they are different types, even though they are the same underlying type, needs a conversion to Celsius or Fahrenheit
	formatter.Printf("%gC\n", FToC(FreezingF))
	formatter.Printf("%gC\n", FToC(BoilingF))

	formatter.Println("===================================")
	// Comparison operators like == and < can also be used to compare a value of a named type to another
	// of the same type, or to a value of the underlying type. But two values of different named types cannot be compared directly
	var c Celsius    // Zero value 0
	var f Fahrenheit // Zero value 0
	formatter.Println(c == 0)
	formatter.Println(f >= 0)
	//fmt.Println(c == f) // type mismatch

	// this is comparing the value 0 not the type
	formatter.Println(c == Celsius(f)) // needs an explicit conversion of one two types

	formatter.Println("================================================")
	d := FToC(BoilingF)
	formatter.Println(d.String()) // Celsius satisfy interface type Stringer, so, the String() function will be used
	formatter.Printf("%v\n", d)
	formatter.Printf("%s\n", d)
	formatter.Println(d)
	formatter.Printf("%g\n", d)   // verb g not call String method
	formatter.Println(float64(d)) // does not call String method
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
