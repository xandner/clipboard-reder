package types

type Clipboard struct {
}

type ReturnClipboard struct {
	Id        int    `json:"id"`
	Datatype  string `json:"type"`
	Data      string `json:"data"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}