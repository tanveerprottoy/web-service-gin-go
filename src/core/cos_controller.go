package core

import (
	"fmt"
	"log"
	"mime/multipart"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
	"github.com/IBM/ibm-cos-sdk-go/service/s3/s3manager"
)

const (
	apiKey            = "2Z__hZEeqeULm1WgxJrK4qZMcRHdCCsUmYsvlrAk1n30"
	serviceInstanceID = "crn:v1:bluemix:public:cloud-object-storage:global:a/1d77be4257264338aaed4c8a3555a276:9322cea1-69a7-499c-8251-88b415d6a56c::"
	authEndpoint      = "https://iam.cloud.ibm.com/identity/token"
	serviceEndpoint   = "https://s3.us.cloud-object-storage.appdomain.cloud"
)

var (
	// Uploader declaration
	Uploader *s3manager.Uploader
	// Session declaration
	Session *session.Session
	// S3Client declaration
	S3Client *s3.S3
)

// InitCOS to use globally
func InitCOS() {
	conf := aws.NewConfig().
		WithRegion(BucketRegion).
		WithEndpoint(serviceEndpoint).
		WithCredentials(
			ibmiam.NewStaticCredentials(
				aws.NewConfig(),
				authEndpoint,
				apiKey,
				serviceInstanceID,
			),
		).WithS3ForcePathStyle(true)
	Session = session.Must(
		session.NewSession(),
	)
	// Create an uploader with the session and default options
	Uploader = s3manager.NewUploader(Session)
	// Create S3 service client
	S3Client = s3.New(Session, conf)
}

// GetBucket determines whether we have this bucket
func GetBucket(bucket string) error {
	// Do we have this Bucket?
	_, err := S3Client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}
	return nil
}

// CreateBucket creates a bucket
func CreateBucket(bucket string) error {
	// Create the S3 Bucket
	_, err := S3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}
	// Wait until bucket is created before finishing
	err = S3Client.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}
	return nil
}

// PutObject uploads (puts) object to COS
func PutObject(
	bucketName string,
	file multipart.File,
	fileName string,
) (string, error) {
	defer file.Close()
	// Upload the file's body to S3 bucket as an object with the key being the
	// same as the filename
	// Call Function to upload (Put) an object
	_, err := S3Client.PutObject(
		&s3.PutObjectInput{
			Bucket: aws.String(bucketName),

			// Can also use the `filepath` standard library package to modify the
			// filename as need for an S3 object key. Such as turning absolute path
			// to a relative path.
			Key: aws.String(fileName),

			// The file to be uploaded. io.ReadSeeker is preferred as the Uploader
			// will be able to optimize memory when uploading large content. io.Reader
			// is supported, but will require buffering of the reader's bytes for
			// each part.
			Body: file,
			ACL:  aws.String("public-read"),
		},
	)
	if err != nil {
		// Print the error and exit.
		log.Println("unable to upload")
		return "", err
	}
	// log.Println(res)
	// fmt.Printf("Successfully uploaded %q to %q\n", fileName, bucketName)
	// https://<endpoint>/<bucket>/<object>
	// or https://<bucket>.<endpoint>/<object>
	url := fmt.Sprintf(
		"%s/%s/%s",
		serviceEndpoint,
		bucketName,
		fileName,
	)
	return url, nil
}
