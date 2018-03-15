package main

import (
	"fmt"
	"strings"
)

func main()  {

	// (i) Arrays
	// Declaring an array requires that we specify the size, and once the size is specified, it cannot grow.
	var scores [10] int
	scores[0] = 399
	fmt.Println(scores[0])

	// We can initialize the array with values
	newscores := [4] int{234, 456, 567, 567}

	var arrayLength = len(newscores)
	fmt.Println(arrayLength)

	for index, value := range newscores {
		fmt.Println(index, value)
	}

	// (ii) Slices
	scoresnew := make([]int, 0, 10)
	scoresnew = scoresnew[0:8]
	scoresnew[7] = 9033
	fmt.Println(scoresnew)

	scoresthe := []int{1,2,3,4,5}
	slice := scoresthe[2:4]
	slice[0] = 999
	fmt.Println(scoresthe)

	// To find the first space in a string (yes, slices work on strings too!) after the first five characters
	haystack := "the spice must flow";
	word := strings.Index(haystack[5:], " ")
	fmt.Println(word)


}