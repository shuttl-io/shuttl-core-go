package core


type ToolkitProps struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Tools *[]ITool `field:"optional" json:"tools" yaml:"tools"`
}

