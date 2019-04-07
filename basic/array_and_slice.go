// https://qiita.com/YumaInaura/items/aec7857518102ff5e27c

package main

import "fmt"

func main() {
	// 配列
	var a1 [3]int           // 宣言のみ
	a2 := [3]int{1, 2, 3}   // 宣言と代入を同時に行う
	a3 := [...]int{1, 2, 3} // 配列サイズを自動設定
	a4 := [3]int{}          // 空配列

	fmt.Printf("a1:%v\n", a1)
	fmt.Printf("a2:%v\n", a2)
	fmt.Printf("a3:%v\n", a3)
	fmt.Printf("a4:%v\n", a4)

	// スライス
	var s1 []int         // 宣言のみ
	s2 := []int{1, 2, 3} // 宣言と代入を同時に行う
	s3 := []int{}        // 空スライス

	fmt.Printf("s1:%v\n", s1)
	fmt.Printf("s2:%v\n", s2)
	fmt.Printf("s3:%v\n", s3)
}
