package main

type printer interface {
	print(s string)
}

type fmtPrinter struct {
}

var _ printer = (*fmtPrinter)(nil)

// func (p *fmtPrinter) print(s string) {
// 	fmt.Println(s)
// }
