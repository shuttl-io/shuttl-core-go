package core


type ApiAuth struct {
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	Value *string `field:"required" json:"value" yaml:"value"`
}

