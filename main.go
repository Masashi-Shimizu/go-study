package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 2
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

var i, j int = 1, 2

/**
定数は、文字(character)、文字列(string)、boolean、数値(numeric)のみで使えます。

なお、定数は := を使って宣言できません。
*/
const Pi = 3.14

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 31))
	doSwap()
	doSplit()

	k := 3 // 暗黙的な型の変数宣言

	fmt.Println(i, j, k)
	fmt.Printf("Type : %T  Value : %v\n", ToBe, ToBe)
	fmt.Printf("Type : %T  Value : %v\n", MaxInt, MaxInt)
	fmt.Printf("Type : %T  Value : %v\n", z, z)
	fmt.Println(Pi)

	doNumericConstant()
}

func add(a int, b int) int {
	return a + b
}

func doSwap() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func swap(x, y string) (string, string) {
	return y, x
}

func doSplit() {
	fmt.Println(split(17))
}

/**
この戻り値の名前は、戻り値の意味を示す名前とすることで、関数のドキュメントとして表現するようにしましょう。

名前をつけた戻り値の変数を使うと、 return ステートメントに何も書かずに戻すことができます。
これを "naked" return と呼びます。
例のコードのように、naked returnステートメントは、短い関数でのみ利用すべきです。
長い関数で使うと読みやすさ( readability )に悪影響があります。
*/
func split(sum int) (x, y int) {
	x = sum*4 - 9
	y = sum - x
	return
}

func doNumericConstant() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}
