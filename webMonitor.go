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

const sleepTime = 30
const testTurn = 3

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()
		fmt.Println("")
		switch comando {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando nao reconhecido!")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Matheus"
	versao := 1.1
	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("")
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O valor da variável comando é:", comandoLido)

	return comandoLido
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	// sites := []string{"http://compasso.com.br", "http://alura.com.br", "https://random-status-code.herokuapp.com"}

	sites := leSite()

	for i := 0; i < testTurn; i++ {
		for _, site := range sites {
			testeSite(site)
		}
		time.Sleep(sleepTime * time.Second)
		fmt.Println("")
	}
}

func testeSite(site string) {
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
		registrarLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta sem acesso, erro:", response.StatusCode)
		registrarLog(site, false)
	}
}

func leSite() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println(err)
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

func showLogs() {
	fmt.Println("Showing Logs...")

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}

func registrarLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " +
		site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
