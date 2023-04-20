package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 2
const delay = 5

func main() {
	HelloWord()

	for {
		ExibeMenu()

		comando := LeComando()

		switch comando {
		case 1:
			IniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa!")
			os.Exit(0)
		default:
			fmt.Println("Erro! Comando digitado não existe :(")
			os.Exit(-1)
		}
	}

}

func HelloWord() {
	nome := "Maria F. Inês"
	versao := 1.2
	fmt.Println("Olá mundo, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func ExibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs...")
	fmt.Println("0- Sair do Programa")
}

func LeComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func IniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := LeSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			TestaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func TestaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		RegistraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
		RegistraLog(site, false)
	}
}

func LeSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()
	return sites
}

func RegistraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}
