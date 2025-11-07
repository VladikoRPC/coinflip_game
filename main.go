package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var avarageWins float64
	fmt.Println("Сыграй в монетку!")
	flipCount, flipWins, history := playGame()
	avarageWins = (float64(flipWins) / float64(flipCount)) * 100
	fmt.Printf("\nРезультаты сессии:\nВсего бросков: %d\nВсего побед: %d\nПроцент побед: %.2f\n", flipCount, flipWins, avarageWins)
	fmt.Println("\n----------------------")
	fmt.Println("История бросков:")
	for index, value := range history {
		fmt.Println(index+1, value)
	}
	fmt.Println("Спасибо за игру! Приходи еще!")
}
func userChoice() string {
	var userChoice string
	for {
		fmt.Print("Выбирай: орел или решка?\n")
		fmt.Scan(&userChoice)
		if userChoice == "орел" || userChoice == "решка" {
			return userChoice
		}
		fmt.Println("Некорректный ввод, попробуйте снова.")
	}
}
func playGame() (int, int, []string) {
	history := []string{}

	var flipCount int
	var flipWins int

	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		options := []string{"орел", "решка"}
		choice := userChoice()
		flip := options[r.Intn(len(options))]
		fmt.Println("Результат броска: ", flip)
		flipCount++
		if choice == flip {
			fmt.Println("Вы угадали! 🎉")
			flipWins++
			history = append(history, "Победа")
		} else {
			fmt.Println("Не повезло! Попробуйте снова.")
			history = append(history, "Поражение")
		}

		repeat := isRepeat()
		if !repeat {
			break
		}
	}
	return flipCount, flipWins, history
}
func isRepeat() bool {
	var userChoice string
	fmt.Println("Хотите повторить? (да/нет)")
	fmt.Scan(&userChoice)
	return userChoice == "да" || userChoice == "ДА"
}
