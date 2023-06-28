package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
)

type downloadFromS3Params struct {
	fileName string
	bucket   string
	svc      *s3.S3
}

type createS3InstanceParams struct {
	secret string
	keyId  string
}

func downloadFromS3(params downloadFromS3Params) error {
	file, err := os.Create(params.fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	resp, err := params.svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(params.bucket),
		Key:    aws.String(params.fileName),
	})
	if err != nil {
		return err
	}

	_, err = file.ReadFrom(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func createS3Instance(params createS3InstanceParams) (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Endpoint:    aws.String("ams3.digitaloceanspaces.com"),
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(params.keyId, params.secret, "")},
	)

	if err != nil {
		return nil, err
	}

	return s3.New(sess), nil
}

func errorHandler(err error) {
	log.Fatal(err)
	os.Exit(1)
}

func uploadToS3(fileName, bucket string, svc *s3.S3) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	params := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   file,
	}

	_, err = svc.PutObject(params)
	if err != nil {
		return err
	}

	return nil
}

func getVideoThumbnail(videoPath string, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ss", "00:00:03", "-vframes", "1", outputPath)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func createThumbnail(fileName string) (string, error) {
	thumbnailName := fileName[:len(fileName)-4] + "_thumbnail.jpg"

	thumbnailErr := getVideoThumbnail(fileName, thumbnailName)

	if thumbnailErr != nil {
		return "", thumbnailErr
	}

	return thumbnailName, nil
}

func handleThumbnail(fileName, bucket string, svc *s3.S3) (thumbnail string, err error) {
	thumbnail, err = createThumbnail(fileName)
	if err != nil {
		errorHandler(err)
	}

	thumbnailUploadErr := uploadToS3(thumbnail, bucket, svc)
	if thumbnailUploadErr != nil {
		errorHandler(thumbnailUploadErr)
	}

	return
}

func Main(args map[string]interface{}) map[string]interface{} {
	godotenv.Load()
	bucketSecret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	bucketSecretId := os.Getenv("AWS_ACCESS_KEY")
	//compressedVideoName := make(chan string)

	// temp
	fileName := "video.mp4"
	tempBucket := "pullappspaces"
	bucket := "pullappspaces"

	svc, s3Error := createS3Instance(createS3InstanceParams{
		secret: bucketSecret,
		keyId:  bucketSecretId,
	})
	if s3Error != nil {
		errorHandler(s3Error)
	}

	err := downloadFromS3(downloadFromS3Params{
		fileName: fileName,
		bucket:   tempBucket,
		svc:      svc,
	})
	if err != nil {
		errorHandler(err)
	}

	//go func() {
	//	defer wg.Done()
	//	handleVideo(fileName, bucket, svc, compressedVideoName)
	//}()
	//
	handleThumbnail(fileName, bucket, svc)

	//produceMessage(message{
	//	CompressedFileName: <-compressedVideoName,
	//	ThumbnailName:      <-thumbnail,
	//	OriginalFileName:   fileName,
	//})

	msg := make(map[string]interface{})
	msg["body"] = "video compressed"
	return msg
}
