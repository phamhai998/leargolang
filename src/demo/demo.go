package main

import "fmt"

type Product interface {
	detail()
}

type Category interface {
	nameCtegory()
}

type ThuongDinh struct {
}

func (t ThuongDinh) detail() {
	fmt.Println("39-den")
}

func (c ThuongDinh) nameCtegory() {
	fmt.Println("thuongdinh")
}

func main() {
	t := ThuongDinh{}
	var p Product = t
	p.detail()
	var c Category = t
	c.nameCtegory()
}
