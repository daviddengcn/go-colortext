go-colortext package [![GoSearch](http://go-search.org/badge?id=github.com%2Fdaviddengcn%2Fgo-colortext)](http://go-search.org/view?id=github.com%2Fdaviddengcn%2Fgo-colortext)
====================

This is a package to change the color of the text and background in the console, working both under Windows and other systems.

Under Windows, the console APIs are used. Otherwise, ANSI texts are output.

Docs: http://godoc.org/github.com/daviddengcn/go-colortext ([packages that import ct](http://go-search.org/view?id=github.com%2fdaviddengcn%2fgo-colortext))

Usage:
```go
ct.Foreground(Green, false)
fmt.Println("Green text starts here...")
ct.ChangeColor(Red, true, White, false)
fmt.Println(...)
ct.ResetColor()
```

To write to a file:
```go
colorFile, err := os.Create("colors.asc")
ct.Writer = colorFile
ct.Foreground(Red, false)
fmt.Fprint(colorFile, "R")
ct.Foreground(Green, false)
fmt.Fprint(colorFile, "G")
ct.Foreground(Blue, false)
fmt.Fprint(colorFile, "B")
ct.ResetColor()
```


LICENSE
=======
BSD/MIT license
