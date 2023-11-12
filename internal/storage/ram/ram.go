package ram

import "github.com/cactusfriday/go-url-shortener/internal/storage/base"

type RamStorage struct {
	StorageURL map[string]string
}

func (i *RamStorage) GetUrl(shortUrl string) (string, error) {
	return i.StorageURL[shortUrl], nil
}

func (i *RamStorage) PutUrl(short_url string, url string) error {
	// Check that key already present
	_, ok := i.StorageURL[short_url]
	if ok {
		return nil
	}

	i.StorageURL[short_url] = url
	return nil
}

// RAM Storage constructor
func NewRamStorage() base.IStorage {
	return &RamStorage{
		StorageURL: make(map[string]string),
	}
}
