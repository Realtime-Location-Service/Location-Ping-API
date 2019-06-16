package model

// Queue ...
type Queue struct {
	Name        string
	Data        []byte
	ContentType string
	Durable     bool
	Exchange    string
}
