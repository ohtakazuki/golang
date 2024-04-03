package main

import "fmt"

func main() {
	subtotal := 899
	cTax := tax(subtotal)
	total := subtotal + cTax

	fmt.Printf("%d円の商品の税込価格は%d円（消費税は%d円）です。\n", subtotal, total, cTax)
}

func tax(subtotal int) int {
	rate := 0.1
	ret := float64(subtotal) * rate
	return int(ret)
}
