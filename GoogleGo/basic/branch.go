package basic

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	var g string
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "output.txt"
	if context, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", context)

	}
	fmt.Println(
		grade(191),
		grade(59),
		grade(68),
		grade(88),
		grade(100),
	)
}
