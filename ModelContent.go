package core


type ModelContent struct {
	Content interface{} `field:"required" json:"content" yaml:"content"`
	Role *string `field:"required" json:"role" yaml:"role"`
}

