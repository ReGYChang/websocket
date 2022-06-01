package storage

type Service interface {
	Close() error
	Save(interface{}) error
	//Load(string) (string, error)
	//LoadInfo(string) (*Item, error)
}

type Item struct {
	Id      uint64 `json:"id" redis:"id"`
	URL     string `json:"url" redis:"url"`
	Expires string `json:"expires" redis:"expires"`
	Visits  int    `json:"visits" redis:"visits"`
}

type ErrNoLink struct{}
