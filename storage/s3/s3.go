package s3

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	a "github.com/pwnartist/onestop/aws"
)

type S3StorageContext struct {
	sess   *session.Session
	region string
	bucket string
	path   string
}

func New(region string, bucket string, path string) (S3StorageContext, error) {
	sess, err := a.Connect(region)

	if err != nil {
		return S3StorageContext{}, storage.GroundStorageHandlers{}, err
	}

	return S3StorageContext{sess: sess, region: region, bucket: bucket, path: path}, nil
}

func (l S3StorageContext) BucketName() string {
	return l.bucket
}

func (l S3StorageContext) Path() string {
	return l.path
}

func (l S3StorageContext) Region() string {
	return l.region
}

func (l S3StorageContext) Session() interface{} {
	return l.sess
}

func (l S3StorageContext) Store(data storage.OnestopStorageDataContext) error {
	svc := s3.New(l.sess)

	params := &s3.PutObjectInput{
		Bucket:        aws.String(s.BucketName()),
		Key:           aws.String(s.Path() + "/" + data.FileName()),
		Body:          data.Reader(),
		ContentLength: aws.Int64(data.Size()),
		ContentType:   aws.String(data.FileType()),
	}

	if data.IsPublic() {
		params.ACL = aws.String("public-read")
	}

	_, err := svc.PutObject(params)
	return err
}

func (l S3StorageContext) Address(data storage.OnestopStorageDataContext) string {
	return "https://s3." + l.Region() + ".amazonaws.com" + "/" + l.BucketName() + "/" + l.Path() + "/" + data.FileName()
}

func UploadPublicFile(sess *session.Session, bucket string, path string, filename string, file *bytes.Reader, fileType string, size int64) error {
	svc := s3.New(sess)

	params := &s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(path + "/" + filename),
		Body:          file,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
		ACL:           aws.String("public-read"),
	}
	_, err := svc.PutObject(params)
	return err
}

func UploadPrivateFile(sess *session.Session, bucket string, path string, filename string, file *bytes.Reader, fileType string, size int64) error {
	svc := s3.New(sess)

	params := &s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(path + "/" + filename),
		Body:          file,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}
	_, err := svc.PutObject(params)
	return err
}

func ImageLink(region string, bucket string, path string, filename string) string {
	return "https://s3." + region + ".amazonaws.com" + "/" + bucket + path + "/" + filename
}
