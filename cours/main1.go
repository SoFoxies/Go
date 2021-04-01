package main //package main obliger d'avoir func main

import (
	//"fmt"
	//"wsf.fr/cours/animal"

	"wsf.fr/cours/pokemon"
)

func main() {

	pokedex := []pokemon.Pokemon{
		&pokemon.Bulbizarre{},
		&pokemon.Bulbizarre{},
		&pokemon.Carapuce{},
		&pokemon.Carapuce{},
		&pokemon.Salameche{},
		&pokemon.Salameche{},
	}

	for _, p := range pokedex {
		p.Attack()
	}
}
	//	var txt string
	//	fmt.Println(txt)

	//	var age uint = 19 // u = unsined integer
	//	if age < 18 {
	//		fmt.Println("Not authorized")
	//	} else {
	//		fmt.Println("Authorized")
	//	}

	//var animals = []string{
	//"Medor",
	//"waf waf",
	//} //animal va etre tableau de string, donc toute les valeurs vont etre string

	//fmt.Println(animals)

	//	var ages = [5]int{
	//		7,8,3,
	//	}

	//fmt.Println(ages)
	//for i, age := range ages { //:= = pas de variable definit comme i et age pas definit
	//	fmt.Println("Index:", i)
	//	fmt.Println("age:", age)
	//}

	//	var animals = map[string]int{
	//		"Medor":  7,
	//		"Machin": 8,
	//		"Truc":   3,
	//	}

	//	for name, age := range animals {
	//		fmt.Println("animal :", name, age)
	//	}

//	medor := &animal.Animal{ //& = definir un nouveau pointeur, medor nest plus type animal mais pointeur dee type animal, animal. c'est mon fichier
//		Name: "Medor",
//		Age:  7,
//	}

	//	machin := Animal{
	//		Name: "Machin",
	//		Age:  8,
	//	}

	//	var animals = []Animal{
	//		medor,
	//		machin,
	//	}
	//	for i, a := range animals {
	//		if a.Name == "" {
	//			fmt.Println("animal n'a pas de nom, index", i)
	//		}
	//	}

//	fmt.Println("Medor fete son anniv, c'est la teuf")
//	fmt.Println("Medor a", medor.Age)
//	medor.Birthday()
//	fmt.Println("Medor a", medor.Age)




