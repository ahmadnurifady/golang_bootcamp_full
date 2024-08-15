package application

import (
	"excercise1/employee"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func App() {

	var BucketEmployee = []employee.Employee{}

	var programContinue = true

	for programContinue {

		fmt.Printf("Silahkan pilih menu yang diinginkan :\n1. Menambahkan karyawan\n2. Melihat semua data karyawan\n3. Update kehadiran karyawan\n4. Hapus data karyawan\n5. Nggak ngapa-ngapain\n6. Keluar aplikasi\n")

		var selectedMenu int
		fmt.Scan(&selectedMenu)
		switch selectedMenu {
		case 1:

			var namaKaryawan string
			fmt.Println("Masukan nama karyawan :")
			fmt.Scan(&namaKaryawan)

			randomNumber := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

			inputUser := employee.Employee{
				Id:        strconv.Itoa(randomNumber.Int()),
				Nama:      namaKaryawan,
				Kehadiran: false,
			}

			employee.AddEmployee(&BucketEmployee, inputUser)
			fmt.Println("Data seluruh karyawan :")
			employee.GetAllEmployee(BucketEmployee)
			fmt.Printf("\n\n\n")
		case 2:
			fmt.Println("Data seluruh karyawan :")
			employee.GetAllEmployee(BucketEmployee)
			fmt.Printf("\n\n")
		case 3:
			var idKaryawan string
			fmt.Println("Masukan id karyawan :")
			fmt.Scan(&idKaryawan)
			employee.UpdateKehadiran(BucketEmployee, idKaryawan)
			fmt.Println("Data seluruh karyawan :")
			employee.GetAllEmployee(BucketEmployee)
			fmt.Printf("\n\n")
		case 4:
			var idKaryawan string
			fmt.Println("Masukan id karyawan :")
			fmt.Scan(&idKaryawan)
			employee.DeleteEmployee(&BucketEmployee, idKaryawan)
			fmt.Println("Data seluruh karyawan :")
			employee.GetAllEmployee(BucketEmployee)
			fmt.Printf("\n\n")

		case 5:
			ea := 100000
			for ea >= 0 {
				fmt.Println("EEEEEEEEEEEAAAAAAAAAAAAAAAAA")
				ea--
			}
		case 6:
			programContinue = false
		}
	}

}
