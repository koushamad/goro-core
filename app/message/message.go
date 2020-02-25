package message

type Message struct {
	Header      string
	ContentType string
	Body        []byte
}

const (
	TEXT = "text/plain"
	HTML = "text/html"
	JS   = "text/javascript"
	JSON = "application/json"
)

func (m Message) GetBody() string {
	return string(m.Body)
}
