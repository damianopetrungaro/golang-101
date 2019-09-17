package main

import "fmt"

type printer interface {
	print(s string)
}

type fmtPrinter struct {
}

func (p *fmtPrinter) print(s string) {
	fmt.Println(s)
}

func main() {
	p := &fmtPrinter{}
	printerFunc("Mario", p)
}

func printerFunc(s string, p printer) {
	p.print(s)
}
