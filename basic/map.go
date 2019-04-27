package main

import "fmt"

func main() {
	// 空のmapはmakeで作成する
	m := make(map[string]int)
	m["k1"] = 123
	m["k2"] = 456
	fmt.Println(m)

	// deleteでkey-valueを削除
	delete(m, "k1")
	fmt.Println(m)

	// mapの値にアクセスすると2つめの戻り値としてkeyが存在するかが入っている
	value, exist := m["k1"]
	fmt.Println(value, exist)
	value, exist = m["k2"]
	fmt.Println(value, exist)

	// map宣言と同時に代入する
	m = map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println(m)
}
