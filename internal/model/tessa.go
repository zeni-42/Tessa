package model

type Clipboard struct {
	data	[]byte
	mime	string
}

func (c *Clipboard) Set(data []byte, mime string) {
    c.data = data
    c.mime = mime
}

func (c *Clipboard) Get() ([]byte, string) {
    return c.data, c.mime
}

func (c *Clipboard) Clear() {
    c.data = nil
    c.mime = ""
}