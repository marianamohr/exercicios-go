package main

import (
	"fmt"
	"math/rand"
)

type Cliente struct {
	Nome      string
	Sobrenome string
}

func (c Cliente) NomeCompleto() string {
	return fmt.Sprintf("%s  %s", c.Nome, c.Sobrenome)
}

func main() {
	num := rand.Intn(100)
	fmt.Println(num)

	if num >= 50 {
		fmt.Println("maior que 50")
	} else {
		fmt.Println("menor que 50")
	}

	cliente := Cliente{"Mari", "Mohr"}
	fmt.Println(cliente.NomeCompleto())

	grades := []float32{8.5, 7, 6.5}

	var media float32

	for _, grade := range grades {
		media += grade
	}
	fmt.Println(media / float32(len(grades)))

	//for {
	//	counter := rand.Intn(100)
	//	if counter == 50 {
	//		break
	//	}
	//	fmt.Println(counter)
	//}

	var count int

	for count < 50 {
		count = rand.Intn(100)
		fmt.Println(count)

	}
}
