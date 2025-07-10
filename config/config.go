package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"urlshort/shorter"
)

// читачь строки
var reader *bufio.Reader = bufio.NewReader(os.Stdin)

// приймає текстове введеня
func getStrInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// функція ініціалізаціі программи
func Init() {
	fmt.Println("🔗 URL Shortener")
	// головний цикл
	for {
		// запросс url у користувача
		fmt.Println("\nВведіть URL для скорочення (або 'exit' для виходу): ")
		inputs := getStrInput("> ")

		// перевірка на ввід 'exit'
		if inputs == "exit" {
			return
		}
		fmt.Println("\n📡 Відправляю запит до TinyURL API...")

		// передавання url користувача до функції shortingurl
		res, err := shorter.ShortingUrl(inputs)
		// обробка помилок
		if err != nil {
			if strings.HasPrefix(err.Error(), "api: ") {
				fmt.Println("❌ Помилка від TinyURL:", strings.TrimPrefix(err.Error(), "api: "))
			} else if strings.HasPrefix(err.Error(), "network: ") {
				fmt.Println("❌ Мережева помилка:", strings.TrimPrefix(err.Error(), "network: "))
			} else {
				fmt.Println("❌ Невідома помилка: ", err)
			}
			continue
		}
		// вивід інформації користувачу про успишну скорочення
		fmt.Println("\n✅ Успішно скорочено!")
		fmt.Println("🔗 Оригінал: ", inputs)
		fmt.Println("⭐ Короткий: ", res)
	}
}
