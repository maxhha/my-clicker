package ports

type LinkStorage interface {
	Create(redirect string) (string, error)
	GetRedirect(link string) (string, error)
	GetCounter(link string) (uint64, error)
	IncrementCounter(link string) error
}
