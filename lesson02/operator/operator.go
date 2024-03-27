package main

import "fmt"

func main() {
	// 算術演算子
	a := 10
	b := 3

	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b
	remainder := a % b

	fmt.Println("加算: ", sum)       // 13
	fmt.Println("減算: ", diff)      // 7
	fmt.Println("乗算: ", product)   // 30
	fmt.Println("割算: ", quotient)  // 3
	fmt.Println("余り: ", remainder) // 1

	// 複合代入演算子
	x := 10
	x += 5                         // x = x + 5
	fmt.Printf("x += 5: %d\n", x)  // 15
	x -= 3                         // x = x - 3
	fmt.Printf("x -= 3: %d\n", x)  // 12
	x *= 2                         // x = x \* 2
	fmt.Printf("x *= 2: %d\n", x)  // 24
	x /= 4                         // x = x / 4
	fmt.Printf("x /= 4: %d\n", x)  // 6
	x %= 5                         // x = x % 5
	fmt.Printf("x %%= 5: %d\n", x) // 1

	// インクリメント・デクリメント演算子
	x = 5
	x++                        // x = x + 1
	fmt.Printf("x++: %d\n", x) // 6
	x--                        // x = x - 1
	fmt.Printf("x--: %d\n", x) // 5

	// 比較演算子
	a = 10
	b = 5
	fmt.Printf("a > b: %t\n", a > b)   // true
	fmt.Printf("a < b: %t\n", a < b)   // false
	fmt.Printf("a >= b: %t\n", a >= b) // true
	fmt.Printf("a <= b: %t\n", a <= b) // false
	fmt.Printf("a == b: %t\n", a == b) // false
	fmt.Printf("a != b: %t\n", a != b) // true

	// 論理演算子
	x = 15
	fmt.Printf("10 < x && x < 20: %t\n", 10 < x && x < 20) // true
	fmt.Printf("x < 10 || x > 20: %t\n", x < 10 || x > 20) // false
	fmt.Printf("!(x == 15): %t\n", !(x == 15))             // false

	// ビット演算子
	a = 0b1100                       // 10進数で12を2進数表示
	b = 0b1010                       // 10進数で10を2進数表示
	fmt.Printf("a & b: %b\n", a&b)   // 1000
	fmt.Printf("a | b: %b\n", a|b)   // 1110
	fmt.Printf("a ^ b: %b\n", a^b)   // 0110
	fmt.Printf("a << 1: %b\n", a<<1) // 11000
	fmt.Printf("a >> 1: %b\n", a>>1) // 110
	fmt.Printf("a &^ b: %b\n", a&^b) // 100
}
