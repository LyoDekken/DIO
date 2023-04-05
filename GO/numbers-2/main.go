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

func isDivisibleByThree(n int) bool {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit
		n /= 10
	}
	return sum%3 == 0
}

func isDivisibleByFive(n int) bool {
	return n%5 == 0
}

func main() {
	showIntroAnimation()

	reader := bufio.NewReader(os.Stdin)

	// ler nomes dos usuários
	var username1, username2 string
	for {
		fmt.Print("Digite o nome do primeiro usuário: ")
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler entrada do usuário:", err)
			continue
		}
		username1 = strings.TrimSpace(name)
		break
	}
	for {
		fmt.Print("Digite o nome do segundo usuário: ")
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler entrada do usuário:", err)
			continue
		}
		username2 = strings.TrimSpace(name)
		break
	}

	// ler número escolhido pelos usuários
	var n int
	for {
		fmt.Print("Digite um número: ")
		num, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler entrada do usuário:", err)
			continue
		}
		num = strings.TrimSpace(num)
		n, err = strconv.Atoi(num)
		if err != nil {
			fmt.Println("Valor inválido. Por favor, tente novamente.")
			continue
		}
		break
	}

	// criar os arrays para armazenar os valores de PIN e PAN
	pinValues := make([]int, n+1)
	panValues := make([]int, n+1)

	// contar divisores de 3 e 5 e armazenar os valores nos arrays
	var pinIndex, panIndex int
	for i := 0; i <= n; i++ {
		if isDivisibleByThree(i) {
			pinValues[pinIndex] = i
			pinIndex++
		}
		if isDivisibleByFive(i) {
			panValues[panIndex] = i
			panIndex++
		}
	}

	// exibir resultados
	fmt.Printf("%s, o número de divisores de 3 de 0 a %d são: %d (PIN: %v)\n", username1, n, pinIndex, pinValues[:pinIndex])
	fmt.Printf("%s, o número de divisores de 5 de 0 a %d são: %d (PAN: %v)\n", username2, n, panIndex, panValues[:panIndex])
}
