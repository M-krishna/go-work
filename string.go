package main

import "fmt"

/*
	strings in Go are slice of bytes.
	strings can be created by enclosing contents inside " "
	strings in Go are in unicode compliant and are UTF-8 encoded. 
	strings are immutable in Go.
*/

// Printing bytes in a string
func printBytes(s string){
	fmt.Println("Length of a string", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) // %x is used to represent hexadecimal format
	}
	fmt.Println()
}

// printing characters in a string
func printCharacters(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i]) // %c is used to represent character format
	}
	fmt.Println()
}

// printing characters using rune
func printCharactersUsingRune(s string){
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c starts at byte %d\n ", runes[i], runes[i])
	}
	fmt.Println()
}

// for range loop on a string
func forRangeOnString(s string){
	for i, rune := range s {
		fmt.Printf("%c starts at byte %d\n ", rune, i)
	}
	fmt.Println()
	fmt.Println()
}

// trying to mutate a string
/*func mutate(s string){
	s[0] = "Y"
}*/

// trying to mutate a string by converting it into slice of runes
func mutateByRunes(s []rune) string {
	s[0] = 'Y' // notice here we are using ' '. because runes are represented as ''.
	return string(s)
}

func main() {
	hello := "Hello World"
	fmt.Println(hello)

	/* 
		Accessing each bytes in a string. 
		since string is a slice of bytes.
	*/
	fmt.Println("Printing each bytes in a string")
	printBytes(hello)

	fmt.Println("Printing Characters in a string")
	printCharacters(hello)

	/*
		but at some point a character may occupy 2 bytes due to its character encoding.
		this where rune comes to help us.
	*/

	fmt.Println()
	fmt.Println()
	fmt.Println("Rune")

	/*
		rune is a built-in datatype in Go.
		It's the alias of int32. rune represents a unicode point in Go.
		It does not matter how many bytes the code point occupies, it can be represented by a rune.
	*/
	printCharactersUsingRune(hello)

	// using for range loop on a string
	fmt.Println("Using for range loop on a string")
	forRangeOnString(hello)

	/*
		Constructing a string from slice of bytes
	*/
	fmt.Println("Constructing a string from slice of bytes")
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9} // hex bytes of the string Café
	str := string(byteSlice)
	fmt.Println(str)

	// what if we have decimal equivalent of hex values
	fmt.Println("By using decimal equivalent of hex values")
	byteSlice1 := []byte{67, 97, 102, 195, 169}
	str1 := string(byteSlice1)
	fmt.Println(str1)

	/*
		Constructing a string from slice of runes
	*/
	fmt.Println("Constructing a string from slice of runes")
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str2 := string(runeSlice)
	fmt.Println(str2)

	// we cannot change a string once its created.
	newstr := "Krishna"
	// mutate(newstr) // cannot assign to s[0].

	/*
		To workaround string immutability, we can turn the string into slice of runes.
		then that slice is mutated for whatever changes and converted back to a new string.
	*/
	fmt.Println("Mutating a string by turning it into slice of runes")
	fmt.Println(mutateByRunes([]rune(newstr)))
	

}