package main

import (
	"fmt"
	"reflect"
)

type Students struct {
	name string
}

type Course struct {
	nome   string
	alunos []Students
}
type Ead struct {
	curso   Course
	webSite string
}

func (c *Course) register(s Students) {
	c.alunos = append(c.alunos, s)
}

func main() {
	fmt.Println("Hello")

	java := Course{"Java", make([]Students, 0)}

	newEad := Ead{java, "batinha"}

	aluno2 := Students{"Cami"}
	pst := &aluno2
	fmt.Println(reflect.TypeOf(aluno2), aluno2.name)
	fmt.Println(reflect.TypeOf(pst), pst.name)

	aluno := Students{"Mari"}
	aluno3 := Students{"Gael"}
	newEad.curso.register(aluno)
	newEad.curso.register(aluno2)
	newEad.curso.register(aluno3)

	fmt.Println(newEad)
}
