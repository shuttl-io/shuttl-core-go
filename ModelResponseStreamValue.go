package core


type ModelResponseStreamValue struct {
	Done *bool `field:"required" json:"done" yaml:"done"`
	Value *ModelResponse `field:"optional" json:"value" yaml:"value"`
}

