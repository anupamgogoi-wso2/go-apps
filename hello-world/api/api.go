package api

type Greet interface {
	SayHello() string
	SayBye() string
}

func HelloWorld(g Greet) string {
	return g.SayHello()
}
