// hasher_test.go
package sha256plus1

import (
	"testing"
)

func TestHash(t *testing.T) {
	h := New()
	
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", ""},  // Проверим, что заканчивается на "1"
		{"world", ""},
		{"", ""},
	}

	for _, tt := range tests {
		result := h.HashString(tt.input)
		
		// Проверяем, что заканчивается на "1"
		if result[len(result)-1:] != "1" {
			t.Errorf("Hash должен заканчиваться на '1', получено: %s", result)
		}
		
		// Проверяем длину (SHA256 = 64 символа + "1" = 65)
		if len(result) != 65 {
			t.Errorf("Длина хеша должна быть 65, получено: %d", len(result))
		}
	}
}

func TestVerifyHash(t *testing.T) {
	h := New()
	data := []byte("test data")
	hash := h.Hash(data)

	if !h.VerifyHash(data, hash) {
		t.Error("Верификация хеша не прошла")
	}

	if h.VerifyHash([]byte("wrong data"), hash) {
		t.Error("Верификация должна была не пройти для неправильных данных")
	}
}

func TestBatchHash(t *testing.T) {
	h := New()
	dataList := [][]byte{
		[]byte("data1"),
		[]byte("data2"),
		[]byte("data3"),
	}

	results := h.BatchHash(dataList)

	if len(results) != len(dataList) {
		t.Errorf("Количество результатов должно быть %d, получено: %d", len(dataList), len(results))
	}

	for _, hash := range results {
		if len(hash) != 65 {
			t.Errorf("Каждый хеш должен быть 65 символов, получено: %d", len(hash))
		}
	}
}

func BenchmarkHash(b *testing.B) {
	h := New()
	data := []byte("benchmark test data")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Hash(data)
	}
}
