package database

type Todo struct {
	Id      string `json:"id"`
	Todo    string `json:"todo"`
	Deleted bool   `json:"delete"`
}
