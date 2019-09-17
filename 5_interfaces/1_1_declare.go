package main

import "fmt"

type printer interface {
	print(s string)
}

type reader interface {
	read(s string)
}

type printerReader interface {
	printer
	reader
}

type fmtPrinterReader struct {
	s string
}

func (p *fmtPrinterReader) print(s string) {
	fmt.Println(fmt.Sprintln(p.s, s))
}

func (p *fmtPrinterReader) read(s string) {
	p.s = fmt.Sprintln(p.s, s)
}

func main() {
	p := &fmtPrinterReader{}
	printerFunc("Super", "Mario", p)
}

func printerFunc(read, print string, p printerReader) {
	p.read(read)
	p.print(print)
}
