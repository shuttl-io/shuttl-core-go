package core


type InputContent struct {
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	File *string `field:"optional" json:"file" yaml:"file"`
	// For file type, can include base64 content.
	FileData *FileAttachment `field:"optional" json:"fileData" yaml:"fileData"`
	Image *string `field:"optional" json:"image" yaml:"image"`
	Text *string `field:"optional" json:"text" yaml:"text"`
}

