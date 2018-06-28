package interfaces
import (
	"fmt"
)

type Humaner interface {
	sayhi()
}

type Studnt struct {
	name string
	id int
}

type Teacher struct {
	addr string
	group string
}

func (std *Studnt) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi\n", std.name, std.id)	
}

func (tch *Teacher) sayhi() {
	fmt.Printf("Teacher[%s, %s] sayhi\n", tch.addr, tch.group)
}

type Mystr string

func (mystr *Mystr) sayhi() {
	fmt.Printf("Mystr[%s] sayhi\n", *mystr)
}

func WhoSayHi(i Humaner) {
	i.sayhi()
} 
func Interfaces() {
	fmt.Println("---> In Interfaces")
	var i Humaner

	s := &Studnt{"Sara", 32730}
	i = s
	i.sayhi()

	t := &Teacher{"Colazon", "Go"}
	i = t
	i.sayhi()

	var r Mystr = "How Long Times"
	i = &r
	i.sayhi()
}

func InterfacesNew() {
	fmt.Println("---> In InterfacesNew 01")

	s := &Studnt{"Sara", 32730}

	t := &Teacher{"Colazon", "Go"}

	var r Mystr = "How Long Times"

	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&r)

	//slice
	fmt.Println("---> In InterfacesNew 02")
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &r

	for _, i := range x {
		i.sayhi()
	}
}