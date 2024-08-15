package main

import "consume-api-soap-rest/internal/provider/server"

func main() {

	// consumeapi.ConsumeApiCalculatorSoap()
	// consumeapi.ConsumeEvent()

	server.NewServer().Run()

}
