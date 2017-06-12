package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"gitlab.artofthings.org/platform/ground/pkg/aws"
	"gitlab.artofthings.org/platform/ground/pkg/storage"
)

type S3StorageContext struct {
	sess   *session.Session
	region string
	bucket string
	path   string
}

func New(region string, bucket string, path string) (S3StorageContext, storage.GroundStorageHandlers, error) {
	sess, err := aws.Connect(region)

	if err != nil {
		return S3StorageContext{}, storage.GroundStorageHandlers{}, err
	}

	return S3StorageContext{sess: sess, region: region, bucket: bucket, path: path}, storage.GroundStorageHandlers{Store: Store, Address: Address}, nil
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
