package core


type ModelToolOutput struct {
	Arguments *map[string]interface{} `field:"required" json:"arguments" yaml:"arguments"`
	CallId *string `field:"required" json:"callId" yaml:"callId"`
	Name *string `field:"required" json:"name" yaml:"name"`
	OutputType *string `field:"required" json:"outputType" yaml:"outputType"`
}

