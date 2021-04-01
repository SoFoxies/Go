package animal

type Animal struct {
	Name string
	Age  int
}

func (a *Animal) Birthday() { //si on met pas pointeur accrementation marche pas, pointeur donne acces a la valeur d'une variable qui existe deja
	a.Age++

} //methode
