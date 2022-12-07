package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "alex",
		lastName:  "anderson",
		contact: contactInfo{
			email: "alex@gmail.com",
			zip:   12345}}
	alex.updatefname("alexy")
	alex.print()
}

func (p *person) updatefname(newfn string) {
	(*p).firstName = newfn
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
