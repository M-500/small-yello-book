package models

import "gorm.io/gorm"

type ImageModel struct {
	gorm.Model
	NoteId  uint   `gorm:"type:int;column:note_id;not null;comment:图片对应的NoteID"`         // NoteID
	OssUrl  string `gorm:"type:varchar(255);column:oss_url;not null;comment:OSS对应的URL"`  // OssUrl地址
	Width   int    `gorm:"type:int;column:width;not null;comment:图片宽度"`                  // 宽度
	Height  int    `gorm:"type:int;column:height;not null;comment:图片高度"`                 // 高度
	Size    int    `gorm:"type:float;column:size;not null;comment:图片大小 KB单位"`            // 大小
	Format  string `gorm:"type:varchar(32);column:format;not null;comment:图片格式"`         // 格式
	Hash    string `gorm:"type:char(32);column:img_hash;not null;comment:图片OSSURL对应的哈希"` // Hash
	OldPath string `gorm:"type:varchar(255);column:old_path;not null;comment:图片的绝对路径"`   // 原来的本地路径
}

func (ImageModel) TableName() string {
	return "images"
}
