package history

// структура urlHistory
type UrlHistory struct {
	Original string // url який вводе користувачь
	Short    string // скорочений url
}

// слайс для збегання історії url
var HistoryList []UrlHistory
