package main

import (
	"fmt"
	str "strings"
)

func main() {
	p := fmt.Println
	p("Contains:			", str.Contains("test", "st"))
	p("Count:			", str.Count("test", "t"))
	p("HasPrefix:		", str.HasPrefix("test", "te"))
	p("HasSuffix:		", str.HasSuffix("test", "st"))
	p("Index:			", str.Index("test", "e"))
	p("Join:				", str.Join([]string{"a", "b", "c"}, "-"))
	p("Repeat:			", str.Repeat("c", 6))
	p("Replace:			", str.Replace("f00", "0", "o", -1))
	p("Split:			", str.Split("a-b-c-d-e-f", "-"))
	p("ToLower:			", str.ToLower("Test"))
	p("ToUpper:			", str.ToUpper("Test"))
	p()
	p("Len:				", len("hell0"))
	p("Char:				", "hello"[0])
}
