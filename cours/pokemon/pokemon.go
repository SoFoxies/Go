package pokemon

import "fmt"

type Pokemon interface {
	Attack()
}

type Carapuce struct {
	Name string
	Age  int

	HasGlasses       bool
	NumberOfEcailles int
}

func (c *Carapuce) Attack() { //attack methode accesible type carapuce
	fmt.Println("Je suis Carapuce, et je t'attaque!! ")
}

type Bulbizarre struct {
	Name string
	Age  int

	NumberOfPetales int
}

func (b *Bulbizarre) Attack() {
	fmt.Println("Mwamwa tu ne me fais pas peur")
}

type Salameche struct {
	Name string
	Age  int
}

func (s *Salameche) Attack() {
	fmt.Println("Je te crame les poils")
}
