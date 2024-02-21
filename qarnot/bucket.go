package qarnot

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Bucket struct {
	Name         string
	CreationDate time.Time
}

type BucketObject struct {
	Name         string
	LastModified time.Time
	Size         int64
	ETag         string
}

// Create a new bucket
func (c *Client) CreateBucket(bucketName string) error {
	_, err := c.s3.CreateBucket(
		context.TODO(),
		&s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
		},
	)
	if err != nil {
		return fmt.Errorf("could not create bucket (%v) due to the following error : %v", bucketName, err)
	}

	return nil
}

// Delete a bucket
func (c *Client) DeleteBucket(bucketName string) error {
	_, err := c.s3.DeleteBucket(
		context.TODO(),
		&s3.DeleteBucketInput{
			Bucket: aws.String(bucketName),
		},
	)
	if err != nil {
		return fmt.Errorf("could not delete bucket (%v) due to the following error : %v", bucketName, err)
	}

	return nil
}

// List the buckets of the authenticated user
func (c *Client) ListBuckets() (*[]Bucket, error) {
	bucketsRaw, err := c.s3.ListBuckets(
		context.TODO(),
		&s3.ListBucketsInput{},
	)
	if err != nil {
		return nil, fmt.Errorf("could not list buckets due to the following error : %v", err)
	}

	var buckets []Bucket
	for _, bu := range bucketsRaw.Buckets {
		buckets = append(buckets, Bucket{Name: *bu.Name, CreationDate: *bu.CreationDate})
	}
	return &buckets, nil
}

// List objects inside of a bucket
func (c *Client) ListObjects(bucketName string) (*[]BucketObject, error) {
	objectsRaw, err := c.s3.ListObjectsV2(
		context.TODO(),
		&s3.ListObjectsV2Input{
			Bucket: &bucketName,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("could not list objects in bucket (%v) due to the following error : %v", bucketName, err)
	}

	var bucketObjects []BucketObject
	for _, obj := range objectsRaw.Contents {
		bucketObjects = append(
			bucketObjects,
			BucketObject{
				Name:         *obj.Key,
				LastModified: *obj.LastModified,
				Size:         *obj.Size,
				ETag:         *obj.ETag,
			})
	}

	return &bucketObjects, nil
}

// Input for uploading an object into a bucket
type ObjectToUpload struct {
	Bucket    string
	LocalPath string
	Key       string
}

// Upload objects in bucket
func (c *Client) UploadObject(object ObjectToUpload) error {
	body, err := os.Open(object.LocalPath)
	if err != nil {
		return fmt.Errorf("could not upload object to bucket due to the following error : %v", err)
	}

	_, err = c.s3.PutObject(
		context.TODO(),
		&s3.PutObjectInput{
			Bucket: &object.Bucket,
			Key:    &object.Key,
			Body:   body,
		},
	)
	if err != nil {
		return fmt.Errorf("could not upload object to bucket due to the following error : %v", err)
	}

	return nil
}

// Input for deleting an object in bucket
type ObjectToDelete struct {
	Bucket string
	Key    string
}

// Delete object in bucket
func (c *Client) DeleteObject(object ObjectToDelete) error {
	_, err := c.s3.DeleteObject(
		context.TODO(),
		&s3.DeleteObjectInput{
			Bucket: &object.Bucket,
			Key:    &object.Key,
		},
	)
	if err != nil {
		return fmt.Errorf("could not delete object in bucket due to the following error : %v", err)
	}

	return nil
}

// Input for getting object head in bucket
type ObjectToGetHead struct {
	Bucket string
	Key    string
}

// Represent object head
type ObjectHead struct {
	PartsCount    int32
	ETag          string
	ContentLength int64
	ContentType   string
}

// Get object head from bucket
func (c *Client) GetObjectHead(object ObjectToGetHead) (*ObjectHead, error) {
	head, err := c.s3.HeadObject(
		context.TODO(),
		&s3.HeadObjectInput{
			Bucket: &object.Bucket,
			Key:    &object.Key,
		},
	)
	if err != nil {
		return &ObjectHead{}, fmt.Errorf("could not get HEAD for the object in bucket due to the following error : %v", err)
	}

	objectHead := ObjectHead{
		PartsCount: func() int32 {
			if head.PartsCount != nil {
				return *head.PartsCount
			}
			return 0
		}(),
		ETag:          *head.ETag,
		ContentLength: *head.ContentLength,
		ContentType:   *head.ContentType,
	}

	return &objectHead, nil
}
