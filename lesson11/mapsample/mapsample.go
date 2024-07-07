package main

import "fmt"

// MapWrapper は、ジェネリックなマップをラップする構造体です
type MapWrapper[K comparable, V any] struct {
	data map[K]V
}

// NewMapWrapper は、新しい MapWrapper のインスタンスを作成します
func NewMapWrapper[K comparable, V any]() *MapWrapper[K, V] {
	return &MapWrapper[K, V]{data: make(map[K]V)}
}

// Set は、指定されたキーと値をマップに設定します
func (w *MapWrapper[K, V]) Set(key K, value V) {
	w.data[key] = value
}

// Get は、指定されたキーに対応する値とその存在を示すブール値を返します
func (w *MapWrapper[K, V]) Get(key K) (V, bool) {
	value, ok := w.data[key]
	return value, ok
}

// Delete は、指定されたキーに対応する要素をマップから削除します
func (w *MapWrapper[K, V]) Delete(key K) {
	delete(w.data, key)
}

func main() {
	intStringMap := NewMapWrapper[int, string]()
	intStringMap.Set(1, "one")
	intStringMap.Set(2, "two")

	value, ok := intStringMap.Get(1)
	fmt.Println(value, ok) // 出力: one true

	intStringMap.Delete(1)
	value, ok = intStringMap.Get(1)
	fmt.Println(value, ok) // 出力:  false
}
