package main

import "fmt"

/*
	Interface defines the behavior of the object.
	It only specifies what the object is supposed to do.

	In Go an interface is a set of method signatures.
	Interface specifies what methods a type should have and the type decides how to implement these methods.

	(Eg.) WashingMachine can be a interface with method signatures cleaning() and drying().
	Any type which provides definition for cleaning() and drying() is said to implement WashingMachine
	interface.
*/

// interface definition
type VowelsFinder interface{
	FindVowels() []rune
}

type MyString string

func (s MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range s{
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main(){

	name := MyString("Krishna")
	var v VowelsFinder
	v = name
	fmt.Printf("Vowels are %c\n", v.FindVowels())

}