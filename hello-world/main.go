package main

import (
	"fmt"

	a "anupam.com/hello-world/api"
	t "anupam.com/hello-world/types"
)

func main() {
	e := t.Employee{FirstName: "Anupam"}
	c := t.Client{Location: "SP"}
	a.HelloWorld(c)
	fmt.Println(a.HelloWorld(e))
}
