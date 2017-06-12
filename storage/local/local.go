package local

import (
	"os"

	"github.com/pwnartist/onestop/storage"
)

type LocalStorageContext struct {
	basePath string
	path     string
}

func New(base string, path string) (LocalStorageContext, error) {
	if err := makeDirectory(base + "/" + path); err != nil {
		return LocalStorageContext{}, err
	}

	return LocalStorageContext{basePath: base, path: path}, nil
}

func (l LocalStorageContext) BucketName() string {
	return l.basePath
}

func (l LocalStorageContext) Path() string {
	return l.path
}

func (l LocalStorageContext) Region() string {
	return ""
}

func (l LocalStorageContext) Session() interface{} {
	return ""
}

func (l LocalStorageContext) Store(data storage.OnestopStorageDataContext) error {
	f, err := os.Create(l.BucketName() + "/" + l.Path() + "/" + data.FileName())

	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data.Bytes())

	if err != nil {
		return err
	}

	return nil
}

func (l LocalStorageContext) Address(data storage.OnestopStorageDataContext) string {
	return "file://" + l.BucketName() + "/" + l.Path() + "/" + data.FileName()
}

func makeDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}

	return nil
}
