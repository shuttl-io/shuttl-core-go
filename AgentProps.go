package core


type AgentProps struct {
	Model IModelFactory `field:"required" json:"model" yaml:"model"`
	Name *string `field:"required" json:"name" yaml:"name"`
	SystemPrompt *string `field:"required" json:"systemPrompt" yaml:"systemPrompt"`
	Toolkits *[]Toolkit `field:"required" json:"toolkits" yaml:"toolkits"`
	Outcomes *[]IOutcome `field:"optional" json:"outcomes" yaml:"outcomes"`
	Tools *[]ITool `field:"optional" json:"tools" yaml:"tools"`
	Triggers *[]ITrigger `field:"optional" json:"triggers" yaml:"triggers"`
}

