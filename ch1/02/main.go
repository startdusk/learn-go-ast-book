package main

import (
	"fmt"
	"go/token"
)

func main() {
	a := token.Position{Filename: "hello.go", Line: 1, Column: 2}
	b := token.Position{Filename: "hello.go", Line: 1}
	c := token.Position{Filename: "hello.go"}

	d := token.Position{Line: 1, Column: 2}
	e := token.Position{Line: 1}
	f := token.Position{Column: 2}

	fmt.Println(a.String())
	fmt.Println(b.String())
	fmt.Println(c.String())
	fmt.Println(d.String())
	fmt.Println(e.String())
	fmt.Println(f.String())

	// Output:
	// hello.go:1:2
	// hello.go:1
	// hello.go
	// 1:2
	// 1
	// -

	// 行号从1开始，是必须的信息，如果缺少行号则输出“-”表示无效的位置
}
