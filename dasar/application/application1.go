package application

import "fmt"

func Start() {
	var inputStart int
	for {
		fmt.Println("Selamat datang di aplikasi")
		fmt.Printf("Silahkan pilih menu yang diinginkan\n ")
		scanStart, _ := fmt.Scan(&inputStart)
		if scanStart == 0 {
			fmt.Println("Masukan pilihan menu yang benar")
		}
		switch inputStart {
		case 1:
			fmt.Println("MENU 1")
		case 2:
			fmt.Println("MENU 2")
		case 3:
			fmt.Println("MENU 3")
			break
		}
	}
}
