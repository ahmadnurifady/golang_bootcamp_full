package employee

import "fmt"

// type Repo struct {
// 	DataEmployee *[]Employee
// }

func AddEmployee(dataEmployee *[]Employee, employee Employee) (string, []Employee) {
	var inputEmployee = Employee{
		Id:        employee.Id,
		Nama:      employee.Nama,
		Kehadiran: false,
	}

	*dataEmployee = append(*dataEmployee, inputEmployee)
	fmt.Println(len(*dataEmployee))
	return "SUCCESS ADD", *dataEmployee
}

func GetAllEmployee(dataEmployee []Employee) {
	fmt.Println(dataEmployee)
}

func UpdateKehadiran(dataEmployee []Employee, idEmployee string) {

	for i, employee := range dataEmployee {
		if employee.Id == idEmployee {
			dataEmployee[i].Kehadiran = true
		} else {
			fmt.Println("Id karyawan tidak ditemukan")
		}
	}

}

func DeleteEmployee(dataEmployee *[]Employee, idEmployee string) {
	var elementToDelete int

	for i, employee := range *dataEmployee {
		if employee.Id == idEmployee {
			elementToDelete = i

			fmt.Println(elementToDelete)
			*dataEmployee = append((*dataEmployee)[:elementToDelete], (*dataEmployee)[elementToDelete+1:]...)
		} else {
			fmt.Println("Id karyawan tidak ditemukan")
		}
	}

}
