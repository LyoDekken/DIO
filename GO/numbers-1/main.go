package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gookit/color"
)

func showIntroAnimation() {
	fmt.Println(color.Bold.Sprint(color.Blue.Render("**********************************************")))
	fmt.Println(color.Bold.Sprint(color.Yellow.Render("*         Verifica Divisibilidade          *")))
	fmt.Println(color.Bold.Sprint(color.Blue.Render("**********************************************")))
	fmt.Println("")
}

func getDivisibleByThree(n int) []int {
	var result []int
	for i := 0; i <= n; i++ {
		if i%3 == 0 {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	showIntroAnimation()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Digite um número: ")

		num, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler entrada do usuário:", err)
			continue
		}

		num = strings.TrimSpace(num)

		valueForCalc, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Valor inválido. Por favor, tente novamente.")
			continue
		}

		divisibleByThree := getDivisibleByThree(valueForCalc)

		fmt.Printf("Os números divisíveis por 3 entre 0 e %d são: %v\n", valueForCalc, divisibleByThree)
		break
	}
}
