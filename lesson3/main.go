package main

import (
	. "fmt"
	. "lesson3/lruCache"
	"time"
)

// Реализовать LRU-cache (Без использования сторонних библиотек)
// При создании, необходимо указывать capacity кеша
// Для каждого элемента, указывать TTL при добавлении в кеш
// Реализовать в виде go module
// Добавить тесты, в т.ч. тестирование производительности и потребления памяти

func Printfln(str string, args ...interface{}) {
	Printf(str, args...)
	Println()
}

func main() {
	testTime := 5 * time.Second
	cache := NewLRUCache(10)
	cache.Add("key1", 50, testTime)

	res, ok := cache.Get("key1")
	res2, notOk := cache.Get("key2")
	time.Sleep(testTime)
	res3, notOk2 := cache.Get("key1")

	Printfln("Result 1 %d %v", res, ok)
	Printfln("Result 2 %d %v", res2, notOk)
	Printfln("Result 3 %d %v", res3, notOk2)
}
