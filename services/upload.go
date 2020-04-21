package services

import (
	"archie/utils"
	"archie/utils/configer"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

type QiNiu struct {
	AK     string
	SK     string
	Bucket string `json:"bucket"`
}

// 读取 Qiniu AK/SK
func (qiniu *QiNiu) New() {
	config := configer.LoadQiNiuConfig()

	utils.CpStruct(&config, qiniu)
}

func (qiniu *QiNiu) GenToken() string {
	putPolicy := storage.PutPolicy{
		Scope: qiniu.Bucket,
	}
	mac := qbox.NewMac(qiniu.AK, qiniu.SK)

	return putPolicy.UploadToken(mac)
}

//func uploadByForm(key string) {
//	token := genToken()
//
//	config := storage.Config{
//		Zone:          &storage.ZoneHuanan,
//		UseHTTPS:      false,
//		UseCdnDomains: false,
//	}
//
//	formUploader := storage.NewFormUploader(&config)
//	ret := storage.PutRet{}
//
//	formUploader.PutFile(ctx.Background(), &ret, token, key)
//}
