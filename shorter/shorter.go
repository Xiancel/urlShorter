package shorter

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func ShortingUrl(originalUrl string) (string, error) {

	if !checkUrl(originalUrl) {
		return "", fmt.Errorf("api: Error: Invalid URL")
	}

	originalUrl = url.QueryEscape(originalUrl)
	apiUrl := "http://tinyurl.com/api-create.php?url=" + originalUrl

	resp, err := http.Get(apiUrl)
	if err != nil {
		return "", fmt.Errorf("network: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("network: %v", err)
	}

	result := string(body)

	if strings.HasPrefix(result, "Error") {
		return "", fmt.Errorf("api: %s", result)
	}
	return result, nil
}

func checkUrl(url string) bool {
	return strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://")
}
