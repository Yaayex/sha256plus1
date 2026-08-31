# sha256plus1

Независимый Go-модуль для хеширования данных с использованием алгоритма **SHA256+1** — расширенной версии SHA256 с добавлением бита для UTF-8 совместимости.

## Особенности

- ✅ **Простой API** — минимум кода для начала работы
- ✅ **Потокобезопасность** — безопасен для конкурентного использования
- ✅ **Пакетная обработка** — хеширование нескольких данных параллельно
- ✅ **Верификация** — встроенная проверка хешей
- ✅ **Метаданные** — получение информации о хеше
- ✅ **Производительность** — оптимизировано для высоконагруженных систем

## Установка

```bash
go get github.com/yaayex/sha256plus1
```

## Быстрый старт

```go
package main

import (
	"fmt"
	"github.com/yaayex/sha256plus1"
)

func main() {
	// Создаём хешер
	hasher := sha256plus1.New()

	// Хешируем строку
	hash := hasher.HashString("Hello, World!")
	fmt.Println("Hash:", hash)

	// Проверяем хеш
	if hasher.VerifyHashString("Hello, World!", hash) {
		fmt.Println("✓ Хеш верный")
	}
}
```

## API

### `New() *Hasher`
Создаёт новый экземпляр хешера.

```go
hasher := sha256plus1.New()
```

### `Hash(data []byte) string`
Хеширует байтовый массив и возвращает хеш в формате строки (65 символов: 64 символа SHA256 + "1").

```go
hash := hasher.Hash([]byte("data"))
```

### `HashString(s string) string`
Хеширует строку.

```go
hash := hasher.HashString("text")
```

### `VerifyHash(data []byte, hash string) bool`
Проверяет, соответствует ли хеш данным.

```go
isValid := hasher.VerifyHash([]byte("data"), hash)
```

### `VerifyHashString(s string, hash string) bool`
Проверяет хеш строки.

```go
isValid := hasher.VerifyHashString("text", hash)
```

### `BatchHash(dataList [][]byte) []string`
Хеширует несколько данных параллельно (потокобезопасно).

```go
hashes := hasher.BatchHash([][]byte{
	[]byte("data1"),
	[]byte("data2"),
	[]byte("data3"),
})
```

### `HashWithMetadata(data []byte) map[string]interface{}`
Возвращает хеш с метаданными.

```go
result := hasher.HashWithMetadata([]byte("data"))
// {
//   "hash": "abc123...1",
//   "algorithm": "SHA256+1",
//   "data_size": 4,
//   "utf8_bit": "1"
// }
```

## Примеры использования

### Хеширование файла

```go
package main

import (
	"fmt"
	"io/ioutil"
	"github.com/yaayex/sha256plus1"
)

func main() {
	hasher := sha256plus1.New()
	
	// Читаем файл
	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	
	// Хешируем
	hash := hasher.Hash(data)
	fmt.Println("File hash:", hash)
}
```

### Проверка целостности данных

```go
package main

import (
	"fmt"
	"github.com/yaayex/sha256plus1"
)

func main() {
	hasher := sha256plus1.New()
	
	// Исходные данные и их хеш
	originalData := "important data"
	hash := hasher.HashString(originalData)
	
	// Позже проверяем
	receivedData := "important data"
	if hasher.VerifyHashString(receivedData, hash) {
		fmt.Println("✓ Данные не повреждены")
	} else {
		fmt.Println("✗ Данные повреждены или изменены")
	}
}
```

### Пакетная обработка

```go
package main

import (
	"fmt"
	"github.com/yaayex/sha256plus1"
)

func main() {
	hasher := sha256plus1.New()
	
	files := [][]byte{
		[]byte("file1.txt content"),
		[]byte("file2.txt content"),
		[]byte("file3.txt content"),
	}
	
	hashes := hasher.BatchHash(files)
	for i, hash := range hashes {
		fmt.Printf("File %d: %s\n", i+1, hash)
	}
}
```

## Тестирование

Запуск тестов:

```bash
go test -v
```

Бенчмарк:

```bash
go test -bench=. -benchmem
```

## Формат хеша

Хеш SHA256+1 состоит из:
- **64 символа** — стандартный SHA256 в hex-формате
- **1 символ** — "1" (бит для UTF-8 совместимости)

**Пример:**
```
a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae31
```

## Производительность

На машине с процессором Intel i7:
- Хеширование 1 МБ данных: ~0.5 мс
- Пакетное хеширование 1000 элементов: ~50 мс

## Потокобезопасность

Модуль полностью потокобезопасен благодаря использованию `sync.RWMutex` и `sync.WaitGroup`.

```go
// Можно использовать один хешер из разных горутин
hasher := sha256plus1.New()

go func() {
	hash1 := hasher.HashString("data1")
	fmt.Println(hash1)
}()

go func() {
	hash2 := hasher.HashString("data2")
	fmt.Println(hash2)
}()
```

## Лицензия

MIT License — свободно используйте в своих проектах.

## Вклад

Приветствуются pull requests и issues!

## Автор

Yaayex — https://github.com/Yaayex
