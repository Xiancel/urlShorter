package shorter

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// функція скорочення url
func ShortingUrl(originalUrl string) (string, error) {

	// перевірка на валідність url
	if !checkUrl(originalUrl) {
		return "", fmt.Errorf("api: Error: Invalid URL")
	}

	// формування url для запроса до tinyurl
	originalUrl = url.QueryEscape(originalUrl)
	apiUrl := "http://tinyurl.com/api-create.php?url=" + originalUrl

	// get запрос до tinyurl
	resp, err := http.Get(apiUrl)
	// обробка помилки
	if err != nil {
		return "", fmt.Errorf("network: %v", err)
	}
	// закриття запроса
	defer resp.Body.Close()

	// читання сформованого url
	body, err := io.ReadAll(resp.Body)
	// обробка помилки
	if err != nil {
		return "", fmt.Errorf("network: %v", err)
	}

	// створення result
	result := string(body)

	// обробка помилки
	if strings.HasPrefix(result, "Error") {
		return "", fmt.Errorf("api: %s", result)
	}
	return result, nil
}

// функція checkUrl для перевірки url на наявність протокола
func checkUrl(url string) bool {
	return strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://")
}
