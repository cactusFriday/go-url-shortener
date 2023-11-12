package base

// Interface that every storage should implement
type IStorage interface {
	GetUrl(shortUrl string) (string, error)
	PutUrl(shortUrl, url string) error
}
