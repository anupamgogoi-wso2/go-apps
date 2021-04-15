package types

type Employee struct {
	FirstName string
}

func (e Employee) SayHello() string {
	return "Hello " + e.FirstName
}
func (e Employee) SayBye() string {
	return "Hello " + e.FirstName
}
