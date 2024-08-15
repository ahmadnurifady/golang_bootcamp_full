package pakett

import "fmt"

func addFruitsInMiddlSlice(silice []string, newElemen string) ([]string, string) {
	cutAndReplaceElement := silice[len(silice)-1]

	newElementSlice := append(silice, cutAndReplaceElement)

	newElementSlice[len(newElementSlice)-2] = newElemen

	return newElementSlice, newElemen
}

func Pakett() {

	nangka := "nangka"

	var fruits = []string{"apel", "mangga", "pisang", "semangka", nangka}
	fmt.Println(fruits)

	fmt.Println(addFruitsInMiddlSlice(fruits, "anggur"))

	// var newSlice = fruits[:1]
	// var sliceNangka = nangka[1:]

	// fmt.Println(newSlice)
	// fmt.Println(sliceNangka)

	// appendFruits := append(fruits, "rambutan")

	// fmt.Println(appendFruits)

}
