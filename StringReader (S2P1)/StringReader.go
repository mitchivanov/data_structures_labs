package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Открываем файл с именем text.txt и читаем его содержимое
	file, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Print(err) // Если произошла ошибка, выводим ее на экран
	}
	// Преобразуем содержимое файла в строку
	content := string(file)
	// Удаляем пробелы в начале и в конце строки и разбиваем ее на подстроки по пробелам
	words := strings.Split(strings.TrimSpace(content), " ")
	// Выводим каждое слово на отдельной строке
	for _, word := range words {
		fmt.Println(word)
	}
}
