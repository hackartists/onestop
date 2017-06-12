package storage

import "bytes"

type OnestopStorageDataContext interface {
	FileName() string
	Bytes() []byte
	Reader() *bytes.Reader
	FileType() string
	Size() int64
	IsPublic() bool
}

type OnestopStorageContext interface {
	Region() string
	BucketName() string
	Path() string
	Session() interface{}
	Store(OnestopStorageDataContext) error
	Address(OnestopStorageDataContext) string
}
