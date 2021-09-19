package core

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	log "github.com/sirupsen/logrus"
	"os"
)

func UploadS3(key string, path string) error {
	cfg := GetConfig()
	session, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
		Credentials: credentials.NewEnvCredentials(),
	})

	log.WithFields(log.Fields{
		"path": path,
	}).Error("Inside Upload")
	file, err := os.Open(path)

	if err != nil {
		log.WithFields(log.Fields{
			"outfile": path,
		}).Error("Cannot Open File")
		return err
	}

	defer file.Close()

	uploader := s3manager.NewUploader(session)
	acl := "public-read"
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cfg.Bucket),
		Key:    aws.String(key),
		Body:   file,
		ACL:    &acl,
	})

	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"bucket": cfg.Bucket,
			"key":    key,
			"path":   path,
		}).Error("Unable To Upload File To S3")
		return err
	}
	return nil
}
