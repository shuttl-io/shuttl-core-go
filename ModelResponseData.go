package core


type ModelResponseData struct {
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	OutputText *ModelTextOutput `field:"optional" json:"outputText" yaml:"outputText"`
	OutputTextDelta *ModelDeltaOutput `field:"optional" json:"outputTextDelta" yaml:"outputTextDelta"`
	Requested interface{} `field:"optional" json:"requested" yaml:"requested"`
	ThreadId *string `field:"optional" json:"threadId" yaml:"threadId"`
	ToolCall *ModelToolOutput `field:"optional" json:"toolCall" yaml:"toolCall"`
}

