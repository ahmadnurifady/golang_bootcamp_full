package application

import (
	"fmt"
)

type calcula struct {
	Val1 float64
	Val2 float64
}

func (c calcula) plus() float64 {
	return c.Val1 + c.Val2
}

func (c calcula) minus() float64 {
	return c.Val1 - c.Val2
}

func (c calcula) decade() float64 {
	return c.Val1 * c.Val2
}

func (c calcula) bagi() float64 {
	return c.Val1 / c.Val2
}

func Application() {

	var bucketResult float64

	programCotinue := true

	for programCotinue != false {
		var inputUser1 int
		fmt.Println("masukkan angka pertama")
		fmt.Scan(&inputUser1)

		var operasi string
		fmt.Println("masukkan perintah operasi")
		fmt.Scan(&operasi)

		var inputUser2 int
		fmt.Println("masukkan angka kedua")
		fmt.Scan(&inputUser2)

		input1 := calcula{
			Val1: float64(inputUser1),
			Val2: float64(inputUser2),
		}

		switch operasi {

		case "+":
			bucketResult = input1.plus()
			fmt.Println(bucketResult)
		case "-":
			bucketResult = input1.minus()
			fmt.Println(bucketResult)
		case "*":
			bucketResult = input1.decade()
			fmt.Println(bucketResult)
		case "/":
			bucketResult = input1.bagi()
			fmt.Println(bucketResult)
		case "exit":
			programCotinue = false
		}

	}

	// fmt.Println(input1.decade())

}
