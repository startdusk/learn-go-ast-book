package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	// src是要分析的代码
	var src = []byte(`println("你好, 世界")`)
	// 创建一个文件集，Token的位置信息必须通过文件集定位，
	// 并且需要通过文件集创建扫描器的Init方法需要的File参数
	var fset = token.NewFileSet()
	// 向fset文件集添加一个新的文件，文件名为“hello.go”，文件的长度就是src要分析代码的长度
	var file = fset.AddFile("hello.go", fset.Base(), len(src))

	// 初始化扫描器。
	// Init的第一个参数就是刚刚添加到fset的文件对象，
	// 第二个参数是要分析的代码，
	// 第三个nil参数表示没有自定义的错误处理函数，
	// 最后的scanner.ScanComments参数表示不用忽略注释Token
	var s scanner.Scanner
	s.Init(file, src, nil, scanner.ScanComments)
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)

		// Output:
		// hello.go:1:1    IDENT   "println"
		// hello.go:1:8    (       ""
		// hello.go:1:9    STRING  "\"你好, 世界\""
		// hello.go:1:25   )       ""
		// hello.go:1:26   ;       "\n"
	}
}
