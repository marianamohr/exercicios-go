package main

import (
	"fmt"
)

func validadorFor(arr []int, alvo int) (int, int, bool) {
	for indexOut, valorOut := range arr {

		for IndexIn, valorIn := range arr {
			soma := valorIn + valorOut
			if soma == alvo {
				fmt.Println(valorOut, "  + ", valorIn, " =  ", soma)
				return indexOut, IndexIn, false
			}
		}
	}
	return 0, 0, true
}

func validador(arr []int, alvo int) ([]int, bool) {
	dict := make(map[int]int)
	ret := make([]int, 0)
	for indexOut, valorOut := range arr {
		dict[valorOut] = indexOut
	}
	fmt.Println(dict)
	for value, _ := range dict {
		complemento := alvo - value
		_, err := dict[complemento]
		if err == true {
			fmt.Println(complemento, "está na posição", dict[complemento])
			ret = append(ret, dict[complemento])
		}

	}
	return ret, true
}

func main() {
	fmt.Println("Hello")

	arr := []int{5, 2, 9, 11}
	// Retorno esperado 0, 3
	arr, conflito := validador(arr, 20)
	fmt.Println(arr, conflito)

}
