// package main

// import (
// 	"dasar-golang-phincon/goroutiness/soal"
// 	"math/rand"
// 	"time"
// )

// // catchProbability checks if catching is successful given a percentage rate
// func catchProbability(rate int) string {
// 	if rate < 0 || rate > 100 {
// 		return "Invalid rate. Please provide a rate between 0 and 100."
// 	}

// 	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
// 	chance := rand.Intn(100) + 1     // Generate a random number between 1 and 100

// 	if chance <= rate {
// 		return "SUCCESS, you caught it"
// 	}
// 	return "FAIL, it got away"
// }

// // func sum(val1, val2 int) int {
// // 	sumof := val1 + val2
// // 	return sumof
// // }

// type Product struct {
// 	name  string
// 	stock int
// }

// func nambahStock(product *Product) {
// 	product.stock++
// }

// func ngurangStock(product *Product) {
// 	product.stock--
// }

// func addStock(product []Product, productName string) {
// 	for i, value := range product {
// 		if value.name == productName {
// 			product[i].stock++
// 		}
// 	}

// }

// func minusStock(product []Product, productName string) {
// 	for i, value := range product {
// 		if value.name == productName {
// 			product[i].stock--
// 		}
// 	}

// }

// func main() {

// 	soal.StartRaceCondition()

// 	// goroutiness.RaceConditions()

// 	// goroutiness.StartGorou()
// 	// materi.InterMuka()

// 	// scaner := bufio.NewScanner(os.Stdin)
// 	// scaner.Scan()
// 	// data := scaner.Text()
// 	// fmt.Println(data)

// 	// application.Start()
// 	// application.Application()

// 	// pakett.Pakett()

// 	// var satu int = 11
// 	// var dua *int = &satu

// 	// fmt.Println(dua)
// 	// fmt.Println(&dua)
// 	// fmt.Println(*dua)

// 	// *dua = 20

// 	// fmt.Println(satu)
// 	// fmt.Println()

// 	// var product1 = Product{
// 	// 	name:  "sepatu",
// 	// 	stock: 10,
// 	// }

// 	// var product2 = Product{
// 	// 	name:  "celana",
// 	// 	stock: 1,
// 	// }

// 	// var allProducts []Product = []Product{product1, product2}

// 	// fmt.Println(allProducts)
// 	// addStock(allProducts, product1.name)
// 	// fmt.Println(product2)
// 	// fmt.Println("after addStock", allProducts)
// 	// minusStock(allProducts, product1.name)
// 	// fmt.Println("after minusStock", allProducts)

// 	// var pointerProduct = &product1

// 	// fmt.Println(*pointerProduct)

// 	// var value int = 10
// 	// var value1 *int = &value

// 	// fmt.Println(value)
// 	// fmt.Println(*value1)

// 	// result := sum(2, 2)

// 	// fmt.Println(result)

// 	// now := time.Now()

// 	// // Create sample Pokemon
// 	// pokemons := []model.Pokemon{
// 	// 	{ID: 1, Name: "Pikachu", Type: "Electric", CatchRate: 70, IsRare: false, RegisteredDate: now},
// 	// 	{ID: 2, Name: "Charmander", Type: "Fire", CatchRate: 40, IsRare: true, RegisteredDate: now.AddDate(0, -1, 0)},
// 	// 	{ID: 3, Name: "Bulbasaur", Type: "Grass/Poison", CatchRate: 10, IsRare: true, RegisteredDate: now.AddDate(0, -6, 0)},
// 	// 	{ID: 4, Name: "Dragonite", Type: "Dragon/Flying", CatchRate: 30, IsRare: true, RegisteredDate: now.AddDate(0, -8, 0)},
// 	// 	{ID: 5, Name: "Mew", Type: "Psychic", CatchRate: 1, IsRare: true, RegisteredDate: now.AddDate(0, -10, 0)},
// 	// 	{ID: 6, Name: "KAPALLAWD", Type: "Water", CatchRate: 90, IsRare: false, RegisteredDate: now.AddDate(0, -11, 0)},
// 	// }

// 	// var turn int
// 	// var bucketResult = make([]string, 10)

// 	// fmt.Println(turn)

// 	// for {
// 	// 	if turn <= 10 {

// 	// 		turn++

// 	// 		fmt.Println("Available Pokémon:")
// 	// 		for _, pokemon := range pokemons {
// 	// 			fmt.Printf("ID: %d, Name: %s, Type: %s, Catch Rate: %d%%, Is Rare: %t, Registered Date: %s\n",
// 	// 				pokemon.ID, pokemon.Name, pokemon.Type, pokemon.CatchRate, pokemon.IsRare, pokemon.RegisteredDate.Format(time.RFC1123))
// 	// 		}

// 	// 		var id int
// 	// 		fmt.Print("Enter Pokémon ID to attempt to catch: ")
// 	// 		scanId, _ := fmt.Scan(&id)
// 	// 		if scanId == 0 || scanId > len(pokemons) {
// 	// 			fmt.Printf("jumlah gacha sebelum diubah = %d \n", turn)
// 	// 			turn--
// 	// 			fmt.Printf("jumlah gacha setelah diubah = %d \n", turn)

// 	// 		}

// 	// 		var selectedPokemon *model.Pokemon
// 	// 		for _, pokemon := range pokemons {
// 	// 			if pokemon.ID == id {
// 	// 				selectedPokemon = &pokemon
// 	// 				break
// 	// 			}
// 	// 		}

// 	// 		if selectedPokemon != nil {

// 	// 			result := catchProbability(selectedPokemon.CatchRate)

// 	// 			cacthResult := fmt.Sprintf("attemp = %d, pokemon name = %s, result = %s\n", turn, selectedPokemon.Name, result)

// 	// 			bucketResult[turn-1] = cacthResult

// 	// 			fmt.Printf("You attempted to catch %s (%s type) with a catch rate of %d%%: %s\n",
// 	// 				selectedPokemon.Name, selectedPokemon.Type, selectedPokemon.CatchRate, result)

// 	// 			fmt.Println(bucketResult)
// 	// 		} else {
// 	// 			fmt.Println("Invalid Pokémon ID. Please try again.")
// 	// 		}

// 	// 	} else {
// 	// 		fmt.Println("JATAH GACHA HABIS, SILAKAN TOP UP LAGI")
// 	// 		break
// 	// 	}
// 	// }

// }
package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
	wg      sync.WaitGroup
)

func increment() {
	defer wg.Done()
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	numGoroutines := 100

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go increment()
	}
	wg.Wait()

	fmt.Println("Final Counter:", counter)
}
