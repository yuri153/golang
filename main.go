package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitorCount = 5
const delayInSeconds = 5

func main() {
	for {
		showIntroduction()

		showMenu()

		switch readCommand() {
		case 1:
			monitor()
		case 2:
			showLogs()
		case 0:
			exit()
		default:
			exitError()
		}
	}
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func showIntroduction() {
	name := "Yuri"
	version := 1.0
	fmt.Println("Olá sr.", name)
	fmt.Println("Este programa está na versão", version)
	fmt.Println()
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println()
	fmt.Println("O comando escolhido foi", command)

	return command
}

func exit() {
	fmt.Println("Saindo do programa")
	os.Exit(0)
}

func exitError() {
	fmt.Println("Não conheço este comando")
	os.Exit(-1)
}

func monitor() {
	fmt.Println("Monitorando...")
	sites := getSitesToMonitor()

	for i := 0; i < monitorCount; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			siteTest(site)
		}
		time.Sleep(delayInSeconds * time.Second)
		fmt.Println()
	}
}

func siteTest(site string) {
	response, err := http.Get(site)

	verifyError(err)

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", response.StatusCode)
	}
}

func showLogs() {
	fmt.Println("Exibindo Logs...")
}

func getSitesToMonitor() []string {
	var sites []string

	file, err := os.Open("sites-to-monitor.txt")

	verifyError(err)

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func verifyError(err error) {
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
}
