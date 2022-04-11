package main

import s3_service "s3-copyobject/s3-service"

func main() {
	err := s3_service.CopyS3Object()
	if err != nil {
		panic(err)
	}
}
