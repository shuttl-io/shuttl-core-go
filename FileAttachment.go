package core


// Represents a file attachment sent with a chat message.
type FileAttachment struct {
	// Base64 encoded file content.
	Content *string `field:"required" json:"content" yaml:"content"`
	// The file name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// MIME type of the file.
	MimeType *string `field:"optional" json:"mimeType" yaml:"mimeType"`
	// The file path (optional, may not be relevant in all contexts).
	Path *string `field:"optional" json:"path" yaml:"path"`
}

