package core


type ToolArg struct {
	ArgType *string `field:"required" json:"argType" yaml:"argType"`
	Description *string `field:"required" json:"description" yaml:"description"`
	Required *bool `field:"required" json:"required" yaml:"required"`
	DefaultValue interface{} `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	EnumValues *[]*string `field:"optional" json:"enumValues" yaml:"enumValues"`
	Name *string `field:"optional" json:"name" yaml:"name"`
}

