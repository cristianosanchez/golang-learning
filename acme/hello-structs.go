package acme

import (
	"fmt"
	"log"
)

// Echo interface
type Echo interface {
	say() string
}

// Hello contains a single message
type Hello struct {
	msg string
}

// Goodbye contains also a single message
type Goodbye struct {
	Hello
}

// Polimorphyc method, any concrete implementation of Echo may be used.
func echo(e Echo) string {
	return e.say()
}

// COMPILER ERROR:
// we're implementing methods, so, since interface does not contain implementations,
// a receiver must be a pointer/value to a concrete type (struct).
//
// func (e *Echo) echo() string {
// 	return e.say()
// }

// callbackHello receives a string and a function f that receives a string and returns a string.
// the returning type of anotherEcho is a string.
func callbackHello(s string, f func(string) string) string {
	return f(s)
}

// say method, has a receiver of type Hello and returns a string.
// so Hello became also an Echo interface implementation.
func (h Hello) say() string {
	return "hello: " + h.msg + " " + h.msg
}

func (h Goodbye) say() string {
	return "goodbye: " + h.msg + " " + h.msg
}

func (h Hello) sayHelloValueReceiver() string {
	return "say hello value receiver: " + h.msg
}

func (h *Hello) sayHelloPointerReceiver() string {
	return "say hello pointer receiver: " + h.msg
}

func sayHelloPointerArgument(h *Hello) string {
	return "say hello pointer argument: " + h.msg
}

func sayHelloValueArgument(h Hello) string {
	return "say hello value argument: " + h.msg
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
}

func p(s string) {
	fmt.Println("\t", s)
}

// HowStructsWorks will play around declared structs
func HowStructsWorks() {
	fmt.Println("calling methods from value")
	h := Hello{"Cristiano"}
	p(h.sayHelloValueReceiver())
	p(h.sayHelloPointerReceiver())

	fmt.Println("calling methods from pointer")
	ph := &h
	p(ph.sayHelloValueReceiver())
	p(ph.sayHelloPointerReceiver())

	fmt.Println("calling echo from value")
	p(echo(h))

	fmt.Println("calling echo from pointer")
	p(echo(&h))

	// an inner function
	f := func(s string) string {
		return "inner message " + s
	}

	fmt.Println("calling from callback")
	p(callbackHello("hello", f))

	var h3 = Hello{"Cecilia"}
	p(echo(h3))

	h4 := Goodbye{Hello{"Marcelo"}}
	p(echo(h4))

	p(echo(h4.Hello))

}
