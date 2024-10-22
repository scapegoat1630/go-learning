package system_design

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// 内存中的短链接存储，实际生产环境应使用数据库
type URLShortener struct {
	mu          sync.RWMutex
	shortToLong map[string]string
	longToShort map[string]string
}

// 初始化URLShortener
func NewURLShortener() *URLShortener {
	return &URLShortener{
		shortToLong: make(map[string]string),
		longToShort: make(map[string]string),
	}
}

// 生成短链接
func (u *URLShortener) generateShortURL(longURL string) string {
	// 这里使用base62编码来生成短链接，可以根据需要调整长度
	const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const shortURLLength = 6
	var encoded bytes.Buffer
	encoder := base64.NewEncoder(base64.NewEncoding(charset), &encoded)
	counter := uint64(1) // 简单的计数器，实际应使用更复杂的唯一标识符生成策略
	u.mu.Lock()
	defer u.mu.Unlock()
	for {
		encoder.Write([]byte(longURL + strconv.FormatUint(counter, 10)))
		shortURL := encoded.String()[:shortURLLength] // 截取前6位作为短链接
		if _, exists := u.shortToLong[shortURL]; !exists {
			u.shortToLong[shortURL] = longURL // 占位，实际存储长链接时再赋值
			u.longToShort[longURL] = shortURL // 占位，用于反向查找（这里需要更复杂的逻辑）
			return shortURL
		}
		counter++
		encoded.Reset() // 重置缓冲区
	}

}

// 添加长链接到短链接映射
func (u *URLShortener) addURL(longURL string) string {
	shortURL := u.generateShortURL(longURL)
	u.mu.Lock()
	u.shortToLong[shortURL] = longURL
	u.longToShort[longURL] = shortURL
	u.mu.Unlock()
	return shortURL
}

// 处理短链接请求
func (u *URLShortener) shortURLHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:] // 去掉开头的斜杠
	u.mu.RLock()
	longURL, exists := u.shortToLong[shortURL]
	u.mu.RUnlock()
	if exists {
		http.Redirect(w, r, longURL, http.StatusMovedPermanently)
	} else {
		http.Error(w, "Short URL not found", http.StatusNotFound)
	}
}

// 处理添加长链接请求
func (u *URLShortener) addURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}
	shortURL := u.addURL(longURL)
	fmt.Fprintf(w, "Short URL: /%s\n", shortURL)
}

func main() {
	urlShortener := NewURLShortener()

	http.HandleFunc("/", urlShortener.shortURLHandler)
	http.HandleFunc("/add", urlShortener.addURLHandler)

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
