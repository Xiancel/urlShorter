package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"urlshort/history"
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
		fmt.Println("\nВведіть URL для скорочення (або 'exit' / 'history'): ")
		inputs := getStrInput("> ")

		// перевірка на ввід 'exit'
		if inputs == "exit" {
			fmt.Println("\n😭 Ceeeee-saaaaar...")
			time.Sleep(1 * time.Millisecond)
			fmt.Println("Сподіваємось, вам сподобався наш URLshortener!")
			fmt.Println("👋 До побачення!")
			return
		}
		// перевірка на ввід 'history'
		if inputs == "history" {
			// перевірка на наявність історії
			if len(history.HistoryList) == 0 {
				fmt.Println("Історія порожня.")
				continue
			} else {
				// вивід історії скорочування
				fmt.Println("🗂 MUDA MUDA MUDA... а ні, це просто історія скорочень.")
				time.Sleep(1 * time.Second)
				fmt.Println("\n📜 Історія скорочень: ")
				for i, h := range history.HistoryList {
					fmt.Printf("%d.\n🔗 %s\n⭐ %s\n", i+1, h.Original, h.Short)
				}
				continue
			}
		}
		fmt.Println("Oh? You’re approaching the API?")

		time.Sleep(1 * time.Second)
		fmt.Println("\n📡 Відправляю запит до TinyURL API...")

		// передавання url користувача до функції shortingurl
		res, err := shorter.ShortingUrl(inputs)
		// обробка помилок
		if err != nil {
			fmt.Println("\nYare Yare Daze... Сталася помилка!\n")

			time.Sleep(1 * time.Second)
			if strings.HasPrefix(err.Error(), "api: ") {
				fmt.Println("❌ Помилка від TinyURL:", strings.TrimPrefix(err.Error(), "api: "))
			} else if strings.HasPrefix(err.Error(), "network: ") {
				fmt.Println("❌ Мережева помилка:", strings.TrimPrefix(err.Error(), "network: "))
			} else {
				fmt.Println("❌ Невідома помилка: ", err)
			}
			continue
		}

		// додавання результату скорочення до історії
		history.HistoryList = append(history.HistoryList, history.UrlHistory{Original: inputs, Short: res})

		// вивід інформації користувачу про успишну скорочення

		fmt.Println("\n💥 ORA ORA ORA!")

		time.Sleep(1 * time.Second)

		fmt.Println("💥 OORAAA")
		fmt.Println("\n✅ Успішно скорочено!")
		fmt.Println("🔗 Оригінал: ", inputs)
		fmt.Println("⭐ Короткий: ", res)
	}
}
