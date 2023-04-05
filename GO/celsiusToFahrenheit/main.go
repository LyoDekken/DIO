package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gookit/color"
)

func main() {
	showIntroAnimation()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nQual conversão você deseja fazer?")
		fmt.Println("1 - Celsius para Fahrenheit")
		fmt.Println("2 - Fahrenheit para Celsius")
		fmt.Print("Escolha uma opção (1 ou 2): ")

		option, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler entrada do usuário:", err)
			continue
		}

		option = strings.TrimSpace(option)

		if option != "1" && option != "2" {
			fmt.Println("Opção inválida. Por favor, tente novamente.")
			continue
		}

		fmt.Print("Digite a temperatura a ser convertida: ")
		temperatureStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler entrada do usuário:", err)
			continue
		}

		temperatureStr = strings.TrimSpace(temperatureStr)

		temperature, err := strconv.ParseFloat(temperatureStr, 64)
		if err != nil {
			fmt.Println("Valor inválido. Por favor, tente novamente.")
			continue
		}

		var convertedTemp float64

		if option == "1" {
			convertedTemp = celsiusToFahrenheit(temperature)
			fmt.Printf("%.2f°C equivale a %.2f°F\n", temperature, convertedTemp)
		} else {
			convertedTemp = fahrenheitToCelsius(temperature)
			fmt.Printf("%.2f°F equivale a %.2f°C\n", temperature, convertedTemp)
		}

		break
	}
}

func showIntroAnimation() {
	fmt.Println(color.Bold.Sprint(color.Blue.Render("**********************************************")))
	fmt.Println(color.Bold.Sprint(color.Yellow.Render("*            CONVERSOR DE TEMPERATURA         *")))
	fmt.Println(color.Bold.Sprint(color.Blue.Render("**********************************************")))
	fmt.Println("")
}

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 1.8) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) / 1.8
}