package core


type Usage struct {
	InputTokens *float64 `field:"required" json:"inputTokens" yaml:"inputTokens"`
	InputTokensDetails *InputTokensDetails `field:"required" json:"inputTokensDetails" yaml:"inputTokensDetails"`
	OutputTokens *float64 `field:"required" json:"outputTokens" yaml:"outputTokens"`
	OutputTokensDetails *OutputTokensDetails `field:"required" json:"outputTokensDetails" yaml:"outputTokensDetails"`
	TotalTokens *float64 `field:"required" json:"totalTokens" yaml:"totalTokens"`
}

