package core


type ModelTextOutput struct {
	OutputType *string `field:"required" json:"outputType" yaml:"outputType"`
	Text *string `field:"required" json:"text" yaml:"text"`
}

