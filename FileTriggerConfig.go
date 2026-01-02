package core


type FileTriggerConfig struct {
	AllowedExtensions *[]*string `field:"required" json:"allowedExtensions" yaml:"allowedExtensions"`
	MaxFileSize *float64 `field:"required" json:"maxFileSize" yaml:"maxFileSize"`
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	UploadPath *string `field:"required" json:"uploadPath" yaml:"uploadPath"`
}

