package core


type ModelDeltaOutput struct {
	Delta *string `field:"required" json:"delta" yaml:"delta"`
	OutputType *string `field:"required" json:"outputType" yaml:"outputType"`
	SequenceNumber *float64 `field:"required" json:"sequenceNumber" yaml:"sequenceNumber"`
}

