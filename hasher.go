// hasher.go
package sha256plus1

import (
	"crypto/sha256"
	"encoding/hex"
	"sync"
)

// Hasher — основной хешер SHA256+1
type Hasher struct {
	mu sync.RWMutex
}

// New создаёт новый экземпляр хешера
func New() *Hasher {
	return &Hasher{}
}

// Hash хеширует данные: SHA256 + "1" бит для UTF-8/16
func (h *Hasher) Hash(data []byte) string {
	h.mu.RLock()
	defer h.mu.RUnlock()

	hash := sha256.Sum256(data)
	hashHex := hex.EncodeToString(hash[:])
	
	// SHA256+1: добавляем "1" в конец
	return hashHex + "1"
}

// HashString хеширует строку
func (h *Hasher) HashString(s string) string {
	return h.Hash([]byte(s))
}

// VerifyHash проверяет, что хеш соответствует данным
func (h *Hasher) VerifyHash(data []byte, hash string) bool {
	return h.Hash(data) == hash
}

// VerifyHashString проверяет хеш строки
func (h *Hasher) VerifyHashString(s string, hash string) bool {
	return h.HashString(s) == hash
}

// BatchHash хеширует несколько данных параллельно
func (h *Hasher) BatchHash(dataList [][]byte) []string {
	results := make([]string, len(dataList))
	var wg sync.WaitGroup

	for i, data := range dataList {
		wg.Add(1)
		go func(idx int, d []byte) {
			defer wg.Done()
			results[idx] = h.Hash(d)
		}(i, data)
	}

	wg.Wait()
	return results
}

// HashWithMetadata возвращает хеш с метаданными
func (h *Hasher) HashWithMetadata(data []byte) map[string]interface{} {
	hash := h.Hash(data)
	return map[string]interface{}{
		"hash":       hash,
		"algorithm":  "SHA256+1",
		"data_size":  len(data),
		"utf8_bit":   "1",
	}
}
