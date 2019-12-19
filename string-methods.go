package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main(){
	p("Contains: ", strings.Contains("Krishna", "na"))
	p("Count: ", strings.Count("Malayalam", "a"))
	p("HasPrefix: ", strings.HasPrefix("Krishna", "Kr"))
	p("HasSuffix: ", strings.HasSuffix("Krishna", "na"))
	p("Index: ", strings.Index("Krishna", "s"))
	p("Join: ", strings.Join([]string{"k", "r", "i", "s", "h"}, "-"))
	p("Repeat: ", strings.Repeat("R", 5))
	p("Replace: ", strings.Replace("Krishna", "K", "R", 1))
	p("Split: ", strings.Split("k-r-i-s-h-n-a", "-"))
	p("ToLower: ", strings.ToLower("KRISHNA"))
	p("ToUpper: ", strings.ToUpper("krishna"))

	p("Len: ", len("Krishna"))
	p("Char: ", "krishna"[1])
}