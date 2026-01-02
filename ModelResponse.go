package core


type ModelResponse struct {
	Data interface{} `field:"required" json:"data" yaml:"data"`
	EventName *string `field:"required" json:"eventName" yaml:"eventName"`
	ModelInstance IModel `field:"optional" json:"modelInstance" yaml:"modelInstance"`
	ThreadId *string `field:"optional" json:"threadId" yaml:"threadId"`
	Usage *Usage `field:"optional" json:"usage" yaml:"usage"`
}

