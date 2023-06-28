package main

func Main(args map[string]interface{}) map[string]interface{} {
	//godotenv.Load()
	//bucketSecret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	//bucketSecretId := os.Getenv("AWS_ACCESS_KEY")
	////compressedVideoName := make(chan string)
	////thumbnail := make(chan string)
	//
	//// temp
	//fileName := "video.mp4"
	//tempBucket := "pullapptemp"
	////bucket := "pullappspaces"
	//
	////var wg sync.WaitGroup
	//
	////wg.Add(2)
	//
	//svc, s3Error := createS3Instance(createS3InstanceParams{
	//	secret: bucketSecret,
	//	keyId:  bucketSecretId,
	//})
	//if s3Error != nil {
	//	errorHandler(s3Error)
	//}
	//
	//err := downloadFromS3(downloadFromS3Params{
	//	fileName: fileName,
	//	bucket:   tempBucket,
	//	svc:      svc,
	//})
	//if err != nil {
	//	errorHandler(err)
	//}

	//go func() {
	//	defer wg.Done()
	//	handleVideo(fileName, bucket, svc, compressedVideoName)
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	handleThumbnail(fileName, bucket, svc, thumbnail)
	//}()
	//
	//defer wg.Wait()

	//produceMessage(message{
	//	CompressedFileName: <-compressedVideoName,
	//	ThumbnailName:      <-thumbnail,
	//	OriginalFileName:   fileName,
	//})

	msg := make(map[string]interface{})
	msg["body"] = "video compressed"
	return msg
}
