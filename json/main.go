package main

import (
	"encoding/json"
	"fmt"
	"json-golang/domain"
	testhttp "json-golang/test_http"
)

func jsonMarshall(inputUser domain.Address) []byte {

	data, _ := json.Marshal(inputUser)

	fmt.Println("RESULT MARSHALL ==", string(data))

	return data

}

func jsonUnmarshall(inputJson []byte) {
	var user domain.User

	fmt.Println(user)

	err := json.Unmarshal(inputJson, &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("RESULT UNMARSHALL", user)

}

func main() {
	// orang1 := domain.User{
	// 	Name:  "anjay",
	// 	Email: "anjaymail.com",
	// 	Age:   0,
	// }
	// address1 := domain.Address{
	// 	Jalan:  "anjay jalannya tolol",
	// 	Alamat: "anjay alamatnya tolol",
	// }

	// result := jsonMarshall(address1)

	// jsonUnmarshall(result)

	// http.HandleFunc("/user", httpweb.Handler)
	// http.ListenAndServe(":8080", nil)
	testhttp.SimpleHttp()
}
