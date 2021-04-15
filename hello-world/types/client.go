package types

type Client struct {
	Location string
}

func (c Client) SayHello() string {
	return "Hello " + c.Location
}
func (c Client) SayBye() string {
	return "Hello " + c.Location
}
