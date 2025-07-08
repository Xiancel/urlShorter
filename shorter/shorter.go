package shorter

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func ShortingUrl(originalUrl string) (string, error) {

	originalUrl = url.QueryEscape(originalUrl)
	apiUrl := "http://tinyurl.com/api-create.php?url=" + originalUrl

	resp, err := http.Get(apiUrl)
	if err != nil {
		return "", fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}
	return string(body), nil
}
