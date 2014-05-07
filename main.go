package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

const VERSION = "0.2.0"

type name struct {
	first string
	last  string
}

type Stringer interface {
	ToS() string
}

func main() {
	var opts struct {
		Name    string `short:"n" long:"name" description:"Name to greet" default:"world"`
		Version bool   `long:"version" description:"Display version and exit"`
	}

	parser := flags.NewParser(&opts, flags.Default)
	parser.Usage = "[options]"
	if _, err := parser.Parse(); err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing command line")
		os.Exit(1)
	}

	if opts.Version {
		fmt.Println("go-hello " + VERSION)
		os.Exit(0)
	}

	fmt.Printf("Hello, %s!\n", opts.Name)

	jack := name{first: "Jack", last: "Christensen"}

	fmt.Println(nameFullName(jack))

	fmt.Println(jack.fullName())

	fmt.Println(jack.pFullName())

	fmt.Println(nameToS(jack))

	puts(&jack)
}

func nameFullName(n name) string {
	return n.first + " " + n.last
}

func nameToS(n *name) string {
	return n.first + " " + n.last
}

func (n name) fullName() string {
	return n.first + " " + n.last
}

func (n *name) pFullName() string {
	return n.first + " " + n.last
}

func (n *name) ToS() string {
	return n.fullName()
}

func puts(v Stringer) {
	fmt.Println(v.ToS())
}
