package main

import "fmt"

// Petインターフェースを定義
type Pet interface {
	Name() string
	Eat()
	Play()
	Sleep()
}

// Dog構造体を定義
type Dog struct {
	name   string
	energy int
}

// Dog構造体にPetインターフェースのメソッドを実装
func (d *Dog) Name() string {
	return d.name
}

func (d *Dog) Eat() {
	d.energy += 10
	fmt.Printf("%sは食事をしています。エネルギー: %d\n", d.name, d.energy)
}

func (d *Dog) Play() {
	d.energy -= 5
	fmt.Printf("%sは遊んでいます。エネルギー: %d\n", d.name, d.energy)
}

func (d *Dog) Sleep() {
	d.energy += 20
	fmt.Printf("%sは眠っています。エネルギー: %d\n", d.name, d.energy)
}

// Cat構造体を定義
type Cat struct {
	name string
	mood string
}

// Cat構造体にPetインターフェースのメソッドを実装
func (c *Cat) Name() string {
	return c.name
}

func (c *Cat) Eat() {
	c.mood = "幸せ"
	fmt.Printf("%sは食事をしています。気分: %s\n", c.name, c.mood)
}

func (c *Cat) Play() {
	c.mood = "遊び心"
	fmt.Printf("%sは遊んでいます。気分: %s\n", c.name, c.mood)
}

func (c *Cat) Sleep() {
	c.mood = "眠い"
	fmt.Printf("%sは眠っています。気分: %s\n", c.name, c.mood)
}

func main() {
	// ペットのスライスを作成
	pets := []Pet{
		&Dog{name: "バディ", energy: 50},
		&Cat{name: "ウィスカー", mood: "普通"},
	}

	// ペットの情報を出力し、各アクションを実行
	for _, pet := range pets {
		fmt.Printf("%sの情報:\n", pet.Name())
		pet.Eat()
		pet.Play()
		pet.Sleep()
		fmt.Println()
	}
}
