package Service

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
)

type s3Service struct {
	client *s3.S3
}

var instance *s3Service
var config = make(map[string]interface{})

func GetS3Service() *s3Service {
	if instance == nil {
		instance = new(s3Service)
	}

	return instance
}

func (this *s3Service) GetS3Client() (*s3.S3, error) {
	if this.client == nil {
		awsSession, err := session.NewSession(&aws.Config{
			Region: aws.String("us-east-1")},
		)
		if err != nil {
			return nil, err
		}
		// Create S3 service client
		s3Client := s3.New(awsSession)
		this.client = s3Client
	}

	return this.client, nil
}

func (this *s3Service) GetConfig(env string) (map[string]interface{}, error) {
	s3Client, err := this.GetS3Client()
	if err != nil {
		return nil, err
	}

	resp, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("sando-dbs"),
		Key:    aws.String(env + ".config.json"),
	})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	file, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}