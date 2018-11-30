package diagram

type Message struct {
	From, To string
	Contents string
}

type Diagram struct {
	Actors   []string
	Messages []Message
}
