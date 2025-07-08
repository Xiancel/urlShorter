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
func Init() {
	fmt.Println("🔗 URL Shortener")
	for {
		fmt.Println("\nВведіть URL для скорочення (або 'exit' для виходу): ")
		inputs := getStrInput("> ")

		fmt.Println("\n📡 Відправляю запит до TinyURL API...")

		if inputs == "exit" {
			return
		}

		// if !checkUrl(inputs) {
		// 	fmt.Println("❌ Invalid URL")
		// 	return
		// }

		res, err := shorter.ShortingUrl(inputs)
		if err != nil {
			fmt.Println("❌ Errors: ", err)
			return
		}
		fmt.Println("\n✅ Успішно скорочено!")
		fmt.Println("🔗 Оригінал: ", inputs)
		fmt.Println("⭐ Короткий: ", res)
	}
}

func checkUrl(url string) bool {
	return strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://")
}
