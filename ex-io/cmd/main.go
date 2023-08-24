package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your name: ")
	text, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %v how are you today?\n", text)
	newName := fmt.Sprintf("%v", strings.TrimSpace(text))
	fmt.Printf("Now, %v, without the extra line\n", newName)

	// Integers
	fmt.Print("Now, type an integer: ")
	nints, _ := reader.ReadString('\n')
	nint, errInt := strconv.ParseInt(strings.TrimSpace(nints), 10, 64)
	fmt.Println("Error ", errInt)
	if errInt != nil {
		log.Fatal("Invalid integer")
	}
	fmt.Printf("You typed: %d\n", nint)

	// Floating point
	fmt.Print("Now, type a float: ")
	nfs, _ := reader.ReadString('\n')
	nf, errf := strconv.ParseFloat(strings.TrimSpace(nfs), 64)
	if errf != nil {
		log.Fatal("Invalid float number")
	}
	fmt.Printf("You typed: %f\n", nf)

	// Numeros decimais i18n:

	p := message.NewPrinter(language.BrazilianPortuguese)
	p.Printf("%f", nf)

	// Creating a text file
	stringArr := []byte("Good morning.\nEat an Apple!\nHave a nice day!\n")
	// Permission: -rw-r--r--
	err := ioutil.WriteFile("/tmp/arq.txt", stringArr, 0644)
	check(err)

	// Reading a text file
	data, err := ioutil.ReadFile("/tmp/arq.txt")
	check(err)
	fmt.Printf("\nType: %T\n", data)
	fmt.Printf("\ndata: %v\n", data)
	batatinha := []uint8{71, 111, 111, 100, 32, 109, 100}
	fmt.Println(string(batatinha))
	textContent := string(data)
	var xablau string
	xablau = newName
	converted, err := strconv.Atoi(xablau)
	fmt.Println(textContent)
	fmt.Println(converted, err)

	list := make([]uint8, len(xablau))

	for i, char := range xablau {
		list[i] = uint8(char)
	}

	fmt.Println(xablau, "em uint8 é assim => ", list)
}
