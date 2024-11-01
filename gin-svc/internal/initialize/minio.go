package initialize

import (
	"gin-svc/internal/conf/cfg"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)


func SetupMinio(c *cfg.MinioCfg)(*minio.Client,error){
	minioClient, err := minio.New(c.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.AccessKeyID, c.SecretAccessKey, ""),
		Secure: c.UseSSL,
	})
	if err != nil {
		return nil,err
	}
	return minioClient,nil
}


func MustSetupMinio(c *cfg.MinioCfg)*minio.Client {
	client,err:=SetupMinio(c)
	if err!=nil{
		panic(err)
	}
	return client
}

