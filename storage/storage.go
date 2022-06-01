package storage

type Service interface {
	Close() error
	Save(interface{}) error
	Load() (string, error)
}

type ErrNoLink struct{}
