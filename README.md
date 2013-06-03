go-colortext package
====================

This is a package to change the color of the text and background in the console, working both in windows and other systems.

Under windows, the console APIs are used, and otherwise ANSI text is used.

Docs: http://godoc.org/github.com/daviddengcn/go-colortext ([packages that import ct](http://godoc.org/github.com/daviddengcn/go-colortext?importers))

Usage:
```go
ChangeColor(Red, true, White, false)
ChangeColor(Green, false, None, false)
ResetColor()
```
