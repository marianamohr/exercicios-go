package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func introducao() {
	var nome string = "Mari"
	var versao float32 = 1.1
	fmt.Println("Olá sra.", nome)
	fmt.Println("Versao do programa", versao)
}

func leComando() int {
	var comando int
	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")
	fmt.Println("* 1 - Iniciar monitoramento *")
	fmt.Println("* 2 - Exibir Logs           *")
	fmt.Println("* 0 - Sair                  *")
	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")
	fmt.Scan(&comando)
	return comando
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer arquivo.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " site: " + site + " online: " + strconv.FormatBool(status) + "\n")

}
func lerSitesdoArquivo() []string {

	arquivo, err := os.Open("sites.txt")
	defer arquivo.Close()

	if err != nil {
		fmt.Println(err)
	}
	leitor := bufio.NewReader(arquivo)
	var sitelist []string
	for {
		linha, err := leitor.ReadString('\n')
		if err == io.EOF {
			break
		}
		linha = strings.TrimSpace(linha)
		sitelist = append(sitelist, linha)
	}
	return sitelist

}

func monitoramento() {
	fmt.Println("Monitorando...")

	sites := lerSitesdoArquivo()

	for _, value := range sites {

		response, _ := http.Get(value)

		if response.StatusCode == 200 {

			fmt.Println("O site", value, "acessado com sucesso")
			registraLog(value, true)
		} else {
			fmt.Println("O site da ", value, "está com problema", response.StatusCode)
			registraLog(value, false)
		}
	}

}
func imprimeLogs() {
	arquivo, err := os.ReadFile("logs.txt")

	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(string(arquivo))
}

func main() {

	introducao()

	for {
		comandoLido := leComando()
		switch comandoLido {
		case 1:

			monitoramento()
		case 2:
			imprimeLogs()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("opção nao encontrada")
			os.Exit(-1)

		}
	}

}
